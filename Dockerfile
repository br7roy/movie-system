FROM centos:7

MAINTAINER by Tak(br7roy@gmail.com)

#生成秘钥、公钥
RUN yum install openssh-server -y
RUN echo y|ssh-keygen -t rsa -b 2048 -P '' -f /root/.ssh/id_rsa
RUN cat /root/.ssh/id_rsa.pub > /root/.ssh/authorized_keys
RUN /bin/cp -rf /root/.ssh/id_rsa /etc/ssh/ssh_host_rsa_key
RUN /bin/cp -rf /root/.ssh/id_rsa.pub /etc/ssh/ssh_host_rsa_key.pub

#
RUN yum install -y epel-release
RUN yum install -y docker-io
RUN yum provides '*/applydeltarpm'
RUN yum install deltarpm -y
RUN yum install -y gcc
RUN yum install -y go
RUN yum install git -y

ENV GOROOT /usr/local/go
ENV GOPATH /data/gopath
ENV PATH $GOROOT/bin:$PATH
RUN curl -s -o go.tar.gz https://storage.googleapis.com/golang/go1.15.5.linux-amd64.tar.gz
RUN tar --remove-files -C /usr/local/ -zxf go.tar.gz
RUN mkdir -p /data/go
RUN mkdir -p /data/gopath

#RUN ln -sv /usr/local/go/bin/go /bin
#ADD src/ /data/gopath/src
#ADD pkg/ /data/gopath/pkg

ENV prj_dir=/data/gopath/moviex
RUN git clone https://github.com/br7roy/moviex.git /data/gopath/moviex --branch master --verbose
WORKDIR ${prj_dir}
RUN go build -o bin/moviex movie.go

#定义时区参数
ENV TZ=Asia/Shanghai
#设置时区
RUN ln -snf /usr/share/zoneinfo/$TZ /etc/localtime && echo '$TZ' > /etc/timezone

#将ssh服务启动脚本复制到/usr/local/sbin目录中，并改变权限为755
ADD run.sh /usr/local/sbin/run.sh
RUN chmod 755 /usr/local/sbin/run.sh

#变更root密码为root
RUN echo "root:root"| chpasswd

EXPOSE 22
EXPOSE 8080

CMD ["/usr/local/sbin/run.sh"]
