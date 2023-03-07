package tasks

import (
	"task/types"

	"github.com/boltdb/bolt"
)

func Delete(db *bolt.DB, id types.ID) error {
	return db.Update(func(tx *bolt.Tx) error {
		bucket, err := tx.CreateBucketIfNotExists([]byte(bucketName))
		if err != nil {
			return err
		}
		return bucket.Delete(id.ToBytes())
	})
}
