// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package ec2

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Provides a resource to associate additional IPv6 CIDR blocks with a VPC.
//
// The `ec2.VpcIpv6CidrBlockAssociation` resource allows IPv6 CIDR blocks to be added to the VPC.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/ec2"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			testVpc, err := ec2.NewVpc(ctx, "testVpc", &ec2.VpcArgs{
//				CidrBlock: pulumi.String("10.0.0.0/16"),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = ec2.NewVpcIpv6CidrBlockAssociation(ctx, "testVpcIpv6CidrBlockAssociation", &ec2.VpcIpv6CidrBlockAssociationArgs{
//				Ipv6IpamPoolId: pulumi.Any(aws_vpc_ipam_pool.Test.Id),
//				VpcId:          testVpc.ID(),
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
//
// ## Import
//
// Using `pulumi import`, import `aws_vpc_ipv6_cidr_block_association` using the VPC CIDR Association ID. For example:
//
// ```sh
//
//	$ pulumi import aws:ec2/vpcIpv6CidrBlockAssociation:VpcIpv6CidrBlockAssociation example vpc-cidr-assoc-xxxxxxxx
//
// ```
type VpcIpv6CidrBlockAssociation struct {
	pulumi.CustomResourceState

	// The IPv6 CIDR block for the VPC. CIDR can be explicitly set or it can be derived from IPAM using `ipv6NetmaskLength`. This parameter is required if `ipv6NetmaskLength` is not set and he IPAM pool does not have `allocationDefaultNetmask` set.
	Ipv6CidrBlock pulumi.StringOutput `pulumi:"ipv6CidrBlock"`
	// The ID of an IPv6 IPAM pool you want to use for allocating this VPC's CIDR. IPAM is a VPC feature that you can use to automate your IP address management workflows including assigning, tracking, troubleshooting, and auditing IP addresses across AWS Regions and accounts.
	Ipv6IpamPoolId pulumi.StringOutput `pulumi:"ipv6IpamPoolId"`
	// The netmask length of the IPv6 CIDR you want to allocate to this VPC. Requires specifying a `ipv6IpamPoolId`. This parameter is optional if the IPAM pool has `allocationDefaultNetmask` set, otherwise it or `cidrBlock` are required
	Ipv6NetmaskLength pulumi.IntPtrOutput `pulumi:"ipv6NetmaskLength"`
	// The ID of the VPC to make the association with.
	VpcId pulumi.StringOutput `pulumi:"vpcId"`
}

