package goutil

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strings"
	"time"
)

const (
	Hourly  = "hourly"
	Daily   = "daily"
	Weekly  = "weekly"
	Monthly = "monthly"
	Yearly  = "yearly"
)

func WriteStringToFile(data, path, fileName, ext, rotation string) error {
	path = strings.TrimSpace(strings.Trim(path, "/"))

	if len(path) <= 0 {
		return errors.New("file path is empty")
	}

	fileName = strings.TrimSpace(fileName)

	if len(fileName) <= 0 {
		return errors.New("filename is empty")
	}

	ext = strings.TrimSpace(strings.TrimLeft(ext, "."))

	rotation = strings.ToLower(strings.TrimSpace(rotation))

	switch rotation {
	case Hourly:
		fileName = fmt.Sprintf("%s-%s", fileName, time.Now().Format("2006-01-02-15"))
	case Daily:
		fileName = fmt.Sprintf("%s-%s", fileName, time.Now().Format("2006-01-02"))
	case Weekly:
		fileName = fmt.Sprintf("%s-%s", fileName, time.Now().Format("2006-W"))
	case Monthly:
		fileName = fmt.Sprintf("%s-%s", fileName, time.Now().Format("2006-01"))
	case Yearly:
		fileName = fmt.Sprintf("%s-%s", fileName, time.Now().Format("2006"))
	}

	if len(ext) > 0 {
		ext = fmt.Sprintf(".%s", ext)
	}

	fileName = fmt.Sprintf("%s/%s%s", path, fileName, ext)

	f, err := os.OpenFile(fileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}

	defer func() {
		_ = f.Close()
	}()

	_, err = f.WriteString(fmt.Sprintf("%s\n", data))
	if err != nil {
		return err
	}

	return nil
}

func WriteJsonToFile(data interface{}, path, fileName, ext, rotation string) error {
	byteData, err := json.Marshal(data)
	if err != nil {
		return err
	}

	return WriteStringToFile(string(byteData), path, fileName, ext, rotation)
}
