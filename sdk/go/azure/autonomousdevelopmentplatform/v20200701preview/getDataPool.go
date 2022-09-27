


package v20200701preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)


func LookupDataPool(ctx *pulumi.Context, args *LookupDataPoolArgs, opts ...pulumi.InvokeOption) (*LookupDataPoolResult, error) {
	var rv LookupDataPoolResult
	err := ctx.Invoke("azure-native:autonomousdevelopmentplatform/v20200701preview:getDataPool", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupDataPoolArgs struct {
	AccountName       string `pulumi:"accountName"`
	DataPoolName      string `pulumi:"dataPoolName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
}


type LookupDataPoolResult struct {
	DataPoolId        string                     `pulumi:"dataPoolId"`
	Id                string                     `pulumi:"id"`
	Locations         []DataPoolLocationResponse `pulumi:"locations"`
	Name              string                     `pulumi:"name"`
	ProvisioningState string                     `pulumi:"provisioningState"`
	SystemData        SystemDataResponse         `pulumi:"systemData"`
	Type              string                     `pulumi:"type"`
}

func LookupDataPoolOutput(ctx *pulumi.Context, args LookupDataPoolOutputArgs, opts ...pulumi.InvokeOption) LookupDataPoolResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupDataPoolResult, error) {
			args := v.(LookupDataPoolArgs)
			r, err := LookupDataPool(ctx, &args, opts...)
			var s LookupDataPoolResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupDataPoolResultOutput)
}

type LookupDataPoolOutputArgs struct {
	AccountName       pulumi.StringInput `pulumi:"accountName"`
	DataPoolName      pulumi.StringInput `pulumi:"dataPoolName"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupDataPoolOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupDataPoolArgs)(nil)).Elem()
}


type LookupDataPoolResultOutput struct{ *pulumi.OutputState }

func (LookupDataPoolResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupDataPoolResult)(nil)).Elem()
}

func (o LookupDataPoolResultOutput) ToLookupDataPoolResultOutput() LookupDataPoolResultOutput {
	return o
}

func (o LookupDataPoolResultOutput) ToLookupDataPoolResultOutputWithContext(ctx context.Context) LookupDataPoolResultOutput {
	return o
}

func (o LookupDataPoolResultOutput) DataPoolId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDataPoolResult) string { return v.DataPoolId }).(pulumi.StringOutput)
}

func (o LookupDataPoolResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDataPoolResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupDataPoolResultOutput) Locations() DataPoolLocationResponseArrayOutput {
	return o.ApplyT(func(v LookupDataPoolResult) []DataPoolLocationResponse { return v.Locations }).(DataPoolLocationResponseArrayOutput)
}

func (o LookupDataPoolResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDataPoolResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupDataPoolResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDataPoolResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o LookupDataPoolResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupDataPoolResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

func (o LookupDataPoolResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDataPoolResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupDataPoolResultOutput{})
}
