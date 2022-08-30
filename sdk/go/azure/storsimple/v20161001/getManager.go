


package v20161001

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)


func LookupManager(ctx *pulumi.Context, args *LookupManagerArgs, opts ...pulumi.InvokeOption) (*LookupManagerResult, error) {
	var rv LookupManagerResult
	err := ctx.Invoke("azure-native:storsimple/v20161001:getManager", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupManagerArgs struct {
	ManagerName       string `pulumi:"managerName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
}


type LookupManagerResult struct {
	CisIntrinsicSettings *ManagerIntrinsicSettingsResponse `pulumi:"cisIntrinsicSettings"`
	Etag                 *string                           `pulumi:"etag"`
	Id                   string                            `pulumi:"id"`
	Location             string                            `pulumi:"location"`
	Name                 string                            `pulumi:"name"`
	ProvisioningState    string                            `pulumi:"provisioningState"`
	Sku                  *ManagerSkuResponse               `pulumi:"sku"`
	Tags                 map[string]string                 `pulumi:"tags"`
	Type                 string                            `pulumi:"type"`
}

func LookupManagerOutput(ctx *pulumi.Context, args LookupManagerOutputArgs, opts ...pulumi.InvokeOption) LookupManagerResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupManagerResult, error) {
			args := v.(LookupManagerArgs)
			r, err := LookupManager(ctx, &args, opts...)
			var s LookupManagerResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupManagerResultOutput)
}

type LookupManagerOutputArgs struct {
	ManagerName       pulumi.StringInput `pulumi:"managerName"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupManagerOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupManagerArgs)(nil)).Elem()
}


type LookupManagerResultOutput struct{ *pulumi.OutputState }

func (LookupManagerResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupManagerResult)(nil)).Elem()
}

func (o LookupManagerResultOutput) ToLookupManagerResultOutput() LookupManagerResultOutput {
	return o
}

func (o LookupManagerResultOutput) ToLookupManagerResultOutputWithContext(ctx context.Context) LookupManagerResultOutput {
	return o
}

func (o LookupManagerResultOutput) CisIntrinsicSettings() ManagerIntrinsicSettingsResponsePtrOutput {
	return o.ApplyT(func(v LookupManagerResult) *ManagerIntrinsicSettingsResponse { return v.CisIntrinsicSettings }).(ManagerIntrinsicSettingsResponsePtrOutput)
}

func (o LookupManagerResultOutput) Etag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupManagerResult) *string { return v.Etag }).(pulumi.StringPtrOutput)
}

func (o LookupManagerResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupManagerResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupManagerResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupManagerResult) string { return v.Location }).(pulumi.StringOutput)
}

func (o LookupManagerResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupManagerResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupManagerResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupManagerResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o LookupManagerResultOutput) Sku() ManagerSkuResponsePtrOutput {
	return o.ApplyT(func(v LookupManagerResult) *ManagerSkuResponse { return v.Sku }).(ManagerSkuResponsePtrOutput)
}

func (o LookupManagerResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupManagerResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o LookupManagerResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupManagerResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupManagerResultOutput{})
}
