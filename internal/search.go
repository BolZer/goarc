package internal

import (
	"bufio"
	"errors"
	"fmt"
	"os/exec"
	"regexp"
	"strings"
)

const Gw2ExecutableName = "Gw2-64.exe"

func SearchForLocalGuildWarsInstallation() (string, error) {
	disks, err := listDiscOfOs()

	if err != nil {
		return "", fmt.Errorf("error encountered while listing disks of os: %w", err)
	}

	result := ""

	for _, disk := range disks {
		result = searchForGuildWars2ExecutableOnDisk(disk)

		if result != "" {
			return strings.TrimSpace(strings.ReplaceAll(result, "\\"+Gw2ExecutableName, "")), nil
		}
	}

	return "", errors.New("couldn't find installation of guild wars 2")
}

func listDiscOfOs() ([]string, error) {
	out, err := exec.Command("wmic", "logicaldisk", "get", "name").Output()

	if err != nil {
		return []string{}, fmt.Errorf("command for listing os disks failed: %w", err)
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
		return ""
	}

	scanner := bufio.NewScanner(stdout)

	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		return scanner.Text()
	}

	err = command.Wait()

	return ""
}
