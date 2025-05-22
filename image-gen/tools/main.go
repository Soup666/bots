package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/urfave/cli/v3"
	"soup666.com/bots"
)

func main() {
	cmd := &cli.Command{
		Name:  "Soup666 Image Gen",
		Usage: "Generate images intended for twitter bots",
		Flags: []cli.Flag{
			&cli.BoolFlag{
				Name:  "generate",
				Usage: "Use Gemini to generate a prompt",
			},
			&cli.BoolFlag{
				Name:  "ship",
				Usage: "Generate Ship Image",
			},
			&cli.BoolFlag{
				Name:  "atat",
				Usage: "Generate ATAT Image",
			},

			&cli.BoolFlag{
				Name:  "list",
				Usage: "List Gemini Models",
			},
			&cli.BoolFlag{
				Name:  "drawthing",
				Usage: "Generate Drawthing Image",
			},
		},
		Action: func(ctx context.Context, c *cli.Command) error {

			// Handle System variables
			err := godotenv.Load()
			if err != nil {
				log.Fatalf("err loading: %v", err)
			}

			// List gemini models
			if c.Bool("list") {
				bot, err := LoadGeminiAPIKey()

				if err != nil {
					log.Fatalf("Error loading Gemini API key: %v", err)
				}

				bot.ListModels()
				return nil
			}

			if c.Bool("ship") {
				bot, err := LoadGeminiAPIKey()

				if err != nil {
					log.Fatalf("Error loading Gemini API key: %v", err)
				}

				bot.GenImage("Generate me a dramatic, realistic star wars image of a sith star destroyer. A large planet is looming in the background. The lighting is very harsh, creating a tense atmosphere. The image should give off serious vibes.")
				return nil
			}

			if c.Bool("atat") {
				bot, err := LoadGeminiAPIKey()

				if err != nil {
					log.Fatalf("Error loading Gemini API key: %v", err)
				}

				bot.GenImage("Generate me a dramatic, realistic star wars image of an ATAT on hoth. A large moon is looming in the background. The lighting is very harsh, creating a tense atmosphere. The image should give off serious vibes.")
				return nil
			}

			if c.Bool("drawthing") {
				fileUtils := bots.NewFileUtils()
				apiUtils := bots.NewAPIUtils()
				stringUtils := bots.NewStringUtils()
				dt := bots.NewDrawThing(
					fileUtils,
					apiUtils,
					stringUtils,
				)

				var prompt string

				if !c.Bool("generate") {
					fmt.Println("No prompt provided, using default.")
					prompt = "Generate me a dramatic, realistic star wars image of an ATAT on hoth. A large moon is looming in the background. The lighting is very harsh, creating a tense atmosphere. The image should give off serious vibes."
				} else {
					fmt.Println("Generate prompt defined.")
					bot, err := LoadGeminiAPIKey()
					if err != nil {
						log.Fatalf("Error loading Gemini API key: %v", err)
					}
					prompt = bot.GeneratePrompt("You are a bot that generates prompts for drawthing. Generate me a prompt that is a description for a star wars image. It should be empire focused, maybe a ship or character, or even a scene. Return only the prompt, no other text.")
					prompt += " The image should be dramatic, realistic, and give off serious vibes. The lighting should be very harsh, creating a tense atmosphere."
				}

				fmt.Printf("Prompt: %s\n", prompt)

				if err := dt.GenerateAndSave(prompt); err != nil {
					fmt.Printf("Error: %v\n", err)
					os.Exit(1)
				}
			}

			return nil
		},
	}

	if err := cmd.Run(context.Background(), os.Args); err != nil {
		log.Fatal(err)
	}
}

func LoadGeminiAPIKey() (*bots.Bot, error) {
	apiKey := os.Getenv("GEMINI_API_KEY")
	if apiKey == "" {
		return nil, errors.New("GEMINI_API_KEY environment variable not set")
	}

	return bots.NewBot(apiKey), nil
}
