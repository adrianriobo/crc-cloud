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

// Allocates (reserves) a CIDR from an IPAM address pool, preventing usage by IPAM. Only works for private IPv4.
//
// ## Example Usage
//
// Basic usage:
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws"
//	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/ec2"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			current, err := aws.GetRegion(ctx, nil, nil)
//			if err != nil {
//				return err
//			}
//			exampleVpcIpam, err := ec2.NewVpcIpam(ctx, "exampleVpcIpam", &ec2.VpcIpamArgs{
//				OperatingRegions: ec2.VpcIpamOperatingRegionArray{
//					&ec2.VpcIpamOperatingRegionArgs{
//						RegionName: *pulumi.String(current.Name),
//					},
//				},
//			})
//			if err != nil {
//				return err
//			}
//			exampleVpcIpamPool, err := ec2.NewVpcIpamPool(ctx, "exampleVpcIpamPool", &ec2.VpcIpamPoolArgs{
//				AddressFamily: pulumi.String("ipv4"),
//				IpamScopeId:   exampleVpcIpam.PrivateDefaultScopeId,
//				Locale:        *pulumi.String(current.Name),
//			})
//			if err != nil {
//				return err
//			}
//			exampleVpcIpamPoolCidr, err := ec2.NewVpcIpamPoolCidr(ctx, "exampleVpcIpamPoolCidr", &ec2.VpcIpamPoolCidrArgs{
//				IpamPoolId: exampleVpcIpamPool.ID(),
//				Cidr:       pulumi.String("172.20.0.0/16"),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = ec2.NewVpcIpamPoolCidrAllocation(ctx, "exampleVpcIpamPoolCidrAllocation", &ec2.VpcIpamPoolCidrAllocationArgs{
//				IpamPoolId: exampleVpcIpamPool.ID(),
//				Cidr:       pulumi.String("172.20.0.0/24"),
//			}, pulumi.DependsOn([]pulumi.Resource{
//				exampleVpcIpamPoolCidr,
//			}))
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
//
// With the `disallowedCidrs` attribute:
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws"
//	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/ec2"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			current, err := aws.GetRegion(ctx, nil, nil)
//			if err != nil {
//				return err
//			}
//			exampleVpcIpam, err := ec2.NewVpcIpam(ctx, "exampleVpcIpam", &ec2.VpcIpamArgs{
//				OperatingRegions: ec2.VpcIpamOperatingRegionArray{
//					&ec2.VpcIpamOperatingRegionArgs{
//						RegionName: *pulumi.String(current.Name),
//					},
//				},
//			})
//			if err != nil {
//				return err
//			}
//			exampleVpcIpamPool, err := ec2.NewVpcIpamPool(ctx, "exampleVpcIpamPool", &ec2.VpcIpamPoolArgs{
//				AddressFamily: pulumi.String("ipv4"),
//				IpamScopeId:   exampleVpcIpam.PrivateDefaultScopeId,
//				Locale:        *pulumi.String(current.Name),
//			})
//			if err != nil {
//				return err
//			}
//			exampleVpcIpamPoolCidr, err := ec2.NewVpcIpamPoolCidr(ctx, "exampleVpcIpamPoolCidr", &ec2.VpcIpamPoolCidrArgs{
//				IpamPoolId: exampleVpcIpamPool.ID(),
//				Cidr:       pulumi.String("172.20.0.0/16"),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = ec2.NewVpcIpamPoolCidrAllocation(ctx, "exampleVpcIpamPoolCidrAllocation", &ec2.VpcIpamPoolCidrAllocationArgs{
//				IpamPoolId:    exampleVpcIpamPool.ID(),
//				NetmaskLength: pulumi.Int(28),
//				DisallowedCidrs: pulumi.StringArray{
//					pulumi.String("172.20.0.0/28"),
//				},
//			}, pulumi.DependsOn([]pulumi.Resource{
//				exampleVpcIpamPoolCidr,
//			}))
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
// Using `pulumi import`, import IPAM allocations using the allocation `id` and `pool id`, separated by `_`. For example:
//
// ```sh
//
//	$ pulumi import aws:ec2/vpcIpamPoolCidrAllocation:VpcIpamPoolCidrAllocation example ipam-pool-alloc-0dc6d196509c049ba8b549ff99f639736_ipam-pool-07cfb559e0921fcbe
//
// ```
type VpcIpamPoolCidrAllocation struct {
	pulumi.CustomResourceState

	// The CIDR you want to assign to the pool.
	Cidr pulumi.StringOutput `pulumi:"cidr"`
	// The description for the allocation.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// Exclude a particular CIDR range from being returned by the pool.
	DisallowedCidrs      pulumi.StringArrayOutput `pulumi:"disallowedCidrs"`
	IpamPoolAllocationId pulumi.StringOutput      `pulumi:"ipamPoolAllocationId"`
	// The ID of the pool to which you want to assign a CIDR.
	IpamPoolId pulumi.StringOutput `pulumi:"ipamPoolId"`
	// The netmask length of the CIDR you would like to allocate to the IPAM pool. Valid Values: `0-128`.
	NetmaskLength pulumi.IntPtrOutput `pulumi:"netmaskLength"`
	// The ID of the resource.
	ResourceId pulumi.StringOutput `pulumi:"resourceId"`
	// The owner of the resource.
	ResourceOwner pulumi.StringOutput `pulumi:"resourceOwner"`
	// The type of the resource.
	ResourceType pulumi.StringOutput `pulumi:"resourceType"`
}

