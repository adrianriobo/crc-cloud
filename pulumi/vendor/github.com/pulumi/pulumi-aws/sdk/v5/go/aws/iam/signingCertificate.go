// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package iam

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides an IAM Signing Certificate resource to upload Signing Certificates.
//
// ## Example Usage
//
// **Using certs on file:**
//
// ```go
// package main
//
// import (
//
//	"io/ioutil"
//
//	"github.com/pulumi/pulumi-aws/sdk/v5/go/aws/iam"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func readFileOrPanic(path string) pulumi.StringPtrInput {
//		data, err := ioutil.ReadFile(path)
//		if err != nil {
//			panic(err.Error())
//		}
//		return pulumi.String(string(data))
//	}
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := iam.NewSigningCertificate(ctx, "testCert", &iam.SigningCertificateArgs{
//				Username:        pulumi.String("some_test_cert"),
//				CertificateBody: readFileOrPanic("self-ca-cert.pem"),
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
// **Example with cert in-line:**
//
// ```go
// package main
//
// import (
//
//	"fmt"
//
//	"github.com/pulumi/pulumi-aws/sdk/v5/go/aws/iam"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := iam.NewSigningCertificate(ctx, "testCertAlt", &iam.SigningCertificateArgs{
//				CertificateBody: pulumi.String(fmt.Sprintf("-----BEGIN CERTIFICATE-----\n[......] # cert contents\n-----END CERTIFICATE-----\n\n")),
//				Username:        pulumi.String("some_test_cert"),
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
// IAM Signing Certificates can be imported using the `id`, e.g.,
//
// ```sh
//
//	$ pulumi import aws:iam/signingCertificate:SigningCertificate certificate IDIDIDIDID:user-name
//
// ```
type SigningCertificate struct {
	pulumi.CustomResourceState

	// The contents of the signing certificate in PEM-encoded format.
	CertificateBody pulumi.StringOutput `pulumi:"certificateBody"`
	// The ID for the signing certificate.
	CertificateId pulumi.StringOutput `pulumi:"certificateId"`
	// The status you want to assign to the certificate. `Active` means that the certificate can be used for programmatic calls to Amazon Web Services `Inactive` means that the certificate cannot be used.
	Status pulumi.StringPtrOutput `pulumi:"status"`
	// The name of the user the signing certificate is for.
	UserName pulumi.StringOutput `pulumi:"userName"`
}

// NewSigningCertificate registers a new resource with the given unique name, arguments, and options.
func NewSigningCertificate(ctx *pulumi.Context,
	name string, args *SigningCertificateArgs, opts ...pulumi.ResourceOption) (*SigningCertificate, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.CertificateBody == nil {
		return nil, errors.New("invalid value for required argument 'CertificateBody'")
	}
	if args.UserName == nil {
		return nil, errors.New("invalid value for required argument 'UserName'")
	}
	var resource SigningCertificate
	err := ctx.RegisterResource("aws:iam/signingCertificate:SigningCertificate", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSigningCertificate gets an existing SigningCertificate resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSigningCertificate(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SigningCertificateState, opts ...pulumi.ResourceOption) (*SigningCertificate, error) {
	var resource SigningCertificate
	err := ctx.ReadResource("aws:iam/signingCertificate:SigningCertificate", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SigningCertificate resources.
type signingCertificateState struct {
	// The contents of the signing certificate in PEM-encoded format.
	CertificateBody *string `pulumi:"certificateBody"`
	// The ID for the signing certificate.
	CertificateId *string `pulumi:"certificateId"`
	// The status you want to assign to the certificate. `Active` means that the certificate can be used for programmatic calls to Amazon Web Services `Inactive` means that the certificate cannot be used.
	Status *string `pulumi:"status"`
	// The name of the user the signing certificate is for.
	UserName *string `pulumi:"userName"`
}

type SigningCertificateState struct {
	// The contents of the signing certificate in PEM-encoded format.
	CertificateBody pulumi.StringPtrInput
	// The ID for the signing certificate.
	CertificateId pulumi.StringPtrInput
	// The status you want to assign to the certificate. `Active` means that the certificate can be used for programmatic calls to Amazon Web Services `Inactive` means that the certificate cannot be used.
	Status pulumi.StringPtrInput
	// The name of the user the signing certificate is for.
	UserName pulumi.StringPtrInput
}

func (SigningCertificateState) ElementType() reflect.Type {
	return reflect.TypeOf((*signingCertificateState)(nil)).Elem()
}

type signingCertificateArgs struct {
	// The contents of the signing certificate in PEM-encoded format.
	CertificateBody string `pulumi:"certificateBody"`
	// The status you want to assign to the certificate. `Active` means that the certificate can be used for programmatic calls to Amazon Web Services `Inactive` means that the certificate cannot be used.
	Status *string `pulumi:"status"`
	// The name of the user the signing certificate is for.
	UserName string `pulumi:"userName"`
}

// The set of arguments for constructing a SigningCertificate resource.
type SigningCertificateArgs struct {
	// The contents of the signing certificate in PEM-encoded format.
	CertificateBody pulumi.StringInput
	// The status you want to assign to the certificate. `Active` means that the certificate can be used for programmatic calls to Amazon Web Services `Inactive` means that the certificate cannot be used.
	Status pulumi.StringPtrInput
	// The name of the user the signing certificate is for.
	UserName pulumi.StringInput
}

func (SigningCertificateArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*signingCertificateArgs)(nil)).Elem()
}

type SigningCertificateInput interface {
	pulumi.Input

	ToSigningCertificateOutput() SigningCertificateOutput
	ToSigningCertificateOutputWithContext(ctx context.Context) SigningCertificateOutput
}

func (*SigningCertificate) ElementType() reflect.Type {
	return reflect.TypeOf((**SigningCertificate)(nil)).Elem()
}

func (i *SigningCertificate) ToSigningCertificateOutput() SigningCertificateOutput {
	return i.ToSigningCertificateOutputWithContext(context.Background())
}

func (i *SigningCertificate) ToSigningCertificateOutputWithContext(ctx context.Context) SigningCertificateOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SigningCertificateOutput)
}

// SigningCertificateArrayInput is an input type that accepts SigningCertificateArray and SigningCertificateArrayOutput values.
// You can construct a concrete instance of `SigningCertificateArrayInput` via:
//
//	SigningCertificateArray{ SigningCertificateArgs{...} }
type SigningCertificateArrayInput interface {
	pulumi.Input

	ToSigningCertificateArrayOutput() SigningCertificateArrayOutput
	ToSigningCertificateArrayOutputWithContext(context.Context) SigningCertificateArrayOutput
}

type SigningCertificateArray []SigningCertificateInput

func (SigningCertificateArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SigningCertificate)(nil)).Elem()
}

