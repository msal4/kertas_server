package service

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"image"
	"image/jpeg"
	_ "image/png"
	"io"
	"path"
	"regexp"
	"strings"
	"time"

	"github.com/99designs/gqlgen/graphql"
	"github.com/minio/minio-go/v7"
	"github.com/msal4/hassah_school_server/util"
	"github.com/nfnt/resize"
)

var fileNameRegexp = regexp.MustCompile("[^a-zA-Z0-9]")

func (s *Service) FormatFilename(filename, suffix string) string {
	var b strings.Builder

	b.WriteString(time.Now().Format(s.Config.FilenameTimeFormat))
	formattedName := fileNameRegexp.ReplaceAll([]byte(filename), []byte{})
	if len(formattedName) > s.Config.MaxFilenameLen {
		formattedName = formattedName[:s.Config.MaxFilenameLen]
	}
	b.WriteByte('_')
	b.Write(formattedName)
	b.WriteByte('_')
	b.WriteString(util.RandomString(s.Config.RS, s.Config.RandStringLen))
	b.WriteString(suffix)
	if idx := strings.LastIndex(filename, "."); idx != -1 {
		b.WriteString(strings.ToLower(filename[idx:]))
	}

	return b.String()
}

func (s *Service) SaveImage(ctx context.Context, bucket, dir, filename string, imgSrc graphql.Upload) (minio.UploadInfo, error) {
	if imgSrc.File == nil {
		return minio.UploadInfo{}, errors.New("image file is required")
	}

	filename = s.FormatFilename(filename, "_original")

	info, err := s.MC.PutObject(ctx, s.Config.RootBucket, path.Join(dir, filename), imgSrc.File, imgSrc.Size, minio.PutObjectOptions{
		ContentType: imgSrc.ContentType,
	})
	if err != nil {
		return info, err
	}

	img, _, err := image.Decode(imgSrc.File)
	if err != nil && err != io.EOF {
		s.MC.RemoveObject(ctx, s.Config.RootBucket, info.Key, minio.RemoveObjectOptions{})
		return minio.UploadInfo{}, fmt.Errorf("unsupported image format: %v", err)
	}

	thumb := resize.Thumbnail(s.Config.ThumbnailSize.Width, s.Config.ThumbnailSize.Height, img, resize.Lanczos2)
	var thumbBuf bytes.Buffer
	if err := jpeg.Encode(&thumbBuf, thumb, nil); err != nil {
		return info, fmt.Errorf("encoding thumbnail: %v", err)
	}

	info, err = s.MC.PutObject(ctx, s.Config.RootBucket, path.Join(dir, strings.Replace(filename, "_original.", "_w200.", 1)), &thumbBuf, int64(thumbBuf.Len()), minio.PutObjectOptions{
		ContentType: "image/jpeg",
	})
	if err != nil {
		return info, err
	}

	highQualityImg := resize.Thumbnail(s.Config.HQImageSize.Width, s.Config.HQImageSize.Height, img, resize.Lanczos2)
	var imgBuf bytes.Buffer
	if err := jpeg.Encode(&imgBuf, highQualityImg, nil); err != nil {
		return info, fmt.Errorf("encoding high quality image: %v", err)
	}

	hqImgPath := path.Join(dir, strings.Replace(filename, "_original.", "_w1000.", 1))
	info, err = s.MC.PutObject(ctx, s.Config.RootBucket, hqImgPath, &imgBuf, int64(imgBuf.Len()), minio.PutObjectOptions{
		ContentType: "image/jpeg",
	})
	if err != nil {
		return info, err
	}

	return info, nil
}
