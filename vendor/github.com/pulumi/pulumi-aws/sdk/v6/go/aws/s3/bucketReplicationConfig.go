// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package s3

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Provides an independent configuration resource for S3 bucket [replication configuration](http://docs.aws.amazon.com/AmazonS3/latest/dev/crr.html).
//
// > **NOTE:** S3 Buckets only support a single replication configuration. Declaring multiple `s3.BucketReplicationConfig` resources to the same S3 Bucket will cause a perpetual difference in configuration.
//
// ## Example Usage
// ### Using replication configuration
//
// ```go
// package main
//
// import (
//
//	"fmt"
//
//	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws"
//	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/iam"
//	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/s3"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := aws.NewProvider(ctx, "central", &aws.ProviderArgs{
//				Region: pulumi.String("eu-central-1"),
//			})
//			if err != nil {
//				return err
//			}
//			assumeRole, err := iam.GetPolicyDocument(ctx, &iam.GetPolicyDocumentArgs{
//				Statements: []iam.GetPolicyDocumentStatement{
//					{
//						Effect: pulumi.StringRef("Allow"),
//						Principals: []iam.GetPolicyDocumentStatementPrincipal{
//							{
//								Type: "Service",
//								Identifiers: []string{
//									"s3.amazonaws.com",
//								},
//							},
//						},
//						Actions: []string{
//							"sts:AssumeRole",
//						},
//					},
//				},
//			}, nil)
//			if err != nil {
//				return err
//			}
//			replicationRole, err := iam.NewRole(ctx, "replicationRole", &iam.RoleArgs{
//				AssumeRolePolicy: *pulumi.String(assumeRole.Json),
//			})
//			if err != nil {
//				return err
//			}
//			destinationBucketV2, err := s3.NewBucketV2(ctx, "destinationBucketV2", nil)
//			if err != nil {
//				return err
//			}
//			sourceBucketV2, err := s3.NewBucketV2(ctx, "sourceBucketV2", nil, pulumi.Provider(aws.Central))
//			if err != nil {
//				return err
//			}
//			replicationPolicyDocument := iam.GetPolicyDocumentOutput(ctx, iam.GetPolicyDocumentOutputArgs{
//				Statements: iam.GetPolicyDocumentStatementArray{
//					&iam.GetPolicyDocumentStatementArgs{
//						Effect: pulumi.String("Allow"),
//						Actions: pulumi.StringArray{
//							pulumi.String("s3:GetReplicationConfiguration"),
//							pulumi.String("s3:ListBucket"),
//						},
//						Resources: pulumi.StringArray{
//							sourceBucketV2.Arn,
//						},
//					},
//					&iam.GetPolicyDocumentStatementArgs{
//						Effect: pulumi.String("Allow"),
//						Actions: pulumi.StringArray{
//							pulumi.String("s3:GetObjectVersionForReplication"),
//							pulumi.String("s3:GetObjectVersionAcl"),
//							pulumi.String("s3:GetObjectVersionTagging"),
//						},
//						Resources: pulumi.StringArray{
//							sourceBucketV2.Arn.ApplyT(func(arn string) (string, error) {
//								return fmt.Sprintf("%v/*", arn), nil
//							}).(pulumi.StringOutput),
//						},
//					},
//					&iam.GetPolicyDocumentStatementArgs{
//						Effect: pulumi.String("Allow"),
//						Actions: pulumi.StringArray{
//							pulumi.String("s3:ReplicateObject"),
//							pulumi.String("s3:ReplicateDelete"),
//							pulumi.String("s3:ReplicateTags"),
//						},
//						Resources: pulumi.StringArray{
//							destinationBucketV2.Arn.ApplyT(func(arn string) (string, error) {
//								return fmt.Sprintf("%v/*", arn), nil
//							}).(pulumi.StringOutput),
//						},
//					},
//				},
//			}, nil)
//			replicationPolicy, err := iam.NewPolicy(ctx, "replicationPolicy", &iam.PolicyArgs{
//				Policy: replicationPolicyDocument.ApplyT(func(replicationPolicyDocument iam.GetPolicyDocumentResult) (*string, error) {
//					return &replicationPolicyDocument.Json, nil
//				}).(pulumi.StringPtrOutput),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = iam.NewRolePolicyAttachment(ctx, "replicationRolePolicyAttachment", &iam.RolePolicyAttachmentArgs{
//				Role:      replicationRole.Name,
//				PolicyArn: replicationPolicy.Arn,
//			})
//			if err != nil {
//				return err
//			}
//			_, err = s3.NewBucketVersioningV2(ctx, "destinationBucketVersioningV2", &s3.BucketVersioningV2Args{
//				Bucket: destinationBucketV2.ID(),
//				VersioningConfiguration: &s3.BucketVersioningV2VersioningConfigurationArgs{
//					Status: pulumi.String("Enabled"),
//				},
//			})
//			if err != nil {
//				return err
//			}
//			_, err = s3.NewBucketAclV2(ctx, "sourceBucketAcl", &s3.BucketAclV2Args{
//				Bucket: sourceBucketV2.ID(),
//				Acl:    pulumi.String("private"),
//			}, pulumi.Provider(aws.Central))
//			if err != nil {
//				return err
//			}
//			sourceBucketVersioningV2, err := s3.NewBucketVersioningV2(ctx, "sourceBucketVersioningV2", &s3.BucketVersioningV2Args{
//				Bucket: sourceBucketV2.ID(),
//				VersioningConfiguration: &s3.BucketVersioningV2VersioningConfigurationArgs{
//					Status: pulumi.String("Enabled"),
//				},
//			}, pulumi.Provider(aws.Central))
//			if err != nil {
//				return err
//			}
//			_, err = s3.NewBucketReplicationConfig(ctx, "replicationBucketReplicationConfig", &s3.BucketReplicationConfigArgs{
//				Role:   replicationRole.Arn,
//				Bucket: sourceBucketV2.ID(),
//				Rules: s3.BucketReplicationConfigRuleArray{
//					&s3.BucketReplicationConfigRuleArgs{
//						Id: pulumi.String("foobar"),
//						Filter: &s3.BucketReplicationConfigRuleFilterArgs{
//							Prefix: pulumi.String("foo"),
//						},
//						Status: pulumi.String("Enabled"),
//						Destination: &s3.BucketReplicationConfigRuleDestinationArgs{
//							Bucket:       destinationBucketV2.Arn,
//							StorageClass: pulumi.String("STANDARD"),
//						},
//					},
//				},
//			}, pulumi.Provider(aws.Central), pulumi.DependsOn([]pulumi.Resource{
//				sourceBucketVersioningV2,
//			}))
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
// ### Bi-Directional Replication
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/s3"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			eastBucketV2, err := s3.NewBucketV2(ctx, "eastBucketV2", nil)
//			if err != nil {
//				return err
//			}
//			eastBucketVersioningV2, err := s3.NewBucketVersioningV2(ctx, "eastBucketVersioningV2", &s3.BucketVersioningV2Args{
//				Bucket: eastBucketV2.ID(),
//				VersioningConfiguration: &s3.BucketVersioningV2VersioningConfigurationArgs{
//					Status: pulumi.String("Enabled"),
//				},
//			})
//			if err != nil {
//				return err
//			}
//			westBucketV2, err := s3.NewBucketV2(ctx, "westBucketV2", nil, pulumi.Provider(aws.West))
//			if err != nil {
//				return err
//			}
//			westBucketVersioningV2, err := s3.NewBucketVersioningV2(ctx, "westBucketVersioningV2", &s3.BucketVersioningV2Args{
//				Bucket: westBucketV2.ID(),
//				VersioningConfiguration: &s3.BucketVersioningV2VersioningConfigurationArgs{
//					Status: pulumi.String("Enabled"),
//				},
//			}, pulumi.Provider(aws.West))
//			if err != nil {
//				return err
//			}
//			_, err = s3.NewBucketReplicationConfig(ctx, "eastToWest", &s3.BucketReplicationConfigArgs{
//				Role:   pulumi.Any(aws_iam_role.East_replication.Arn),
//				Bucket: eastBucketV2.ID(),
//				Rules: s3.BucketReplicationConfigRuleArray{
//					&s3.BucketReplicationConfigRuleArgs{
//						Id: pulumi.String("foobar"),
//						Filter: &s3.BucketReplicationConfigRuleFilterArgs{
//							Prefix: pulumi.String("foo"),
//						},
//						Status: pulumi.String("Enabled"),
//						Destination: &s3.BucketReplicationConfigRuleDestinationArgs{
//							Bucket:       westBucketV2.Arn,
//							StorageClass: pulumi.String("STANDARD"),
//						},
//					},
//				},
//			}, pulumi.DependsOn([]pulumi.Resource{
//				eastBucketVersioningV2,
//			}))
//			if err != nil {
//				return err
//			}
//			_, err = s3.NewBucketReplicationConfig(ctx, "westToEast", &s3.BucketReplicationConfigArgs{
//				Role:   pulumi.Any(aws_iam_role.West_replication.Arn),
//				Bucket: westBucketV2.ID(),
//				Rules: s3.BucketReplicationConfigRuleArray{
//					&s3.BucketReplicationConfigRuleArgs{
//						Id: pulumi.String("foobar"),
//						Filter: &s3.BucketReplicationConfigRuleFilterArgs{
//							Prefix: pulumi.String("foo"),
//						},
//						Status: pulumi.String("Enabled"),
//						Destination: &s3.BucketReplicationConfigRuleDestinationArgs{
//							Bucket:       eastBucketV2.Arn,
//							StorageClass: pulumi.String("STANDARD"),
//						},
//					},
//				},
//			}, pulumi.Provider(aws.West), pulumi.DependsOn([]pulumi.Resource{
//				westBucketVersioningV2,
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
// Using `pulumi import`, import S3 bucket replication configuration using the `bucket`. For example:
//
// ```sh
//
//	$ pulumi import aws:s3/bucketReplicationConfig:BucketReplicationConfig replication bucket-name
//
// ```
type BucketReplicationConfig struct {
	pulumi.CustomResourceState

	// Name of the source S3 bucket you want Amazon S3 to monitor.
	Bucket pulumi.StringOutput `pulumi:"bucket"`
	// ARN of the IAM role for Amazon S3 to assume when replicating the objects.
	Role pulumi.StringOutput `pulumi:"role"`
	// List of configuration blocks describing the rules managing the replication. See below.
	Rules BucketReplicationConfigRuleArrayOutput `pulumi:"rules"`
	// Token to allow replication to be enabled on an Object Lock-enabled bucket. You must contact AWS support for the bucket's "Object Lock token".
	// For more details, see [Using S3 Object Lock with replication](https://docs.aws.amazon.com/AmazonS3/latest/userguide/object-lock-managing.html#object-lock-managing-replication).
	Token pulumi.StringPtrOutput `pulumi:"token"`
}

// NewBucketReplicationConfig registers a new resource with the given unique name, arguments, and options.
func NewBucketReplicationConfig(ctx *pulumi.Context,
	name string, args *BucketReplicationConfigArgs, opts ...pulumi.ResourceOption) (*BucketReplicationConfig, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Bucket == nil {
		return nil, errors.New("invalid value for required argument 'Bucket'")
	}
	if args.Role == nil {
		return nil, errors.New("invalid value for required argument 'Role'")
	}
	if args.Rules == nil {
		return nil, errors.New("invalid value for required argument 'Rules'")
	}
	if args.Token != nil {
		args.Token = pulumi.ToSecret(args.Token).(pulumi.StringPtrInput)
	}
	secrets := pulumi.AdditionalSecretOutputs([]string{
		"token",
	})
	opts = append(opts, secrets)
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource BucketReplicationConfig
	err := ctx.RegisterResource("aws:s3/bucketReplicationConfig:BucketReplicationConfig", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetBucketReplicationConfig gets an existing BucketReplicationConfig resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetBucketReplicationConfig(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *BucketReplicationConfigState, opts ...pulumi.ResourceOption) (*BucketReplicationConfig, error) {
	var resource BucketReplicationConfig
	err := ctx.ReadResource("aws:s3/bucketReplicationConfig:BucketReplicationConfig", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering BucketReplicationConfig resources.
type bucketReplicationConfigState struct {
	// Name of the source S3 bucket you want Amazon S3 to monitor.
	Bucket *string `pulumi:"bucket"`
	// ARN of the IAM role for Amazon S3 to assume when replicating the objects.
	Role *string `pulumi:"role"`
	// List of configuration blocks describing the rules managing the replication. See below.
	Rules []BucketReplicationConfigRule `pulumi:"rules"`
	// Token to allow replication to be enabled on an Object Lock-enabled bucket. You must contact AWS support for the bucket's "Object Lock token".
	// For more details, see [Using S3 Object Lock with replication](https://docs.aws.amazon.com/AmazonS3/latest/userguide/object-lock-managing.html#object-lock-managing-replication).
	Token *string `pulumi:"token"`
}

type BucketReplicationConfigState struct {
	// Name of the source S3 bucket you want Amazon S3 to monitor.
	Bucket pulumi.StringPtrInput
	// ARN of the IAM role for Amazon S3 to assume when replicating the objects.
	Role pulumi.StringPtrInput
	// List of configuration blocks describing the rules managing the replication. See below.
	Rules BucketReplicationConfigRuleArrayInput
	// Token to allow replication to be enabled on an Object Lock-enabled bucket. You must contact AWS support for the bucket's "Object Lock token".
	// For more details, see [Using S3 Object Lock with replication](https://docs.aws.amazon.com/AmazonS3/latest/userguide/object-lock-managing.html#object-lock-managing-replication).
	Token pulumi.StringPtrInput
}

func (BucketReplicationConfigState) ElementType() reflect.Type {
	return reflect.TypeOf((*bucketReplicationConfigState)(nil)).Elem()
}

type bucketReplicationConfigArgs struct {
	// Name of the source S3 bucket you want Amazon S3 to monitor.
	Bucket string `pulumi:"bucket"`
	// ARN of the IAM role for Amazon S3 to assume when replicating the objects.
	Role string `pulumi:"role"`
	// List of configuration blocks describing the rules managing the replication. See below.
	Rules []BucketReplicationConfigRule `pulumi:"rules"`
	// Token to allow replication to be enabled on an Object Lock-enabled bucket. You must contact AWS support for the bucket's "Object Lock token".
	// For more details, see [Using S3 Object Lock with replication](https://docs.aws.amazon.com/AmazonS3/latest/userguide/object-lock-managing.html#object-lock-managing-replication).
	Token *string `pulumi:"token"`
}

// The set of arguments for constructing a BucketReplicationConfig resource.
type BucketReplicationConfigArgs struct {
	// Name of the source S3 bucket you want Amazon S3 to monitor.
	Bucket pulumi.StringInput
	// ARN of the IAM role for Amazon S3 to assume when replicating the objects.
	Role pulumi.StringInput
	// List of configuration blocks describing the rules managing the replication. See below.
	Rules BucketReplicationConfigRuleArrayInput
	// Token to allow replication to be enabled on an Object Lock-enabled bucket. You must contact AWS support for the bucket's "Object Lock token".
	// For more details, see [Using S3 Object Lock with replication](https://docs.aws.amazon.com/AmazonS3/latest/userguide/object-lock-managing.html#object-lock-managing-replication).
	Token pulumi.StringPtrInput
}

func (BucketReplicationConfigArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*bucketReplicationConfigArgs)(nil)).Elem()
}

type BucketReplicationConfigInput interface {
	pulumi.Input

	ToBucketReplicationConfigOutput() BucketReplicationConfigOutput
	ToBucketReplicationConfigOutputWithContext(ctx context.Context) BucketReplicationConfigOutput
}

func (*BucketReplicationConfig) ElementType() reflect.Type {
	return reflect.TypeOf((**BucketReplicationConfig)(nil)).Elem()
}

func (i *BucketReplicationConfig) ToBucketReplicationConfigOutput() BucketReplicationConfigOutput {
	return i.ToBucketReplicationConfigOutputWithContext(context.Background())
}

func (i *BucketReplicationConfig) ToBucketReplicationConfigOutputWithContext(ctx context.Context) BucketReplicationConfigOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BucketReplicationConfigOutput)
}

func (i *BucketReplicationConfig) ToOutput(ctx context.Context) pulumix.Output[*BucketReplicationConfig] {
	return pulumix.Output[*BucketReplicationConfig]{
		OutputState: i.ToBucketReplicationConfigOutputWithContext(ctx).OutputState,
	}
}

// BucketReplicationConfigArrayInput is an input type that accepts BucketReplicationConfigArray and BucketReplicationConfigArrayOutput values.
// You can construct a concrete instance of `BucketReplicationConfigArrayInput` via:
//
//	BucketReplicationConfigArray{ BucketReplicationConfigArgs{...} }
type BucketReplicationConfigArrayInput interface {
	pulumi.Input

	ToBucketReplicationConfigArrayOutput() BucketReplicationConfigArrayOutput
	ToBucketReplicationConfigArrayOutputWithContext(context.Context) BucketReplicationConfigArrayOutput
}

type BucketReplicationConfigArray []BucketReplicationConfigInput

func (BucketReplicationConfigArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*BucketReplicationConfig)(nil)).Elem()
}

