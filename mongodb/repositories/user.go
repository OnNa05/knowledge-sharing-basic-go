package repositories

import (
	"context"

	"github.com/OnNa05/knowledge-sharing-basic-go/mongodb/connection"
	"github.com/OnNa05/knowledge-sharing-basic-go/mongodb/entities"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserRepo struct {
	Collection *mongo.Collection
}

type IUserRepo interface {
	Insert(ctx context.Context, value *entities.UserEntitty) (*mongo.InsertOneResult, error)
	FindAll(ctx context.Context) ([]entities.UserEntitty, error)
	FindByEmail(ctx context.Context, email string) (*entities.UserEntitty, error)
	FindByID(ctx context.Context, id string) (*entities.UserEntitty, error)
	FindByEmailAndPassword(ctx context.Context, email string, password string) (*entities.UserEntitty, error)
	UpdateTokenByEmailAndPassword(ctx context.Context, ent *entities.UserEntitty) error
}

func NewUserRepo(ds *connection.MongoDB) IUserRepo {
	return &UserRepo{
		Collection: ds.MongoDB.Database("webstorage").Collection("user_storage"),
	}
}

func (r *UserRepo) FindAll(ctx context.Context) ([]entities.UserEntitty, error) {
	var res []entities.UserEntitty

	c, err := r.Collection.Find(ctx, bson.M{})
	if err != nil {
		return res, err
	}

	err = c.All(ctx, &res)
	return res, err
}

func (r *UserRepo) Insert(ctx context.Context, value *entities.UserEntitty) (*mongo.InsertOneResult, error) {
	return r.Collection.InsertOne(ctx, value)
}

func (r *UserRepo) FindByEmail(ctx context.Context, email string) (*entities.UserEntitty, error) {
	var res *entities.UserEntitty
	err := r.Collection.FindOne(ctx, bson.M{"email": email}).Decode(&res)

	return res, err
}

func (r *UserRepo) FindByEmailAndPassword(ctx context.Context, email string, password string) (*entities.UserEntitty, error) {
	var res *entities.UserEntitty
	err := r.Collection.FindOne(ctx, bson.M{"email": email, "password": password}).Decode(&res)

	return res, err
}

func (r *UserRepo) UpdateTokenByEmailAndPassword(ctx context.Context, ent *entities.UserEntitty) error {
	_, err := r.Collection.UpdateOne(ctx, bson.M{"email": ent.Email, "password": ent.Password}, bson.M{"$set": bson.M{"access_token": ent.AccessToken}})
	return err
}

func (r *UserRepo) FindByID(ctx context.Context, id string) (*entities.UserEntitty, error) {
	var res *entities.UserEntitty

	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return res, err
	}

	err = r.Collection.FindOne(ctx, bson.M{"_id": oid}).Decode(&res)

	return res, err
}
