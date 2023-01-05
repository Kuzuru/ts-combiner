package app

import (
	"fmt"
	"github.com/Kuzuru/ts-combiner/pkg"
	"github.com/alexflint/go-arg"
	"strconv"
)

var Args struct {
	Filename    string `arg:"-f, --filename, required" help:"file URL"`
	LastSegment int    `arg:"-l, --last, required" help:"last *.ts file (download from 1 to Last)"`
	SaveFolder  string `arg:"-s, --save, required" help:"folder to save files"`
	Verbose     bool   `arg:"-v, --verbose" help:"verbosity level"`
}

func Main() error {
	// Parsing command-line arguments
	arg.MustParse(&Args)

	if Args.Verbose {
		fmt.Printf("[LOG] Filename: %s\n", Args.Filename)
		fmt.Printf("[LOG] DL URL: %s*.ts (* <- from 1 to %d)\n", Args.Filename, Args.LastSegment)
	}

	for i := 1; i <= Args.LastSegment; i++ {
		pkg.Download(Args.Filename, Args.SaveFolder, i, Args.Verbose)
	}

	inputs := make([]string, 0)
	for i := 1; i <= Args.LastSegment; i++ {
		inputs = append(inputs, strconv.Itoa(i)+".ts")
	}

	if Args.Verbose {
		fmt.Println("[LOG] Done downloading!")
		fmt.Printf("[LOG] Combining %d segments...", Args.LastSegment)
	}

	pkg.Combine(Args.SaveFolder, inputs, Args.Verbose)

	fmt.Println("[LOG] Done! You can now launch combine.cmd")

	return nil
}
