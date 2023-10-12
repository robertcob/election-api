package main

import "github.com/qiangxue/go-rest-api/pkg/log"

var Version = "1.0.0"

func main() {

	logger := log.New().With(nil, "version", Version)
}
