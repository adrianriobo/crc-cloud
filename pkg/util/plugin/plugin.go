package plugin

import (
	"context"

	"github.com/pulumi/pulumi/sdk/v3/go/auto"
)

// Struct with information about
// plugin to manage provider resources
type Plugin struct {
	Name    string
	Version string
}

func (p *Plugin) Install(ctx context.Context, s *auto.Stack) (err error) {
	if err = s.Workspace().InstallPlugin(ctx, p.Name, p.Version); err != nil {
		return
	}
	_, err = s.Refresh(ctx)
	return
}
