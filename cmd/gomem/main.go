package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/kamisari/gomem"
)

type option struct {
	version     bool
	interactive bool
	workdir     string

	// TODO: impl subcmd for run
	subcmd  string
	subarsg []string
}

var opt option

const version = "0.0"

func (opt *option) init() {
	flag.BoolVar(&opt.version, "version", false, "")
	flag.BoolVar(&opt.version, "v", false, "")

	flag.BoolVar(&opt.interactive, "interactive", false, "")
	flag.BoolVar(&opt.interactive, "i", false, "")

	flag.StringVar(&opt.workdir, "workdir", "", "")
	flag.StringVar(&opt.workdir, "w", "", "")
	flag.Parse()
	if flag.NArg() != 0 {
		// TODO: impl parse subcmd
		log.Fatalf("invalid args: %q", flag.Args())
	}
	if opt.version {
		fmt.Fprintf(os.Stderr, "version %s", version)
	}
	var err error
	opt.workdir, err = filepath.Abs(opt.workdir)
	if err != nil {
		log.Fatal(err)
	}
	if err := os.Chdir(opt.workdir); err != nil {
		log.Fatal(err)
	}
}

func init() {
	log.SetPrefix("gomem:")
	log.SetFlags(log.Lshortfile)

	opt.init()
}

func main() {
	gs, err := gomem.GomemsNew()
	if err != nil {
		log.Fatal(err)
	}

	if opt.interactive {
		if err := interactive(os.Stdin, os.Stdout, "gomem:>", gs); err != nil {
			log.Fatal(err)
		}
		return
	}
	if err := run(os.Stdout, gs); err != nil {
		log.Fatal(err)
	}
}
