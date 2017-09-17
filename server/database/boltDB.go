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
	"github.com/darwinfroese/hacksite/server/utilities"
)

var projectsBucket = []byte("projects")
var accountsBucket = []byte("accounts")
var sessionsBucket = []byte("sessions")

// TODO: Need to limit the amount of operations and logic in the database code
// TODO: Make sure the objects passed aren't being passed by copy so that we can just update
// fields and they're returned instead of having to explicitly return the object
// TODO: Iterations should probably be in their own bucket
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
func (b *boltDB) AddProject(project models.Project) (models.Project, error) {
	db, err := bolt.Open(b.dbLocation, 0644, nil)
	if err != nil {
		return models.Project{}, err
	}
	defer db.Close()

	project.Status = utilities.UpdateProjectStatus(project)

	err = db.Update(func(tx *bolt.Tx) error {
		bucket := tx.Bucket(projectsBucket)
		if bucket == nil {
			return fmt.Errorf("bucket %q not found", projectsBucket)
		}

		id, err := bucket.NextSequence()
		if err != nil {
			return err
		}

		project.ID = int(id)
		project.CurrentIteration.ProjectID = project.ID
		project.CurrentIteration.Number = 1
		project.Iterations = append(project.Iterations, project.CurrentIteration)
		for i, task := range project.CurrentIteration.Tasks {
			task.ProjectID = project.ID
			task.IterationNumber = project.CurrentIteration.Number
			project.CurrentIteration.Tasks[i] = task
		}

		key := itob(int(id))
		value, err := json.Marshal(project)

		if err != nil {
			return err
		}

		return bucket.Put(key, value)
	})

	return project, nil
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

// GetProjects will return all the projects in the database
func (b *boltDB) GetProjects(userID int) ([]models.Project, error) {
	db, err := bolt.Open(b.dbLocation, 0644, nil)
	if err != nil {
		return nil, err
	}
	defer db.Close()

	var account models.Account
	err = db.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket(accountsBucket)
		if bucket == nil {
			return fmt.Errorf("bucket %q not found", accountsBucket)
		}

		acc := bucket.Get(itob(userID))
		return json.Unmarshal(acc, &account)
	})

	var projects []models.Project
	err = db.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket(projectsBucket)
		if bucket == nil {
			return fmt.Errorf("bucket %q not found", projectsBucket)
		}

		for _, pid := range account.ProjectIds {
			var project models.Project

			p := bucket.Get(itob(pid))

			err := json.Unmarshal(p, &project)
			if err != nil {
				return err
			}

			projects = append(projects, project)
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	return projects, nil
}

// UpdateProject will store the new project in the database
func (b *boltDB) UpdateProject(p models.Project) error {
	db, err := bolt.Open(b.dbLocation, 0644, nil)
	if err != nil {
		return err
	}
	defer db.Close()

	p.Status = utilities.UpdateProjectStatus(p)

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

// UpdateTask will update a specfic task in a project
func (b *boltDB) UpdateTask(t models.Task) (models.Project, error) {
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

		val := bucket.Get(itob(t.ProjectID))
		err := json.Unmarshal(val, &project)
		if err != nil {
			return err
		}

		tasks := project.CurrentIteration.Tasks

		for i, task := range tasks {
			if task.ID == t.ID {
				tasks[i] = t
				break
			}
		}

		project.CurrentIteration.Tasks = tasks
		for i, iter := range project.Iterations {
			if iter.Number == project.CurrentIteration.Number {
				project.Iterations[i] = project.CurrentIteration
			}
		}

		project.Status = utilities.UpdateProjectStatus(project)

		v, err := json.Marshal(project)
		if err != nil {
			return err
		}

		err = bucket.Put(itob(t.ProjectID), v)
		if err != nil {
			return err
		}

		return nil
	})

	return project, nil
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

// RemoveTask will remove a task from a project
func (b *boltDB) RemoveTask(t models.Task) (models.Project, error) {
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

		val := bucket.Get(itob(t.ProjectID))

		err := json.Unmarshal(val, &project)
		if err != nil {
			return err
		}

		tasks := project.CurrentIteration.Tasks

		for i, task := range tasks {
			if task.ID == t.ID {
				tasks = append(tasks[:i], tasks[i+1:]...)
				break
			}
		}

		project.CurrentIteration.Tasks = tasks
		for i, iter := range project.Iterations {
			if iter.Number == project.CurrentIteration.Number {
				project.Iterations[i] = project.CurrentIteration
			}
		}

		p, err := json.Marshal(project)
		if err != nil {
			return err
		}

		err = bucket.Put(itob(t.ProjectID), p)
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
		project.Status = utilities.UpdateProjectStatus(project)

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
		project.Status = utilities.UpdateProjectStatus(project)

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

// Helper Functions
func itob(v int) []byte {
	b := make([]byte, 8)
	binary.BigEndian.PutUint64(b, uint64(v))
	return b
}
