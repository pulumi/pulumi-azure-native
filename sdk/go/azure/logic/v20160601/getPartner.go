


package v20160601

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupPartner(ctx *pulumi.Context, args *LookupPartnerArgs, opts ...pulumi.InvokeOption) (*LookupPartnerResult, error) {
	var rv LookupPartnerResult
	err := ctx.Invoke("azure-native:logic/v20160601:getPartner", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupPartnerArgs struct {
	IntegrationAccountName string `pulumi:"integrationAccountName"`
	PartnerName            string `pulumi:"partnerName"`
	ResourceGroupName      string `pulumi:"resourceGroupName"`
}


type LookupPartnerResult struct {
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

func LookupPartnerOutput(ctx *pulumi.Context, args LookupPartnerOutputArgs, opts ...pulumi.InvokeOption) LookupPartnerResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupPartnerResult, error) {
			args := v.(LookupPartnerArgs)
			r, err := LookupPartner(ctx, &args, opts...)
			var s LookupPartnerResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupPartnerResultOutput)
}

type LookupPartnerOutputArgs struct {
	IntegrationAccountName pulumi.StringInput `pulumi:"integrationAccountName"`
	PartnerName            pulumi.StringInput `pulumi:"partnerName"`
	ResourceGroupName      pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupPartnerOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupPartnerArgs)(nil)).Elem()
}


type LookupPartnerResultOutput struct{ *pulumi.OutputState }

func (LookupPartnerResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupPartnerResult)(nil)).Elem()
}

func (o LookupPartnerResultOutput) ToLookupPartnerResultOutput() LookupPartnerResultOutput {
	return o
}

func (o LookupPartnerResultOutput) ToLookupPartnerResultOutputWithContext(ctx context.Context) LookupPartnerResultOutput {
	return o
}

func (o LookupPartnerResultOutput) ChangedTime() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPartnerResult) string { return v.ChangedTime }).(pulumi.StringOutput)
}

func (o LookupPartnerResultOutput) Content() PartnerContentResponseOutput {
	return o.ApplyT(func(v LookupPartnerResult) PartnerContentResponse { return v.Content }).(PartnerContentResponseOutput)
}

func (o LookupPartnerResultOutput) CreatedTime() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPartnerResult) string { return v.CreatedTime }).(pulumi.StringOutput)
}

func (o LookupPartnerResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPartnerResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupPartnerResultOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupPartnerResult) *string { return v.Location }).(pulumi.StringPtrOutput)
}

func (o LookupPartnerResultOutput) Metadata() pulumi.AnyOutput {
	return o.ApplyT(func(v LookupPartnerResult) interface{} { return v.Metadata }).(pulumi.AnyOutput)
}

func (o LookupPartnerResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPartnerResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupPartnerResultOutput) PartnerType() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPartnerResult) string { return v.PartnerType }).(pulumi.StringOutput)
}

func (o LookupPartnerResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupPartnerResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o LookupPartnerResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPartnerResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupPartnerResultOutput{})
}
