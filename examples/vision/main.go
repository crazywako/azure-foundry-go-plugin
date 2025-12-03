package main

import (
	"context"
	"log"

	"github.com/firebase/genkit/go/ai"
	"github.com/firebase/genkit/go/genkit"
	azureaifoundry "github.com/xavidop/genkit-azure-foundry-go"
	"github.com/xavidop/genkit-azure-foundry-go/examples/common"
)

func main() {
	ctx := context.Background()

	// Setup Genkit with Azure AI Foundry
	g, azurePlugin, err := common.SetupGenkit(ctx, nil)
	if err != nil {
		log.Fatalf("Failed to setup Genkit: %v", err)
	}

	// Define a GPT-5 model with vision support
	gpt5Model := azurePlugin.DefineModel(g, azureaifoundry.ModelDefinition{
		Name:          "gpt-5", // Your deployment name in Azure
		Type:          "chat",
		SupportsMedia: true,
	}, nil)

	log.Println("Starting Vision (Multimodal) example...")
	log.Println("This example demonstrates how to analyze images")
	log.Println()

	// Analyze an image from a URL
	log.Println("===Analyzing an image from URL ===")
	imageURL := "https://upload.wikimedia.org/wikipedia/commons/thumb/d/dd/Gfp-wisconsin-madison-the-nature-boardwalk.jpg/2560px-Gfp-wisconsin-madison-the-nature-boardwalk.jpg"

	response1, err := genkit.Generate(ctx, g,
		ai.WithModel(gpt5Model),
		ai.WithMessages(ai.NewUserMessage(
			ai.NewTextPart("What's in this image? Describe it in detail."),
			ai.NewMediaPart("image/jpeg", imageURL),
		)),
	)

	if err != nil {
		log.Printf("Error analyzing image from URL: %v", err)
	} else {
		log.Printf("Response: %s\n", response1.Text())
	}
	log.Println()

	log.Println("Vision example completed!")
}
