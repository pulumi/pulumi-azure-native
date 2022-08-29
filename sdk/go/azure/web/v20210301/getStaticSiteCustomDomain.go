


package v20210301

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupStaticSiteCustomDomain(ctx *pulumi.Context, args *LookupStaticSiteCustomDomainArgs, opts ...pulumi.InvokeOption) (*LookupStaticSiteCustomDomainResult, error) {
	var rv LookupStaticSiteCustomDomainResult
	err := ctx.Invoke("azure-native:web/v20210301:getStaticSiteCustomDomain", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupStaticSiteCustomDomainArgs struct {
	DomainName        string `pulumi:"domainName"`
	Name              string `pulumi:"name"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
}


type LookupStaticSiteCustomDomainResult struct {
	CreatedOn       string  `pulumi:"createdOn"`
	DomainName      string  `pulumi:"domainName"`
	ErrorMessage    string  `pulumi:"errorMessage"`
	Id              string  `pulumi:"id"`
	Kind            *string `pulumi:"kind"`
	Name            string  `pulumi:"name"`
	Status          string  `pulumi:"status"`
	Type            string  `pulumi:"type"`
	ValidationToken string  `pulumi:"validationToken"`
}

func LookupStaticSiteCustomDomainOutput(ctx *pulumi.Context, args LookupStaticSiteCustomDomainOutputArgs, opts ...pulumi.InvokeOption) LookupStaticSiteCustomDomainResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupStaticSiteCustomDomainResult, error) {
			args := v.(LookupStaticSiteCustomDomainArgs)
			r, err := LookupStaticSiteCustomDomain(ctx, &args, opts...)
			var s LookupStaticSiteCustomDomainResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupStaticSiteCustomDomainResultOutput)
}

type LookupStaticSiteCustomDomainOutputArgs struct {
	DomainName        pulumi.StringInput `pulumi:"domainName"`
	Name              pulumi.StringInput `pulumi:"name"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupStaticSiteCustomDomainOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupStaticSiteCustomDomainArgs)(nil)).Elem()
}


type LookupStaticSiteCustomDomainResultOutput struct{ *pulumi.OutputState }

func (LookupStaticSiteCustomDomainResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupStaticSiteCustomDomainResult)(nil)).Elem()
}

func (o LookupStaticSiteCustomDomainResultOutput) ToLookupStaticSiteCustomDomainResultOutput() LookupStaticSiteCustomDomainResultOutput {
	return o
}

func (o LookupStaticSiteCustomDomainResultOutput) ToLookupStaticSiteCustomDomainResultOutputWithContext(ctx context.Context) LookupStaticSiteCustomDomainResultOutput {
	return o
}

func (o LookupStaticSiteCustomDomainResultOutput) CreatedOn() pulumi.StringOutput {
	return o.ApplyT(func(v LookupStaticSiteCustomDomainResult) string { return v.CreatedOn }).(pulumi.StringOutput)
}

func (o LookupStaticSiteCustomDomainResultOutput) DomainName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupStaticSiteCustomDomainResult) string { return v.DomainName }).(pulumi.StringOutput)
}

func (o LookupStaticSiteCustomDomainResultOutput) ErrorMessage() pulumi.StringOutput {
	return o.ApplyT(func(v LookupStaticSiteCustomDomainResult) string { return v.ErrorMessage }).(pulumi.StringOutput)
}

func (o LookupStaticSiteCustomDomainResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupStaticSiteCustomDomainResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupStaticSiteCustomDomainResultOutput) Kind() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupStaticSiteCustomDomainResult) *string { return v.Kind }).(pulumi.StringPtrOutput)
}

func (o LookupStaticSiteCustomDomainResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupStaticSiteCustomDomainResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupStaticSiteCustomDomainResultOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v LookupStaticSiteCustomDomainResult) string { return v.Status }).(pulumi.StringOutput)
}

func (o LookupStaticSiteCustomDomainResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupStaticSiteCustomDomainResult) string { return v.Type }).(pulumi.StringOutput)
}

func (o LookupStaticSiteCustomDomainResultOutput) ValidationToken() pulumi.StringOutput {
	return o.ApplyT(func(v LookupStaticSiteCustomDomainResult) string { return v.ValidationToken }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupStaticSiteCustomDomainResultOutput{})
}
