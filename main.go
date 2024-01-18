package main

import (
	"fmt"
	"os"
	"os/exec"
)

func main() {
	// Get environment variables
	sourceRegistry := os.Getenv("PLUGIN_SOURCE_REGISTRY")
	sourceUsername := os.Getenv("PLUGIN_SOURCE_USERNAME")
	sourcePassword := os.Getenv("PLUGIN_SOURCE_PASSWORD")
	sourceNamespace := os.Getenv("PLUGIN_SOURCE_NAMESPACE")

	destinationRegistry := os.Getenv("PLUGIN_DESTINATION_REGISTRY")
	destinationUsername := os.Getenv("PLUGIN_DESTINATION_USERNAME")
	destinationPassword := os.Getenv("PLUGIN_DESTINATION_PASSWORD")
	destinationNamespace := os.Getenv("PLUGIN_DESTINATION_NAMESPACE")

	imageName := os.Getenv("PLUGIN_IMAGE_NAME")
	imageTag := os.Getenv("PLUGIN_IMAGE_TAG")

	// Authenticate and transfer image
	authenticateAndTransferImage(sourceRegistry, sourceUsername, sourcePassword, sourceNamespace, destinationRegistry, destinationUsername, destinationPassword, destinationNamespace, imageName, imageTag)
}

func authenticateAndTransferImage(sourceRegistry, sourceUsername, sourcePassword, sourceNamespace, destinationRegistry, destinationUsername, destinationPassword, destinationNamespace, imageName, imageTag string) {
	if sourceUsername != "" && sourcePassword != "" {
		execCommand("skopeo", "login", "--username", sourceUsername, "--password", sourcePassword, sourceRegistry)
	}

	if destinationUsername != "" && destinationPassword != "" {
		execCommand("skopeo", "login", "--username", destinationUsername, "--password", destinationPassword, destinationRegistry)
	}

	sourceImage := fmt.Sprintf("docker://%s/%s/%s:%s", sourceRegistry, sourceNamespace, imageName, imageTag)
	destinationImage := fmt.Sprintf("docker://%s/%s", destinationRegistry, destinationNamespace)

	execCommand("skopeo", "copy", sourceImage, destinationImage)
}

func execCommand(name string, arg ...string) {
	cmd := exec.Command(name, arg...)
	output, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Printf("Error executing %s: %s\n", name, err)
		fmt.Println(string(output))
		os.Exit(1)
	}

	fmt.Println(string(output))
}
