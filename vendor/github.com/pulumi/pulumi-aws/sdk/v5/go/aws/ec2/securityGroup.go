// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package ec2

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a security group resource.
//
// > **NOTE on Security Groups and Security Group Rules:** This provider currently
// provides both a standalone Security Group Rule resource (a single `ingress` or
// `egress` rule), and a Security Group resource with `ingress` and `egress` rules
// defined in-line. At this time you cannot use a Security Group with in-line rules
// in conjunction with any Security Group Rule resources. Doing so will cause
// a conflict of rule settings and will overwrite rules.
//
// > **NOTE:** Referencing Security Groups across VPC peering has certain restrictions. More information is available in the [VPC Peering User Guide](https://docs.aws.amazon.com/vpc/latest/peering/vpc-peering-security-groups.html).
//
// > **NOTE:** Due to [AWS Lambda improved VPC networking changes that began deploying in September 2019](https://aws.amazon.com/blogs/compute/announcing-improved-vpc-networking-for-aws-lambda-functions/), security groups associated with Lambda Functions can take up to 45 minutes to successfully delete.
//
// ## Example Usage
// ### Basic Usage
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
//			_, err := ec2.NewSecurityGroup(ctx, "allowTls", &ec2.SecurityGroupArgs{
//				Description: pulumi.String("Allow TLS inbound traffic"),
//				VpcId:       pulumi.Any(aws_vpc.Main.Id),
//				Ingress: ec2.SecurityGroupIngressArray{
//					&ec2.SecurityGroupIngressArgs{
//						Description: pulumi.String("TLS from VPC"),
//						FromPort:    pulumi.Int(443),
//						ToPort:      pulumi.Int(443),
//						Protocol:    pulumi.String("tcp"),
//						CidrBlocks: pulumi.StringArray{
//							aws_vpc.Main.Cidr_block,
//						},
//						Ipv6CidrBlocks: pulumi.StringArray{
//							aws_vpc.Main.Ipv6_cidr_block,
//						},
//					},
//				},
//				Egress: ec2.SecurityGroupEgressArray{
//					&ec2.SecurityGroupEgressArgs{
//						FromPort: pulumi.Int(0),
//						ToPort:   pulumi.Int(0),
//						Protocol: pulumi.String("-1"),
//						CidrBlocks: pulumi.StringArray{
//							pulumi.String("0.0.0.0/0"),
//						},
//						Ipv6CidrBlocks: pulumi.StringArray{
//							pulumi.String("::/0"),
//						},
//					},
//				},
//				Tags: pulumi.StringMap{
//					"Name": pulumi.String("allow_tls"),
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
//
// > **NOTE on Egress rules:** By default, AWS creates an `ALLOW ALL` egress rule when creating a new Security Group inside of a VPC. When creating a new Security Group inside a VPC, **this provider will remove this default rule**, and require you specifically re-create it if you desire that rule. We feel this leads to fewer surprises in terms of controlling your egress rules. If you desire this rule to be in place, you can use this `egress` block:
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
//			_, err := ec2.NewSecurityGroup(ctx, "example", &ec2.SecurityGroupArgs{
//				Egress: ec2.SecurityGroupEgressArray{
//					&ec2.SecurityGroupEgressArgs{
//						CidrBlocks: pulumi.StringArray{
//							pulumi.String("0.0.0.0/0"),
//						},
//						FromPort: pulumi.Int(0),
//						Ipv6CidrBlocks: pulumi.StringArray{
//							pulumi.String("::/0"),
//						},
//						Protocol: pulumi.String("-1"),
//						ToPort:   pulumi.Int(0),
//					},
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
// ### Usage With Prefix List IDs
//
// Prefix Lists are either managed by AWS internally, or created by the customer using a
// Prefix List resource. Prefix Lists provided by
// AWS are associated with a prefix list name, or service name, that is linked to a specific region.
// Prefix list IDs are exported on VPC Endpoints, so you can use this format:
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
//			myEndpoint, err := ec2.NewVpcEndpoint(ctx, "myEndpoint", nil)
//			if err != nil {
//				return err
//			}
//			_, err = ec2.NewSecurityGroup(ctx, "example", &ec2.SecurityGroupArgs{
//				Egress: ec2.SecurityGroupEgressArray{
//					&ec2.SecurityGroupEgressArgs{
//						FromPort: pulumi.Int(0),
//						ToPort:   pulumi.Int(0),
//						Protocol: pulumi.String("-1"),
//						PrefixListIds: pulumi.StringArray{
//							myEndpoint.PrefixListId,
//						},
//					},
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
//
// You can also find a specific Prefix List using the `ec2.getPrefixList` data source.
// ### Change of name or name-prefix value
//
// Security Group's Name [cannot be edited after the resource is created](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/working-with-security-groups.html#creating-security-group). In fact, the `name` and `name-prefix` arguments force the creation of a new Security Group resource when they change value. In that case, this provider first deletes the existing Security Group resource and then it creates a new one. If the existing Security Group is associated to a Network Interface resource, the deletion cannot complete. The reason is that Network Interface resources cannot be left with no Security Group attached and the new one is not yet available at that point.
//
// You must invert the default behavior of the provider. That is, first the new Security Group resource must be created, then associated to possible Network Interface resources and finally the old Security Group can be detached and deleted. To force this behavior, you must set the createBeforeDestroy property:
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
//			_, err := ec2.NewSecurityGroup(ctx, "sgWithChangeableName", nil)
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
// Security Groups can be imported using the `security group id`, e.g.,
//
// ```sh
//
//	$ pulumi import aws:ec2/securityGroup:SecurityGroup elb_sg sg-903004f8
//
// ```
type SecurityGroup struct {
	pulumi.CustomResourceState

	// ARN of the security group.
	Arn pulumi.StringOutput `pulumi:"arn"`
	// Description of this egress rule.
	Description pulumi.StringOutput `pulumi:"description"`
	// Configuration block for egress rules. Can be specified multiple times for each egress rule. Each egress block supports fields documented below. This argument is processed in attribute-as-blocks mode.
	Egress SecurityGroupEgressArrayOutput `pulumi:"egress"`
	// Configuration block for ingress rules. Can be specified multiple times for each ingress rule. Each ingress block supports fields documented below. This argument is processed in attribute-as-blocks mode.
	Ingress SecurityGroupIngressArrayOutput `pulumi:"ingress"`
	// Name of the security group. If omitted, this provider will assign a random, unique name.
	Name pulumi.StringOutput `pulumi:"name"`
	// Creates a unique name beginning with the specified prefix. Conflicts with `name`.
	NamePrefix pulumi.StringOutput `pulumi:"namePrefix"`
	// Owner ID.
	OwnerId pulumi.StringOutput `pulumi:"ownerId"`
	// Instruct the provider to revoke all of the Security Groups attached ingress and egress rules before deleting the rule itself. This is normally not needed, however certain AWS services such as Elastic Map Reduce may automatically add required rules to security groups used with the service, and those rules may contain a cyclic dependency that prevent the security groups from being destroyed without removing the dependency first. Default `false`.
	RevokeRulesOnDelete pulumi.BoolPtrOutput `pulumi:"revokeRulesOnDelete"`
	// Map of tags to assign to the resource. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// A map of tags assigned to the resource, including those inherited from the provider `defaultTags` configuration block.
	TagsAll pulumi.StringMapOutput `pulumi:"tagsAll"`
	// VPC ID.
	// Defaults to the region's default VPC.
	VpcId pulumi.StringOutput `pulumi:"vpcId"`
}

