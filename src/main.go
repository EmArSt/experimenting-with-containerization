package main

import (
	"fmt"
	"os"
	"os/exec"
	"syscall"
	//"syscall"
)

func main() {
	switch os.Args[1] {
	case "run":
		parent()
	case "child":
		child()
	default:
		panic("Invalid Argument: should be 'run' or 'child'")
	}
}

func parent() {
	// '...' is go's spread operator to give elements of a slice seperately in a variadic function

	//	{ //pull and extract an image from dockerhub
	//		cmd := exec.Command("docker", fmt.Sprintf("export $(docker create %s) -o rootfs.tar.gz", os.Args[2]))
	//		cmd.Stdin = os.Stdin
	//		cmd.Stdout = os.Stdout
	//		cmd.Stderr = os.Stderr
	//
	//		if err := cmd.Run(); err != nil {
	//			fmt.Println("Error", err)
	//			os.Exit(1)
	//		}
	//	}

	cmd := exec.Command("/proc/self/exe", append([]string{"child"}, os.Args[2:]...)...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	cmd.SysProcAttr = &syscall.SysProcAttr{
		Cloneflags: syscall.CLONE_NEWUTS | syscall.CLONE_NEWPID | syscall.CLONE_NEWNS,
	}

	if err := cmd.Run(); err != nil {
		fmt.Println("Error", err)
		os.Exit(1)
	}
}

func child() {

	//must(syscall.Mount("rootfs", "rootfs", "tmpfs", syscall.MS_BIND|syscall.MS_REC, ""))
	//must(os.MkdirAll("rootfs/oldrootfs", 0700))
	//must(syscall.PivotRoot("./rootfs", "./rootfs/oldrootfs"))
	must(syscall.Chroot("rootfs"))
	must(os.Chdir("/"))

	must(syscall.Sethostname([]byte("Container")))

	cmd := exec.Command(os.Args[2], os.Args[3:]...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		fmt.Println("Error", err)
		os.Exit(1)
	}

}

func must(err error) {
	if err != nil {
		panic(err)
	}
}
