package config

func RunConfig() {
	InitializeEnv()
	InitializeDatabase()
	InitializeOpenAIClient()
	InitializeReplicateClient()
	InitTurnStileClient()
}
