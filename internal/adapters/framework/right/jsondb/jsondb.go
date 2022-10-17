package jsondb

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"sync"
	"time"
)

type Adapter struct {
	filename string
	mu       *sync.Mutex
}

type dbEntry struct {
	Date      time.Time `json:"date"`
	Answer    int32     `json:"answer"`
	Operation string    `json:"operation"`
}

type dbHistory []dbEntry

func NewAdapter(jsonFilepath string) (*Adapter, error) {

	if !(filepath.Ext(jsonFilepath) == ".json") {
		return nil, errors.New("invalid file extension, must be .json")
	}

	// open file
	file, err := os.OpenFile(jsonFilepath, os.O_RDWR|os.O_CREATE, 0755)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	return &Adapter{
		filename: jsonFilepath,
		mu:       &sync.Mutex{},
	}, nil
}

func (jdba Adapter) CloseDBConnection() error {
	return nil
}

func (jdba Adapter) AddToHistory(answer int32, operation string) error {
	jdba.mu.Lock()
	defer jdba.mu.Unlock()

	oldFileContents, err := os.ReadFile(jdba.filename)
	if err != nil {
		return err
	}

	var history dbHistory
	err = json.Unmarshal(oldFileContents, &history)
	if err != nil {
		fmt.Printf("err: %T\n", err)
		var syntaxError *json.SyntaxError
		var unmarshalTypeError *json.UnmarshalTypeError
		if errors.As(err, &syntaxError) || errors.As(err, &unmarshalTypeError) {
			fmt.Println("json syntax error")
			history = dbHistory{}
		} else {
			return err
		}
	}

	history = append(history, dbEntry{
		Date:      time.Now(),
		Answer:    answer,
		Operation: operation,
	})

	newFileContents, err := json.MarshalIndent(history, "", "  ")
	if err != nil {
		return err
	}

	err = os.WriteFile(jdba.filename, newFileContents, 0755)
	if err != nil {
		return err
	}
	return nil
}
