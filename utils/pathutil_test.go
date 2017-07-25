package utils

import (
	"github.com/stretchr/testify/assert"
	"sort"
	"testing"
	"testing/quick"
)

func TestCurrentPath(t *testing.T) {
	cp := CurrentPath()
	assert.NotNil(t, cp)
}

func TestDirectorySort(t *testing.T) {
	vd := NewVcsDirectory()
	vd.DirPath = append(vd.DirPath,
		"/tmp/123",
		"/tmp",
		"/tmp/abc",
	)

	beforeSorted := []string{"/tmp/123", "/tmp", "/tmp/abc"}

	f := func() bool {
		for idx, element := range vd.DirPath {
			if element != beforeSorted[idx] {
				return false
			}
		}
		return true
	}

	if err := quick.Check(f, &quick.Config{MaxCount: 1}); err != nil {
		t.Error(err)
	}

	if !sort.StringsAreSorted(vd.DirPath) {
		sort.Strings(vd.DirPath)
	}

	sorted := []string{"/tmp", "/tmp/123", "/tmp/abc"}
	f = func() bool {
		for idx, element := range vd.DirPath {
			if element != sorted[idx] {
				return false
			}
		}
		return true
	}

	if err := quick.Check(f, &quick.Config{MaxCount: 1}); err != nil {
		t.Error(err)
	}

}
