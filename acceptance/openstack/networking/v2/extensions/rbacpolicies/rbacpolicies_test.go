// +build acceptance

package rbacpolicies

import (
	"os"
	"testing"

	"github.com/gophercloud/gophercloud/acceptance/clients"
	projects "github.com/gophercloud/gophercloud/acceptance/openstack/identity/v3"
	networking "github.com/gophercloud/gophercloud/acceptance/openstack/networking/v2"
	"github.com/gophercloud/gophercloud/acceptance/tools"
)

func TestRBACPolicyCreate(t *testing.T) {
	username := os.Getenv("OS_USERNAME")
	if username != "admin" {
		t.Skip("must be admin to run this test")
	}

	client, err := clients.NewNetworkV2Client()
	if err != nil {
		t.Fatalf("Unable to create a network client: %v", err)
	}

	// Create a network
	network, err := networking.CreateNetwork(t, client)
	if err != nil {
		t.Fatalf("Unable to create network: %v", err)
	}
	defer networking.DeleteNetwork(t, client, network.ID)

	tools.PrintResource(t, network)

	identityClient, err := clients.NewIdentityV3Client()
	if err != nil {
		t.Fatalf("Unable to obtain an identity client: %v")
	}

	// Create a project/tenant
	project, err := projects.CreateProject(t, identityClient, nil)
	if err != nil {
		t.Fatalf("Unable to create project: %v", err)
	}
	defer projects.DeleteProject(t, identityClient, project.ID)

	tools.PrintResource(t, project)

	// Create a rbac-policy
	rbacPolicy, err := CreateRBACPolicy(t, client, project.ID, network.ID)
	if err != nil {
		t.Fatalf("Unable to create rbac-policy: %v", err)
	}

	tools.PrintResource(t, rbacPolicy)
}
