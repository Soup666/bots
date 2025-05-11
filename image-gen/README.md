Go Image gen

# Go Image Generator

Simple go cli program to generate images for a star wars ai image bot.

# Usage

Add `GEMINI_API_KEY=` to tools/.env file. Then run with `cd tools && go run . -h`.

```
❯ go run . -h
NAME:
   Soup666 Image Gen - Generate images intended for twitter bots

USAGE:
   Soup666 Image Gen [global options]

GLOBAL OPTIONS:
   --ship      Generate Ship Image (default: false)
   --atat      Generate ATAT Image (default: false)
   --list      List Gemini Models (default: false)
   --help, -h  show help
```