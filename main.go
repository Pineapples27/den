package den

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func Save(key string, contents interface{}) error {
	jsonBytes, err := json.Marshal(contents)
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(keyToFileName(key), jsonBytes, os.ModePerm)
	if err != nil {
		return err
	}
	return nil
}

func Fetch(key string, returnPointer interface{}) error {
	bytes, err := ioutil.ReadFile(keyToFileName(key))
	if err != nil {
		return err
	}
	err = json.Unmarshal(bytes, &returnPointer)
	if err != nil {
		return err
	}
	return nil
}

func FetchRaw(key string) (returnedContents interface{}, err error) {
	bytes, err := ioutil.ReadFile(keyToFileName(key))
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(bytes, &returnedContents)
	if err != nil {
		return nil, err
	}
	return
}

func Release(key string) error {
	err := os.Remove(keyToFileName(key))
	if err != nil {
		return err
	}
	return nil
}

func keyToFileName(key string) string {
	if !strings.Contains(key, ".json") {
		key = fmt.Sprintf("%s.json", key)
	}
	return key
}
