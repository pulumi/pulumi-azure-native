


package v20200901

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func ListWebAppHostKeysSlot(ctx *pulumi.Context, args *ListWebAppHostKeysSlotArgs, opts ...pulumi.InvokeOption) (*ListWebAppHostKeysSlotResult, error) {
	var rv ListWebAppHostKeysSlotResult
	err := ctx.Invoke("azure-native:web/v20200901:listWebAppHostKeysSlot", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListWebAppHostKeysSlotArgs struct {
	Name              string `pulumi:"name"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
	Slot              string `pulumi:"slot"`
}


type ListWebAppHostKeysSlotResult struct {
	FunctionKeys map[string]string `pulumi:"functionKeys"`
	MasterKey    *string           `pulumi:"masterKey"`
	SystemKeys   map[string]string `pulumi:"systemKeys"`
}

func ListWebAppHostKeysSlotOutput(ctx *pulumi.Context, args ListWebAppHostKeysSlotOutputArgs, opts ...pulumi.InvokeOption) ListWebAppHostKeysSlotResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (ListWebAppHostKeysSlotResult, error) {
			args := v.(ListWebAppHostKeysSlotArgs)
			r, err := ListWebAppHostKeysSlot(ctx, &args, opts...)
			var s ListWebAppHostKeysSlotResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(ListWebAppHostKeysSlotResultOutput)
}

type ListWebAppHostKeysSlotOutputArgs struct {
	Name              pulumi.StringInput `pulumi:"name"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	Slot              pulumi.StringInput `pulumi:"slot"`
}

func (ListWebAppHostKeysSlotOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ListWebAppHostKeysSlotArgs)(nil)).Elem()
}


type ListWebAppHostKeysSlotResultOutput struct{ *pulumi.OutputState }

func (ListWebAppHostKeysSlotResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ListWebAppHostKeysSlotResult)(nil)).Elem()
}

func (o ListWebAppHostKeysSlotResultOutput) ToListWebAppHostKeysSlotResultOutput() ListWebAppHostKeysSlotResultOutput {
	return o
}

func (o ListWebAppHostKeysSlotResultOutput) ToListWebAppHostKeysSlotResultOutputWithContext(ctx context.Context) ListWebAppHostKeysSlotResultOutput {
	return o
}

func (o ListWebAppHostKeysSlotResultOutput) FunctionKeys() pulumi.StringMapOutput {
	return o.ApplyT(func(v ListWebAppHostKeysSlotResult) map[string]string { return v.FunctionKeys }).(pulumi.StringMapOutput)
}

func (o ListWebAppHostKeysSlotResultOutput) MasterKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ListWebAppHostKeysSlotResult) *string { return v.MasterKey }).(pulumi.StringPtrOutput)
}

func (o ListWebAppHostKeysSlotResultOutput) SystemKeys() pulumi.StringMapOutput {
	return o.ApplyT(func(v ListWebAppHostKeysSlotResult) map[string]string { return v.SystemKeys }).(pulumi.StringMapOutput)
}

func init() {
	pulumi.RegisterOutputType(ListWebAppHostKeysSlotResultOutput{})
}
