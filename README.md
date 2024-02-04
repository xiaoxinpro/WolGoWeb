# WolGoWeb

 WolGoWeb是一款远程唤醒WebAPI工具，主要用于搭建在局域网服务器或NAS中，实现WebAPI唤醒局域网内主机。

 > 在使用该工具前，首先要确认需要唤醒的主机支持WOL功能并且已经开启。

## 开发状态

[![build docker image](https://github.com/xiaoxinpro/WolGoWeb/actions/workflows/docker-image.yml/badge.svg?branch=master)](https://hub.docker.com/r/chishin/wol-go-web)

WolGoWeb 经历了五年的测试已经在诸多测试、生产环境得以稳定运行。

* [master](https://github.com/xiaoxinpro/WolGoWeb/tree/master) 分支用于发布开发版本（稳定性需要进一步测试）
* [release](https://github.com/xiaoxinpro/WolGoWeb/releases) 版本为经测试稳定发布的版本（建议下载最新的 release 版本部署）

在生产环境中建议使用Docker或release版本来部署WolGoWeb。

## 部署说明

### 1、服务器直接部署

无论是Windows还是Linux系统都可以直接下载对应的 release 编译版本直接运行即可，无需安装任何依赖。

```
WolGoWeb_linux_amd64 -port 9090
```

其中参数 `-port` 表示服务端口号，默认是 `9090` 也可以不填。

> 需要注意的是指定端口必须可以访问，可能需要额外配置防火墙。

#### 运行参数说明：

| 参数名称      | 描述          | 备注                                                                                                               |
|-----------|-------------|------------------------------------------------------------------------------------------------------------------|
| -c        | 设置配置源       | `default`命令行、`env`环境变量，默认：default                                                                                |
| -port     | 开放服务端口      | 默认：9090                                                                                                          |
| -web      | 是否启用Web页面   | 默认：`true`                                                                                                        |
| -username | 设置Web页面登陆账号 | 仅在启用Web页面时，且`-username`与`-password`都不为空时有效                                                                       |
| -password | 设置Web页面登陆密码 | 仅在启用Web页面时，且`-username`与`-password`都不为空时有效                                                                       |
| -key      | API权限验证KEY  | 默认：`false`不进行权限验证，详见[API权限验证说明](https://github.com/xiaoxinpro/WolGoWeb#4api%E6%9D%83%E9%99%90%E9%AA%8C%E8%AF%81) |

### 2、服务器Docker-compose部署（推荐）

使用 Docker-compose 可以十分便捷的部署 WolGoWeb 工具，首先要确保服务器中已经安装了 Docker 和 Docker-compose。

创建一个 `docker-compose.yml` 文件：

```yaml
version: '3'
services:
  wol-go-web:
    image: chishin/wol-go-web:latest
    container_name: WolGoWeb
    restart: unless-stopped
    network_mode: host
    environment:
      - PORT=9090
      - KEY=false
```

启动容器：

```bash
docker-compose pull
docker-compose up -d
```

到此部署已经完成，如果需要升级到最新版本，可直接执行以下命令：

```bash
docker-compose down
docker-compose pull
docker-compose up -d
```

### 3、服务器Docker部署

使用 Docker 部署 WolGoWeb 工具：

```
docker run -d --net=host chishin/wol-go-web
```

如果需要指定端口可以使用下面的命令：

```
docker run -d --net=host --env PORT=端口号 chishin/wol-go-web
```

#### 环境说明：

| 参数名称     | 描述          | 备注                                                                                                                |
|----------|-------------|-------------------------------------------------------------------------------------------------------------------|
| PORT     | 开放服务端口      | 默认：9090                                                                                                           |
| WEB      | 是否启用Web页面   | 默认：`true`                                                                                                         |
| USERNAME | 设置Web页面登陆账号 | 仅在启用Web页面时，且`USERNAME`与`PASSWORD`都不为空时有效                                                                          |
| PASSWORD | 设置Web页面登陆密码 | 仅在启用Web页面时，且`USERNAME`与`PASSWORD`都不为空时有效                                                                          |
| KEY      | API权限验证KEY  | 默认：`false`不进行权限验证，详见 [API权限验证说明](https://github.com/xiaoxinpro/WolGoWeb#4api%E6%9D%83%E9%99%90%E9%AA%8C%E8%AF%81) |

### 4、群晖Docker部署

群晖系统可以在Docker应用的 **注册表** 中搜索 `wol-go-web`，即可下载和部署项目。

更多图文教程可以参考 https://github.com/xiaoxinpro/WolGoWeb/blob/master/docker/README.md

## 使用方法

### 1、验证部署成功

完成部署工作即可开始使用，首先使用浏览器访问 `http://服务器IP或域名:9090`，如果修改了端口号请访问对应的端口。

![访问服务地址](https://image.xiaoxin.pro/github/WolGoWeb/%E8%AE%BF%E9%97%AE%E6%9C%8D%E5%8A%A1%E5%9C%B0%E5%9D%80.PNG)

看到以上界面表示服务部署成功。

### 2、发送唤醒请求

可以直接使用浏览器访问 `http://服务器IP或域名:9090/wol?mac=需要唤醒主机的MAC地址` 当出现以下界面表示唤醒命令发送成功。

![发送唤醒请求](https://image.xiaoxin.pro/github/WolGoWeb/%E5%8F%91%E9%80%81%E5%94%A4%E9%86%92%E8%AF%B7%E6%B1%82.PNG)

### 3、唤醒请求参数
| 参数名称 | 描述         | 备注                 |
|------|------------|--------------------|
| mac  | 唤醒主机的MAC地址 | 必填                 |
| ip   | 唤醒主机的IP地址  | 默认：255.255.255.255 |
| port | 唤醒命令发送的端口  | 默认：9               |

### 4、API权限验证

API权限验证用于防止他人触发唤醒指令的发送，是一种唤醒指令安全措施，默认处于关闭状态。

#### 开启API权限验证

在启动项目时传入 `-key` 参数或者Docker增加环境变量 `KEY`，此参数的长度必须大于等于6个字符，否则API权限验证将处于关闭状态。

#### 权限验证请求参数

开启API权限验证后，在发送WOL唤醒请求时必须和唤醒请求一起发送一下参数。

| 参数名称  | 描述              | 备注                      |
|-------|-----------------|-------------------------|
| time  | 发送请求时的时间戳       | 单位：秒                    |
| token | 经过加密后得到的权限Token | token=MD5(key+mac+time) |

例如：设定的`key=123456`，发送请求时的 `time=1594452205`, `mac=00-00-00-00-00-00`，计算token的公式为`MD5("12345600-00-00-00-00-001594452205")`,结果为`token=eb3515003672b3e0324196ecd78438a2`

#### 特别注意

* 对于参数time必须不能小于接收时刻30秒以上，同时也不能大于接收时刻的时间戳。
* 对于多次发送相同mac的唤醒请求time值不允许相同。
* 对于token参数长度必须为32，并且英文字符必须是小写的。
* 对于key长度必须大于6个字符，否则不会进行权限验证。

## 应用实例

### 1、使用iOS快捷指令唤醒（Siri唤醒）

可以自己创建一个快捷指令访问唤醒的URL即可，也可以直接在iOS浏览器中打开下面的链接修改成你的服务器地址和需要唤醒的MAC地址。

[https://www.icloud.com/shortcuts/0931d2a9d4e84984b8d85e977aff8ef9](https://www.icloud.com/shortcuts/0931d2a9d4e84984b8d85e977aff8ef9)

![快捷指令](https://image.xiaoxin.pro/github/WolGoWeb/%E5%BF%AB%E6%8D%B7%E6%8C%87%E4%BB%A4.PNG)

创建完成快捷指令后可以在快捷指令主页用点击 **唤醒电脑** ，或者语音唤醒Siri说出 **唤醒电脑** 即可完成电脑唤醒。

### 2、群晖定时唤醒主机

首先要确保在群晖中已经部署了WolGoWeb，可以访问群晖的IP地址:9090查看是否部署完成。

接下来在群晖里找到`控制面板`中的`任务计划`，新增一个计划任务`用户定义脚本`；

在`常规`界面中随意填写一个任务名称，在`计划`界面中设定好时间；

在`任务设置`界面输入以下唤醒命令，其中`00-00-00-00-00-00`为你要唤醒的主机MAC地址。

```bash
curl http://127.0.0.1:9090/wol?mac=00-00-00-00-00-00
```

如果你的WolGoWeb不是部署在群晖中，需要将上面命令中的`127.0.0.1:9090`替换成你部署的IP和端口。

![1653360092485.png](https://image.xiaoxin.pro/2022/05/24/57e70aa26da8a.png)

最后点击`确定`按钮保存任务。

> 在任务列表里找到刚刚创建的任务，右击`运行`可以立即唤醒主机。

### 3、浏览器收藏夹快捷唤醒

可以在电脑或手机等任意浏览器中创建一个收藏夹或书签，名称随意填写，地址填入：

```
http://192.168.10.10:9090/wol?mac=00-00-00-00-00-00
```

其中`192.168.10.10:9090`是你部署的WolGoWeb，`00-00-00-00-00-00`为你要唤醒的主机MAC地址。

需要唤醒时，直接在收藏夹或书签中点击即可执行唤醒动作。