package internal

import (
	"io"
	"log"
	"net/http"
	"os"
)

const ArcDpsDownloadLink = "https://www.deltaconnected.com/arcdps/x64/d3d9.dll"

func DownloadRemoteArcDPSToDestinationPath(path string) error {
	resp, err := http.Get(ArcDpsDownloadLink)

	if err != nil {
		return err
	}

	defer func() {
		err := resp.Body.Close()

		if err != nil {
			log.Fatal(err)
		}
	}()

	out, err := os.Create(path)

	if err != nil {
		return err
	}

	defer func() {
		err := out.Close()

		if err != nil {
			log.Fatal(err)
		}
	}()

	_, err = io.Copy(out, resp.Body)

	if err != nil {
		return err
	}

	return nil
}
