# v2rayXPlus
![](https://gitlab.com/xiayesuifeng/v2rayxplus/raw/master/resources/v2rayXPlus-64px.svg)

> 使用go写的v2ray gui客户端

[![Go Report Card](https://goreportcard.com/badge/gitlab.com/xiayesuifeng/v2rayxplus)](https://goreportcard.com/report/gitlab.com/xiayesuifeng/v2rayxplus)
[![GoDoc](https://godoc.org/gitlab.com/xiayesuifeng/v2rayxplus?status.svg)](https://godoc.org/gitlab.com/xiayesuifeng/v2rayxplus)
[![Sourcegraph](https://sourcegraph.com/gitlab.com/xiayesuifeng/v2rayxplus/-/badge.svg)](https://sourcegraph.com/gitlab.com/xiayesuifeng/v2rayxplus)

使用iptables+透明代理+v2ray路由实现真全局自动分流

感谢[@linyuan](t.me/linyuan)提供的logo


## Dependencies
[therecipe/qt](https://github.com/therecipe/qt.git)

## Install therecipe/qt

[therecipe/qt install](https://github.com/therecipe/qt/wiki/Installation)

add $GOPATH/go/bin to environment variables

## Build

> bash
```
go get -t gitlab.com/firerainos/firerain-installer

cd $(go env GOPATH)/src/gitlab.com/firerain-installer

qtrcc desktop ./resources/

qtdeploy build desktop 
```

> fish
```
go get -t gitlab.com/firerainos/firerain-installer

cd (go env GOPATH)/src/gitlab.com/firerain-installer

qtrcc desktop ./resources/

qtdeploy build desktop 

```

## License

goblog is licensed under [GPLv3](LICENSE).