func (i BucketReplicationConfigArray) ToBucketReplicationConfigArrayOutput() BucketReplicationConfigArrayOutput {
	return i.ToBucketReplicationConfigArrayOutputWithContext(context.Background())
}

func (i BucketReplicationConfigArray) ToBucketReplicationConfigArrayOutputWithContext(ctx context.Context) BucketReplicationConfigArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BucketReplicationConfigArrayOutput)
}

func (i BucketReplicationConfigArray) ToOutput(ctx context.Context) pulumix.Output[[]*BucketReplicationConfig] {
	return pulumix.Output[[]*BucketReplicationConfig]{
		OutputState: i.ToBucketReplicationConfigArrayOutputWithContext(ctx).OutputState,
	}
}

// BucketReplicationConfigMapInput is an input type that accepts BucketReplicationConfigMap and BucketReplicationConfigMapOutput values.
// You can construct a concrete instance of `BucketReplicationConfigMapInput` via:
//
//	BucketReplicationConfigMap{ "key": BucketReplicationConfigArgs{...} }
type BucketReplicationConfigMapInput interface {
	pulumi.Input

	ToBucketReplicationConfigMapOutput() BucketReplicationConfigMapOutput
	ToBucketReplicationConfigMapOutputWithContext(context.Context) BucketReplicationConfigMapOutput
}

