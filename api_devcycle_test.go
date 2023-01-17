package devcycle

import (
	"context"
	"fmt"
	"os"
	"testing"
	"time"

	"github.com/jarcoal/httpmock"
)

func TestDVCClientService_AllFeatures_Local(t *testing.T) {
	auth := context.WithValue(context.Background(), ContextAPIKey, APIKey{
		Key: "dvc_server_token_hash",
	})
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	httpConfigMock(200)
	c, err := NewDVCClient("dvc_server_token_hash", &DVCOptions{})

	features, err := c.DevCycleApi.AllFeatures(auth,
		DVCUser{UserId: "j_test", Platform: "golang-testing", SdkType: "server", PlatformVersion: "testing", DeviceModel: "testing", SdkVersion: "testing"})
	if err != nil {
		t.Fatal(err)
		return
	}

	fmt.Println(features)

}
func TestDVCClientService_AllVariablesLocal(t *testing.T) {
	auth := context.WithValue(context.Background(), ContextAPIKey, APIKey{
		Key: "dvc_server_token_hash",
	})
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	httpConfigMock(200)
	c, err := NewDVCClient("dvc_server_token_hash", &DVCOptions{})

	variables, err := c.DevCycleApi.AllVariables(auth,
		DVCUser{UserId: "j_test", Platform: "golang-testing", SdkType: "server", PlatformVersion: "testing", DeviceModel: "testing", SdkVersion: "testing"})
	if err != nil {
		t.Fatal(err)
		return
	}

	fmt.Println(variables)
}

func TestDVCClientService_VariableCloud(t *testing.T) {
	auth := context.WithValue(context.Background(), ContextAPIKey, APIKey{
		Key: "dvc_server_token_hash",
	})
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	httpBucketingAPIMock()
	c, err := NewDVCClient("dvc_server_token_hash", &DVCOptions{EnableCloudBucketing: true, ConfigPollingIntervalMS: 10 * time.Second})

	variable, err := c.DevCycleApi.Variable(auth,
		DVCUser{UserId: "j_test", Platform: "golang-testing", SdkType: "server", PlatformVersion: "testing", DeviceModel: "testing", SdkVersion: "testing"},
		"test", true)
	if err != nil {
		t.Fatal(err)
		return
	}

	fmt.Println(variable)
}

func TestDVCClientService_VariableLocal(t *testing.T) {
	auth := context.WithValue(context.Background(), ContextAPIKey, APIKey{
		Key: "dvc_server_token_hash",
	})
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	httpConfigMock(200)

	c, err := NewDVCClient("dvc_server_token_hash", &DVCOptions{})

	variable, err := c.DevCycleApi.Variable(auth,
		DVCUser{UserId: "j_test", Platform: "golang-testing", SdkType: "server", PlatformVersion: "testing", DeviceModel: "testing", SdkVersion: "testing"},
		"test", true)
	if err != nil {
		t.Fatal(err)
		return
	}

	fmt.Println(variable)
}

func TestDVCClientService_VariableLocal_403(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	httpConfigMock(403)

	_, err := NewDVCClient("dvc_server_token_hash", &DVCOptions{})
	if err == nil {
		t.Fatal("Expected error from configmanager")
	}
}

func TestDVCClientService_TrackLocal_QueueEvent(t *testing.T) {
	auth := context.WithValue(context.Background(), ContextAPIKey, APIKey{
		Key: test_environmentKey,
	})
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	httpConfigMock(200)
	dvcOptions := DVCOptions{ConfigPollingIntervalMS: 10 * time.Second}

	c, err := NewDVCClient(test_environmentKey, &dvcOptions)

	track, err := c.DevCycleApi.Track(auth, DVCUser{UserId: "j_test", Platform: "golang-testing", SdkType: "server", PlatformVersion: "testing", DeviceModel: "testing", SdkVersion: "testing"}, DVCEvent{
		Target:      "customEvent",
		Value:       0,
		FeatureVars: nil,
		MetaData:    nil,
	})
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(track)
}

func TestProduction_Local(t *testing.T) {
	environmentKey := os.Getenv("DVC_SERVER_KEY")
	user := DVCUser{UserId: "test"}
	auth := context.WithValue(context.Background(), ContextAPIKey, APIKey{
		Key: environmentKey,
	})
	if environmentKey == "" {
		t.Skip("DVC_SERVER_KEY not set. Not using production tests.")
	}
	dvcOptions := DVCOptions{
		EnableEdgeDB:                 false,
		EnableCloudBucketing:         false,
		EventFlushIntervalMS:         0,
		ConfigPollingIntervalMS:      10 * time.Second,
		RequestTimeout:               10 * time.Second,
		DisableAutomaticEventLogging: false,
		DisableCustomEventLogging:    false,
	}
	client, err := NewDVCClient(environmentKey, &dvcOptions)
	if err != nil {
		t.Fatal(err)
	}

	variables, err := client.DevCycleApi.AllVariables(auth, user)
	if err != nil {
		t.Fatal(err)
	}
	if len(variables) == 0 {
		t.Fatal("No variables returned")
	}
}
