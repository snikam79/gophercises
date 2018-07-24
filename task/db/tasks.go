package db

import (
	"encoding/binary"
	"time"

	"github.com/boltdb/bolt"
)

// we are defining task bucket named "tasks"
// very much similar to table in sql
var taskBucket = []byte("tasks")

//handle to bolt db
// it used for all DML operations
var db *bolt.DB

// Task structure is representation of task.
// simple key-value format
type Task struct {
	Key   int
	Value string
}

// Init function open connection with bolt db
// Also it creates a bucket if it does not exists already.
func Init(dbPath string) error {
	var err error
	db, err = bolt.Open(dbPath, 0600, &bolt.Options{Timeout: 1 * time.Second})

	if err != nil {
		return err
	}

	// use of function as a variable
	// we can also use anonymous function as well but it is kind of hard to read

	dbUpdateFn := func(tx *bolt.Tx) error {
		_, err = tx.CreateBucketIfNotExists(taskBucket)
		return err
	}

	return db.Update(dbUpdateFn)
}

// CreateTask creates new task in db
// For numbering the task it uses auto increament functionality of Bucket interface
func CreateTask(task string) (int, error) {
	var id int

	updateClosure := func(tx *bolt.Tx) error {
		b := tx.Bucket(taskBucket)
		id64, _ := b.NextSequence()
		id = int(id64)
		key := intToByte(id)
		return b.Put(key, []byte(task))
	}

	err := db.Update(updateClosure)
	if err != nil {
		return -1, err
	}

	return id, nil
}

//DeleteTask deletes the task from bucket
func DeleteTask(key int) error {
	deleteFn := func(tx *bolt.Tx) error {
		b := tx.Bucket(taskBucket)
		return b.Delete(intToByte(key))
	}
	return db.Update(deleteFn)

}

func intToByte(i int) []byte {
	b := make([]byte, 8)
	binary.BigEndian.PutUint64(b, uint64(i))
	return b
}

func byteToInt(b []byte) int {
	return int(binary.BigEndian.Uint64(b))
}

//GetAllTasks returns all tasks from bucket
// It returns read only view of tasks to the caller
func GetAllTasks() ([]Task, error) {
	var tasks []Task

	allTasksFn := func(tx *bolt.Tx) error {
		b := tx.Bucket(taskBucket)
		c := b.Cursor()
		for k, v := c.First(); k != nil; k, v = c.Next() {
			tasks = append(tasks, Task{
				Key:   byteToInt(k),
				Value: string(v),
			})
		}
		return nil
	}

	err := db.View(allTasksFn)
	if err != nil {
		return nil, err
	}
	return tasks, nil
}
