


package v20220301

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func ListWebAppHostKeys(ctx *pulumi.Context, args *ListWebAppHostKeysArgs, opts ...pulumi.InvokeOption) (*ListWebAppHostKeysResult, error) {
	var rv ListWebAppHostKeysResult
	err := ctx.Invoke("azure-native:web/v20220301:listWebAppHostKeys", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListWebAppHostKeysArgs struct {
	Name              string `pulumi:"name"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
}


type ListWebAppHostKeysResult struct {
	FunctionKeys map[string]string `pulumi:"functionKeys"`
	MasterKey    *string           `pulumi:"masterKey"`
	SystemKeys   map[string]string `pulumi:"systemKeys"`
}

func ListWebAppHostKeysOutput(ctx *pulumi.Context, args ListWebAppHostKeysOutputArgs, opts ...pulumi.InvokeOption) ListWebAppHostKeysResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (ListWebAppHostKeysResult, error) {
			args := v.(ListWebAppHostKeysArgs)
			r, err := ListWebAppHostKeys(ctx, &args, opts...)
			var s ListWebAppHostKeysResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(ListWebAppHostKeysResultOutput)
}

type ListWebAppHostKeysOutputArgs struct {
	Name              pulumi.StringInput `pulumi:"name"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (ListWebAppHostKeysOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ListWebAppHostKeysArgs)(nil)).Elem()
}


type ListWebAppHostKeysResultOutput struct{ *pulumi.OutputState }

func (ListWebAppHostKeysResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ListWebAppHostKeysResult)(nil)).Elem()
}

func (o ListWebAppHostKeysResultOutput) ToListWebAppHostKeysResultOutput() ListWebAppHostKeysResultOutput {
	return o
}

func (o ListWebAppHostKeysResultOutput) ToListWebAppHostKeysResultOutputWithContext(ctx context.Context) ListWebAppHostKeysResultOutput {
	return o
}

func (o ListWebAppHostKeysResultOutput) FunctionKeys() pulumi.StringMapOutput {
	return o.ApplyT(func(v ListWebAppHostKeysResult) map[string]string { return v.FunctionKeys }).(pulumi.StringMapOutput)
}

func (o ListWebAppHostKeysResultOutput) MasterKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ListWebAppHostKeysResult) *string { return v.MasterKey }).(pulumi.StringPtrOutput)
}

func (o ListWebAppHostKeysResultOutput) SystemKeys() pulumi.StringMapOutput {
	return o.ApplyT(func(v ListWebAppHostKeysResult) map[string]string { return v.SystemKeys }).(pulumi.StringMapOutput)
}

func init() {
	pulumi.RegisterOutputType(ListWebAppHostKeysResultOutput{})
}
