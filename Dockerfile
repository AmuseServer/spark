FROM registry-in.dustess.com:9000/base/centos:7
RUN /bin/cp /usr/share/zoneinfo/Asia/Shanghai /etc/localtime && echo 'Asia/Shanghai' >/etc/timezone
WORKDIR /app
ADD ./go-crm-integration /app
ADD ./pre-stop.sh /app
RUN chmod +x /app/pre-stop.sh
ENV RUN_MODE pro
ENV GIN_MODE release
EXPOSE 8080 50051
CMD ["/app/go-crm-integration"]
