package main

import (
	"time"

	"github.com/jessevdk/go-assets"
)

var _Assets95aa4ed3a0a5918c330d9340efd86426339c5692 = "# 各ソースステップ数一覧({{.Datetime}} 時点)\n\n#### ※ツール（ https://github.com/sky0621/go-stepcounter ）による自動生成\n\nTotal\n\n| TotalStep | {{.TotalStep}} |\n| TotalComment | {{.TotalComment}} |\n\n| Path | Step | Comment |\n{{range .FileStepCounters}}| {{.FilePath}} | {{.Step}} | {{.Comment}} |\n{{end}}\n"

// Assets returns go-assets FileSystem
var Assets = assets.NewFileSystem(map[string][]string{"/": []string{"tmpl"}, "/tmpl": []string{"eachSource.md"}}, map[string]*assets.File{
	"/": &assets.File{
		Path:     "/",
		FileMode: 0x800001ed,
		Mtime:    time.Unix(1523709893, 1523709893798453855),
		Data:     nil,
	}, "/tmpl": &assets.File{
		Path:     "/tmpl",
		FileMode: 0x800001ed,
		Mtime:    time.Unix(1523712493, 1523712493430865905),
		Data:     nil,
	}, "/tmpl/eachSource.md": &assets.File{
		Path:     "/tmpl/eachSource.md",
		FileMode: 0x1b4,
		Mtime:    time.Unix(1523712471, 1523712471743544939),
		Data:     []byte(_Assets95aa4ed3a0a5918c330d9340efd86426339c5692),
	}}, "")