// NewSecurityGroup registers a new resource with the given unique name, arguments, and options.
func NewSecurityGroup(ctx *pulumi.Context,
	name string, args *SecurityGroupArgs, opts ...pulumi.ResourceOption) (*SecurityGroup, error) {
	if args == nil {
		args = &SecurityGroupArgs{}
	}

	if isZero(args.Description) {
		args.Description = pulumi.StringPtr("Managed by Pulumi")
	}
	var resource SecurityGroup
	err := ctx.RegisterResource("aws:ec2/securityGroup:SecurityGroup", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSecurityGroup gets an existing SecurityGroup resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSecurityGroup(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SecurityGroupState, opts ...pulumi.ResourceOption) (*SecurityGroup, error) {
	var resource SecurityGroup
	err := ctx.ReadResource("aws:ec2/securityGroup:SecurityGroup", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SecurityGroup resources.
type securityGroupState struct {
	// ARN of the security group.
	Arn *string `pulumi:"arn"`
	// Description of this egress rule.
	Description *string `pulumi:"description"`
	// Configuration block for egress rules. Can be specified multiple times for each egress rule. Each egress block supports fields documented below. This argument is processed in attribute-as-blocks mode.
	Egress []SecurityGroupEgress `pulumi:"egress"`
	// Configuration block for ingress rules. Can be specified multiple times for each ingress rule. Each ingress block supports fields documented below. This argument is processed in attribute-as-blocks mode.
	Ingress []SecurityGroupIngress `pulumi:"ingress"`
	// Name of the security group. If omitted, this provider will assign a random, unique name.
	Name *string `pulumi:"name"`
	// Creates a unique name beginning with the specified prefix. Conflicts with `name`.
	NamePrefix *string `pulumi:"namePrefix"`
	// Owner ID.
	OwnerId *string `pulumi:"ownerId"`
	// Instruct the provider to revoke all of the Security Groups attached ingress and egress rules before deleting the rule itself. This is normally not needed, however certain AWS services such as Elastic Map Reduce may automatically add required rules to security groups used with the service, and those rules may contain a cyclic dependency that prevent the security groups from being destroyed without removing the dependency first. Default `false`.
	RevokeRulesOnDelete *bool `pulumi:"revokeRulesOnDelete"`
	// Map of tags to assign to the resource. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `pulumi:"tags"`
	// A map of tags assigned to the resource, including those inherited from the provider `defaultTags` configuration block.
	TagsAll map[string]string `pulumi:"tagsAll"`
	// VPC ID.
	// Defaults to the region's default VPC.
	VpcId *string `pulumi:"vpcId"`
}

type SecurityGroupState struct {
	// ARN of the security group.
	Arn pulumi.StringPtrInput
	// Description of this egress rule.
	Description pulumi.StringPtrInput
	// Configuration block for egress rules. Can be specified multiple times for each egress rule. Each egress block supports fields documented below. This argument is processed in attribute-as-blocks mode.
	Egress SecurityGroupEgressArrayInput
	// Configuration block for ingress rules. Can be specified multiple times for each ingress rule. Each ingress block supports fields documented below. This argument is processed in attribute-as-blocks mode.
	Ingress SecurityGroupIngressArrayInput
	// Name of the security group. If omitted, this provider will assign a random, unique name.
	Name pulumi.StringPtrInput
	// Creates a unique name beginning with the specified prefix. Conflicts with `name`.
	NamePrefix pulumi.StringPtrInput
	// Owner ID.
	OwnerId pulumi.StringPtrInput
	// Instruct the provider to revoke all of the Security Groups attached ingress and egress rules before deleting the rule itself. This is normally not needed, however certain AWS services such as Elastic Map Reduce may automatically add required rules to security groups used with the service, and those rules may contain a cyclic dependency that prevent the security groups from being destroyed without removing the dependency first. Default `false`.
	RevokeRulesOnDelete pulumi.BoolPtrInput
	// Map of tags to assign to the resource. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags pulumi.StringMapInput
	// A map of tags assigned to the resource, including those inherited from the provider `defaultTags` configuration block.
	TagsAll pulumi.StringMapInput
	// VPC ID.
	// Defaults to the region's default VPC.
	VpcId pulumi.StringPtrInput
}

func (SecurityGroupState) ElementType() reflect.Type {
	return reflect.TypeOf((*securityGroupState)(nil)).Elem()
}

type securityGroupArgs struct {
	// Description of this egress rule.
	Description *string `pulumi:"description"`
	// Configuration block for egress rules. Can be specified multiple times for each egress rule. Each egress block supports fields documented below. This argument is processed in attribute-as-blocks mode.
	Egress []SecurityGroupEgress `pulumi:"egress"`
	// Configuration block for ingress rules. Can be specified multiple times for each ingress rule. Each ingress block supports fields documented below. This argument is processed in attribute-as-blocks mode.
	Ingress []SecurityGroupIngress `pulumi:"ingress"`
	// Name of the security group. If omitted, this provider will assign a random, unique name.
	Name *string `pulumi:"name"`
	// Creates a unique name beginning with the specified prefix. Conflicts with `name`.
	NamePrefix *string `pulumi:"namePrefix"`
	// Instruct the provider to revoke all of the Security Groups attached ingress and egress rules before deleting the rule itself. This is normally not needed, however certain AWS services such as Elastic Map Reduce may automatically add required rules to security groups used with the service, and those rules may contain a cyclic dependency that prevent the security groups from being destroyed without removing the dependency first. Default `false`.
	RevokeRulesOnDelete *bool `pulumi:"revokeRulesOnDelete"`
	// Map of tags to assign to the resource. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `pulumi:"tags"`
	// VPC ID.
	// Defaults to the region's default VPC.
	VpcId *string `pulumi:"vpcId"`
}

// The set of arguments for constructing a SecurityGroup resource.
type SecurityGroupArgs struct {
	// Description of this egress rule.
	Description pulumi.StringPtrInput
	// Configuration block for egress rules. Can be specified multiple times for each egress rule. Each egress block supports fields documented below. This argument is processed in attribute-as-blocks mode.
	Egress SecurityGroupEgressArrayInput
	// Configuration block for ingress rules. Can be specified multiple times for each ingress rule. Each ingress block supports fields documented below. This argument is processed in attribute-as-blocks mode.
	Ingress SecurityGroupIngressArrayInput
	// Name of the security group. If omitted, this provider will assign a random, unique name.
	Name pulumi.StringPtrInput
	// Creates a unique name beginning with the specified prefix. Conflicts with `name`.
	NamePrefix pulumi.StringPtrInput
	// Instruct the provider to revoke all of the Security Groups attached ingress and egress rules before deleting the rule itself. This is normally not needed, however certain AWS services such as Elastic Map Reduce may automatically add required rules to security groups used with the service, and those rules may contain a cyclic dependency that prevent the security groups from being destroyed without removing the dependency first. Default `false`.
	RevokeRulesOnDelete pulumi.BoolPtrInput
	// Map of tags to assign to the resource. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags pulumi.StringMapInput
	// VPC ID.
	// Defaults to the region's default VPC.
	VpcId pulumi.StringPtrInput
}

func (SecurityGroupArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*securityGroupArgs)(nil)).Elem()
}

type SecurityGroupInput interface {
	pulumi.Input

	ToSecurityGroupOutput() SecurityGroupOutput
	ToSecurityGroupOutputWithContext(ctx context.Context) SecurityGroupOutput
}

func (*SecurityGroup) ElementType() reflect.Type {
	return reflect.TypeOf((**SecurityGroup)(nil)).Elem()
}

func (i *SecurityGroup) ToSecurityGroupOutput() SecurityGroupOutput {
	return i.ToSecurityGroupOutputWithContext(context.Background())
}

func (i *SecurityGroup) ToSecurityGroupOutputWithContext(ctx context.Context) SecurityGroupOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SecurityGroupOutput)
}

