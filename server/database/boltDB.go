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

var projectsBucket = []byte("projects")
var usersBucket = []byte("users")

// TODO: Need to limit the amount of operations and logic
// in the database code

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

		_, err = tx.CreateBucketIfNotExists(usersBucket)
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
			return fmt.Errorf("Bucket %q not found.", projectsBucket)
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
			return fmt.Errorf("Bucket %q not found.", projectsBucket)
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
		bucket := tx.Bucket(projectsBucket)
		if bucket == nil {
			return fmt.Errorf("Bucket %q not found.", projectsBucket)
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
		bucket := tx.Bucket(projectsBucket)
		if bucket == nil {
			return fmt.Errorf("Bucket %q not found.", projectsBucket)
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
			return fmt.Errorf("Bucket %q not found.", projectsBucket)
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
			return fmt.Errorf("Bucket %q not found.", projectsBucket)
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
			return fmt.Errorf("Bucket %q not found.", projectsBucket)
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
			return fmt.Errorf("Bucket %q not found.", projectsBucket)
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
			return fmt.Errorf("Bucket %q not found.", projectsBucket)
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

func itob(v int) []byte {
	b := make([]byte, 8)
	binary.BigEndian.PutUint64(b, uint64(v))
	return b
}
