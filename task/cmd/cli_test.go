package cmd

import (
	"fmt"
	"os/exec"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestList(t *testing.T) {
	out, err := exec.Command("go", "run", "../main.go", "list").Output()
	if err != nil {
		fmt.Println("Something went wrong while invoking list command")
		return
	}
	str := fmt.Sprintf("%s", out)
	flag := strings.Contains(str, "sample task")
	assert.Equal(t, flag, true)

}

func Testadd(t *testing.T) {
	out, err := exec.Command("go", "run", "../main.go", "add", "new task").Output()
	if err != nil {
		fmt.Println("Something went wrong while invoking list command")
		return
	}
	str := fmt.Sprintf("%s", out)
	flag := strings.Contains(str, "new task")
	assert.Equal(t, flag, true)

}

func Testdo(t *testing.T) {
	out, err := exec.Command("go", "run", "../main.go", "do", "2").Output()
	if err != nil {
		fmt.Println("Something went wrong while invoking list command")
		return
	}
	str := fmt.Sprintf("%s", out)
	flag := strings.Contains(str, "Marked")
	assert.Equal(t, flag, true)

}

func TestPlainTask(t *testing.T) {
	out, err := exec.Command("go", "run", "../main.go").Output()
	if err != nil {
		fmt.Println("Something went wrong while invoking list command")
		return
	}
	str := fmt.Sprintf("%s", out)
	flag := strings.Contains(str, "task is a todo manager")
	assert.Equal(t, flag, false)

}
