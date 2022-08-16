


package logic

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupIntegrationAccountPartner(ctx *pulumi.Context, args *LookupIntegrationAccountPartnerArgs, opts ...pulumi.InvokeOption) (*LookupIntegrationAccountPartnerResult, error) {
	var rv LookupIntegrationAccountPartnerResult
	err := ctx.Invoke("azure-native:logic:getIntegrationAccountPartner", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupIntegrationAccountPartnerArgs struct {
	IntegrationAccountName string `pulumi:"integrationAccountName"`
	PartnerName            string `pulumi:"partnerName"`
	ResourceGroupName      string `pulumi:"resourceGroupName"`
}


type LookupIntegrationAccountPartnerResult struct {
	ChangedTime string                 `pulumi:"changedTime"`
	Content     PartnerContentResponse `pulumi:"content"`
	CreatedTime string                 `pulumi:"createdTime"`
	Id          string                 `pulumi:"id"`
	Location    *string                `pulumi:"location"`
	Metadata    interface{}            `pulumi:"metadata"`
	Name        string                 `pulumi:"name"`
	PartnerType string                 `pulumi:"partnerType"`
	Tags        map[string]string      `pulumi:"tags"`
	Type        string                 `pulumi:"type"`
}

func LookupIntegrationAccountPartnerOutput(ctx *pulumi.Context, args LookupIntegrationAccountPartnerOutputArgs, opts ...pulumi.InvokeOption) LookupIntegrationAccountPartnerResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupIntegrationAccountPartnerResult, error) {
			args := v.(LookupIntegrationAccountPartnerArgs)
			r, err := LookupIntegrationAccountPartner(ctx, &args, opts...)
			var s LookupIntegrationAccountPartnerResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupIntegrationAccountPartnerResultOutput)
}

type LookupIntegrationAccountPartnerOutputArgs struct {
	IntegrationAccountName pulumi.StringInput `pulumi:"integrationAccountName"`
	PartnerName            pulumi.StringInput `pulumi:"partnerName"`
	ResourceGroupName      pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupIntegrationAccountPartnerOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupIntegrationAccountPartnerArgs)(nil)).Elem()
}


type LookupIntegrationAccountPartnerResultOutput struct{ *pulumi.OutputState }

func (LookupIntegrationAccountPartnerResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupIntegrationAccountPartnerResult)(nil)).Elem()
}

func (o LookupIntegrationAccountPartnerResultOutput) ToLookupIntegrationAccountPartnerResultOutput() LookupIntegrationAccountPartnerResultOutput {
	return o
}

func (o LookupIntegrationAccountPartnerResultOutput) ToLookupIntegrationAccountPartnerResultOutputWithContext(ctx context.Context) LookupIntegrationAccountPartnerResultOutput {
	return o
}

func (o LookupIntegrationAccountPartnerResultOutput) ChangedTime() pulumi.StringOutput {
	return o.ApplyT(func(v LookupIntegrationAccountPartnerResult) string { return v.ChangedTime }).(pulumi.StringOutput)
}

func (o LookupIntegrationAccountPartnerResultOutput) Content() PartnerContentResponseOutput {
	return o.ApplyT(func(v LookupIntegrationAccountPartnerResult) PartnerContentResponse { return v.Content }).(PartnerContentResponseOutput)
}

func (o LookupIntegrationAccountPartnerResultOutput) CreatedTime() pulumi.StringOutput {
	return o.ApplyT(func(v LookupIntegrationAccountPartnerResult) string { return v.CreatedTime }).(pulumi.StringOutput)
}

func (o LookupIntegrationAccountPartnerResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupIntegrationAccountPartnerResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupIntegrationAccountPartnerResultOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupIntegrationAccountPartnerResult) *string { return v.Location }).(pulumi.StringPtrOutput)
}

func (o LookupIntegrationAccountPartnerResultOutput) Metadata() pulumi.AnyOutput {
	return o.ApplyT(func(v LookupIntegrationAccountPartnerResult) interface{} { return v.Metadata }).(pulumi.AnyOutput)
}

func (o LookupIntegrationAccountPartnerResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupIntegrationAccountPartnerResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupIntegrationAccountPartnerResultOutput) PartnerType() pulumi.StringOutput {
	return o.ApplyT(func(v LookupIntegrationAccountPartnerResult) string { return v.PartnerType }).(pulumi.StringOutput)
}

func (o LookupIntegrationAccountPartnerResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupIntegrationAccountPartnerResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o LookupIntegrationAccountPartnerResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupIntegrationAccountPartnerResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupIntegrationAccountPartnerResultOutput{})
}
