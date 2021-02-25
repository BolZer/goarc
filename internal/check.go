package internal

import (
	"bytes"
	"crypto/md5"
	"io"
	"net/http"
	"os"
)

func CheckIfArcDPSExists(path string) bool {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return false
	}

	return true
}

func CheckIfArcDPSIsOutdated(path string) (bool, error) {

	var existingArcDpsMd5Checksum []byte

	resp, err := http.Get(ARC_DPS_DOWNLOAD_LINK)

	if err != nil {
		return false, err
	}

	defer resp.Body.Close()

	contentOfFetchArcDpsMd5 := md5.New()

	if _, err := io.Copy(contentOfFetchArcDpsMd5, resp.Body); err != nil {
		return false, err
	}

	fetchArcDpsMd5Checksum := contentOfFetchArcDpsMd5.Sum(nil)

	existingArcDpsFile, err := os.Open(path)

	if err != nil {
		return false, err
	}

	defer existingArcDpsFile.Close()

	existingArcDpsMd5 := md5.New()

	if _, err := io.Copy(existingArcDpsMd5, existingArcDpsFile); err != nil {
		return false, err
	}

	existingArcDpsMd5Checksum = existingArcDpsMd5.Sum(nil)

	compareResult := bytes.Compare(fetchArcDpsMd5Checksum, existingArcDpsMd5Checksum)

	if compareResult == 0 {
		return false, nil
	}	

	return true, nil
}