func (i SigningCertificateArray) ToSigningCertificateArrayOutput() SigningCertificateArrayOutput {
	return i.ToSigningCertificateArrayOutputWithContext(context.Background())
}

func (i SigningCertificateArray) ToSigningCertificateArrayOutputWithContext(ctx context.Context) SigningCertificateArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SigningCertificateArrayOutput)
}

// SigningCertificateMapInput is an input type that accepts SigningCertificateMap and SigningCertificateMapOutput values.
// You can construct a concrete instance of `SigningCertificateMapInput` via:
//
//	SigningCertificateMap{ "key": SigningCertificateArgs{...} }
type SigningCertificateMapInput interface {
	pulumi.Input

	ToSigningCertificateMapOutput() SigningCertificateMapOutput
	ToSigningCertificateMapOutputWithContext(context.Context) SigningCertificateMapOutput
}

type SigningCertificateMap map[string]SigningCertificateInput

func (SigningCertificateMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SigningCertificate)(nil)).Elem()
}

func (i SigningCertificateMap) ToSigningCertificateMapOutput() SigningCertificateMapOutput {
	return i.ToSigningCertificateMapOutputWithContext(context.Background())
}

func (i SigningCertificateMap) ToSigningCertificateMapOutputWithContext(ctx context.Context) SigningCertificateMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SigningCertificateMapOutput)
}

