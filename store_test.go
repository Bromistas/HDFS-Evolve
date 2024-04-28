package main

import (
	"bytes"
	"io/ioutil"
	"testing"
)

func TestPathTransformFunc(t *testing.T) {
	key := "momsbestpicture"
	pathKey := CASPathTransformFunc(key)
	expectedOriginalKey := "6804429f74181a63c50c3d81d733a12f14a353ff"
	expectedPathName := "68044/29f74/181a6/3c50c/3d81d/733a1/2f14a/353ff"

	if pathKey.PathName != expectedPathName {
		t.Errorf("Expected %s, got %s", pathKey.PathName, expectedPathName)
	}
	if pathKey.Filename != expectedOriginalKey {
		t.Errorf("Expected %s, got %s", pathKey.Filename, expectedOriginalKey)
	}

}

func TestStoreDeleteKey(t *testing.T) {
	opts := StoreOpts{
		PathTransform: CASPathTransformFunc,
	}

	s := NewStore(opts)

	key := "momsspecials"
	data := []byte("some jpg bytes")

	if err := s.writeStream(key, bytes.NewReader(data)); err != nil {
		t.Error(err)
	}

	if err := s.Delete(key); err != nil {
		t.Error(err)
	}
}

func TestStore(t *testing.T) {
	opts := StoreOpts{
		PathTransform: CASPathTransformFunc,
	}

	s := NewStore(opts)

	key := "momsspecials"
	data := []byte("some jpg bytes")

	if err := s.writeStream(key, bytes.NewReader(data)); err != nil {
		t.Error(err)
	}

	r, err := s.Read(key)
	if err != nil {
		t.Error(err)
	}

	b, _ := ioutil.ReadAll(r)
	if string(b) != string(data) {
		t.Errorf("want %s have %s", data, b)
	}

	if s.Has(key) == false {
		t.Errorf("Expected key to exist")
	}

	s.Delete(key)

	if s.Has(key) {
		t.Errorf("Expected key to be deleted")
	}
}
