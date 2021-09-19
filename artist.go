package main

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

type artist struct {
	gitpath string
}

func NewArtist(gitPath string) artist {
	return artist{
		gitpath: gitPath,
	}
}

func (a artist) GenerateArt() error {
	commitRegex := regexp.MustCompile("([^\\/]+$)")

	if _, err := os.Stat(filepath.Join(a.gitpath, ".git")); err != nil && os.IsNotExist(err) {
		return fmt.Errorf("no git folder found in %s", a.gitpath)
	}

	objectsPath := filepath.Join(a.gitpath, ".git", "objects", "*", "*")
	commits, err := filepath.Glob(objectsPath)
	if err != nil {
		return err
	}

	if len(commits) < 5 {
		fmt.Errorf("not enough paint to work with, I need more than %d commits to work with", len(commits))
	}

	for _, file := range commits {
		commitId := commitRegex.FindString(file)
		if strings.Contains(commitId, "-") {
			continue
		}
		a.paint(commitId)
	}

	return nil
}

func (a artist) paint(stroke string) {
	for _, c := range stroke {
		fmt.Printf("%s  ", colourTags[string(c)])
	}
	// reset the output
	fmt.Println("\x1b[49m")
}