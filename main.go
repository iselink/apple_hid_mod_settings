package main

import (
	"flag"
	"io/ioutil"
	"os"
)

const module_root = "/sys/module/hid_apple/parameters"
const fnParam = "/sys/module/hid_apple/parameters/fnmode"
const isoLayoutParam = "/sys/module/hid_apple/parameters/iso_layout"
const swapFnLeftCtrlParam = "/sys/module/hid_apple/parameters/swap_fn_leftctrl"
const swapOptionCommandParam = "/sys/module/hid_apple/parameters/swap_opt_cmd"

func main() {
	daemon := flag.Bool("daemon", false, "Run progam as daemon")
	fnKeyMode := flag.Int("fn", -1, "Set 1 to enable or 0 to disable fn keys.")

	flag.Parse()

	if *daemon {
		runAsDaemon()
		return
	}

	switch *fnKeyMode {
	case 0:
		{
			key, err := setFnKey(false)
			if err != nil {
				println(err.Error())
			}
			println(key)
			break
		}
	case 1:
		{
			key, err := setFnKey(true)
			if err != nil {
				println(err.Error())
			}
			println(key)
			break
		}

	}
}

func runAsDaemon() {

}

func setFnKey(value bool) (bool, error) {
	var data []byte
	if value {
		data = []byte("1")
	} else {
		data = []byte("0")
	}

	//get perms
	lstat, err := os.Lstat(fnParam)
	if err != nil {
		return false, err
	}

	//write
	err = ioutil.WriteFile(fnParam, data, lstat.Mode())
	if err != nil {
		return false, err
	}

	return true, nil
}
