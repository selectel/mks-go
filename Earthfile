linters-1.16:
    FROM golang:1.16
    DO +LINTER_CMDS

linters-1.17:
    FROM golang:1.17
    DO +LINTER_CMDS

tests-1.16:
    FROM golang:1.16
    DO +TEST_CMDS

tests-1.17:
    FROM golang:1.17
    DO +TEST_CMDS

TEST_CMDS:
    COMMAND
    WORKDIR /mks-go
    COPY . .
    RUN go get github.com/mattn/goveralls; \
        go get github.com/wadey/gocovmerge
    RUN make unittest

LINTER_CMDS:
    COMMAND
    WORKDIR /mks-go
    COPY . .
    RUN curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(go env GOPATH)/bin v1.43.0
    RUN make golangci-lint
