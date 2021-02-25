package main

import (
	"strings"

	helper "github.com/bolZer/goarc/v2/internal"
)

func main() {
	var guildWars2InstallationPath string
	var arcDpsFileDestinationPath string

	helper.OutputToConsoleWithWarningFormatting("Start searching for Guild Wars 2 Installation!")

	guildWars2InstallationPath, err := helper.SearchForGuildWarsInstallation()

	if err != nil {
		helper.OutputToConsoleWithAlertFormatting("No installation of Guild Wars 2 found. Exit.")
		return
	}

	helper.OutputToConsoleWithSuccessFormatting("Installation of Guild Wars 2 found!")
	helper.OutputToConsoleWithWarningFormatting("Checking if ArcDPS exists")

	arcDpsFileDestinationPath = strings.Join([]string{
		guildWars2InstallationPath,
		"bin64",
		"d3d9.dll",
	}, "\\")

	doesArcDpsExist := helper.CheckIfArcDPSExists(arcDpsFileDestinationPath)

	if !doesArcDpsExist {
		helper.OutputToConsoleWithWarningFormatting("ArcDPS does not exists.")
	}

	if doesArcDpsExist {
		helper.OutputToConsoleWithWarningFormatting("ArcDPS exists. Checking if outdated")

		isExistingArcDpsOutdated, err := helper.CheckIfArcDPSIsOutdated(arcDpsFileDestinationPath)

		if err != nil {
			helper.OutputToConsoleWithAlertFormatting(err.Error())
			return
		}

		if !isExistingArcDpsOutdated {
			helper.OutputToConsoleWithAlertFormatting("ArcDPS is not outdated. Exit.")
			return
		}
	
	}

	helper.OutputToConsoleWithWarningFormatting("Downloading ArcDPS")

	err = helper.DownloadArcDPStoDestinationPath(arcDpsFileDestinationPath)

	if err != nil {
		helper.OutputToConsoleWithAlertFormatting(err.Error())
		return
	}

	helper.OutputToConsoleWithSuccessFormatting("Done! ArcDPS is ready to be used")
}
