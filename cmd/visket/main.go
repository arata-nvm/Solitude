package main

import (
	"flag"
	"fmt"
	"github.com/arata-nvm/visket/cmd/visket/build"
	"log"
	"os"
)

const VERSION = "0.0.1"

func main() {
	var (
		//isDebug  = flag.Bool("v", false, "Emit debug information")
		optimize = flag.Bool("O", false, "Enable optimization")
		output   = flag.String("o", "", "Write output to <filename>")
		emitLLVM = flag.Bool("emit-llvm", false, "Generate output in LLVM formats")
	)
	flag.Parse()

	filename := flag.Arg(0)
	if filename == "" {
		fmt.Printf("visket %s\n", VERSION)
		fmt.Println("Usage: visket [options] <filename>")
		os.Exit(1)
	}

	fmt.Printf("Compiling %s\n", filename)

	if *emitLLVM {
		err := build.EmitLLVM(filename, *output, *optimize)
		if err != nil {
			log.Fatal(err)
		}
	} else {
		err := build.Build(filename, *output, *optimize)
		if err != nil {
			log.Fatal(err)
		}
	}

	fmt.Println("Finished")
}
