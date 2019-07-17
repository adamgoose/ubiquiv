package ubiquiv

import (
	"encoding/gob"
	"io"
)

// RemoteRequest is used for serializing TCP requests
type RemoteRequest struct {
	Operation RemoteRequestOperation
	Bucket    []byte
	Key       []byte
	Value     []byte
}

// Encode the Request for transit
func (rr RemoteRequest) Encode(w io.Writer) error {
	return gob.NewEncoder(w).Encode(rr)
}

// Decode the Reqeust from transit
func (rr *RemoteRequest) Decode(r io.Reader) error {
	return gob.NewDecoder(r).Decode(rr)
}

// Execute the request against a KeyValue store
func (rr RemoteRequest) Execute(kv KeyValue) ([]byte, error) {
	switch rr.Operation {
	case GET:
		return kv.Get(rr.Key), nil
	case PUT:
		return nil, kv.Put(rr.Key, rr.Value)
	case DELETE:
		return nil, kv.Delete(rr.Key)
	}
	return nil, nil
}

// RemoteRequestOperation enum
type RemoteRequestOperation int

// RemoteRequestOperation values
const (
	GET    RemoteRequestOperation = iota
	PUT    RemoteRequestOperation = iota
	DELETE RemoteRequestOperation = iota
)

// RemoteResponse is used for serializing TCP responses
type RemoteResponse struct {
	Value []byte
	Error error
}

// Encode the Request for transit
func (rr RemoteResponse) Encode(w io.Writer) error {
	return gob.NewEncoder(w).Encode(rr)
}

// Decode the Reqeust from transit
func (rr *RemoteResponse) Decode(r io.Reader) error {
	return gob.NewDecoder(r).Decode(rr)
}
