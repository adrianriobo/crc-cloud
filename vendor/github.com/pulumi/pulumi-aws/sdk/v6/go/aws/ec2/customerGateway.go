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

// Provides a customer gateway inside a VPC. These objects can be connected to VPN gateways via VPN connections, and allow you to establish tunnels between your network and the VPC.
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
//			_, err := ec2.NewCustomerGateway(ctx, "main", &ec2.CustomerGatewayArgs{
//				BgpAsn:    pulumi.String("65000"),
//				IpAddress: pulumi.String("172.83.124.10"),
//				Tags: pulumi.StringMap{
//					"Name": pulumi.String("main-customer-gateway"),
//				},
//				Type: pulumi.String("ipsec.1"),
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
// Using `pulumi import`, import Customer Gateways using the `id`. For example:
//
// ```sh
//
//	$ pulumi import aws:ec2/customerGateway:CustomerGateway main cgw-b4dc3961
//
// ```
type CustomerGateway struct {
	pulumi.CustomResourceState

	// The ARN of the customer gateway.
	Arn pulumi.StringOutput `pulumi:"arn"`
	// The gateway's Border Gateway Protocol (BGP) Autonomous System Number (ASN).
	BgpAsn pulumi.StringOutput `pulumi:"bgpAsn"`
	// The Amazon Resource Name (ARN) for the customer gateway certificate.
	CertificateArn pulumi.StringPtrOutput `pulumi:"certificateArn"`
	// A name for the customer gateway device.
	DeviceName pulumi.StringPtrOutput `pulumi:"deviceName"`
	// The IPv4 address for the customer gateway device's outside interface.
	IpAddress pulumi.StringPtrOutput `pulumi:"ipAddress"`
	// Tags to apply to the gateway. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// A map of tags assigned to the resource, including those inherited from the provider `defaultTags` configuration block.
	TagsAll pulumi.StringMapOutput `pulumi:"tagsAll"`
	// The type of customer gateway. The only type AWS
	// supports at this time is "ipsec.1".
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewCustomerGateway registers a new resource with the given unique name, arguments, and options.
func NewCustomerGateway(ctx *pulumi.Context,
	name string, args *CustomerGatewayArgs, opts ...pulumi.ResourceOption) (*CustomerGateway, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.BgpAsn == nil {
		return nil, errors.New("invalid value for required argument 'BgpAsn'")
	}
	if args.Type == nil {
		return nil, errors.New("invalid value for required argument 'Type'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource CustomerGateway
	err := ctx.RegisterResource("aws:ec2/customerGateway:CustomerGateway", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetCustomerGateway gets an existing CustomerGateway resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetCustomerGateway(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *CustomerGatewayState, opts ...pulumi.ResourceOption) (*CustomerGateway, error) {
	var resource CustomerGateway
	err := ctx.ReadResource("aws:ec2/customerGateway:CustomerGateway", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering CustomerGateway resources.
type customerGatewayState struct {
	// The ARN of the customer gateway.
	Arn *string `pulumi:"arn"`
	// The gateway's Border Gateway Protocol (BGP) Autonomous System Number (ASN).
	BgpAsn *string `pulumi:"bgpAsn"`
	// The Amazon Resource Name (ARN) for the customer gateway certificate.
	CertificateArn *string `pulumi:"certificateArn"`
	// A name for the customer gateway device.
	DeviceName *string `pulumi:"deviceName"`
	// The IPv4 address for the customer gateway device's outside interface.
	IpAddress *string `pulumi:"ipAddress"`
	// Tags to apply to the gateway. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `pulumi:"tags"`
	// A map of tags assigned to the resource, including those inherited from the provider `defaultTags` configuration block.
	TagsAll map[string]string `pulumi:"tagsAll"`
	// The type of customer gateway. The only type AWS
	// supports at this time is "ipsec.1".
	Type *string `pulumi:"type"`
}

type CustomerGatewayState struct {
	// The ARN of the customer gateway.
	Arn pulumi.StringPtrInput
	// The gateway's Border Gateway Protocol (BGP) Autonomous System Number (ASN).
	BgpAsn pulumi.StringPtrInput
	// The Amazon Resource Name (ARN) for the customer gateway certificate.
	CertificateArn pulumi.StringPtrInput
	// A name for the customer gateway device.
	DeviceName pulumi.StringPtrInput
	// The IPv4 address for the customer gateway device's outside interface.
	IpAddress pulumi.StringPtrInput
	// Tags to apply to the gateway. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags pulumi.StringMapInput
	// A map of tags assigned to the resource, including those inherited from the provider `defaultTags` configuration block.
	TagsAll pulumi.StringMapInput
	// The type of customer gateway. The only type AWS
	// supports at this time is "ipsec.1".
	Type pulumi.StringPtrInput
}

func (CustomerGatewayState) ElementType() reflect.Type {
	return reflect.TypeOf((*customerGatewayState)(nil)).Elem()
}

type customerGatewayArgs struct {
	// The gateway's Border Gateway Protocol (BGP) Autonomous System Number (ASN).
	BgpAsn string `pulumi:"bgpAsn"`
	// The Amazon Resource Name (ARN) for the customer gateway certificate.
	CertificateArn *string `pulumi:"certificateArn"`
	// A name for the customer gateway device.
	DeviceName *string `pulumi:"deviceName"`
	// The IPv4 address for the customer gateway device's outside interface.
	IpAddress *string `pulumi:"ipAddress"`
	// Tags to apply to the gateway. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `pulumi:"tags"`
	// The type of customer gateway. The only type AWS
	// supports at this time is "ipsec.1".
	Type string `pulumi:"type"`
}

// The set of arguments for constructing a CustomerGateway resource.
type CustomerGatewayArgs struct {
	// The gateway's Border Gateway Protocol (BGP) Autonomous System Number (ASN).
	BgpAsn pulumi.StringInput
	// The Amazon Resource Name (ARN) for the customer gateway certificate.
	CertificateArn pulumi.StringPtrInput
	// A name for the customer gateway device.
	DeviceName pulumi.StringPtrInput
	// The IPv4 address for the customer gateway device's outside interface.
	IpAddress pulumi.StringPtrInput
	// Tags to apply to the gateway. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags pulumi.StringMapInput
	// The type of customer gateway. The only type AWS
	// supports at this time is "ipsec.1".
	Type pulumi.StringInput
}

func (CustomerGatewayArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*customerGatewayArgs)(nil)).Elem()
}

type CustomerGatewayInput interface {
	pulumi.Input

	ToCustomerGatewayOutput() CustomerGatewayOutput
	ToCustomerGatewayOutputWithContext(ctx context.Context) CustomerGatewayOutput
}

func (*CustomerGateway) ElementType() reflect.Type {
	return reflect.TypeOf((**CustomerGateway)(nil)).Elem()
}

func (i *CustomerGateway) ToCustomerGatewayOutput() CustomerGatewayOutput {
	return i.ToCustomerGatewayOutputWithContext(context.Background())
}

func (i *CustomerGateway) ToCustomerGatewayOutputWithContext(ctx context.Context) CustomerGatewayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CustomerGatewayOutput)
}

func (i *CustomerGateway) ToOutput(ctx context.Context) pulumix.Output[*CustomerGateway] {
	return pulumix.Output[*CustomerGateway]{
		OutputState: i.ToCustomerGatewayOutputWithContext(ctx).OutputState,
	}
}

// CustomerGatewayArrayInput is an input type that accepts CustomerGatewayArray and CustomerGatewayArrayOutput values.
// You can construct a concrete instance of `CustomerGatewayArrayInput` via:
//
//	CustomerGatewayArray{ CustomerGatewayArgs{...} }
type CustomerGatewayArrayInput interface {
	pulumi.Input

	ToCustomerGatewayArrayOutput() CustomerGatewayArrayOutput
	ToCustomerGatewayArrayOutputWithContext(context.Context) CustomerGatewayArrayOutput
}

type CustomerGatewayArray []CustomerGatewayInput

func (CustomerGatewayArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*CustomerGateway)(nil)).Elem()
}

func (i CustomerGatewayArray) ToCustomerGatewayArrayOutput() CustomerGatewayArrayOutput {
	return i.ToCustomerGatewayArrayOutputWithContext(context.Background())
}

func (i CustomerGatewayArray) ToCustomerGatewayArrayOutputWithContext(ctx context.Context) CustomerGatewayArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CustomerGatewayArrayOutput)
}

