# Markdown Editor
[![Build Status](https://travis-ci.org/linyuy/editor.svg?branch=master)](https://travis-ci.org/linyuy/editor)
Integrate background services on [pandao/editor.md](https://github.com/pandao/editor.md) and [hawtim/editor.md](https://github.com/hawtim/editor.md).

## Build Setup

### Compile Manually 
```shell
> go get github.com/linyuy/editor
> cd $GOPATH/src/github.com/linyuy/editor
> go build && editor
```

### Docker Service
```
docker pull autnihil/editor
docker run -it --rm --name editor -p 80:80 autnihil/editor
```
