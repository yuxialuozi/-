# 检查 go 命令是否存在于 PATH 环境变量中
if (-not (Get-Command go -ErrorAction SilentlyContinue)) {
    Write-Host "错误：go 命令未在 PATH 中找到。请安装或将其添加到 PATH 中。"
    exit 1
}

# 检查 protoc 命令是否存在于 PATH 环境变量中
if (-not (Get-Command protoc -ErrorAction SilentlyContinue)) {
    Write-Host "错误：protoc 命令未在 PATH 中找到。请安装或将其添加到 PATH 中。"
    exit 1
}

# 检查 kitex 命令是否存在于 PATH 环境变量中
if (-not (Get-Command kitex -ErrorAction SilentlyContinue)) {
    Write-Host "错误：kitex 命令未在 PATH 中找到，尝试安装..."
    go install github.com/cloudwego/kitex/tool/cmd/kitex@latest
}

# 创建生成目录
New-Item -ItemType Directory -Force -Path kitex_gen

# 运行 kitex 生成代码
kitex -module "simpledouyin" -I idl/ "./idl/$($args[0]).proto"

# 创建 service 目录
New-Item -ItemType Directory -Force -Path "./service./$($args[0])"

# 进入 service 目录并运行 kitex
Set-Location "service\$($args[0])"
kitex -module "simpledouyin" -service $args[0] -use simpledouyin/kitex_gen/ -I ../../idl/ "../../idl\$($args[0]).proto"

# 执行 go mod tidy 来整理依赖
go mod tidy
