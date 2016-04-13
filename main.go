package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/exec"
	"path"
	"path/filepath"
	"time"
)

const convert = "/usr/bin/convert"

var (
	pathMedia = flag.String("path", "/media", "Base path for storing files")
	port      = flag.String("port", ":8080", "Port for handling requests")
	jsonCfg   = flag.String("json", "presets.json", "Json file with loaded presets")
	thumbPath = "t"
)

var presets = map[string]string{
	"roig300": "-colorspace gray -level +10% +level-colors '#000000','#ffec00' -thumbnail 350x250^ -extent 350x250 -quality 82",
}

func main() {

	flag.Parse()
	err := loadPresets()
	if err != nil {
		log.Fatalf("Can't load presets file %s", err)
	}

	mux := http.NewServeMux()
	mux.Handle("/", http.HandlerFunc(handleMedia))
	log.Printf("Service http file server started at %s", *port)
	if err := http.ListenAndServe(*port, mux); err != nil {
		log.Fatal(err)
	}

}

func handleMedia(w http.ResponseWriter, r *http.Request) {
	file := fmt.Sprintf("%s%s", *pathMedia, r.URL.Path)
	pr := r.FormValue("p")
	if pr != "" {
		// test if preset exists
		if _, ok := presets[pr]; ok {
			nfile := filepath.Join(*pathMedia, thumbPath, pr, r.URL.Path)
			log.Printf("Preset filepath %s", nfile)
			// test if file exists prior to generate it
			if _, err := os.Stat(nfile); err == nil {
				file = nfile
			} else {
				err := convertFile(pr, file, nfile)
				if err == nil {
					file = nfile
				} else {
					log.Printf("Error generating thumbnail %s", err)
				}
			}
		}
	}

	log.Printf("Serving file %s", file)
	if ff, err := os.Stat(file); err == nil && ff.IsDir() == false {
		http.ServeFile(w, r, path.Clean(file))
	}
}

func convertFile(preset, file, nfile string) error {
	start := time.Now()
	var err error
	dir := filepath.Dir(nfile)
	// ensure path exists prior to convert file
	err = os.MkdirAll(dir, 0755)
	if err != nil {
		return err
	}
	command := fmt.Sprintf("%s %s %s %s",
		convert, file, presets[preset], nfile)
	//err := runOut(os.Stdout, "sh -c", command)
	cmd := exec.Command("sh", "-c", command)
	log.Printf("[thumb] command: sh -c %s", command)
	cmd.Stdin = nil
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err = cmd.Run()
	log.Printf("[thumb] %s %s generated", time.Since(start), nfile)
	return err
}

func loadPresets() error {
	file, err := ioutil.ReadFile(*jsonCfg)
	if err != nil {
		log.Fatalf("Invalid config %s: %s", *jsonCfg, err)
		return err
	}
	//var m map[string]string
	err = json.Unmarshal(file, &presets)
	if err != nil {
		log.Fatalf("Invalid config %s: %s", *jsonCfg, err)
		return err
	}
	log.Println("Loaded presets")
	for k, v := range presets {
		log.Printf("%s --> %s", k, v)
	}

	return nil
}
