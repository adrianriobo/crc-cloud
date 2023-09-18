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

// > **Note:** There is only a single subscription allowed per account.
//
// To help you understand the charges for your Spot instances, Amazon EC2 provides a data feed that describes your Spot instance usage and pricing.
// This data feed is sent to an Amazon S3 bucket that you specify when you subscribe to the data feed.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/ec2"
//	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/s3"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			defaultBucketV2, err := s3.NewBucketV2(ctx, "defaultBucketV2", nil)
//			if err != nil {
//				return err
//			}
//			_, err = ec2.NewSpotDatafeedSubscription(ctx, "defaultSpotDatafeedSubscription", &ec2.SpotDatafeedSubscriptionArgs{
//				Bucket: defaultBucketV2.ID(),
//				Prefix: pulumi.String("my_subdirectory"),
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
// Using `pulumi import`, import a Spot Datafeed Subscription using the word `spot-datafeed-subscription`. For example:
//
// ```sh
//
//	$ pulumi import aws:ec2/spotDatafeedSubscription:SpotDatafeedSubscription mysubscription spot-datafeed-subscription
//
// ```
type SpotDatafeedSubscription struct {
	pulumi.CustomResourceState

	// The Amazon S3 bucket in which to store the Spot instance data feed.
	Bucket pulumi.StringOutput `pulumi:"bucket"`
	// Path of folder inside bucket to place spot pricing data.
	Prefix pulumi.StringPtrOutput `pulumi:"prefix"`
}

// NewSpotDatafeedSubscription registers a new resource with the given unique name, arguments, and options.
func NewSpotDatafeedSubscription(ctx *pulumi.Context,
	name string, args *SpotDatafeedSubscriptionArgs, opts ...pulumi.ResourceOption) (*SpotDatafeedSubscription, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Bucket == nil {
		return nil, errors.New("invalid value for required argument 'Bucket'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource SpotDatafeedSubscription
	err := ctx.RegisterResource("aws:ec2/spotDatafeedSubscription:SpotDatafeedSubscription", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSpotDatafeedSubscription gets an existing SpotDatafeedSubscription resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSpotDatafeedSubscription(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SpotDatafeedSubscriptionState, opts ...pulumi.ResourceOption) (*SpotDatafeedSubscription, error) {
	var resource SpotDatafeedSubscription
	err := ctx.ReadResource("aws:ec2/spotDatafeedSubscription:SpotDatafeedSubscription", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SpotDatafeedSubscription resources.
type spotDatafeedSubscriptionState struct {
	// The Amazon S3 bucket in which to store the Spot instance data feed.
	Bucket *string `pulumi:"bucket"`
	// Path of folder inside bucket to place spot pricing data.
	Prefix *string `pulumi:"prefix"`
}

type SpotDatafeedSubscriptionState struct {
	// The Amazon S3 bucket in which to store the Spot instance data feed.
	Bucket pulumi.StringPtrInput
	// Path of folder inside bucket to place spot pricing data.
	Prefix pulumi.StringPtrInput
}

func (SpotDatafeedSubscriptionState) ElementType() reflect.Type {
	return reflect.TypeOf((*spotDatafeedSubscriptionState)(nil)).Elem()
}

type spotDatafeedSubscriptionArgs struct {
	// The Amazon S3 bucket in which to store the Spot instance data feed.
	Bucket string `pulumi:"bucket"`
	// Path of folder inside bucket to place spot pricing data.
	Prefix *string `pulumi:"prefix"`
}

// The set of arguments for constructing a SpotDatafeedSubscription resource.
type SpotDatafeedSubscriptionArgs struct {
	// The Amazon S3 bucket in which to store the Spot instance data feed.
	Bucket pulumi.StringInput
	// Path of folder inside bucket to place spot pricing data.
	Prefix pulumi.StringPtrInput
}

func (SpotDatafeedSubscriptionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*spotDatafeedSubscriptionArgs)(nil)).Elem()
}

type SpotDatafeedSubscriptionInput interface {
	pulumi.Input

	ToSpotDatafeedSubscriptionOutput() SpotDatafeedSubscriptionOutput
	ToSpotDatafeedSubscriptionOutputWithContext(ctx context.Context) SpotDatafeedSubscriptionOutput
}

func (*SpotDatafeedSubscription) ElementType() reflect.Type {
	return reflect.TypeOf((**SpotDatafeedSubscription)(nil)).Elem()
}

func (i *SpotDatafeedSubscription) ToSpotDatafeedSubscriptionOutput() SpotDatafeedSubscriptionOutput {
	return i.ToSpotDatafeedSubscriptionOutputWithContext(context.Background())
}

func (i *SpotDatafeedSubscription) ToSpotDatafeedSubscriptionOutputWithContext(ctx context.Context) SpotDatafeedSubscriptionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SpotDatafeedSubscriptionOutput)
}

func (i *SpotDatafeedSubscription) ToOutput(ctx context.Context) pulumix.Output[*SpotDatafeedSubscription] {
	return pulumix.Output[*SpotDatafeedSubscription]{
		OutputState: i.ToSpotDatafeedSubscriptionOutputWithContext(ctx).OutputState,
	}
}

// SpotDatafeedSubscriptionArrayInput is an input type that accepts SpotDatafeedSubscriptionArray and SpotDatafeedSubscriptionArrayOutput values.
// You can construct a concrete instance of `SpotDatafeedSubscriptionArrayInput` via:
//
//	SpotDatafeedSubscriptionArray{ SpotDatafeedSubscriptionArgs{...} }
type SpotDatafeedSubscriptionArrayInput interface {
	pulumi.Input

	ToSpotDatafeedSubscriptionArrayOutput() SpotDatafeedSubscriptionArrayOutput
	ToSpotDatafeedSubscriptionArrayOutputWithContext(context.Context) SpotDatafeedSubscriptionArrayOutput
}

type SpotDatafeedSubscriptionArray []SpotDatafeedSubscriptionInput

func (SpotDatafeedSubscriptionArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SpotDatafeedSubscription)(nil)).Elem()
}

