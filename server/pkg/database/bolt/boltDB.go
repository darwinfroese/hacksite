package bolt

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"os"

	"github.com/boltdb/bolt"
	"github.com/darwinfroese/hacksite/server/models"
	"github.com/darwinfroese/hacksite/server/pkg/database"
)

type boltDB struct {
	dbLocation string
}

var projectsBucket = []byte("projects")
var accountsBucket = []byte("accounts")
var sessionsBucket = []byte("sessions")

// TODO: Make sure the objects passed aren't being passed by copy so that we can just update
// fields and they're returned instead of having to explicitly return the object
// TODO: Iterations and tasks should be in their own buckets
// TODO: Wrap db calls better - take function as argument, call after opening db and getting bucket

// New creates a basic database struct
func New() database.Database {
	db := boltDB{
		dbLocation: "database.db",
	}

	createBuckets(db)
	return &db
}

func createBuckets(b boltDB) {
	db, err := bolt.Open(b.dbLocation, 0644, nil)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %s\n", err.Error())
		return
	}
	defer db.Close()

	err = db.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists(projectsBucket)
		if err != nil {
			return err
		}

		_, err = tx.CreateBucketIfNotExists(accountsBucket)
		if err != nil {
			return err
		}

		_, err = tx.CreateBucketIfNotExists(sessionsBucket)
		if err != nil {
			return err
		}

		return nil
	})

	// TODO: This should probably crash the program? or attempt to recover?
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %s\n", err.Error())
	}
}

// AddProject will add a project to the database
func (b *boltDB) AddProject(project models.Project) error {
	db, err := bolt.Open(b.dbLocation, 0644, nil)
	if err != nil {
		return err
	}
	defer db.Close()

	return db.Update(func(tx *bolt.Tx) error {
		bucket := tx.Bucket(projectsBucket)
		if bucket == nil {
			return fmt.Errorf("bucket %q not found", projectsBucket)
		}

		key := []byte(project.ID)
		value, err := json.Marshal(project)

		if err != nil {
			return err
		}

		return bucket.Put(key, value)
	})
}

// GetProject will lookup a project by id
func (b *boltDB) GetProject(id string) (models.Project, error) {
	db, err := bolt.Open(b.dbLocation, 0644, nil)
	if err != nil {
		return models.Project{}, err
	}
	defer db.Close()

	var project models.Project
	err = db.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket(projectsBucket)
		if bucket == nil {
			return fmt.Errorf("bucket %q not found", projectsBucket)
		}

		v := bucket.Get([]byte(id))
		err := json.Unmarshal(v, &project)
		if err != nil {
			return err
		}

		return nil
	})

	return project, err
}

// UpdateProject will store the new project in the database
func (b *boltDB) UpdateProject(p models.Project) error {
	db, err := bolt.Open(b.dbLocation, 0644, nil)
	if err != nil {
		return err
	}
	defer db.Close()

	return db.Update(func(tx *bolt.Tx) error {
		bucket := tx.Bucket(projectsBucket)
		if bucket == nil {
			return fmt.Errorf("bucket %q not found", projectsBucket)
		}

		v, err := json.Marshal(p)
		if err != nil {
			return err
		}

		err = bucket.Put([]byte(p.ID), v)
		if err != nil {
			return err
		}

		return nil
	})
}

// RemoveProject will remove a project from the database
func (b *boltDB) RemoveProject(id string) error {
	db, err := bolt.Open(b.dbLocation, 0644, nil)
	if err != nil {
		return err
	}
	defer db.Close()

	return db.Update(func(tx *bolt.Tx) error {
		bucket := tx.Bucket(projectsBucket)
		if bucket == nil {
			return fmt.Errorf("bucket %q not found", projectsBucket)
		}

		err := bucket.Delete([]byte(id))
		if err != nil {
			return err
		}

		return nil
	})
}

// CreateAccount creates a user in the database
func (b *boltDB) CreateAccount(account models.Account) error {
	db, err := bolt.Open(b.dbLocation, 0644, nil)
	if err != nil {
		return err
	}
	defer db.Close()

	return db.Update(func(tx *bolt.Tx) error {
		bucket := tx.Bucket(accountsBucket)
		if bucket == nil {
			return fmt.Errorf("bucket %q not found", accountsBucket)
		}

		a, err := json.Marshal(account)
		if err != nil {
			return err
		}

		return bucket.Put([]byte(account.Username), a)
	})
}

