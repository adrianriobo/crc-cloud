// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package iam

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use this data source to lookup information about IAM Server Certificates.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-aws/sdk/v5/go/aws/elb"
//	"github.com/pulumi/pulumi-aws/sdk/v5/go/aws/iam"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			my_domain, err := iam.LookupServerCertificate(ctx, &iam.LookupServerCertificateArgs{
//				NamePrefix: pulumi.StringRef("my-domain.org"),
//				Latest:     pulumi.BoolRef(true),
//			}, nil)
//			if err != nil {
//				return err
//			}
//			_, err = elb.NewLoadBalancer(ctx, "elb", &elb.LoadBalancerArgs{
//				Listeners: elb.LoadBalancerListenerArray{
//					&elb.LoadBalancerListenerArgs{
//						InstancePort:     pulumi.Int(8000),
//						InstanceProtocol: pulumi.String("https"),
//						LbPort:           pulumi.Int(443),
//						LbProtocol:       pulumi.String("https"),
//						SslCertificateId: *pulumi.String(my_domain.Arn),
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
func LookupServerCertificate(ctx *pulumi.Context, args *LookupServerCertificateArgs, opts ...pulumi.InvokeOption) (*LookupServerCertificateResult, error) {
	var rv LookupServerCertificateResult
	err := ctx.Invoke("aws:iam/getServerCertificate:getServerCertificate", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getServerCertificate.
type LookupServerCertificateArgs struct {
	// sort results by expiration date. returns the certificate with expiration date in furthest in the future.
	Latest *bool `pulumi:"latest"`
	// exact name of the cert to lookup
	Name *string `pulumi:"name"`
	// prefix of cert to filter by
	NamePrefix *string `pulumi:"namePrefix"`
	// prefix of path to filter by
	PathPrefix *string `pulumi:"pathPrefix"`
}

// A collection of values returned by getServerCertificate.
type LookupServerCertificateResult struct {
	Arn              string `pulumi:"arn"`
	CertificateBody  string `pulumi:"certificateBody"`
	CertificateChain string `pulumi:"certificateChain"`
	ExpirationDate   string `pulumi:"expirationDate"`
	// The provider-assigned unique ID for this managed resource.
	Id         string  `pulumi:"id"`
	Latest     *bool   `pulumi:"latest"`
	Name       string  `pulumi:"name"`
	NamePrefix *string `pulumi:"namePrefix"`
	Path       string  `pulumi:"path"`
	PathPrefix *string `pulumi:"pathPrefix"`
	UploadDate string  `pulumi:"uploadDate"`
}

func LookupServerCertificateOutput(ctx *pulumi.Context, args LookupServerCertificateOutputArgs, opts ...pulumi.InvokeOption) LookupServerCertificateResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupServerCertificateResult, error) {
			args := v.(LookupServerCertificateArgs)
			r, err := LookupServerCertificate(ctx, &args, opts...)
			var s LookupServerCertificateResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupServerCertificateResultOutput)
}

// A collection of arguments for invoking getServerCertificate.
type LookupServerCertificateOutputArgs struct {
	// sort results by expiration date. returns the certificate with expiration date in furthest in the future.
	Latest pulumi.BoolPtrInput `pulumi:"latest"`
	// exact name of the cert to lookup
	Name pulumi.StringPtrInput `pulumi:"name"`
	// prefix of cert to filter by
	NamePrefix pulumi.StringPtrInput `pulumi:"namePrefix"`
	// prefix of path to filter by
	PathPrefix pulumi.StringPtrInput `pulumi:"pathPrefix"`
}

func (LookupServerCertificateOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupServerCertificateArgs)(nil)).Elem()
}

// A collection of values returned by getServerCertificate.
type LookupServerCertificateResultOutput struct{ *pulumi.OutputState }

func (LookupServerCertificateResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupServerCertificateResult)(nil)).Elem()
}

func (o LookupServerCertificateResultOutput) ToLookupServerCertificateResultOutput() LookupServerCertificateResultOutput {
	return o
}

func (o LookupServerCertificateResultOutput) ToLookupServerCertificateResultOutputWithContext(ctx context.Context) LookupServerCertificateResultOutput {
	return o
}

func (o LookupServerCertificateResultOutput) Arn() pulumi.StringOutput {
	return o.ApplyT(func(v LookupServerCertificateResult) string { return v.Arn }).(pulumi.StringOutput)
}

func (o LookupServerCertificateResultOutput) CertificateBody() pulumi.StringOutput {
	return o.ApplyT(func(v LookupServerCertificateResult) string { return v.CertificateBody }).(pulumi.StringOutput)
}

func (o LookupServerCertificateResultOutput) CertificateChain() pulumi.StringOutput {
	return o.ApplyT(func(v LookupServerCertificateResult) string { return v.CertificateChain }).(pulumi.StringOutput)
}

func (o LookupServerCertificateResultOutput) ExpirationDate() pulumi.StringOutput {
	return o.ApplyT(func(v LookupServerCertificateResult) string { return v.ExpirationDate }).(pulumi.StringOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupServerCertificateResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupServerCertificateResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupServerCertificateResultOutput) Latest() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupServerCertificateResult) *bool { return v.Latest }).(pulumi.BoolPtrOutput)
}

func (o LookupServerCertificateResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupServerCertificateResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupServerCertificateResultOutput) NamePrefix() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupServerCertificateResult) *string { return v.NamePrefix }).(pulumi.StringPtrOutput)
}

func (o LookupServerCertificateResultOutput) Path() pulumi.StringOutput {
	return o.ApplyT(func(v LookupServerCertificateResult) string { return v.Path }).(pulumi.StringOutput)
}

func (o LookupServerCertificateResultOutput) PathPrefix() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupServerCertificateResult) *string { return v.PathPrefix }).(pulumi.StringPtrOutput)
}

func (o LookupServerCertificateResultOutput) UploadDate() pulumi.StringOutput {
	return o.ApplyT(func(v LookupServerCertificateResult) string { return v.UploadDate }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupServerCertificateResultOutput{})
}
