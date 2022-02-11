package main

import (
	"ProjectManager/utils"
	"github.com/AlecAivazis/survey/v2"
	"github.com/fatih/color"
	"os"
	"os/exec"
	"strings"
)

func main() {
	action := ""
	if utils.HandleError(survey.AskOne(utils.ActionUI, &action)) {
		return
	}

	project := ""
	if utils.HandleError(survey.AskOne(utils.ProjectNameUI, &project)) {
		return
	}

	switch action {
	case "Create project":
		if utils.DirExists(project) {
			color.HiRed("This project already exists!")
		} else {
			color.HiGreen("Create project...")
			if utils.HandleError(os.Mkdir(project, 0)) {
				break
			}
			initCmd := exec.Command("go", "mod", "init", project)
			initCmd.Dir = project

			output, initErr := initCmd.CombinedOutput()
			color.HiBlue(string(output[:]))
			if !utils.HandleError(initErr) {
				color.HiGreen("Successfully created project!")
			}
		}
	case "Delete project":
		if !utils.DirExists(project) {
			color.HiRed("This project doesn't exists!")
		} else {
			color.HiGreen("Delete project...")
			if utils.HandleError(os.RemoveAll(project)) {
				break
			}
			color.HiGreen("Successfully deleted project!")
		}
	case "Build project":
		if !utils.DirExists(project) {
			color.HiRed("This project doesn't exists!")
		} else {

			color.HiGreen("Build project...")

			platform := ""
			if utils.HandleError(survey.AskOne(utils.PlatformUI, &platform)) {
				break
			}

			buildCmd := exec.Command("go", "build")
			buildCmd.Dir = project
			buildCmd.Env = os.Environ()
			buildCmd.Env = append(buildCmd.Env, []string{
				"GOOS=" + strings.Split(platform, "/")[0],
				"GOARCH=" + strings.Split(platform, "/")[1],
				"GOPATH=" + os.Getenv("GOPATH")}...)

			output, buildErr := buildCmd.CombinedOutput()
			color.HiBlue(string(output[:]))
			if !utils.HandleError(buildErr) {
				color.HiGreen("Successfully build project!")
			}
		}
	case "Sync dependencies":
		if !utils.DirExists(project) {
			color.HiRed("This project doesn't exists!")
		} else {
			color.HiGreen("Sync dependencies...")
			syncCmd := exec.Command("go", "mod", "tidy")
			syncCmd.Dir = project

			output, syncErr := syncCmd.CombinedOutput()
			color.HiBlue(string(output[:]))
			if !utils.HandleError(syncErr) {
				color.HiGreen("Successfully sync dependencies!")
			}
		}
	case "Open project in VS-Code":
		if !utils.DirExists(project) {
			color.HiRed("This project doesn't exists!")
		} else {
			color.HiGreen("Open project in VS-Code...")
			openCmd := exec.Command("code", project)

			output, openErr := openCmd.CombinedOutput()
			color.HiBlue(string(output[:]))
			if !utils.HandleError(openErr) {
				color.HiGreen("Successfully open VS-Code!")
			}
		}
	}
}
