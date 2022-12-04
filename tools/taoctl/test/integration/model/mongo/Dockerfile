FROM golang:1.18

ENV TZ Asia/Shanghai
ENV GOPROXY https://goproxy.cn,direct

WORKDIR /app
ADD taoctl /usr/bin/taoctl
ADD cmd.sh .

RUN chmod +x /usr/bin/taoctl
RUN chmod +x cmd.sh
CMD ["/bin/bash", "cmd.sh"]
