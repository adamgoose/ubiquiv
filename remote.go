package ubiquiv

import (
	"log"
	"net"
)

// RemoteInterface uses an HTTP Client for implementation
type RemoteInterface struct {
	*Client
	ServerAddr string
}

// Get gets the given key
func (c RemoteInterface) Get(key []byte) []byte {
	req := RemoteRequest{
		Bucket:    c.Bucket,
		Operation: GET,
		Key:       key,
	}

	resp := RemoteResponse{}
	if err := c.doRemoteRequest(&resp, req); err != nil {
		return nil
	}

	return resp.Value
}

// Put puts the given key and value
func (c RemoteInterface) Put(key, value []byte) error {
	req := RemoteRequest{
		Bucket:    c.Bucket,
		Operation: PUT,
		Key:       key,
		Value:     value,
	}

	resp := RemoteResponse{}
	if err := c.doRemoteRequest(&resp, req); err != nil {
		return err
	}

	return resp.Error
}

// Delete deletes the given key
func (c RemoteInterface) Delete(key []byte) error {
	req := RemoteRequest{
		Bucket:    c.Bucket,
		Operation: DELETE,
		Key:       key,
	}

	resp := RemoteResponse{}
	if err := c.doRemoteRequest(&resp, req); err != nil {
		return err
	}

	return resp.Error
}

// Close does nothing for the remote interface, for now
func (c RemoteInterface) Close() error {
	return nil
}

func (c RemoteInterface) doRemoteRequest(resp *RemoteResponse, req RemoteRequest) error {
	conn, err := net.Dial("tcp", c.ServerAddr)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	if err := req.Encode(conn); err != nil {
		return err
	}
	if err := resp.Decode(conn); err != nil {
		return err
	}

	return nil
}
