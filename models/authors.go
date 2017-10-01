package models

import (
	"errors"
	"time"

	"github.com/opiumated/jinPod/config"

	"gopkg.in/mgo.v2/bson"
)

var (
	//InvalidID An Invalid ID
	InvalidID = errors.New("Invalid ID")
)

//Author Represents an Author writing a podcast
type Author struct {
	ID          bson.ObjectId `json:"id" bson:"_id"`
	Name        string        `json:"name" bson:"name"`
	AvatarURL   string        `json:"avatar_url" bson:"avatar_url"`
	Country     string        `json:"country" bson:"country"`
	DateCreated time.Time     `json:"dateCreated" bson:"dateCreated"`
	DateUpdated time.Time     `json:"dateUpdated" bson:"dateUpdated"`
}

//Get Get's a single author based on the authors id
func (a Author) Get(cfg *config.Config, id string) (Author, error) {
	if !bson.IsObjectIdHex(id) {
		return a, InvalidID
	}
	var author Author
	session := cfg.Session.Copy()
	if err := cfg.Database.C(AuthorCollection).Find(bson.M{"_id": id}).One(&author); err != nil {
		return author, err
	}
	defer session.Close()
	return author, nil
}

//GetAll Gets all the authors
func (a Author) GetAll(cfg *config.Config) ([]*Author, error) {
	session := cfg.Session.Copy()
	var authors []*Author
	if err := cfg.Database.C(AuthorCollection).Find(bson.M{}).Sort("-dateCreated").All(&authors); err != nil {
		return authors, err
	}
	defer session.Close()
	return authors, nil
}

//Remove Removes an author from the collection
func (a Author) Remove(cfg *config.Config, authorID string) error {
	if !bson.IsObjectIdHex(authorID) {
		return InvalidID
	}
	session := cfg.Session.Copy()
	defer session.Close()
	if err := cfg.Database.C(AuthorCollection).Remove(bson.M{"authorId": authorID}); err != nil {
		return err
	}
	return nil
}

//Update Updates an author's data based on the given authorID
func (a Author) Update(cfg *config.Config, authorID string, data map[string]string) {

}
