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

// Provides a resource to manage VPC peering connection options.
//
// > **NOTE on VPC Peering Connections and VPC Peering Connection Options:** This provider provides
// both a standalone VPC Peering Connection Options and a VPC Peering Connection
// resource with `accepter` and `requester` attributes. Do not manage options for the same VPC peering
// connection in both a VPC Peering Connection resource and a VPC Peering Connection Options resource.
// Doing so will cause a conflict of options and will overwrite the options.
// Using a VPC Peering Connection Options resource decouples management of the connection options from
// management of the VPC Peering Connection and allows options to be set correctly in cross-region and
// cross-account scenarios.
//
// ## Example Usage
// ### Basic Usage
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
//			fooVpc, err := ec2.NewVpc(ctx, "fooVpc", &ec2.VpcArgs{
//				CidrBlock: pulumi.String("10.0.0.0/16"),
//			})
//			if err != nil {
//				return err
//			}
//			bar, err := ec2.NewVpc(ctx, "bar", &ec2.VpcArgs{
//				CidrBlock: pulumi.String("10.1.0.0/16"),
//			})
//			if err != nil {
//				return err
//			}
//			fooVpcPeeringConnection, err := ec2.NewVpcPeeringConnection(ctx, "fooVpcPeeringConnection", &ec2.VpcPeeringConnectionArgs{
//				VpcId:      fooVpc.ID(),
//				PeerVpcId:  bar.ID(),
//				AutoAccept: pulumi.Bool(true),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = ec2.NewPeeringConnectionOptions(ctx, "fooPeeringConnectionOptions", &ec2.PeeringConnectionOptionsArgs{
//				VpcPeeringConnectionId: fooVpcPeeringConnection.ID(),
//				Accepter: &ec2.PeeringConnectionOptionsAccepterArgs{
//					AllowRemoteVpcDnsResolution: pulumi.Bool(true),
//				},
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
// ### Cross-Account Usage
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
//			_, err := aws.NewProvider(ctx, "requester", nil)
//			if err != nil {
//				return err
//			}
//			_, err = aws.NewProvider(ctx, "accepter", nil)
//			if err != nil {
//				return err
//			}
//			main, err := ec2.NewVpc(ctx, "main", &ec2.VpcArgs{
//				CidrBlock:          pulumi.String("10.0.0.0/16"),
//				EnableDnsSupport:   pulumi.Bool(true),
//				EnableDnsHostnames: pulumi.Bool(true),
//			}, pulumi.Provider(aws.Requester))
//			if err != nil {
//				return err
//			}
//			peerVpc, err := ec2.NewVpc(ctx, "peerVpc", &ec2.VpcArgs{
//				CidrBlock:          pulumi.String("10.1.0.0/16"),
//				EnableDnsSupport:   pulumi.Bool(true),
//				EnableDnsHostnames: pulumi.Bool(true),
//			}, pulumi.Provider(aws.Accepter))
//			if err != nil {
//				return err
//			}
//			peerCallerIdentity, err := aws.GetCallerIdentity(ctx, nil, nil)
//			if err != nil {
//				return err
//			}
//			peerVpcPeeringConnection, err := ec2.NewVpcPeeringConnection(ctx, "peerVpcPeeringConnection", &ec2.VpcPeeringConnectionArgs{
//				VpcId:       main.ID(),
//				PeerVpcId:   peerVpc.ID(),
//				PeerOwnerId: *pulumi.String(peerCallerIdentity.AccountId),
//				AutoAccept:  pulumi.Bool(false),
//				Tags: pulumi.StringMap{
//					"Side": pulumi.String("Requester"),
//				},
//			}, pulumi.Provider(aws.Requester))
//			if err != nil {
//				return err
//			}
//			peerVpcPeeringConnectionAccepter, err := ec2.NewVpcPeeringConnectionAccepter(ctx, "peerVpcPeeringConnectionAccepter", &ec2.VpcPeeringConnectionAccepterArgs{
//				VpcPeeringConnectionId: peerVpcPeeringConnection.ID(),
//				AutoAccept:             pulumi.Bool(true),
//				Tags: pulumi.StringMap{
//					"Side": pulumi.String("Accepter"),
//				},
//			}, pulumi.Provider(aws.Accepter))
//			if err != nil {
//				return err
//			}
//			_, err = ec2.NewPeeringConnectionOptions(ctx, "requesterPeeringConnectionOptions", &ec2.PeeringConnectionOptionsArgs{
//				VpcPeeringConnectionId: peerVpcPeeringConnectionAccepter.ID(),
//				Requester: &ec2.PeeringConnectionOptionsRequesterArgs{
//					AllowRemoteVpcDnsResolution: pulumi.Bool(true),
//				},
//			}, pulumi.Provider(aws.Requester))
//			if err != nil {
//				return err
//			}
//			_, err = ec2.NewPeeringConnectionOptions(ctx, "accepterPeeringConnectionOptions", &ec2.PeeringConnectionOptionsArgs{
//				VpcPeeringConnectionId: peerVpcPeeringConnectionAccepter.ID(),
//				Accepter: &ec2.PeeringConnectionOptionsAccepterArgs{
//					AllowRemoteVpcDnsResolution: pulumi.Bool(true),
//				},
//			}, pulumi.Provider(aws.Accepter))
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
// Using `pulumi import`, import VPC Peering Connection Options using the VPC peering `id`. For example:
//
// ```sh
//
//	$ pulumi import aws:ec2/peeringConnectionOptions:PeeringConnectionOptions foo pcx-111aaa111
//
// ```
type PeeringConnectionOptions struct {
	pulumi.CustomResourceState

	// An optional configuration block that allows for [VPC Peering Connection](https://docs.aws.amazon.com/vpc/latest/peering/what-is-vpc-peering.html) options to be set for the VPC that acceptsthe peering connection (a maximum of one).
	Accepter PeeringConnectionOptionsAccepterOutput `pulumi:"accepter"`
	// A optional configuration block that allows for [VPC Peering Connection](https://docs.aws.amazon.com/vpc/latest/peering/what-is-vpc-peering.html) options to be set for the VPC that requeststhe peering connection (a maximum of one).
	Requester PeeringConnectionOptionsRequesterOutput `pulumi:"requester"`
	// The ID of the requester VPC peering connection.
	VpcPeeringConnectionId pulumi.StringOutput `pulumi:"vpcPeeringConnectionId"`
}

