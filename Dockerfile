FROM centos:7

MAINTAINER by Tak(br7roy@gmail.com)

#
RUN yum install -y epel-release
RUN yum install -y docker-io
RUN yum provides '*/applydeltarpm'
RUN yum install deltarpm -y
RUN yum install -y gcc
RUN yum install -y go
RUN yum install git -y

#
ENV GOROOT /usr/local/go
ENV GOPATH /data/gopath
ENV PATH $GOROOT/bin:$PATH
RUN curl -s -o go.tar.gz https://storage.googleapis.com/golang/go1.15.5.linux-amd64.tar.gz
RUN tar --remove-files -C /usr/local/ -zxf go.tar.gz
RUN mkdir -p /data/go
RUN mkdir -p /data/gopath


RUN git clone https://github.com/br7roy/moviex.git /data/gopath/moviex --branch master --verbose
WORKDIR /data/gopath/moviex
RUN go build -o bin/moviex movie.go
RUN chmod 755 bin/moviex

#
ENV TZ=Asia/Shanghai
#
RUN ln -snf /usr/share/zoneinfo/$TZ /etc/localtime && echo '$TZ' > /etc/timezone


EXPOSE 8080

CMD ["/data/gopath/moviex/bin/moviex"]
