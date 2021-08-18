package service

import (
	"context"
	"fmt"

	"github.com/minio/minio-go/v7"
)

func (s *Service) RemoveDir(ctx context.Context, dir string) error {
	objCh := s.MC.ListObjects(ctx, s.Config.RootBucket, minio.ListObjectsOptions{Prefix: dir, Recursive: true})
	for obj := range s.MC.RemoveObjects(ctx, s.Config.RootBucket, objCh, minio.RemoveObjectsOptions{}) {
		if obj.Err != nil {
			return fmt.Errorf("removing object %q: %v", obj.ObjectName, obj.Err)
		}
	}

	err := s.MC.RemoveObject(ctx, s.Config.RootBucket, dir, minio.RemoveObjectOptions{})
	if err != nil {
		return fmt.Errorf("removing directory: %v", err)
	}

	return nil
}
