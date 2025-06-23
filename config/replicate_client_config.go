package config

import (
	"github.com/replicate/replicate-go"
)

var ReplicateClient *replicate.Client

func InitializeReplicateClient() {
	ReplicateClient, err = replicate.NewClient(replicate.WithTokenFromEnv())
	if err != nil {
		panic(err)
	}
}
