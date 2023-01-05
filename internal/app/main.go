package app

import (
	"fmt"
	"github.com/Kuzuru/ts-combiner/pkg"
	"github.com/alexflint/go-arg"
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

	done := make(chan bool, Args.LastSegment)

	for i := 1; i <= Args.LastSegment; i++ {
		go pkg.Download(Args.Filename, Args.SaveFolder, i, Args.Verbose, done)
	}

	for i := 1; i <= Args.LastSegment; i++ {
		<-done
	}

	return nil
}
