package groupcount_test

import (
	"testing"

	"github.com/briangreenhill/groupcount"
)

func TestWaitGroupCount(t *testing.T) {
	var wg groupcount.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
	}()
	wg.Wait()
	if wg.Count() != 0 {
		t.Errorf("Count() = %v, want 0", wg.Count())
	}
}

func TestWaitGroupAdd(t *testing.T) {
	var wg groupcount.WaitGroup
	wg.Add(1)
	wg.Add(1)
	if wg.Count() != 2 {
		t.Errorf("Count() = %v, want 2", wg.Count())
	}
}

func TestWaitGroupDone(t *testing.T) {
	var wg groupcount.WaitGroup
	wg.Add(1)
	wg.Done()
	if wg.Count() != 0 {
		t.Errorf("Count() = %v, want 0", wg.Count())
	}
}
