package provider

import (
	"github.com/crc/crc-cloud/pkg/util/plugin"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

const (
	OutputKey      string = "private-key"
	OutputHost     string = "host"
	OutputUsername string = "username"
	OutputPassword string = "password"

	OutputBootKey string = "bootkey"
	OutputImageID string = "image-id"
)

type Provider interface {
	// Plugin information, required to dynamically install the plugin
	// for the specific provider
	GetPlugin() *plugin.Plugin

	// Manage all the image import process for the specific provider
	ImportImageRunFunc(bundleDownloadURL, shasumfileDownloadURL string) (pulumi.RunFunc, error)

	// Set of params tied to provider to customize the create operation
	CreateParams() map[string]string
	// Subset of create params which are mandatory
	CreateParamsMandatory() []string
	// Creates all resources for the specific provider required on the create operation
	CreateRunFunc(bootingPrivateKeyFilePath, ocpPullSecretFilePath string,
		args map[string]string) (pulumi.RunFunc, error)
}
