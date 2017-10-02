package database

import (
	"encoding/base64"
	"encoding/binary"
	"encoding/json"
	"fmt"
	"os"
	"time"

	"github.com/boltdb/bolt"
	"github.com/darwinfroese/hacksite/server/models"
)

var projectsBucket = []byte("projects")
var accountsBucket = []byte("accounts")
var sessionsBucket = []byte("sessions")

// TODO: Make sure the objects passed aren't being passed by copy so that we can just update
// fields and they're returned instead of having to explicitly return the object
// TODO: Iterations and tasks should be in their own buckets
// TODO: Wrap db calls better - take function as argument, call after opening db and getting bucket

// CreateBoltDB creates a basic database struct
func createBoltDB() Database {
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

		key := itob(int(project.ID))
		value, err := json.Marshal(project)

		if err != nil {
			return err
		}

		return bucket.Put(key, value)
	})
}

// GetProject will lookup a project by id
func (b *boltDB) GetProject(id int) (models.Project, error) {
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

		v := bucket.Get(itob(id))
		err := json.Unmarshal(v, &project)
		if err != nil {
			return err
		}

		return nil
	})
	if err != nil {
		return models.Project{}, err
	}

	return project, nil
}

// GetNextProjectID returns the next sequence ID for the ProjectsBucket
func (b *boltDB) GetNextProjectID() (int, error) {
	return getNextID(b, projectsBucket)
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

		err = bucket.Put(itob(p.ID), v)
		if err != nil {
			return err
		}

		return nil
	})
}

// RemoveProject will remove a project from the database
func (b *boltDB) RemoveProject(id int) error {
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

		err := bucket.Delete(itob(id))
		if err != nil {
			return err
		}

		return nil
	})
}

// AddIteration will add an iteration to the project and update it in the database
func (b *boltDB) AddIteration(iteration models.Iteration) (models.Project, error) {
	db, err := bolt.Open(b.dbLocation, 0644, nil)
	if err != nil {
		return models.Project{}, err
	}
	defer db.Close()

	var project models.Project
	err = db.Update(func(tx *bolt.Tx) error {
		bucket := tx.Bucket(projectsBucket)
		if bucket == nil {
			return fmt.Errorf("bucket %q not found", projectsBucket)
		}

		val := bucket.Get(itob(iteration.ProjectID))

		err := json.Unmarshal(val, &project)
		if err != nil {
			return err
		}

		project.CurrentIteration = iteration
		project.Iterations = append(project.Iterations, iteration)
		// project.Status = projects.UpdateProjectStatus(project)

		p, err := json.Marshal(project)
		if err != nil {
			return err
		}

		err = bucket.Put(itob(iteration.ProjectID), p)
		if err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		return models.Project{}, err
	}

	return project, nil
}

// SwapCurrentIteration will set the iteration in the argument as the current iteration
func (b *boltDB) SwapCurrentIteration(iteration models.Iteration) (models.Project, error) {
	db, err := bolt.Open(b.dbLocation, 0644, nil)
	if err != nil {
		return models.Project{}, err
	}
	defer db.Close()

	var project models.Project
	err = db.Update(func(tx *bolt.Tx) error {
		bucket := tx.Bucket(projectsBucket)
		if bucket == nil {
			return fmt.Errorf("bucket %q not found", projectsBucket)
		}

		val := bucket.Get(itob(iteration.ProjectID))

		err := json.Unmarshal(val, &project)
		if err != nil {
			return err
		}

		project.CurrentIteration = iteration
		// project.Status = projects.UpdateProjectStatus(project)

		p, err := json.Marshal(project)
		if err != nil {
			return err
		}

		err = bucket.Put(itob(iteration.ProjectID), p)
		if err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		return models.Project{}, err
	}

	return project, nil
}

// CreateAccount creates a user in the database
func (b *boltDB) CreateAccount(account models.Account) (int, error) {
	db, err := bolt.Open(b.dbLocation, 0644, nil)
	if err != nil {
		return -1, err
	}
	defer db.Close()

	err = db.Update(func(tx *bolt.Tx) error {
		bucket := tx.Bucket(accountsBucket)
		if bucket == nil {
			return fmt.Errorf("bucket %q not found", accountsBucket)
		}

		id, err := bucket.NextSequence()
		if err != nil {
			return err
		}

		account.ID = int(id)
		a, err := json.Marshal(account)
		if err != nil {
			return err
		}

		key := itob(int(id))
		return bucket.Put(key, a)
	})

	if err != nil {
		return -1, err
	}

	return account.ID, nil
}

// GetAccount finds an account in the database if there is a matching username
func (b *boltDB) GetAccount(username string) (models.Account, error) {
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

		return fmt.Errorf("no matching account found")
	})

	return account, err
}

// GetAccountByID looksup an account by using the userID key
func (b *boltDB) GetAccountByID(userID int) (models.Account, error) {
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

		a := bucket.Get(itob(userID))

		return json.Unmarshal(a, &account)
	})

	return account, err
}

// GetNextAccountID returns the next sequence for the AccountsBucket
func (b *boltDB) GetNextAccountID() (int, error) {
	return getNextID(b, accountsBucket)
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

		key := itob(account.ID)
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

// GetNextSessionID returns the next sequence for the sessionBucket
func (b *boltDB) GetNextSessionID() (int, error) {
	return getNextID(b, sessionsBucket)
}

// CleanSessions removes all sessions that are expired from the database and returns the
// number of sessions removed
func (b *boltDB) CleanSessions() (int, error) {
	db, err := bolt.Open(b.dbLocation, 0644, nil)
	if err != nil {
		return -1, err
	}
	defer db.Close()

	count := 0
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

			if time.Now().After(sesh.Expiration) {
				err := c.Delete()
				if err != nil {
					return err
				}

				count++
			}
		}

		return nil
	})

	return count, err
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

// Helper Functions
func itob(v int) []byte {
	b := make([]byte, 8)
	binary.BigEndian.PutUint64(b, uint64(v))
	return b
}

func getNextID(b *boltDB, bucketName []byte) (int, error) {
	db, err := bolt.Open(b.dbLocation, 0644, nil)
	if err != nil {
		return -1, err
	}
	defer db.Close()

	id := -1
	err = db.Update(func(tx *bolt.Tx) error {
		bucket := tx.Bucket(bucketName)
		if bucket == nil {
			return fmt.Errorf("bucket %q not found", bucketName)
		}

		i, err := bucket.NextSequence()
		if err != nil {
			return err
		}

		id = int(i)
		return nil
	})

	return id, err
}
