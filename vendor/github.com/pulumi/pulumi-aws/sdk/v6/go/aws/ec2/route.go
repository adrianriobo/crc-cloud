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

// Provides a resource to create a routing table entry (a route) in a VPC routing table.
//
// > **NOTE on Route Tables and Routes:** This provider currently provides both a standalone Route resource and a Route Table resource with routes defined in-line. At this time you cannot use a Route Table with in-line routes in conjunction with any Route resources. Doing so will cause a conflict of rule settings and will overwrite rules.
//
// > **NOTE on `gatewayId` attribute:** The AWS API is very forgiving with the resource ID passed in the `gatewayId` attribute. For example an `ec2.Route` resource can be created with an `ec2.NatGateway` or `ec2.EgressOnlyInternetGateway` ID specified for the `gatewayId` attribute. Specifying anything other than an `ec2.InternetGateway` or `ec2.VpnGateway` ID will lead to this provider reporting a permanent diff between your configuration and recorded state, as the AWS API returns the more-specific attribute. If you are experiencing constant diffs with an `ec2.Route` resource, the first thing to check is that the correct attribute is being specified.
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
//			_, err := ec2.NewRoute(ctx, "route", &ec2.RouteArgs{
//				RouteTableId:           pulumi.String("rtb-4fbb3ac4"),
//				DestinationCidrBlock:   pulumi.String("10.0.1.0/22"),
//				VpcPeeringConnectionId: pulumi.String("pcx-45ff3dc1"),
//			}, pulumi.DependsOn([]pulumi.Resource{
//				aws_route_table.Testing,
//			}))
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
// ## Example IPv6 Usage
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
//			vpc, err := ec2.NewVpc(ctx, "vpc", &ec2.VpcArgs{
//				CidrBlock:                    pulumi.String("10.1.0.0/16"),
//				AssignGeneratedIpv6CidrBlock: pulumi.Bool(true),
//			})
//			if err != nil {
//				return err
//			}
//			egress, err := ec2.NewEgressOnlyInternetGateway(ctx, "egress", &ec2.EgressOnlyInternetGatewayArgs{
//				VpcId: vpc.ID(),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = ec2.NewRoute(ctx, "route", &ec2.RouteArgs{
//				RouteTableId:             pulumi.String("rtb-4fbb3ac4"),
//				DestinationIpv6CidrBlock: pulumi.String("::/0"),
//				EgressOnlyGatewayId:      egress.ID(),
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
// Import a route in route table `rtb-656C65616E6F72` with an IPv4 destination CIDR of `10.42.0.0/16`:
//
// Import a route in route table `rtb-656C65616E6F72` with an IPv6 destination CIDR of `2620:0:2d0:200::8/125`:
//
// Import a route in route table `rtb-656C65616E6F72` with a managed prefix list destination of `pl-0570a1d2d725c16be`:
//
// __Using `pulumi import` to import__ individual routes using `ROUTETABLEID_DESTINATION`. Import [local routes](https://docs.aws.amazon.com/vpc/latest/userguide/VPC_Route_Tables.html#RouteTables) using the VPC's IPv4 or IPv6 CIDR blocks. For example:
//
// Import a route in route table `rtb-656C65616E6F72` with an IPv4 destination CIDR of `10.42.0.0/16`:
//
// ```sh
//
//	$ pulumi import aws:ec2/route:Route my_route rtb-656C65616E6F72_10.42.0.0/16
//
// ```
//
//	Import a route in route table `rtb-656C65616E6F72` with an IPv6 destination CIDR of `2620:0:2d0:200::8/125`:
//
// ```sh
//
//	$ pulumi import aws:ec2/route:Route my_route rtb-656C65616E6F72_2620:0:2d0:200::8/125
//
// ```
//
//	Import a route in route table `rtb-656C65616E6F72` with a managed prefix list destination of `pl-0570a1d2d725c16be`:
//
// ```sh
//
//	$ pulumi import aws:ec2/route:Route my_route rtb-656C65616E6F72_pl-0570a1d2d725c16be
//
// ```
type Route struct {
	pulumi.CustomResourceState

	// Identifier of a carrier gateway. This attribute can only be used when the VPC contains a subnet which is associated with a Wavelength Zone.
	CarrierGatewayId pulumi.StringPtrOutput `pulumi:"carrierGatewayId"`
	// The Amazon Resource Name (ARN) of a core network.
	CoreNetworkArn pulumi.StringPtrOutput `pulumi:"coreNetworkArn"`
	// The destination CIDR block.
	DestinationCidrBlock pulumi.StringPtrOutput `pulumi:"destinationCidrBlock"`
	// The destination IPv6 CIDR block.
	DestinationIpv6CidrBlock pulumi.StringPtrOutput `pulumi:"destinationIpv6CidrBlock"`
	// The ID of a managed prefix list destination.
	//
	// One of the following target arguments must be supplied:
	DestinationPrefixListId pulumi.StringPtrOutput `pulumi:"destinationPrefixListId"`
	// Identifier of a VPC Egress Only Internet Gateway.
	EgressOnlyGatewayId pulumi.StringPtrOutput `pulumi:"egressOnlyGatewayId"`
	// Identifier of a VPC internet gateway or a virtual private gateway. Specify `local` when updating a previously imported local route.
	GatewayId pulumi.StringPtrOutput `pulumi:"gatewayId"`
	// Identifier of an EC2 instance.
	InstanceId pulumi.StringOutput `pulumi:"instanceId"`
	// The AWS account ID of the owner of the EC2 instance.
	InstanceOwnerId pulumi.StringOutput `pulumi:"instanceOwnerId"`
	// Identifier of a Outpost local gateway.
	LocalGatewayId pulumi.StringPtrOutput `pulumi:"localGatewayId"`
	// Identifier of a VPC NAT gateway.
	NatGatewayId pulumi.StringPtrOutput `pulumi:"natGatewayId"`
	// Identifier of an EC2 network interface.
	NetworkInterfaceId pulumi.StringOutput `pulumi:"networkInterfaceId"`
	// How the route was created - `CreateRouteTable`, `CreateRoute` or `EnableVgwRoutePropagation`.
	Origin pulumi.StringOutput `pulumi:"origin"`
	// The ID of the routing table.
	//
	// One of the following destination arguments must be supplied:
	RouteTableId pulumi.StringOutput `pulumi:"routeTableId"`
	// The state of the route - `active` or `blackhole`.
	State pulumi.StringOutput `pulumi:"state"`
	// Identifier of an EC2 Transit Gateway.
	TransitGatewayId pulumi.StringPtrOutput `pulumi:"transitGatewayId"`
	// Identifier of a VPC Endpoint.
	VpcEndpointId pulumi.StringPtrOutput `pulumi:"vpcEndpointId"`
	// Identifier of a VPC peering connection.
	//
	// Note that the default route, mapping the VPC's CIDR block to "local", is created implicitly and cannot be specified.
	VpcPeeringConnectionId pulumi.StringPtrOutput `pulumi:"vpcPeeringConnectionId"`
}

