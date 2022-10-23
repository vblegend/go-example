
# App Sidecar

服务器管理台，siteweb6平台看门狗，监控服务器资源装及平台各服务的资源状态，自动重启异常服务
- 服务器资源监控
- S6服务资源监控
- 定时任务处理
- S6平台日志收集与展示
- 数据备份/恢复功能
- 定时数据备份功能
- License 授权申请文件导出

#### ✨ 拉取代码
```
git clone https://gitlab.com/vertiv-co/apac/s6-service-monitoring/application/siteweb6-appsidecar.git
```
#### ✨ 前端
用`vscode` 打开 `./siteweb6-appsidecar/web` 目录
``` bash
# 安装npm依赖包
npm i
# 调试 端口9527
npm run dev
# 编译
npm run build:prod
```

- 调试模式下 后端IP存放于 `.env.development`文件中的`VUE_APP_BASE_API`选项
- 发布模式下 应删除或注释 `.env.production` 文件中的`VUE_APP_BASE_API`选项




#### 💎 后端
提前安装好 golang开发环境 并配置好 `GOPATH` `GOROOT` 环境变量。
用`vscode` 打开 `./siteweb6-appsidecar/backend` 目录
- #### 方法一 命令行编译法
  - 1.更新go代码依赖 执行命令`go mod tidy`
  - 2.编译 在backend 目录 执行 `make clean linux`命令 生成项目  也可以调用命令行`go build -o siteweb-manager`
  - 3.执行命令`siteweb-manager server -c ./config/settings.yml`启动服务器管理台


- #### 方法二 VSCode法
  - 更新go代码依赖 执行命令`go mod tidy`
  - 1.`vscode` 增加 下方`GOLANG DEBUG`脚本,在vscode运行和调试中选中 `Launch Go` F5运行,启动服务器管理台


#### 🤝 访问页面
- 服务器管理台发布时是前后端聚合的，前端页面复用了后端端口8000
- 当前端调试时请使用 http://127.0.0.1:9527 访问
- 当通过后端 http://127.0.0.1:8000 访问时，需要在前端web目录执行`npm run build:prod`编译前端代码

默认账号：`admin`
默认密码：`adminadmin`


#### ⚠️ 使用说明
- 前端执行 `npm run build:prod` 命令时将执行编译脚本将编译好的文件放入`backend/static/www`目录
- 后端新拉下来的代码是没有包含 static/www目录的，所以无法访问页面，需要使用npm命令编译前端即可调试。





#### 附录 GOLANG DEBUG

.vscode 目录下 新增launch.json 写入如下内容
Launch Go 调试运行
环境变量 `RUN_MODE`    表示强制运行模式  dev prod 
环境变量 `DOCKER_ENV`  表示docker环境，为空时连接本机docker守护，填写环境地址连接目标环境
``` json 
{
    "version": "0.2.0",
    "configurations": [
        {
            "name": "Launch Go",
            "type": "go",
            "request": "launch",
            "mode": "auto",
            "program": "main.go",
            "args": [
                "server",
                "-c",
                "./config/settings.yml"
            ],
            "env": {
                "GIN_MODE": "release",
                "RUN_MODE" : "dev",
                "DOCKER_ENV": "http://10.163.100.128:2375"
            }
        }
    ]
}
```





#### 📦 发布编译

- 安装make工具 
- 在backend目录执行 `make clean linux` 命令即可，或在backend目录执行 `go build -o siteweb-manager` 命令




# CLI 打包指令 

## 备份包
通过下方指令对输入目录`/root/backup`执行打包备份包，生成备份文件`/root/out.gzip`
- -b 表示构建备份包
- -i 表示输入目录
- -o 表示输出文件
- 输入目录内必须包含`files`目录或`tables`目录
- `files`目录为文件系统的/根目录。
- `tables`目录下仅支持.sql文件（小写），且文件名由`{库名}.{表名}.sql`或 `{库名}.sql` 组成

``` bash
    siteweb-manager build -b -i /root/backup -o /root/out.gzip
```


## 升级包

通过使用下方指令对目录进行打包为zip升级包

``` bash
    siteweb-manager.exe build -u -i ./root  -o ./out/upgrade.package.zip -t 哎，就是玩er  -d 此次更新了XXX + xxx
```
siteweb-manager.exe `build` 参数如下所示
- -u 表示构建为升级包
- -i 表示输入文件目录
- -o 表示输出文件名称
- -t 表示升级包标题名
- -d 表示升级包说明

输入文件目录下存放着每个打包项目的文件夹（比如redis一个文件夹、backend一个文件夹、frontend 一个文件夹）打包时将按照文件夹名字进行排序打包步骤,除排序外文件名本身并无其他用途。
项目目录下应存在 `project.yml`文件，若无此文件则不会处理该目录的打包流程，`project.yml`文件内容如下所示
``` yaml
Project: redis
Type: container
Version: 1.7.5
Tag: notest
BuildDate: "2022-01-01 22:00:00"
Images:
    - ./rmu
    - ./rmu.dll
    - ./rmu.so
    - ./lib/network.so
Deploys:
    - ./config/application.properties:/siteweb/redis/config/application.properties
    - ./config/application-prod.properties:/siteweb/redis/config/application-prod.properties
    - ./config/logback-spring.xml:/siteweb/redis/config/logback-spring.xml
Command: /siteweb/rmu/rmu service
```

|字段|说明|必须|
|-|-|-|
|Project|服务名、项目名（sql、bash等项目）|
|Type|项目类型，`container`、`service`、`sql`、`bash` 其中一个|
|Version|项目版本 一般用于说明服务的版本（`sql`、`bash`项目未使用）|
|Tag|指定项目标签  `notest`、`release`、`debug`|
|BuildDate|项目打包日期（一般指服务tar包打包日期）|
|Images|需要`Sha256`校验的文件列表（当前目录的相对路径）|
|Deploys|需要部署的文件目录列表  `文件目录`:`部署目录`，文件目录为当前目录的相对路径，部署目录为目标主机的绝对目录|
|Command|当Type == service 时，此项为 服务的启动命令|


`container`类型项目，项目目录下应包含以下文件
{$Project}.tar      `服务的tar包，文件名应该与Project一致`
docker-compose.yml  `服务docker-compose 配置`        
changelog.md        `软件更新日志`

`service`类型项目，项目目录下应包含以下文件
changelog.md        `软件更新日志`

`sql`与`bash`类型项目，项目目录下应仅包含一个 sh或sql文件



##服务控制 

- 执行`siteweb-manager service install ` 命令将会把服务器管理台注册进系统服务随机启动，同时将 siteweb-manager 注册为全局命令 
- 执行`siteweb-manager service uninstall ` 命令将会把siteweb-manager从服务中删除，同时删除全局命令 siteweb-manager
- 执行`siteweb-manager service start` 命令 会启动服务器管理台服务
- 执行`siteweb-manager service restart` 命令 会重新启动服务器管理台服务
- 执行`siteweb-manager service stop` 命令 会停止服务器管理台服务
- 执行`siteweb-manager service status` 命令 查看服务器管理台状态
- 执行`siteweb-manager service logs` 命令 查看服务器管理台最新日志


#### ✨ 性能分析调试 pprof
页面地址 `http://localhost:8000/debug`
在生产环境下访问该页面需要输入以下账号密码
账号 `siteweb`
密码 `123456`