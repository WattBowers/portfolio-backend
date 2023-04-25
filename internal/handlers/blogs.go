package handlers

import (
	"context"
	"fmt"
	"portfolio-backend/internal/db"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Blog struct {
	ID        primitive.ObjectID `json:"_id" bson:"_id" validate:"required"`
	CreatedAt time.Time          `json:"createdAt" bson:"created_at" validate:"required"`
	UpdatedAt time.Time          `json:"updatedAt" bson:"updated_at" validate:"required"`
	Title     string             `json:"title" bson:"title" validate:"required,min=3"`
	Content   string             `json:"content" bson:"content" validate:"required,min=12"`
}

type ErrorResponse struct {
	FailedField string
	Tag         string
	Value       string
}

func ValidateBlogStruct(p Blog) []*ErrorResponse {
	var errors []*ErrorResponse
	validate := validator.New()
	err := validate.Struct(p)

	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			var element ErrorResponse
			element.FailedField = err.StructNamespace()
			element.Tag = err.Tag()
			element.Value = err.Param()
			errors = append(errors, &element)
		}
	}

	return errors
}

func CreateBlog(c *fiber.Ctx) error {
	blog := Blog{
		ID:        primitive.NewObjectID(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	fmt.Println(blog)
	if err := c.BodyParser(&blog); err != nil {
		return err
	}

	errors := ValidateBlogStruct(blog)

	if errors != nil {
		return c.JSON(errors)
	}
	fmt.Println(blog)
	client, err := db.GetMongoClient()

	if err != nil {
		return err
	}
	fmt.Println(blog)
	collection := client.Database(db.Database).Collection(string(db.BlogsCollection))

	_, err = collection.InsertOne(context.TODO(), blog)

	if err != nil {
		return err
	}
	fmt.Println(blog)
	return c.JSON(blog)
}

func GetAllBlogs(c *fiber.Ctx) error {
	client, err := db.GetMongoClient()

	var blogs []*Blog

	if err != nil {
		return err
	}

	collection := client.Database(db.Database).Collection(string(db.BlogsCollection))

	cur, err := collection.Find(context.TODO(), bson.D{
		primitive.E{},
	})

	if err != nil {
		return err
	}

	for cur.Next(context.TODO()) {
		var p Blog
		err := cur.Decode(&p)

		if err != nil {
			return err
		}

		blogs = append(blogs, &p)
	}

	return c.JSON(blogs)
}
