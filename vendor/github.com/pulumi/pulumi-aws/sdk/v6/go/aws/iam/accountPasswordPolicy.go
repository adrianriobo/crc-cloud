// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package iam

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// > **Note:** There is only a single policy allowed per AWS account. An existing policy will be lost when using this resource as an effect of this limitation.
//
// Manages Password Policy for the AWS Account.
// See more about [Account Password Policy](http://docs.aws.amazon.com/IAM/latest/UserGuide/id_credentials_passwords_account-policy.html)
// in the official AWS docs.
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
//			_, err := iam.NewAccountPasswordPolicy(ctx, "strict", &iam.AccountPasswordPolicyArgs{
//				AllowUsersToChangePassword: pulumi.Bool(true),
//				MinimumPasswordLength:      pulumi.Int(8),
//				RequireLowercaseCharacters: pulumi.Bool(true),
//				RequireNumbers:             pulumi.Bool(true),
//				RequireSymbols:             pulumi.Bool(true),
//				RequireUppercaseCharacters: pulumi.Bool(true),
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
// Using `pulumi import`, import IAM Account Password Policy using the word `iam-account-password-policy`. For example:
//
// ```sh
//
//	$ pulumi import aws:iam/accountPasswordPolicy:AccountPasswordPolicy strict iam-account-password-policy
//
// ```
type AccountPasswordPolicy struct {
	pulumi.CustomResourceState

	// Whether to allow users to change their own password
	AllowUsersToChangePassword pulumi.BoolPtrOutput `pulumi:"allowUsersToChangePassword"`
	// Indicates whether passwords in the account expire. Returns `true` if `maxPasswordAge` contains a value greater than `0`. Returns `false` if it is `0` or _not present_.
	ExpirePasswords pulumi.BoolOutput `pulumi:"expirePasswords"`
	// Whether users are prevented from setting a new password after their password has expired (i.e., require administrator reset)
	HardExpiry pulumi.BoolOutput `pulumi:"hardExpiry"`
	// The number of days that an user password is valid.
	MaxPasswordAge pulumi.IntOutput `pulumi:"maxPasswordAge"`
	// Minimum length to require for user passwords.
	MinimumPasswordLength pulumi.IntPtrOutput `pulumi:"minimumPasswordLength"`
	// The number of previous passwords that users are prevented from reusing.
	PasswordReusePrevention pulumi.IntOutput `pulumi:"passwordReusePrevention"`
	// Whether to require lowercase characters for user passwords.
	RequireLowercaseCharacters pulumi.BoolOutput `pulumi:"requireLowercaseCharacters"`
	// Whether to require numbers for user passwords.
	RequireNumbers pulumi.BoolOutput `pulumi:"requireNumbers"`
	// Whether to require symbols for user passwords.
	RequireSymbols pulumi.BoolOutput `pulumi:"requireSymbols"`
	// Whether to require uppercase characters for user passwords.
	RequireUppercaseCharacters pulumi.BoolOutput `pulumi:"requireUppercaseCharacters"`
}

