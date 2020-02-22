FROM golang:latest

ENV GO111MODULE=on

RUN go get -u \
       github.com/reviewdog/reviewdog/cmd/reviewdog \
       golang.org/x/tools/cmd/goimports \
       golang.org/x/lint/golint \
       github.com/daisuzu/gsc \
       honnef.co/go/tools/cmd/staticcheck

COPY entrypoint.sh /entrypoint.sh
ENTRYPOINT ["/entrypoint.sh"]