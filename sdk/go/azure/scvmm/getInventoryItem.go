


package scvmm

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupInventoryItem(ctx *pulumi.Context, args *LookupInventoryItemArgs, opts ...pulumi.InvokeOption) (*LookupInventoryItemResult, error) {
	var rv LookupInventoryItemResult
	err := ctx.Invoke("azure-native:scvmm:getInventoryItem", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupInventoryItemArgs struct {
	InventoryItemName string `pulumi:"inventoryItemName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
	VmmServerName     string `pulumi:"vmmServerName"`
}


type LookupInventoryItemResult struct {
	Id                string             `pulumi:"id"`
	InventoryItemName string             `pulumi:"inventoryItemName"`
	InventoryType     string             `pulumi:"inventoryType"`
	Kind              *string            `pulumi:"kind"`
	ManagedResourceId string             `pulumi:"managedResourceId"`
	Name              string             `pulumi:"name"`
	ProvisioningState string             `pulumi:"provisioningState"`
	SystemData        SystemDataResponse `pulumi:"systemData"`
	Type              string             `pulumi:"type"`
	Uuid              string             `pulumi:"uuid"`
}

func LookupInventoryItemOutput(ctx *pulumi.Context, args LookupInventoryItemOutputArgs, opts ...pulumi.InvokeOption) LookupInventoryItemResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupInventoryItemResult, error) {
			args := v.(LookupInventoryItemArgs)
			r, err := LookupInventoryItem(ctx, &args, opts...)
			var s LookupInventoryItemResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupInventoryItemResultOutput)
}

type LookupInventoryItemOutputArgs struct {
	InventoryItemName pulumi.StringInput `pulumi:"inventoryItemName"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	VmmServerName     pulumi.StringInput `pulumi:"vmmServerName"`
}

func (LookupInventoryItemOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupInventoryItemArgs)(nil)).Elem()
}


type LookupInventoryItemResultOutput struct{ *pulumi.OutputState }

func (LookupInventoryItemResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupInventoryItemResult)(nil)).Elem()
}

func (o LookupInventoryItemResultOutput) ToLookupInventoryItemResultOutput() LookupInventoryItemResultOutput {
	return o
}

func (o LookupInventoryItemResultOutput) ToLookupInventoryItemResultOutputWithContext(ctx context.Context) LookupInventoryItemResultOutput {
	return o
}

func (o LookupInventoryItemResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupInventoryItemResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupInventoryItemResultOutput) InventoryItemName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupInventoryItemResult) string { return v.InventoryItemName }).(pulumi.StringOutput)
}

func (o LookupInventoryItemResultOutput) InventoryType() pulumi.StringOutput {
	return o.ApplyT(func(v LookupInventoryItemResult) string { return v.InventoryType }).(pulumi.StringOutput)
}

func (o LookupInventoryItemResultOutput) Kind() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupInventoryItemResult) *string { return v.Kind }).(pulumi.StringPtrOutput)
}

func (o LookupInventoryItemResultOutput) ManagedResourceId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupInventoryItemResult) string { return v.ManagedResourceId }).(pulumi.StringOutput)
}

func (o LookupInventoryItemResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupInventoryItemResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupInventoryItemResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupInventoryItemResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o LookupInventoryItemResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupInventoryItemResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

func (o LookupInventoryItemResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupInventoryItemResult) string { return v.Type }).(pulumi.StringOutput)
}

func (o LookupInventoryItemResultOutput) Uuid() pulumi.StringOutput {
	return o.ApplyT(func(v LookupInventoryItemResult) string { return v.Uuid }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupInventoryItemResultOutput{})
}
