FROM centos
MAINTAINER "vincent321x@gmail.com"
WORKDIR /var/local/XcessAlipay
RUN mkdir /etc/xcesspay
COPY configuration /etc/xcesspay
COPY . .
EXPOSE 28888
ENTRYPOINT ["/bin/bash","/var/local/XcessAlipay/service/main"]