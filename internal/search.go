package internal

import (
	"bufio"
	"errors"
	"log"
	"os/exec"
	"regexp"
	"strings"
)

const Gw2ExecutableName = "Gw2-64.exe"

func SearchForGuildWarsInstallation() (string, error) {

	disks, err := listDiscOfOs()

	if err != nil {
		return "", err
	}

	result := ""

	for _, disk := range disks {
		result = searchForGuildWars2ExecutableOnDisk(disk)

		if result != "" {
			return strings.TrimSpace(strings.ReplaceAll(result, "\\"+Gw2ExecutableName, "")), nil
		}
	}

	return "", errors.New("Couldn't find Guild Wars 2 Installation")
}

func listDiscOfOs() ([]string, error) {

	out, err := exec.Command("wmic", "logicaldisk", "get", "name").Output()

	if err != nil {
		return []string{}, err
	}

	regex := regexp.MustCompile(`(?m)^.:`)

	return regex.FindAllString(string(out), -1), nil
}

func searchForGuildWars2ExecutableOnDisk(disk string) string {
	sanitizedDiskString := disk + "\\"

	command := exec.Command("WHERE", "/R", sanitizedDiskString, Gw2ExecutableName)

	stdout, _ := command.StdoutPipe()

	err := command.Start()

	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(stdout)

	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		return scanner.Text()
	}

	err = command.Wait()

	if err != nil {
		log.Fatal(err)
	}

	return ""
}