// NewVpcIpv6CidrBlockAssociation registers a new resource with the given unique name, arguments, and options.
func NewVpcIpv6CidrBlockAssociation(ctx *pulumi.Context,
	name string, args *VpcIpv6CidrBlockAssociationArgs, opts ...pulumi.ResourceOption) (*VpcIpv6CidrBlockAssociation, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Ipv6IpamPoolId == nil {
		return nil, errors.New("invalid value for required argument 'Ipv6IpamPoolId'")
	}
	if args.VpcId == nil {
		return nil, errors.New("invalid value for required argument 'VpcId'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource VpcIpv6CidrBlockAssociation
	err := ctx.RegisterResource("aws:ec2/vpcIpv6CidrBlockAssociation:VpcIpv6CidrBlockAssociation", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetVpcIpv6CidrBlockAssociation gets an existing VpcIpv6CidrBlockAssociation resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetVpcIpv6CidrBlockAssociation(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *VpcIpv6CidrBlockAssociationState, opts ...pulumi.ResourceOption) (*VpcIpv6CidrBlockAssociation, error) {
	var resource VpcIpv6CidrBlockAssociation
	err := ctx.ReadResource("aws:ec2/vpcIpv6CidrBlockAssociation:VpcIpv6CidrBlockAssociation", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering VpcIpv6CidrBlockAssociation resources.
type vpcIpv6CidrBlockAssociationState struct {
	// The IPv6 CIDR block for the VPC. CIDR can be explicitly set or it can be derived from IPAM using `ipv6NetmaskLength`. This parameter is required if `ipv6NetmaskLength` is not set and he IPAM pool does not have `allocationDefaultNetmask` set.
	Ipv6CidrBlock *string `pulumi:"ipv6CidrBlock"`
	// The ID of an IPv6 IPAM pool you want to use for allocating this VPC's CIDR. IPAM is a VPC feature that you can use to automate your IP address management workflows including assigning, tracking, troubleshooting, and auditing IP addresses across AWS Regions and accounts.
	Ipv6IpamPoolId *string `pulumi:"ipv6IpamPoolId"`
	// The netmask length of the IPv6 CIDR you want to allocate to this VPC. Requires specifying a `ipv6IpamPoolId`. This parameter is optional if the IPAM pool has `allocationDefaultNetmask` set, otherwise it or `cidrBlock` are required
	Ipv6NetmaskLength *int `pulumi:"ipv6NetmaskLength"`
	// The ID of the VPC to make the association with.
	VpcId *string `pulumi:"vpcId"`
}

type VpcIpv6CidrBlockAssociationState struct {
	// The IPv6 CIDR block for the VPC. CIDR can be explicitly set or it can be derived from IPAM using `ipv6NetmaskLength`. This parameter is required if `ipv6NetmaskLength` is not set and he IPAM pool does not have `allocationDefaultNetmask` set.
	Ipv6CidrBlock pulumi.StringPtrInput
	// The ID of an IPv6 IPAM pool you want to use for allocating this VPC's CIDR. IPAM is a VPC feature that you can use to automate your IP address management workflows including assigning, tracking, troubleshooting, and auditing IP addresses across AWS Regions and accounts.
	Ipv6IpamPoolId pulumi.StringPtrInput
	// The netmask length of the IPv6 CIDR you want to allocate to this VPC. Requires specifying a `ipv6IpamPoolId`. This parameter is optional if the IPAM pool has `allocationDefaultNetmask` set, otherwise it or `cidrBlock` are required
	Ipv6NetmaskLength pulumi.IntPtrInput
	// The ID of the VPC to make the association with.
	VpcId pulumi.StringPtrInput
}

func (VpcIpv6CidrBlockAssociationState) ElementType() reflect.Type {
	return reflect.TypeOf((*vpcIpv6CidrBlockAssociationState)(nil)).Elem()
}

type vpcIpv6CidrBlockAssociationArgs struct {
	// The IPv6 CIDR block for the VPC. CIDR can be explicitly set or it can be derived from IPAM using `ipv6NetmaskLength`. This parameter is required if `ipv6NetmaskLength` is not set and he IPAM pool does not have `allocationDefaultNetmask` set.
	Ipv6CidrBlock *string `pulumi:"ipv6CidrBlock"`
	// The ID of an IPv6 IPAM pool you want to use for allocating this VPC's CIDR. IPAM is a VPC feature that you can use to automate your IP address management workflows including assigning, tracking, troubleshooting, and auditing IP addresses across AWS Regions and accounts.
	Ipv6IpamPoolId string `pulumi:"ipv6IpamPoolId"`
	// The netmask length of the IPv6 CIDR you want to allocate to this VPC. Requires specifying a `ipv6IpamPoolId`. This parameter is optional if the IPAM pool has `allocationDefaultNetmask` set, otherwise it or `cidrBlock` are required
	Ipv6NetmaskLength *int `pulumi:"ipv6NetmaskLength"`
	// The ID of the VPC to make the association with.
	VpcId string `pulumi:"vpcId"`
}

// The set of arguments for constructing a VpcIpv6CidrBlockAssociation resource.
type VpcIpv6CidrBlockAssociationArgs struct {
	// The IPv6 CIDR block for the VPC. CIDR can be explicitly set or it can be derived from IPAM using `ipv6NetmaskLength`. This parameter is required if `ipv6NetmaskLength` is not set and he IPAM pool does not have `allocationDefaultNetmask` set.
	Ipv6CidrBlock pulumi.StringPtrInput
	// The ID of an IPv6 IPAM pool you want to use for allocating this VPC's CIDR. IPAM is a VPC feature that you can use to automate your IP address management workflows including assigning, tracking, troubleshooting, and auditing IP addresses across AWS Regions and accounts.
	Ipv6IpamPoolId pulumi.StringInput
	// The netmask length of the IPv6 CIDR you want to allocate to this VPC. Requires specifying a `ipv6IpamPoolId`. This parameter is optional if the IPAM pool has `allocationDefaultNetmask` set, otherwise it or `cidrBlock` are required
	Ipv6NetmaskLength pulumi.IntPtrInput
	// The ID of the VPC to make the association with.
	VpcId pulumi.StringInput
}

func (VpcIpv6CidrBlockAssociationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*vpcIpv6CidrBlockAssociationArgs)(nil)).Elem()
}

type VpcIpv6CidrBlockAssociationInput interface {
	pulumi.Input

	ToVpcIpv6CidrBlockAssociationOutput() VpcIpv6CidrBlockAssociationOutput
	ToVpcIpv6CidrBlockAssociationOutputWithContext(ctx context.Context) VpcIpv6CidrBlockAssociationOutput
}

func (*VpcIpv6CidrBlockAssociation) ElementType() reflect.Type {
	return reflect.TypeOf((**VpcIpv6CidrBlockAssociation)(nil)).Elem()
}

func (i *VpcIpv6CidrBlockAssociation) ToVpcIpv6CidrBlockAssociationOutput() VpcIpv6CidrBlockAssociationOutput {
	return i.ToVpcIpv6CidrBlockAssociationOutputWithContext(context.Background())
}

func (i *VpcIpv6CidrBlockAssociation) ToVpcIpv6CidrBlockAssociationOutputWithContext(ctx context.Context) VpcIpv6CidrBlockAssociationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VpcIpv6CidrBlockAssociationOutput)
}

func (i *VpcIpv6CidrBlockAssociation) ToOutput(ctx context.Context) pulumix.Output[*VpcIpv6CidrBlockAssociation] {
	return pulumix.Output[*VpcIpv6CidrBlockAssociation]{
		OutputState: i.ToVpcIpv6CidrBlockAssociationOutputWithContext(ctx).OutputState,
	}
}

// VpcIpv6CidrBlockAssociationArrayInput is an input type that accepts VpcIpv6CidrBlockAssociationArray and VpcIpv6CidrBlockAssociationArrayOutput values.
// You can construct a concrete instance of `VpcIpv6CidrBlockAssociationArrayInput` via:
//
//	VpcIpv6CidrBlockAssociationArray{ VpcIpv6CidrBlockAssociationArgs{...} }
type VpcIpv6CidrBlockAssociationArrayInput interface {
	pulumi.Input

	ToVpcIpv6CidrBlockAssociationArrayOutput() VpcIpv6CidrBlockAssociationArrayOutput
	ToVpcIpv6CidrBlockAssociationArrayOutputWithContext(context.Context) VpcIpv6CidrBlockAssociationArrayOutput
}

type VpcIpv6CidrBlockAssociationArray []VpcIpv6CidrBlockAssociationInput

func (VpcIpv6CidrBlockAssociationArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*VpcIpv6CidrBlockAssociation)(nil)).Elem()
}

