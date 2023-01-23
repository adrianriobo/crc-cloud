// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package ec2

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// `ec2.Eip` provides details about a specific Elastic IP.
//
// ## Example Usage
// ### Search By Allocation ID (VPC only)
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-aws/sdk/v5/go/aws/ec2"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := ec2.GetElasticIp(ctx, &ec2.GetElasticIpArgs{
//				Id: pulumi.StringRef("eipalloc-12345678"),
//			}, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
// ### Search By Filters (EC2-Classic or VPC)
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-aws/sdk/v5/go/aws/ec2"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := ec2.GetElasticIp(ctx, &ec2.GetElasticIpArgs{
//				Filters: []ec2.GetElasticIpFilter{
//					{
//						Name: "tag:Name",
//						Values: []string{
//							"exampleNameTagValue",
//						},
//					},
//				},
//			}, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
// ### Search By Public IP (EC2-Classic or VPC)
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-aws/sdk/v5/go/aws/ec2"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := ec2.GetElasticIp(ctx, &ec2.GetElasticIpArgs{
//				PublicIp: pulumi.StringRef("1.2.3.4"),
//			}, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
// ### Search By Tags (EC2-Classic or VPC)
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-aws/sdk/v5/go/aws/ec2"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := ec2.GetElasticIp(ctx, &ec2.GetElasticIpArgs{
//				Tags: map[string]interface{}{
//					"Name": "exampleNameTagValue",
//				},
//			}, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
func GetElasticIp(ctx *pulumi.Context, args *GetElasticIpArgs, opts ...pulumi.InvokeOption) (*GetElasticIpResult, error) {
	var rv GetElasticIpResult
	err := ctx.Invoke("aws:ec2/getElasticIp:getElasticIp", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getElasticIp.
type GetElasticIpArgs struct {
	// One or more name/value pairs to use as filters. There are several valid keys, for a full reference, check out the [EC2 API Reference](https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_DescribeAddresses.html).
	Filters []GetElasticIpFilter `pulumi:"filters"`
	// Allocation ID of the specific VPC EIP to retrieve. If a classic EIP is required, do NOT set `id`, only set `publicIp`
	Id *string `pulumi:"id"`
	// Public IP of the specific EIP to retrieve.
	PublicIp *string `pulumi:"publicIp"`
	// Map of tags, each pair of which must exactly match a pair on the desired Elastic IP
	Tags map[string]string `pulumi:"tags"`
}

// A collection of values returned by getElasticIp.
type GetElasticIpResult struct {
	// ID representing the association of the address with an instance in a VPC.
	AssociationId string `pulumi:"associationId"`
	// Carrier IP address.
	CarrierIp string `pulumi:"carrierIp"`
	// Customer Owned IP.
	CustomerOwnedIp string `pulumi:"customerOwnedIp"`
	// The ID of a Customer Owned IP Pool. For more on customer owned IP addressed check out [Customer-owned IP addresses guide](https://docs.aws.amazon.com/outposts/latest/userguide/outposts-networking-components.html#ip-addressing)
	CustomerOwnedIpv4Pool string `pulumi:"customerOwnedIpv4Pool"`
	// Whether the address is for use in EC2-Classic (standard) or in a VPC (vpc).
	Domain  string               `pulumi:"domain"`
	Filters []GetElasticIpFilter `pulumi:"filters"`
	// If VPC Elastic IP, the allocation identifier. If EC2-Classic Elastic IP, the public IP address.
	Id string `pulumi:"id"`
	// ID of the instance that the address is associated with (if any).
	InstanceId string `pulumi:"instanceId"`
	// The ID of the network interface.
	NetworkInterfaceId string `pulumi:"networkInterfaceId"`
	// The ID of the AWS account that owns the network interface.
	NetworkInterfaceOwnerId string `pulumi:"networkInterfaceOwnerId"`
	// Private DNS associated with the Elastic IP address.
	PrivateDns string `pulumi:"privateDns"`
	// Private IP address associated with the Elastic IP address.
	PrivateIp string `pulumi:"privateIp"`
	// Public DNS associated with the Elastic IP address.
	PublicDns string `pulumi:"publicDns"`
	// Public IP address of Elastic IP.
	PublicIp string `pulumi:"publicIp"`
	// ID of an address pool.
	PublicIpv4Pool string `pulumi:"publicIpv4Pool"`
	// Key-value map of tags associated with Elastic IP.
	Tags map[string]string `pulumi:"tags"`
}

func GetElasticIpOutput(ctx *pulumi.Context, args GetElasticIpOutputArgs, opts ...pulumi.InvokeOption) GetElasticIpResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetElasticIpResult, error) {
			args := v.(GetElasticIpArgs)
			r, err := GetElasticIp(ctx, &args, opts...)
			var s GetElasticIpResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetElasticIpResultOutput)
}

// A collection of arguments for invoking getElasticIp.
type GetElasticIpOutputArgs struct {
	// One or more name/value pairs to use as filters. There are several valid keys, for a full reference, check out the [EC2 API Reference](https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_DescribeAddresses.html).
	Filters GetElasticIpFilterArrayInput `pulumi:"filters"`
	// Allocation ID of the specific VPC EIP to retrieve. If a classic EIP is required, do NOT set `id`, only set `publicIp`
	Id pulumi.StringPtrInput `pulumi:"id"`
	// Public IP of the specific EIP to retrieve.
	PublicIp pulumi.StringPtrInput `pulumi:"publicIp"`
	// Map of tags, each pair of which must exactly match a pair on the desired Elastic IP
	Tags pulumi.StringMapInput `pulumi:"tags"`
}