type BucketReplicationConfigMap map[string]BucketReplicationConfigInput

func (BucketReplicationConfigMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*BucketReplicationConfig)(nil)).Elem()
}

func (i BucketReplicationConfigMap) ToBucketReplicationConfigMapOutput() BucketReplicationConfigMapOutput {
	return i.ToBucketReplicationConfigMapOutputWithContext(context.Background())
}

func (i BucketReplicationConfigMap) ToBucketReplicationConfigMapOutputWithContext(ctx context.Context) BucketReplicationConfigMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BucketReplicationConfigMapOutput)
}

func (i BucketReplicationConfigMap) ToOutput(ctx context.Context) pulumix.Output[map[string]*BucketReplicationConfig] {
	return pulumix.Output[map[string]*BucketReplicationConfig]{
		OutputState: i.ToBucketReplicationConfigMapOutputWithContext(ctx).OutputState,
	}
}

type BucketReplicationConfigOutput struct{ *pulumi.OutputState }

func (BucketReplicationConfigOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**BucketReplicationConfig)(nil)).Elem()
}

func (o BucketReplicationConfigOutput) ToBucketReplicationConfigOutput() BucketReplicationConfigOutput {
	return o
}

func (o BucketReplicationConfigOutput) ToBucketReplicationConfigOutputWithContext(ctx context.Context) BucketReplicationConfigOutput {
	return o
}

