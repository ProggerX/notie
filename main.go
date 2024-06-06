package main

import (
	"fmt"
	"os"
	"strconv"
	"sync"

	_ "embed"

	"github.com/BurntSushi/toml"
	"github.com/ProggerX/notie/editor"
	"github.com/ProggerX/notie/viewer"
)

type Config struct {
	EditorPort int `toml:"editor_port"`
	ViewerPort int `toml:"viewer_port"`
}

func isFileExist(path string) bool {
	if _, err := os.Stat(path); err == nil {
		return true
	}
	return false
}

//go:embed example_config.toml
var basicConfig string

func main() {
	if !isFileExist(".notie") {
		fmt.Println(".notie dir is going to be created...")
		os.MkdirAll(".notie/notes", os.ModePerm)
		os.WriteFile(".notie/config.toml", []byte(basicConfig), os.ModePerm)
	}
	bts, _ := os.ReadFile(".notie/config.toml")
	var config Config
	_, _ = toml.Decode(string(bts), &config)

	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		editor.StartEditor("localhost", strconv.Itoa(config.EditorPort))
		wg.Done()
	}()
	go func() {
		viewer.StartViewer("localhost", strconv.Itoa(config.ViewerPort))
		wg.Done()
	}()

	wg.Wait()
}
