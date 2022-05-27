package main

import (
	"flag"
	"fmt"
	"os"

	stackdriver "github.com/TV4/logrus-stackdriver-formatter"
	"github.com/Thospol/go-learning/docs"
	"github.com/Thospol/go-learning/internal/core/config"
	"github.com/Thospol/go-learning/internal/core/sql"
	"github.com/Thospol/go-learning/internal/handlers/routes"
	"github.com/sirupsen/logrus"
)

func main() {
	configs := flag.String("config", "configs", "set configs path, default as: 'configs'")
	flag.Parse()

	// Init configuration
	err := config.InitConfig(*configs)
	if err != nil {
		panic(err)
	}
	//=======================================================
	port := os.Getenv("PORT")
	if port != "" {
		config.CF.App.Port = port
	}

	// programatically set swagger info
	docs.SwaggerInfo.Title = config.CF.Swagger.Title
	docs.SwaggerInfo.Description = config.CF.Swagger.Description
	docs.SwaggerInfo.Version = config.CF.Swagger.Version
	docs.SwaggerInfo.Host = fmt.Sprintf("%s%s", config.CF.Swagger.Host, config.CF.Swagger.BaseURL)
	//=======================================================

	// set logrus
	logrus.SetReportCaller(true)
	if config.CF.App.Release {
		logrus.SetFormatter(stackdriver.NewFormatter(
			stackdriver.WithService("api"),
			stackdriver.WithVersion("v1.0.0")))
	} else {
		logrus.SetFormatter(&logrus.TextFormatter{})
	}
	logrus.Infof("Initial 'Configuration'. %+v", config.CF)
	//=======================================================

	// Init return result
	err = config.InitReturnResult("configs")
	if err != nil {
		panic(err)
	}
	//=======================================================

	// Init connection mysql
	configuration := sql.Configuration{
		Host:     config.CF.SQL.Host,
		Port:     config.CF.SQL.Port,
		Username: config.CF.SQL.Username,
		Password: config.CF.SQL.Password,
	}
	configuration.DatabaseName = config.CF.SQL.DatabaseName
	session, err := sql.InitConnectionMysql(configuration)
	if err != nil {
		panic(err)
	}
	sql.Database = session.Database

	// New router
	routes.NewRouter()
	//========================================================
}
