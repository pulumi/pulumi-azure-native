


package v20210801

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)


func LookupDataSetMapping(ctx *pulumi.Context, args *LookupDataSetMappingArgs, opts ...pulumi.InvokeOption) (*LookupDataSetMappingResult, error) {
	var rv LookupDataSetMappingResult
	err := ctx.Invoke("azure-native:datashare/v20210801:getDataSetMapping", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupDataSetMappingArgs struct {
	AccountName           string `pulumi:"accountName"`
	DataSetMappingName    string `pulumi:"dataSetMappingName"`
	ResourceGroupName     string `pulumi:"resourceGroupName"`
	ShareSubscriptionName string `pulumi:"shareSubscriptionName"`
}


type LookupDataSetMappingResult struct {
	Id         string             `pulumi:"id"`
	Kind       string             `pulumi:"kind"`
	Name       string             `pulumi:"name"`
	SystemData SystemDataResponse `pulumi:"systemData"`
	Type       string             `pulumi:"type"`
}

func LookupDataSetMappingOutput(ctx *pulumi.Context, args LookupDataSetMappingOutputArgs, opts ...pulumi.InvokeOption) LookupDataSetMappingResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupDataSetMappingResult, error) {
			args := v.(LookupDataSetMappingArgs)
			r, err := LookupDataSetMapping(ctx, &args, opts...)
			var s LookupDataSetMappingResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupDataSetMappingResultOutput)
}

type LookupDataSetMappingOutputArgs struct {
	AccountName           pulumi.StringInput `pulumi:"accountName"`
	DataSetMappingName    pulumi.StringInput `pulumi:"dataSetMappingName"`
	ResourceGroupName     pulumi.StringInput `pulumi:"resourceGroupName"`
	ShareSubscriptionName pulumi.StringInput `pulumi:"shareSubscriptionName"`
}

func (LookupDataSetMappingOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupDataSetMappingArgs)(nil)).Elem()
}


type LookupDataSetMappingResultOutput struct{ *pulumi.OutputState }

func (LookupDataSetMappingResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupDataSetMappingResult)(nil)).Elem()
}

func (o LookupDataSetMappingResultOutput) ToLookupDataSetMappingResultOutput() LookupDataSetMappingResultOutput {
	return o
}

func (o LookupDataSetMappingResultOutput) ToLookupDataSetMappingResultOutputWithContext(ctx context.Context) LookupDataSetMappingResultOutput {
	return o
}

func (o LookupDataSetMappingResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDataSetMappingResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupDataSetMappingResultOutput) Kind() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDataSetMappingResult) string { return v.Kind }).(pulumi.StringOutput)
}

func (o LookupDataSetMappingResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDataSetMappingResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupDataSetMappingResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupDataSetMappingResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

func (o LookupDataSetMappingResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDataSetMappingResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupDataSetMappingResultOutput{})
}
