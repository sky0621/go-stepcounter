# go-stepcounter

指定ディレクトリ配下のgoソースのステップカウンター

## パラメータ

##### [実行例] cmd 配下にて下記コマンド実行

<pre> go run main.go ..\\_sampleproject </pre>

## TODO
 
##### ・出力内容をテンプレート化

##### ・外部から使いやすいように出力方法を見直す

## dep

dep init

dep ensure

## go-bindata

go-bindata tmpl/

## goxc

goxc -os="linux darwin windows" -arch="386 amd64" -d=dist -tasks="clean-destination,xc,archive-zip,rmbin" -wc

goxc

## ghr

ghr v0.0.x dist/snapshot/
