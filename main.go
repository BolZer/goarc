package main

import (
	"fmt"
	"log"
	"strings"

	"github.com/bolZer/goarc/v2/internal"
)

func main() {
	var guildWars2InstallationPath string
	var arcDpsFileDestinationPath string

	fmt.Println("Start searching for Guild Wars 2 Installation!")

	guildWars2InstallationPath, err := internal.SearchForLocalGuildWarsInstallation()

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Installation of Guild Wars 2 found! Checking if ArcDPS exists.")

	arcDpsFileDestinationPath = strings.Join([]string{
		guildWars2InstallationPath,
		"bin64",
		"d3d9.dll",
	}, "\\")

	doesArcDpsExist := internal.CheckIfLocalArcDPSExists(arcDpsFileDestinationPath)

	if !doesArcDpsExist {
		fmt.Println("ArcDPS does not exists.")
	}

	if doesArcDpsExist {
		fmt.Println("ArcDPS exists. Checking if it's outdated.")

		isExistingArcDpsOutdated, err := internal.CheckIfLocalArcDpsDiffersFromRemoteOne(arcDpsFileDestinationPath)

		if err != nil {
			log.Fatal(err)
		}

		if !isExistingArcDpsOutdated {
			fmt.Println("ArcDPS is not outdated. Exit.")
			return
		}
	}

	fmt.Println("Downloading ArcDPS")

	err = internal.DownloadRemoteArcDPSToDestinationPath(arcDpsFileDestinationPath)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Done! ArcDPS is ready to be used")
}
