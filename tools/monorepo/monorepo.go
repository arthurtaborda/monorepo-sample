package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"
)

func runCommand(cmd string, errorMessage string) (string, error) {
	//fmt.Println("executing command:", cmd)
	commandSplit := strings.Split(cmd, " ")
	command := exec.Command(commandSplit[0], commandSplit[1:]...)

	output, err := command.Output()
	if err != nil {
		log.Printf(errorMessage+": %v", err)
	}
	return strings.TrimSuffix(string(output), "\n"), err
}

func runCommandAttached(cmd string, errorMessage string) error {
	commandSplit := strings.Split(cmd, " ")
	command := exec.Command(commandSplit[0], commandSplit[1:]...)
	command.Stdout = os.Stdout
	command.Stderr = os.Stderr
	err := command.Run()
	if err != nil {
		log.Printf(errorMessage+": %v", err)
	}
	return err
}

func main() {
	var path string
	var command string
	var commitRange string
	flag.StringVar(&path, "path", "", "The path of the service that should be built")
	flag.StringVar(&command, "command", "", "The command that should run if the path changed")
	flag.StringVar(&commitRange, "commitRange", "", "The commit range to be used to check if the path changed")

	flag.Parse()

	//fmt.Println("path:", path)
	//fmt.Println("command:", command)
	//fmt.Println("commitRange:", commitRange)

	if _, err := os.Stat(path); os.IsNotExist(err) {
		fmt.Printf("Project in path %s does not exist", path)
		return
	}

	if commitRange == "" {
		currentBranch, _ := runCommand("git rev-parse --abbrev-ref HEAD", "Couldn't get current branch")
		headHash, _ := runCommand("git rev-parse HEAD", "Couldn't get head hash")

		hashToCompare := ""
		if currentBranch == "master" {
			hashToCompare, _ = runCommand("git rev-parse "+headHash+"~1", "Couldn't get previous git hash")
		} else {
			hashToCompare, _ = runCommand("git rev-parse origin/master", "Couldn't get master hash")
		}

		commitRange = hashToCompare + ".." + headHash

		//fmt.Println("currentBranch:", currentBranch)
		//fmt.Println("headHash:", headHash)
		//fmt.Println("hashToCompare:", hashToCompare)
		//fmt.Println("commitRange:", commitRange)
	}

	changedPathsCmd, _ := runCommand("git diff --name-only "+commitRange, "Couldn't get changed paths")
	changedPaths := strings.Split(changedPathsCmd, "\n")

	fmt.Println("changedPaths:", changedPaths)

	shouldBuild := false
	for _, changedPath := range changedPaths {
		if strings.HasPrefix(changedPath, path) {
			shouldBuild = true
		}
	}

	if shouldBuild {
		fmt.Printf("Building project in path %s\n", path)
		_ = runCommandAttached(command, "Couldn't build project")
	} else {
		fmt.Printf("Project in path %s didn't change, skipping...\n", path)
	}
}