func (i VpcIpv6CidrBlockAssociationArray) ToVpcIpv6CidrBlockAssociationArrayOutput() VpcIpv6CidrBlockAssociationArrayOutput {
	return i.ToVpcIpv6CidrBlockAssociationArrayOutputWithContext(context.Background())
}

func (i VpcIpv6CidrBlockAssociationArray) ToVpcIpv6CidrBlockAssociationArrayOutputWithContext(ctx context.Context) VpcIpv6CidrBlockAssociationArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VpcIpv6CidrBlockAssociationArrayOutput)
}

func (i VpcIpv6CidrBlockAssociationArray) ToOutput(ctx context.Context) pulumix.Output[[]*VpcIpv6CidrBlockAssociation] {
	return pulumix.Output[[]*VpcIpv6CidrBlockAssociation]{
		OutputState: i.ToVpcIpv6CidrBlockAssociationArrayOutputWithContext(ctx).OutputState,
	}
}

// VpcIpv6CidrBlockAssociationMapInput is an input type that accepts VpcIpv6CidrBlockAssociationMap and VpcIpv6CidrBlockAssociationMapOutput values.
// You can construct a concrete instance of `VpcIpv6CidrBlockAssociationMapInput` via:
//
//	VpcIpv6CidrBlockAssociationMap{ "key": VpcIpv6CidrBlockAssociationArgs{...} }
type VpcIpv6CidrBlockAssociationMapInput interface {
	pulumi.Input

	ToVpcIpv6CidrBlockAssociationMapOutput() VpcIpv6CidrBlockAssociationMapOutput
	ToVpcIpv6CidrBlockAssociationMapOutputWithContext(context.Context) VpcIpv6CidrBlockAssociationMapOutput
}

