package graph

import (
	"bytes"
	"context"
	"fmt"
	"image"
	"image/jpeg"
	_ "image/png"
	"io"
	"math/rand"
	"regexp"
	"strings"
	"time"

	"github.com/99designs/gqlgen/graphql"
	"github.com/minio/minio-go/v7"
	"github.com/msal4/hassah_school_server/ent"
	"github.com/msal4/hassah_school_server/graph/generated"
	"github.com/msal4/hassah_school_server/util"
	"github.com/nfnt/resize"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	Client     *ent.Client
	Minio      *minio.Client
	RandSource rand.Source
}

// TODO: create a Config struct in Resolver to hold these values.
const (
	maxFilenameLen = 6
	randStringLen  = 6
	timeFormat     = "02-01-06_15-04"
)

var fileNameRegexp = regexp.MustCompile("[^a-zA-Z0-9]")

func (r *Resolver) FormatFilename(filename, suffix string) string {
	var b strings.Builder

	b.WriteString(time.Now().Format(timeFormat))
	formattedName := fileNameRegexp.ReplaceAll([]byte(filename), []byte{})
	if len(formattedName) > maxFilenameLen {
		formattedName = formattedName[:maxFilenameLen]
	}
	b.WriteByte('_')
	b.Write(formattedName)
	b.WriteByte('_')
	b.WriteString(util.RandomString(r.RandSource, randStringLen))
	b.WriteString(suffix)
	if idx := strings.LastIndex(filename, "."); idx != -1 {
		b.WriteString(strings.ToLower(filename[idx:]))
	}

	return b.String()
}

const (
	thumbnailWidth  = 200
	thumbnailHeight = 200
	imageWidth      = 1000
	imageHeight     = 1000
)

func (r *Resolver) SaveImage(ctx context.Context, bucket, dir, filename string, imgSrc graphql.Upload) (minio.UploadInfo, error) {
	filename = r.FormatFilename(filename, "_original")

	info, err := r.Minio.PutObject(ctx, "images", filename, imgSrc.File, imgSrc.Size, minio.PutObjectOptions{
		ContentType: imgSrc.ContentType,
	})
	if err != nil {
		return info, err
	}

	img, _, err := image.Decode(imgSrc.File)
	if err != nil && err != io.EOF {
		return minio.UploadInfo{}, fmt.Errorf("unsupported image format: %v", err)
	}

	thumb := resize.Thumbnail(thumbnailWidth, thumbnailHeight, img, resize.Lanczos2)
	var thumbBuf bytes.Buffer
	if err := jpeg.Encode(&thumbBuf, thumb, nil); err != nil {
		return info, fmt.Errorf("encoding thumbnail: %v", err)
	}

	info, err = r.Minio.PutObject(ctx, "images", strings.Replace(filename, "_original.", "_w200.", 1), &thumbBuf, int64(thumbBuf.Len()), minio.PutObjectOptions{
		ContentType: "image/jpeg",
	})
	if err != nil {
		return info, err
	}

	highQImg := resize.Thumbnail(imageWidth, imageHeight, img, resize.Lanczos2)
	var imgBuf bytes.Buffer
	if err := jpeg.Encode(&imgBuf, highQImg, nil); err != nil {
		return info, fmt.Errorf("encoding high quality image: %v", err)
	}

	info, err = r.Minio.PutObject(ctx, "images", strings.Replace(filename, "_original.", "_w1000.", 1), &imgBuf, int64(imgBuf.Len()), minio.PutObjectOptions{
		ContentType: "image/jpeg",
	})
	if err != nil {
		return info, err
	}

	return info, nil
}

func NewSchema(ec *ent.Client, mc *minio.Client, src rand.Source) graphql.ExecutableSchema {
	return generated.NewExecutableSchema(generated.Config{
		Resolvers: &Resolver{
			Client:     ec,
			Minio:      mc,
			RandSource: src,
		},
	})
}
