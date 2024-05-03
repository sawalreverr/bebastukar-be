package helper

import (
	"context"
	"log"
	"time"

	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
	"github.com/sawalreverr/bebastukar-be/config"
)

func UploadToCloudinary(file interface{}, folderPath string) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	conf := config.GetConfig().Cloudinary
	cld, err := cloudinary.NewFromParams(conf.CloudName, conf.ApiKey, conf.ApiSecret)
	if err != nil {
		log.Fatal("Cloudinary Enviroment Needed!")
		return "", err
	}

	resp, err := cld.Upload.Upload(ctx, file, uploader.UploadParams{Folder: folderPath})
	if err != nil {
		return "", err
	}

	return resp.SecureURL, nil
}
