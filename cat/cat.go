//
// cat.go (go-coreutils) 0.1
// Copyright (C) 2014, The GO-Coreutils Developers.
//
// Written By: Abram C. Isola
//
package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"net"
	"os"
	"strings"
)

const (
	help_text string = `
    Usage: cat [OPTIONS] [FILE]...

    Concatenate and print the content of files.

    With no FILE, or when FILE is -, read standard input.

        -b, --number-nonblank    number nonempty output lines, overrides -n
        -E, --show-ends          display $ at end of each line
        -n, --number             number all output lines
        -s, --squeeze-blank      suppress repeated empty output lines
        -T, --show-tabs          display TAB characters as ^I
        -u                       (ignored)

        --help        display this help and exit
        --version     output version information and exit
    `
	version_text = `
    cat (go-coreutils) 0.1

    Copyright (C) 2014, The GO-Coreutils Developers.
    This program comes with ABSOLUTELY NO WARRANTY; for details see
    LICENSE. This is free software, and you are welcome to redistribute
    it under certain conditions in LICENSE.
`
)

var (
	countNonBlank         = flag.Bool("b", false, "Number the non-blank output lines, starting at 1.")
	countNonBlankLong     = flag.Bool("number-nonblank", false, "Number the non-blank output lines, starting at 1.")
	ignoredU              = flag.Bool("u", false, "(ignored)")
	numberOutput          = flag.Bool("n", false, "Number the output lines, starting at 1.")
	showEnds              = flag.Bool("E", false, "Display $ at end of each line.")
	showEndsLong          = flag.Bool("show-ends", false, "Display $ at end of each line.")
	showEndsNonPrint      = flag.Bool("e", false, "Equivalent to -vE")
	showTabs              = flag.Bool("T", false, "Display TAB characters as ^I")
	showTabsLong          = flag.Bool("show-tabs", false, "Display TAB characters as ^I")
	squeezeEmptyLines     = flag.Bool("s", false, "Squeeze multiple adjacent empty lines, causing the output to be single spaced.")
	squeezeEmptyLinesLong = flag.Bool("squeeze-blank", false, "Squeeze multiple adjacent empty lines, causing the output to be single spaced.")
)

func openFile(s string) (io.ReadWriteCloser, error) {
	fi, err := os.Stat(s)
	if err != nil {
		return nil, err
	}
	if fi.Mode()&os.ModeSocket != 0 {
		return net.Dial("unix", s)
	}
	return os.Open(s)
}

func dumpLines(w io.Writer, r io.Reader) (n int64, err error) {
	var lastline, line string
	br := bufio.NewReader(r)
	nr := 0
	for {
		line, err = br.ReadString('\n')
		if err != nil {
			return
		}
		if *squeezeEmptyLines && lastline == "\n" && line == "\n" {
			continue
		}
		if *showEnds {
			line = strings.ReplaceAll(line, "\n", "$\n")
		}
		if *showTabs {
			line = strings.ReplaceAll(line, "\t", "^I")
		}
		if *countNonBlank && line == "\n" || line == "" {
			fmt.Fprint(w, line)
		} else if *countNonBlank || *numberOutput {
			nr++
			fmt.Fprintf(w, "%6d\t%s", nr, line)
		} else {
			fmt.Fprint(w, line)
		}
		lastline = line
	}
	return
}

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

	// process the long versions of the flags
	if *countNonBlankLong {
		*countNonBlank = true
	}
	if *showEndsLong {
		*showEnds = true
	}
	if *showEndsNonPrint {
		*showEnds = true
		// FIXME enable show-non-printing, ie -v
	}
	if *showTabsLong {
		*showTabs = true
	}
	if *squeezeEmptyLinesLong {
		*squeezeEmptyLines = true
	}

	// -b overrides -n
	if *countNonBlank {
		*numberOutput = false
	}

	catenatingFunc := io.Copy
	// some flags cause the input data to be modified before being output
	if *countNonBlank || *numberOutput || *squeezeEmptyLines || *showEnds || *showTabs {
		catenatingFunc = dumpLines
	}

	fileList := flag.Args()
	// read from stdin if no files were specified
	if 0 == len(fileList) {
		fileList = append(fileList, "-")
	}

	for _, fname := range fileList {
		if fname == "-" {
			catenatingFunc(os.Stdout, os.Stdin)
		} else {
			f, err := openFile(fname)
			if err != nil {
				fmt.Fprintln(os.Stderr, err)
				continue
			}
			catenatingFunc(os.Stdout, f)
			f.Close()
		}
	}
}
