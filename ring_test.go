// +build linux

package iouring

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestNew(t *testing.T) {
	var p Params
	r, err := New(2048, &p)
	require.NoError(t, err)
	require.NotNil(t, r)

	require.NotZero(t, r.sq.Size)
	require.NotNil(t, r.sq.Head)
	require.NotNil(t, r.sq.Tail)
	require.NotNil(t, r.sq.Mask)
	require.NotNil(t, r.sq.Entries)
	require.NotNil(t, r.sq.Flags)
	require.NotNil(t, r.sq.Dropped)
	require.NotNil(t, r.sq.Entries)

	require.NotZero(t, r.cq.Size)
	require.NotNil(t, r.cq.Head)
	require.NotNil(t, r.cq.Tail)
	require.NotNil(t, r.cq.Mask)
	require.NotNil(t, r.cq.Entries)
	require.NotNil(t, r.cq.Entries)
}

func TestNewRingInvalidSize(t *testing.T) {
	var p Params
	_, err := New(99999, &p)
	require.Error(t, err)
}
