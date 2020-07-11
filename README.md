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

#### 运行参数说明：

|参数名称|描述|备注
|---|---|---|
|-port|开放服务端口|默认：9090|
|-key|API权限验证KEY|默认关闭，详见[API权限验证说明](https://github.com/xiaoxinpro/WolGoWeb#4、API权限验证)|

### 2、服务器Docker部署

使用 Docker 可以更加便捷的部署 WolGoWeb 工具。

```
docker run -d --net=host chishin/wol-go-web
```

如果需要指定端口可以使用下面的命令：

```
docker run -d --net=host --env PORT=端口号 chishin/wol-go-web
```

#### 环境说明：

|参数名称|描述|备注
|---|---|---|
|PORT|开放服务端口|默认：9090|
|KEY|API权限验证KEY|默认关闭，详见 [API权限验证说明](https://github.com/xiaoxinpro/WolGoWeb#4、API权限验证)|

### 3、群晖Docker部署

群晖系统可以在Docker应用的 **注册表** 中搜索 `wol-go-web`，即可下载和部署项目。

更多图文教程可以参考 https://github.com/xiaoxinpro/WolGoWeb/blob/master/docker/README.md

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

### 4、API权限验证

API权限验证用于防止他人触发唤醒指令的发送，是一种唤醒指令安全措施，默认处于关闭状态。

#### 开启API权限验证

再启动软件是传入 `-key` 参数或者Docker增加环境变量 `KEY`，此参数的长度必须大于等于6个字符，否则API权限验证将处于关闭状态。

#### 权限验证请求参数

开启API权限验证后，在发送WOL唤醒请求时必须和唤醒请求一起发送一下参数。

|参数名称|描述|备注
|---|---|---|
|time|发送请求时的时间戳|单位：秒|
|token|经过加密后得到的权限Token|token=MD5(key+mac+time)|

例如：设定的`key=123456`，发送请求时的 `time=1594452205`, `mac=00-00-00-00-00-00`，计算token的公式为`MD5("12345600-00-00-00-00-001594452205")`,结果为`token=eb3515003672b3e0324196ecd78438a2`

#### 特别注意

* 对于参数time必须不能小于接收时刻30秒以上，同时也不能大于接收时刻的时间戳。
* 对于多次发送相同mac的唤醒请求time值不允许相同。
* 对于token参数长度必须为32，并且英文字符必须是小写的。
* 对于key长度必须大于6个字符，否则不会进行权限验证。

## 应用实例实例

### 1、使用iOS快捷指令唤醒（Siri唤醒）

可以自己创建一个快捷指令访问唤醒的URL即可，也可以直接在iOS浏览器中打开下面的链接修改成你的服务器地址和需要唤醒的MAC地址。

[https://www.icloud.com/shortcuts/0931d2a9d4e84984b8d85e977aff8ef9](https://www.icloud.com/shortcuts/0931d2a9d4e84984b8d85e977aff8ef9)

![快捷指令](https://upload-images.jianshu.io/upload_images/1568014-9304d5e3506cd536.png?imageMogr2/auto-orient/strip%7CimageView2/2/w/1240)

创建完成快捷指令后可以在快捷指令主页用点击 **唤醒电脑** ，或者语音唤醒Siri说出 **唤醒电脑** 即可完成电脑唤醒。

