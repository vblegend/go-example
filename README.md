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

## ✨ 性能分析调试 pprof
页面地址 `http://localhost:8000/debug`
在生产环境下访问该页面需要输入以下账号密码
账号 `gogogo`
密码 `123456`
