package main

import (
	"election-api/internal/config"
	"github.com/qiangxue/go-rest-api/pkg/log"
)

var Version = "1.0.0"

func main() {
	logger := log.New().With(nil, "version", Version)
}

func buildHandler(logger log.Logger, cfg *config.Config) {

	// TODO we can load our config here

	// TODO we can create our dynamo db instance here

}
