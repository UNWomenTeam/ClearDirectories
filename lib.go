package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"time"
)

func BinDir() (string, error) {
	ex, err := os.Executable()
	if err != nil {
		return "", err
	}
	return filepath.Dir(ex), nil
}
func Fck(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "ERROR: %s", err)
		ExitWithSecTimeout(1)
	}
}
func FckText(text string, err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "ERROR HEAD: %s\n ERROT TEXT: %s", text, err)
		ExitWithSecTimeout(1)
	}
}
func ExitWithSecTimeout(status int) {
	// 0 - norm
	// 1 - error
	fmt.Println("\n Завершение работы программы")
	time.Sleep(200 * time.Millisecond)
	os.Exit(status)
}
func OpenReadFile(filePath string) (dat []byte) {
	dat, err := ioutil.ReadFile(filePath)
	Fck(err)
	return
}

func ExecuteRemover(remDir string) {
	cmd := exec.Command("rm", "-r", "-f", remDir)
	var outb, errb bytes.Buffer
	cmd.Stdout = &outb
	cmd.Stderr = &errb
	err := cmd.Run()
	fmt.Println("[rm][", remDir, "] | out:", outb.String(), "| bin/bash:", errb.String())
	if err != nil {
		log.Fatal("error exit!")
	}

	cmd = exec.Command("mkdir", remDir)
	outb.Reset()
	errb.Reset()
	cmd.Stdout = &outb
	cmd.Stderr = &errb
	err = cmd.Run()
	fmt.Println("[mkdir] out:", outb.String(), "bin/bash:", errb.String())
	if err != nil {
		log.Fatal("error exit!")
	}

	cmd = exec.Command("chmod", "777", remDir)
	outb.Reset()
	errb.Reset()
	cmd.Stdout = &outb
	cmd.Stderr = &errb
	err = cmd.Run()
	fmt.Println("[mkdir] out:", outb.String(), "bin/bash:", errb.String())
	if err != nil {
		log.Fatal("error exit!")
	}
}
