# wol-go-web

 wol-go-web是WolGoWeb部署再Docker中的项目，主要用于搭建在局域网服务器或NAS的Docker中，用于方便快捷的实现WebAPI唤醒局域网内主机。

 > 在使用该工具前，首先要确认需要唤醒的主机支持WOL功能并且已经开启。

## 服务器Docker部署

使用 Docker 可以更加便捷的部署 WolGoWeb 工具。

```
docker run -d --net=host chishin/wol-go-web
```

如果需要指定端口可以使用下面的命令：

```
docker run -d --net=host --env PORT=端口号 chishin/wol-go-web
```

## 群晖Docker部署

首先你的群晖必须已经安装好Docker，打开Docker应用，在 **注册表** 中搜索 `wol-go-web`，搜索到下图这个 `chishin/wol-go-web` 右击下载。

![搜索 wol-go-web](https://upload-images.jianshu.io/upload_images/1568014-c462aca21a3dc594.png?imageMogr2/auto-orient/strip%7CimageView2/2/w/1240)

下载完成后再 **映像** 中找到 `chishin/wol-go-web` ,双击进行配置。

![配置界面](https://upload-images.jianshu.io/upload_images/1568014-c06273400c1a8eed.png?imageMogr2/auto-orient/strip%7CimageView2/2/w/1240)

依次点击 **高级设置 → 网络 → 使用与 Docker Host 相同的网络**

![网络配置](https://upload-images.jianshu.io/upload_images/1568014-d9332400d2fdc946.png?imageMogr2/auto-orient/strip%7CimageView2/2/w/1240)

接下来切换到 **环境** 界面，根据需要设置服务端口

![端口配置](https://upload-images.jianshu.io/upload_images/1568014-66b09b6547ca8ed6.png?imageMogr2/auto-orient/strip%7CimageView2/2/w/1240)

最后点击 **应用** 完成部署工作

![完成部署](https://upload-images.jianshu.io/upload_images/1568014-9353dee69d237ce5.png?imageMogr2/auto-orient/strip%7CimageView2/2/w/1240)

可以看到WolGoWeb的系统占用非常低！

