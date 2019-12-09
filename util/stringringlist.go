package util

import (
	"errors"
	"os"
	"path/filepath"
)

// StringRingList represents a ringed list of strings.
// You can retrieve items by calling Next().
type StringRingList struct {
	Items []string

	current int
}

// Next returns the next item form the ring list.
// Wrap around is done automatically, so no need to check current index.
// It will throw an error, if the list is empty.
func (list *StringRingList) Next() (string, error) {
	if len(list.Items) == 0 {
		return "", errors.New("Cannot call Next() on empty list")
	}

	item := list.Items[list.current]

	if list.current >= len(list.Items)-1 {
		list.current = 0
	} else {
		list.current = list.current + 1
	}

	return item, nil
}

// NewFileRingList will construct a new StringRingList containing the relative pathes to
// all files within the given directory.
func NewFileRingList(root string) StringRingList {
	var files []string

	filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if info.IsDir() {
			return nil
		}
		files = append(files, path)
		return nil
	})

	return StringRingList{
		Items: files,
	}
}
