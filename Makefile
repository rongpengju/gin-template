# 项目名称
PROJECT_NAME=gin-template

# 替换模板项目名称
# 使用方法：make project REPLACE_STRING="your_project_name"；示例：make project REPLACE_STRING="user-center"
.PHONY: project
REPLACE_STRING ?= gin-template
project:
ifneq ($(PROJECT_NAME),gin-template)
 $(error "该项目有可能不是模板文件 请检查项目状态")
endif
	@echo "开始替换模板中的名称，新项目名称为：$(REPLACE_STRING)"
	@chmod +x replace.sh
	@rm -rf .git
	@rm README.md
	@./replace.sh "$(REPLACE_STRING)"
	@rm replace.sh
	@echo "替换完成，请移至上一级目录查看"
	@echo "注意！！！如果替换失败，请检查是否出现目录重复，或者重新 clone 项目然后再次生成"

# 生成数据库对应的 gorm-gen 代码
.PHONY: orm
orm:
	go run cmd/gen/generate.go

# 获取当前时间戳（用于Docker Image Tag），格式为 yyyymmdd.hhmm
CURRENT_TIME := $(shell date +%Y%m%d.%H%M)

# 打包 docker 镜像
.PHONY: image
image:
	docker build --pull=false -t registry/gin-template:$(CURRENT_TIME) .

# 推送 docker 镜像
.PHONY: push
push: image
	docker push registry/gin-template:$(CURRENT_TIME)
