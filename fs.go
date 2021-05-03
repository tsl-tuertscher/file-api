package main

import (
	"os"
    "strings"
	"io"
    "fmt"
    "net/http"
    "archive/zip"
    "path/filepath"
)

func UnzipFile(src string, target string) error {
    zipReader, _ := zip.OpenReader(src)
	for _, file := range zipReader.Reader.File {
 
		zippedFile, err := file.Open()
		if err != nil {
            return err
		}
		defer zippedFile.Close()
 
		extractedFilePath := filepath.Join(
			target,
			file.Name,
		)
 
		if file.FileInfo().IsDir() {
			os.MkdirAll(extractedFilePath, file.Mode())
		} else { 
			outputFile, err := os.OpenFile(
				extractedFilePath,
				os.O_WRONLY|os.O_CREATE|os.O_TRUNC,
				file.Mode(),
			)
			if err != nil {
				return err
			}
			defer outputFile.Close()
 
			_, err = io.Copy(outputFile, zippedFile)
			if err != nil {
				return err
			}
		}
	}
    return nil
}

func DownloadTileSource(config Config, source string, url string) (string, error) {
    // Create the file
    dir := [2]string{config.WorkingDirectory, source}
    err := os.MkdirAll(strings.Join((dir)[:], "/"), 0755)
    if err != nil {
        return "", err
    }
    
    dest := [3]string{config.WorkingDirectory, source, "default.zip"}

    out, err := os.Create(strings.Join(dest[:], "/"))
    if err != nil  {
        return "", err
    }
    defer out.Close()

    // Get the data
    resp, err := http.Get(url)
    if err != nil {
        return "", err
    }
    defer resp.Body.Close()

    // Check server response
    if resp.StatusCode != http.StatusOK {
        return "", fmt.Errorf("bad status: %s", resp.Status)
    }

    // Writer the body to file
    _, err = io.Copy(out, resp.Body)
    if err != nil  {
        return "", err
    }

    return strings.Join(dest[:], "/"), nil  
}

func GetTileUrl(config Config, source string, z string, x string, y string, fileType string) string {
    src := [5]string{config.Offset, source, z, x, y + "." + fileType}
	if FileExists(strings.Join(src[:], "/")) {
        return strings.Join(src[:], "/")

	} else {
        src := [3]string{config.Offset, source, "default." + fileType}
        return strings.Join(src[:], "/")
    }
}

func GetMimeTypeFromFileType(fileType string) string {
    switch fileType {
    case "jpg":
    case "jpeg":
        return "image/jpeg"
    case "png":
        return "image/png"
    case "json":
        return "application/json" 
    }
    return "application/octet-stream"
}

func FileExists(filename string) bool {
    info, err := os.Stat(filename)
    if os.IsNotExist(err) {
        return false
    }
    return !info.IsDir()
}
