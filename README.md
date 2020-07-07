# WolGoWeb

 WolGoWeb是一款远程唤醒WebAPI工具，主要用于搭建在局域网服务器或NAS中，实现WebAPI唤醒局域网内主机。

 > 在使用该工具前，首先要确认需要唤醒的主机支持WOL功能并且已经开启。

## 开发状态

WolGoWeb 仍然处于开发阶段，未经充分测试与验证，不推荐用于生产环境。

* master 分支用于发布开发版本（稳定性需要进一步测试）
* release 版本为经测试稳定发布的版本（建议下载最新的 release 版本部署）

目前WebAPI可能随时改变，不保证向后兼容，升级是需注意WebAPI是否准确。

## 部署说明

### 1、服务器直接部署

无论是Windows还是Linux系统都可以直接下载对应的 release 编译版本直接运行即可，无需安装任何依赖。

```
WolGoWeb_0.0.1_linux_amd64 -port 9090
```

其中参数 `-port` 表示服务端口号，默认是 `9090` 也可以不填。

> 需要注意的是指定端口必须可以访问，可能需要额外配置防火墙。

### 2、服务器Docker部署

使用 Docker 可以更加便捷的部署 WolGoWeb 工具。

```
docker run -d --net=host chishin/wol-go-web
```

如果需要指定端口可以使用下面的命令：

```
docker run -d --net=host --env PORT=端口号 chishin/wol-go-web
```

### 3、群晖Docker部署

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

## 使用方法

### 1、验证部署成功

完成部署工作即可开始使用，首先使用浏览器访问 `http://服务器IP或域名:9090`，如果修改了端口号请访问对应的端口。

![访问服务地址](https://upload-images.jianshu.io/upload_images/1568014-d02b340a42a433aa.png?imageMogr2/auto-orient/strip%7CimageView2/2/w/1240)

看到以上界面表示服务部署成功。

### 2、发送唤醒请求

可以直接使用浏览器访问 `http://服务器IP或域名:9090/wol?mac=需要唤醒主机的MAC地址` 当出现以下界面表示唤醒命令发送成功。

![发送唤醒请求](https://upload-images.jianshu.io/upload_images/1568014-6d1abfdbb644a986.png?imageMogr2/auto-orient/strip%7CimageView2/2/w/1240)

### 3、唤醒请求参数
|参数名称|描述|备注
|---|---|---|
|mac|唤醒主机的MAC地址|必填|
|ip|唤醒主机的IP地址|默认：255.255.255.255|
|port|唤醒命令发送的端口|默认：9|

## 更多使用实例

不断更新中。。。

