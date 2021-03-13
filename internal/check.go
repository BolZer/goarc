package internal

import (
	"bytes"
	"crypto/md5"
	"io"
	"log"
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
	var remoteArcDpsMd5Checksum []byte
	var destinationArcDpsMd5Checksum []byte

	resp, err := http.Get(ArcDpsDownloadLink)

	if err != nil {
		return false, err
	}

	defer func() {
		err := resp.Body.Close()

		if err != nil {
			log.Fatal(err)
		}
	}()

	body, err := io.ReadAll(resp.Body)

	remoteArcDpsMd5 := md5.New()

	if _, err := io.Copy(remoteArcDpsMd5, bytes.NewReader(body)); err != nil {
		return false, err
	}

	remoteArcDpsMd5Checksum = remoteArcDpsMd5.Sum(nil)

	if err != nil {
		return false, err
	}

	existingArcDpsFile, err := os.Open(path)

	if err != nil {
		return false, err
	}

	defer func() {
		err := existingArcDpsFile.Close()

		if err != nil {
			log.Fatal(err)
		}
	}()

	existingArcDpsMd5 := md5.New()

	if _, err := io.Copy(existingArcDpsMd5, existingArcDpsFile); err != nil {
		return false, err
	}

	destinationArcDpsMd5Checksum = existingArcDpsMd5.Sum(nil)

	compareResult := bytes.Compare(remoteArcDpsMd5Checksum, destinationArcDpsMd5Checksum)

	if compareResult == 0 {
		return false, nil
	}

	return true, nil
}