// NewRoute registers a new resource with the given unique name, arguments, and options.
func NewRoute(ctx *pulumi.Context,
	name string, args *RouteArgs, opts ...pulumi.ResourceOption) (*Route, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.RouteTableId == nil {
		return nil, errors.New("invalid value for required argument 'RouteTableId'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Route
	err := ctx.RegisterResource("aws:ec2/route:Route", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetRoute gets an existing Route resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetRoute(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *RouteState, opts ...pulumi.ResourceOption) (*Route, error) {
	var resource Route
	err := ctx.ReadResource("aws:ec2/route:Route", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Route resources.
type routeState struct {
	// Identifier of a carrier gateway. This attribute can only be used when the VPC contains a subnet which is associated with a Wavelength Zone.
	CarrierGatewayId *string `pulumi:"carrierGatewayId"`
	// The Amazon Resource Name (ARN) of a core network.
	CoreNetworkArn *string `pulumi:"coreNetworkArn"`
	// The destination CIDR block.
	DestinationCidrBlock *string `pulumi:"destinationCidrBlock"`
	// The destination IPv6 CIDR block.
	DestinationIpv6CidrBlock *string `pulumi:"destinationIpv6CidrBlock"`
	// The ID of a managed prefix list destination.
	//
	// One of the following target arguments must be supplied:
	DestinationPrefixListId *string `pulumi:"destinationPrefixListId"`
	// Identifier of a VPC Egress Only Internet Gateway.
	EgressOnlyGatewayId *string `pulumi:"egressOnlyGatewayId"`
	// Identifier of a VPC internet gateway or a virtual private gateway. Specify `local` when updating a previously imported local route.
	GatewayId *string `pulumi:"gatewayId"`
	// Identifier of an EC2 instance.
	InstanceId *string `pulumi:"instanceId"`
	// The AWS account ID of the owner of the EC2 instance.
	InstanceOwnerId *string `pulumi:"instanceOwnerId"`
	// Identifier of a Outpost local gateway.
	LocalGatewayId *string `pulumi:"localGatewayId"`
	// Identifier of a VPC NAT gateway.
	NatGatewayId *string `pulumi:"natGatewayId"`
	// Identifier of an EC2 network interface.
	NetworkInterfaceId *string `pulumi:"networkInterfaceId"`
	// How the route was created - `CreateRouteTable`, `CreateRoute` or `EnableVgwRoutePropagation`.
	Origin *string `pulumi:"origin"`
	// The ID of the routing table.
	//
	// One of the following destination arguments must be supplied:
	RouteTableId *string `pulumi:"routeTableId"`
	// The state of the route - `active` or `blackhole`.
	State *string `pulumi:"state"`
	// Identifier of an EC2 Transit Gateway.
	TransitGatewayId *string `pulumi:"transitGatewayId"`
	// Identifier of a VPC Endpoint.
	VpcEndpointId *string `pulumi:"vpcEndpointId"`
	// Identifier of a VPC peering connection.
	//
	// Note that the default route, mapping the VPC's CIDR block to "local", is created implicitly and cannot be specified.
	VpcPeeringConnectionId *string `pulumi:"vpcPeeringConnectionId"`
}

type RouteState struct {
	// Identifier of a carrier gateway. This attribute can only be used when the VPC contains a subnet which is associated with a Wavelength Zone.
	CarrierGatewayId pulumi.StringPtrInput
	// The Amazon Resource Name (ARN) of a core network.
	CoreNetworkArn pulumi.StringPtrInput
	// The destination CIDR block.
	DestinationCidrBlock pulumi.StringPtrInput
	// The destination IPv6 CIDR block.
	DestinationIpv6CidrBlock pulumi.StringPtrInput
	// The ID of a managed prefix list destination.
	//
	// One of the following target arguments must be supplied:
	DestinationPrefixListId pulumi.StringPtrInput
	// Identifier of a VPC Egress Only Internet Gateway.
	EgressOnlyGatewayId pulumi.StringPtrInput
	// Identifier of a VPC internet gateway or a virtual private gateway. Specify `local` when updating a previously imported local route.
	GatewayId pulumi.StringPtrInput
	// Identifier of an EC2 instance.
	InstanceId pulumi.StringPtrInput
	// The AWS account ID of the owner of the EC2 instance.
	InstanceOwnerId pulumi.StringPtrInput
	// Identifier of a Outpost local gateway.
	LocalGatewayId pulumi.StringPtrInput
	// Identifier of a VPC NAT gateway.
	NatGatewayId pulumi.StringPtrInput
	// Identifier of an EC2 network interface.
	NetworkInterfaceId pulumi.StringPtrInput
	// How the route was created - `CreateRouteTable`, `CreateRoute` or `EnableVgwRoutePropagation`.
	Origin pulumi.StringPtrInput
	// The ID of the routing table.
	//
	// One of the following destination arguments must be supplied:
	RouteTableId pulumi.StringPtrInput
	// The state of the route - `active` or `blackhole`.
	State pulumi.StringPtrInput
	// Identifier of an EC2 Transit Gateway.
	TransitGatewayId pulumi.StringPtrInput
	// Identifier of a VPC Endpoint.
	VpcEndpointId pulumi.StringPtrInput
	// Identifier of a VPC peering connection.
	//
	// Note that the default route, mapping the VPC's CIDR block to "local", is created implicitly and cannot be specified.
	VpcPeeringConnectionId pulumi.StringPtrInput
}

func (RouteState) ElementType() reflect.Type {
	return reflect.TypeOf((*routeState)(nil)).Elem()
}

type routeArgs struct {
	// Identifier of a carrier gateway. This attribute can only be used when the VPC contains a subnet which is associated with a Wavelength Zone.
	CarrierGatewayId *string `pulumi:"carrierGatewayId"`
	// The Amazon Resource Name (ARN) of a core network.
	CoreNetworkArn *string `pulumi:"coreNetworkArn"`
	// The destination CIDR block.
	DestinationCidrBlock *string `pulumi:"destinationCidrBlock"`
	// The destination IPv6 CIDR block.
	DestinationIpv6CidrBlock *string `pulumi:"destinationIpv6CidrBlock"`
	// The ID of a managed prefix list destination.
	//
	// One of the following target arguments must be supplied:
	DestinationPrefixListId *string `pulumi:"destinationPrefixListId"`
	// Identifier of a VPC Egress Only Internet Gateway.
	EgressOnlyGatewayId *string `pulumi:"egressOnlyGatewayId"`
	// Identifier of a VPC internet gateway or a virtual private gateway. Specify `local` when updating a previously imported local route.
	GatewayId *string `pulumi:"gatewayId"`
	// Identifier of a Outpost local gateway.
	LocalGatewayId *string `pulumi:"localGatewayId"`
	// Identifier of a VPC NAT gateway.
	NatGatewayId *string `pulumi:"natGatewayId"`
	// Identifier of an EC2 network interface.
	NetworkInterfaceId *string `pulumi:"networkInterfaceId"`
	// The ID of the routing table.
	//
	// One of the following destination arguments must be supplied:
	RouteTableId string `pulumi:"routeTableId"`
	// Identifier of an EC2 Transit Gateway.
	TransitGatewayId *string `pulumi:"transitGatewayId"`
	// Identifier of a VPC Endpoint.
	VpcEndpointId *string `pulumi:"vpcEndpointId"`
	// Identifier of a VPC peering connection.
	//
	// Note that the default route, mapping the VPC's CIDR block to "local", is created implicitly and cannot be specified.
	VpcPeeringConnectionId *string `pulumi:"vpcPeeringConnectionId"`
}

// The set of arguments for constructing a Route resource.
type RouteArgs struct {
	// Identifier of a carrier gateway. This attribute can only be used when the VPC contains a subnet which is associated with a Wavelength Zone.
	CarrierGatewayId pulumi.StringPtrInput
	// The Amazon Resource Name (ARN) of a core network.
	CoreNetworkArn pulumi.StringPtrInput
	// The destination CIDR block.
	DestinationCidrBlock pulumi.StringPtrInput
	// The destination IPv6 CIDR block.
	DestinationIpv6CidrBlock pulumi.StringPtrInput
	// The ID of a managed prefix list destination.
	//
	// One of the following target arguments must be supplied:
	DestinationPrefixListId pulumi.StringPtrInput
	// Identifier of a VPC Egress Only Internet Gateway.
	EgressOnlyGatewayId pulumi.StringPtrInput
	// Identifier of a VPC internet gateway or a virtual private gateway. Specify `local` when updating a previously imported local route.
	GatewayId pulumi.StringPtrInput
	// Identifier of a Outpost local gateway.
	LocalGatewayId pulumi.StringPtrInput
	// Identifier of a VPC NAT gateway.
	NatGatewayId pulumi.StringPtrInput
	// Identifier of an EC2 network interface.
	NetworkInterfaceId pulumi.StringPtrInput
	// The ID of the routing table.
	//
	// One of the following destination arguments must be supplied:
	RouteTableId pulumi.StringInput
	// Identifier of an EC2 Transit Gateway.
	TransitGatewayId pulumi.StringPtrInput
	// Identifier of a VPC Endpoint.
	VpcEndpointId pulumi.StringPtrInput
	// Identifier of a VPC peering connection.
	//
	// Note that the default route, mapping the VPC's CIDR block to "local", is created implicitly and cannot be specified.
	VpcPeeringConnectionId pulumi.StringPtrInput
}

func (RouteArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*routeArgs)(nil)).Elem()
}

type RouteInput interface {
	pulumi.Input

	ToRouteOutput() RouteOutput
	ToRouteOutputWithContext(ctx context.Context) RouteOutput
}

func (*Route) ElementType() reflect.Type {
	return reflect.TypeOf((**Route)(nil)).Elem()
}

func (i *Route) ToRouteOutput() RouteOutput {
	return i.ToRouteOutputWithContext(context.Background())
}

func (i *Route) ToRouteOutputWithContext(ctx context.Context) RouteOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RouteOutput)
}

func (i *Route) ToOutput(ctx context.Context) pulumix.Output[*Route] {
	return pulumix.Output[*Route]{
		OutputState: i.ToRouteOutputWithContext(ctx).OutputState,
	}
}

// RouteArrayInput is an input type that accepts RouteArray and RouteArrayOutput values.
// You can construct a concrete instance of `RouteArrayInput` via:
//
//	RouteArray{ RouteArgs{...} }
type RouteArrayInput interface {
	pulumi.Input

	ToRouteArrayOutput() RouteArrayOutput
	ToRouteArrayOutputWithContext(context.Context) RouteArrayOutput
}

type RouteArray []RouteInput

func (RouteArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Route)(nil)).Elem()
}

