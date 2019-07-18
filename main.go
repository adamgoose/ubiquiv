package ubiquiv

import (
	"context"

	proto "github.com/adamgoose/ubiquiv/proto"
	"github.com/boltdb/bolt"
	"google.golang.org/grpc"
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
	conn, err := grpc.Dial(addr, grpc.WithInsecure())
	if err != nil {
		return nil, err
	}

	return &GRPCInterface{
		Client: &Client{
			Bucket: []byte(bucket),
		},
		Context: context.Background(),
		Conn:    conn,
		GRPC:    proto.NewKeyValueClient(conn),
	}, nil
}
