package db

import (
	"sync"

	"github.com/hashicorp/go-multierror"
	"github.com/pkg/errors"
	"go.etcd.io/bbolt"
)

const (
	BucketSessions     = `sessions`
	BucketImages       = `images`
	BucketImageUploads = `images-up`
	BucketUploads      = `uploads`
)

var (
	dbLock      sync.RWMutex
	db          *bbolt.DB
	bucketNames = [][]byte{
		[]byte(BucketSessions),
		[]byte(BucketImages),
		[]byte(BucketImageUploads),
		[]byte(BucketUploads),
	}
)

func Init(path string) error {
	dbLock.Lock()
	defer dbLock.Unlock()
	var err error
	if db, err = bbolt.Open(path, 0660, &bbolt.Options{}); err != nil {
		return errors.Wrap(err, "failed to open database "+path)
	}
	err = db.Update(func(tx *bbolt.Tx) error {
		var result, e error
		for _, bucketName := range bucketNames {
			if _, e = tx.CreateBucketIfNotExists(bucketName); e != nil {
				result = multierror.Append(result, errors.Wrap(e, "failed to create bucket "+string(bucketName)))
			}
		}
		return result
	})
	return errors.Wrap(err, "failed to init database "+path)
}

func Get() *bbolt.DB {
	dbLock.RLock()
	defer dbLock.RUnlock()
	return db
}

func Close() error {
	dbLock.Lock()
	defer dbLock.Unlock()
	return db.Close()
}
