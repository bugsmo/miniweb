# Systemd 配置、安装和启动

1. [Systemd 配置、安装和启动](#systemd-配置安装和启动)
	1. [1. 前置操作](#1-前置操作)
	2. [2. 创建 miniweb systemd unit 模板文件](#2-创建-miniweb-systemd-unit-模板文件)
	3. [3. 复制 systemd unit 模板文件到 sysmted 配置目录](#3-复制-systemd-unit-模板文件到-sysmted-配置目录)
	4. [4. 启动 systemd 服务](#4-启动-systemd-服务)

## 1. 前置操作

1. 创建需要的目录 

```bash
sudo mkdir -p /data/miniweb /opt/miniweb/bin /etc/miniweb /var/log/miniweb
```

2. 编译构建 `miniweb` 二进制文件

```bash
make build # 编译源码生成 miniweb 二进制文件
```

3. 将 `miniweb` 可执行文件安装在 `bin` 目录下

```bash
sudo cp _output/platforms/linux/amd64/miniweb /opt/miniweb/bin # 安装二进制文件
```

4. 安装 `miniweb` 配置文件

```bash
sed 's/.\/_output/\/etc\/miniweb/g' configs/miniweb.yaml > miniweb.sed.yaml # 替换 CA 文件路径
sudo cp miniweb.sed.yaml /etc/miniweb/ # 安装配置文件
```

5. 安装 CA 文件

```bash
make ca # 创建 CA 文件
sudo cp -a _output/cert/ /etc/miniweb/ # 将 CA 文件复制到 miniweb 配置文件目录
```

## 2. 创建 miniweb systemd unit 模板文件

执行如下 shell 脚本生成 `miniweb.service.template`

```bash
cat > miniweb.service.template <<EOF
[Unit]
Description=APIServer for web platform.
Documentation=https://github.com/bugsmo/miniweb/blob/master/init/README.md

[Service]
WorkingDirectory=/data/miniweb
ExecStartPre=/usr/bin/mkdir -p /data/miniweb
ExecStartPre=/usr/bin/mkdir -p /var/log/miniweb
ExecStart=/opt/miniweb/bin/miniweb --config=/etc/miniweb/miniweb.yaml
Restart=always
RestartSec=5
StartLimitInterval=0

[Install]
WantedBy=multi-user.target
EOF
```

## 3. 复制 systemd unit 模板文件到 sysmted 配置目录

```bash
sudo cp miniweb.service.template /etc/systemd/system/miniweb.service
```

## 4. 启动 systemd 服务

```bash
sudo systemctl daemon-reload && systemctl enable miniweb && systemctl restart miniweb
```
