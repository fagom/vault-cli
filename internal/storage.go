package internal

import (
	"fmt"
	"log"

	"go.etcd.io/bbolt"
)

func InitDb() *bbolt.DB {
	db, err := bbolt.Open("passwords.db", 0600, nil)
	if err != nil {
		log.Fatal(err)
	}

	return db
}

func CreatePassword(db *bbolt.DB, key string, pass string) error {
	return db.Update(func(tx *bbolt.Tx) error {
		b, err := tx.CreateBucketIfNotExists([]byte("passwords"))
		if err != nil {
			return err
		}
		return b.Put([]byte(key), []byte(pass))
	})
}

func GetPassword(db *bbolt.DB, key string) (string, error) {
	var pass []byte
	err := db.View(func(tx *bbolt.Tx) error {
		b := tx.Bucket([]byte("passwords"))
		if b == nil {
			return fmt.Errorf("Password not fount")
		}
		pass = b.Get([]byte(key))

		return nil
	})

	if err != nil {
		return "", err
	}
	return string(pass), nil
}
