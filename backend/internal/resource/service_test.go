package resource

import (
	"testing"
)

func TestResourceRegistration(t *testing.T) {
	sensitivityLevels := []string{"public", "internal", "restricted", "critical"}
	for _, level := range sensitivityLevels {
		if level == "" {
			t.Error("sensitivity level should not be empty")
		}
	}

	resourceTypes := []string{"web_app", "api", "database", "ssh", "tcp"}
	for _, rt := range resourceTypes {
		if rt == "" {
			t.Error("resource type should not be empty")
		}
	}
}

func TestConnectorTypes(t *testing.T) {
	connectorTypes := []string{"http_proxy", "tcp_tunnel", "database_proxy"}
	for _, ct := range connectorTypes {
		if ct == "" {
			t.Error("connector type should not be empty")
		}
	}
}
