package bots

import (
	"encoding/base64"
	"fmt"
	"os"
	"regexp"
	"strings"
	"sync"
	"time"
)

type FileUtils struct {
}

func NewFileUtils() *FileUtils {
	return &FileUtils{}
}

func (fu *FileUtils) GenerateFileName(prompt string, timestamp int64, idx int) string {
	cleanPrompt := fu.CleanPrompt(prompt)
	if len(cleanPrompt) > MAX_FILE_NAME_LEN {
		cleanPrompt = cleanPrompt[:MAX_FILE_NAME_LEN]
	}
	return fmt.Sprintf("%s_%d_%d.png", cleanPrompt, timestamp, idx)
}

// SaveImages saves multiple images concurrently
func (fu *FileUtils) SaveImages(images []string, prompt string) error {
	timestamp := time.Now().Unix()
	var wg sync.WaitGroup
	errChan := make(chan error, len(images))

	for idx, imageData := range images {
		wg.Add(1)
		go func(imgData string, index int) {
			defer wg.Done()
			if err := fu.SaveImage(imgData, index, prompt, timestamp); err != nil {
				errChan <- fmt.Errorf("failed to save image %d: %w", index, err)
			}
		}(imageData, idx)
	}

	wg.Wait()
	close(errChan)

	// Check for any errors
	var errs []error
	for err := range errChan {
		errs = append(errs, err)
	}

	if len(errs) > 0 {
		return fmt.Errorf("encountered %d errors while saving images: %v", len(errs), errs[0])
	}

	return nil
}

// SaveImage saves a single image from base64 data
func (fu *FileUtils) SaveImage(imageData string, idx int, prompt string, timestamp int64) error {
	fileName := fu.GenerateFileName(prompt, timestamp, idx)

	imgBytes, err := base64.StdEncoding.DecodeString(imageData)
	if err != nil {
		return fmt.Errorf("failed to decode base64 image data: %w", err)
	}

	if err := os.WriteFile(fileName, imgBytes, 0644); err != nil {
		return fmt.Errorf("failed to write image file: %w", err)
	}

	fmt.Printf("Saved: %s\n", fileName)
	return nil
}

func (fu *FileUtils) CleanPrompt(prompt string) string {
	re := regexp.MustCompile(`[^\w]+`)
	cleaned := re.ReplaceAllString(prompt, "_")
	cleaned = strings.Trim(cleaned, "_")
	return strings.ToLower(cleaned)
}
