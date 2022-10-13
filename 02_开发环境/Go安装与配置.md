# 安装

### Mac 下安装

可以通过 brew 方式安装，也可以直接在官网下载可执行文件，然后双击安装包，不停下一步就可以了


### Linux 下安装

下载安装包：

wget https://golang.google.cn/dl/go1.16.6.linux-amd64.tar.gz
解压到 /usr/local 目录：

sudo tar -zxvf go1.16.6.linux-amd64.tar.gz -C /usr/local

# 配置

打开 $HOME/.bash_profile 文件，增加下面两行代码：

```
export GOROOT=/usr/local/go
export PATH=$PATH:$GOROOT/bin
```

最后使环境变量生效：

source $HOME/.bash_profile
安装完成后，在终端执行查看版本命令，如果能正确输出版本信息，那就说明安装成功了。

go version
go version go1.16.6 linux/amd64