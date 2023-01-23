// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package resources

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Deployment information.
// API Version: 2021-01-01.
func LookupDeploymentAtSubscriptionScope(ctx *pulumi.Context, args *LookupDeploymentAtSubscriptionScopeArgs, opts ...pulumi.InvokeOption) (*LookupDeploymentAtSubscriptionScopeResult, error) {
	var rv LookupDeploymentAtSubscriptionScopeResult
	err := ctx.Invoke("azure-native:resources:getDeploymentAtSubscriptionScope", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupDeploymentAtSubscriptionScopeArgs struct {
	// The name of the deployment.
	DeploymentName string `pulumi:"deploymentName"`
}

// Deployment information.
type LookupDeploymentAtSubscriptionScopeResult struct {
	// The ID of the deployment.
	Id string `pulumi:"id"`
	// the location of the deployment.
	Location *string `pulumi:"location"`
	// The name of the deployment.
	Name string `pulumi:"name"`
	// Deployment properties.
	Properties DeploymentPropertiesExtendedResponse `pulumi:"properties"`
	// Deployment tags
	Tags map[string]string `pulumi:"tags"`
	// The type of the deployment.
	Type string `pulumi:"type"`
}

func LookupDeploymentAtSubscriptionScopeOutput(ctx *pulumi.Context, args LookupDeploymentAtSubscriptionScopeOutputArgs, opts ...pulumi.InvokeOption) LookupDeploymentAtSubscriptionScopeResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupDeploymentAtSubscriptionScopeResult, error) {
			args := v.(LookupDeploymentAtSubscriptionScopeArgs)
			r, err := LookupDeploymentAtSubscriptionScope(ctx, &args, opts...)
			var s LookupDeploymentAtSubscriptionScopeResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupDeploymentAtSubscriptionScopeResultOutput)
}

type LookupDeploymentAtSubscriptionScopeOutputArgs struct {
	// The name of the deployment.
	DeploymentName pulumi.StringInput `pulumi:"deploymentName"`
}

func (LookupDeploymentAtSubscriptionScopeOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupDeploymentAtSubscriptionScopeArgs)(nil)).Elem()
}

// Deployment information.
type LookupDeploymentAtSubscriptionScopeResultOutput struct{ *pulumi.OutputState }

func (LookupDeploymentAtSubscriptionScopeResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupDeploymentAtSubscriptionScopeResult)(nil)).Elem()
}

func (o LookupDeploymentAtSubscriptionScopeResultOutput) ToLookupDeploymentAtSubscriptionScopeResultOutput() LookupDeploymentAtSubscriptionScopeResultOutput {
	return o
}

func (o LookupDeploymentAtSubscriptionScopeResultOutput) ToLookupDeploymentAtSubscriptionScopeResultOutputWithContext(ctx context.Context) LookupDeploymentAtSubscriptionScopeResultOutput {
	return o
}

// The ID of the deployment.
func (o LookupDeploymentAtSubscriptionScopeResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDeploymentAtSubscriptionScopeResult) string { return v.Id }).(pulumi.StringOutput)
}

// the location of the deployment.
func (o LookupDeploymentAtSubscriptionScopeResultOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupDeploymentAtSubscriptionScopeResult) *string { return v.Location }).(pulumi.StringPtrOutput)
}

// The name of the deployment.
func (o LookupDeploymentAtSubscriptionScopeResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDeploymentAtSubscriptionScopeResult) string { return v.Name }).(pulumi.StringOutput)
}

// Deployment properties.
func (o LookupDeploymentAtSubscriptionScopeResultOutput) Properties() DeploymentPropertiesExtendedResponseOutput {
	return o.ApplyT(func(v LookupDeploymentAtSubscriptionScopeResult) DeploymentPropertiesExtendedResponse {
		return v.Properties
	}).(DeploymentPropertiesExtendedResponseOutput)
}

// Deployment tags
func (o LookupDeploymentAtSubscriptionScopeResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupDeploymentAtSubscriptionScopeResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

// The type of the deployment.
func (o LookupDeploymentAtSubscriptionScopeResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDeploymentAtSubscriptionScopeResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupDeploymentAtSubscriptionScopeResultOutput{})
}