// NewPeeringConnectionOptions registers a new resource with the given unique name, arguments, and options.
func NewPeeringConnectionOptions(ctx *pulumi.Context,
	name string, args *PeeringConnectionOptionsArgs, opts ...pulumi.ResourceOption) (*PeeringConnectionOptions, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.VpcPeeringConnectionId == nil {
		return nil, errors.New("invalid value for required argument 'VpcPeeringConnectionId'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource PeeringConnectionOptions
	err := ctx.RegisterResource("aws:ec2/peeringConnectionOptions:PeeringConnectionOptions", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetPeeringConnectionOptions gets an existing PeeringConnectionOptions resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetPeeringConnectionOptions(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *PeeringConnectionOptionsState, opts ...pulumi.ResourceOption) (*PeeringConnectionOptions, error) {
	var resource PeeringConnectionOptions
	err := ctx.ReadResource("aws:ec2/peeringConnectionOptions:PeeringConnectionOptions", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering PeeringConnectionOptions resources.
type peeringConnectionOptionsState struct {
	// An optional configuration block that allows for [VPC Peering Connection](https://docs.aws.amazon.com/vpc/latest/peering/what-is-vpc-peering.html) options to be set for the VPC that acceptsthe peering connection (a maximum of one).
	Accepter *PeeringConnectionOptionsAccepter `pulumi:"accepter"`
	// A optional configuration block that allows for [VPC Peering Connection](https://docs.aws.amazon.com/vpc/latest/peering/what-is-vpc-peering.html) options to be set for the VPC that requeststhe peering connection (a maximum of one).
	Requester *PeeringConnectionOptionsRequester `pulumi:"requester"`
	// The ID of the requester VPC peering connection.
	VpcPeeringConnectionId *string `pulumi:"vpcPeeringConnectionId"`
}

type PeeringConnectionOptionsState struct {
	// An optional configuration block that allows for [VPC Peering Connection](https://docs.aws.amazon.com/vpc/latest/peering/what-is-vpc-peering.html) options to be set for the VPC that acceptsthe peering connection (a maximum of one).
	Accepter PeeringConnectionOptionsAccepterPtrInput
	// A optional configuration block that allows for [VPC Peering Connection](https://docs.aws.amazon.com/vpc/latest/peering/what-is-vpc-peering.html) options to be set for the VPC that requeststhe peering connection (a maximum of one).
	Requester PeeringConnectionOptionsRequesterPtrInput
	// The ID of the requester VPC peering connection.
	VpcPeeringConnectionId pulumi.StringPtrInput
}

func (PeeringConnectionOptionsState) ElementType() reflect.Type {
	return reflect.TypeOf((*peeringConnectionOptionsState)(nil)).Elem()
}

type peeringConnectionOptionsArgs struct {
	// An optional configuration block that allows for [VPC Peering Connection](https://docs.aws.amazon.com/vpc/latest/peering/what-is-vpc-peering.html) options to be set for the VPC that acceptsthe peering connection (a maximum of one).
	Accepter *PeeringConnectionOptionsAccepter `pulumi:"accepter"`
	// A optional configuration block that allows for [VPC Peering Connection](https://docs.aws.amazon.com/vpc/latest/peering/what-is-vpc-peering.html) options to be set for the VPC that requeststhe peering connection (a maximum of one).
	Requester *PeeringConnectionOptionsRequester `pulumi:"requester"`
	// The ID of the requester VPC peering connection.
	VpcPeeringConnectionId string `pulumi:"vpcPeeringConnectionId"`
}

// The set of arguments for constructing a PeeringConnectionOptions resource.
type PeeringConnectionOptionsArgs struct {
	// An optional configuration block that allows for [VPC Peering Connection](https://docs.aws.amazon.com/vpc/latest/peering/what-is-vpc-peering.html) options to be set for the VPC that acceptsthe peering connection (a maximum of one).
	Accepter PeeringConnectionOptionsAccepterPtrInput
	// A optional configuration block that allows for [VPC Peering Connection](https://docs.aws.amazon.com/vpc/latest/peering/what-is-vpc-peering.html) options to be set for the VPC that requeststhe peering connection (a maximum of one).
	Requester PeeringConnectionOptionsRequesterPtrInput
	// The ID of the requester VPC peering connection.
	VpcPeeringConnectionId pulumi.StringInput
}

func (PeeringConnectionOptionsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*peeringConnectionOptionsArgs)(nil)).Elem()
}

type PeeringConnectionOptionsInput interface {
	pulumi.Input

	ToPeeringConnectionOptionsOutput() PeeringConnectionOptionsOutput
	ToPeeringConnectionOptionsOutputWithContext(ctx context.Context) PeeringConnectionOptionsOutput
}

func (*PeeringConnectionOptions) ElementType() reflect.Type {
	return reflect.TypeOf((**PeeringConnectionOptions)(nil)).Elem()
}

func (i *PeeringConnectionOptions) ToPeeringConnectionOptionsOutput() PeeringConnectionOptionsOutput {
	return i.ToPeeringConnectionOptionsOutputWithContext(context.Background())
}

func (i *PeeringConnectionOptions) ToPeeringConnectionOptionsOutputWithContext(ctx context.Context) PeeringConnectionOptionsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PeeringConnectionOptionsOutput)
}

func (i *PeeringConnectionOptions) ToOutput(ctx context.Context) pulumix.Output[*PeeringConnectionOptions] {
	return pulumix.Output[*PeeringConnectionOptions]{
		OutputState: i.ToPeeringConnectionOptionsOutputWithContext(ctx).OutputState,
	}
}

// PeeringConnectionOptionsArrayInput is an input type that accepts PeeringConnectionOptionsArray and PeeringConnectionOptionsArrayOutput values.
// You can construct a concrete instance of `PeeringConnectionOptionsArrayInput` via:
//
//	PeeringConnectionOptionsArray{ PeeringConnectionOptionsArgs{...} }
type PeeringConnectionOptionsArrayInput interface {
	pulumi.Input

	ToPeeringConnectionOptionsArrayOutput() PeeringConnectionOptionsArrayOutput
	ToPeeringConnectionOptionsArrayOutputWithContext(context.Context) PeeringConnectionOptionsArrayOutput
}

type PeeringConnectionOptionsArray []PeeringConnectionOptionsInput

func (PeeringConnectionOptionsArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*PeeringConnectionOptions)(nil)).Elem()
}

