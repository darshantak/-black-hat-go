package main

import (
	"io"
	"os"
	"path/filepath"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"gopkg.in/gographics/imagick.v3/imagick"
)

func convertImage(format string, inputPath string, outputPath string) error {
	mw := imagick.NewMagickWand()
	defer mw.Destroy()

	if err := mw.ReadImage(inputPath); err != nil {
		return err
	}

	mw.SetImageFormat(format)

	if err := mw.WriteImage(outputPath); err != nil {
		return err
	}

	return nil
}

func uploadAndConvertHandler(c echo.Context, format string) error {
	file, err := c.FormFile("image")
	if err != nil {
		return err
	}

	src, err := file.Open()
	if err != nil {
		return err
	}
	defer src.Close()

	inputPath := filepath.Join("uploads", uuid.New().String()+filepath.Ext(file.Filename))
	outputPath := inputPath + "." + format

	dst, err := os.Create(inputPath)
	if err != nil {
		return err
	}
	defer dst.Close()

	if _, err = io.Copy(dst, src); err != nil {
		return err
	}

	if err := convertImage(format, inputPath, outputPath); err != nil {
		return err
	}

	return c.File(outputPath)
}

func convertToJXLHandler(c echo.Context) error {
	return uploadAndConvertHandler(c, "jxl")
}

func convertToWebPHandler(c echo.Context) error {
	return uploadAndConvertHandler(c, "webp")
}

func main() {
	imagick.Initialize()
	defer imagick.Terminate()

	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.POST("/convert/jxl", convertToJXLHandler)
	e.POST("/convert/webp", convertToWebPHandler)

	e.Logger.Fatal(e.Start(":8080"))
}
