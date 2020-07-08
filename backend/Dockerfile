FROM golang

# for go get
env GOPROXY https://goproxy.io

ADD . /backend

# build
RUN cd /backend/cmd && go build

ENTRYPOINT ["/backend/cmd/cmd"]
