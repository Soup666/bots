package main

import (
	"context"
	"errors"
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
		},
		Action: func(ctx context.Context, c *cli.Command) error {

			// Handle System variables
			err := godotenv.Load()
			if err != nil {
				log.Fatalf("err loading: %v", err)
			}

			geminiKey := os.Getenv("GEMINI_API_KEY")
			if geminiKey == "" {
				return errors.New("missing Gemini Key")
			}

			bot := bots.NewBot(geminiKey)

			// List gemini models
			if c.Bool("list") {
				bot.ListModels()
				return nil
			}

			if c.Bool("ship") {
				bot.GenImage("Generate me a dramatic, realistic star wars image of a sith star destroyer. A large planet is looming in the background. The lighting is very harsh, creating a tense atmosphere. The image should give off serious vibes.")
				return nil
			}

			if c.Bool("atat") {
				bot.GenImage("Generate me a dramatic, realistic star wars image of an ATAT on hoth. A large moon is looming in the background. The lighting is very harsh, creating a tense atmosphere. The image should give off serious vibes.")
				return nil
			}

			return nil
		},
	}

	if err := cmd.Run(context.Background(), os.Args); err != nil {
		log.Fatal(err)
	}
}