// NewVpcIpamPoolCidrAllocation registers a new resource with the given unique name, arguments, and options.
func NewVpcIpamPoolCidrAllocation(ctx *pulumi.Context,
	name string, args *VpcIpamPoolCidrAllocationArgs, opts ...pulumi.ResourceOption) (*VpcIpamPoolCidrAllocation, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.IpamPoolId == nil {
		return nil, errors.New("invalid value for required argument 'IpamPoolId'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource VpcIpamPoolCidrAllocation
	err := ctx.RegisterResource("aws:ec2/vpcIpamPoolCidrAllocation:VpcIpamPoolCidrAllocation", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetVpcIpamPoolCidrAllocation gets an existing VpcIpamPoolCidrAllocation resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetVpcIpamPoolCidrAllocation(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *VpcIpamPoolCidrAllocationState, opts ...pulumi.ResourceOption) (*VpcIpamPoolCidrAllocation, error) {
	var resource VpcIpamPoolCidrAllocation
	err := ctx.ReadResource("aws:ec2/vpcIpamPoolCidrAllocation:VpcIpamPoolCidrAllocation", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering VpcIpamPoolCidrAllocation resources.
type vpcIpamPoolCidrAllocationState struct {
	// The CIDR you want to assign to the pool.
	Cidr *string `pulumi:"cidr"`
	// The description for the allocation.
	Description *string `pulumi:"description"`
	// Exclude a particular CIDR range from being returned by the pool.
	DisallowedCidrs      []string `pulumi:"disallowedCidrs"`
	IpamPoolAllocationId *string  `pulumi:"ipamPoolAllocationId"`
	// The ID of the pool to which you want to assign a CIDR.
	IpamPoolId *string `pulumi:"ipamPoolId"`
	// The netmask length of the CIDR you would like to allocate to the IPAM pool. Valid Values: `0-128`.
	NetmaskLength *int `pulumi:"netmaskLength"`
	// The ID of the resource.
	ResourceId *string `pulumi:"resourceId"`
	// The owner of the resource.
	ResourceOwner *string `pulumi:"resourceOwner"`
	// The type of the resource.
	ResourceType *string `pulumi:"resourceType"`
}

type VpcIpamPoolCidrAllocationState struct {
	// The CIDR you want to assign to the pool.
	Cidr pulumi.StringPtrInput
	// The description for the allocation.
	Description pulumi.StringPtrInput
	// Exclude a particular CIDR range from being returned by the pool.
	DisallowedCidrs      pulumi.StringArrayInput
	IpamPoolAllocationId pulumi.StringPtrInput
	// The ID of the pool to which you want to assign a CIDR.
	IpamPoolId pulumi.StringPtrInput
	// The netmask length of the CIDR you would like to allocate to the IPAM pool. Valid Values: `0-128`.
	NetmaskLength pulumi.IntPtrInput
	// The ID of the resource.
	ResourceId pulumi.StringPtrInput
	// The owner of the resource.
	ResourceOwner pulumi.StringPtrInput
	// The type of the resource.
	ResourceType pulumi.StringPtrInput
}

func (VpcIpamPoolCidrAllocationState) ElementType() reflect.Type {
	return reflect.TypeOf((*vpcIpamPoolCidrAllocationState)(nil)).Elem()
}

type vpcIpamPoolCidrAllocationArgs struct {
	// The CIDR you want to assign to the pool.
	Cidr *string `pulumi:"cidr"`
	// The description for the allocation.
	Description *string `pulumi:"description"`
	// Exclude a particular CIDR range from being returned by the pool.
	DisallowedCidrs []string `pulumi:"disallowedCidrs"`
	// The ID of the pool to which you want to assign a CIDR.
	IpamPoolId string `pulumi:"ipamPoolId"`
	// The netmask length of the CIDR you would like to allocate to the IPAM pool. Valid Values: `0-128`.
	NetmaskLength *int `pulumi:"netmaskLength"`
}

// The set of arguments for constructing a VpcIpamPoolCidrAllocation resource.
type VpcIpamPoolCidrAllocationArgs struct {
	// The CIDR you want to assign to the pool.
	Cidr pulumi.StringPtrInput
	// The description for the allocation.
	Description pulumi.StringPtrInput
	// Exclude a particular CIDR range from being returned by the pool.
	DisallowedCidrs pulumi.StringArrayInput
	// The ID of the pool to which you want to assign a CIDR.
	IpamPoolId pulumi.StringInput
	// The netmask length of the CIDR you would like to allocate to the IPAM pool. Valid Values: `0-128`.
	NetmaskLength pulumi.IntPtrInput
}

func (VpcIpamPoolCidrAllocationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*vpcIpamPoolCidrAllocationArgs)(nil)).Elem()
}

type VpcIpamPoolCidrAllocationInput interface {
	pulumi.Input

	ToVpcIpamPoolCidrAllocationOutput() VpcIpamPoolCidrAllocationOutput
	ToVpcIpamPoolCidrAllocationOutputWithContext(ctx context.Context) VpcIpamPoolCidrAllocationOutput
}

func (*VpcIpamPoolCidrAllocation) ElementType() reflect.Type {
	return reflect.TypeOf((**VpcIpamPoolCidrAllocation)(nil)).Elem()
}

func (i *VpcIpamPoolCidrAllocation) ToVpcIpamPoolCidrAllocationOutput() VpcIpamPoolCidrAllocationOutput {
	return i.ToVpcIpamPoolCidrAllocationOutputWithContext(context.Background())
}

func (i *VpcIpamPoolCidrAllocation) ToVpcIpamPoolCidrAllocationOutputWithContext(ctx context.Context) VpcIpamPoolCidrAllocationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VpcIpamPoolCidrAllocationOutput)
}

func (i *VpcIpamPoolCidrAllocation) ToOutput(ctx context.Context) pulumix.Output[*VpcIpamPoolCidrAllocation] {
	return pulumix.Output[*VpcIpamPoolCidrAllocation]{
		OutputState: i.ToVpcIpamPoolCidrAllocationOutputWithContext(ctx).OutputState,
	}
}

// VpcIpamPoolCidrAllocationArrayInput is an input type that accepts VpcIpamPoolCidrAllocationArray and VpcIpamPoolCidrAllocationArrayOutput values.
// You can construct a concrete instance of `VpcIpamPoolCidrAllocationArrayInput` via:
//
//	VpcIpamPoolCidrAllocationArray{ VpcIpamPoolCidrAllocationArgs{...} }
type VpcIpamPoolCidrAllocationArrayInput interface {
	pulumi.Input

	ToVpcIpamPoolCidrAllocationArrayOutput() VpcIpamPoolCidrAllocationArrayOutput
	ToVpcIpamPoolCidrAllocationArrayOutputWithContext(context.Context) VpcIpamPoolCidrAllocationArrayOutput
}

type VpcIpamPoolCidrAllocationArray []VpcIpamPoolCidrAllocationInput

func (VpcIpamPoolCidrAllocationArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*VpcIpamPoolCidrAllocation)(nil)).Elem()
}

