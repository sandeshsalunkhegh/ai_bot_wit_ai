package main

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	witai "github.com/wit-ai/wit-go/v2"
)

func main() {
	godotenv.Load(".env")

	client := witai.NewClient(os.Getenv("BEARER_TOKEN"))
	msg, _ := client.Parse(&witai.MessageRequest{
		Query: "What should one have in the morning for healthy breakfast?",
	})
	fmt.Printf("%v\n", msg.Text)
}