// NewAccountPasswordPolicy registers a new resource with the given unique name, arguments, and options.
func NewAccountPasswordPolicy(ctx *pulumi.Context,
	name string, args *AccountPasswordPolicyArgs, opts ...pulumi.ResourceOption) (*AccountPasswordPolicy, error) {
	if args == nil {
		args = &AccountPasswordPolicyArgs{}
	}

	opts = internal.PkgResourceDefaultOpts(opts)
	var resource AccountPasswordPolicy
	err := ctx.RegisterResource("aws:iam/accountPasswordPolicy:AccountPasswordPolicy", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetAccountPasswordPolicy gets an existing AccountPasswordPolicy resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetAccountPasswordPolicy(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AccountPasswordPolicyState, opts ...pulumi.ResourceOption) (*AccountPasswordPolicy, error) {
	var resource AccountPasswordPolicy
	err := ctx.ReadResource("aws:iam/accountPasswordPolicy:AccountPasswordPolicy", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering AccountPasswordPolicy resources.
type accountPasswordPolicyState struct {
	// Whether to allow users to change their own password
	AllowUsersToChangePassword *bool `pulumi:"allowUsersToChangePassword"`
	// Indicates whether passwords in the account expire. Returns `true` if `maxPasswordAge` contains a value greater than `0`. Returns `false` if it is `0` or _not present_.
	ExpirePasswords *bool `pulumi:"expirePasswords"`
	// Whether users are prevented from setting a new password after their password has expired (i.e., require administrator reset)
	HardExpiry *bool `pulumi:"hardExpiry"`
	// The number of days that an user password is valid.
	MaxPasswordAge *int `pulumi:"maxPasswordAge"`
	// Minimum length to require for user passwords.
	MinimumPasswordLength *int `pulumi:"minimumPasswordLength"`
	// The number of previous passwords that users are prevented from reusing.
	PasswordReusePrevention *int `pulumi:"passwordReusePrevention"`
	// Whether to require lowercase characters for user passwords.
	RequireLowercaseCharacters *bool `pulumi:"requireLowercaseCharacters"`
	// Whether to require numbers for user passwords.
	RequireNumbers *bool `pulumi:"requireNumbers"`
	// Whether to require symbols for user passwords.
	RequireSymbols *bool `pulumi:"requireSymbols"`
	// Whether to require uppercase characters for user passwords.
	RequireUppercaseCharacters *bool `pulumi:"requireUppercaseCharacters"`
}

type AccountPasswordPolicyState struct {
	// Whether to allow users to change their own password
	AllowUsersToChangePassword pulumi.BoolPtrInput
	// Indicates whether passwords in the account expire. Returns `true` if `maxPasswordAge` contains a value greater than `0`. Returns `false` if it is `0` or _not present_.
	ExpirePasswords pulumi.BoolPtrInput
	// Whether users are prevented from setting a new password after their password has expired (i.e., require administrator reset)
	HardExpiry pulumi.BoolPtrInput
	// The number of days that an user password is valid.
	MaxPasswordAge pulumi.IntPtrInput
	// Minimum length to require for user passwords.
	MinimumPasswordLength pulumi.IntPtrInput
	// The number of previous passwords that users are prevented from reusing.
	PasswordReusePrevention pulumi.IntPtrInput
	// Whether to require lowercase characters for user passwords.
	RequireLowercaseCharacters pulumi.BoolPtrInput
	// Whether to require numbers for user passwords.
	RequireNumbers pulumi.BoolPtrInput
	// Whether to require symbols for user passwords.
	RequireSymbols pulumi.BoolPtrInput
	// Whether to require uppercase characters for user passwords.
	RequireUppercaseCharacters pulumi.BoolPtrInput
}

func (AccountPasswordPolicyState) ElementType() reflect.Type {
	return reflect.TypeOf((*accountPasswordPolicyState)(nil)).Elem()
}

type accountPasswordPolicyArgs struct {
	// Whether to allow users to change their own password
	AllowUsersToChangePassword *bool `pulumi:"allowUsersToChangePassword"`
	// Whether users are prevented from setting a new password after their password has expired (i.e., require administrator reset)
	HardExpiry *bool `pulumi:"hardExpiry"`
	// The number of days that an user password is valid.
	MaxPasswordAge *int `pulumi:"maxPasswordAge"`
	// Minimum length to require for user passwords.
	MinimumPasswordLength *int `pulumi:"minimumPasswordLength"`
	// The number of previous passwords that users are prevented from reusing.
	PasswordReusePrevention *int `pulumi:"passwordReusePrevention"`
	// Whether to require lowercase characters for user passwords.
	RequireLowercaseCharacters *bool `pulumi:"requireLowercaseCharacters"`
	// Whether to require numbers for user passwords.
	RequireNumbers *bool `pulumi:"requireNumbers"`
	// Whether to require symbols for user passwords.
	RequireSymbols *bool `pulumi:"requireSymbols"`
	// Whether to require uppercase characters for user passwords.
	RequireUppercaseCharacters *bool `pulumi:"requireUppercaseCharacters"`
}

// The set of arguments for constructing a AccountPasswordPolicy resource.
type AccountPasswordPolicyArgs struct {
	// Whether to allow users to change their own password
	AllowUsersToChangePassword pulumi.BoolPtrInput
	// Whether users are prevented from setting a new password after their password has expired (i.e., require administrator reset)
	HardExpiry pulumi.BoolPtrInput
	// The number of days that an user password is valid.
	MaxPasswordAge pulumi.IntPtrInput
	// Minimum length to require for user passwords.
	MinimumPasswordLength pulumi.IntPtrInput
	// The number of previous passwords that users are prevented from reusing.
	PasswordReusePrevention pulumi.IntPtrInput
	// Whether to require lowercase characters for user passwords.
	RequireLowercaseCharacters pulumi.BoolPtrInput
	// Whether to require numbers for user passwords.
	RequireNumbers pulumi.BoolPtrInput
	// Whether to require symbols for user passwords.
	RequireSymbols pulumi.BoolPtrInput
	// Whether to require uppercase characters for user passwords.
	RequireUppercaseCharacters pulumi.BoolPtrInput
}

func (AccountPasswordPolicyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*accountPasswordPolicyArgs)(nil)).Elem()
}

type AccountPasswordPolicyInput interface {
	pulumi.Input

	ToAccountPasswordPolicyOutput() AccountPasswordPolicyOutput
	ToAccountPasswordPolicyOutputWithContext(ctx context.Context) AccountPasswordPolicyOutput
}

func (*AccountPasswordPolicy) ElementType() reflect.Type {
	return reflect.TypeOf((**AccountPasswordPolicy)(nil)).Elem()
}

func (i *AccountPasswordPolicy) ToAccountPasswordPolicyOutput() AccountPasswordPolicyOutput {
	return i.ToAccountPasswordPolicyOutputWithContext(context.Background())
}

func (i *AccountPasswordPolicy) ToAccountPasswordPolicyOutputWithContext(ctx context.Context) AccountPasswordPolicyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AccountPasswordPolicyOutput)
}

