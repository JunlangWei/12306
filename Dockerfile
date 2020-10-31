FROM ubuntu:18.04

ENV LANG C.UTF-8

RUN echo "">/etc/apt/sources.list \
    && echo "deb http://mirrors.aliyun.com/ubuntu/ bionic main restricted universe multiverse">>/etc/apt/sources.list \
    && echo "deb http://mirrors.aliyun.com/ubuntu/ bionic-security main restricted universe multiverse">>/etc/apt/sources.list \
    && echo "deb http://mirrors.aliyun.com/ubuntu/ bionic-updates main restricted universe multiverse">>/etc/apt/sources.list \
    && echo "deb http://mirrors.aliyun.com/ubuntu/ bionic-proposed main restricted universe multiverse">>/etc/apt/sources.list \
    && echo "deb http://mirrors.aliyun.com/ubuntu/ bionic-backports main restricted universe multiverse">>/etc/apt/sources.list \
    && echo "deb-src http://mirrors.aliyun.com/ubuntu/ bionic main restricted universe multiverse">>/etc/apt/sources.list \
    && echo "deb-src http://mirrors.aliyun.com/ubuntu/ bionic-security main restricted universe multiverse">>/etc/apt/sources.list \
    && echo "deb-src http://mirrors.aliyun.com/ubuntu/ bionic-updates main restricted universe multiverse">>/etc/apt/sources.list \
    && echo "deb-src http://mirrors.aliyun.com/ubuntu/ bionic-proposed main restricted universe multiverse">>/etc/apt/sources.list \
    && echo "deb-src http://mirrors.aliyun.com/ubuntu/ bionic-backports main restricted universe multiverse">>/etc/apt/sources.list \
    && apt-get clean && apt-get update

RUN DEBIAN_FRONTEND=noninteractive \
    TZ=Asia/Shanghai \
    apt-get install -y vim && \
    apt-get install -y wget && \
    apt-get install -y dpkg && \
    apt-get install -y unzip && \
    apt-get install -y apt-utils && \
    apt-get install -y golang-go && \
    apt-get autoremove -y

ENTRYPOINT [ "go", "run", "/src/main.go" ]