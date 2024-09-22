package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/writeconcern"

	"github.com/jibaru/ctx-transaction/internal/orders/application"
	"github.com/jibaru/ctx-transaction/internal/orders/domain"
	"github.com/jibaru/ctx-transaction/internal/orders/infrastructure/handlers"
	mongo_repo "github.com/jibaru/ctx-transaction/internal/orders/infrastructure/repositories/mongo"
	"github.com/jibaru/ctx-transaction/internal/orders/infrastructure/repositories/mysql"
	mongo_tx "github.com/jibaru/ctx-transaction/internal/shared/infrastructure/transactional/mongo"
	mysql_tx "github.com/jibaru/ctx-transaction/internal/shared/infrastructure/transactional/mysql"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
		return
	}

	var orderRepo domain.OrderRepository
	var orderLineRepo domain.OrderLineRepository
	var transactional application.Transactional

	dbType := os.Getenv("DB")

	if dbType == "mysql" {
		mysqlDB, err := sql.Open(
			"mysql",
			fmt.Sprintf(
				"%s:%s@tcp(%s:%s)/%s?parseTime=true",
				os.Getenv("MYSQL_USER"),
				os.Getenv("MYSQL_PASSWORD"),
				os.Getenv("MYSQL_HOST"),
				os.Getenv("MYSQL_PORT"),
				os.Getenv("MYSQL_NAME"),
			),
		)
		if err != nil {
			log.Fatal(err)
		}

		orderRepo = mysql.NewMySQLOrderRepository(mysqlDB)
		orderLineRepo = mysql.NewMySQLOrderLineRepository(mysqlDB)
		transactional = mysql_tx.NewMySQLTransactional(mysqlDB)
	} else if dbType == "mongo" {
		wc := writeconcern.Majority()
		opts := options.Client().ApplyURI(os.Getenv("MONGO_URI")).SetWriteConcern(wc)
		mongoClient, err := mongo.Connect(context.Background(), opts)
		if err != nil {
			log.Fatal(err)
		}

		mongoDB := mongoClient.Database("orders")
		orderRepo = mongo_repo.NewMongoOrderRepository(mongoDB)
		orderLineRepo = mongo_repo.NewMongoOrderLineRepository(mongoDB)
		transactional = mongo_tx.NewMongoTransactional(mongoClient)
	} else {
		log.Fatal("Unsupported database type. Use 'mysql' or 'mongo'.")
	}

	service := application.NewCreateOrderService(orderRepo, orderLineRepo)
	txService := application.NewCreateOrderServiceTx(service, transactional)

	handler := handlers.NewCreateOrderHandler(txService)

	http.Handle("POST /orders", handler)
	log.Printf("Server running on %s\n", os.Getenv("APP_PORT"))
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", os.Getenv("APP_PORT")), nil))
}
