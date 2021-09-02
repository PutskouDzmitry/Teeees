package cmd

import (
	"os"

	"github.com/tonymontanapaffpaff/timescale-postgres-load-testing/pkg/api"
	"github.com/tonymontanapaffpaff/timescale-postgres-load-testing/pkg/data"
	"github.com/tonymontanapaffpaff/timescale-postgres-load-testing/pkg/db"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

var (
	dbPort     = os.Getenv("DB_PORT")
	dbUser     = os.Getenv("DB_USER")
	dbName     = os.Getenv("DB_NAME")
	dbPassword = os.Getenv("DB_PASSWORD")
	sslMode    = os.Getenv("SSL_MODE")
)

func init() {
	if dbPort == "" {
		dbPort = "5432"
	}
	if dbUser == "" {
		dbUser = "postgres"
	}
	if dbName == "" {
		dbName = "test"
	}
	if dbPassword == "" {
		dbPassword = "root"
	}
	if sslMode == "" {
		sslMode = "disable"
	}
}

func start(dbHost string) {
	// Creates database connection
	conn, err := db.GetConnection(dbHost, dbPort, dbUser, dbName, dbPassword, sslMode)
	if err != nil {
		log.Fatalf("Can't connect to database, error: %v", err)
	}

	// Creates a default gin router
	r := gin.Default()
	visitData := data.NewVisitData(conn)
	api.ServeVisitResource(r, *visitData)

	// Listen and serve on 0.0.0.0:8080
	err = r.Run(":8080")
	if err != nil {
		log.Fatal("Server has been crashed...")
	}
}
