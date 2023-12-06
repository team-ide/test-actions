FROM golang:1.19.4 as go-1.19
FROM docker:20.10.12 as docker

FROM centos:7

ENV GO_HOME /usr/local/go-1.19
ENV PATH $GO_HOME/bin:$PATH
COPY --from=go-1.19 /usr/local/go $GO_HOME

RUN go env -w GOPROXY=https://goproxy.cn,direct

# 修改时区
RUN cp -f /usr/share/zoneinfo/Asia/Shanghai /etc/localtime
RUN yum -y install kde-l10n-Chinese
RUN localedef -c -f UTF-8 -i zh_CN zh_CN.utf8
RUN echo 'LANG="zh_CN.UTF-8"'>/etc/locale.conf

WORKDIR /home
ENV LANG zh_CN.UTF-8
ENV LC_ALL zh_CN.UTF-8

RUN yum remove -y unixODBC unixODBC-devel
# 升级odbc
RUN rpm -Uvh https://packages.microsoft.com/config/rhel/7/packages-microsoft-prod.rpm
RUN yum install -y gcc gcc-c++ unixODBC unixODBC-devel wget
RUN yum install -y zip unzip lrzsz vim* net-tools.x86_64 net-tools tree tcpdump
RUN yum install -y make openssl-devel curl-devel expat-devel gettext-devel zlib-devel asciidoc perl-ExtUtils-MakeMaker
RUN yum clean all

COPY --from=docker /usr/local/bin/docker /usr/local/bin/docker

