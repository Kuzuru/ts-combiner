package pkg

import (
	"fmt"
	"io/ioutil"
	"log"
)

func Combine(saveFolder, output string, inputs []string, isVerbose bool) error {
	// Set the input string for ffmpeg
	concatInput := "concat:"

	for _, input := range inputs {
		concatInput += input + "|"
	}

	// Remove the last "|"
	concatInput = concatInput[0 : len(concatInput)-1]

	// Create the ffmpeg command
	createCombineScript(saveFolder, concatInput, output)

	if isVerbose {
		fmt.Println("[LOG] Script created successfully")
	}

	return nil
}

func createCombineScript(name, input, output string) error {
	script := fmt.Sprintf("ffmpeg.exe -i %q -c copy %s", input, output)

	err := ioutil.WriteFile("./"+name+"/"+"combine.cmd", []byte(script), 0644)
	if err != nil {
		log.Fatalln("[ERR] Error writing to file:", err)
		return err
	}

	return nil
}
