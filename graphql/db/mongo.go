package db

import (
	"github.com/pkg/errors"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// UserProfile represents a user profile in NoSQL database.
type UserProfile struct {
	Permissions []string `json:"permissions"`
}

// Mongo represents an implementation of NoSQL database.
type Mongo struct {
	conn   *mgo.Session
	dbName string
}

// NewMongo returns an instance of NoSQL database.
func NewMongo(dbURL, dbName string) (*Mongo, error) {
	conn, err := mgo.Dial("localhost")
	if err != nil {
		return nil, errors.Wrapf(err, "failed to connect to database at %s", dbURL)
	}

	return &Mongo{
		conn:   conn,
		dbName: dbName,
	}, nil
}

// Close closes database connection.
func (m Mongo) Close() {
	m.conn.Close()
}

func (m Mongo) profiles() *mgo.Collection {
	return m.conn.DB(m.dbName).C("profiles")
}

// GetPermissions returns a list of permissions for one (random) user profile.
func (m Mongo) GetPermissions() ([]string, error) {
	var profile UserProfile
	if err := m.profiles().Find(bson.M{}).One(&profile); err != nil {
		return profile.Permissions, errors.Wrap(err, "failed to get user permissions")
	}

	return profile.Permissions, nil
}

// GetUserPermissions returns a list of permissions of a profile represented by given username.
func (m Mongo) GetUserPermissions(username string) ([]string, error) {
	var profile UserProfile
	if err := m.profiles().Find(bson.M{"_id": username}).One(&profile); err != nil {
		return profile.Permissions, errors.Wrap(err, "failed to get user permissions")
	}

	return profile.Permissions, nil
}

// AddUserPermission adds a permission to the profile represented by given username.
func (m Mongo) AddUserPermission(username, permission string) error {
	query := bson.M{"_id": username}
	update := bson.M{
		"$addToSet": bson.M{
			"permissions": permission,
		},
	}

	if err := m.profiles().Update(query, update); err != nil {
		return errors.Wrap(err, "failed to update permissions list")
	}

	return nil
}

/*
mongo, err := db.NewMongo("localhost", "graphql-example")
if err != nil {
    log.Fatalf("failed to connect to Mongo: %v", err)
}
defer mongo.Close()
*/