func (i SpotDatafeedSubscriptionArray) ToSpotDatafeedSubscriptionArrayOutput() SpotDatafeedSubscriptionArrayOutput {
	return i.ToSpotDatafeedSubscriptionArrayOutputWithContext(context.Background())
}

func (i SpotDatafeedSubscriptionArray) ToSpotDatafeedSubscriptionArrayOutputWithContext(ctx context.Context) SpotDatafeedSubscriptionArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SpotDatafeedSubscriptionArrayOutput)
}

func (i SpotDatafeedSubscriptionArray) ToOutput(ctx context.Context) pulumix.Output[[]*SpotDatafeedSubscription] {
	return pulumix.Output[[]*SpotDatafeedSubscription]{
		OutputState: i.ToSpotDatafeedSubscriptionArrayOutputWithContext(ctx).OutputState,
	}
}

// SpotDatafeedSubscriptionMapInput is an input type that accepts SpotDatafeedSubscriptionMap and SpotDatafeedSubscriptionMapOutput values.
// You can construct a concrete instance of `SpotDatafeedSubscriptionMapInput` via:
//
//	SpotDatafeedSubscriptionMap{ "key": SpotDatafeedSubscriptionArgs{...} }
type SpotDatafeedSubscriptionMapInput interface {
	pulumi.Input

	ToSpotDatafeedSubscriptionMapOutput() SpotDatafeedSubscriptionMapOutput
	ToSpotDatafeedSubscriptionMapOutputWithContext(context.Context) SpotDatafeedSubscriptionMapOutput
}

type SpotDatafeedSubscriptionMap map[string]SpotDatafeedSubscriptionInput

func (SpotDatafeedSubscriptionMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SpotDatafeedSubscription)(nil)).Elem()
}

func (i SpotDatafeedSubscriptionMap) ToSpotDatafeedSubscriptionMapOutput() SpotDatafeedSubscriptionMapOutput {
	return i.ToSpotDatafeedSubscriptionMapOutputWithContext(context.Background())
}

func (i SpotDatafeedSubscriptionMap) ToSpotDatafeedSubscriptionMapOutputWithContext(ctx context.Context) SpotDatafeedSubscriptionMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SpotDatafeedSubscriptionMapOutput)
}

func (i SpotDatafeedSubscriptionMap) ToOutput(ctx context.Context) pulumix.Output[map[string]*SpotDatafeedSubscription] {
	return pulumix.Output[map[string]*SpotDatafeedSubscription]{
		OutputState: i.ToSpotDatafeedSubscriptionMapOutputWithContext(ctx).OutputState,
	}
}

type SpotDatafeedSubscriptionOutput struct{ *pulumi.OutputState }

func (SpotDatafeedSubscriptionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SpotDatafeedSubscription)(nil)).Elem()
}

func (o SpotDatafeedSubscriptionOutput) ToSpotDatafeedSubscriptionOutput() SpotDatafeedSubscriptionOutput {
	return o
}

func (o SpotDatafeedSubscriptionOutput) ToSpotDatafeedSubscriptionOutputWithContext(ctx context.Context) SpotDatafeedSubscriptionOutput {
	return o
}

func (o SpotDatafeedSubscriptionOutput) ToOutput(ctx context.Context) pulumix.Output[*SpotDatafeedSubscription] {
	return pulumix.Output[*SpotDatafeedSubscription]{
		OutputState: o.OutputState,
	}
}

// The Amazon S3 bucket in which to store the Spot instance data feed.
func (o SpotDatafeedSubscriptionOutput) Bucket() pulumi.StringOutput {
	return o.ApplyT(func(v *SpotDatafeedSubscription) pulumi.StringOutput { return v.Bucket }).(pulumi.StringOutput)
}

// Path of folder inside bucket to place spot pricing data.
func (o SpotDatafeedSubscriptionOutput) Prefix() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SpotDatafeedSubscription) pulumi.StringPtrOutput { return v.Prefix }).(pulumi.StringPtrOutput)
}

type SpotDatafeedSubscriptionArrayOutput struct{ *pulumi.OutputState }

func (SpotDatafeedSubscriptionArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SpotDatafeedSubscription)(nil)).Elem()
}

func (o SpotDatafeedSubscriptionArrayOutput) ToSpotDatafeedSubscriptionArrayOutput() SpotDatafeedSubscriptionArrayOutput {
	return o
}

func (o SpotDatafeedSubscriptionArrayOutput) ToSpotDatafeedSubscriptionArrayOutputWithContext(ctx context.Context) SpotDatafeedSubscriptionArrayOutput {
	return o
}

func (o SpotDatafeedSubscriptionArrayOutput) ToOutput(ctx context.Context) pulumix.Output[[]*SpotDatafeedSubscription] {
	return pulumix.Output[[]*SpotDatafeedSubscription]{
		OutputState: o.OutputState,
	}
}

func (o SpotDatafeedSubscriptionArrayOutput) Index(i pulumi.IntInput) SpotDatafeedSubscriptionOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *SpotDatafeedSubscription {
		return vs[0].([]*SpotDatafeedSubscription)[vs[1].(int)]
	}).(SpotDatafeedSubscriptionOutput)
}

type SpotDatafeedSubscriptionMapOutput struct{ *pulumi.OutputState }

func (SpotDatafeedSubscriptionMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SpotDatafeedSubscription)(nil)).Elem()
}

func (o SpotDatafeedSubscriptionMapOutput) ToSpotDatafeedSubscriptionMapOutput() SpotDatafeedSubscriptionMapOutput {
	return o
}

func (o SpotDatafeedSubscriptionMapOutput) ToSpotDatafeedSubscriptionMapOutputWithContext(ctx context.Context) SpotDatafeedSubscriptionMapOutput {
	return o
}

func (o SpotDatafeedSubscriptionMapOutput) ToOutput(ctx context.Context) pulumix.Output[map[string]*SpotDatafeedSubscription] {
	return pulumix.Output[map[string]*SpotDatafeedSubscription]{
		OutputState: o.OutputState,
	}
}

func (o SpotDatafeedSubscriptionMapOutput) MapIndex(k pulumi.StringInput) SpotDatafeedSubscriptionOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *SpotDatafeedSubscription {
		return vs[0].(map[string]*SpotDatafeedSubscription)[vs[1].(string)]
	}).(SpotDatafeedSubscriptionOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*SpotDatafeedSubscriptionInput)(nil)).Elem(), &SpotDatafeedSubscription{})
	pulumi.RegisterInputType(reflect.TypeOf((*SpotDatafeedSubscriptionArrayInput)(nil)).Elem(), SpotDatafeedSubscriptionArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*SpotDatafeedSubscriptionMapInput)(nil)).Elem(), SpotDatafeedSubscriptionMap{})
	pulumi.RegisterOutputType(SpotDatafeedSubscriptionOutput{})
	pulumi.RegisterOutputType(SpotDatafeedSubscriptionArrayOutput{})
	pulumi.RegisterOutputType(SpotDatafeedSubscriptionMapOutput{})
}