func (i PeeringConnectionOptionsArray) ToPeeringConnectionOptionsArrayOutput() PeeringConnectionOptionsArrayOutput {
	return i.ToPeeringConnectionOptionsArrayOutputWithContext(context.Background())
}

func (i PeeringConnectionOptionsArray) ToPeeringConnectionOptionsArrayOutputWithContext(ctx context.Context) PeeringConnectionOptionsArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PeeringConnectionOptionsArrayOutput)
}

func (i PeeringConnectionOptionsArray) ToOutput(ctx context.Context) pulumix.Output[[]*PeeringConnectionOptions] {
	return pulumix.Output[[]*PeeringConnectionOptions]{
		OutputState: i.ToPeeringConnectionOptionsArrayOutputWithContext(ctx).OutputState,
	}
}

// PeeringConnectionOptionsMapInput is an input type that accepts PeeringConnectionOptionsMap and PeeringConnectionOptionsMapOutput values.
// You can construct a concrete instance of `PeeringConnectionOptionsMapInput` via:
//
//	PeeringConnectionOptionsMap{ "key": PeeringConnectionOptionsArgs{...} }
type PeeringConnectionOptionsMapInput interface {
	pulumi.Input

	ToPeeringConnectionOptionsMapOutput() PeeringConnectionOptionsMapOutput
	ToPeeringConnectionOptionsMapOutputWithContext(context.Context) PeeringConnectionOptionsMapOutput
}

