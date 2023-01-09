FROM scratch 
# 载入一个空的初始化环境

WORKDIR $GOPATH/src/github.com/EDDYCJY/go-gin-example
# 把工作目录设置到 这个下面
COPY . $GOPATH/src/github.com/EDDYCJY/go-gin-example
# 复制源代码过去


EXPOSE 8000
# EXPOSE 指令是声明运行时容器提供服务端口，
# 这只是一个声明，在运行时并不会因为这个声明应用就会开启这个端口的服务
ENTRYPOINT ["./go-gin-example"]