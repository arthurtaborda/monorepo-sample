package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"
)

//var monorepoYaml = `
//projects:
//  - path: services/server
//    build: ./gradlew -p services/server assemble
//    test: ./gradlew -p services/server test
//  - path: services/client
//    build: ./gradlew -p services/client assemble
//    test: ./gradlew -p services/client test
//  - path: libraries/common
//    build: ./gradlew -p libraries/common assemble
//    test: ./gradlew -p libraries/common test
//  - path: libraries/logging
//    build: ./gradlew -p libraries/logging assemble
//    test: ./gradlew -p libraries/logging test
//  - path: tools/gradle-plugins/gradle-hello-plugin
//    build: ./gradlew -p tools/gradle-plugins/gradle-hello-plugin assemble
//    test: ./gradlew -p tools/gradle-plugins/gradle-hello-plugin test
//`
//
//type MonorepoConfig struct {
//	Projects []Project `yaml:"projects"`
//}

//type Project struct {
//	Path  string
//	Build string
//	Test  string
//}

func runCommand(cmd string, errorMessage string) (string, error) {
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
	flag.StringVar(&path, "path", "", "The path of the service that should be built")
	flag.StringVar(&command, "command", "", "The command that should run if the path changed")

	flag.Parse()

	fmt.Println("path:", path)
	fmt.Println("command:", command)

	if _, err := os.Stat(path); os.IsNotExist(err) {
		fmt.Printf("Project in path %s does not exist", path)
		return
	}

	currentBranch, _ := runCommand("git rev-parse --abbrev-ref HEAD", "Couldn't get current branch")
	headHash, _ := runCommand("git rev-parse HEAD", "Couldn't get head hash")

	hashToCompare := ""
	if currentBranch == "master" {
		hashToCompare, _ = runCommand("git rev-parse "+headHash+"~1", "Couldn't get previous git hash")
	} else {
		hashToCompare, _ = runCommand("git rev-parse origin/master", "Couldn't get master hash")
	}

	commitRange := hashToCompare + ".." + headHash

	changedPathsCmd, _ := runCommand("git diff --name-only "+commitRange, "Couldn't get changed paths")
	changedPaths := strings.Split(changedPathsCmd, "\n")

	fmt.Println("currentBranch:", currentBranch)
	fmt.Println("headHash:", headHash)
	fmt.Println("hashToCompare:", hashToCompare)
	fmt.Println("commitRange:", commitRange)
	fmt.Println("changedPaths:", changedPaths)

	//config := MonorepoConfig{}
	//err := yaml.Unmarshal([]byte(monorepoYaml), &config)
	//if err != nil {
	//	log.Fatalf("error: %v", err)
	//}

	shouldBuild := false
	for _, changedPath := range changedPaths {
		if strings.HasPrefix(changedPath, path) {
			shouldBuild = true
		}
	}

	if shouldBuild {
		fmt.Printf("Building project in path %s", path)
		_ = runCommandAttached(command, "Couldn't build project")
	} else {
		fmt.Printf("Project in path %s didn't change, skipping...", path)
	}

	//projectsToTest := make(map[string]Project)
	//for _, project := range config.Projects {
	//	for _, changedPath := range changedPaths {
	//		if strings.HasPrefix(changedPath, project.Path) {
	//			projectsToTest[project.Path] = project
	//		}
	//	}
	//}

	//fmt.Println(projectsToTest)

	//fmt.Printf("--- config:\n%v\n\n", config)

	//for _, project := range config.Projects {
	//	if _, exists := projectsToTest[project.Path]; exists {
	//		log.Printf("Running build for project in " + project.Path + "\n\n")
	//		_ = runCommandAttached(project.Build, "Couldn't build project "+project.Path)
	//	} else {
	//		log.Printf("Project " + project.Path + " should not be built, skipping...\n\n")
	//	}
	//}
	//
	//if err != nil {
	//	log.Printf("Command finished with error: %v", err)
	//}

	//d, err := yaml.Marshal(&config)
	//if err != nil {
	//	log.Fatalf("error: %v", err)
	//}
	//fmt.Printf("--- t dump:\n%s\n\n", string(d))
	//
	//m := make(map[interface{}]interface{})
	//
	//err = yaml.Unmarshal([]byte(monorepoYaml), &m)
	//if err != nil {
	//	log.Fatalf("error: %v", err)
	//}
	//fmt.Printf("--- m:\n%v\n\n", m)
	//
	//d, err = yaml.Marshal(&m)
	//if err != nil {
	//	log.Fatalf("error: %v", err)
	//}
	//fmt.Printf("--- m dump:\n%s\n\n", string(d))
}