type PeeringConnectionOptionsMap map[string]PeeringConnectionOptionsInput

func (PeeringConnectionOptionsMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*PeeringConnectionOptions)(nil)).Elem()
}

func (i PeeringConnectionOptionsMap) ToPeeringConnectionOptionsMapOutput() PeeringConnectionOptionsMapOutput {
	return i.ToPeeringConnectionOptionsMapOutputWithContext(context.Background())
}

func (i PeeringConnectionOptionsMap) ToPeeringConnectionOptionsMapOutputWithContext(ctx context.Context) PeeringConnectionOptionsMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PeeringConnectionOptionsMapOutput)
}

func (i PeeringConnectionOptionsMap) ToOutput(ctx context.Context) pulumix.Output[map[string]*PeeringConnectionOptions] {
	return pulumix.Output[map[string]*PeeringConnectionOptions]{
		OutputState: i.ToPeeringConnectionOptionsMapOutputWithContext(ctx).OutputState,
	}
}

type PeeringConnectionOptionsOutput struct{ *pulumi.OutputState }

func (PeeringConnectionOptionsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PeeringConnectionOptions)(nil)).Elem()
}

func (o PeeringConnectionOptionsOutput) ToPeeringConnectionOptionsOutput() PeeringConnectionOptionsOutput {
	return o
}

func (o PeeringConnectionOptionsOutput) ToPeeringConnectionOptionsOutputWithContext(ctx context.Context) PeeringConnectionOptionsOutput {
	return o
}

func (o PeeringConnectionOptionsOutput) ToOutput(ctx context.Context) pulumix.Output[*PeeringConnectionOptions] {
	return pulumix.Output[*PeeringConnectionOptions]{
		OutputState: o.OutputState,
	}
}