type SigningCertificateOutput struct{ *pulumi.OutputState }

func (SigningCertificateOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SigningCertificate)(nil)).Elem()
}

func (o SigningCertificateOutput) ToSigningCertificateOutput() SigningCertificateOutput {
	return o
}

func (o SigningCertificateOutput) ToSigningCertificateOutputWithContext(ctx context.Context) SigningCertificateOutput {
	return o
}

// The contents of the signing certificate in PEM-encoded format.
func (o SigningCertificateOutput) CertificateBody() pulumi.StringOutput {
	return o.ApplyT(func(v *SigningCertificate) pulumi.StringOutput { return v.CertificateBody }).(pulumi.StringOutput)
}

// The ID for the signing certificate.
func (o SigningCertificateOutput) CertificateId() pulumi.StringOutput {
	return o.ApplyT(func(v *SigningCertificate) pulumi.StringOutput { return v.CertificateId }).(pulumi.StringOutput)
}

// The status you want to assign to the certificate. `Active` means that the certificate can be used for programmatic calls to Amazon Web Services `Inactive` means that the certificate cannot be used.
func (o SigningCertificateOutput) Status() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SigningCertificate) pulumi.StringPtrOutput { return v.Status }).(pulumi.StringPtrOutput)
}

// The name of the user the signing certificate is for.
func (o SigningCertificateOutput) UserName() pulumi.StringOutput {
	return o.ApplyT(func(v *SigningCertificate) pulumi.StringOutput { return v.UserName }).(pulumi.StringOutput)
}

type SigningCertificateArrayOutput struct{ *pulumi.OutputState }

func (SigningCertificateArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SigningCertificate)(nil)).Elem()
}

func (o SigningCertificateArrayOutput) ToSigningCertificateArrayOutput() SigningCertificateArrayOutput {
	return o
}

func (o SigningCertificateArrayOutput) ToSigningCertificateArrayOutputWithContext(ctx context.Context) SigningCertificateArrayOutput {
	return o
}

func (o SigningCertificateArrayOutput) Index(i pulumi.IntInput) SigningCertificateOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *SigningCertificate {
		return vs[0].([]*SigningCertificate)[vs[1].(int)]
	}).(SigningCertificateOutput)
}

type SigningCertificateMapOutput struct{ *pulumi.OutputState }

func (SigningCertificateMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SigningCertificate)(nil)).Elem()
}

func (o SigningCertificateMapOutput) ToSigningCertificateMapOutput() SigningCertificateMapOutput {
	return o
}

func (o SigningCertificateMapOutput) ToSigningCertificateMapOutputWithContext(ctx context.Context) SigningCertificateMapOutput {
	return o
}

func (o SigningCertificateMapOutput) MapIndex(k pulumi.StringInput) SigningCertificateOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *SigningCertificate {
		return vs[0].(map[string]*SigningCertificate)[vs[1].(string)]
	}).(SigningCertificateOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*SigningCertificateInput)(nil)).Elem(), &SigningCertificate{})
	pulumi.RegisterInputType(reflect.TypeOf((*SigningCertificateArrayInput)(nil)).Elem(), SigningCertificateArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*SigningCertificateMapInput)(nil)).Elem(), SigningCertificateMap{})
	pulumi.RegisterOutputType(SigningCertificateOutput{})
	pulumi.RegisterOutputType(SigningCertificateArrayOutput{})
	pulumi.RegisterOutputType(SigningCertificateMapOutput{})
}
