package composeit

import (
	"testing"

	compose "github.com/libkermit/compose/testing"
)

func TestServicesProject(t *testing.T) {
	project := compose.CreateProject(t, "services", "../assets/services.yml")
	project.Start(t, "hello")

	container := project.Container(t, "hello")
	if container.Name != "/services_hello_1" {
		t.Fatalf("expected name '/services_hello_1', got %s", container.Name)
	}

	//"No container found for '%s' service
	project.NoContainer(t, "other")

	project.Stop(t, "hello")
}
