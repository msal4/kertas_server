package service_test

import (
	"context"
	"strings"
	"testing"

	"github.com/99designs/gqlgen/graphql"
	"github.com/minio/minio-go/v7"
	"github.com/msal4/hassah_school_server/service"
	"github.com/msal4/hassah_school_server/testutil"
	"github.com/stretchr/testify/require"
)

func TestPutImage(t *testing.T) {
	s := newService(t)
	ctx := context.Background()

	f := testutil.OpenFile(t, "../testfiles/harvard.jpg")

	t.Run("top level", func(t *testing.T) {
		f.Seek(0, 0)
		info, err := s.PutImage(ctx, service.PutImageOptions{
			Upload: graphql.Upload{File: f, Filename: f.File.Name(), Size: f.Size(), ContentType: f.ContentType},
		})
		require.NoError(t, err)
		require.Contains(t, info.Key, "_w200.jpg")

		_, err = s.MC.StatObject(ctx, s.Config.RootBucket, info.Key, minio.StatObjectOptions{})
		require.NoError(t, err)

		_, err = s.MC.StatObject(ctx, s.Config.RootBucket, strings.Replace(info.Key, "_w200", "_w1000", 1), minio.StatObjectOptions{})
		require.NoError(t, err)
	})

	t.Run("in directory", func(t *testing.T) {
		f.Seek(0, 0)

		const dir = "parentdir1"
		info, err := s.PutImage(ctx, service.PutImageOptions{
			ParentDir: dir,
			Upload:    graphql.Upload{File: f, Filename: f.File.Name(), Size: f.Size(), ContentType: f.ContentType},
		})
		require.NoError(t, err)
		require.Contains(t, info.Key, "_w200.jpg")
		require.Equal(t, info.Key[:10], dir)

		_, err = s.MC.StatObject(ctx, s.Config.RootBucket, info.Key, minio.StatObjectOptions{})
		require.NoError(t, err)

		_, err = s.MC.StatObject(ctx, s.Config.RootBucket, strings.Replace(info.Key, "_w200", "_w1000", 1), minio.StatObjectOptions{})
		require.NoError(t, err)
	})

	t.Run("invalid", func(t *testing.T) {
		f := testutil.OpenFile(t, "../testfiles/file.txt")

		_, err := s.PutImage(ctx, service.PutImageOptions{
			Upload: graphql.Upload{File: f, Filename: f.File.Name(), Size: f.Size(), ContentType: f.ContentType},
		})
		require.Error(t, err)
	})

	t.Run("custom name & path", func(t *testing.T) {
		f.Seek(0, 0)

		const dir = "parentdir2"
		info, err := s.PutImage(ctx, service.PutImageOptions{
			ParentDir: dir,
			Filename:  "image1.jpg",
			Upload:    graphql.Upload{File: f, Filename: f.File.Name(), Size: f.Size(), ContentType: f.ContentType},
		})
		require.NoError(t, err)
		require.Equal(t, dir+"/image1_w200.jpg", info.Key)

		_, err = s.MC.StatObject(ctx, s.Config.RootBucket, info.Key, minio.StatObjectOptions{})
		require.NoError(t, err)

		_, err = s.MC.StatObject(ctx, s.Config.RootBucket, strings.Replace(info.Key, "_w200", "_w1000", 1), minio.StatObjectOptions{})
		require.NoError(t, err)

		f.Seek(0, 0)

		const dir2 = "parentdir3"
		info, err = s.PutImage(ctx, service.PutImageOptions{
			ParentDir: dir2,
			Filename:  "image2_w200.jpg",
			Upload:    graphql.Upload{File: f, Filename: f.File.Name(), Size: f.Size(), ContentType: f.ContentType},
		})
		require.NoError(t, err)
		require.Equal(t, dir2+"/image2_w200.jpg", info.Key)

		_, err = s.MC.StatObject(ctx, s.Config.RootBucket, info.Key, minio.StatObjectOptions{})
		require.NoError(t, err)

		_, err = s.MC.StatObject(ctx, s.Config.RootBucket, strings.Replace(info.Key, "_w200", "_w1000", 1), minio.StatObjectOptions{})
		require.NoError(t, err)

	})
}