func (o BucketReplicationConfigOutput) ToOutput(ctx context.Context) pulumix.Output[*BucketReplicationConfig] {
	return pulumix.Output[*BucketReplicationConfig]{
		OutputState: o.OutputState,
	}
}

// Name of the source S3 bucket you want Amazon S3 to monitor.
func (o BucketReplicationConfigOutput) Bucket() pulumi.StringOutput {
	return o.ApplyT(func(v *BucketReplicationConfig) pulumi.StringOutput { return v.Bucket }).(pulumi.StringOutput)
}

// ARN of the IAM role for Amazon S3 to assume when replicating the objects.
func (o BucketReplicationConfigOutput) Role() pulumi.StringOutput {
	return o.ApplyT(func(v *BucketReplicationConfig) pulumi.StringOutput { return v.Role }).(pulumi.StringOutput)
}

// List of configuration blocks describing the rules managing the replication. See below.
func (o BucketReplicationConfigOutput) Rules() BucketReplicationConfigRuleArrayOutput {
	return o.ApplyT(func(v *BucketReplicationConfig) BucketReplicationConfigRuleArrayOutput { return v.Rules }).(BucketReplicationConfigRuleArrayOutput)
}

// Token to allow replication to be enabled on an Object Lock-enabled bucket. You must contact AWS support for the bucket's "Object Lock token".
// For more details, see [Using S3 Object Lock with replication](https://docs.aws.amazon.com/AmazonS3/latest/userguide/object-lock-managing.html#object-lock-managing-replication).
func (o BucketReplicationConfigOutput) Token() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *BucketReplicationConfig) pulumi.StringPtrOutput { return v.Token }).(pulumi.StringPtrOutput)
}