type VpcIpv6CidrBlockAssociationMap map[string]VpcIpv6CidrBlockAssociationInput

func (VpcIpv6CidrBlockAssociationMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*VpcIpv6CidrBlockAssociation)(nil)).Elem()
}

func (i VpcIpv6CidrBlockAssociationMap) ToVpcIpv6CidrBlockAssociationMapOutput() VpcIpv6CidrBlockAssociationMapOutput {
	return i.ToVpcIpv6CidrBlockAssociationMapOutputWithContext(context.Background())
}

func (i VpcIpv6CidrBlockAssociationMap) ToVpcIpv6CidrBlockAssociationMapOutputWithContext(ctx context.Context) VpcIpv6CidrBlockAssociationMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VpcIpv6CidrBlockAssociationMapOutput)
}

func (i VpcIpv6CidrBlockAssociationMap) ToOutput(ctx context.Context) pulumix.Output[map[string]*VpcIpv6CidrBlockAssociation] {
	return pulumix.Output[map[string]*VpcIpv6CidrBlockAssociation]{
		OutputState: i.ToVpcIpv6CidrBlockAssociationMapOutputWithContext(ctx).OutputState,
	}
}

type VpcIpv6CidrBlockAssociationOutput struct{ *pulumi.OutputState }

func (VpcIpv6CidrBlockAssociationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**VpcIpv6CidrBlockAssociation)(nil)).Elem()
}

func (o VpcIpv6CidrBlockAssociationOutput) ToVpcIpv6CidrBlockAssociationOutput() VpcIpv6CidrBlockAssociationOutput {
	return o
}

func (o VpcIpv6CidrBlockAssociationOutput) ToVpcIpv6CidrBlockAssociationOutputWithContext(ctx context.Context) VpcIpv6CidrBlockAssociationOutput {
	return o
}

func (o VpcIpv6CidrBlockAssociationOutput) ToOutput(ctx context.Context) pulumix.Output[*VpcIpv6CidrBlockAssociation] {
	return pulumix.Output[*VpcIpv6CidrBlockAssociation]{
		OutputState: o.OutputState,
	}
}

// The IPv6 CIDR block for the VPC. CIDR can be explicitly set or it can be derived from IPAM using `ipv6NetmaskLength`. This parameter is required if `ipv6NetmaskLength` is not set and he IPAM pool does not have `allocationDefaultNetmask` set.
func (o VpcIpv6CidrBlockAssociationOutput) Ipv6CidrBlock() pulumi.StringOutput {
	return o.ApplyT(func(v *VpcIpv6CidrBlockAssociation) pulumi.StringOutput { return v.Ipv6CidrBlock }).(pulumi.StringOutput)
}

// The ID of an IPv6 IPAM pool you want to use for allocating this VPC's CIDR. IPAM is a VPC feature that you can use to automate your IP address management workflows including assigning, tracking, troubleshooting, and auditing IP addresses across AWS Regions and accounts.
func (o VpcIpv6CidrBlockAssociationOutput) Ipv6IpamPoolId() pulumi.StringOutput {
	return o.ApplyT(func(v *VpcIpv6CidrBlockAssociation) pulumi.StringOutput { return v.Ipv6IpamPoolId }).(pulumi.StringOutput)
}