// SecurityGroupArrayInput is an input type that accepts SecurityGroupArray and SecurityGroupArrayOutput values.
// You can construct a concrete instance of `SecurityGroupArrayInput` via:
//
//	SecurityGroupArray{ SecurityGroupArgs{...} }
type SecurityGroupArrayInput interface {
	pulumi.Input

	ToSecurityGroupArrayOutput() SecurityGroupArrayOutput
	ToSecurityGroupArrayOutputWithContext(context.Context) SecurityGroupArrayOutput
}

type SecurityGroupArray []SecurityGroupInput

func (SecurityGroupArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SecurityGroup)(nil)).Elem()
}

func (i SecurityGroupArray) ToSecurityGroupArrayOutput() SecurityGroupArrayOutput {
	return i.ToSecurityGroupArrayOutputWithContext(context.Background())
}

func (i SecurityGroupArray) ToSecurityGroupArrayOutputWithContext(ctx context.Context) SecurityGroupArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SecurityGroupArrayOutput)
}

// SecurityGroupMapInput is an input type that accepts SecurityGroupMap and SecurityGroupMapOutput values.
// You can construct a concrete instance of `SecurityGroupMapInput` via:
//
//	SecurityGroupMap{ "key": SecurityGroupArgs{...} }
type SecurityGroupMapInput interface {
	pulumi.Input

	ToSecurityGroupMapOutput() SecurityGroupMapOutput
	ToSecurityGroupMapOutputWithContext(context.Context) SecurityGroupMapOutput
}

