package models

import (
	"log"
	"time"

	"github.com/opiumated/jinPod/config"
	"gopkg.in/mgo.v2/bson"
)

//Podcasts  Represents a podcast in the database
type Podcasts struct {
	ID          bson.ObjectId `json:"_id,omitempty" bson:"_id"`
	Title       string        `json:"title" bson:"title"`
	Slug        string        `json:"slug" bson:"slug"`
	Description string        `json:"description" bson:"description"`
	Body        string        `json:"body" bson:"body"`
	Likes       int           `json:"likes"`
	Author      Author        `json:",omitempty" bson:",omitempty"`
	PodcastsURL string        `json:"podcast_url" bson:"podcast_url"`
	DateCreated time.Time     `json:"dateCreated" bson:"dateCreated"`
	DateUpdated time.Time     `json:"dateUpdated" bson:"dateUpdated"`
}

type podCastLikes struct {
	PodcastID bson.ObjectId
	Count     int
}

//Get Gets a Podcast by ID
func (p Podcasts) Get(cfg *config.Config, podcastID bson.ObjectId) (Podcasts, error) {
	session := cfg.Session.Copy()
	defer session.Close()
	var podCast Podcasts

	if err := cfg.Database.C(PodcastCollection).Find(bson.M{"_id": podcastID}).One(&podCast); err != nil {
		log.Fatalf("%v\n", err.Error())
		return podCast, err
	}
	return podCast, nil
}

//GetBySlug Get's a podcast by it's slug
func (p Podcasts) GetBySlug(cfg *config.Config, slug string) (Podcasts, error) {
	session := cfg.Session.Copy()
	defer session.Close()
	var podCast Podcasts

	if err := cfg.Database.C(PodcastCollection).Find(bson.M{"slug": slug}).One(&podCast); err != nil {
		log.Fatal("Error retrieving that podcast", err)
		return podCast, err
	}
	return podCast, nil
}

//GetAll Get's all the podcasts
func (p Podcasts) GetAll(cfg *config.Config) ([]*Podcasts, error) {
	session := cfg.Session.Copy()
	defer session.Close()

	var podcasts []*Podcasts
	if err := cfg.Database.C(PodcastCollection).Find(bson.M{}).All(&podcasts); err != nil {
		log.Fatalf("%v\n", err)
		return podcasts, err
	}
	return podcasts, nil
}

//Remove Remove's a podcast
func (p Podcasts) Remove(cfg *config.Config, slug string) error {
	session := cfg.Session.Copy()
	defer session.Close()

	if err := cfg.Database.C(PodcastCollection).Remove(bson.M{"slug": slug}); err != nil {
		log.Fatalf("%v\n", err)
		return err
	}
	return nil
}

//Add Adds a podcast
func (p Podcasts) Add(cfg *config.Config) error {
	session := cfg.Session.Copy()
	defer session.Close()
	if err := cfg.Database.C(PodcastCollection).Insert(p); err != nil {
		log.Fatal("Error Adding a new podcast: ", err)
		return err
	}
	return nil
}

//Update Updates a podcast
func (p Podcasts) Update(cfg *config.Config, ID bson.ObjectId) error {
	session := cfg.Session.Copy()
	defer session.Close()
	p.DateUpdated = time.Now()
	if err := cfg.Database.C(PodcastCollection).Update(bson.M{"_id": ID}, p); err != nil {
		log.Fatal("Error Updating podcast")
		return err
	}
	return nil
}