func (i RouteArray) ToRouteArrayOutput() RouteArrayOutput {
	return i.ToRouteArrayOutputWithContext(context.Background())
}

func (i RouteArray) ToRouteArrayOutputWithContext(ctx context.Context) RouteArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RouteArrayOutput)
}

func (i RouteArray) ToOutput(ctx context.Context) pulumix.Output[[]*Route] {
	return pulumix.Output[[]*Route]{
		OutputState: i.ToRouteArrayOutputWithContext(ctx).OutputState,
	}
}

// RouteMapInput is an input type that accepts RouteMap and RouteMapOutput values.
// You can construct a concrete instance of `RouteMapInput` via:
//
//	RouteMap{ "key": RouteArgs{...} }
type RouteMapInput interface {
	pulumi.Input

	ToRouteMapOutput() RouteMapOutput
	ToRouteMapOutputWithContext(context.Context) RouteMapOutput
}

type RouteMap map[string]RouteInput

func (RouteMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Route)(nil)).Elem()
}

func (i RouteMap) ToRouteMapOutput() RouteMapOutput {
	return i.ToRouteMapOutputWithContext(context.Background())
}

func (i RouteMap) ToRouteMapOutputWithContext(ctx context.Context) RouteMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RouteMapOutput)
}

func (i RouteMap) ToOutput(ctx context.Context) pulumix.Output[map[string]*Route] {
	return pulumix.Output[map[string]*Route]{
		OutputState: i.ToRouteMapOutputWithContext(ctx).OutputState,
	}
}

