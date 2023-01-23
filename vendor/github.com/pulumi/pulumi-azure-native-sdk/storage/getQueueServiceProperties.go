// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package storage

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// The properties of a storage account’s Queue service.
// API Version: 2021-02-01.
func LookupQueueServiceProperties(ctx *pulumi.Context, args *LookupQueueServicePropertiesArgs, opts ...pulumi.InvokeOption) (*LookupQueueServicePropertiesResult, error) {
	var rv LookupQueueServicePropertiesResult
	err := ctx.Invoke("azure-native:storage:getQueueServiceProperties", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupQueueServicePropertiesArgs struct {
	// The name of the storage account within the specified resource group. Storage account names must be between 3 and 24 characters in length and use numbers and lower-case letters only.
	AccountName string `pulumi:"accountName"`
	// The name of the Queue Service within the specified storage account. Queue Service Name must be 'default'
	QueueServiceName string `pulumi:"queueServiceName"`
	// The name of the resource group within the user's subscription. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// The properties of a storage account’s Queue service.
type LookupQueueServicePropertiesResult struct {
	// Specifies CORS rules for the Queue service. You can include up to five CorsRule elements in the request. If no CorsRule elements are included in the request body, all CORS rules will be deleted, and CORS will be disabled for the Queue service.
	Cors *CorsRulesResponse `pulumi:"cors"`
	// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	Id string `pulumi:"id"`
	// The name of the resource
	Name string `pulumi:"name"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type string `pulumi:"type"`
}

func LookupQueueServicePropertiesOutput(ctx *pulumi.Context, args LookupQueueServicePropertiesOutputArgs, opts ...pulumi.InvokeOption) LookupQueueServicePropertiesResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupQueueServicePropertiesResult, error) {
			args := v.(LookupQueueServicePropertiesArgs)
			r, err := LookupQueueServiceProperties(ctx, &args, opts...)
			var s LookupQueueServicePropertiesResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupQueueServicePropertiesResultOutput)
}

type LookupQueueServicePropertiesOutputArgs struct {
	// The name of the storage account within the specified resource group. Storage account names must be between 3 and 24 characters in length and use numbers and lower-case letters only.
	AccountName pulumi.StringInput `pulumi:"accountName"`
	// The name of the Queue Service within the specified storage account. Queue Service Name must be 'default'
	QueueServiceName pulumi.StringInput `pulumi:"queueServiceName"`
	// The name of the resource group within the user's subscription. The name is case insensitive.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupQueueServicePropertiesOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupQueueServicePropertiesArgs)(nil)).Elem()
}

// The properties of a storage account’s Queue service.
type LookupQueueServicePropertiesResultOutput struct{ *pulumi.OutputState }

func (LookupQueueServicePropertiesResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupQueueServicePropertiesResult)(nil)).Elem()
}

func (o LookupQueueServicePropertiesResultOutput) ToLookupQueueServicePropertiesResultOutput() LookupQueueServicePropertiesResultOutput {
	return o
}

func (o LookupQueueServicePropertiesResultOutput) ToLookupQueueServicePropertiesResultOutputWithContext(ctx context.Context) LookupQueueServicePropertiesResultOutput {
	return o
}

// Specifies CORS rules for the Queue service. You can include up to five CorsRule elements in the request. If no CorsRule elements are included in the request body, all CORS rules will be deleted, and CORS will be disabled for the Queue service.
func (o LookupQueueServicePropertiesResultOutput) Cors() CorsRulesResponsePtrOutput {
	return o.ApplyT(func(v LookupQueueServicePropertiesResult) *CorsRulesResponse { return v.Cors }).(CorsRulesResponsePtrOutput)
}

// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
func (o LookupQueueServicePropertiesResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupQueueServicePropertiesResult) string { return v.Id }).(pulumi.StringOutput)
}

// The name of the resource
func (o LookupQueueServicePropertiesResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupQueueServicePropertiesResult) string { return v.Name }).(pulumi.StringOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o LookupQueueServicePropertiesResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupQueueServicePropertiesResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupQueueServicePropertiesResultOutput{})
}
