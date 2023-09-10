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
	msgs := [2]string{"What should one have in the morning for healthy breakfast?", "What should one have in the morning for light breakfast?"}
	for i := 0; i < len(msgs); i++ {
		msg, _ := client.Parse(&witai.MessageRequest{
			Query: msgs[i],
		})
		fmt.Printf("ME: %v\nBOT: %v\n", msgs[i], msg.Traits["wit_breakfast"][0].Value)
	}
}
