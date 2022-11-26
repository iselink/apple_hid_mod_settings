package main

import (
	"flag"
	"io/ioutil"
	"log"
	"net"
	"os"
	"sync"
)

const module_root = "/sys/module/hid_apple/parameters"
const fnParam = "/sys/module/hid_apple/parameters/fnmode"
const isoLayoutParam = "/sys/module/hid_apple/parameters/iso_layout"
const swapFnLeftCtrlParam = "/sys/module/hid_apple/parameters/swap_fn_leftctrl"
const swapOptionCommandParam = "/sys/module/hid_apple/parameters/swap_opt_cmd"
const sockPath = "/tmp/hidDaemon.sock"

func main() {
	daemon := flag.Bool("daemon", false, "Run progam as daemon")
	daemonFrontend := flag.Bool("daemon-ui", false, "Run progam as daemon in tray (show nice clickable checkboxes).") //TODO: make UI frontend
	fnKeyMode := flag.Int("fn", -1, "Set 1 to enable or 0 to disable fn keys.")
	isoLayoutMode := flag.Int("iso", -1, "Set 1 to enable or 0 to disable ISO layout.")
	swapFnLeftCtrlMode := flag.Int("fn", -1, "Set 1 to enable or 0 to disable swapping Fn and left Ctrl.")
	swapOptCmdMode := flag.Int("fn", -1, "Set 1 to enable or 0 to disable swapping Opt and Cmd keys..")

	flag.Parse()

	if *daemon {
		runAsDaemon()
	} else if *daemonFrontend {
		runAsUiFrontend()
	} else {
		setValues(fnKeyMode, isoLayoutMode, swapFnLeftCtrlMode, swapOptCmdMode)
	}

}

func setValues(fnKeyMode *int, isoLayoutMode *int, swapFnLeftCtrlMode *int, swapOptCmdMode *int) {
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

	//TODO: complete func with all parameters
}

func runAsUiFrontend() {
	wait := sync.WaitGroup{}
	wait.Add(1)

	wait.Wait()

}

func runAsDaemon() {
	if err := os.Remove(sockPath); err != nil {
		log.Fatal(err)
	}

	l, err := net.Listen("unix", sockPath)
	if err != nil {
		log.Fatal("listen error:", err)
	}
	defer l.Close()

	for {
		// Accept new connections, dispatching them to echoServer
		// in a goroutine.
		conn, err := l.Accept()
		if err != nil {
			log.Fatal("accept error:", err)
		} else {
			go func(con net.Conn) {
				//TODO: complete
			}(conn)
		}
	}
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
