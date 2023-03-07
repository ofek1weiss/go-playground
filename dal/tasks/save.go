package tasks

import (
	"encoding/json"
	"task/types"

	"github.com/boltdb/bolt"
)

func Save(db *bolt.DB, task *types.Task) error {
	return db.Update(func(tx *bolt.Tx) error {
		bucket, err := tx.CreateBucketIfNotExists([]byte(bucketName))
		if err != nil {
			return err
		}
		if task.Id == 0 {
			id, _ := bucket.NextSequence()
			task.Id = types.ID(id)
		}
		data, err := json.Marshal(task)
		if err != nil {
			return err
		}
		return bucket.Put(task.Id.ToBytes(), data)
	})
}
