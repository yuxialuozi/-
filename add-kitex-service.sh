#!/bin/bash

# 检查 go 命令是否存在于 PATH 环境变量中
if ! command -v go &>/dev/null; then
  echo "错误：go 命令未在 PATH 中找到。请安装或将其添加到 PATH 中。"
  exit 1
fi

# 检查 protoc 命令是否存在于 PATH 环境变量中
if ! command -v protoc &>/dev/null; then
  echo "错误：protoc 命令未在 PATH 中找到。请安装或将其添加到 PATH 中。"
  exit 1
fi

# 检查 kitex 命令是否存在于 PATH 环境变量中
if ! command -v kitex &>/dev/null; then
  echo "错误：kitex 命令未在 PATH 中找到，尝试安装..."
  go install github.com/cloudwego/kitex/tool/cmd/kitex@latest
fi

# 再次检查 kitex 命令是否存在于 PATH 环境变量中
if ! command -v kitex &>/dev/null; then
  echo "错误：kitex 命令未在 PATH 中找到，看起来安装失败了。请手动安装。"
  exit 1
fi

# 创建生成目录
mkdir -p kitex_gen

# 运行 kitex 生成代码
kitex -module "simpledouyin" -I idl/ "idl/$1.proto"

# 创建 service 目录
mkdir -p "service/$1"

# 进入 service 目录并运行 kitex
cd "service/$1" && kitex -module "simpledouyin" -service "$1" -use simpledouyin/kitex_gen/ -I ../../idl/ "../../idl/$1.proto"

# 执行 go mod tidy 来整理依赖
go mod tidy
