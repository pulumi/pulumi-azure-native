


package v20210630preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupReferenceDataSet(ctx *pulumi.Context, args *LookupReferenceDataSetArgs, opts ...pulumi.InvokeOption) (*LookupReferenceDataSetResult, error) {
	var rv LookupReferenceDataSetResult
	err := ctx.Invoke("azure-native:timeseriesinsights/v20210630preview:getReferenceDataSet", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupReferenceDataSetArgs struct {
	EnvironmentName      string `pulumi:"environmentName"`
	ReferenceDataSetName string `pulumi:"referenceDataSetName"`
	ResourceGroupName    string `pulumi:"resourceGroupName"`
}


type LookupReferenceDataSetResult struct {
	CreationTime                 string                                `pulumi:"creationTime"`
	DataStringComparisonBehavior *string                               `pulumi:"dataStringComparisonBehavior"`
	Id                           string                                `pulumi:"id"`
	KeyProperties                []ReferenceDataSetKeyPropertyResponse `pulumi:"keyProperties"`
	Location                     string                                `pulumi:"location"`
	Name                         string                                `pulumi:"name"`
	ProvisioningState            string                                `pulumi:"provisioningState"`
	Tags                         map[string]string                     `pulumi:"tags"`
	Type                         string                                `pulumi:"type"`
}

func LookupReferenceDataSetOutput(ctx *pulumi.Context, args LookupReferenceDataSetOutputArgs, opts ...pulumi.InvokeOption) LookupReferenceDataSetResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupReferenceDataSetResult, error) {
			args := v.(LookupReferenceDataSetArgs)
			r, err := LookupReferenceDataSet(ctx, &args, opts...)
			var s LookupReferenceDataSetResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupReferenceDataSetResultOutput)
}

type LookupReferenceDataSetOutputArgs struct {
	EnvironmentName      pulumi.StringInput `pulumi:"environmentName"`
	ReferenceDataSetName pulumi.StringInput `pulumi:"referenceDataSetName"`
	ResourceGroupName    pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupReferenceDataSetOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupReferenceDataSetArgs)(nil)).Elem()
}


type LookupReferenceDataSetResultOutput struct{ *pulumi.OutputState }

func (LookupReferenceDataSetResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupReferenceDataSetResult)(nil)).Elem()
}

func (o LookupReferenceDataSetResultOutput) ToLookupReferenceDataSetResultOutput() LookupReferenceDataSetResultOutput {
	return o
}

func (o LookupReferenceDataSetResultOutput) ToLookupReferenceDataSetResultOutputWithContext(ctx context.Context) LookupReferenceDataSetResultOutput {
	return o
}

func (o LookupReferenceDataSetResultOutput) CreationTime() pulumi.StringOutput {
	return o.ApplyT(func(v LookupReferenceDataSetResult) string { return v.CreationTime }).(pulumi.StringOutput)
}

func (o LookupReferenceDataSetResultOutput) DataStringComparisonBehavior() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupReferenceDataSetResult) *string { return v.DataStringComparisonBehavior }).(pulumi.StringPtrOutput)
}

func (o LookupReferenceDataSetResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupReferenceDataSetResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupReferenceDataSetResultOutput) KeyProperties() ReferenceDataSetKeyPropertyResponseArrayOutput {
	return o.ApplyT(func(v LookupReferenceDataSetResult) []ReferenceDataSetKeyPropertyResponse { return v.KeyProperties }).(ReferenceDataSetKeyPropertyResponseArrayOutput)
}

func (o LookupReferenceDataSetResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupReferenceDataSetResult) string { return v.Location }).(pulumi.StringOutput)
}

func (o LookupReferenceDataSetResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupReferenceDataSetResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupReferenceDataSetResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupReferenceDataSetResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o LookupReferenceDataSetResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupReferenceDataSetResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o LookupReferenceDataSetResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupReferenceDataSetResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupReferenceDataSetResultOutput{})
}
