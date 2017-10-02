package models

import (
	"time"

	"github.com/opiumated/jinPod/config"

	"gopkg.in/mgo.v2/bson"
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

//Add Adds an Author to our collection
func (a Author) Add(cfg *config.Config) error {
	session := cfg.Session.Copy()
	defer session.Close()
	if err := cfg.Database.C(AuthorCollection).Insert(a); err != nil {
		return err
	}
	return nil
}

//Get Get's a single author based on the authors ID
func (a Author) Get(cfg *config.Config, id string) (Author, error) {
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

//GetByName Get's an author through it's name
func (a Author) GetByName(cfg *config.Config, name string) (Author, error) {
	session := cfg.Session.Copy()
	defer session.Close()

	var author Author
	if err := cfg.Database.C(AuthorCollection).Find(bson.M{"name": name}).One(&author); err != nil {
		return author, err
	}
	return author, nil
}

//Remove Removes an author from the collection
func (a Author) Remove(cfg *config.Config, authorID string) error {
	session := cfg.Session.Copy()
	defer session.Close()
	if err := cfg.Database.C(AuthorCollection).Remove(bson.M{"_id": authorID}); err != nil {
		return err
	}
	return nil
}

//Update Updates an author's data based on the given authorID
func (a Author) Update(cfg *config.Config, authorID string, data map[string]string) {

}
