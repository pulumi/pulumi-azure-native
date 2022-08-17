


package v20200601

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupDomainTopic(ctx *pulumi.Context, args *LookupDomainTopicArgs, opts ...pulumi.InvokeOption) (*LookupDomainTopicResult, error) {
	var rv LookupDomainTopicResult
	err := ctx.Invoke("azure-native:eventgrid/v20200601:getDomainTopic", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupDomainTopicArgs struct {
	DomainName        string `pulumi:"domainName"`
	DomainTopicName   string `pulumi:"domainTopicName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
}


type LookupDomainTopicResult struct {
	Id                string             `pulumi:"id"`
	Name              string             `pulumi:"name"`
	ProvisioningState string             `pulumi:"provisioningState"`
	SystemData        SystemDataResponse `pulumi:"systemData"`
	Type              string             `pulumi:"type"`
}

func LookupDomainTopicOutput(ctx *pulumi.Context, args LookupDomainTopicOutputArgs, opts ...pulumi.InvokeOption) LookupDomainTopicResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupDomainTopicResult, error) {
			args := v.(LookupDomainTopicArgs)
			r, err := LookupDomainTopic(ctx, &args, opts...)
			var s LookupDomainTopicResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupDomainTopicResultOutput)
}

type LookupDomainTopicOutputArgs struct {
	DomainName        pulumi.StringInput `pulumi:"domainName"`
	DomainTopicName   pulumi.StringInput `pulumi:"domainTopicName"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupDomainTopicOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupDomainTopicArgs)(nil)).Elem()
}


type LookupDomainTopicResultOutput struct{ *pulumi.OutputState }

func (LookupDomainTopicResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupDomainTopicResult)(nil)).Elem()
}

func (o LookupDomainTopicResultOutput) ToLookupDomainTopicResultOutput() LookupDomainTopicResultOutput {
	return o
}

func (o LookupDomainTopicResultOutput) ToLookupDomainTopicResultOutputWithContext(ctx context.Context) LookupDomainTopicResultOutput {
	return o
}

func (o LookupDomainTopicResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDomainTopicResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupDomainTopicResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDomainTopicResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupDomainTopicResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDomainTopicResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o LookupDomainTopicResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupDomainTopicResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

func (o LookupDomainTopicResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDomainTopicResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupDomainTopicResultOutput{})
}
