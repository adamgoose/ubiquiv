package ubivolt

import (
	"context"

	proto "github.com/adamgoose/ubivolt/proto"
	"google.golang.org/grpc"
)

// GRPCInterface uses gRPC for implementation
type GRPCInterface struct {
	*Client
	Context context.Context
	Conn    *grpc.ClientConn
	GRPC    proto.KeyValueClient
}

// Get gets the given key
func (b GRPCInterface) Get(key []byte) []byte {
	r, _ := b.GRPC.Get(b.Context, &proto.Operation{
		Bucket: b.Bucket,
		Key:    key,
	})

	return r.Value
}

// Put puts the given key and value
func (b GRPCInterface) Put(key, value []byte) error {
	_, err := b.GRPC.Put(b.Context, &proto.Operation{
		Bucket: b.Bucket,
		Key:    key,
		Value:  value,
	})

	return err
}

// Delete deletes the given key
func (b GRPCInterface) Delete(key []byte) error {
	_, err := b.GRPC.Delete(b.Context, &proto.Operation{
		Bucket: b.Bucket,
		Key:    key,
	})

	return err
}

// List lists all keys
func (b GRPCInterface) List() [][]byte {
	r, _ := b.GRPC.List(b.Context, &proto.Operation{
		Bucket: b.Bucket,
	})

	return r.Keys
}

// Close closes the connection
func (b GRPCInterface) Close() error {
	return b.Conn.Close()
}
