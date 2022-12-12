


package v20220801preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupWebPubSubCustomDomain(ctx *pulumi.Context, args *LookupWebPubSubCustomDomainArgs, opts ...pulumi.InvokeOption) (*LookupWebPubSubCustomDomainResult, error) {
	var rv LookupWebPubSubCustomDomainResult
	err := ctx.Invoke("azure-native:webpubsub/v20220801preview:getWebPubSubCustomDomain", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupWebPubSubCustomDomainArgs struct {
	Name              string `pulumi:"name"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
	ResourceName      string `pulumi:"resourceName"`
}


type LookupWebPubSubCustomDomainResult struct {
	CustomCertificate ResourceReferenceResponse `pulumi:"customCertificate"`
	DomainName        string                    `pulumi:"domainName"`
	Id                string                    `pulumi:"id"`
	Name              string                    `pulumi:"name"`
	ProvisioningState string                    `pulumi:"provisioningState"`
	SystemData        SystemDataResponse        `pulumi:"systemData"`
	Type              string                    `pulumi:"type"`
}

func LookupWebPubSubCustomDomainOutput(ctx *pulumi.Context, args LookupWebPubSubCustomDomainOutputArgs, opts ...pulumi.InvokeOption) LookupWebPubSubCustomDomainResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupWebPubSubCustomDomainResult, error) {
			args := v.(LookupWebPubSubCustomDomainArgs)
			r, err := LookupWebPubSubCustomDomain(ctx, &args, opts...)
			var s LookupWebPubSubCustomDomainResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupWebPubSubCustomDomainResultOutput)
}

type LookupWebPubSubCustomDomainOutputArgs struct {
	Name              pulumi.StringInput `pulumi:"name"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	ResourceName      pulumi.StringInput `pulumi:"resourceName"`
}

func (LookupWebPubSubCustomDomainOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupWebPubSubCustomDomainArgs)(nil)).Elem()
}


type LookupWebPubSubCustomDomainResultOutput struct{ *pulumi.OutputState }

func (LookupWebPubSubCustomDomainResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupWebPubSubCustomDomainResult)(nil)).Elem()
}

func (o LookupWebPubSubCustomDomainResultOutput) ToLookupWebPubSubCustomDomainResultOutput() LookupWebPubSubCustomDomainResultOutput {
	return o
}

func (o LookupWebPubSubCustomDomainResultOutput) ToLookupWebPubSubCustomDomainResultOutputWithContext(ctx context.Context) LookupWebPubSubCustomDomainResultOutput {
	return o
}

func (o LookupWebPubSubCustomDomainResultOutput) CustomCertificate() ResourceReferenceResponseOutput {
	return o.ApplyT(func(v LookupWebPubSubCustomDomainResult) ResourceReferenceResponse { return v.CustomCertificate }).(ResourceReferenceResponseOutput)
}

func (o LookupWebPubSubCustomDomainResultOutput) DomainName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWebPubSubCustomDomainResult) string { return v.DomainName }).(pulumi.StringOutput)
}

func (o LookupWebPubSubCustomDomainResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWebPubSubCustomDomainResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupWebPubSubCustomDomainResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWebPubSubCustomDomainResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupWebPubSubCustomDomainResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWebPubSubCustomDomainResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o LookupWebPubSubCustomDomainResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupWebPubSubCustomDomainResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

func (o LookupWebPubSubCustomDomainResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWebPubSubCustomDomainResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupWebPubSubCustomDomainResultOutput{})
}