type SecurityGroupMap map[string]SecurityGroupInput

func (SecurityGroupMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SecurityGroup)(nil)).Elem()
}

func (i SecurityGroupMap) ToSecurityGroupMapOutput() SecurityGroupMapOutput {
	return i.ToSecurityGroupMapOutputWithContext(context.Background())
}

func (i SecurityGroupMap) ToSecurityGroupMapOutputWithContext(ctx context.Context) SecurityGroupMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SecurityGroupMapOutput)
}

type SecurityGroupOutput struct{ *pulumi.OutputState }

func (SecurityGroupOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SecurityGroup)(nil)).Elem()
}

func (o SecurityGroupOutput) ToSecurityGroupOutput() SecurityGroupOutput {
	return o
}

func (o SecurityGroupOutput) ToSecurityGroupOutputWithContext(ctx context.Context) SecurityGroupOutput {
	return o
}

// ARN of the security group.
func (o SecurityGroupOutput) Arn() pulumi.StringOutput {
	return o.ApplyT(func(v *SecurityGroup) pulumi.StringOutput { return v.Arn }).(pulumi.StringOutput)
}

// Description of this egress rule.
func (o SecurityGroupOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v *SecurityGroup) pulumi.StringOutput { return v.Description }).(pulumi.StringOutput)
}

