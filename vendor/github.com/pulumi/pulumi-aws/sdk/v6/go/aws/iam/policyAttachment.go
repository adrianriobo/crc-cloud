// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package iam

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Attaches a Managed IAM Policy to user(s), role(s), and/or group(s)
//
// !> **WARNING:** The iam.PolicyAttachment resource creates **exclusive** attachments of IAM policies. Across the entire AWS account, all of the users/roles/groups to which a single policy is attached must be declared by a single iam.PolicyAttachment resource. This means that even any users/roles/groups that have the attached policy via any other mechanism (including other resources managed by this provider) will have that attached policy revoked by this resource. Consider `iam.RolePolicyAttachment`, `iam.UserPolicyAttachment`, or `iam.GroupPolicyAttachment` instead. These resources do not enforce exclusive attachment of an IAM policy.
//
// > **NOTE:** The usage of this resource conflicts with the `iam.GroupPolicyAttachment`, `iam.RolePolicyAttachment`, and `iam.UserPolicyAttachment` resources and will permanently show a difference if both are defined.
//
// > **NOTE:** For a given role, this resource is incompatible with using the `iam.Role` resource `managedPolicyArns` argument. When using that argument and this resource, both will attempt to manage the role's managed policy attachments and the provider will show a permanent difference.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/iam"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			user, err := iam.NewUser(ctx, "user", nil)
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
//									"ec2.amazonaws.com",
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
//			role, err := iam.NewRole(ctx, "role", &iam.RoleArgs{
//				AssumeRolePolicy: *pulumi.String(assumeRole.Json),
//			})
//			if err != nil {
//				return err
//			}
//			group, err := iam.NewGroup(ctx, "group", nil)
//			if err != nil {
//				return err
//			}
//			policyPolicyDocument, err := iam.GetPolicyDocument(ctx, &iam.GetPolicyDocumentArgs{
//				Statements: []iam.GetPolicyDocumentStatement{
//					{
//						Effect: pulumi.StringRef("Allow"),
//						Actions: []string{
//							"ec2:Describe*",
//						},
//						Resources: []string{
//							"*",
//						},
//					},
//				},
//			}, nil)
//			if err != nil {
//				return err
//			}
//			policyPolicy, err := iam.NewPolicy(ctx, "policyPolicy", &iam.PolicyArgs{
//				Description: pulumi.String("A test policy"),
//				Policy:      *pulumi.String(policyPolicyDocument.Json),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = iam.NewPolicyAttachment(ctx, "test-attach", &iam.PolicyAttachmentArgs{
//				Users: pulumi.AnyArray{
//					user.Name,
//				},
//				Roles: pulumi.AnyArray{
//					role.Name,
//				},
//				Groups: pulumi.AnyArray{
//					group.Name,
//				},
//				PolicyArn: policyPolicy.Arn,
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
type PolicyAttachment struct {
	pulumi.CustomResourceState

	// The group(s) the policy should be applied to
	Groups pulumi.StringArrayOutput `pulumi:"groups"`
	// The name of the attachment. This cannot be an empty string.
	Name pulumi.StringOutput `pulumi:"name"`
	// The ARN of the policy you want to apply
	PolicyArn pulumi.StringOutput `pulumi:"policyArn"`
	// The role(s) the policy should be applied to
	Roles pulumi.StringArrayOutput `pulumi:"roles"`
	// The user(s) the policy should be applied to
	Users pulumi.StringArrayOutput `pulumi:"users"`
}