type RouteOutput struct{ *pulumi.OutputState }

func (RouteOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Route)(nil)).Elem()
}

func (o RouteOutput) ToRouteOutput() RouteOutput {
	return o
}

func (o RouteOutput) ToRouteOutputWithContext(ctx context.Context) RouteOutput {
	return o
}

func (o RouteOutput) ToOutput(ctx context.Context) pulumix.Output[*Route] {
	return pulumix.Output[*Route]{
		OutputState: o.OutputState,
	}
}

// Identifier of a carrier gateway. This attribute can only be used when the VPC contains a subnet which is associated with a Wavelength Zone.
func (o RouteOutput) CarrierGatewayId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Route) pulumi.StringPtrOutput { return v.CarrierGatewayId }).(pulumi.StringPtrOutput)
}

// The Amazon Resource Name (ARN) of a core network.
func (o RouteOutput) CoreNetworkArn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Route) pulumi.StringPtrOutput { return v.CoreNetworkArn }).(pulumi.StringPtrOutput)
}

// The destination CIDR block.
func (o RouteOutput) DestinationCidrBlock() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Route) pulumi.StringPtrOutput { return v.DestinationCidrBlock }).(pulumi.StringPtrOutput)
}

// The destination IPv6 CIDR block.
func (o RouteOutput) DestinationIpv6CidrBlock() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Route) pulumi.StringPtrOutput { return v.DestinationIpv6CidrBlock }).(pulumi.StringPtrOutput)
}

