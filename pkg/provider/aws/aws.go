package aws

import (
	"github.com/crc/crc-cloud/pkg/util/plugin"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type Provider struct{}

func GetProvider() *Provider {
	return &Provider{}
}

func (a *Provider) GetPlugin() *plugin.Plugin {
	return &plugin.Plugin{
		Name:    "aws",
		Version: "v5.27.0"}
}

func (a *Provider) ImportImageRunFunc(bundleDownloadURL, shasumfileDownloadURL string) (pulumi.RunFunc, error) {
	r, err := fillImportRequest(bundleDownloadURL, shasumfileDownloadURL)
	if err != nil {
		return nil, err
	}
	return (pulumi.RunFunc)(r.runFunc), nil
}

func (a *Provider) CreateParams() map[string]string {
	return map[string]string{
		amiID: amiIDDesc,
	}
}

func (a *Provider) CreateParamsMandatory() []string {
	return []string{amiID}
}

func (a *Provider) CreateRunFunc(bootingPrivateKeyFilePath, ocpPullSecretFilePath string,
	args map[string]string) (pulumi.RunFunc, error) {
	r, err := fillCreateRequest(bootingPrivateKeyFilePath, ocpPullSecretFilePath, args)
	if err != nil {
		return nil, err
	}
	return (pulumi.RunFunc)(r.runFunc), nil
}
