//go:generate protoc --go_out=plugins=grpc:./proto --js_out=import_style=commonjs:./proto --grpc-web_out=import_style=commonjs,mode=grpcwebtext:./proto --csharp_out=./proto ./keyvalue.proto

package ubiquiv

import (
	"context"
	"errors"

	proto "github.com/adamgoose/ubiquiv/proto"
	"github.com/boltdb/bolt"
)

// GRPCServer implements the KeyValue gRPC Service Protocol
type GRPCServer struct {
	DB *bolt.DB
}

// Get a value from the store
func (s GRPCServer) Get(ctx context.Context, op *proto.Operation) (*proto.Result, error) {
	if len(op.Bucket) == 0 || len(op.Key) == 0 {
		return nil, errors.New("bucket and key are required")
	}

	var value []byte
	s.DB.View(func(tx *bolt.Tx) error {
		value = tx.Bucket(op.Bucket).Get(op.Key)
		return nil
	})

	return &proto.Result{Value: value}, nil
}

// Put a value to the store
func (s GRPCServer) Put(ctx context.Context, op *proto.Operation) (*proto.Result, error) {
	if len(op.Bucket) == 0 || len(op.Key) == 0 || len(op.Value) == 0 {
		return nil, errors.New("bucket, key, and value are required")
	}

	if err := s.DB.Update(func(tx *bolt.Tx) error {
		return tx.Bucket(op.Bucket).Put(op.Key, op.Value)
	}); err != nil {
		return nil, err
	}

	return &proto.Result{}, nil
}

// Delete a value from the store
func (s GRPCServer) Delete(ctx context.Context, op *proto.Operation) (*proto.Result, error) {
	if len(op.Bucket) == 0 || len(op.Key) == 0 {
		return nil, errors.New("bucket and key are required")
	}

	if err := s.DB.Update(func(tx *bolt.Tx) error {
		return tx.Bucket(op.Bucket).Delete(op.Key)
	}); err != nil {
		return nil, err
	}

	return &proto.Result{}, nil
}