// The ID of a managed prefix list destination.
//
// One of the following target arguments must be supplied:
func (o RouteOutput) DestinationPrefixListId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Route) pulumi.StringPtrOutput { return v.DestinationPrefixListId }).(pulumi.StringPtrOutput)
}

// Identifier of a VPC Egress Only Internet Gateway.
func (o RouteOutput) EgressOnlyGatewayId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Route) pulumi.StringPtrOutput { return v.EgressOnlyGatewayId }).(pulumi.StringPtrOutput)
}

// Identifier of a VPC internet gateway or a virtual private gateway. Specify `local` when updating a previously imported local route.
func (o RouteOutput) GatewayId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Route) pulumi.StringPtrOutput { return v.GatewayId }).(pulumi.StringPtrOutput)
}

// Identifier of an EC2 instance.
func (o RouteOutput) InstanceId() pulumi.StringOutput {
	return o.ApplyT(func(v *Route) pulumi.StringOutput { return v.InstanceId }).(pulumi.StringOutput)
}

// The AWS account ID of the owner of the EC2 instance.
func (o RouteOutput) InstanceOwnerId() pulumi.StringOutput {
	return o.ApplyT(func(v *Route) pulumi.StringOutput { return v.InstanceOwnerId }).(pulumi.StringOutput)
}