func (i VpcIpamPoolCidrAllocationArray) ToVpcIpamPoolCidrAllocationArrayOutput() VpcIpamPoolCidrAllocationArrayOutput {
	return i.ToVpcIpamPoolCidrAllocationArrayOutputWithContext(context.Background())
}

func (i VpcIpamPoolCidrAllocationArray) ToVpcIpamPoolCidrAllocationArrayOutputWithContext(ctx context.Context) VpcIpamPoolCidrAllocationArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VpcIpamPoolCidrAllocationArrayOutput)
}

func (i VpcIpamPoolCidrAllocationArray) ToOutput(ctx context.Context) pulumix.Output[[]*VpcIpamPoolCidrAllocation] {
	return pulumix.Output[[]*VpcIpamPoolCidrAllocation]{
		OutputState: i.ToVpcIpamPoolCidrAllocationArrayOutputWithContext(ctx).OutputState,
	}
}

// VpcIpamPoolCidrAllocationMapInput is an input type that accepts VpcIpamPoolCidrAllocationMap and VpcIpamPoolCidrAllocationMapOutput values.
// You can construct a concrete instance of `VpcIpamPoolCidrAllocationMapInput` via:
//
//	VpcIpamPoolCidrAllocationMap{ "key": VpcIpamPoolCidrAllocationArgs{...} }
type VpcIpamPoolCidrAllocationMapInput interface {
	pulumi.Input

	ToVpcIpamPoolCidrAllocationMapOutput() VpcIpamPoolCidrAllocationMapOutput
	ToVpcIpamPoolCidrAllocationMapOutputWithContext(context.Context) VpcIpamPoolCidrAllocationMapOutput
}