type BucketReplicationConfigArrayOutput struct{ *pulumi.OutputState }

func (BucketReplicationConfigArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*BucketReplicationConfig)(nil)).Elem()
}

func (o BucketReplicationConfigArrayOutput) ToBucketReplicationConfigArrayOutput() BucketReplicationConfigArrayOutput {
	return o
}

func (o BucketReplicationConfigArrayOutput) ToBucketReplicationConfigArrayOutputWithContext(ctx context.Context) BucketReplicationConfigArrayOutput {
	return o
}

func (o BucketReplicationConfigArrayOutput) ToOutput(ctx context.Context) pulumix.Output[[]*BucketReplicationConfig] {
	return pulumix.Output[[]*BucketReplicationConfig]{
		OutputState: o.OutputState,
	}
}

func (o BucketReplicationConfigArrayOutput) Index(i pulumi.IntInput) BucketReplicationConfigOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *BucketReplicationConfig {
		return vs[0].([]*BucketReplicationConfig)[vs[1].(int)]
	}).(BucketReplicationConfigOutput)
}

type BucketReplicationConfigMapOutput struct{ *pulumi.OutputState }

func (BucketReplicationConfigMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*BucketReplicationConfig)(nil)).Elem()
}

func (o BucketReplicationConfigMapOutput) ToBucketReplicationConfigMapOutput() BucketReplicationConfigMapOutput {
	return o
}

