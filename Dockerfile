FROM golang:latest

ENV GO111MODULE=on

RUN apt update &
    apt -y install jq &
    go get -u github.com/reviewdog/reviewdog/cmd/reviewdog

COPY entrypoint.sh /entrypoint.sh
ENTRYPOINT ["/entrypoint.sh"]