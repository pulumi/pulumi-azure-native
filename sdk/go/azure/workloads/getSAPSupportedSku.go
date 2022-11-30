


package workloads

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func GetSAPSupportedSku(ctx *pulumi.Context, args *GetSAPSupportedSkuArgs, opts ...pulumi.InvokeOption) (*GetSAPSupportedSkuResult, error) {
	var rv GetSAPSupportedSkuResult
	err := ctx.Invoke("azure-native:workloads:getSAPSupportedSku", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type GetSAPSupportedSkuArgs struct {
	AppLocation          string  `pulumi:"appLocation"`
	DatabaseType         string  `pulumi:"databaseType"`
	DeploymentType       string  `pulumi:"deploymentType"`
	Environment          string  `pulumi:"environment"`
	HighAvailabilityType *string `pulumi:"highAvailabilityType"`
	Location             string  `pulumi:"location"`
	SapProduct           string  `pulumi:"sapProduct"`
}


type GetSAPSupportedSkuResult struct {
	SupportedSkus []SAPSupportedSkuResponse `pulumi:"supportedSkus"`
}

func GetSAPSupportedSkuOutput(ctx *pulumi.Context, args GetSAPSupportedSkuOutputArgs, opts ...pulumi.InvokeOption) GetSAPSupportedSkuResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetSAPSupportedSkuResult, error) {
			args := v.(GetSAPSupportedSkuArgs)
			r, err := GetSAPSupportedSku(ctx, &args, opts...)
			var s GetSAPSupportedSkuResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetSAPSupportedSkuResultOutput)
}

type GetSAPSupportedSkuOutputArgs struct {
	AppLocation          pulumi.StringInput    `pulumi:"appLocation"`
	DatabaseType         pulumi.StringInput    `pulumi:"databaseType"`
	DeploymentType       pulumi.StringInput    `pulumi:"deploymentType"`
	Environment          pulumi.StringInput    `pulumi:"environment"`
	HighAvailabilityType pulumi.StringPtrInput `pulumi:"highAvailabilityType"`
	Location             pulumi.StringInput    `pulumi:"location"`
	SapProduct           pulumi.StringInput    `pulumi:"sapProduct"`
}

func (GetSAPSupportedSkuOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetSAPSupportedSkuArgs)(nil)).Elem()
}


type GetSAPSupportedSkuResultOutput struct{ *pulumi.OutputState }

func (GetSAPSupportedSkuResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetSAPSupportedSkuResult)(nil)).Elem()
}

func (o GetSAPSupportedSkuResultOutput) ToGetSAPSupportedSkuResultOutput() GetSAPSupportedSkuResultOutput {
	return o
}

func (o GetSAPSupportedSkuResultOutput) ToGetSAPSupportedSkuResultOutputWithContext(ctx context.Context) GetSAPSupportedSkuResultOutput {
	return o
}

func (o GetSAPSupportedSkuResultOutput) SupportedSkus() SAPSupportedSkuResponseArrayOutput {
	return o.ApplyT(func(v GetSAPSupportedSkuResult) []SAPSupportedSkuResponse { return v.SupportedSkus }).(SAPSupportedSkuResponseArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(GetSAPSupportedSkuResultOutput{})
}
