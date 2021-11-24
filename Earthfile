linter-1.13:
    FROM golang:1.13
    DO +LINTER_CMDS

linter-1.14:
    FROM golang:1.14
    DO +LINTER_CMDS

LINTER_CMDS:
    COMMAND
    WORKDIR /mks-go
    COPY . .
    RUN go get github.com/mattn/goveralls; \
        go get github.com/wadey/gocovmerge; \
        curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(go env GOPATH)/bin v1.23.3
    RUN make golangci-lint