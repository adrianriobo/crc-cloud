package manager

import (
	"context"
	"fmt"
	"os"
	"path"

	crcContext "github.com/crc/crc-cloud/pkg/manager/context"
	"github.com/crc/crc-cloud/pkg/util/plugin"
	"github.com/pulumi/pulumi/sdk/v3/go/auto"
	"github.com/pulumi/pulumi/sdk/v3/go/auto/optdestroy"
	"github.com/pulumi/pulumi/sdk/v3/go/auto/optup"
	"github.com/pulumi/pulumi/sdk/v3/go/common/tokens"
	"github.com/pulumi/pulumi/sdk/v3/go/common/util/logging"
	"github.com/pulumi/pulumi/sdk/v3/go/common/workspace"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type stack struct {
	name           string
	backedURL      string
	deployFunc     pulumi.RunFunc
	providerPlugin plugin.Plugin
}

func newStack(name, backedURL string,
	deployFunc pulumi.RunFunc,
	providerPlugin plugin.Plugin) stack {
	return stack{
		name:           name,
		backedURL:      backedURL,
		deployFunc:     deployFunc,
		providerPlugin: providerPlugin,
	}
}

func validateParams(args map[string]string, required []string) error {
	var requiredMissing []string
	for _, r := range required {
		_, ok := args[r]
		if !ok {
			requiredMissing = append(requiredMissing, r)
		}
	}

	if len(requiredMissing) > 0 {
		return fmt.Errorf("required fields missing: %v", requiredMissing)
	}
	return nil
}

func upStack(s stack) (auto.UpResult, error) {
	ctx := context.Background()
	objectStack := getStack(ctx, s)
	stdoutStreamer := optup.ProgressStreams(os.Stdout)
	return objectStack.Up(ctx, stdoutStreamer)
}

func destroyStack(s stack) (err error) {
	ctx := context.Background()
	objectStack := getStack(ctx, s)
	stdoutStreamer := optdestroy.ProgressStreams(os.Stdout)
	if _, err = objectStack.Destroy(ctx, stdoutStreamer); err != nil {
		return
	}
	err = objectStack.Workspace().RemoveStack(ctx, s.name)
	return
}

// this function gets our stack ready for update/destroy by prepping the workspace, init/selecting the stack
// and doing a refresh to make sure state and cloud resources are in sync
func getStack(ctx context.Context, s stack) auto.Stack {
	// create or select a stack with an inline Pulumi program
	ps, err := auto.UpsertStackInlineSource(ctx, s.name,
		crcContext.GetName(), s.deployFunc, getOpts(s)...)
	if err != nil {
		logging.Errorf("%v", err)
		os.Exit(1)
	}
	if err = s.providerPlugin.Install(ctx, &ps); err != nil {
		logging.Errorf("%v", err)
		os.Exit(1)
	}
	return ps
}

func getOpts(s stack) []auto.LocalWorkspaceOption {
	return []auto.LocalWorkspaceOption{
		auto.Project(workspace.Project{
			Name:    tokens.PackageName(crcContext.GetName()),
			Runtime: workspace.NewProjectRuntimeInfo("go", nil),
			Backend: &workspace.ProjectBackend{
				URL: s.backedURL,
			},
		}),
		auto.WorkDir("."),
		// auto.SecretsProvider("awskms://alias/pulumi-secret-encryption"),
	}
}

func writeOutputs(stackResult auto.UpResult,
	destinationFolder string, results map[string]string) (err error) {
	for k, v := range results {
		if err = writeOutput(stackResult, k, destinationFolder, v); err != nil {
			return err
		}
	}
	return
}

func writeOutput(stackResult auto.UpResult, outputkey,
	destinationFolder, destinationFilename string) error {
	value, ok := stackResult.Outputs[outputkey].Value.(string)
	if ok {
		err := os.WriteFile(path.Join(destinationFolder, destinationFilename), []byte(value), 0600)
		if err != nil {
			return err
		}
	} else {
		return fmt.Errorf("output value %s not found", outputkey)
	}
	return nil
}