func (i CustomerGatewayArray) ToOutput(ctx context.Context) pulumix.Output[[]*CustomerGateway] {
	return pulumix.Output[[]*CustomerGateway]{
		OutputState: i.ToCustomerGatewayArrayOutputWithContext(ctx).OutputState,
	}
}

// CustomerGatewayMapInput is an input type that accepts CustomerGatewayMap and CustomerGatewayMapOutput values.
// You can construct a concrete instance of `CustomerGatewayMapInput` via:
//
//	CustomerGatewayMap{ "key": CustomerGatewayArgs{...} }
type CustomerGatewayMapInput interface {
	pulumi.Input

	ToCustomerGatewayMapOutput() CustomerGatewayMapOutput
	ToCustomerGatewayMapOutputWithContext(context.Context) CustomerGatewayMapOutput
}

type CustomerGatewayMap map[string]CustomerGatewayInput

func (CustomerGatewayMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*CustomerGateway)(nil)).Elem()
}

func (i CustomerGatewayMap) ToCustomerGatewayMapOutput() CustomerGatewayMapOutput {
	return i.ToCustomerGatewayMapOutputWithContext(context.Background())
}

func (i CustomerGatewayMap) ToCustomerGatewayMapOutputWithContext(ctx context.Context) CustomerGatewayMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CustomerGatewayMapOutput)
}

func (i CustomerGatewayMap) ToOutput(ctx context.Context) pulumix.Output[map[string]*CustomerGateway] {
	return pulumix.Output[map[string]*CustomerGateway]{
		OutputState: i.ToCustomerGatewayMapOutputWithContext(ctx).OutputState,
	}
}