// Identifier of a Outpost local gateway.
func (o RouteOutput) LocalGatewayId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Route) pulumi.StringPtrOutput { return v.LocalGatewayId }).(pulumi.StringPtrOutput)
}

// Identifier of a VPC NAT gateway.
func (o RouteOutput) NatGatewayId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Route) pulumi.StringPtrOutput { return v.NatGatewayId }).(pulumi.StringPtrOutput)
}

// Identifier of an EC2 network interface.
func (o RouteOutput) NetworkInterfaceId() pulumi.StringOutput {
	return o.ApplyT(func(v *Route) pulumi.StringOutput { return v.NetworkInterfaceId }).(pulumi.StringOutput)
}

// How the route was created - `CreateRouteTable`, `CreateRoute` or `EnableVgwRoutePropagation`.
func (o RouteOutput) Origin() pulumi.StringOutput {
	return o.ApplyT(func(v *Route) pulumi.StringOutput { return v.Origin }).(pulumi.StringOutput)
}

// The ID of the routing table.
//
// One of the following destination arguments must be supplied:
func (o RouteOutput) RouteTableId() pulumi.StringOutput {
	return o.ApplyT(func(v *Route) pulumi.StringOutput { return v.RouteTableId }).(pulumi.StringOutput)
}

// The state of the route - `active` or `blackhole`.
func (o RouteOutput) State() pulumi.StringOutput {
	return o.ApplyT(func(v *Route) pulumi.StringOutput { return v.State }).(pulumi.StringOutput)
}

// Identifier of an EC2 Transit Gateway.
func (o RouteOutput) TransitGatewayId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Route) pulumi.StringPtrOutput { return v.TransitGatewayId }).(pulumi.StringPtrOutput)
}

// Identifier of a VPC Endpoint.
func (o RouteOutput) VpcEndpointId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Route) pulumi.StringPtrOutput { return v.VpcEndpointId }).(pulumi.StringPtrOutput)
}

// Identifier of a VPC peering connection.
//
// Note that the default route, mapping the VPC's CIDR block to "local", is created implicitly and cannot be specified.
func (o RouteOutput) VpcPeeringConnectionId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Route) pulumi.StringPtrOutput { return v.VpcPeeringConnectionId }).(pulumi.StringPtrOutput)
}

type RouteArrayOutput struct{ *pulumi.OutputState }

func (RouteArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Route)(nil)).Elem()
}

func (o RouteArrayOutput) ToRouteArrayOutput() RouteArrayOutput {
	return o
}

func (o RouteArrayOutput) ToRouteArrayOutputWithContext(ctx context.Context) RouteArrayOutput {
	return o
}

func (o RouteArrayOutput) ToOutput(ctx context.Context) pulumix.Output[[]*Route] {
	return pulumix.Output[[]*Route]{
		OutputState: o.OutputState,
	}
}

func (o RouteArrayOutput) Index(i pulumi.IntInput) RouteOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Route {
		return vs[0].([]*Route)[vs[1].(int)]
	}).(RouteOutput)
}

type RouteMapOutput struct{ *pulumi.OutputState }

func (RouteMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Route)(nil)).Elem()
}

func (o RouteMapOutput) ToRouteMapOutput() RouteMapOutput {
	return o
}

func (o RouteMapOutput) ToRouteMapOutputWithContext(ctx context.Context) RouteMapOutput {
	return o
}

func (o RouteMapOutput) ToOutput(ctx context.Context) pulumix.Output[map[string]*Route] {
	return pulumix.Output[map[string]*Route]{
		OutputState: o.OutputState,
	}
}

func (o RouteMapOutput) MapIndex(k pulumi.StringInput) RouteOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Route {
		return vs[0].(map[string]*Route)[vs[1].(string)]
	}).(RouteOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*RouteInput)(nil)).Elem(), &Route{})
	pulumi.RegisterInputType(reflect.TypeOf((*RouteArrayInput)(nil)).Elem(), RouteArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*RouteMapInput)(nil)).Elem(), RouteMap{})
	pulumi.RegisterOutputType(RouteOutput{})
	pulumi.RegisterOutputType(RouteArrayOutput{})
	pulumi.RegisterOutputType(RouteMapOutput{})
}
