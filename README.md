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

## docker

sudo docker version
Client:
 Version:	18.03.0-ce
 API version:	1.37
 Go version:	go1.9.4
 Git commit:	0520e24
 Built:	Wed Mar 21 23:10:09 2018
 OS/Arch:	linux/amd64
 Experimental:	false
 Orchestrator:	swarm

Server:
 Engine:
  Version:	18.03.0-ce
  API version:	1.37 (minimum version 1.12)
  Go version:	go1.9.4
  Git commit:	0520e24
  Built:	Wed Mar 21 23:08:36 2018
  OS/Arch:	linux/amd64
  Experimental:	false

## docker build

sudo docker build -t sky0621dhub/go-stepcounter:0.1 .

## docker run

sudo docker run sky0621dhub/go-stepcounter:0.1 _sampleproject/

## docker login -> push

sudo docker login

sudo docker push sky0621dhub/go-stepcounter

## docker hub

https://hub.docker.com/r/sky0621dhub/go-stepcounter/

## docker pull

sudo docker pull sky0621dhub/go-stepcounter:0.1

## docker tag

sudo docker tag go-stepcounter asia.gcr.io/{project-name}/go-stepcounter

