package main

import (
	"bytes"
	"flag"
	"html/template"
	"os"
	"path/filepath"

	"regexp"

	"fmt"

	"bufio"

	"strings"

	"github.com/Sirupsen/logrus"
)

var (
	allStepCount    int64 = 0
	allCommentCount int64 = 0
)

type StepCounter struct {
	TotalStep        int64
	TotalComment     int64
	FileStepCounters []*FileStepCounter
}

type FileStepCounter struct {
	FilePath string
	Step     int64
	Comment  int64
}

var result *StepCounter = &StepCounter{FileStepCounters: []*FileStepCounter{}}

func main() {
	target := flag.String("target", "_sampleproject", "Parse Target")
	flag.Parse()

	err := filepath.Walk(*target, Apply)
	if err != nil {
		logrus.Error(err)
	}
	result.TotalStep = allStepCount
	result.TotalComment = allCommentCount

	tmpl := template.Must(template.ParseFiles("tmpl.csv"))
	buf := &bytes.Buffer{}
	err = tmpl.Execute(buf, result)
	if err != nil {
		panic(err)
	}

	fmt.Println(buf.String())
}

func Apply(path string, info os.FileInfo, err error) error {
	if err != nil {
		return err
	}

	if !filter(path, info) {
		return nil
	}

	fp, err := os.Open(path)
	if err != nil {
		return err
	}
	defer func() {
		if fp != nil {
			fp.Close()
		}
	}()

	var stepCount int64 = 0
	var commentCount int64 = 0

	inComment := false

	scanner := bufio.NewScanner(fp)
	for scanner.Scan() {
		txt := strings.TrimSpace(scanner.Text())
		txt2 := strings.Trim(txt, "\t")

		if strings.HasPrefix(txt2, "/*") && strings.HasSuffix(txt2, "*/") {
			inComment = false
			commentCount = commentCount + 1
			continue
		}

		if strings.HasPrefix(txt2, "/*") {
			inComment = true
			commentCount = commentCount + 1
			continue
		}

		if strings.HasPrefix(txt2, "*/") {
			inComment = false
			commentCount = commentCount + 1
			continue
		}

		if inComment {
			commentCount = commentCount + 1
			continue
		}

		if strings.HasPrefix(txt2, "//") || strings.HasPrefix(txt2, "*") {
			commentCount = commentCount + 1
			continue
		}

		if txt2 == "" {
			continue
		}

		stepCount = stepCount + 1
	}

	result.FileStepCounters = append(result.FileStepCounters, &FileStepCounter{FilePath: path, Step: stepCount, Comment: commentCount})

	allStepCount = allStepCount + stepCount
	allCommentCount = allCommentCount + commentCount

	return nil
}

func filter(path string, info os.FileInfo) bool {
	if info.IsDir() {
		return false
	}

	absPath, err := filepath.Abs(path)
	if err != nil {
		return false
	}

	outDirExp, err := regexp.Compile("vendor")
	if err != nil {
		return false
	}
	if outDirExp.MatchString(absPath) {
		return false
	}

	outDirExp2, err := regexp.Compile("\\.git")
	if err != nil {
		return false
	}
	if outDirExp2.MatchString(absPath) {
		return false
	}

	outFileExp, err := regexp.Compile(".*test.*")
	if err != nil {
		return false
	}
	if outFileExp.MatchString(path) {
		return false
	}

	inFileExp, err := regexp.Compile(".*\\.go")
	if err != nil {
		return false
	}
	if !inFileExp.MatchString(path) {
		return false
	}

	return true
}
