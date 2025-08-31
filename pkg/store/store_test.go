package store

import (
	"bytes"
	"errors"
	"testing"
)

func TestNewStore_IsEmpty(t *testing.T) {
	s := NewStore()

	if v, ok := s.Get("missing"); ok || v != nil {
		t.Fatalf("new store should have no keys; got ok=%v v=%v", ok, v)
	}
	if s.Exists("missing") {
		t.Fatalf("Exists should be false on new store")
	}
}

func TestSetThenGet_RoundTrip(t *testing.T) {
	s := NewStore()

	want := []byte("value")
	if err := s.Set("k", want); err != nil {
		t.Fatalf("Set error: %v", err)
	}
	got, ok := s.Get("k")
	if !ok {
		t.Fatalf("expected ok=true")
	}
	if !bytes.Equal(got, want) {
		t.Fatalf("Get mismatch: got=%q want=%q", got, want)
	}
}

func TestSet_EmptyKey(t *testing.T) {
	s := NewStore()

	err := s.Set("", []byte("x"))
	if !errors.Is(err, ErrEmptyKey) {
		t.Fatalf("expected ErrEmptyKey, got %v", err)
	}
}

func TestGet_MissingAndEmptyKey(t *testing.T) {
	s := NewStore()

	if v, ok := s.Get("missing"); ok || v != nil {
		t.Fatalf("missing key: want (nil,false), got v=%v ok=%v", v, ok)
	}
	if v, ok := s.Get(""); ok || v != nil {
		t.Fatalf("empty key should behave like missing: got v=%v ok=%v", v, ok)
	}
}

func TestExists_Sequence(t *testing.T) {
	s := NewStore()

	if s.Exists("k") {
		t.Fatalf("exists should be false before set")
	}
	_ = s.Set("k", []byte("v"))
	if !s.Exists("k") {
		t.Fatalf("exists should be true after set")
	}
	_ = s.Del("k")
	if s.Exists("k") {
		t.Fatalf("exists should be false after delete")
	}
}

func TestDel_Basics(t *testing.T) {
	s := NewStore()

	// deleting a missing key
	if s.Del("missing") {
		t.Fatalf("Del should return false for missing key")
	}

	// delete after set
	_ = s.Set("k", []byte("v"))
	if !s.Del("k") {
		t.Fatalf("Del should return true when key existed")
	}
	if v, ok := s.Get("k"); ok || v != nil {
		t.Fatalf("expected key removed; got v=%v ok=%v", v, ok)
	}
}

func TestEmptyKeyBehavior_ExistsAndDel(t *testing.T) {
	s := NewStore()

	if s.Exists("") {
		t.Fatalf("Exists(\"\") should be false")
	}
	if s.Del("") {
		t.Fatalf("Del(\"\") should be false")
	}
}

func TestCopyOnWrite_SetDoesNotAliasCallerSlice(t *testing.T) {
	s := NewStore()

	buf := []byte("abc")
	_ = s.Set("k", buf)
	buf[0] = 'X' // mutate caller's slice after Set

	got, ok := s.Get("k")
	if !ok {
		t.Fatalf("expected ok=true")
	}
	if !bytes.Equal(got, []byte("abc")) {
		t.Fatalf("copy-on-write violated: got=%q want=%q", got, "abc")
	}
}

func TestCopyOnRead_GetReturnsFreshSlice(t *testing.T) {
	s := NewStore()

	_ = s.Set("k", []byte("abc"))
	a, ok := s.Get("k")
	if !ok {
		t.Fatalf("expected ok=true")
	}
	a[0] = 'X' // mutate returned slice

	b, ok := s.Get("k")
	if !ok {
		t.Fatalf("expected ok=true on second get")
	}
	if !bytes.Equal(b, []byte("abc")) {
		t.Fatalf("copy-on-read violated: got=%q want=%q", b, "abc")
	}
}

func TestNilValue_RoundTrip(t *testing.T) {
	s := NewStore()

	if err := s.Set("k", nil); err != nil {
		t.Fatalf("Set error: %v", err)
	}
	v, ok := s.Get("k")
	if !ok {
		t.Fatalf("expected ok=true for nil value")
	}
	if v != nil {
		t.Fatalf("expected nil slice back, got non-nil len=%d", len(v))
	}
}
