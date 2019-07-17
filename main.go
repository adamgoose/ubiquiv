package ubiquiv

import (
	"github.com/boltdb/bolt"
)

// Client wraps the KeyValue interface with the provided connection
type Client struct {
	Bucket []byte
}

// KeyValue is the interface with which you will interact!
type KeyValue interface {
	Get(key []byte) []byte
	Put(key, value []byte) error
	Delete(key []byte) error
	Close() error
}

// OpenFile creates a new Client against a given filepath
func OpenFile(path, bucket string) (KeyValue, error) {
	db, err := bolt.Open(path, 0600, nil)
	if err != nil {
		return nil, err
	}

	if err := db.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists([]byte(bucket))
		return err
	}); err != nil {
		return nil, err
	}

	return &LocalInterface{
		Client: &Client{
			Bucket: []byte(bucket),
		},
		DB: db,
	}, nil
}

// OpenServer creates a new Client against a given server address
func OpenServer(addr, bucket string) (KeyValue, error) {
	return &RemoteInterface{
		Client: &Client{
			Bucket: []byte(bucket),
		},
		ServerAddr: addr,
	}, nil
}