type VpcIpamPoolCidrAllocationMap map[string]VpcIpamPoolCidrAllocationInput

func (VpcIpamPoolCidrAllocationMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*VpcIpamPoolCidrAllocation)(nil)).Elem()
}

func (i VpcIpamPoolCidrAllocationMap) ToVpcIpamPoolCidrAllocationMapOutput() VpcIpamPoolCidrAllocationMapOutput {
	return i.ToVpcIpamPoolCidrAllocationMapOutputWithContext(context.Background())
}

func (i VpcIpamPoolCidrAllocationMap) ToVpcIpamPoolCidrAllocationMapOutputWithContext(ctx context.Context) VpcIpamPoolCidrAllocationMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VpcIpamPoolCidrAllocationMapOutput)
}

func (i VpcIpamPoolCidrAllocationMap) ToOutput(ctx context.Context) pulumix.Output[map[string]*VpcIpamPoolCidrAllocation] {
	return pulumix.Output[map[string]*VpcIpamPoolCidrAllocation]{
		OutputState: i.ToVpcIpamPoolCidrAllocationMapOutputWithContext(ctx).OutputState,
	}
}

type VpcIpamPoolCidrAllocationOutput struct{ *pulumi.OutputState }

func (VpcIpamPoolCidrAllocationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**VpcIpamPoolCidrAllocation)(nil)).Elem()
}

func (o VpcIpamPoolCidrAllocationOutput) ToVpcIpamPoolCidrAllocationOutput() VpcIpamPoolCidrAllocationOutput {
	return o
}

func (o VpcIpamPoolCidrAllocationOutput) ToVpcIpamPoolCidrAllocationOutputWithContext(ctx context.Context) VpcIpamPoolCidrAllocationOutput {
	return o
}

func (o VpcIpamPoolCidrAllocationOutput) ToOutput(ctx context.Context) pulumix.Output[*VpcIpamPoolCidrAllocation] {
	return pulumix.Output[*VpcIpamPoolCidrAllocation]{
		OutputState: o.OutputState,
	}
}

// The CIDR you want to assign to the pool.
func (o VpcIpamPoolCidrAllocationOutput) Cidr() pulumi.StringOutput {
	return o.ApplyT(func(v *VpcIpamPoolCidrAllocation) pulumi.StringOutput { return v.Cidr }).(pulumi.StringOutput)
}

// The description for the allocation.
func (o VpcIpamPoolCidrAllocationOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VpcIpamPoolCidrAllocation) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// Exclude a particular CIDR range from being returned by the pool.
func (o VpcIpamPoolCidrAllocationOutput) DisallowedCidrs() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *VpcIpamPoolCidrAllocation) pulumi.StringArrayOutput { return v.DisallowedCidrs }).(pulumi.StringArrayOutput)
}

func (o VpcIpamPoolCidrAllocationOutput) IpamPoolAllocationId() pulumi.StringOutput {
	return o.ApplyT(func(v *VpcIpamPoolCidrAllocation) pulumi.StringOutput { return v.IpamPoolAllocationId }).(pulumi.StringOutput)
}

// The ID of the pool to which you want to assign a CIDR.
func (o VpcIpamPoolCidrAllocationOutput) IpamPoolId() pulumi.StringOutput {
	return o.ApplyT(func(v *VpcIpamPoolCidrAllocation) pulumi.StringOutput { return v.IpamPoolId }).(pulumi.StringOutput)
}

