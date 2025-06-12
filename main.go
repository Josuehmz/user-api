package main

import (
	"log"
	"os"

	"github.com/Josuehmz/user-api/internal/domain"
	"github.com/Josuehmz/user-api/internal/handler"
	"github.com/Josuehmz/user-api/internal/infraestructure/mongo"
	"github.com/Josuehmz/user-api/internal/infraestructure/persistence"
	"github.com/joho/godotenv"

	"github.com/Josuehmz/user-api/internal/router"
	"github.com/Josuehmz/user-api/internal/usecase"

	_ "github.com/Josuehmz/user-api/docs"
)

// @title           User API
// @version         1.0
// @description     API RESTful para gestión de usuarios
// @host            localhost:8080
// @BasePath        /
func main() {

	if err := godotenv.Load(); err != nil {
		log.Println("No se pudo cargar el archivo .env, usando valores por defecto")
	}

	mongoURI := os.Getenv("MONGO_URI")
	if mongoURI == "" {
		mongoURI = "mongodb://localhost:27017"
	}

	dbName := os.Getenv("DB_NAME")
	if dbName == "" {
		dbName = "userdb"
	}

	collectionName := os.Getenv("DB_COLLECTION")
	if collectionName == "" {
		collectionName = "users"
	}

	client, err := mongo.ConnectMongo(mongoURI)
	if err != nil {
		log.Fatalf("Error conectando a Mongo: %v", err)
	}
	log.Printf("Conectado a MongoDB en %s", mongoURI)

	db := client.Database(dbName)
	userCollection := db.Collection(collectionName)

	var userRepo domain.UserRepository = persistence.NewUserMongoRepository(userCollection)

	userUC := usecase.NewUserUseCase(userRepo)

	userHandler := handler.NewUserHandler(userUC)

	r := router.NewRouter(userHandler)
	log.Println("Servidor iniciado en http://localhost:8080")
	log.Println("Swagger UI disponible en http://localhost:8080/swagger/index.html")

	if err := r.Run(":8080"); err != nil {
		log.Fatalf("Error iniciando servidor: %v", err)
	}
}
