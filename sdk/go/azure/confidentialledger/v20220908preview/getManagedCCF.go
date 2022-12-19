


package v20220908preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupManagedCCF(ctx *pulumi.Context, args *LookupManagedCCFArgs, opts ...pulumi.InvokeOption) (*LookupManagedCCFResult, error) {
	var rv LookupManagedCCFResult
	err := ctx.Invoke("azure-native:confidentialledger/v20220908preview:getManagedCCF", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupManagedCCFArgs struct {
	AppName           string `pulumi:"appName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
}


type LookupManagedCCFResult struct {
	Id         string                       `pulumi:"id"`
	Location   string                       `pulumi:"location"`
	Name       string                       `pulumi:"name"`
	Properties ManagedCCFPropertiesResponse `pulumi:"properties"`
	SystemData SystemDataResponse           `pulumi:"systemData"`
	Tags       map[string]string            `pulumi:"tags"`
	Type       string                       `pulumi:"type"`
}

func LookupManagedCCFOutput(ctx *pulumi.Context, args LookupManagedCCFOutputArgs, opts ...pulumi.InvokeOption) LookupManagedCCFResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupManagedCCFResult, error) {
			args := v.(LookupManagedCCFArgs)
			r, err := LookupManagedCCF(ctx, &args, opts...)
			var s LookupManagedCCFResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupManagedCCFResultOutput)
}

type LookupManagedCCFOutputArgs struct {
	AppName           pulumi.StringInput `pulumi:"appName"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupManagedCCFOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupManagedCCFArgs)(nil)).Elem()
}


type LookupManagedCCFResultOutput struct{ *pulumi.OutputState }

func (LookupManagedCCFResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupManagedCCFResult)(nil)).Elem()
}

func (o LookupManagedCCFResultOutput) ToLookupManagedCCFResultOutput() LookupManagedCCFResultOutput {
	return o
}

func (o LookupManagedCCFResultOutput) ToLookupManagedCCFResultOutputWithContext(ctx context.Context) LookupManagedCCFResultOutput {
	return o
}

func (o LookupManagedCCFResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupManagedCCFResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupManagedCCFResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupManagedCCFResult) string { return v.Location }).(pulumi.StringOutput)
}

func (o LookupManagedCCFResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupManagedCCFResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupManagedCCFResultOutput) Properties() ManagedCCFPropertiesResponseOutput {
	return o.ApplyT(func(v LookupManagedCCFResult) ManagedCCFPropertiesResponse { return v.Properties }).(ManagedCCFPropertiesResponseOutput)
}

func (o LookupManagedCCFResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupManagedCCFResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

func (o LookupManagedCCFResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupManagedCCFResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o LookupManagedCCFResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupManagedCCFResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupManagedCCFResultOutput{})
}
