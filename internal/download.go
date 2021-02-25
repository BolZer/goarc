package internal

import (
	"io"
	"net/http"
	"os"
)

const ARC_DPS_DOWNLOAD_LINK = "https://www.deltaconnected.com/arcdps/x64/d3d9.dll"

func DownloadArcDPStoDestinationPath(path string) error {
	resp, err := http.Get(ARC_DPS_DOWNLOAD_LINK)

	if err != nil {
		return err
	}

	defer resp.Body.Close()

	out, err := os.Create(path)

	if err != nil {
		return err
	}

	defer out.Close()

	_, err = io.Copy(out, resp.Body)

	if err != nil {
		return err
	}

	return nil
}