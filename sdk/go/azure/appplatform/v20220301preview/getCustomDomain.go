


package v20220301preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupCustomDomain(ctx *pulumi.Context, args *LookupCustomDomainArgs, opts ...pulumi.InvokeOption) (*LookupCustomDomainResult, error) {
	var rv LookupCustomDomainResult
	err := ctx.Invoke("azure-native:appplatform/v20220301preview:getCustomDomain", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupCustomDomainArgs struct {
	AppName           string `pulumi:"appName"`
	DomainName        string `pulumi:"domainName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
	ServiceName       string `pulumi:"serviceName"`
}


type LookupCustomDomainResult struct {
	Id         string                         `pulumi:"id"`
	Name       string                         `pulumi:"name"`
	Properties CustomDomainPropertiesResponse `pulumi:"properties"`
	SystemData SystemDataResponse             `pulumi:"systemData"`
	Type       string                         `pulumi:"type"`
}

func LookupCustomDomainOutput(ctx *pulumi.Context, args LookupCustomDomainOutputArgs, opts ...pulumi.InvokeOption) LookupCustomDomainResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupCustomDomainResult, error) {
			args := v.(LookupCustomDomainArgs)
			r, err := LookupCustomDomain(ctx, &args, opts...)
			var s LookupCustomDomainResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupCustomDomainResultOutput)
}

type LookupCustomDomainOutputArgs struct {
	AppName           pulumi.StringInput `pulumi:"appName"`
	DomainName        pulumi.StringInput `pulumi:"domainName"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	ServiceName       pulumi.StringInput `pulumi:"serviceName"`
}

func (LookupCustomDomainOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupCustomDomainArgs)(nil)).Elem()
}


type LookupCustomDomainResultOutput struct{ *pulumi.OutputState }

func (LookupCustomDomainResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupCustomDomainResult)(nil)).Elem()
}

func (o LookupCustomDomainResultOutput) ToLookupCustomDomainResultOutput() LookupCustomDomainResultOutput {
	return o
}

func (o LookupCustomDomainResultOutput) ToLookupCustomDomainResultOutputWithContext(ctx context.Context) LookupCustomDomainResultOutput {
	return o
}

func (o LookupCustomDomainResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCustomDomainResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupCustomDomainResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCustomDomainResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupCustomDomainResultOutput) Properties() CustomDomainPropertiesResponseOutput {
	return o.ApplyT(func(v LookupCustomDomainResult) CustomDomainPropertiesResponse { return v.Properties }).(CustomDomainPropertiesResponseOutput)
}

func (o LookupCustomDomainResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupCustomDomainResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

func (o LookupCustomDomainResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCustomDomainResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupCustomDomainResultOutput{})
}