// NewPolicyAttachment registers a new resource with the given unique name, arguments, and options.
func NewPolicyAttachment(ctx *pulumi.Context,
	name string, args *PolicyAttachmentArgs, opts ...pulumi.ResourceOption) (*PolicyAttachment, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.PolicyArn == nil {
		return nil, errors.New("invalid value for required argument 'PolicyArn'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource PolicyAttachment
	err := ctx.RegisterResource("aws:iam/policyAttachment:PolicyAttachment", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetPolicyAttachment gets an existing PolicyAttachment resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetPolicyAttachment(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *PolicyAttachmentState, opts ...pulumi.ResourceOption) (*PolicyAttachment, error) {
	var resource PolicyAttachment
	err := ctx.ReadResource("aws:iam/policyAttachment:PolicyAttachment", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering PolicyAttachment resources.
type policyAttachmentState struct {
	// The group(s) the policy should be applied to
	Groups []interface{} `pulumi:"groups"`
	// The name of the attachment. This cannot be an empty string.
	Name *string `pulumi:"name"`
	// The ARN of the policy you want to apply
	PolicyArn *string `pulumi:"policyArn"`
	// The role(s) the policy should be applied to
	Roles []interface{} `pulumi:"roles"`
	// The user(s) the policy should be applied to
	Users []interface{} `pulumi:"users"`
}

type PolicyAttachmentState struct {
	// The group(s) the policy should be applied to
	Groups pulumi.ArrayInput
	// The name of the attachment. This cannot be an empty string.
	Name pulumi.StringPtrInput
	// The ARN of the policy you want to apply
	PolicyArn pulumi.StringPtrInput
	// The role(s) the policy should be applied to
	Roles pulumi.ArrayInput
	// The user(s) the policy should be applied to
	Users pulumi.ArrayInput
}

func (PolicyAttachmentState) ElementType() reflect.Type {
	return reflect.TypeOf((*policyAttachmentState)(nil)).Elem()
}

type policyAttachmentArgs struct {
	// The group(s) the policy should be applied to
	Groups []interface{} `pulumi:"groups"`
	// The name of the attachment. This cannot be an empty string.
	Name *string `pulumi:"name"`
	// The ARN of the policy you want to apply
	PolicyArn string `pulumi:"policyArn"`
	// The role(s) the policy should be applied to
	Roles []interface{} `pulumi:"roles"`
	// The user(s) the policy should be applied to
	Users []interface{} `pulumi:"users"`
}

// The set of arguments for constructing a PolicyAttachment resource.
type PolicyAttachmentArgs struct {
	// The group(s) the policy should be applied to
	Groups pulumi.ArrayInput
	// The name of the attachment. This cannot be an empty string.
	Name pulumi.StringPtrInput
	// The ARN of the policy you want to apply
	PolicyArn pulumi.StringInput
	// The role(s) the policy should be applied to
	Roles pulumi.ArrayInput
	// The user(s) the policy should be applied to
	Users pulumi.ArrayInput
}

func (PolicyAttachmentArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*policyAttachmentArgs)(nil)).Elem()
}

type PolicyAttachmentInput interface {
	pulumi.Input

	ToPolicyAttachmentOutput() PolicyAttachmentOutput
	ToPolicyAttachmentOutputWithContext(ctx context.Context) PolicyAttachmentOutput
}

func (*PolicyAttachment) ElementType() reflect.Type {
	return reflect.TypeOf((**PolicyAttachment)(nil)).Elem()
}

func (i *PolicyAttachment) ToPolicyAttachmentOutput() PolicyAttachmentOutput {
	return i.ToPolicyAttachmentOutputWithContext(context.Background())
}

func (i *PolicyAttachment) ToPolicyAttachmentOutputWithContext(ctx context.Context) PolicyAttachmentOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PolicyAttachmentOutput)
}

func (i *PolicyAttachment) ToOutput(ctx context.Context) pulumix.Output[*PolicyAttachment] {
	return pulumix.Output[*PolicyAttachment]{
		OutputState: i.ToPolicyAttachmentOutputWithContext(ctx).OutputState,
	}
}

// PolicyAttachmentArrayInput is an input type that accepts PolicyAttachmentArray and PolicyAttachmentArrayOutput values.
// You can construct a concrete instance of `PolicyAttachmentArrayInput` via:
//
//	PolicyAttachmentArray{ PolicyAttachmentArgs{...} }
type PolicyAttachmentArrayInput interface {
	pulumi.Input

	ToPolicyAttachmentArrayOutput() PolicyAttachmentArrayOutput
	ToPolicyAttachmentArrayOutputWithContext(context.Context) PolicyAttachmentArrayOutput
}

type PolicyAttachmentArray []PolicyAttachmentInput

func (PolicyAttachmentArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*PolicyAttachment)(nil)).Elem()
}

func (i PolicyAttachmentArray) ToPolicyAttachmentArrayOutput() PolicyAttachmentArrayOutput {
	return i.ToPolicyAttachmentArrayOutputWithContext(context.Background())
}

func (i PolicyAttachmentArray) ToPolicyAttachmentArrayOutputWithContext(ctx context.Context) PolicyAttachmentArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PolicyAttachmentArrayOutput)
}

func (i PolicyAttachmentArray) ToOutput(ctx context.Context) pulumix.Output[[]*PolicyAttachment] {
	return pulumix.Output[[]*PolicyAttachment]{
		OutputState: i.ToPolicyAttachmentArrayOutputWithContext(ctx).OutputState,
	}
}

// PolicyAttachmentMapInput is an input type that accepts PolicyAttachmentMap and PolicyAttachmentMapOutput values.
// You can construct a concrete instance of `PolicyAttachmentMapInput` via:
//
//	PolicyAttachmentMap{ "key": PolicyAttachmentArgs{...} }
type PolicyAttachmentMapInput interface {
	pulumi.Input

	ToPolicyAttachmentMapOutput() PolicyAttachmentMapOutput
	ToPolicyAttachmentMapOutputWithContext(context.Context) PolicyAttachmentMapOutput
}

type PolicyAttachmentMap map[string]PolicyAttachmentInput

func (PolicyAttachmentMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*PolicyAttachment)(nil)).Elem()
}

