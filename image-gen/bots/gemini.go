package bots

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"google.golang.org/genai"
)

type Bot struct {
	GeminiAPIKey string
}

func NewBot(apiKey string) *Bot {
	return &Bot{GeminiAPIKey: apiKey}
}

func (b *Bot) CreateClient(ctx context.Context) *genai.Client {
	client, err := genai.NewClient(ctx, &genai.ClientConfig{
		APIKey:  b.GeminiAPIKey,
		Backend: genai.BackendGeminiAPI,
	})

	if err != nil {
		log.Fatalf("failed to create genai client: %v", err)
	}

	return client
}

func (b *Bot) ListModels() {
	ctx := context.Background()
	client := b.CreateClient(ctx)

	iter := client.Models.All(ctx)

	for model, err := range iter {
		if err != nil {
			fmt.Println("Error:", err)
			continue
		}
		fmt.Printf("Name: %s\n", model.Name)
		fmt.Printf("Description: %s\n", model.Description)
		fmt.Printf("SupportedActions: %s\n", model.SupportedActions)
		fmt.Println()

	}

	fmt.Println("Exiting")
}

func (b *Bot) GeneratePrompt(prompt string) string {
	ctx := context.Background()
	client := b.CreateClient(ctx)

	result, err := client.Models.GenerateContent(
		ctx,
		"gemini-2.0-flash",
		genai.Text(prompt),
		nil,
	)

	if err != nil {
		log.Fatalf("failed to generate text: %v", err)
	}

	fmt.Println("Generated Prompt:", result.Text())
	return result.Text()
}

// GenImage uses Gemini to generate an image for twitter bots
func (b *Bot) GenImage(prompt string) {
	ctx := context.Background()
	client := b.CreateClient(ctx)

	config := &genai.GenerateImagesConfig{
		NumberOfImages: 1,
	}

	result, err := client.Models.GenerateImages(
		ctx,
		"imagen-3.0-generate-002",
		prompt,
		config,
	)

	if err != nil {
		log.Fatalf("failed to generate image: %v", err)
	}

	for _, part := range result.GeneratedImages {
		outputFilename := fmt.Sprintf("gemini_image_%d.png", time.Now().Unix())
		if err := os.WriteFile(outputFilename, part.Image.ImageBytes, 0644); err != nil {
			log.Printf("failed to write file: %v", err)
		}
	}
}
