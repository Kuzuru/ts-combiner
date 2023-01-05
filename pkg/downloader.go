package pkg

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
)

func Download(url string, segmentNum int) error {
	fmt.Println("Downloading", url)

	out, err := os.Create(strconv.Itoa(segmentNum) + ".ts")
	if err != nil {
		fmt.Println("Error creating file:", err)
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
		fmt.Println("Error downloading:", err)
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
		fmt.Println("Error writing to file:", err)
		return err
	}

	return nil
}
