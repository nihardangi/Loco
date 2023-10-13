package config

/*
EnvConfig type struct
	contains information about environment of application, logger details.
*/
type EnvConfig struct {
	// LogBaseURL holds location of logs base path
	LogBaseURL string
	Env        string
	LogLevel   string
	Logger     string
	// HTTPPort holds port number on which application will listen for http requests
	HTTPPort string
}
