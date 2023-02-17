package model

import (
	"encoding/json"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Restaurant struct {
	ID          primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Name        string             `json:"name,omitempty" bson:"name,omitempty"`
	Address     string             `json:"address,omitempty" bson:"address,omitempty"`
	City        string             `json:"city,omitempty" bson:"city,omitempty"`
	PhoneNumber string             `json:"_phoneNumber,omitempty" bson:"_phoneNumber,omitempty"`
	IsOpened    bool               `json:"_isOpened,omitempty" bson:"_isOpened,omitempty"`
}

func (r Restaurant) MarshalJSON() ([]byte, error) {
	type Alias Restaurant
	return json.Marshal(&struct {
		ID string `json:"_id,omitempty"`
		*Alias
	}{
		ID:    r.ID.Hex(),
		Alias: (*Alias)(&r),
	})
}
