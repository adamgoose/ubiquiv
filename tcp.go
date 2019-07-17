package ubiquiv

import (
	"log"
	"net"

	"github.com/boltdb/bolt"
)

func (c LocalInterface) TCP(addr string) error {
	// Listen for incoming connections.
	l, err := net.Listen("tcp", addr)
	if err != nil {
		return err
	}

	// Close the listener when the application closes.
	defer l.Close()

	log.Println("Listening on " + addr)
	for {
		// Listen for an incoming connection.
		conn, err := l.Accept()
		if err != nil {
			return err
		}

		// Handle connections in a new goroutine.
		go func(conn net.Conn) {
			req := RemoteRequest{}
			if err := req.Decode(conn); err != nil {
				conn.Close()
				return
			}

			var value []byte
			var err error
			if err := c.Update(func(tx *bolt.Tx) error {
				b, err := tx.CreateBucketIfNotExists(req.Bucket)
				switch req.Operation {
				case GET:
					value = b.Get(req.Key)
					break
				case PUT:
					err = b.Put(req.Key, req.Value)
					break
				case DELETE:
					err = b.Delete(req.Key)
					break
				}
				return err
			}); err != nil {
				RemoteResponse{
					Error: err,
				}.Encode(conn)
				return
			}

			RemoteResponse{
				Value: value,
				Error: err,
			}.Encode(conn)

			conn.Close()
		}(conn)
	}
}