// An optional configuration block that allows for [VPC Peering Connection](https://docs.aws.amazon.com/vpc/latest/peering/what-is-vpc-peering.html) options to be set for the VPC that acceptsthe peering connection (a maximum of one).
func (o PeeringConnectionOptionsOutput) Accepter() PeeringConnectionOptionsAccepterOutput {
	return o.ApplyT(func(v *PeeringConnectionOptions) PeeringConnectionOptionsAccepterOutput { return v.Accepter }).(PeeringConnectionOptionsAccepterOutput)
}

// A optional configuration block that allows for [VPC Peering Connection](https://docs.aws.amazon.com/vpc/latest/peering/what-is-vpc-peering.html) options to be set for the VPC that requeststhe peering connection (a maximum of one).
func (o PeeringConnectionOptionsOutput) Requester() PeeringConnectionOptionsRequesterOutput {
	return o.ApplyT(func(v *PeeringConnectionOptions) PeeringConnectionOptionsRequesterOutput { return v.Requester }).(PeeringConnectionOptionsRequesterOutput)
}

// The ID of the requester VPC peering connection.
func (o PeeringConnectionOptionsOutput) VpcPeeringConnectionId() pulumi.StringOutput {
	return o.ApplyT(func(v *PeeringConnectionOptions) pulumi.StringOutput { return v.VpcPeeringConnectionId }).(pulumi.StringOutput)
}

type PeeringConnectionOptionsArrayOutput struct{ *pulumi.OutputState }

func (PeeringConnectionOptionsArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*PeeringConnectionOptions)(nil)).Elem()
}

func (o PeeringConnectionOptionsArrayOutput) ToPeeringConnectionOptionsArrayOutput() PeeringConnectionOptionsArrayOutput {
	return o
}

func (o PeeringConnectionOptionsArrayOutput) ToPeeringConnectionOptionsArrayOutputWithContext(ctx context.Context) PeeringConnectionOptionsArrayOutput {
	return o
}

func (o PeeringConnectionOptionsArrayOutput) ToOutput(ctx context.Context) pulumix.Output[[]*PeeringConnectionOptions] {
	return pulumix.Output[[]*PeeringConnectionOptions]{
		OutputState: o.OutputState,
	}
}

func (o PeeringConnectionOptionsArrayOutput) Index(i pulumi.IntInput) PeeringConnectionOptionsOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *PeeringConnectionOptions {
		return vs[0].([]*PeeringConnectionOptions)[vs[1].(int)]
	}).(PeeringConnectionOptionsOutput)
}

type PeeringConnectionOptionsMapOutput struct{ *pulumi.OutputState }

func (PeeringConnectionOptionsMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*PeeringConnectionOptions)(nil)).Elem()
}

func (o PeeringConnectionOptionsMapOutput) ToPeeringConnectionOptionsMapOutput() PeeringConnectionOptionsMapOutput {
	return o
}

func (o PeeringConnectionOptionsMapOutput) ToPeeringConnectionOptionsMapOutputWithContext(ctx context.Context) PeeringConnectionOptionsMapOutput {
	return o
}

func (o PeeringConnectionOptionsMapOutput) ToOutput(ctx context.Context) pulumix.Output[map[string]*PeeringConnectionOptions] {
	return pulumix.Output[map[string]*PeeringConnectionOptions]{
		OutputState: o.OutputState,
	}
}

func (o PeeringConnectionOptionsMapOutput) MapIndex(k pulumi.StringInput) PeeringConnectionOptionsOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *PeeringConnectionOptions {
		return vs[0].(map[string]*PeeringConnectionOptions)[vs[1].(string)]
	}).(PeeringConnectionOptionsOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*PeeringConnectionOptionsInput)(nil)).Elem(), &PeeringConnectionOptions{})
	pulumi.RegisterInputType(reflect.TypeOf((*PeeringConnectionOptionsArrayInput)(nil)).Elem(), PeeringConnectionOptionsArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*PeeringConnectionOptionsMapInput)(nil)).Elem(), PeeringConnectionOptionsMap{})
	pulumi.RegisterOutputType(PeeringConnectionOptionsOutput{})
	pulumi.RegisterOutputType(PeeringConnectionOptionsArrayOutput{})
	pulumi.RegisterOutputType(PeeringConnectionOptionsMapOutput{})
}
