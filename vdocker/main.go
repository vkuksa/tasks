package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"syscall"
)

const (
	appName = "vdocker"
	rootfs  = "/home/waterfall/" + appName
	helpMsg = `vdocker <command> <attributes>
	Avaliable commands:
		run
		fork`
)

func main() {
	switch os.Args[1] {
	case "run":
		run()
	case "fork":
		fork()
	default:
		fmt.Print(helpMsg)
	}
}

func run() {
	fmt.Printf("Running %v \n", os.Args[2:])

	cmd := exec.Command("/proc/self/exe", append([]string{"fork"}, os.Args[2:]...)...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.SysProcAttr = &syscall.SysProcAttr{
		Cloneflags: syscall.CLONE_NEWUTS | syscall.CLONE_NEWPID, // | syscall.CLONE_NEWNS,
	}

	must(cmd.Run())
}

func fork() {
	fmt.Printf("Forked %v \n", os.Args[2:])

	prepareCGroup()

	cmd := exec.Command(os.Args[2], os.Args[3:]...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	must(syscall.Sethostname([]byte(appName)))
	must(syscall.Chroot(rootfs))
	must(os.Chdir("/"))
	must(syscall.Mount("proc", "proc", "proc", 0, ""))
	defer syscall.Unmount("proc", 0)

	os.Mkdir("/mytemp", 0755)
	must(syscall.Mount("mytemp", "mytemp", "tmpfs", 0, ""))
	defer syscall.Unmount("mytemp", 0)

	must(cmd.Run())
}

func prepareCGroup() {
	cgroups := "/sys/fs/cgroup/"
	appRoot := filepath.Join(filepath.Join(cgroups, "pids.max"), appName)
	os.Mkdir(appRoot, 0755)
	// Cgroup settings here
	must(ioutil.WriteFile(filepath.Join(appRoot, "/notify_on_release"), []byte("1"), 0700)) // https://www.man7.org/linux/man-pages/man7/cgroups.7.html
	// Move our pid into control group
	must(ioutil.WriteFile(filepath.Join(appRoot, "/cgroups.procs"), []byte(strconv.Itoa(os.Getpid())), 0700)) // https://www.man7.org/linux/man-pages/man7/cgroups.7.html
	// ...
}

func must(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
