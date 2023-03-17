package manager

import (
	"github.com/crc/crc-cloud/pkg/manager/context"
	"github.com/crc/crc-cloud/pkg/manager/provider"
	"github.com/pulumi/pulumi/sdk/v3/go/auto"
	"golang.org/x/exp/maps"
)

const (
	stackCreate string = "crcCloud-Create"
)

func CreateParams() (params map[string]string) {
	params = map[string]string{}
	for _, p := range GetSupportedProviders() {
		maps.Copy(params, p.CreateParams())
	}
	return
}

func Create(projectName, backerURL, outputFoler string,
	sp Provider, providerArgs map[string]string,
	ocpPullSecretFilePath, bootKeyFilePath string,
	tags map[string]string) error {

	// Initialize context
	context.Init(projectName, tags)

	// this will return a provider which implements the api.Provider interface
	p, err := GetProvider(sp)
	if err != nil {
		return err
	}
	// TODO think best option to pass params to provider
	// may serialize all params and let provider validate and pick the required
	// as a provider the params are specs and manager requires to know about them?
	err = validateParams(providerArgs, p.CreateParamsMandatory())
	if err != nil {
		return err
	}

	createFunc, err :=
		p.CreateRunFunc(bootKeyFilePath, ocpPullSecretFilePath,
			providerArgs)
	if err != nil {
		return err
	}

	stackResult, err := upStack(
		newStack(stackCreate, backerURL,
			createFunc, *p.GetPlugin()))
	if err != nil {
		return err
	}
	return manageCreateResults(stackResult, outputFoler)
}

func manageCreateResults(stackResult auto.UpResult, destinationFolder string) error {
	if err := writeOutputs(stackResult, destinationFolder, map[string]string{
		provider.OutputKey:      "id_rsa",
		provider.OutputHost:     "host",
		provider.OutputUsername: "username",
		provider.OutputPassword: "password",
	}); err != nil {
		return err
	}
	return nil
}
