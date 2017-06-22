package main

import (
	"flag"
	"os"
	"path/filepath"

	"regexp"

	"fmt"

	"bufio"

	"strings"

	"github.com/Sirupsen/logrus"
)

func main() {
	target := flag.String("target", "../sampleproject", "Parse Target")
	flag.Parse()

	err := filepath.Walk(*target, Apply)
	if err != nil {
		logrus.Error(err)
	}
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

	stepCount := 0
	commentCount := 0

	inComment := false

	scanner := bufio.NewScanner(fp)
	for scanner.Scan() {
		txt := strings.TrimSpace(scanner.Text())
		txt2 := strings.Trim(txt, "\t")

		//fmt.Printf("[inComment: %v][stepCount: %d][commentCount: %d] %v\n", inComment, stepCount, commentCount, txt2)

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

	fmt.Printf("[%v] step: %d, comment: %d\n", path, stepCount, commentCount)
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

	outFileExp, err := regexp.Compile(".*test.*")
	if err != nil {
		return false
	}
	if outFileExp.MatchString(path) {
		return false
	}

	inFileExp, err := regexp.Compile(".*.go")
	if err != nil {
		return false
	}
	if !inFileExp.MatchString(path) {
		return false
	}

	return true
}
