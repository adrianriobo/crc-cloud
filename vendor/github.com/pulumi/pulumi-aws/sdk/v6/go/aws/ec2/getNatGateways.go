// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package ec2

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// This resource can be useful for getting back a list of NAT gateway ids to be referenced elsewhere.
func GetNatGateways(ctx *pulumi.Context, args *GetNatGatewaysArgs, opts ...pulumi.InvokeOption) (*GetNatGatewaysResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetNatGatewaysResult
	err := ctx.Invoke("aws:ec2/getNatGateways:getNatGateways", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getNatGateways.
type GetNatGatewaysArgs struct {
	// Custom filter block as described below.
	Filters []GetNatGatewaysFilter `pulumi:"filters"`
	// Map of tags, each pair of which must exactly match
	// a pair on the desired NAT Gateways.
	//
	// More complex filters can be expressed using one or more `filter` sub-blocks,
	// which take the following arguments:
	Tags map[string]string `pulumi:"tags"`
	// VPC ID that you want to filter from.
	VpcId *string `pulumi:"vpcId"`
}

// A collection of values returned by getNatGateways.
type GetNatGatewaysResult struct {
	Filters []GetNatGatewaysFilter `pulumi:"filters"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// List of all the NAT gateway ids found.
	Ids   []string          `pulumi:"ids"`
	Tags  map[string]string `pulumi:"tags"`
	VpcId *string           `pulumi:"vpcId"`
}

func GetNatGatewaysOutput(ctx *pulumi.Context, args GetNatGatewaysOutputArgs, opts ...pulumi.InvokeOption) GetNatGatewaysResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetNatGatewaysResult, error) {
			args := v.(GetNatGatewaysArgs)
			r, err := GetNatGateways(ctx, &args, opts...)
			var s GetNatGatewaysResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetNatGatewaysResultOutput)
}

// A collection of arguments for invoking getNatGateways.
type GetNatGatewaysOutputArgs struct {
	// Custom filter block as described below.
	Filters GetNatGatewaysFilterArrayInput `pulumi:"filters"`
	// Map of tags, each pair of which must exactly match
	// a pair on the desired NAT Gateways.
	//
	// More complex filters can be expressed using one or more `filter` sub-blocks,
	// which take the following arguments:
	Tags pulumi.StringMapInput `pulumi:"tags"`
	// VPC ID that you want to filter from.
	VpcId pulumi.StringPtrInput `pulumi:"vpcId"`
}

func (GetNatGatewaysOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetNatGatewaysArgs)(nil)).Elem()
}

// A collection of values returned by getNatGateways.
type GetNatGatewaysResultOutput struct{ *pulumi.OutputState }

func (GetNatGatewaysResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetNatGatewaysResult)(nil)).Elem()
}

func (o GetNatGatewaysResultOutput) ToGetNatGatewaysResultOutput() GetNatGatewaysResultOutput {
	return o
}

func (o GetNatGatewaysResultOutput) ToGetNatGatewaysResultOutputWithContext(ctx context.Context) GetNatGatewaysResultOutput {
	return o
}

func (o GetNatGatewaysResultOutput) ToOutput(ctx context.Context) pulumix.Output[GetNatGatewaysResult] {
	return pulumix.Output[GetNatGatewaysResult]{
		OutputState: o.OutputState,
	}
}

func (o GetNatGatewaysResultOutput) Filters() GetNatGatewaysFilterArrayOutput {
	return o.ApplyT(func(v GetNatGatewaysResult) []GetNatGatewaysFilter { return v.Filters }).(GetNatGatewaysFilterArrayOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetNatGatewaysResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetNatGatewaysResult) string { return v.Id }).(pulumi.StringOutput)
}

// List of all the NAT gateway ids found.
func (o GetNatGatewaysResultOutput) Ids() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetNatGatewaysResult) []string { return v.Ids }).(pulumi.StringArrayOutput)
}

func (o GetNatGatewaysResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v GetNatGatewaysResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o GetNatGatewaysResultOutput) VpcId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetNatGatewaysResult) *string { return v.VpcId }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(GetNatGatewaysResultOutput{})
}
