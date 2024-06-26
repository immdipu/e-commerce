package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID             primitive.ObjectID `json:"_id" bson:"_id"`
	FirstName      string             `json:"first_name"`
	LastName       string             `json:"last_name"`
	Password       string             `json:"password" `
	Email          string             `json:"email"`
	Phone          string             `json:"phone"`
	Token          string             `json:"token"`
	RefreshToken   string             `json:"refresh_token"`
	CreatedAt      time.Time          `json:"created_at"`
	UpdatedAt      time.Time          `json:"updated_at"`
	UserCart       []ProductUser      `json:"usercart" bson:"usercart"`
	AddressDetails []Address          `json:"address" bson:"address"`
	OrderStatus    []Order            `json:"orders" bson:"orders"`
}

type Product struct {
	ProductID   primitive.ObjectID `bson:"_id"`
	ProductName *string            `json:"product_name" bson:"product_name"`
	Price       *uint64            `json:"price" bson:"pri"`
	Rating      *uint8             `json:"rating" bson:"rating"`
	Image       *string            `json:"image" bson:"image"`
}

type ProductUser struct {
	ProductID   primitive.ObjectID `json:"product_id"`
	ProductName *string            `json:"product_name"`
	Price       *uint64            `json:"price"`
	Rating      *uint8             `json:"rating"`
	Image       *string            `json:"image"`
}

type Address struct {
	AddressID primitive.ObjectID `bson:"_id"`
	House     *string            `json:"house" bson:"house_name"`
	Street    *string            `json:"street" bson:"street_name"`
	City      *string            `json:"city_name" bson:"city_name" `
	Pincode   *string            `json:"pin_code" bson:"pin_code"`
}

type Order struct {
	OrderID       primitive.ObjectID `bson:"_id"`
	OrderCart     []ProductUser      `json:"order_list" bson:"order_list"`
	OrderedAt     time.Time          `json:"ordered_at" bson:"ordered_at"`
	Price         *uint64            `json:"total_price" bson:"total_price`
	Discount      *int               `json:"discount" bson:"discount"`
	PaymentMethod Payment            `json:"payment_method" bson:"payment_method"`
}

type Payment struct {
	Digital bool `json:"digital"`
	COD     bool `json:"cod"`
}
