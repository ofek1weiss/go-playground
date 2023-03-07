package tasks

import (
	"encoding/json"
	"errors"
	"task/types"
	"time"

	"github.com/boltdb/bolt"
)

func GetByFilters(db *bolt.DB, filter func(*types.Task) bool) ([]*types.Task, error) {
	var tasks []*types.Task
	err := db.Update(func(tx *bolt.Tx) error {
		bucket, err := tx.CreateBucketIfNotExists([]byte(bucketName))
		if err != nil {
			return err
		}
		return bucket.ForEach(func(key, value []byte) error {
			task := new(types.Task)
			if err := json.Unmarshal(value, task); err != nil {
				return err
			}
			if filter(task) {
				tasks = append(tasks, task)
			}
			return nil
		})
	})
	return tasks, err
}

func GetByID(db *bolt.DB, id types.ID) (*types.Task, error) {
	task := new(types.Task)
	err := db.Update(func(tx *bolt.Tx) error {
		bucket, err := tx.CreateBucketIfNotExists([]byte(bucketName))
		if err != nil {
			return err
		}
		data := bucket.Get(id.ToBytes())
		if data == nil {
			return errors.New("task not found")
		}
		if err := json.Unmarshal(data, task); err != nil {
			return err
		}
		return nil
	})
	return task, err
}

func GetActive(db *bolt.DB) ([]*types.Task, error) {
	return GetByFilters(db, func(task *types.Task) bool {
		return !task.IsComplete()
	})
}

func GetCompleted(db *bolt.DB, timeBack time.Duration) ([]*types.Task, error) {
	minimumComletionTime := time.Now().Add(-timeBack)
	return GetByFilters(db, func(task *types.Task) bool {
		if !task.IsComplete() {
			return false
		}
		return time.Time(*task.CompletionTime).After(minimumComletionTime)
	})
}
