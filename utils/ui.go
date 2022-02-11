package utils

import (
	"github.com/AlecAivazis/survey/v2"
	"os"
	"strings"
)

var ProjectNameUI = &survey.Input{
	Message: "Project:",
	Suggest: func(toComplete string) []string {
		folders := []string{}
		dirContent, dirErr := os.ReadDir(".")
		if !HandleError(dirErr) {
			for _, entry := range dirContent {
				if strings.HasPrefix(strings.ToLower(entry.Name()), strings.ToLower(toComplete)) && entry.IsDir() {
					folders = append(folders, entry.Name())
				}
			}
		}

		return folders
	},
}

var ActionUI = &survey.Select{
	Message: "What would you do?",
	Options: []string{"Create project", "Build project", "Delete project", "Sync dependencies", "Open project in VS-Code"},
}

var PlatformUI = &survey.Select{
	Message: "Choose the Build-Platform:",
	Options: AllPlatforms,
}
