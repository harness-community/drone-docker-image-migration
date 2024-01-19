package env

import (
	"fmt"
	"os"
)

func VerifyEnvVars() error {
	// Get environment variables
	sourceRegistry := os.Getenv("PLUGIN_SOURCE_REGISTRY")
	sourceUsername := os.Getenv("PLUGIN_SOURCE_USERNAME")
	sourcePassword := os.Getenv("PLUGIN_SOURCE_PASSWORD")
	sourceNamespace := os.Getenv("PLUGIN_SOURCE_NAMESPACE")

	destinationRegistry := os.Getenv("PLUGIN_DESTINATION_REGISTRY")
	destinationUsername := os.Getenv("PLUGIN_DESTINATION_USERNAME")
	destinationNamespace := os.Getenv("PLUGIN_DESTINATION_NAMESPACE")

	imageName := os.Getenv("PLUGIN_IMAGE_NAME")
	imageTag := os.Getenv("PLUGIN_IMAGE_TAG")

	// verify source variables
	if sourceRegistry == "" || sourceUsername == "" || sourcePassword == "" || sourceNamespace == "" {
		return fmt.Errorf("missing source variables")
	}

	// verify image variables
	if imageName == "" || imageTag == "" {
		return fmt.Errorf("missing image variables")
	}

	// verify destination variables
	if destinationRegistry == "" || destinationUsername == "" || destinationNamespace == "" {
		return fmt.Errorf("missing destination variables")
	}

	return nil
}
