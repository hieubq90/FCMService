package database

import (
	"fmt"
	"time"

	"encoding/json"

	"github.com/boltdb/bolt"
)

var db *bolt.DB

func init() {
	var err error
	db, err = bolt.Open("fcm.db", 0600, &bolt.Options{Timeout: 1 * time.Second})
	if err != nil {
		fmt.Println("[FCMService] Error on opening fcm.db", err)
	}

	// init bucket
	db.Update(func(tx *bolt.Tx) error {
		b, err := tx.CreateBucket([]byte("users"))
		if b != nil && err == nil {
			fmt.Println("Init users bucket successful!")
		} else {
			fmt.Println("Can not init users bucket!")
		}
		return err
	})
}

func AddNewToken(phone, token string) error {
	return db.Update(func(tx *bolt.Tx) error {
		// Retrieve the users bucket.
		// This should be created when the DB is first opened.
		b := tx.Bucket([]byte("users"))

		// get current data
		v := b.Get([]byte(phone))
		user := &User{}
		err := json.Unmarshal(v, user)
		if err == nil {
			// do update
			user.Tokens = append(user.Tokens, token)
		} else {
			// create new user
			user = &User{
				Phone: phone,
				Tokens: []string {token},
			}

		}
		// Marshal user data into bytes.
		buf, err := json.Marshal(user)
		if err != nil {
			return err
		}

		// Persist bytes to users bucket.
		return b.Put([]byte(phone), buf)
	})
}

func GetUserByPhone(phone string) (user *User, err error) {
	user = &User{
		Phone: phone,
		Tokens: []string {},
	}
	err = db.View(func(tx *bolt.Tx) (err error) {
		// Retrieve the users bucket.
		// This should be created when the DB is first opened.
		b := tx.Bucket([]byte("users"))

		// get current data
		v := b.Get([]byte(phone))
		err = json.Unmarshal(v, user)
		if err != nil {
			// create new user
			user = &User{
				Phone: phone,
				Tokens: []string {},
			}

		}
		return
	})
	return
}