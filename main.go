package main

import (
	"flag"
	"io/fs"
	"log"
	"os"
	"strings"
	"sync"
)

func graph(email *string) {
	_ = ""
}

func GetPath(dir *string, file *os.File) error {
	return fs.WalkDir(os.DirFS(*dir), ".", func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		// BAD CODE
		if d.Name() == "node_modules" || d.Name() == ".cargo" || d.Name() == "pip" || d.Name() == ".cache" || d.Name() == ".config" || d.Name() == ".local" {
			return fs.SkipDir
		}

		if d.Name() == ".git" {
			// shit line
			_, err := file.WriteString(strings.TrimRight(*dir, "/") + "/" + path[:len(path)-4] + "\n")
			if err != nil {
				return err
			}
			return fs.SkipDir
		}

		return nil
	})
}

func Getdir(dir *string, email string) error {
	homeDir, err := os.UserHomeDir()
	_ = homeDir
	if err != nil {
		return err
	}

	//workingpath := homeDir + "/.gitgraph/" + email + ".gitgraoh"
	workingpath := email + ".gitgraph"
	file, err := os.OpenFile(workingpath, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)

	if err != nil {
		return err
	}
	defer file.Close()
	return GetPath(dir, file)
}

func main() {
	var wg sync.WaitGroup
	dir := flag.String("add", "", "add sub directory")
	email := flag.String("email", "your@email.com", "please provide your email")
	flag.Parse()

	wg.Add(1)

	if *dir != "" {
		go func() {
			defer wg.Done()
			if err := Getdir(dir, *email); err != nil {
				log.Println(err)
			}
		}()
		wg.Wait()
	}

	graph(email)
}
