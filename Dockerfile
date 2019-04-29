FROM centos
MAINTAINER "vincent321x@gmail.com"
WORKDIR /var/local/XcessAlipay/service
RUN mkdir /etc/xcesspay
COPY configuration /etc/xcesspay
COPY . .
EXPOSE 28888
ENTRYPOINT ["/bin/bash", ".main"]