func (i *AccountPasswordPolicy) ToOutput(ctx context.Context) pulumix.Output[*AccountPasswordPolicy] {
	return pulumix.Output[*AccountPasswordPolicy]{
		OutputState: i.ToAccountPasswordPolicyOutputWithContext(ctx).OutputState,
	}
}

// AccountPasswordPolicyArrayInput is an input type that accepts AccountPasswordPolicyArray and AccountPasswordPolicyArrayOutput values.
// You can construct a concrete instance of `AccountPasswordPolicyArrayInput` via:
//
//	AccountPasswordPolicyArray{ AccountPasswordPolicyArgs{...} }
type AccountPasswordPolicyArrayInput interface {
	pulumi.Input

	ToAccountPasswordPolicyArrayOutput() AccountPasswordPolicyArrayOutput
	ToAccountPasswordPolicyArrayOutputWithContext(context.Context) AccountPasswordPolicyArrayOutput
}

type AccountPasswordPolicyArray []AccountPasswordPolicyInput

func (AccountPasswordPolicyArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*AccountPasswordPolicy)(nil)).Elem()
}

func (i AccountPasswordPolicyArray) ToAccountPasswordPolicyArrayOutput() AccountPasswordPolicyArrayOutput {
	return i.ToAccountPasswordPolicyArrayOutputWithContext(context.Background())
}

func (i AccountPasswordPolicyArray) ToAccountPasswordPolicyArrayOutputWithContext(ctx context.Context) AccountPasswordPolicyArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AccountPasswordPolicyArrayOutput)
}

func (i AccountPasswordPolicyArray) ToOutput(ctx context.Context) pulumix.Output[[]*AccountPasswordPolicy] {
	return pulumix.Output[[]*AccountPasswordPolicy]{
		OutputState: i.ToAccountPasswordPolicyArrayOutputWithContext(ctx).OutputState,
	}
}

