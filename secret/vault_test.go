package secret

import (
	"path/filepath"
	"testing"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/stretchr/testify/assert"
)

//TestFile basic test case
func TestFile(t *testing.T) {
	home, _ := homedir.Dir()
	//	return filepath.Join(home, "test.secrets")
	filePath := filepath.Join(home, "test.secrets")

	v := File("test", filePath)
	assert.Equal(t, "test", v.encodingKey)

	key1 := "key1"
	value1 := "value1"
	v.SetKey(key1, value1)

	answer1, _ := v.GetValue(key1)
	assert.Equal(t, answer1, value1)

	// now try to get value for non existing key
	s, err := v.GetValue("")
	assert.Equal(t, s, "")
	assert.Equal(t, err.Error, "secret: no value for that key")

}
