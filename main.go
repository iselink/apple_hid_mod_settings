package main

import (
	"errors"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"syscall"
)

const module_root = "/sys/module/hid_apple/parameters"
const fnParam = "/sys/module/hid_apple/parameters/fnmode"
const isoLayoutParam = "/sys/module/hid_apple/parameters/iso_layout"
const swapFnLeftCtrlParam = "/sys/module/hid_apple/parameters/swap_fn_leftctrl"
const swapOptionCommandParam = "/sys/module/hid_apple/parameters/swap_opt_cmd"

func main() {
	//daemon := flag.Bool("daemon", false, "Run progam as daemon")
	//daemonFrontend := flag.Bool("daemon-ui", false, "Run progam as daemon in tray (show nice clickable checkboxes).") //TODO: make UI frontend
	skipAccessCheck := flag.Bool("skip-check", false, "Allow to skip access (permission) check.")
	fnKeyMode := flag.Int("fn", -1, "Set 1 to enable or 0 to disable fn keys.")
	isoLayoutMode := flag.Int("iso", -1, "Set 1 to enable or 0 to disable ISO layout.")
	swapFnLeftCtrlMode := flag.Int("swap-fn-lctrl", -1, "Set 1 to enable or 0 to disable swapping Fn and left Ctrl.")
	swapOptCmdMode := flag.Int("swap-opt-cmd", -1, "Set 1 to enable or 0 to disable swapping Opt and Cmd keys..")

	flag.Parse()

	if !*skipAccessCheck {
		accessAvailable, err := validateAccess()
		if err != nil {
			println(err.Error())
			return
		}

		if !accessAvailable {
			println("It looks like program doesn't have access to module parameters.")
			return
		}
	}

	setValues(fnKeyMode, isoLayoutMode, swapFnLeftCtrlMode, swapOptCmdMode)

}

func validateAccess() (bool, error) {
	dir, err := os.ReadDir(module_root)
	if err != nil {
		return false, err
	}

	if len(dir) == 0 {
		return false, errors.New("fs is empty in module directory")
	}

	for _, s := range []string{fnParam, isoLayoutParam, swapOptionCommandParam, swapFnLeftCtrlParam} {
		info, err := os.Stat(s)
		if err != nil {
			return false, err
		}

		stat := info.Sys().(*syscall.Stat_t)

		if uint32(os.Geteuid()) == stat.Uid && stat.Mode&syscall.S_IWUSR != 0 {
			//owner have write access
			continue
		} else if uint32(os.Getegid()) == stat.Gid && stat.Mode&syscall.S_IWGRP != 0 {
			//group have write access
			continue
		} else if stat.Mode&syscall.S_IWOTH != 0 {
			//everyone else have access, wtf?
			continue
		} else {
			return false, errors.New(fmt.Sprint("For ", s, " program doesn't have access."))
		}

	}
	return true, nil
}

func setValues(fnKeyMode *int, isoLayoutMode *int, swapFnLeftCtrlMode *int, swapOptCmdMode *int) {
	if *fnKeyMode != -1 {
		key, err := writeParam(fnParam, !(*fnKeyMode == 0))
		if err != nil {
			println(err.Error())
		}
		println(key)
	}

	if *isoLayoutMode != -1 {
		key, err := writeParam(isoLayoutParam, !(*isoLayoutMode == 0))
		if err != nil {
			println(err.Error())
		}
		println(key)
	}

	if *swapFnLeftCtrlMode != -1 {
		key, err := writeParam(swapFnLeftCtrlParam, !(*swapFnLeftCtrlMode == 0))
		if err != nil {
			println(err.Error())
		}
		println(key)
	}

	if *swapOptCmdMode != -1 {
		key, err := writeParam(swapOptionCommandParam, !(*swapOptCmdMode == 0))
		if err != nil {
			println(err.Error())
		}
		println(key)
	}

}

func writeParam(path string, value bool) (bool, error) {
	var data []byte
	if value {
		data = []byte("1")
	} else {
		data = []byte("0")
	}

	//get perms
	lstat, err := os.Lstat(path)
	if err != nil {
		return false, err
	}

	//write
	err = ioutil.WriteFile(path, data, lstat.Mode())
	if err != nil {
		return false, err
	}

	println(path, "=", string(data))

	return true, nil
}
