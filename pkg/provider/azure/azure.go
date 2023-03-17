package azure

import (
	"github.com/crc/crc-cloud/pkg/util/plugin"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

const (
	azureNativePluginName    = "azure-native"
	azureNativePluginVersion = "v1.92.0"
)

type Provider struct{}

func GetProvider() *Provider {
	return &Provider{}
}

func (a *Provider) GetPlugin() *plugin.Plugin {
	return &plugin.Plugin{
		Name:    azureNativePluginName,
		Version: azureNativePluginVersion}
}

func (a *Provider) ImportImageRunFunc(projectName, bundleDownloadURL, shasumfileDownloadURL string) (pulumi.RunFunc, error) {
	r, err := fillImportRequest(projectName, bundleDownloadURL, shasumfileDownloadURL)
	if err != nil {
		return nil, err
	}
	return (pulumi.RunFunc)(r.runFunc), nil
}
