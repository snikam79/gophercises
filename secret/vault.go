package secret

import (
	"encoding/json"
	"errors"
	"io"
	"os"
	"sync"

	"secret/cipher"
)

// Vault is public struct which holds-
// encodingKey - used for encoding
// filePath - path on file system
// mutex - used for synchronizing read and write calls to Vault
// keyValues - simple key value map
type Vault struct {
	encodingKey string
	filepath    string
	mutex       sync.Mutex
	keyValues   map[string]string
}

// File is thin simple factory method for creating new Vault
func File(key, path string) *Vault {
	return &Vault{
		encodingKey: key,
		filepath:    path,
	}
}

func (v *Vault) readKeyValues(r io.Reader) error {
	dec := json.NewDecoder(r)
	return dec.Decode(&v.keyValues)
}

// as it's name suggests it loads vault
// if the map does not exists it creates one using calling make()
func (v *Vault) load() error {
	f, err := os.Open(v.filepath)
	if err != nil {
		v.keyValues = make(map[string]string)
		return nil
	}
	defer f.Close()

	reader, err := cipher.DecryptReader(v.encodingKey, f)
	if err != nil {
		return err
	}

	return v.readKeyValues(reader)
}

//save function writes encoding key to vault
//for this it opens vault file r/w or create mode
//so that if the file does not exists it creates one
func (v *Vault) save() error {
	f, err := os.OpenFile(v.filepath, os.O_RDWR|os.O_CREATE, 0755)
	if err != nil {
		return err
	}
	defer f.Close()

	w, err := cipher.EncryptWriter(v.encodingKey, f)
	if err != nil {
		return err
	}
	return v.writeKeyValues(w)
}

func (v *Vault) writeKeyValues(w io.Writer) error {
	enc := json.NewEncoder(w)
	return enc.Encode(v.keyValues)
}

//SetKey is used for stroing key:value pair in Vault
//It is thread safe as we have used mutex here.
func (v *Vault) SetKey(key, value string) error {
	v.mutex.Lock()
	defer v.mutex.Unlock()

	err := v.load()
	if err != nil {
		return err
	}
	v.keyValues[key] = value
	err = v.save()
	return err
}

//GetValue is public interface for retriving value
// for a specified key. This is thread safe implementation.
func (v *Vault) GetValue(key string) (string, error) {
	v.mutex.Lock()
	defer v.mutex.Unlock()

	err := v.load()
	if err != nil {
		return "", err
	}

	value, ok := v.keyValues[key]
	if !ok {
		return "", errors.New("secret: no value for that key")
	}
	return value, nil
}
