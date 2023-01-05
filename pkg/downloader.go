package pkg

import (
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
)

func Download(url, saveFolder string, segmentNum int, isVerbose bool, done chan bool) error {
	url += strconv.Itoa(segmentNum) + ".ts"

	if isVerbose {
		fmt.Println("[LOG] Downloading", url)
	}

	// If save folder does not exist, create it
	if _, err := os.Stat(saveFolder); errors.Is(err, os.ErrNotExist) {
		err := os.Mkdir(saveFolder, os.ModePerm)
		if err != nil {
			log.Fatalln("[ERR] Error creating save folder:", err)
		}
	}

	// TODO: Make smth better than this shit
	out, err := os.Create("./" + saveFolder + "/" + strconv.Itoa(segmentNum) + ".ts")
	if err != nil {
		log.Fatalln("[ERR] Error creating file:", err)
		return err
	}

	defer func(out *os.File) {
		err := out.Close()
		if err != nil {
			return
		}
	}(out)

	resp, err := http.Get(url)
	if err != nil {
		log.Fatalln("[ERR] Error downloading:", err)
		return err
	}

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			return
		}
	}(resp.Body)

	_, err = io.Copy(out, resp.Body)
	if err != nil {
		log.Fatalln("[ERR] Error writing to file:", err)
		return err
	}

	done <- true

	return nil
}
