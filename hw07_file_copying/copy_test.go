package main

import (
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

var tempFiles []*os.File

func TestFailCopy(t *testing.T) {
	defer cleanFiles(t)

	t.Run("Unsupported file", func(t *testing.T) {
		err := Copy("/dev/urandom", "", 0, 0)
		require.Equal(t, ErrUnsupportedFile, err)
	})

	t.Run("Nonexistent file", func(t *testing.T) {
		err := Copy("nonexistent.file", "", 0, 0)
		require.Error(t, err)
	})

	t.Run("Offset larger than file size", func(t *testing.T) {
		in := createTemp(t, "")
		err := Copy(in.Name(), "", 5, 0)
		require.Equal(t, ErrOffsetExceedsFileSize, err)
	})

	t.Run("Same file", func(t *testing.T) {
		in := createTemp(t, "")
		err := Copy(in.Name(), in.Name(), 0, 0)
		require.Equal(t, ErrSameFile, err)
	})
}

func TestSuccessCopy(t *testing.T) {
	defer cleanFiles(t)

	t.Run("Simple copy", func(t *testing.T) {
		in := createTemp(t, "test content")
		out := createTemp(t, "")
		err := Copy(in.Name(), out.Name(), 0, 0)
		require.NoError(t, err)
		content, err := os.ReadFile(out.Name())
		if err != nil {
			t.Fatal(err)
		}
		require.Equal(t, "test content", string(content))
	})

	t.Run("Offset copy", func(t *testing.T) {
		in := createTemp(t, "test content")
		out := createTemp(t, "")
		err := Copy(in.Name(), out.Name(), 5, 0)
		require.NoError(t, err)
		content, err := os.ReadFile(out.Name())
		if err != nil {
			t.Fatal(err)
		}
		require.Equal(t, "content", string(content))
	})

	t.Run("Limit copy", func(t *testing.T) {
		in := createTemp(t, "test content")
		out := createTemp(t, "")
		err := Copy(in.Name(), out.Name(), 0, 4)
		require.NoError(t, err)
		content, err := os.ReadFile(out.Name())
		if err != nil {
			t.Fatal(err)
		}
		require.Equal(t, "test", string(content))
	})

	t.Run("Offset limit copy", func(t *testing.T) {
		in := createTemp(t, "test content")
		out := createTemp(t, "")
		err := Copy(in.Name(), out.Name(), 5, 4)
		require.NoError(t, err)
		content, err := os.ReadFile(out.Name())
		if err != nil {
			t.Fatal(err)
		}
		require.Equal(t, "cont", string(content))
	})
}

func createTemp(t *testing.T, content string) *os.File {
	t.Helper()
	f, err := os.CreateTemp("", "example.*.txt")
	if err != nil {
		t.Fatal(err)
	}

	if len(content) > 0 {
		if _, err := f.Write([]byte(content)); err != nil {
			t.Fatal(err)
		}
	}

	tempFiles = append(tempFiles, f)

	return f
}

func cleanFiles(t *testing.T) {
	t.Helper()
	for _, file := range tempFiles {
		err := file.Close()
		if err != nil {
			t.Fatal(err)
		}
		err = os.Remove(file.Name())
		if err != nil {
			t.Fatal(err)
		}
	}

	tempFiles = []*os.File{}
}
