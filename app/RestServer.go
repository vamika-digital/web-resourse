package app

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/spf13/viper"
	"github.com/vamika-digital/wms-lib/logger"
	"github.com/vamika-digital/wms-resourse/app/handlers/rest"
	"github.com/vamika-digital/wms-resourse/config"
	"net/http"
)

func sanityCheck() {
	serverEnvProps := []string{
		"servers.api_server.address",
		"servers.api_server.port",
		"rdbms.type",
	}
	for _, k := range serverEnvProps {
		if viper.GetString(k) == "" {
			logger.Fatal(fmt.Sprintf("Environment variable %s not defined. Terminating application...", k))
		}
	}
	dbEnvProps := map[string][]string{
		"sqlite3": {
			"dbPath",
		},
		"mysql": {
			"host",
			"port",
			"username",
			"password",
			"dbname",
			"connMaxLifetime",
			"maxOpenConns",
			"maxIdleConns",
		},
		"pgsql": {
			"host",
			"port",
			"username",
			"password",
			"dbname",
			"connMaxLifetime",
			"maxOpenConns",
			"maxIdleConns",
		},
	}
	dbType := viper.GetString("rdbms.type")
	if dbTypeProps, exists := dbEnvProps[dbType]; exists {
		for _, k := range dbTypeProps {
			if viper.GetString(fmt.Sprintf("rdbms.%s.%s", dbType, k)) == "" {
				logger.Fatal(fmt.Sprintf("Environment variable %s not defined for %s. Terminating application...", k, dbType))
			}
		}
	} else {
		logger.Fatal(fmt.Sprintf("No configuration details found for database driver: %s. Terminating application...", dbType))
	}
}

func StartRestServer() {
	sanityCheck()
	router := mux.NewRouter()
	dbClient := config.GetDBInstance()

	productFamilyHander := rest.ProductFamilyHandlers{}

	//wiring
	//ch := CustomerHandlers{service.NewCustomerService(domain.NewCustomerRepositoryStub())}
	//customerRepositoryDb := domainNewCustomerRepositoryDb(dbClient)
	//accountRepositoryDb := domain.NewAccountRepositoryDb(dbClient)
	//ch := CustomerHandlers{service.NewCustomerService(customerRepositoryDb)}
	//ah := AccountHandler{service.NewAccountService(accountRepositoryDb)}

	// define routes
	//router.
	//	HandleFunc("/customers", ch.getAllCustomers).
	//	Methods(http.MethodGet).
	//	Name("GetAllCustomers")
	//router.
	//	HandleFunc("/customers/{customer_id:[0-9]+}", ch.getCustomer).
	//	Methods(http.MethodGet).
	//	Name("GetCustomer")
	//router.
	//	HandleFunc("/customers/{customer_id:[0-9]+}/account", ah.NewAccount).
	//	Methods(http.MethodPost).
	//	Name("NewAccount")
	//router.
	//	HandleFunc("/customers/{customer_id:[0-9]+}/account/{account_id:[0-9]+}", ah.MakeTransaction).
	//	Methods(http.MethodPost).
	//	Name("NewTransaction")
	//
	//am := AuthMiddleware{domain.NewAuthRepository()}
	//router.Use(am.authorizationHandler())
	// starting server

	address := viper.GetString("servers.api_server.address")
	port := viper.GetString("servers.api_server.port")
	logger.Info(fmt.Sprintf("Starting server on %s:%s ...", address, port))
	err := http.ListenAndServe(fmt.Sprintf("%s:%s", address, port), router)
	if err != nil {
		logger.Fatal(fmt.Sprintf("Failed to listen and start the server: %s", err))
	}

}

//func getDbClient() *sqlx.DB {
//	dbUser := os.Getenv("DB_USER")
//	dbPasswd := os.Getenv("DB_PASSWD")
//	dbAddr := os.Getenv("DB_ADDR")
//	dbPort := os.Getenv("DB_PORT")
//	dbName := os.Getenv("DB_NAME")
//
//	dataSource := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbUser, dbPasswd, dbAddr, dbPort, dbName)
//	client, err := sqlx.Open("mysql", dataSource)
//	if err != nil {
//		panic(err)
//	}
//	// See "Important settings" section.
//	client.SetConnMaxLifetime(time.Minute * 3)
//	client.SetMaxOpenConns(10)
//	client.SetMaxIdleConns(10)
//	return client
//}
