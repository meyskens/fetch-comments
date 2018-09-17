package main

import (
	"io/ioutil"
	"os"
	"path"
	"regexp"

	"gopkg.in/src-d/go-billy.v4"

	git "gopkg.in/src-d/go-git.v4"
)

// RepoFiles contains a list of files per language of a repo
type RepoFiles struct {
	GoFiles         []string
	PythonFiles     []string
	JavaFiles       []string
	JavascriptFiles []string
}

var goFile = regexp.MustCompile(`\.go$`)
var pythonFile = regexp.MustCompile(`\.py$`)
var javaFile = regexp.MustCompile(`\.java$`)
var javascriptFile = regexp.MustCompile(`\.js$`)

func getFiles(url string) (*RepoFiles, billy.Filesystem, error) {
	dir, err := ioutil.TempDir(os.TempDir(), "fetch-comments")
	if err != nil {
		return nil, nil, err
	}
	r, err := git.PlainClone(dir, false, &git.CloneOptions{
		URL: url,
	})
	if err != nil {
		return nil, nil, err
	}

	worktree, err := r.Worktree()
	if err != nil {
		return nil, nil, err
	}

	files, err := worktree.Filesystem.ReadDir("/")
	if err != nil {
		return nil, nil, err
	}

	fileNames := []string{}

	for _, file := range files {
		if file.IsDir() {
			fileNames = append(fileNames, getAllDirectoryFiles(worktree.Filesystem, file.Name())...)
		} else {
			fileNames = append(fileNames, file.Name())
		}
	}

	repoFiles := RepoFiles{
		GoFiles:         []string{},
		PythonFiles:     []string{},
		JavaFiles:       []string{},
		JavascriptFiles: []string{},
	}

	for _, name := range fileNames {
		if goFile.MatchString(name) {
			repoFiles.GoFiles = append(repoFiles.GoFiles, name)
		}
		if pythonFile.MatchString(name) {
			repoFiles.PythonFiles = append(repoFiles.PythonFiles, name)
		}
		if javaFile.MatchString(name) {
			repoFiles.JavaFiles = append(repoFiles.JavaFiles, name)
		}
		if javascriptFile.MatchString(name) {
			repoFiles.JavascriptFiles = append(repoFiles.JavascriptFiles, name)
		}
	}

	return &repoFiles, worktree.Filesystem, nil
}

func getAllDirectoryFiles(fs billy.Filesystem, p string) []string {
	out := []string{}
	files, err := fs.ReadDir(p)
	if err != nil {
		return out
	}

	for _, file := range files {
		if file.IsDir() {
			out = append(out, getAllDirectoryFiles(fs, path.Join(p, file.Name()))...)
		}
		out = append(out, path.Join(p, file.Name()))
	}

	return out
}
