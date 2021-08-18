package service_test

import (
	"bufio"
	"context"
	"testing"

	"github.com/minio/minio-go/v7"
	"github.com/stretchr/testify/require"
)

func TestRemoveDir(t *testing.T) {
	s := newService(t)
	ctx := context.Background()

	t.Run("non existing dir", func(t *testing.T) {
		err := s.RemoveDir(ctx, "this/dir/does/not/exist")
		require.NoError(t, err)
	})

	t.Run("dir with objects", func(t *testing.T) {
		require := require.New(t)

		dir := s.FormatFilename("commondir34534534322222", "test") + "/"
		dir2 := s.FormatFilename("anotherdir", "test") + "/"

		obj1, err := s.MC.PutObject(ctx, s.Config.RootBucket, dir+"dir/file.whatever", bufio.NewReader(nil), 0, minio.PutObjectOptions{})
		require.NoError(err)
		obj2, err := s.MC.PutObject(ctx, s.Config.RootBucket, dir+"dir/file2.whatever", bufio.NewReader(nil), 0, minio.PutObjectOptions{})
		require.NoError(err)
		obj3, err := s.MC.PutObject(ctx, s.Config.RootBucket, dir+"dir/file3.whatever", bufio.NewReader(nil), 0, minio.PutObjectOptions{})
		require.NoError(err)
		obj4, err := s.MC.PutObject(ctx, s.Config.RootBucket, dir2+"dir/file3.whatever", bufio.NewReader(nil), 0, minio.PutObjectOptions{})
		require.NoError(err)

		_, err = s.MC.StatObject(ctx, s.Config.RootBucket, obj1.Key, minio.StatObjectOptions{})
		require.NoError(err)
		_, err = s.MC.StatObject(ctx, s.Config.RootBucket, obj2.Key, minio.StatObjectOptions{})
		require.NoError(err)
		_, err = s.MC.StatObject(ctx, s.Config.RootBucket, obj3.Key, minio.StatObjectOptions{})
		require.NoError(err)
		_, err = s.MC.StatObject(ctx, s.Config.RootBucket, obj4.Key, minio.StatObjectOptions{})
		require.NoError(err)

		err = s.RemoveDir(ctx, dir)
		require.NoError(err)

		_, err = s.MC.StatObject(ctx, s.Config.RootBucket, obj1.Key, minio.StatObjectOptions{})
		require.Error(err)
		_, err = s.MC.StatObject(ctx, s.Config.RootBucket, obj2.Key, minio.StatObjectOptions{})
		require.Error(err)
		_, err = s.MC.StatObject(ctx, s.Config.RootBucket, obj3.Key, minio.StatObjectOptions{})
		require.Error(err)
		_, err = s.MC.StatObject(ctx, s.Config.RootBucket, obj4.Key, minio.StatObjectOptions{})
		require.NoError(err)
	})
}
