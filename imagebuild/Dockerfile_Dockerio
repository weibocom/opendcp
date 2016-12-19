FROM docker:dind

RUN apk add --no-cache bash git subversion openssh-client ca-certificates

# 程序目录
COPY code /imagebuild/code
COPY globle_config /imagebuild/globle_config
RUN ["/bin/sh", "/imagebuild/code/delete_gofile.sh"]
RUN mkdir -p /imagebuild/project

# 端口
EXPOSE 8080

# 工作目录
WORKDIR "/imagebuild/code/web"

# entrypoint
ENTRYPOINT ["/bin/sh", "/imagebuild/code/entrypoint.sh"]

