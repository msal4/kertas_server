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

type PutImageOptions struct {
	ParentDir string
	Filename  string
	Upload    graphql.Upload
}

const (
	thumbnailSuffix = "_w200"
	hqSuffix        = "_w1000"
	jpgExt          = ".jpg"
)

func (s *Service) PutImage(ctx context.Context, opts PutImageOptions) (minio.UploadInfo, error) {
	if opts.Upload.File == nil {
		return minio.UploadInfo{}, errors.New("image file is required")
	}

	if opts.Filename == "" {
		opts.Filename = s.FormatFilename(opts.Upload.Filename, thumbnailSuffix)
	}
	opts.Filename = path.Join(opts.ParentDir, opts.Filename)

	if !strings.Contains(opts.Filename, "_w200") {
		idx := strings.LastIndex(opts.Filename, ".")
		opts.Filename = opts.Filename[:idx] + "_w200" + opts.Filename[idx:]
	}

	img, _, err := image.Decode(opts.Upload.File)
	if err != nil && err != io.EOF {
		return minio.UploadInfo{}, fmt.Errorf("unsupported image format: %v", err)
	}

	thumb := resize.Thumbnail(s.Config.ThumbnailSize.Width, s.Config.ThumbnailSize.Height, img, resize.Lanczos2)
	var thumbBuf bytes.Buffer
	if err := jpeg.Encode(&thumbBuf, thumb, nil); err != nil {
		return minio.UploadInfo{}, fmt.Errorf("encoding thumbnail: %v", err)
	}

	ext := opts.Filename[strings.LastIndex(opts.Filename, "."):]

	thumbInfo, err := s.MC.PutObject(ctx, s.Config.RootBucket, strings.Replace(opts.Filename, ext, jpgExt, 1),
		&thumbBuf, int64(thumbBuf.Len()), minio.PutObjectOptions{ContentType: "image/jpeg"})
	if err != nil {
		return thumbInfo, err
	}

	highQualityImg := resize.Thumbnail(s.Config.HQImageSize.Width, s.Config.HQImageSize.Height, img, resize.Lanczos2)
	var imgBuf bytes.Buffer
	if err := jpeg.Encode(&imgBuf, highQualityImg, nil); err != nil {
		return thumbInfo, fmt.Errorf("encoding high quality image: %v", err)
	}

	info, err := s.MC.PutObject(ctx, s.Config.RootBucket, strings.Replace(opts.Filename, thumbnailSuffix+ext, hqSuffix+jpgExt, 1),
		&imgBuf, int64(imgBuf.Len()), minio.PutObjectOptions{ContentType: "image/jpeg"})
	if err != nil {
		return info, err
	}

	return thumbInfo, nil
}
