package infrastructure

import "emerson-argueta/m/v2/env"

// GlobalConfig creates a new Config struct.
var GlobalConfig = newConfig()

// Config is a struct that contains configuration variables.
type Config struct {
	Environment    string
	HTTPServer     *HTTPServer
	Database       *Database
	PaymentGateway *PaymentGateway
	Authorization  *Authorization
}

// Database is a struct that contains the DB's configuration variables.
type Database struct {
	Host     string
	Port     string
	User     string
	DB       string
	Schema   string
	Password string
	Dialect  string
}

// HTTPServer contains configuration variables for the http server.
type HTTPServer struct {
	Host       string
	Port       string
	APIBaseURL string
}

// PaymentGateway contains configuration variables for the payment gateway client.
type PaymentGateway struct {
	APIKey string
}

// Authorization contains configuration variables for the auth client.
type Authorization struct {
	APIKey         string
	Secret         string
	PrivateKeyPath string
	PublicKeyPath  string
}

func newConfig() *Config {
	env.CheckDotEnv()
	httpServer := &HTTPServer{
		Port:       env.MustGet("PORT"),
		Host:       env.MustGet("HOST"),
		APIBaseURL: env.MustGet("API_BASE_URL"),
	}

	database := &Database{
		Host:     env.Get("DB_HOST"),
		Port:     env.Get("DB_PORT"),
		User:     env.Get("DB_USER"),
		DB:       env.Get("DB_DB"),
		Schema:   env.Get("DB_SCHEMA"),
		Password: env.Get("DB_PASSWORD"),
		Dialect:  env.Get("DB_DIALECT"),
	}

	paymentGateway := &PaymentGateway{
		APIKey: env.Get("PAYMENT_GATEWAY_API_KEY"),
	}
	authorization := &Authorization{
		APIKey:         env.Get("AUTH_API_KEY"),
		Secret:         env.Get("SECRET"),
		PrivateKeyPath: env.Get("PRIVATE_KEY_PATH"),
		PublicKeyPath:  env.Get("PUBLIC_KEY_PATH"),
	}

	return &Config{
		Environment:    env.MustGet("ENV"),
		HTTPServer:     httpServer,
		Database:       database,
		PaymentGateway: paymentGateway,
		Authorization:  authorization,
	}
}
