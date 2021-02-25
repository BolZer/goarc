package internal

import (
	"bufio"
	"errors"
	"os/exec"
	"regexp"
	"strings"
)

const GW2_EXECUTABLE_NAME = "Gw2-64.exe"

func SearchForGuildWarsInstallation() (string, error) {

	disks, err := listDiscOfOs()

	if err != nil {
		return "", err
	}

	result := ""

	for _, disk := range disks {
		result = searchForGuildWars2ExecutableOnDisk(disk)

		if result != ""{
			return strings.TrimSpace(strings.ReplaceAll(result, "\\" +  GW2_EXECUTABLE_NAME, "")), nil
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

	command := exec.Command("WHERE", "/R", sanitizedDiskString, GW2_EXECUTABLE_NAME)

    stdout, _ := command.StdoutPipe()

    command.Start()

	scanner := bufio.NewScanner(stdout)

    scanner.Split(bufio.ScanLines)

    for scanner.Scan() {
        return scanner.Text()
    }

	command.Wait()

	return ""
}
