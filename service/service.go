package service

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"sync"
	"time"

	"github.com/google/uuid"
	"github.com/minio/minio-go/v7"
	"github.com/msal4/hassah_school_server/auth"
	"github.com/msal4/hassah_school_server/ent"
	"github.com/segmentio/ksuid"
)

type Config struct {
	// RS is the random source used to generate unique filenames.
	//
	// Default: `rand.NewSource(time.Now().Unix())`.
	RS rand.Source

	// RootBucket is the bucket used to store all the schools binary data.
	//
	// If not provided defaultRootBucket is used.
	RootBucket string `yaml:"root_bucket" env:"ROOT_BUCKET"`

	// MaxFilenameLen is used to clip the original filename.
	//
	// If not provided defaultMaxFilenameLen is used.
	MaxFilenameLen int `yaml:"max_filename_len" env:"MAX_FILENAME_LEN"`

	// RandStringLen is the length of the generated random string that is appended to the filename to make it unique.
	//
	// If not provided defaultRandStringLen is used.
	RandStringLen int `yaml:"rand_string_len" env:"RAND_STRING_LEN"`

	// FilenameTimeFormat is the format used for the time prefix in filename to increase the probability of it being unique.
	//
	// If not provided defaultFilenameTimeFormat is used.
	FilenameTimeFormat string `yaml:"filename_time_format" env:"FILENAME_TIME_FORMAT"`

	// ThumbnailSize is the size of the image thumbnail.
	//
	// If not provided defaultThumbnailSize is used.
	ThumbnailSize struct {
		Width  uint `yaml:"width" env:"THUMBNAIL_IMAGE_WIDTH"`
		Height uint `yaml:"height" env:"THUMBNAIL_IMAGE_HEIGHT"`
	} `yaml:"thumbnail_image_size"`

	// HQImageSize is the size of the high quality version of the image.
	//
	// If not provided defaultHQImageSize is used.
	HQImageSize struct {
		Width  uint `yaml:"width" env:"HQ_IMAGE_WIDTH"`
		Height uint `yaml:"height" env:"HQ_IMAGE_HEIGHT"`
	} `yaml:"hq_image_size"`

	// AuthConfig is used to configure the token generation.
	auth.AuthConfig `yaml:"auth"`
}

type Service struct {
	// EC is the entity client used to interact with the database.
	EC *ent.Client

	// MC is the minio client used to interact with the s3 compatible store.
	MC *minio.Client

	// Config is all of the server configuration.
	Config *Config

	msgChannels map[uuid.UUID]map[ksuid.KSUID]chan *ent.Message
	mu          sync.Mutex
}

// New creates a new initialized and configured service.
func New(ec *ent.Client, mc *minio.Client, cfg *Config) (*Service, error) {
	cfg = getConfig(cfg)

	ctx := context.Background()
	exists, err := mc.BucketExists(ctx, cfg.RootBucket)
	if err != nil {
		return nil, fmt.Errorf("checking if %q bucket exists: %v", cfg.RootBucket, err)
	}
	if !exists {
		log.Printf("bucket %q does not exist, creating one...\n", cfg.RootBucket)
		err := mc.MakeBucket(ctx, cfg.RootBucket, minio.MakeBucketOptions{})
		if err != nil {
			return nil, fmt.Errorf(`making %q bucket: %v`, cfg.RootBucket, err)
		}
		log.Printf("created bucket %q.\n", cfg.RootBucket)
	}

	return &Service{EC: ec, MC: mc, Config: cfg, msgChannels: make(map[uuid.UUID]map[ksuid.KSUID]chan *ent.Message)}, nil
}

// Config defaults.
const (
	defaultFilenameTimeFormat   = "02-01-06_15-04"
	defaultRootBucket           = "root"
	defaultMaxFilenameLen       = 6
	defaultRandStringLen        = 6
	defaultThumbnailSize        = 200
	defaultHQImageSize          = 1000
	defaultAccessKey            = "dontusethedefaultaccesskey"
	defaultRefreshKey           = "dontusethedefaultrefreshkey"
	defaultAccessTokenLifetime  = 2 * time.Minute
	defaultRefreshTokenLifetime = 1 * time.Hour
)

func getConfig(cfg *Config) *Config {
	if cfg == nil {
		cfg = &Config{}
	}

	if cfg.RS == nil {
		cfg.RS = rand.NewSource(time.Now().Unix())
	}

	if cfg.RootBucket == "" {
		cfg.RootBucket = defaultRootBucket
	}

	if cfg.FilenameTimeFormat == "" {
		cfg.FilenameTimeFormat = defaultFilenameTimeFormat
	}

	if cfg.RandStringLen <= 0 {
		cfg.RandStringLen = defaultRandStringLen
	}

	if cfg.MaxFilenameLen <= 0 {
		cfg.MaxFilenameLen = defaultMaxFilenameLen
	}

	if cfg.ThumbnailSize.Width == 0 {
		cfg.ThumbnailSize.Width = defaultHQImageSize
	}

	if cfg.ThumbnailSize.Height == 0 {
		cfg.ThumbnailSize.Height = defaultHQImageSize
	}

	if cfg.HQImageSize.Width == 0 {
		cfg.HQImageSize.Width = defaultHQImageSize
	}

	if cfg.HQImageSize.Height == 0 {
		cfg.HQImageSize.Height = defaultHQImageSize
	}

	if len(cfg.AccessSecretKey) == 0 {
		cfg.AccessSecretKey = defaultAccessKey
	}

	if len(cfg.RefreshSecretKey) == 0 {
		cfg.RefreshSecretKey = defaultRefreshKey
	}

	if cfg.AccessTokenLifetime == 0 {
		cfg.AccessTokenLifetime = defaultAccessTokenLifetime
	}

	if cfg.RefreshTokenLifetime == 0 {
		cfg.RefreshTokenLifetime = defaultRefreshTokenLifetime
	}

	return cfg
}
