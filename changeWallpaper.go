package main

import (
	"flag"
	"io/ioutil"
	"log"
	"math/rand"
	"os/exec"
	"path/filepath"
	"strings"
	"time"
)

func hasSuffixes(s string, suffixes []string) bool {
	for _, suffix := range suffixes {
		if strings.HasSuffix(s, suffix) {
			return true
		}
	}
	return false
}
func getImageFilePaths(root string) []string {
	fileInfos, err := ioutil.ReadDir(root)
	if err != nil {
		log.Fatalf("Error: %v", err)
	}
	paths := []string{}
	suffixes := []string{".jpg", ".jpeg", ".png"}
	for _, fileInfo := range fileInfos {
		fileName := fileInfo.Name()
		if hasSuffixes(fileName, suffixes) {
			paths = append(paths, fileName)
		}
	}
	return paths
}
func main() {
	flag.Parse()
	root, err := filepath.Abs(flag.Arg(0))
	log.Printf("Specified root path is %v", root)
	if err != nil {
		log.Fatalf("Error: %v", err)
	}
	paths := getImageFilePaths(root)
	path := ""
	if len(paths) != 0 {
		rand.Seed(int64(time.Now().Nanosecond()))
		r := rand.Intn(len(paths))
		path = filepath.Join(root, paths[r])
	}
	command := exec.Command(
		"sh",
		"-c",
		"dconf write /org/gnome/desktop/background/picture-uri \"'"+path+"'\"",
	)
	out, err := command.Output()
	if err != nil {
		log.Fatalf("Error: %v", err)
	}
	log.Printf("Command output %v", out)
	log.Printf("Set wallpaper %v", path)
}
