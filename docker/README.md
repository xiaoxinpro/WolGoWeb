# wol-go-web

 wol-go-web是WolGoWeb部署再Docker中的项目，主要用于搭建在局域网服务器或NAS的Docker中，用于方便快捷的实现WebAPI唤醒局域网内主机。

 > 在使用该工具前，首先要确认需要唤醒的主机支持WOL功能并且已经开启。

## 服务器Docker-compose部署（推荐）

使用 Docker-compose 可以十分便捷的部署 WolGoWeb 工具，首先要确保服务器中已经安装了 Docker 和 Docker-compose。

然后创建一个 `docker-compose.yml` 文件：

```
version: '3'
services:
  wol-go-web:
    image: chishin/wol-go-web
    container_name: WolGoWeb
    restart: always
    network_mode: host
    environment:
      - PORT=9090
      - KEY=false
```

最后启动容器：

```
docker-compose pull
docker-compose up -d
```

## 服务器Docker部署

使用 Docker 部署 WolGoWeb 工具：

```
docker run -d --net=host chishin/wol-go-web
```

如果需要指定端口可以使用下面的命令：

```
docker run -d --net=host --env PORT=端口号 chishin/wol-go-web
```

### 环境变量说明：

|参数名称|描述|备注
|---|---|---|
|PORT|开放服务端口|默认：9090|
|KEY|API权限验证KEY|默认关闭，详见 [API权限验证说明](https://github.com/xiaoxinpro/WolGoWeb#4api%E6%9D%83%E9%99%90%E9%AA%8C%E8%AF%81)|

## 群晖Docker部署

首先你的群晖必须已经安装好Docker，打开Docker应用，在 **注册表** 中搜索 `wol-go-web`，搜索到下图这个 `chishin/wol-go-web` 右击下载。

![搜索 wol-go-web](https://image.xiaoxin.pro/github/WolGoWeb/%E6%90%9C%E7%B4%A2wol-go-web.png)

下载完成后再 **映像** 中找到 `chishin/wol-go-web` ,双击进行配置。

![配置界面](https://image.xiaoxin.pro/github/WolGoWeb/%E9%85%8D%E7%BD%AE%E7%95%8C%E9%9D%A2.png)

依次点击 **高级设置 → 网络 → 使用与 Docker Host 相同的网络**

![网络配置](https://image.xiaoxin.pro/github/WolGoWeb/%E7%BD%91%E7%BB%9C%E9%85%8D%E7%BD%AE.png)

接下来切换到 **环境** 界面，根据需要设置服务端口

![端口配置](https://image.xiaoxin.pro/github/WolGoWeb/%E7%AB%AF%E5%8F%A3%E9%85%8D%E7%BD%AE.png)

最后点击 **应用** 完成部署工作

![完成部署](https://image.xiaoxin.pro/github/WolGoWeb/%E5%AE%8C%E6%88%90%E9%83%A8%E7%BD%B2.png)

可以看到WolGoWeb的系统占用非常低！

## Github

[https://github.com/xiaoxinpro/WolGoWeb](https://github.com/xiaoxinpro/WolGoWeb)

