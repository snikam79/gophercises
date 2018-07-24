package db

import (
	"path/filepath"
	"testing"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/stretchr/testify/assert"
)

const dbPath = "c:\\Users\\gs-1025\\tasks.db"

// func TestAdd(t *testing.T) {
// 	out, err := exec.Command("go", "run", "../main.go", "add", "Watch Movie").Output()

// 	if err != nil {
// 		fmt.Printf("something went wrong : %s", err)
// 	}
// 	//fmt.Printf("The output  is %s\n", out)
// 	str := fmt.Sprintf("%s", out)
// 	//fmt.Println(str)

// 	// command output must contain Added "asdf " task to task list.
// 	flag := strings.Contains(str, "Watch Movie")
// 	assert.Equal(t, flag, true)
// 	//fmt.Println(flag)

// 	//now give
// }

// func TestList(t *testing.T) {
// 	out, err := exec.Command("go", "run", "../main.go", "list").Output()

// 	if err != nil {
// 		fmt.Printf("something went wrong : %s", err)
// 	}

// 	str := fmt.Sprintf("%s", out)
// 	//fmt.Println(str)

// 	// command output must contain Added "asdf " task to task list.
// 	flag := strings.Contains(str, "Watch Movie")
// 	assert.Equal(t, flag, true)
// }

func TestInit(t *testing.T) {
	home, _ := homedir.Dir()
	dbPath := filepath.Join(home, "tasks.db")
	err := Init(dbPath)
	assert.Equal(t, err, nil)
}

func TestCreateTask(t *testing.T) {
	testTask := "sample task"
	_, err := CreateTask(testTask)
	assert.Equal(t, err, nil)
}

func TestDeleteTask(t *testing.T) {
	err := DeleteTask(3)
	assert.Equal(t, err, nil)
}

func TestGetAllTasks(t *testing.T) {
	tasks, err := GetAllTasks()
	assert.Equal(t, err, nil)
	notNil := (tasks != nil)
	assert.Equal(t, notNil, true)

}