type CustomerGatewayOutput struct{ *pulumi.OutputState }

func (CustomerGatewayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**CustomerGateway)(nil)).Elem()
}

func (o CustomerGatewayOutput) ToCustomerGatewayOutput() CustomerGatewayOutput {
	return o
}

func (o CustomerGatewayOutput) ToCustomerGatewayOutputWithContext(ctx context.Context) CustomerGatewayOutput {
	return o
}

func (o CustomerGatewayOutput) ToOutput(ctx context.Context) pulumix.Output[*CustomerGateway] {
	return pulumix.Output[*CustomerGateway]{
		OutputState: o.OutputState,
	}
}

// The ARN of the customer gateway.
func (o CustomerGatewayOutput) Arn() pulumi.StringOutput {
	return o.ApplyT(func(v *CustomerGateway) pulumi.StringOutput { return v.Arn }).(pulumi.StringOutput)
}

// The gateway's Border Gateway Protocol (BGP) Autonomous System Number (ASN).
func (o CustomerGatewayOutput) BgpAsn() pulumi.StringOutput {
	return o.ApplyT(func(v *CustomerGateway) pulumi.StringOutput { return v.BgpAsn }).(pulumi.StringOutput)
}

// The Amazon Resource Name (ARN) for the customer gateway certificate.
func (o CustomerGatewayOutput) CertificateArn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CustomerGateway) pulumi.StringPtrOutput { return v.CertificateArn }).(pulumi.StringPtrOutput)
}

// A name for the customer gateway device.
func (o CustomerGatewayOutput) DeviceName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CustomerGateway) pulumi.StringPtrOutput { return v.DeviceName }).(pulumi.StringPtrOutput)
}

// The IPv4 address for the customer gateway device's outside interface.
func (o CustomerGatewayOutput) IpAddress() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CustomerGateway) pulumi.StringPtrOutput { return v.IpAddress }).(pulumi.StringPtrOutput)
}

// Tags to apply to the gateway. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
func (o CustomerGatewayOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *CustomerGateway) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

// A map of tags assigned to the resource, including those inherited from the provider `defaultTags` configuration block.
func (o CustomerGatewayOutput) TagsAll() pulumi.StringMapOutput {
	return o.ApplyT(func(v *CustomerGateway) pulumi.StringMapOutput { return v.TagsAll }).(pulumi.StringMapOutput)
}

// The type of customer gateway. The only type AWS
// supports at this time is "ipsec.1".
func (o CustomerGatewayOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *CustomerGateway) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

type CustomerGatewayArrayOutput struct{ *pulumi.OutputState }

func (CustomerGatewayArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*CustomerGateway)(nil)).Elem()
}

func (o CustomerGatewayArrayOutput) ToCustomerGatewayArrayOutput() CustomerGatewayArrayOutput {
	return o
}

func (o CustomerGatewayArrayOutput) ToCustomerGatewayArrayOutputWithContext(ctx context.Context) CustomerGatewayArrayOutput {
	return o
}

func (o CustomerGatewayArrayOutput) ToOutput(ctx context.Context) pulumix.Output[[]*CustomerGateway] {
	return pulumix.Output[[]*CustomerGateway]{
		OutputState: o.OutputState,
	}
}

func (o CustomerGatewayArrayOutput) Index(i pulumi.IntInput) CustomerGatewayOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *CustomerGateway {
		return vs[0].([]*CustomerGateway)[vs[1].(int)]
	}).(CustomerGatewayOutput)
}

type CustomerGatewayMapOutput struct{ *pulumi.OutputState }

func (CustomerGatewayMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*CustomerGateway)(nil)).Elem()
}

func (o CustomerGatewayMapOutput) ToCustomerGatewayMapOutput() CustomerGatewayMapOutput {
	return o
}

func (o CustomerGatewayMapOutput) ToCustomerGatewayMapOutputWithContext(ctx context.Context) CustomerGatewayMapOutput {
	return o
}

func (o CustomerGatewayMapOutput) ToOutput(ctx context.Context) pulumix.Output[map[string]*CustomerGateway] {
	return pulumix.Output[map[string]*CustomerGateway]{
		OutputState: o.OutputState,
	}
}

func (o CustomerGatewayMapOutput) MapIndex(k pulumi.StringInput) CustomerGatewayOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *CustomerGateway {
		return vs[0].(map[string]*CustomerGateway)[vs[1].(string)]
	}).(CustomerGatewayOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*CustomerGatewayInput)(nil)).Elem(), &CustomerGateway{})
	pulumi.RegisterInputType(reflect.TypeOf((*CustomerGatewayArrayInput)(nil)).Elem(), CustomerGatewayArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*CustomerGatewayMapInput)(nil)).Elem(), CustomerGatewayMap{})
	pulumi.RegisterOutputType(CustomerGatewayOutput{})
	pulumi.RegisterOutputType(CustomerGatewayArrayOutput{})
	pulumi.RegisterOutputType(CustomerGatewayMapOutput{})
}
