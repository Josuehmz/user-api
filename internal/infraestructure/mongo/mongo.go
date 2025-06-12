package mongo

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// ConnectMongo establece una conexión con MongoDB usando la URI proporcionada.
// Retorna una instancia del cliente o un error si falla la conexión.
func ConnectMongo(mongoURI string) (*mongo.Client, error) {
	// Configura las opciones del cliente con la URI
	clientOpts := options.Client().ApplyURI(mongoURI)

	// Crea una nueva instancia del cliente MongoDB
	client, err := mongo.NewClient(clientOpts)
	if err != nil {
		return nil, err
	}

	// Define un contexto con timeout para la conexión
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Conecta el cliente al servidor
	if err := client.Connect(ctx); err != nil {
		return nil, err
	}

	// Verifica la conexión con un ping
	if err := client.Ping(ctx, nil); err != nil {
		return nil, err
	}

	return client, nil
}
