package bots

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

/*
func main() {
	dt := NewDrawThing()

	prompt := "bunch of carrots"

	if err := dt.GenerateAndSave(prompt); err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}
}
*/

const (
	DRAW_THINGS_URL = "http://127.0.0.1:7860/sdapi/v1/txt2img"
	BATCH_COUNT     = 5
	IMG_SIZE        = 512
)

type DrawThing struct {
	FileUtils   *FileUtils
	APIUtils    *APIUtils
	StringUtils *StringUtils
}

func NewDrawThing(fileUtils *FileUtils, apiUtils *APIUtils, stringUtils *StringUtils) *DrawThing {
	return &DrawThing{
		FileUtils:   fileUtils,
		APIUtils:    apiUtils,
		StringUtils: stringUtils,
	}
}

type RequestParams struct {
	Prompt         string `json:"prompt"`
	NegativePrompt string `json:"negative_prompt"`
	Seed           int    `json:"seed"`
	Steps          int    `json:"steps"`
	GuidanceScale  int    `json:"guidance_scale"`
	BatchCount     int    `json:"batch_count"`
	Width          int    `json:"width"`
	Height         int    `json:"height"`
}

type Response struct {
	Images []string `json:"images"`
}

type ImageGenerationConfig struct {
	Prompt         string
	NegativePrompt string
	BatchCount     int
	Steps          int
	GuidanceScale  int
	Width          int
	Height         int
	Seed           int
}

// DefaultConfig returns a default configuration for image generation
func DefaultConfig() ImageGenerationConfig {
	return ImageGenerationConfig{
		NegativePrompt: "(worst quality, low quality, normal quality, (variations):1.4), blur:1.5",
		BatchCount:     BATCH_COUNT,
		Steps:          20,
		GuidanceScale:  4,
		Width:          IMG_SIZE,
		Height:         IMG_SIZE,
		Seed:           -1,
	}
}

// GenerateImages generates images using the Draw Things API
func (dt *DrawThing) GenerateImages(config ImageGenerationConfig) ([]string, error) {
	params := RequestParams{
		Prompt:         config.Prompt,
		NegativePrompt: config.NegativePrompt,
		Seed:           config.Seed,
		Steps:          config.Steps,
		GuidanceScale:  config.GuidanceScale,
		BatchCount:     config.BatchCount,
		Width:          config.Width,
		Height:         config.Height,
	}

	jsonData, err := json.Marshal(params)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal request params: %w", err)
	}

	resp, err := http.Post(DRAW_THINGS_URL, "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, fmt.Errorf("failed to make API request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("API request failed with status: %d", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}

	var data Response
	if err := json.Unmarshal(body, &data); err != nil {
		return nil, fmt.Errorf("failed to parse JSON response: %w", err)
	}

	return data.Images, nil
}

// GenerateAndSave is a convenience method that generates and saves images
func (dt *DrawThing) GenerateAndSave(prompt string) error {
	config := DefaultConfig()
	config.Prompt = prompt

	images, err := dt.GenerateImages(config)
	if err != nil {
		return fmt.Errorf("failed to generate images: %w", err)
	}

	if err := dt.FileUtils.SaveImages(images, prompt); err != nil {
		return fmt.Errorf("failed to save images: %w", err)
	}

	fmt.Printf("Successfully generated and saved %d images for prompt: %s\n", len(images), prompt)
	return nil
}
