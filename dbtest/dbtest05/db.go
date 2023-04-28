package main

import (
	"context"
	"log"

	"github.com/ava-labs/avalanchego/database"
	"github.com/ava-labs/hypersdk/pebble"

	"github.com/ava-labs/hypersdk/chain"
)

var _ chain.Database = (*MyDatabase)(nil)

type MyDatabase struct {
	db database.Database
}

func NewMyDatabase() (*MyDatabase, error) {
	cfg := pebble.NewDefaultConfig()
	d, e := pebble.New("demo", cfg)
	return &MyDatabase{
		db: d,
	}, e
}
func (s *MyDatabase) GetValue(ctx context.Context, key []byte) ([]byte, error) {
	b, err := s.db.Get(key)
	if err != nil {
		log.Printf("db.Get err: %v", err)
		return nil, err
	}
	return b, nil

}

func (d *MyDatabase) Insert(ctx context.Context, key []byte, value []byte) error {
	log.Printf("Insert: key=%x", key)
	log.Printf("Insert: value=%v", value)
	// b, err := d.Has(key)
	// if b == false {
	// 	d.Put(key, uint64ToBytes(0))
	// 	if err != nil {
	// 		log.Printf("err: %v", err)
	// 		return err
	// 	}
	// 	return nil
	// }
	// if err != nil {
	// 	log.Printf("db.Has err: %v", err)
	// 	return err
	// }
	// log.Printf("Insert: key=%v", key)
	// // v, err := d.GetValue(ctx, key)
	// // if err != nil {
	// // 	log.Printf("d.GetValue err: %v", err)
	// // }
	// log.Printf("Insert:Put value=%v", value)

	// // err = d.Put(key, uint64ToBytes(bytesToUint64(v)+bytesToUint64(value)))
	// err = d.Put(key, value)
	// if err != nil {
	// 	log.Printf("d.Put err: %v", err)
	// 	return err
	// }

	log.Printf("db.put(key,val)")
	log.Printf("key=%v", key)
	log.Printf("value=%v", value)
	err := d.db.Put(key, value)
	if err != nil {
		log.Printf("d.db.Put err: %v", err)
	}
	return err

	// return nil
}
func (d *MyDatabase) Remove(ctx context.Context, key []byte) error {
	return nil
}
func (d *MyDatabase) Close() error {
	return d.db.Close()
}

func (d *MyDatabase) Has(key []byte) (bool, error) {
	return d.db.Has(key)
}

func (d *MyDatabase) Put(key []byte, val []byte) error {
	return d.db.Put(key, val)

}

func (d *MyDatabase) Get(key []byte) ([]byte, error) {
	return d.db.Get(key)
}
