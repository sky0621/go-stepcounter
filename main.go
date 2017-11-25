package main

import (
	"bytes"
	"html/template"
	"os"
	"path/filepath"

	"regexp"

	"fmt"

	"bufio"

	"strings"

	"go.uber.org/zap"
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
	logger, err := zap.NewProduction()
	if err != nil {
		panic(err)
	}
	defer logger.Sync()

	if len(os.Args) < 2 {
		logger.Error("引数[ターゲットディレクトリのパス]が必要です")
		os.Exit(-1)
	}
	target := os.Args[1]

	err = filepath.Walk(target, Apply)
	if err != nil {
		logger.Error("", zap.String("error", err.Error()))
		os.Exit(-1)
	}

	result.TotalStep = allStepCount
	result.TotalComment = allCommentCount

	tmplPath := "tmpl/tmpl.csv"
	a, err := Asset(tmplPath)
	if err != nil {
		logger.Error("", zap.String("error", err.Error()))
		os.Exit(-1)
	}

	tmpl := template.Must(template.New(tmplPath).Parse(string(a)))
	buf := &bytes.Buffer{}
	err = tmpl.Execute(buf, result)
	if err != nil {
		logger.Error("", zap.String("error", err.Error()))
		os.Exit(-1)
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