func (i PolicyAttachmentMap) ToPolicyAttachmentMapOutput() PolicyAttachmentMapOutput {
	return i.ToPolicyAttachmentMapOutputWithContext(context.Background())
}

func (i PolicyAttachmentMap) ToPolicyAttachmentMapOutputWithContext(ctx context.Context) PolicyAttachmentMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PolicyAttachmentMapOutput)
}

func (i PolicyAttachmentMap) ToOutput(ctx context.Context) pulumix.Output[map[string]*PolicyAttachment] {
	return pulumix.Output[map[string]*PolicyAttachment]{
		OutputState: i.ToPolicyAttachmentMapOutputWithContext(ctx).OutputState,
	}
}

type PolicyAttachmentOutput struct{ *pulumi.OutputState }

func (PolicyAttachmentOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PolicyAttachment)(nil)).Elem()
}

func (o PolicyAttachmentOutput) ToPolicyAttachmentOutput() PolicyAttachmentOutput {
	return o
}

func (o PolicyAttachmentOutput) ToPolicyAttachmentOutputWithContext(ctx context.Context) PolicyAttachmentOutput {
	return o
}

func (o PolicyAttachmentOutput) ToOutput(ctx context.Context) pulumix.Output[*PolicyAttachment] {
	return pulumix.Output[*PolicyAttachment]{
		OutputState: o.OutputState,
	}
}

// The group(s) the policy should be applied to
func (o PolicyAttachmentOutput) Groups() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *PolicyAttachment) pulumi.StringArrayOutput { return v.Groups }).(pulumi.StringArrayOutput)
}

// The name of the attachment. This cannot be an empty string.
func (o PolicyAttachmentOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *PolicyAttachment) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The ARN of the policy you want to apply
func (o PolicyAttachmentOutput) PolicyArn() pulumi.StringOutput {
	return o.ApplyT(func(v *PolicyAttachment) pulumi.StringOutput { return v.PolicyArn }).(pulumi.StringOutput)
}

// The role(s) the policy should be applied to
func (o PolicyAttachmentOutput) Roles() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *PolicyAttachment) pulumi.StringArrayOutput { return v.Roles }).(pulumi.StringArrayOutput)
}

// The user(s) the policy should be applied to
func (o PolicyAttachmentOutput) Users() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *PolicyAttachment) pulumi.StringArrayOutput { return v.Users }).(pulumi.StringArrayOutput)
}

type PolicyAttachmentArrayOutput struct{ *pulumi.OutputState }

func (PolicyAttachmentArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*PolicyAttachment)(nil)).Elem()
}

func (o PolicyAttachmentArrayOutput) ToPolicyAttachmentArrayOutput() PolicyAttachmentArrayOutput {
	return o
}

func (o PolicyAttachmentArrayOutput) ToPolicyAttachmentArrayOutputWithContext(ctx context.Context) PolicyAttachmentArrayOutput {
	return o
}

func (o PolicyAttachmentArrayOutput) ToOutput(ctx context.Context) pulumix.Output[[]*PolicyAttachment] {
	return pulumix.Output[[]*PolicyAttachment]{
		OutputState: o.OutputState,
	}
}

func (o PolicyAttachmentArrayOutput) Index(i pulumi.IntInput) PolicyAttachmentOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *PolicyAttachment {
		return vs[0].([]*PolicyAttachment)[vs[1].(int)]
	}).(PolicyAttachmentOutput)
}

type PolicyAttachmentMapOutput struct{ *pulumi.OutputState }

func (PolicyAttachmentMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*PolicyAttachment)(nil)).Elem()
}

func (o PolicyAttachmentMapOutput) ToPolicyAttachmentMapOutput() PolicyAttachmentMapOutput {
	return o
}

func (o PolicyAttachmentMapOutput) ToPolicyAttachmentMapOutputWithContext(ctx context.Context) PolicyAttachmentMapOutput {
	return o
}

func (o PolicyAttachmentMapOutput) ToOutput(ctx context.Context) pulumix.Output[map[string]*PolicyAttachment] {
	return pulumix.Output[map[string]*PolicyAttachment]{
		OutputState: o.OutputState,
	}
}

func (o PolicyAttachmentMapOutput) MapIndex(k pulumi.StringInput) PolicyAttachmentOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *PolicyAttachment {
		return vs[0].(map[string]*PolicyAttachment)[vs[1].(string)]
	}).(PolicyAttachmentOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*PolicyAttachmentInput)(nil)).Elem(), &PolicyAttachment{})
	pulumi.RegisterInputType(reflect.TypeOf((*PolicyAttachmentArrayInput)(nil)).Elem(), PolicyAttachmentArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*PolicyAttachmentMapInput)(nil)).Elem(), PolicyAttachmentMap{})
	pulumi.RegisterOutputType(PolicyAttachmentOutput{})
	pulumi.RegisterOutputType(PolicyAttachmentArrayOutput{})
	pulumi.RegisterOutputType(PolicyAttachmentMapOutput{})
}