// The netmask length of the IPv6 CIDR you want to allocate to this VPC. Requires specifying a `ipv6IpamPoolId`. This parameter is optional if the IPAM pool has `allocationDefaultNetmask` set, otherwise it or `cidrBlock` are required
func (o VpcIpv6CidrBlockAssociationOutput) Ipv6NetmaskLength() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *VpcIpv6CidrBlockAssociation) pulumi.IntPtrOutput { return v.Ipv6NetmaskLength }).(pulumi.IntPtrOutput)
}

// The ID of the VPC to make the association with.
func (o VpcIpv6CidrBlockAssociationOutput) VpcId() pulumi.StringOutput {
	return o.ApplyT(func(v *VpcIpv6CidrBlockAssociation) pulumi.StringOutput { return v.VpcId }).(pulumi.StringOutput)
}

type VpcIpv6CidrBlockAssociationArrayOutput struct{ *pulumi.OutputState }

func (VpcIpv6CidrBlockAssociationArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*VpcIpv6CidrBlockAssociation)(nil)).Elem()
}

func (o VpcIpv6CidrBlockAssociationArrayOutput) ToVpcIpv6CidrBlockAssociationArrayOutput() VpcIpv6CidrBlockAssociationArrayOutput {
	return o
}

func (o VpcIpv6CidrBlockAssociationArrayOutput) ToVpcIpv6CidrBlockAssociationArrayOutputWithContext(ctx context.Context) VpcIpv6CidrBlockAssociationArrayOutput {
	return o
}

func (o VpcIpv6CidrBlockAssociationArrayOutput) ToOutput(ctx context.Context) pulumix.Output[[]*VpcIpv6CidrBlockAssociation] {
	return pulumix.Output[[]*VpcIpv6CidrBlockAssociation]{
		OutputState: o.OutputState,
	}
}

func (o VpcIpv6CidrBlockAssociationArrayOutput) Index(i pulumi.IntInput) VpcIpv6CidrBlockAssociationOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *VpcIpv6CidrBlockAssociation {
		return vs[0].([]*VpcIpv6CidrBlockAssociation)[vs[1].(int)]
	}).(VpcIpv6CidrBlockAssociationOutput)
}

type VpcIpv6CidrBlockAssociationMapOutput struct{ *pulumi.OutputState }

func (VpcIpv6CidrBlockAssociationMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*VpcIpv6CidrBlockAssociation)(nil)).Elem()
}

func (o VpcIpv6CidrBlockAssociationMapOutput) ToVpcIpv6CidrBlockAssociationMapOutput() VpcIpv6CidrBlockAssociationMapOutput {
	return o
}

func (o VpcIpv6CidrBlockAssociationMapOutput) ToVpcIpv6CidrBlockAssociationMapOutputWithContext(ctx context.Context) VpcIpv6CidrBlockAssociationMapOutput {
	return o
}

func (o VpcIpv6CidrBlockAssociationMapOutput) ToOutput(ctx context.Context) pulumix.Output[map[string]*VpcIpv6CidrBlockAssociation] {
	return pulumix.Output[map[string]*VpcIpv6CidrBlockAssociation]{
		OutputState: o.OutputState,
	}
}

func (o VpcIpv6CidrBlockAssociationMapOutput) MapIndex(k pulumi.StringInput) VpcIpv6CidrBlockAssociationOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *VpcIpv6CidrBlockAssociation {
		return vs[0].(map[string]*VpcIpv6CidrBlockAssociation)[vs[1].(string)]
	}).(VpcIpv6CidrBlockAssociationOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*VpcIpv6CidrBlockAssociationInput)(nil)).Elem(), &VpcIpv6CidrBlockAssociation{})
	pulumi.RegisterInputType(reflect.TypeOf((*VpcIpv6CidrBlockAssociationArrayInput)(nil)).Elem(), VpcIpv6CidrBlockAssociationArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*VpcIpv6CidrBlockAssociationMapInput)(nil)).Elem(), VpcIpv6CidrBlockAssociationMap{})
	pulumi.RegisterOutputType(VpcIpv6CidrBlockAssociationOutput{})
	pulumi.RegisterOutputType(VpcIpv6CidrBlockAssociationArrayOutput{})
	pulumi.RegisterOutputType(VpcIpv6CidrBlockAssociationMapOutput{})
}
