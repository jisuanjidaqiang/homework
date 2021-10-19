FROM centos
ADD bin/amd64/main /main
EXPOSE 80
ENTRYPOINT /main