# v2rayXPlus
![](https://gitlab.com/xiayesuifeng/v2rayxplus/raw/master/resources/v2rayXPlus-64px.svg)

> 使用go写的v2ray gui客户端

[![pipeline status](https://gitlab.com/xiayesuifeng/v2rayxplus/badges/master/pipeline.svg)](https://gitlab.com/xiayesuifeng/v2rayxplus/commits/master)
[![Go Report Card](https://goreportcard.com/badge/gitlab.com/xiayesuifeng/v2rayxplus)](https://goreportcard.com/report/gitlab.com/xiayesuifeng/v2rayxplus)
[![GoDoc](https://godoc.org/gitlab.com/xiayesuifeng/v2rayxplus?status.svg)](https://godoc.org/gitlab.com/xiayesuifeng/v2rayxplus)
[![Sourcegraph](https://sourcegraph.com/gitlab.com/xiayesuifeng/v2rayxplus/-/badge.svg)](https://sourcegraph.com/gitlab.com/xiayesuifeng/v2rayxplus)

使用iptables+透明代理+v2ray路由实现真全局自动分流

感谢[@linyuan](https://t.me/linyuan)提供的logo

## Dependencies
[therecipe/qt](https://github.com/therecipe/qt.git)

## Install v2rayxplus for archlinux
```bash
wget https://mirrors.firerain.me/x86_64/v2rayxplus-1.0.0-1-x86_64.pkg.tar.xz
sudo pacman -U ./v2rayxplus-1.0.0-1-x86_64.pkg.tar.xz
```

## Install v2rayxplus for other GNU/Linux (qt 5.13 +)
```bash
wget https://mirrors.firerain.me/x86_64/v2rayxplus-1.0.0-1-x86_64.pkg.tar.xz
mkdir tmp
cd tmp
tar -xvf ../v2rayxplus-1.0.0-1-x86_64.pkg.tar.xz
rm .*
mv -f ./usr/* /usr
sudo systemctl daemon-reload
```

## Build v2rayxplus install

### 1. Install therecipe/qt

[https://blog.firerain.me/article/6](https://blog.firerain.me/article/6) or
[therecipe/qt install](https://github.com/therecipe/qt/wiki/Installation)

add $GOPATH/go/bin to environment variables

### 2. Build or [v2rayxplus.zip](https://gitlab.com/xiayesuifeng/v2rayxplus/builds/artifacts/master/download?job=run-build) (qt 5.13 +)

> bash
```
go get -t gitlab.com/xiayesuifeng/v2rayxplus

cd $(go env GOPATH)/src/gitlab.com/xiayesuifeng/v2rayxplus

qtrcc desktop ./resources/

qtdeploy build desktop 
```

> fish
```
go get -t gitlab.com/xiayesuifeng/v2rayxplus

cd (go env GOPATH)/src/gitlab.com/xiayesuifeng/v2rayxplus

qtrcc desktop ./resources/

qtdeploy build desktop 

```

# Install
```
sudo wget https://gitlab.com/firerainos/firerain-package/tree/master/v2rayxplus/v2rayxplus.service -O /usr/lib/systemd/system/v2rayxplus.service
sudo wget https://gitlab.com/firerainos/firerain-package/tree/master/v2rayxplus/v2rayxplus.desktop -O /usr/share/applications/v2rayxplus.desktop
sudo wget https://gitlab.com/firerainos/firerain-package/tree/master/v2rayxplus/v2rayxplus.svg -O /usr/share/icons/hicolor/96x96/apps/v2rayxplus.svg
sudo cp ./deploy/linux/v2rayxplus /usr/bin
sudo systemctl daemon-reload
```

## License

v2rayXPlus is licensed under [GPLv3](LICENSE).