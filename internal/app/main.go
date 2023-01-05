package app

import (
	"fmt"
	"github.com/Kuzuru/ts-combiner/pkg"
	"github.com/alexflint/go-arg"
	"log"
)

var Args struct {
	Filename    string `arg:"-f, --filename, required" help:"file URL"`
	LastSegment int    `arg:"-l, --last, required" help:"last *.ts file (download from 1 to Last)"`
	Verbose     bool   `arg:"-v, --verbose" help:"verbosity level"`
}

func Main() error {
	// Parsing command-line arguments
	arg.MustParse(&Args)

	fmt.Printf("Filename: %s\n", Args.Filename)
	fmt.Printf("DL URL: %s*.ts (* <- from 1 to %d)\n", Args.Filename, Args.LastSegment)

	err := pkg.Download(Args.Filename, 1)
	if err != nil {
		log.Fatalln("error while downloading file:", err)
		return err
	}

	return nil
}