// AccountPasswordPolicyMapInput is an input type that accepts AccountPasswordPolicyMap and AccountPasswordPolicyMapOutput values.
// You can construct a concrete instance of `AccountPasswordPolicyMapInput` via:
//
//	AccountPasswordPolicyMap{ "key": AccountPasswordPolicyArgs{...} }
type AccountPasswordPolicyMapInput interface {
	pulumi.Input

	ToAccountPasswordPolicyMapOutput() AccountPasswordPolicyMapOutput
	ToAccountPasswordPolicyMapOutputWithContext(context.Context) AccountPasswordPolicyMapOutput
}

type AccountPasswordPolicyMap map[string]AccountPasswordPolicyInput

func (AccountPasswordPolicyMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*AccountPasswordPolicy)(nil)).Elem()
}

func (i AccountPasswordPolicyMap) ToAccountPasswordPolicyMapOutput() AccountPasswordPolicyMapOutput {
	return i.ToAccountPasswordPolicyMapOutputWithContext(context.Background())
}

func (i AccountPasswordPolicyMap) ToAccountPasswordPolicyMapOutputWithContext(ctx context.Context) AccountPasswordPolicyMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AccountPasswordPolicyMapOutput)
}

func (i AccountPasswordPolicyMap) ToOutput(ctx context.Context) pulumix.Output[map[string]*AccountPasswordPolicy] {
	return pulumix.Output[map[string]*AccountPasswordPolicy]{
		OutputState: i.ToAccountPasswordPolicyMapOutputWithContext(ctx).OutputState,
	}
}

type AccountPasswordPolicyOutput struct{ *pulumi.OutputState }

func (AccountPasswordPolicyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AccountPasswordPolicy)(nil)).Elem()
}

func (o AccountPasswordPolicyOutput) ToAccountPasswordPolicyOutput() AccountPasswordPolicyOutput {
	return o
}

func (o AccountPasswordPolicyOutput) ToAccountPasswordPolicyOutputWithContext(ctx context.Context) AccountPasswordPolicyOutput {
	return o
}

func (o AccountPasswordPolicyOutput) ToOutput(ctx context.Context) pulumix.Output[*AccountPasswordPolicy] {
	return pulumix.Output[*AccountPasswordPolicy]{
		OutputState: o.OutputState,
	}
}

// Whether to allow users to change their own password
func (o AccountPasswordPolicyOutput) AllowUsersToChangePassword() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *AccountPasswordPolicy) pulumi.BoolPtrOutput { return v.AllowUsersToChangePassword }).(pulumi.BoolPtrOutput)
}

// Indicates whether passwords in the account expire. Returns `true` if `maxPasswordAge` contains a value greater than `0`. Returns `false` if it is `0` or _not present_.
func (o AccountPasswordPolicyOutput) ExpirePasswords() pulumi.BoolOutput {
	return o.ApplyT(func(v *AccountPasswordPolicy) pulumi.BoolOutput { return v.ExpirePasswords }).(pulumi.BoolOutput)
}

// Whether users are prevented from setting a new password after their password has expired (i.e., require administrator reset)
func (o AccountPasswordPolicyOutput) HardExpiry() pulumi.BoolOutput {
	return o.ApplyT(func(v *AccountPasswordPolicy) pulumi.BoolOutput { return v.HardExpiry }).(pulumi.BoolOutput)
}

// The number of days that an user password is valid.
func (o AccountPasswordPolicyOutput) MaxPasswordAge() pulumi.IntOutput {
	return o.ApplyT(func(v *AccountPasswordPolicy) pulumi.IntOutput { return v.MaxPasswordAge }).(pulumi.IntOutput)
}

// Minimum length to require for user passwords.
func (o AccountPasswordPolicyOutput) MinimumPasswordLength() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *AccountPasswordPolicy) pulumi.IntPtrOutput { return v.MinimumPasswordLength }).(pulumi.IntPtrOutput)
}

// The number of previous passwords that users are prevented from reusing.
func (o AccountPasswordPolicyOutput) PasswordReusePrevention() pulumi.IntOutput {
	return o.ApplyT(func(v *AccountPasswordPolicy) pulumi.IntOutput { return v.PasswordReusePrevention }).(pulumi.IntOutput)
}

