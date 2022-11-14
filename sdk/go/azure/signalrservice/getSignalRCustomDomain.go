


package signalrservice

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupSignalRCustomDomain(ctx *pulumi.Context, args *LookupSignalRCustomDomainArgs, opts ...pulumi.InvokeOption) (*LookupSignalRCustomDomainResult, error) {
	var rv LookupSignalRCustomDomainResult
	err := ctx.Invoke("azure-native:signalrservice:getSignalRCustomDomain", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupSignalRCustomDomainArgs struct {
	Name              string `pulumi:"name"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
	ResourceName      string `pulumi:"resourceName"`
}


type LookupSignalRCustomDomainResult struct {
	CustomCertificate ResourceReferenceResponse `pulumi:"customCertificate"`
	DomainName        string                    `pulumi:"domainName"`
	Id                string                    `pulumi:"id"`
	Name              string                    `pulumi:"name"`
	ProvisioningState string                    `pulumi:"provisioningState"`
	SystemData        SystemDataResponse        `pulumi:"systemData"`
	Type              string                    `pulumi:"type"`
}

func LookupSignalRCustomDomainOutput(ctx *pulumi.Context, args LookupSignalRCustomDomainOutputArgs, opts ...pulumi.InvokeOption) LookupSignalRCustomDomainResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupSignalRCustomDomainResult, error) {
			args := v.(LookupSignalRCustomDomainArgs)
			r, err := LookupSignalRCustomDomain(ctx, &args, opts...)
			var s LookupSignalRCustomDomainResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupSignalRCustomDomainResultOutput)
}

type LookupSignalRCustomDomainOutputArgs struct {
	Name              pulumi.StringInput `pulumi:"name"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	ResourceName      pulumi.StringInput `pulumi:"resourceName"`
}

func (LookupSignalRCustomDomainOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSignalRCustomDomainArgs)(nil)).Elem()
}


type LookupSignalRCustomDomainResultOutput struct{ *pulumi.OutputState }

func (LookupSignalRCustomDomainResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSignalRCustomDomainResult)(nil)).Elem()
}

func (o LookupSignalRCustomDomainResultOutput) ToLookupSignalRCustomDomainResultOutput() LookupSignalRCustomDomainResultOutput {
	return o
}

func (o LookupSignalRCustomDomainResultOutput) ToLookupSignalRCustomDomainResultOutputWithContext(ctx context.Context) LookupSignalRCustomDomainResultOutput {
	return o
}

func (o LookupSignalRCustomDomainResultOutput) CustomCertificate() ResourceReferenceResponseOutput {
	return o.ApplyT(func(v LookupSignalRCustomDomainResult) ResourceReferenceResponse { return v.CustomCertificate }).(ResourceReferenceResponseOutput)
}

func (o LookupSignalRCustomDomainResultOutput) DomainName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSignalRCustomDomainResult) string { return v.DomainName }).(pulumi.StringOutput)
}

func (o LookupSignalRCustomDomainResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSignalRCustomDomainResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupSignalRCustomDomainResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSignalRCustomDomainResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupSignalRCustomDomainResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSignalRCustomDomainResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o LookupSignalRCustomDomainResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupSignalRCustomDomainResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

func (o LookupSignalRCustomDomainResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSignalRCustomDomainResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupSignalRCustomDomainResultOutput{})
}
