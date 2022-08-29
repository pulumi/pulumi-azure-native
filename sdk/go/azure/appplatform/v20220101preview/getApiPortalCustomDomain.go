


package v20220101preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupApiPortalCustomDomain(ctx *pulumi.Context, args *LookupApiPortalCustomDomainArgs, opts ...pulumi.InvokeOption) (*LookupApiPortalCustomDomainResult, error) {
	var rv LookupApiPortalCustomDomainResult
	err := ctx.Invoke("azure-native:appplatform/v20220101preview:getApiPortalCustomDomain", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupApiPortalCustomDomainArgs struct {
	ApiPortalName     string `pulumi:"apiPortalName"`
	DomainName        string `pulumi:"domainName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
	ServiceName       string `pulumi:"serviceName"`
}


type LookupApiPortalCustomDomainResult struct {
	Id         string                                  `pulumi:"id"`
	Name       string                                  `pulumi:"name"`
	Properties ApiPortalCustomDomainPropertiesResponse `pulumi:"properties"`
	SystemData SystemDataResponse                      `pulumi:"systemData"`
	Type       string                                  `pulumi:"type"`
}

func LookupApiPortalCustomDomainOutput(ctx *pulumi.Context, args LookupApiPortalCustomDomainOutputArgs, opts ...pulumi.InvokeOption) LookupApiPortalCustomDomainResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupApiPortalCustomDomainResult, error) {
			args := v.(LookupApiPortalCustomDomainArgs)
			r, err := LookupApiPortalCustomDomain(ctx, &args, opts...)
			var s LookupApiPortalCustomDomainResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupApiPortalCustomDomainResultOutput)
}

type LookupApiPortalCustomDomainOutputArgs struct {
	ApiPortalName     pulumi.StringInput `pulumi:"apiPortalName"`
	DomainName        pulumi.StringInput `pulumi:"domainName"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	ServiceName       pulumi.StringInput `pulumi:"serviceName"`
}

func (LookupApiPortalCustomDomainOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupApiPortalCustomDomainArgs)(nil)).Elem()
}


type LookupApiPortalCustomDomainResultOutput struct{ *pulumi.OutputState }

func (LookupApiPortalCustomDomainResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupApiPortalCustomDomainResult)(nil)).Elem()
}

func (o LookupApiPortalCustomDomainResultOutput) ToLookupApiPortalCustomDomainResultOutput() LookupApiPortalCustomDomainResultOutput {
	return o
}

func (o LookupApiPortalCustomDomainResultOutput) ToLookupApiPortalCustomDomainResultOutputWithContext(ctx context.Context) LookupApiPortalCustomDomainResultOutput {
	return o
}

func (o LookupApiPortalCustomDomainResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupApiPortalCustomDomainResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupApiPortalCustomDomainResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupApiPortalCustomDomainResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupApiPortalCustomDomainResultOutput) Properties() ApiPortalCustomDomainPropertiesResponseOutput {
	return o.ApplyT(func(v LookupApiPortalCustomDomainResult) ApiPortalCustomDomainPropertiesResponse { return v.Properties }).(ApiPortalCustomDomainPropertiesResponseOutput)
}

func (o LookupApiPortalCustomDomainResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupApiPortalCustomDomainResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

func (o LookupApiPortalCustomDomainResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupApiPortalCustomDomainResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupApiPortalCustomDomainResultOutput{})
}
