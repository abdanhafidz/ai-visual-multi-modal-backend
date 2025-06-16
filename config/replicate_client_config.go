package config

import (
	"github.com/replicate/replicate-go"
)

var ReplicateClient *replicate.Client

func init() {
	ReplicateClient, err = replicate.NewClient(replicate.WithTokenFromEnv())
	if err != nil {
		panic(err)
	}
}
