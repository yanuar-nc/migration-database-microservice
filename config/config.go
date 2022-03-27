package config

import (
	"errors"
	"os"
	"strconv"

	dotenv "github.com/joho/godotenv"
)

var (

	// Development env checking, this env for debug purpose
	Development string

	// HTTPPort config
	HTTPPort uint16

	// GRPCPort config
	GRPCPort string

	// WriteDBHost config
	WriteDBHost string
	// WriteDBName config
	WriteDBName string
	// WriteDBUser config
	WriteDBUser string
	// WriteDBPassword config
	WriteDBPassword string
	// WriteDBPort config
	WriteDBPort int

	// ReadDBHost config
	ReadDBHost string
	// ReadDBName config
	ReadDBName string
	// ReadDBUser config
	ReadDBUser string
	// ReadDBPassword config
	ReadDBPassword string
	// ReadDBPort config
	ReadDBPort int
)

// Load function will load all config from environment variable
func Load() error {
	// load .env
	err := dotenv.Load(".env")
	if err != nil {
		return errors.New(".env is not loaded properly")
	}

	development, ok := os.LookupEnv("DEVELOPMENT")
	if !ok {
		return errors.New("DEVELOPMENT env is not loaded")
	}

	// set Development
	Development = development

	// ------------------------------------

	httpPortStr, ok := os.LookupEnv("HTTP_PORT")
	if !ok {
		return errors.New("HTTP_PORT env is not loaded")
	}

	httpPort, err := strconv.Atoi(httpPortStr)
	if err != nil {
		return errors.New("HTTP_PORT env is not valid")
	}

	// set http port
	HTTPPort = uint16(httpPort)

	writeDBHost, ok := os.LookupEnv("WRITE_DB_HOST")
	if !ok {
		return errors.New("WRITE_DB_HOST env is not loaded")
	}

	// set WriteDBHost
	WriteDBHost = writeDBHost

	writeDBName, ok := os.LookupEnv("WRITE_DB_NAME")
	if !ok {
		return errors.New("WRITE_DB_NAME env is not loaded")
	}

	// set WriteDBName
	WriteDBName = writeDBName

	writeDBUser, ok := os.LookupEnv("WRITE_DB_USER")
	if !ok {
		return errors.New("WRITE_DB_USER env is not loaded")
	}

	// set WriteDBUser
	WriteDBUser = writeDBUser

	writeDBPassword, ok := os.LookupEnv("WRITE_DB_PASSWORD")
	if !ok {
		return errors.New("WRITE_DB_PASSWORD env is not loaded")
	}

	// set WriteDBPassword
	WriteDBPassword = writeDBPassword

	writeDBPort, ok := os.LookupEnv("WRITE_DB_PORT")
	if !ok {
		return errors.New("WRITE_DB_PORT env is not loaded")
	}

	// set WriteDBPort
	WriteDBPort, _ = strconv.Atoi(writeDBPort)
	// ------------------------------------
	return nil
}
