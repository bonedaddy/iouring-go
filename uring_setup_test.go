// +build linux

package iouring

import (
	"syscall"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSetupInvalidEntries(t *testing.T) {
	var p Params
	fd, err := Setup(0, &p)
	if err == nil {
		t.Fatal("expected Setup to fail")
	}
	defer require.NoError(t, syscall.Close(fd))
	_, err = Setup(8192, &p)
	if err == nil {
		t.Fatal("expected Setup to fail")
	}
	_, err = Setup(9999, &p)
	if err == nil {
		t.Fatal("expected Setup to fail")
	}
}

func TestSetupValidEntries(t *testing.T) {
	var p Params
	fd, err := Setup(1024, &p)
	if err != nil {
		t.Fatal(err)
	}
	defer require.NoError(t, syscall.Close(fd))
	if fd <= 0 {
		t.Fatalf("expected valid fd, got: %d", fd)
	}
}

func TestMmapSubmitRing(t *testing.T) {
	var p Params
	fd, err := Setup(1024, &p)
	if err != nil {
		t.Fatal(err)
	}
	defer require.NoError(t, syscall.Close(fd))
	var sq SubmitQueue
	if err := MmapSubmitRing(fd, &p, &sq); err != nil {
		t.Fatal(err)
	}
}

func TestMmapCompletionRing(t *testing.T) {
	var p Params
	fd, err := Setup(1024, &p)
	if err != nil {
		t.Fatal(err)
	}
	defer require.NoError(t, syscall.Close(fd))
	var cq CompletionQueue
	if err := MmapCompletionRing(fd, &p, &cq); err != nil {
		t.Fatal(err)
	}
}
