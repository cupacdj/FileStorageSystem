package main

import (
	"bytes"
	"fmt"
	"io"
	"testing"
)

func TestPathTransformFunc(t *testing.T) {
	key := "bestpicture"
	pathKey := CASPathTransformFunc(key)
	expectedFilename := "71056ad8aa24742ea41ea36fa2e3452a31636e82"
	expectedPathName := "71056/ad8aa/24742/ea41e/a36fa/2e345/2a316/36e82"
	if pathKey.PathName != expectedPathName {
		t.Errorf("Expected %s, got %s", pathKey.PathName, expectedPathName)
	}
	if pathKey.Filename != expectedFilename {
		t.Errorf("Expected %s, got %s", pathKey.Filename, expectedFilename)
	}
}

func TestStore(t *testing.T) {

	s := newStore()
	defer teardown(t, s)

	for i := 0; i < 50; i++ {

		key := fmt.Sprintf("foo_%d", i)
		data := []byte("some jpg bytes")

		if _, err := s.writeStream(key, bytes.NewReader(data)); err != nil {
			t.Error(err)
		}

		if ok := s.Has(key); !ok {
			t.Errorf("store should have key: %s", key)
		}

		_, r, err := s.Read(key)
		if err != nil {
			t.Error(err)
		}

		b, _ := io.ReadAll(r)
		if string(b) != string(data) {
			t.Errorf("want %s have %s", string(data), string(b))
		}

		//fmt.Printf("read data: %s\n", string(b))

		if err := s.Delete(key); err != nil {
			t.Error(err)
		}

		if ok := s.Has(key); ok {
			t.Errorf("store should not have key: %s", key)
		}

	}
}

func newStore() *Store {
	opts := StoreOpts{
		PathTransformFunc: CASPathTransformFunc,
	}
	return NewStore(opts)
}

func teardown(t *testing.T, s *Store) {
	fmt.Printf("teardown store with root: %s\n", s.Root)
	if err := s.Clear(); err != nil {
		t.Error(err)
	}
}