func (o BucketReplicationConfigMapOutput) ToBucketReplicationConfigMapOutputWithContext(ctx context.Context) BucketReplicationConfigMapOutput {
	return o
}

func (o BucketReplicationConfigMapOutput) ToOutput(ctx context.Context) pulumix.Output[map[string]*BucketReplicationConfig] {
	return pulumix.Output[map[string]*BucketReplicationConfig]{
		OutputState: o.OutputState,
	}
}

func (o BucketReplicationConfigMapOutput) MapIndex(k pulumi.StringInput) BucketReplicationConfigOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *BucketReplicationConfig {
		return vs[0].(map[string]*BucketReplicationConfig)[vs[1].(string)]
	}).(BucketReplicationConfigOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*BucketReplicationConfigInput)(nil)).Elem(), &BucketReplicationConfig{})
	pulumi.RegisterInputType(reflect.TypeOf((*BucketReplicationConfigArrayInput)(nil)).Elem(), BucketReplicationConfigArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*BucketReplicationConfigMapInput)(nil)).Elem(), BucketReplicationConfigMap{})
	pulumi.RegisterOutputType(BucketReplicationConfigOutput{})
	pulumi.RegisterOutputType(BucketReplicationConfigArrayOutput{})
	pulumi.RegisterOutputType(BucketReplicationConfigMapOutput{})
}
