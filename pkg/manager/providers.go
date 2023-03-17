package manager

import (
	"fmt"

	"github.com/crc/crc-cloud/pkg/manager/provider"
	"github.com/crc/crc-cloud/pkg/provider/aws"
)

type Provider string

const (
	AWS Provider = "aws"
	GCP Provider = "gcp"
	AZ  Provider = "azure"
)

func GetProvider(p Provider) (provider.Provider, error) {
	switch p {
	case AWS, AZ:
		return aws.GetProvider(), nil
	}
	return nil, fmt.Errorf("%s: provider not supported", p)
}

func GetSupportedProviders() (sp []provider.Provider) {
	sp = append(sp, aws.GetProvider())
	return
}
