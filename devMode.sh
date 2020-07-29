#!/usr/bin/env bash

# 开发模式 dev/prod
# 此处修改模式
# 执行该脚本
devMode=${1}
#version=1.0

# 模块
# 此处新增模块
modules=(client/cmd server)

# 默认dev
if [[ ${devMode} = '' ]]; then
   devMode=dev
fi

# 配置文件地址
# 修改各个模块下app.yaml文件开发模式
#confFiles=()
# 工作目录
workDir=$(pwd)

# prod 模式自动构建docker
# 可注释, 通过docker.sh执行构建
if [[ ${devMode} = 'dev' ]]; then
    echo "prod 开始构建docker"
#    dor="BEGIN{print \"./docker.sh -v ${version} -m"
    for module in ${modules[*]}
    #也可以写成for element in ${array[@]}
    do
#        dor="${dor} ${module}"
        echo "工作目录: ${workDir}"
        echo -e "\n模块: ""${module} 开始构建docker"
        cd "${workDir}"/"${module}" && ./docker.sh
    done
fi