// Configuration block for egress rules. Can be specified multiple times for each egress rule. Each egress block supports fields documented below. This argument is processed in attribute-as-blocks mode.
func (o SecurityGroupOutput) Egress() SecurityGroupEgressArrayOutput {
	return o.ApplyT(func(v *SecurityGroup) SecurityGroupEgressArrayOutput { return v.Egress }).(SecurityGroupEgressArrayOutput)
}

// Configuration block for ingress rules. Can be specified multiple times for each ingress rule. Each ingress block supports fields documented below. This argument is processed in attribute-as-blocks mode.
func (o SecurityGroupOutput) Ingress() SecurityGroupIngressArrayOutput {
	return o.ApplyT(func(v *SecurityGroup) SecurityGroupIngressArrayOutput { return v.Ingress }).(SecurityGroupIngressArrayOutput)
}

// Name of the security group. If omitted, this provider will assign a random, unique name.
func (o SecurityGroupOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *SecurityGroup) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Creates a unique name beginning with the specified prefix. Conflicts with `name`.
func (o SecurityGroupOutput) NamePrefix() pulumi.StringOutput {
	return o.ApplyT(func(v *SecurityGroup) pulumi.StringOutput { return v.NamePrefix }).(pulumi.StringOutput)
}

// Owner ID.
func (o SecurityGroupOutput) OwnerId() pulumi.StringOutput {
	return o.ApplyT(func(v *SecurityGroup) pulumi.StringOutput { return v.OwnerId }).(pulumi.StringOutput)
}

// Instruct the provider to revoke all of the Security Groups attached ingress and egress rules before deleting the rule itself. This is normally not needed, however certain AWS services such as Elastic Map Reduce may automatically add required rules to security groups used with the service, and those rules may contain a cyclic dependency that prevent the security groups from being destroyed without removing the dependency first. Default `false`.
func (o SecurityGroupOutput) RevokeRulesOnDelete() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *SecurityGroup) pulumi.BoolPtrOutput { return v.RevokeRulesOnDelete }).(pulumi.BoolPtrOutput)
}

// Map of tags to assign to the resource. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
func (o SecurityGroupOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *SecurityGroup) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

// A map of tags assigned to the resource, including those inherited from the provider `defaultTags` configuration block.
func (o SecurityGroupOutput) TagsAll() pulumi.StringMapOutput {
	return o.ApplyT(func(v *SecurityGroup) pulumi.StringMapOutput { return v.TagsAll }).(pulumi.StringMapOutput)
}

// VPC ID.
// Defaults to the region's default VPC.
func (o SecurityGroupOutput) VpcId() pulumi.StringOutput {
	return o.ApplyT(func(v *SecurityGroup) pulumi.StringOutput { return v.VpcId }).(pulumi.StringOutput)
}

type SecurityGroupArrayOutput struct{ *pulumi.OutputState }

func (SecurityGroupArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SecurityGroup)(nil)).Elem()
}

func (o SecurityGroupArrayOutput) ToSecurityGroupArrayOutput() SecurityGroupArrayOutput {
	return o
}

func (o SecurityGroupArrayOutput) ToSecurityGroupArrayOutputWithContext(ctx context.Context) SecurityGroupArrayOutput {
	return o
}

func (o SecurityGroupArrayOutput) Index(i pulumi.IntInput) SecurityGroupOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *SecurityGroup {
		return vs[0].([]*SecurityGroup)[vs[1].(int)]
	}).(SecurityGroupOutput)
}

type SecurityGroupMapOutput struct{ *pulumi.OutputState }

func (SecurityGroupMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SecurityGroup)(nil)).Elem()
}

func (o SecurityGroupMapOutput) ToSecurityGroupMapOutput() SecurityGroupMapOutput {
	return o
}

func (o SecurityGroupMapOutput) ToSecurityGroupMapOutputWithContext(ctx context.Context) SecurityGroupMapOutput {
	return o
}

func (o SecurityGroupMapOutput) MapIndex(k pulumi.StringInput) SecurityGroupOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *SecurityGroup {
		return vs[0].(map[string]*SecurityGroup)[vs[1].(string)]
	}).(SecurityGroupOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*SecurityGroupInput)(nil)).Elem(), &SecurityGroup{})
	pulumi.RegisterInputType(reflect.TypeOf((*SecurityGroupArrayInput)(nil)).Elem(), SecurityGroupArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*SecurityGroupMapInput)(nil)).Elem(), SecurityGroupMap{})
	pulumi.RegisterOutputType(SecurityGroupOutput{})
	pulumi.RegisterOutputType(SecurityGroupArrayOutput{})
	pulumi.RegisterOutputType(SecurityGroupMapOutput{})
}
