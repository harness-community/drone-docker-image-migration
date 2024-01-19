package main

import (
	"drone/plugin/image-migration/env"
	"os"
	"testing"
)

func TestMain_EnvVarsNotSet(t *testing.T) {
	// Save current environment variables
	sourceRegistry := os.Getenv("PLUGIN_SOURCE_REGISTRY")
	sourceUsername := os.Getenv("PLUGIN_SOURCE_USERNAME")
	sourcePassword := os.Getenv("PLUGIN_SOURCE_PASSWORD")
	sourceNamespace := os.Getenv("PLUGIN_SOURCE_NAMESPACE")

	destinationRegistry := os.Getenv("PLUGIN_DESTINATION_REGISTRY")
	destinationUsername := os.Getenv("PLUGIN_DESTINATION_USERNAME")
	destinationNamespace := os.Getenv("PLUGIN_DESTINATION_NAMESPACE")

	imageName := os.Getenv("PLUGIN_IMAGE_NAME")
	imageTag := os.Getenv("PLUGIN_IMAGE_TAG")

	// Clear environment variables
	os.Setenv("PLUGIN_SOURCE_REGISTRY", "")
	os.Setenv("PLUGIN_SOURCE_USERNAME", "")
	os.Setenv("PLUGIN_SOURCE_PASSWORD", "")
	os.Setenv("PLUGIN_SOURCE_NAMESPACE", "")

	os.Setenv("PLUGIN_DESTINATION_REGISTRY", "")
	os.Setenv("PLUGIN_DESTINATION_USERNAME", "")
	os.Setenv("PLUGIN_DESTINATION_NAMESPACE", "")

	os.Setenv("PLUGIN_IMAGE_NAME", "")
	os.Setenv("PLUGIN_IMAGE_TAG", "")

	defer func() {
		// Restore original environment variables
		os.Setenv("PLUGIN_SOURCE_REGISTRY", sourceRegistry)
		os.Setenv("PLUGIN_SOURCE_USERNAME", sourceUsername)
		os.Setenv("PLUGIN_SOURCE_PASSWORD", sourcePassword)
		os.Setenv("PLUGIN_SOURCE_NAMESPACE", sourceNamespace)

		os.Setenv("PLUGIN_DESTINATION_REGISTRY", destinationRegistry)
		os.Setenv("PLUGIN_DESTINATION_USERNAME", destinationUsername)
		os.Setenv("PLUGIN_DESTINATION_NAMESPACE", destinationNamespace)

		os.Setenv("PLUGIN_IMAGE_NAME", imageName)
		os.Setenv("PLUGIN_IMAGE_TAG", imageTag)
	}()

	err := env.VerifyEnvVars()
	if err == nil {
		t.Error("Expected error, but got nil")
	}

	expectedErrorMessage := "missing source variables"
	if got := err.Error(); got != expectedErrorMessage {
		t.Errorf("Expected error message %q, but got %q", expectedErrorMessage, got)
	}
}

func TestMain_EnvVarsSet(t *testing.T) {
	// Save current environment variables
	sourceRegistry := os.Getenv("PLUGIN_SOURCE_REGISTRY")
	sourceUsername := os.Getenv("PLUGIN_SOURCE_USERNAME")
	sourcePassword := os.Getenv("PLUGIN_SOURCE_PASSWORD")
	sourceNamespace := os.Getenv("PLUGIN_SOURCE_NAMESPACE")

	destinationRegistry := os.Getenv("PLUGIN_DESTINATION_REGISTRY")
	destinationUsername := os.Getenv("PLUGIN_DESTINATION_USERNAME")
	destinationNamespace := os.Getenv("PLUGIN_DESTINATION_NAMESPACE")

	imageName := os.Getenv("PLUGIN_IMAGE_NAME")
	imageTag := os.Getenv("PLUGIN_IMAGE_TAG")

	// Set environment variables
	os.Setenv("PLUGIN_SOURCE_REGISTRY", "docker.io")
	os.Setenv("PLUGIN_SOURCE_USERNAME", "test")
	os.Setenv("PLUGIN_SOURCE_PASSWORD", "pass")
	os.Setenv("PLUGIN_SOURCE_NAMESPACE", "test")

	os.Setenv("PLUGIN_DESTINATION_REGISTRY", "aws")
	os.Setenv("PLUGIN_DESTINATION_USERNAME", "test")
	os.Setenv("PLUGIN_DESTINATION_NAMESPACE", "test")

	os.Setenv("PLUGIN_IMAGE_NAME", "image")
	os.Setenv("PLUGIN_IMAGE_TAG", "latest")

	defer func() {
		// Restore original environment variables
		os.Setenv("PLUGIN_SOURCE_REGISTRY", sourceRegistry)
		os.Setenv("PLUGIN_SOURCE_USERNAME", sourceUsername)
		os.Setenv("PLUGIN_SOURCE_PASSWORD", sourcePassword)
		os.Setenv("PLUGIN_SOURCE_NAMESPACE", sourceNamespace)

		os.Setenv("PLUGIN_DESTINATION_REGISTRY", destinationRegistry)
		os.Setenv("PLUGIN_DESTINATION_USERNAME", destinationUsername)
		os.Setenv("PLUGIN_DESTINATION_NAMESPACE", destinationNamespace)

		os.Setenv("PLUGIN_IMAGE_NAME", imageName)
		os.Setenv("PLUGIN_IMAGE_TAG", imageTag)
	}()

	err := env.VerifyEnvVars()
	if err != nil {
		t.Errorf("Expected nil, but got %v", err)
	}
}
