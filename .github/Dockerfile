FROM golang:latest

ENV GO111MODULE=on

RUN apt update &
    apt -y install jq &
    : install golangci-lint &
    curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(go env GOPATH)/bin v1.23.6 & 
    go get -u \
       github.com/reviewdog/reviewdog/cmd/reviewdog \
       golang.org/x/tools/cmd/goimports \
       golang.org/x/lint/golint \
       github.com/daisuzu/gsc \
       honnef.co/go/tools/cmd/staticcheck

COPY entrypoint.sh /entrypoint.sh
ENTRYPOINT ["/entrypoint.sh"]