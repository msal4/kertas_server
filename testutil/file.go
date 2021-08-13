package testutil

import (
	"io/fs"
	"os"
	"testing"

	"github.com/msal4/hassah_school_server/util"
	"github.com/stretchr/testify/require"
)

type File struct {
	*os.File
	fs.FileInfo
	ContentType string
}

func OpenFile(t testing.TB, name string) *File {
	t.Helper()

	f, err := os.Open(name)
	require.NoError(t, err)
	contentType, err := util.GetFileContentType(f)
	require.NoError(t, err)
	stat, err := f.Stat()
	require.NoError(t, err)
	f.Seek(0, 0)

	return &File{File: f, ContentType: contentType, FileInfo: stat}
}
