#!/bin/bash

shellExit()
{
if [ $1 -eq 1 ]; then
    printf "\nfailed!!!\n\n"
    exit 1
fi
}

printf "\n Regenerating file \n\n"
time go run -v ./cmd/mysqlmd/main.go  -addr $1 -user $2 -pass $3 -name $4 -tables $5
shellExit $?

printf "\n create curd code : \n"
time go build -o gormgen ./cmd/gormgen/main.go
shellExit $?

if [ ! -d $GOPATH/bin ];then
   mkdir -p $GOPATH/bin
fi

mv gormgen $GOPATH/bin
shellExit $?

go generate ./...
shellExit $?

printf "\n Formatting code \n\n"
time go run -v ./cmd/mfmt/main.go
shellExit $?

printf "\n Done. \n\n"
