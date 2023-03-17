package manager

import (
	"github.com/crc/crc-cloud/pkg/manager/context"
	"github.com/crc/crc-cloud/pkg/manager/provider"
	"github.com/pulumi/pulumi/sdk/v3/go/auto"
)

const (
	stackImportImage string = "crcCloud-ImageImport"
)

func Import(projectName, backerURL, outputFoler string,
	bundleDownloadURL, shasumfileDownloadURL string, sp Provider,
	tags map[string]string) error {

	// Initialize context
	context.Init(projectName, tags)

	// Pick the import function according to the provider
	p, err := GetProvider(sp)
	if err != nil {
		return err
	}
	importRunFunc, err :=
		p.ImportImageRunFunc(bundleDownloadURL, shasumfileDownloadURL)
	if err != nil {
		return err
	}
	// Create a stack based on the import function and create it
	stack := newStack(stackImportImage, backerURL,
		importRunFunc, *p.GetPlugin())

	stackResult, err := upStack(stack)
	if err != nil {
		return err
	}
	err = manageImageImportResults(stackResult, outputFoler)
	if err != nil {
		return nil
	}

	// Current exec create temporary resources to enable the import
	// we delete it as they are only temporary
	return destroyStack(stack)
}

func manageImageImportResults(stackResult auto.UpResult, destinationFolder string) error {
	if err := writeOutputs(stackResult, destinationFolder, map[string]string{
		provider.OutputBootKey: "id_ecdsa",
		provider.OutputImageID: "image-id",
	}); err != nil {
		return err
	}
	return nil
}