func (GetElasticIpOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetElasticIpArgs)(nil)).Elem()
}

// A collection of values returned by getElasticIp.
type GetElasticIpResultOutput struct{ *pulumi.OutputState }

func (GetElasticIpResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetElasticIpResult)(nil)).Elem()
}

func (o GetElasticIpResultOutput) ToGetElasticIpResultOutput() GetElasticIpResultOutput {
	return o
}

func (o GetElasticIpResultOutput) ToGetElasticIpResultOutputWithContext(ctx context.Context) GetElasticIpResultOutput {
	return o
}

// ID representing the association of the address with an instance in a VPC.
func (o GetElasticIpResultOutput) AssociationId() pulumi.StringOutput {
	return o.ApplyT(func(v GetElasticIpResult) string { return v.AssociationId }).(pulumi.StringOutput)
}

// Carrier IP address.
func (o GetElasticIpResultOutput) CarrierIp() pulumi.StringOutput {
	return o.ApplyT(func(v GetElasticIpResult) string { return v.CarrierIp }).(pulumi.StringOutput)
}

// Customer Owned IP.
func (o GetElasticIpResultOutput) CustomerOwnedIp() pulumi.StringOutput {
	return o.ApplyT(func(v GetElasticIpResult) string { return v.CustomerOwnedIp }).(pulumi.StringOutput)
}

// The ID of a Customer Owned IP Pool. For more on customer owned IP addressed check out [Customer-owned IP addresses guide](https://docs.aws.amazon.com/outposts/latest/userguide/outposts-networking-components.html#ip-addressing)
func (o GetElasticIpResultOutput) CustomerOwnedIpv4Pool() pulumi.StringOutput {
	return o.ApplyT(func(v GetElasticIpResult) string { return v.CustomerOwnedIpv4Pool }).(pulumi.StringOutput)
}

// Whether the address is for use in EC2-Classic (standard) or in a VPC (vpc).
func (o GetElasticIpResultOutput) Domain() pulumi.StringOutput {
	return o.ApplyT(func(v GetElasticIpResult) string { return v.Domain }).(pulumi.StringOutput)
}

func (o GetElasticIpResultOutput) Filters() GetElasticIpFilterArrayOutput {
	return o.ApplyT(func(v GetElasticIpResult) []GetElasticIpFilter { return v.Filters }).(GetElasticIpFilterArrayOutput)
}

// If VPC Elastic IP, the allocation identifier. If EC2-Classic Elastic IP, the public IP address.
func (o GetElasticIpResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetElasticIpResult) string { return v.Id }).(pulumi.StringOutput)
}

// ID of the instance that the address is associated with (if any).
func (o GetElasticIpResultOutput) InstanceId() pulumi.StringOutput {
	return o.ApplyT(func(v GetElasticIpResult) string { return v.InstanceId }).(pulumi.StringOutput)
}

// The ID of the network interface.
func (o GetElasticIpResultOutput) NetworkInterfaceId() pulumi.StringOutput {
	return o.ApplyT(func(v GetElasticIpResult) string { return v.NetworkInterfaceId }).(pulumi.StringOutput)
}

// The ID of the AWS account that owns the network interface.
func (o GetElasticIpResultOutput) NetworkInterfaceOwnerId() pulumi.StringOutput {
	return o.ApplyT(func(v GetElasticIpResult) string { return v.NetworkInterfaceOwnerId }).(pulumi.StringOutput)
}

// Private DNS associated with the Elastic IP address.
func (o GetElasticIpResultOutput) PrivateDns() pulumi.StringOutput {
	return o.ApplyT(func(v GetElasticIpResult) string { return v.PrivateDns }).(pulumi.StringOutput)
}

// Private IP address associated with the Elastic IP address.
func (o GetElasticIpResultOutput) PrivateIp() pulumi.StringOutput {
	return o.ApplyT(func(v GetElasticIpResult) string { return v.PrivateIp }).(pulumi.StringOutput)
}

// Public DNS associated with the Elastic IP address.
func (o GetElasticIpResultOutput) PublicDns() pulumi.StringOutput {
	return o.ApplyT(func(v GetElasticIpResult) string { return v.PublicDns }).(pulumi.StringOutput)
}

// Public IP address of Elastic IP.
func (o GetElasticIpResultOutput) PublicIp() pulumi.StringOutput {
	return o.ApplyT(func(v GetElasticIpResult) string { return v.PublicIp }).(pulumi.StringOutput)
}

// ID of an address pool.
func (o GetElasticIpResultOutput) PublicIpv4Pool() pulumi.StringOutput {
	return o.ApplyT(func(v GetElasticIpResult) string { return v.PublicIpv4Pool }).(pulumi.StringOutput)
}

// Key-value map of tags associated with Elastic IP.
func (o GetElasticIpResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v GetElasticIpResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func init() {
	pulumi.RegisterOutputType(GetElasticIpResultOutput{})
}
