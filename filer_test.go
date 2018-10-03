package filer

import (
	"bytes"
	"testing"
)

// quick test

func TestTouch(t *testing.T) {
	// README.md does exist.
	Touch("README.md")

	// "test123" does not exist"
	Touch("test123")

}

func TestAppend(t *testing.T) {
	tc := []byte("test append works\n")
	Remove("test1234")
	Touch("test1234")
	Append("test1234", tc)
	Append("test1234", tc)
	b := Read("test1234")
	if bytes.Compare(b, append(tc, tc...)) != 0 {
		t.Logf("Got: %q, wanted %q", string(b), string(tc))
		t.Fail()
	}
}

func TestWrite(t *testing.T) {
	tc := []byte("test write works\n")
	Touch("test123")
	Write("test123", tc)
	b := Read("test123")
	if bytes.Compare(b, tc) != 0 {
		t.Logf("Got: %q, wanted %q", string(b), string(tc))
		t.Fail()
	}

}

func TestRemove(t *testing.T) {
	if err := Remove("test123"); err != nil {
		t.Error(err)
	}
}
