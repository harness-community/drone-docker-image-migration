package main

import (
	"fmt"
	"os"
	"os/exec"
)

func main() {
	//  Get plugin settings
	sourceDockerRegistry := os.Getenv("PLUGIN_SOURCE_DOCKER_REGISTRY")
	destinationDockerRegistry := os.Getenv("PLUGIN_DESTINATION_DOCKER_REGISTRY")

	sourceUsername := os.Getenv("PLUGIN_SOURCE_USERNAME")
	sourcePassword := os.Getenv("PLUGIN_SOURCE_PASSWORD")
	sourceNamespace := os.Getenv("PLUGIN_SOURCE_NAMESPACE")

	destinationUsername := os.Getenv("PLUGIN_DESTINATION_USERNAME")
	destinationPassword := os.Getenv("PLUGIN_DESTINATION_PASSWORD")
	destinationNamespace := os.Getenv("PLUGIN_DESTINATION_NAMESPACE")

	imageName := os.Getenv("PLUGIN_IMAGE_NAME")
	imageTag := os.Getenv("PLUGIN_IMAGE_TAG")

	if imageTag == "" {
		imageTag = "latest"
	}

	if sourceDockerRegistry == "" || sourceUsername == "" || sourcePassword == "" || sourceNamespace == "" {
		fmt.Println("Source docker registry, namespace, username and password are required")
		os.Exit(1)
	}

	if destinationDockerRegistry == "" || destinationUsername == "" || destinationPassword == "" || destinationNamespace == "" {
		fmt.Println("Destination docker registry, namespace, username and password are required")
		os.Exit(1)
	}

	loginToDockerRegistry(sourceDockerRegistry, sourceUsername, sourcePassword)

	pullImage(sourceDockerRegistry, imageName, imageTag, sourceNamespace)

	tagImage(sourceDockerRegistry, destinationDockerRegistry, imageName, imageTag, sourceNamespace, destinationNamespace)

	loginToDockerRegistry(destinationDockerRegistry, destinationUsername, destinationPassword)

	pushImage(destinationDockerRegistry, imageName, imageTag, destinationNamespace)
}

func loginToDockerRegistry(dockerRegistry string, username string, password string) {
	cmd := exec.Command("docker", "login", dockerRegistry, "-u", username, "-p", password)
	output, err := cmd.Output()

	if err != nil {
		fmt.Println("Error logging in to docker registry")
		os.Exit(1)
	}

	fmt.Println(string(output))
}

func pullImage(dockerRegistry string, imageName string, imageTag string, namespace string) {
	cmd := exec.Command("docker", "pull", dockerRegistry+"/"+namespace+"/"+imageName+":"+imageTag)
	output, err := cmd.Output()

	if err != nil {
		fmt.Println("Error pulling image")
		os.Exit(1)
	}

	fmt.Println(string(output))
}

func tagImage(sourceDockerRegistry string, destinationDockerRegistry string, imageName string, imageTag string, sourceNamespace string, destinationNamespace string) {
	cmd := exec.Command("docker", "tag", sourceDockerRegistry+"/"+sourceNamespace+"/"+imageName+":"+imageTag, destinationDockerRegistry+"/"+destinationNamespace+"/"+imageName+":"+imageTag)
	output, err := cmd.Output()

	if err != nil {
		fmt.Println("Error tagging image")
		os.Exit(1)
	}

	fmt.Println(string(output))
}

func pushImage(dockerRegistry string, imageName string, imageTag string, namespace string) {
	cmd := exec.Command("docker", "push", dockerRegistry+"/"+namespace+"/"+imageName+":"+imageTag)
	output, err := cmd.Output()

	if err != nil {
		fmt.Println("Error pushing image")
		os.Exit(1)
	}

	fmt.Println(string(output))
}
