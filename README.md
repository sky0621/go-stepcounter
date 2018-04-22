# go-stepcounter

指定ディレクトリ配下のgoソースのステップカウンター

## パラメータ

##### [実行例] cmd 配下にて下記コマンド実行

<pre> go run main.go template.go ./_sampleproject </pre>

## TODO
 
##### ・出力内容をテンプレート化

##### ・外部から使いやすいように出力方法を見直す

## dep

https://github.com/golang/dep

dep init

dep ensure

## go-assets-builder

go get -v github.com/jessevdk/go-assets-builder

go-assets-builder tmpl/ > template.go

## gox

https://github.com/mitchellh/gox

gox -os="linux darwin windows" -arch="amd64"

## ghr

https://github.com/tcnksm/ghr

git config --global github.token "....."

export GITHUB_API=http://github.company.com/api/v3/

ghr v0.1.0 dist/
