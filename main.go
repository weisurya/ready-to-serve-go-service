package main

import (
	"os"

	"github.com/weisurya/ready-to-serve-go-service/config"
	"github.com/weisurya/ready-to-serve-go-service/endpoints"
	"github.com/weisurya/ready-to-serve-go-service/setting"
)

var (
	// Server-related
	port           = os.Getenv("PORT")
	certificateDir = os.Getenv("DIR_CERTIFICATE")
	keyDir         = os.Getenv("DIR_PRIVATE_KEY")

	// Database-related
	dbUser = os.Getenv("DATABASE_USERNAME")
	dbPass = os.Getenv("DATABASE_PASSWORD")
	dbHost = os.Getenv("DATABASE_HOST")
	dbPort = os.Getenv("DATABASE_PORT")
	dbName = os.Getenv("DATABASE_NAME")
)

func main() {
	logger := config.InitiateLog("log-")
	setting.SetLog(logger)

	db, err := config.DatabaseSetting{
		Host:     dbHost,
		Port:     dbPort,
		Name:     dbName,
		Username: dbUser,
		Password: dbPass,
	}.InitiateDatabase()

	if err != nil {
		logger.Fatalf("Failed to connect database: %s", err)
	}
	setting.SetDB(db)

	mux := endpoints.InitiateRouter()

	server := config.ServerSetting{
		Port:         port,
		ReadTimeout:  5,
		WriteTimeout: 10,
		IdleTimeout:  120,
	}.InitiateHTTPSServer(mux)

	if err := server.ListenAndServeTLS(certificateDir, keyDir); err != nil {
		logger.Fatalf("Server failed to start: %s", err)
		os.Exit(1)
	}
}
