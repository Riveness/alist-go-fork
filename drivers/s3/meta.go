package s3

import (
	"github.com/alist-org/alist/v3/internal/driver"
	"github.com/alist-org/alist/v3/internal/op"
)

type Addition struct {
	driver.RootFolderPath
	Bucket            string `json:"bucket" required:"true"`
	Endpoint          string `json:"endpoint" required:"true"`
	Region            string `json:"region"`
	AccessKeyID       string `json:"access_key_id" required:"true"`
	SecretAccessKey   string `json:"secret_access_key" required:"true"`
	CustomHost        string `json:"custom_host"`
	SignURLExpire     int    `json:"sign_url_expire" type:"number" default:"4"`
	Placeholder       string `json:"placeholder"`
	ForcePathStyle    bool   `json:"force_path_style"`
	ListObjectVersion string `json:"list_object_version" type:"select" options:"v1,v2" default:"v1"`
}

var config = driver.Config{
	Name:        "S3",
	LocalSort:   true,
	CheckStatus: true,
}

func New() driver.Driver {
	return &S3{}
}

func init() {
	op.RegisterDriver(config, New)
}
