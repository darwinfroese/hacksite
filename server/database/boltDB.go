package database

import (
	"encoding/binary"
	"encoding/json"
	"fmt"
	"os"

	"github.com/boltdb/bolt"
	"github.com/darwinfroese/hacksite/server/models"
	"github.com/darwinfroese/hacksite/server/utilities"
)

var projectBucket = []byte("projects")

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
		_, err := tx.CreateBucketIfNotExists(projectBucket)
		if err != nil {
			return err
		}

		return nil
	})

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
	project.Iteration = models.Iteration{Number: 1}

	err = db.Update(func(tx *bolt.Tx) error {
		bucket := tx.Bucket(projectBucket)
		if bucket == nil {
			return fmt.Errorf("Bucket %q not fonud.", projectBucket)
		}

		id, err := bucket.NextSequence()
		if err != nil {
			return err
		}

		project.ID = int(id)
		for i, task := range project.Tasks {
			task.ProjectID = int(id)
			project.Tasks[i] = task
			task.IterationNumber = project.Iteration.Number
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
		bucket := tx.Bucket(projectBucket)
		if bucket == nil {
			return fmt.Errorf("Bucket %q not found.", projectBucket)
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
func (b *boltDB) GetProjects() ([]models.Project, error) {
	db, err := bolt.Open(b.dbLocation, 0644, nil)
	if err != nil {
		return nil, err
	}
	defer db.Close()

	var projects []models.Project
	err = db.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket(projectBucket)
		if bucket == nil {
			return fmt.Errorf("Bucket %q not found.", projectBucket)
		}

		c := bucket.Cursor()
		for k, v := c.First(); k != nil; k, v = c.Next() {
			var project models.Project
			err = json.Unmarshal(v, &project)

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
		bucket := tx.Bucket(projectBucket)
		if bucket == nil {
			return fmt.Errorf("Bucket %q not found.", projectBucket)
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
		bucket := tx.Bucket(projectBucket)
		if bucket == nil {
			return fmt.Errorf("Bucket %q not found.", projectBucket)
		}

		val := bucket.Get(itob(t.ProjectID))
		err := json.Unmarshal(val, &project)
		if err != nil {
			return err
		}

		for i, task := range project.Tasks {
			if task.ID == t.ID {
				project.Tasks[i] = t
				break
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
		bucket := tx.Bucket(projectBucket)
		if bucket == nil {
			return fmt.Errorf("Bucket %q not found.", projectBucket)
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
		bucket := tx.Bucket(projectBucket)
		if bucket == nil {
			return fmt.Errorf("Bucket %q not found.", projectBucket)
		}

		val := bucket.Get(itob(t.ProjectID))

		err := json.Unmarshal(val, &project)
		if err != nil {
			return err
		}

		for i, task := range project.Tasks {
			if task.ID == t.ID {
				project.Tasks = append(project.Tasks[:i], project.Tasks[i+1:]...)
				break
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

func itob(v int) []byte {
	b := make([]byte, 8)
	binary.BigEndian.PutUint64(b, uint64(v))
	return b
}
