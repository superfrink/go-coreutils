//
// arch.go (go-coreutils) 0.1
// Copyright (C) 2014, The GO-Coreutils Developers.
//
// Written By: Abram C. Isola
//
package main

import "flag"
import "fmt"
import "os"
import "syscall"

const (
	help_text string = `
    Usage: arch
    
    print the system architecture

        --help        display this help and exit
        --version     output version information and exit
    `
	version_text = `
    arch (go-coreutils) 0.1

    Copyright (C) 2014, The GO-Coreutils Developers.
    This program comes with ABSOLUTELY NO WARRANTY; for details see
    LICENSE. This is free software, and you are welcome to redistribute 
    it under certain conditions in LICENSE.
`
)

func main() {
	help := flag.Bool("help", false, help_text)
	version := flag.Bool("version", false, version_text)
	flag.Parse()

	if *help {
		fmt.Println(help_text)
		os.Exit(0)
	}

	if *version {
		fmt.Println(version_text)
		os.Exit(0)
	}

	var utsname syscall.Utsname
	syscall.Uname(&utsname)

	// is there a cleaner way to find the null bytes?
	n := 0
	found: for ; n < len(utsname.Machine); n++ {
		if 0 == utsname.Machine[n] {
			break found
		}
	}
	fmt.Printf("%s\n", string(utsname.Machine[:n]))
}
