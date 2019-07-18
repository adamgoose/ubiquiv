package ubivolt

import "github.com/boltdb/bolt"

// LocalInterface uses a local BoltDB file for implementation
type LocalInterface struct {
	*Client
	*bolt.DB
}

// Get gets the given key
func (b LocalInterface) Get(key []byte) (value []byte) {
	b.View(func(tx *bolt.Tx) error {
		value = tx.Bucket([]byte(b.Bucket)).Get(key)
		return nil
	})

	return
}

// Put puts the given key and value
func (b LocalInterface) Put(key, value []byte) error {
	return b.Update(func(tx *bolt.Tx) error {
		return tx.Bucket(b.Bucket).Put(key, value)
	})
}

// Delete deletes the given key
func (b LocalInterface) Delete(key []byte) error {
	return b.Update(func(tx *bolt.Tx) error {
		return tx.Bucket(b.Bucket).Delete(key)
	})
}

// List lists all keys
func (b LocalInterface) List() [][]byte {
	keys := [][]byte{}
	b.View(func(tx *bolt.Tx) error {
		tx.Bucket(b.Bucket).
			ForEach(func(k, v []byte) error {
				keys = append(keys, k)
				return nil
			})
		return nil
	})
	return keys
}
