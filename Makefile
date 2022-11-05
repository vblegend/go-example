.PHONY: web publish archive key


#============================================================
# make 说明
#============================================================
# make release   - 构建发布的版本至$(pwd)/${PUBLISH_DIR} 目录
# make linux     - 编译linux二进制文件
# make win       - 编译windows二进制文件
# make web       - 编译web项目
# make tidy      - 执行go mod tidy
# make clean     - 移除已编译的二进制文件
# make archive   - 将发布的文件夹归档至tar
#============================================================




APP_FILENAME	:= go-example
BUILD_VERSION   := 1.1.1
BUILD_TIME      := $(shell date "+%Y-%m-%d %H:%M:%S")
COMMIT_ID		:= $(shell git rev-parse --short HEAD)
OUT_DIR			:= ./publish
PUBLISH_DIR     := ${OUT_DIR}/${APP_FILENAME}


define clean
	@echo "正在清理目录.."
	@rm -rf ${OUT_DIR}
	@mkdir -p ${PUBLISH_DIR}
	@if [ -f ./server/${APP_FILENAME} ] ; then rm ./server/${APP_FILENAME}; fi
	@if [ -f ./server/${APP_FILENAME}.exe ] ; then rm ./server/${APP_FILENAME}.exe; fi
endef

define depend
	@echo "正在恢复golang项目依赖.."
	@cd server; go mod tidy
	@echo "正在恢复web项目依赖.."
	@cd web; npm install --force
endef


define buildWeb
	@echo "正在编译Web项目..."
	@cd ./web; npm run build:prod
endef

define buildPlugin
	@echo "正在编译项目插件..."
	@cd plugins/pprof; CGO_ENABLED=1 GOOS=linux GOARCH=amd64 go build  -o pprof.so -buildmode=plugin ./main.go

endef
# -race
define buildLinux
	@echo "正在编译后端项目..."
	@echo "当前版本：${BUILD_VERSION}"
	@echo "构建时间：${BUILD_TIME}"
	@echo "提交记录：${COMMIT_ID}"
	@cd server; CGO_ENABLED=1 GOOS=linux GOARCH=amd64 go build -ldflags "-X 'server/common/assembly.Version=${BUILD_VERSION}' -X 'server/common/assembly.BuildTime=${BUILD_TIME}' -X 'server/common/assembly.CommitID=${COMMIT_ID}'" -o ./${APP_FILENAME}  ./main.go
	@echo "编译完毕..."
endef

define buildWindows
	@echo "正在编译Windows架构可执行文件..."
	@CGO_ENABLED=1 GOOS=windows GOARCH=amd64 go build $(RACE) -o ./${APP_FILENAME}.exe ./server/main.go
	@echo "编译完毕..."
endef

define archive
	@cd publish; tar -cvf ${APP_FILENAME}.tar ${APP_FILENAME}
endef

define tidy
	@cd server; go mod tidy
endef


define publish
	@echo "正在发布文件..."
	@/bin/cp -rf ./server/${APP_FILENAME} ${PUBLISH_DIR}/${APP_FILENAME}
	@chmod 777 ${PUBLISH_DIR}/${APP_FILENAME}
	@/bin/cp -rf ./server/bin ${PUBLISH_DIR}/bin
	@/bin/cp -rf ./server/config ${PUBLISH_DIR}/config
	@/bin/cp -rf ./server/static ${PUBLISH_DIR}/static
	@/bin/cp -rf ./server/template ${PUBLISH_DIR}/template
	@/bin/cp -rf ./server/version ${PUBLISH_DIR}/version
	@rm -rf ${PUBLISH_DIR}/bin/windows
	@rm -rf ${PUBLISH_DIR}/static/form-generator
	@${PUBLISH_DIR}/${APP_FILENAME} config reset
endef


#server.pfx
define makePem
	openssl pkcs12 -in server.pfx -nocerts -out key.pem -nodes
	openssl pkcs12 -in server.pfx -nokeys -out server.pem
	openssl rsa -in key.pem -out server.key
endef

 #生成 rsa 私钥/公钥
KEY_PASSWORD    :=123456
define makeKey
	rm -rf private.pem
	rm -rf public.pem
	ssh-keygen -t rsa -f private.pem -m pem -P "${KEY_PASSWORD}"
    openssl rsa -in private.pem -pubout -out public.pem -passin pass:${KEY_PASSWORD}
	rm -rf private.pem.pub
endef



#发布文件一条龙服务
release:
	@echo "正在准备发布..."
	@$(clean)
	@$(depend)
	@$(buildWeb)
	@$(buildLinux)
	@$(publish)
	@$(archive)
	@echo "====================================================="
	@echo "文件已发布至${PUBLISH_DIR}目录..."
	@echo "====================================================="



#仅编译linux版本服务
linux:
	@$(buildLinux)

#仅编译windows版本服务
windows:
	@$(buildWindows)

#清理已编译的软件和目录
clean:
	@$(clean)

#仅编译web项目
web:
	@$(buildWeb)

#安装项目所有依赖项
depend:
	@$(depend)

#仅处理golang依赖项
tidy:
	@$(tidy)

#发布已编译文件至输出目录
publish:
	@$(publish)

#将发布的文件夹归档至tar包
archive:
	@$(archive)

#将发布的文件夹归档至tar包
plugin:
	@$(buildPlugin)


key:
	@$(makeKey)