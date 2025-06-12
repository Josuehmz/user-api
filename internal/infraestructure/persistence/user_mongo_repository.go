package persistence

import (
	"context"
	"errors"
	"time"

	"github.com/Josuehmz/user-api/internal/domain"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// userMongoRepository implementa la interfaz UserRepository usando MongoDB.
type userMongoRepository struct {
	collection *mongo.Collection
}

// NewUserMongoRepository crea un nuevo repositorio de usuarios con MongoDB.
func NewUserMongoRepository(col *mongo.Collection) domain.UserRepository {
	return &userMongoRepository{collection: col}
}

// Create inserta un nuevo usuario en la colección.
func (r *userMongoRepository) Create(ctx context.Context, user *domain.User) error {
	user.CreatedAt = time.Now()
	user.UpdatedAt = time.Now()

	result, err := r.collection.InsertOne(ctx, user)
	if err != nil {
		return err
	}

	id := result.InsertedID.(primitive.ObjectID)
	user.ID = id.Hex()
	return nil
}

// GetAll retorna todos los usuarios de la colección.
func (r *userMongoRepository) GetAll(ctx context.Context) ([]domain.User, error) {
	cursor, err := r.collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var users []domain.User
	for cursor.Next(ctx) {
		var raw bson.M
		if err := cursor.Decode(&raw); err != nil {
			return nil, err
		}

		objID, ok := raw["_id"].(primitive.ObjectID)
		if !ok {
			return nil, errors.New("invalid _id format")
		}

		user := domain.User{
			ID:        objID.Hex(),
			Name:      raw["name"].(string),
			Email:     raw["email"].(string),
			CreatedAt: raw["createdat"].(primitive.DateTime).Time(),
			UpdatedAt: raw["updatedat"].(primitive.DateTime).Time(),
		}

		users = append(users, user)
	}

	return users, nil
}

// GetByID busca un usuario por su ID en MongoDB.
func (r *userMongoRepository) GetByID(ctx context.Context, id string) (*domain.User, error) {
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	var user domain.User
	if err := r.collection.FindOne(ctx, bson.M{"_id": objID}).Decode(&user); err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, nil
		}
		return nil, err
	}

	user.ID = objID.Hex()
	return &user, nil
}

// Update reemplaza todos los campos de un usuario por su ID.
func (r *userMongoRepository) Update(ctx context.Context, user *domain.User) error {
	objID, err := primitive.ObjectIDFromHex(user.ID)
	if err != nil {
		return err
	}

	user.UpdatedAt = time.Now()

	_, err = r.collection.UpdateOne(ctx, bson.M{"_id": objID}, bson.M{
		"$set": bson.M{
			"name":      user.Name,
			"email":     user.Email,
			"updatedat": user.UpdatedAt,
		},
	})

	return err
}

// Delete elimina un usuario de la colección por su ID.
func (r *userMongoRepository) Delete(ctx context.Context, id string) error {
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	_, err = r.collection.DeleteOne(ctx, bson.M{"_id": objID})
	return err
}

// UpdatePartial actualiza solo campos específicos de un usuario.
func (r *userMongoRepository) UpdatePartial(ctx context.Context, id string, fields map[string]interface{}) error {
	if len(fields) == 0 {
		return nil
	}

	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	fields["updatedat"] = time.Now()

	_, err = r.collection.UpdateOne(ctx, bson.M{"_id": objID}, bson.M{"$set": fields})
	return err
}
