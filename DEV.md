# go-stepcounter

## go module

```bash
$ go mod init github.com/sky0621/go-stepcounter
$ go mod download
```

## go-assets-builder

https://github.com/jessevdk/go-assets-builder

```bash
$ go get -v github.com/jessevdk/go-assets-builder
$
$ go-assets-builder tmpl/ > template.go
```

## goreleaser

https://goreleaser.com

```bash
$ git tag v0.1.0
$
$ git push origin v0.1.0
$
$ goreleaser
```

## docker

```bash
$ sudo docker version
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
```

## docker build

```bash
$ sudo docker build -t sky0621dhub/go-stepcounter:0.1 .
```

## docker run

```bash
$ sudo docker run sky0621dhub/go-stepcounter:0.1 _sampleproject/
```

## docker login -> push

```bash
$ sudo docker login
$
$ sudo docker push sky0621dhub/go-stepcounter
```

## docker hub

https://hub.docker.com/r/sky0621dhub/go-stepcounter/

## docker pull

```bash
$ sudo docker pull sky0621dhub/go-stepcounter:0.1
```

## docker tag

```bash
$ sudo docker tag go-stepcounter asia.gcr.io/{project-name}/go-stepcounter
```
