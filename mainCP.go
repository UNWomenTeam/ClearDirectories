package main

import (
	"fmt"
	"path/filepath"
	"regexp"
)

func main() {
	fmt.Printf("ಠ┗(▀̿Ĺ̯▀̿ ̿)┓ \n")
	fmt.Printf("       Clean Picture v1.01\n")
	fmt.Printf("        	 ❥ очистка от картинок которые накопились\n")
	fmt.Printf("        	 ❥ расположите ПО в папке ci_cd проекта \"Shoplifter 42\"\n")

	dir, err := BinDir()
	Fck(err)
	fmt.Printf("        	 ❥ расположение в проекте \"%s\"\n", dir)
	dockerCompseFName := "docker-compose.yml"
	fmt.Printf("        	 ❥ открываем  \"%s\"\n", dockerCompseFName)
	datComposeProjectFile := OpenReadFile(filepath.Join(dir, dockerCompseFName))
	findDir := "\\/opt\\/rainbow_video"
	exression := fmt.Sprint("-\\s*(.+):", findDir)
	fmt.Printf("        	 ❥ ищем какая директория у  \"%s\"\n", findDir)
	var compRegEx = regexp.MustCompile(exression)
	matches := compRegEx.FindAllStringSubmatch(string(datComposeProjectFile), -1)
	var localDirRainbow string
	if len(matches) > 0 {
		if len(matches[0]) > 1 {
			localDirRainbow = matches[0][1]
		}
	}
	if localDirRainbow == "" {
		Fck(fmt.Errorf("не найдена папка для очистки rx:%s", exression))
	}

	fmt.Printf("        	 ❥ найдена начальная директория:\"%s\"\n", localDirRainbow)
	cleanFolderName := "rendering/pictures"
	fmt.Printf("        	 ❥ очистка, этап 1 папка:\"%s\"\n", cleanFolderName)
	cleanFolderPth := filepath.Join(localDirRainbow, cleanFolderName)
	ExecuteRemover(cleanFolderPth)
	fmt.Printf("        	 ❥ очистка, этап 1 :OK\n")
	FckText("не удалось выполнить команду CMD", err)

	cleanFolderName = "rendering/videos/"
	fmt.Printf("        	 ❥ очистка, этап 2 папка:\"%s\"\n", cleanFolderName)
	cleanFolderPth = filepath.Join(localDirRainbow, cleanFolderName)
	ExecuteRemover(cleanFolderPth)
	fmt.Printf("        	 ❥ очистка, этап 2 :OK\n")

	cleanFolderName = "streaming/"
	fmt.Printf("        	 ❥ очистка, этап 3 папка:\"%s\"\n", cleanFolderName)
	cleanFolderPth = filepath.Join(localDirRainbow, cleanFolderName)
	ExecuteRemover(cleanFolderPth)
	fmt.Printf("        	 ❥ очистка, этап 3 :OK\n")
	fmt.Printf("    Good by  ٩◔̯◔۶\n")
}
