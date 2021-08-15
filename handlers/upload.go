package handlers

import (
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/gabriel-vasile/mimetype"
	"github.com/kzw200015/picbed/cache"
	"github.com/kzw200015/picbed/config"
	"github.com/labstack/echo/v4"
	uuid "github.com/satori/go.uuid"
)

type Image struct {
	Id   string `json:"id,omitempty" yaml:"id,omitempty"`
	Name string `json:"name,omitempty" yaml:"name,omitempty"`
	Url  string `json:"url,omitempty" yaml:"url,omitempty"`
}

func HandleUpload(c echo.Context) error {
	file, err := c.FormFile("file")
	if err != nil {
		return err
	}

	imgPath, dstPath, err := saveImage(file)
	if err != nil {
		os.Remove(dstPath)
		return err
	}

	id := uuid.Must(uuid.NewV4()).String()
	cache.Set(id, dstPath)

	return c.JSON(200, Image{
		Id:   id,
		Name: file.Filename,
		Url:  imgPath,
	})
}

func saveImage(file *multipart.FileHeader) (string, string, error) {
	src, err := file.Open()
	if err != nil {
		return "", "", err
	}

	dstFileName := uuid.Must(uuid.NewV4()).String() + filepath.Ext(file.Filename)
	dstPath := filepath.Join(config.Get().ImgSrc, dstFileName)
	dst, err := os.Create(dstPath)
	if err != nil {
		return "", "", err
	}

	if _, err = io.Copy(dst, src); err != nil {
		return "", "", err
	}
	src.Close()
	dst.Close()

	b, err := isImage(dstPath)
	if err != nil {
		return "", "", err
	}

	if !b {
		return "", "", echo.NewHTTPError(http.StatusBadRequest, "not image")
	}

	return filepath.ToSlash(filepath.Join(config.IMG_ROUTE_PREFIX, dstFileName)), dstPath, nil
}

func isImage(p string) (bool, error) {
	b, err := mimetype.DetectFile(p)
	if err != nil {
		return false, err
	}

	return strings.HasPrefix(b.String(), "image"), nil
}
