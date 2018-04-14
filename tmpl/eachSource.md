# 各ソースステップ数一覧({{.Datetime}} 時点)

#### ※ツール（ https://github.com/sky0621/go-stepcounter ）による自動生成

Total

| TotalStep | {{.TotalStep}} |
| TotalComment | {{.TotalComment}} |

| Path | Step | Comment |
{{range .FileStepCounters}}| {{.FilePath}} | {{.Step}} | {{.Comment}} |
{{end}}
