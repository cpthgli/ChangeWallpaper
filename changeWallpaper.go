package main

import (
	"flag"
	"io"
	"io/ioutil"
	"log"
	"math/rand"
	"os"
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
		log.Fatalf("ERROR %v", err)
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
	isLoging := flag.Bool("log", false, "write log for bool")
	flag.Parse()
	if *isLoging {
		logfile, err := os.OpenFile("./debug.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
		if err != nil {
			panic("cannot open debug.log:" + err.Error())
		}
		defer logfile.Close()
		log.SetOutput(io.MultiWriter(logfile, os.Stdout))
	}
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
	root, err := filepath.Abs(flag.Arg(0))
	log.Printf("INFO Specified root path is %v", root)
	if err != nil {
		log.Fatalf("ERROR %v", err)
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
		log.Fatalf("ERROR %v", err)
	}
	log.Printf("INFO Command output %v", out)
	log.Printf("INFO Set wallpaper %v", path)
}