// GetAccountByUsername finds an account in the database if there is a matching username
func (b *boltDB) GetAccountByUsername(username string) (models.Account, error) {
	db, err := bolt.Open(b.dbLocation, 0644, nil)
	if err != nil {
		return models.Account{}, err
	}
	defer db.Close()

	var account models.Account
	err = db.Update(func(tx *bolt.Tx) error {
		bucket := tx.Bucket(accountsBucket)
		if bucket == nil {
			return fmt.Errorf("bucket %q not found", accountsBucket)
		}

		c := bucket.Cursor()
		for k, v := c.First(); k != nil; k, v = c.Next() {
			var acc models.Account
			err = json.Unmarshal(v, &acc)

			if err != nil {
				return err
			}

			if acc.Username == username {
				account = acc
				return nil
			}

		}

		return nil
	})

	return account, err
}

// GetAccountByUsername finds an account in the database if there is a matching username
func (b *boltDB) GetAccountByEmail(email string) (models.Account, error) {
	db, err := bolt.Open(b.dbLocation, 0644, nil)
	if err != nil {
		return models.Account{}, err
	}
	defer db.Close()

	var account models.Account
	err = db.Update(func(tx *bolt.Tx) error {
		bucket := tx.Bucket(accountsBucket)
		if bucket == nil {
			return fmt.Errorf("bucket %q not found", accountsBucket)
		}

		c := bucket.Cursor()
		for k, v := c.First(); k != nil; k, v = c.Next() {
			var acc models.Account
			err = json.Unmarshal(v, &acc)

			if err != nil {
				return err
			}

			if acc.Email == email {
				account = acc
				return nil
			}

		}

		return nil
	})

	return account, err
}

// UpdateAccount inserts a new account into the accounts location in the bucket
func (b *boltDB) UpdateAccount(account models.Account) error {
	db, err := bolt.Open(b.dbLocation, 0644, nil)
	if err != nil {
		return err
	}
	defer db.Close()

	return db.Update(func(tx *bolt.Tx) error {
		bucket := tx.Bucket(accountsBucket)
		if bucket == nil {
			return fmt.Errorf("bucket %q not found", accountsBucket)
		}

		key := []byte(account.Username)
		value, err := json.Marshal(account)
		if err != nil {
			return err
		}

		return bucket.Put(key, value)
	})
}

// StoreSession inserts a session into the sessions bucket
func (b *boltDB) StoreSession(session models.Session) error {
	db, err := bolt.Open(b.dbLocation, 0644, nil)
	if err != nil {
		return err
	}
	defer db.Close()

	return db.Update(func(tx *bolt.Tx) error {
		bucket := tx.Bucket(sessionsBucket)
		if bucket == nil {
			return fmt.Errorf("bucket %q not found", sessionsBucket)
		}

		key, err := base64.StdEncoding.DecodeString(session.Token)
		if err != nil {
			return err
		}
		val, err := json.Marshal(session)
		if err != nil {
			return err
		}

		return bucket.Put(key, val)
	})
}

// GetSession looks up a session in the database
func (b *boltDB) GetSession(sessionToken string) (models.Session, error) {
	db, err := bolt.Open(b.dbLocation, 0644, nil)
	if err != nil {
		return models.Session{}, err
	}
	defer db.Close()

	var session models.Session
	err = db.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket(sessionsBucket)
		if bucket == nil {
			return fmt.Errorf("bucket %q not found", sessionsBucket)
		}

		key, err := base64.StdEncoding.DecodeString(sessionToken)
		if err != nil {
			return err
		}

		val := bucket.Get(key)
		return json.Unmarshal(val, &session)
	})

	return session, err
}

// GetAllSessions returns all the sessions currently in the database
func (b *boltDB) GetAllSessions() ([]models.Session, error) {
	db, err := bolt.Open(b.dbLocation, 0644, nil)
	if err != nil {
		return nil, err
	}
	defer db.Close()

	var sessions []models.Session
	err = db.Update(func(tx *bolt.Tx) error {
		bucket := tx.Bucket(sessionsBucket)
		if bucket == nil {
			return fmt.Errorf("bucket %q not found", sessionsBucket)
		}

		c := bucket.Cursor()
		for k, v := c.First(); k != nil; k, v = c.Next() {
			var sesh models.Session
			err = json.Unmarshal(v, &sesh)

			if err != nil {
				return err
			}

			sessions = append(sessions, sesh)
		}

		return nil
	})

	return sessions, err
}

// RemoveSession removes the token from the database, typically for loguout
func (b *boltDB) RemoveSession(sessionToken string) error {
	db, err := bolt.Open(b.dbLocation, 0644, nil)
	if err != nil {
		return err
	}
	defer db.Close()

	return db.Update(func(tx *bolt.Tx) error {
		bucket := tx.Bucket(sessionsBucket)
		if bucket == nil {
			return fmt.Errorf("bucket %q not found", sessionsBucket)
		}

		key, err := base64.StdEncoding.DecodeString(sessionToken)
		if err != nil {
			return err
		}

		return bucket.Delete(key)
	})
}