// Whether to require lowercase characters for user passwords.
func (o AccountPasswordPolicyOutput) RequireLowercaseCharacters() pulumi.BoolOutput {
	return o.ApplyT(func(v *AccountPasswordPolicy) pulumi.BoolOutput { return v.RequireLowercaseCharacters }).(pulumi.BoolOutput)
}

// Whether to require numbers for user passwords.
func (o AccountPasswordPolicyOutput) RequireNumbers() pulumi.BoolOutput {
	return o.ApplyT(func(v *AccountPasswordPolicy) pulumi.BoolOutput { return v.RequireNumbers }).(pulumi.BoolOutput)
}

// Whether to require symbols for user passwords.
func (o AccountPasswordPolicyOutput) RequireSymbols() pulumi.BoolOutput {
	return o.ApplyT(func(v *AccountPasswordPolicy) pulumi.BoolOutput { return v.RequireSymbols }).(pulumi.BoolOutput)
}

// Whether to require uppercase characters for user passwords.
func (o AccountPasswordPolicyOutput) RequireUppercaseCharacters() pulumi.BoolOutput {
	return o.ApplyT(func(v *AccountPasswordPolicy) pulumi.BoolOutput { return v.RequireUppercaseCharacters }).(pulumi.BoolOutput)
}

type AccountPasswordPolicyArrayOutput struct{ *pulumi.OutputState }

func (AccountPasswordPolicyArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*AccountPasswordPolicy)(nil)).Elem()
}

func (o AccountPasswordPolicyArrayOutput) ToAccountPasswordPolicyArrayOutput() AccountPasswordPolicyArrayOutput {
	return o
}

func (o AccountPasswordPolicyArrayOutput) ToAccountPasswordPolicyArrayOutputWithContext(ctx context.Context) AccountPasswordPolicyArrayOutput {
	return o
}

func (o AccountPasswordPolicyArrayOutput) ToOutput(ctx context.Context) pulumix.Output[[]*AccountPasswordPolicy] {
	return pulumix.Output[[]*AccountPasswordPolicy]{
		OutputState: o.OutputState,
	}
}

func (o AccountPasswordPolicyArrayOutput) Index(i pulumi.IntInput) AccountPasswordPolicyOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *AccountPasswordPolicy {
		return vs[0].([]*AccountPasswordPolicy)[vs[1].(int)]
	}).(AccountPasswordPolicyOutput)
}

type AccountPasswordPolicyMapOutput struct{ *pulumi.OutputState }

func (AccountPasswordPolicyMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*AccountPasswordPolicy)(nil)).Elem()
}

func (o AccountPasswordPolicyMapOutput) ToAccountPasswordPolicyMapOutput() AccountPasswordPolicyMapOutput {
	return o
}

func (o AccountPasswordPolicyMapOutput) ToAccountPasswordPolicyMapOutputWithContext(ctx context.Context) AccountPasswordPolicyMapOutput {
	return o
}

func (o AccountPasswordPolicyMapOutput) ToOutput(ctx context.Context) pulumix.Output[map[string]*AccountPasswordPolicy] {
	return pulumix.Output[map[string]*AccountPasswordPolicy]{
		OutputState: o.OutputState,
	}
}

func (o AccountPasswordPolicyMapOutput) MapIndex(k pulumi.StringInput) AccountPasswordPolicyOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *AccountPasswordPolicy {
		return vs[0].(map[string]*AccountPasswordPolicy)[vs[1].(string)]
	}).(AccountPasswordPolicyOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*AccountPasswordPolicyInput)(nil)).Elem(), &AccountPasswordPolicy{})
	pulumi.RegisterInputType(reflect.TypeOf((*AccountPasswordPolicyArrayInput)(nil)).Elem(), AccountPasswordPolicyArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*AccountPasswordPolicyMapInput)(nil)).Elem(), AccountPasswordPolicyMap{})
	pulumi.RegisterOutputType(AccountPasswordPolicyOutput{})
	pulumi.RegisterOutputType(AccountPasswordPolicyArrayOutput{})
	pulumi.RegisterOutputType(AccountPasswordPolicyMapOutput{})
}