// The netmask length of the CIDR you would like to allocate to the IPAM pool. Valid Values: `0-128`.
func (o VpcIpamPoolCidrAllocationOutput) NetmaskLength() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *VpcIpamPoolCidrAllocation) pulumi.IntPtrOutput { return v.NetmaskLength }).(pulumi.IntPtrOutput)
}

// The ID of the resource.
func (o VpcIpamPoolCidrAllocationOutput) ResourceId() pulumi.StringOutput {
	return o.ApplyT(func(v *VpcIpamPoolCidrAllocation) pulumi.StringOutput { return v.ResourceId }).(pulumi.StringOutput)
}

// The owner of the resource.
func (o VpcIpamPoolCidrAllocationOutput) ResourceOwner() pulumi.StringOutput {
	return o.ApplyT(func(v *VpcIpamPoolCidrAllocation) pulumi.StringOutput { return v.ResourceOwner }).(pulumi.StringOutput)
}

// The type of the resource.
func (o VpcIpamPoolCidrAllocationOutput) ResourceType() pulumi.StringOutput {
	return o.ApplyT(func(v *VpcIpamPoolCidrAllocation) pulumi.StringOutput { return v.ResourceType }).(pulumi.StringOutput)
}

type VpcIpamPoolCidrAllocationArrayOutput struct{ *pulumi.OutputState }

func (VpcIpamPoolCidrAllocationArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*VpcIpamPoolCidrAllocation)(nil)).Elem()
}

func (o VpcIpamPoolCidrAllocationArrayOutput) ToVpcIpamPoolCidrAllocationArrayOutput() VpcIpamPoolCidrAllocationArrayOutput {
	return o
}

func (o VpcIpamPoolCidrAllocationArrayOutput) ToVpcIpamPoolCidrAllocationArrayOutputWithContext(ctx context.Context) VpcIpamPoolCidrAllocationArrayOutput {
	return o
}

func (o VpcIpamPoolCidrAllocationArrayOutput) ToOutput(ctx context.Context) pulumix.Output[[]*VpcIpamPoolCidrAllocation] {
	return pulumix.Output[[]*VpcIpamPoolCidrAllocation]{
		OutputState: o.OutputState,
	}
}

func (o VpcIpamPoolCidrAllocationArrayOutput) Index(i pulumi.IntInput) VpcIpamPoolCidrAllocationOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *VpcIpamPoolCidrAllocation {
		return vs[0].([]*VpcIpamPoolCidrAllocation)[vs[1].(int)]
	}).(VpcIpamPoolCidrAllocationOutput)
}

type VpcIpamPoolCidrAllocationMapOutput struct{ *pulumi.OutputState }

func (VpcIpamPoolCidrAllocationMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*VpcIpamPoolCidrAllocation)(nil)).Elem()
}

func (o VpcIpamPoolCidrAllocationMapOutput) ToVpcIpamPoolCidrAllocationMapOutput() VpcIpamPoolCidrAllocationMapOutput {
	return o
}

func (o VpcIpamPoolCidrAllocationMapOutput) ToVpcIpamPoolCidrAllocationMapOutputWithContext(ctx context.Context) VpcIpamPoolCidrAllocationMapOutput {
	return o
}

func (o VpcIpamPoolCidrAllocationMapOutput) ToOutput(ctx context.Context) pulumix.Output[map[string]*VpcIpamPoolCidrAllocation] {
	return pulumix.Output[map[string]*VpcIpamPoolCidrAllocation]{
		OutputState: o.OutputState,
	}
}

func (o VpcIpamPoolCidrAllocationMapOutput) MapIndex(k pulumi.StringInput) VpcIpamPoolCidrAllocationOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *VpcIpamPoolCidrAllocation {
		return vs[0].(map[string]*VpcIpamPoolCidrAllocation)[vs[1].(string)]
	}).(VpcIpamPoolCidrAllocationOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*VpcIpamPoolCidrAllocationInput)(nil)).Elem(), &VpcIpamPoolCidrAllocation{})
	pulumi.RegisterInputType(reflect.TypeOf((*VpcIpamPoolCidrAllocationArrayInput)(nil)).Elem(), VpcIpamPoolCidrAllocationArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*VpcIpamPoolCidrAllocationMapInput)(nil)).Elem(), VpcIpamPoolCidrAllocationMap{})
	pulumi.RegisterOutputType(VpcIpamPoolCidrAllocationOutput{})
	pulumi.RegisterOutputType(VpcIpamPoolCidrAllocationArrayOutput{})
	pulumi.RegisterOutputType(VpcIpamPoolCidrAllocationMapOutput{})
}
