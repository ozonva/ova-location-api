package config

type serverConfig struct {
	GrpcPort int
	HttpPort int
}

func getServerConfig() serverConfig {
	return serverConfig{
		HttpPort: getEnvAsInt("SERVER_HTTP_PORT", 8081),
		GrpcPort: getEnvAsInt("SERVER_GRPC_PORT", 8082),
	}
}
