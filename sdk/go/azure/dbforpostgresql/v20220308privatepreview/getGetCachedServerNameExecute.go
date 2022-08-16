


package v20220308privatepreview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func GetGetCachedServerNameExecute(ctx *pulumi.Context, args *GetGetCachedServerNameExecuteArgs, opts ...pulumi.InvokeOption) (*GetGetCachedServerNameExecuteResult, error) {
	var rv GetGetCachedServerNameExecuteResult
	err := ctx.Invoke("azure-native:dbforpostgresql/v20220308privatepreview:getGetCachedServerNameExecute", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type GetGetCachedServerNameExecuteArgs struct {
	LocationName      string  `pulumi:"locationName"`
	ResourceGroupName string  `pulumi:"resourceGroupName"`
	Sku               Sku     `pulumi:"sku"`
	Storage           Storage `pulumi:"storage"`
	Version           string  `pulumi:"version"`
}


type GetGetCachedServerNameExecuteResult struct {
	Name string `pulumi:"name"`
}

func GetGetCachedServerNameExecuteOutput(ctx *pulumi.Context, args GetGetCachedServerNameExecuteOutputArgs, opts ...pulumi.InvokeOption) GetGetCachedServerNameExecuteResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetGetCachedServerNameExecuteResult, error) {
			args := v.(GetGetCachedServerNameExecuteArgs)
			r, err := GetGetCachedServerNameExecute(ctx, &args, opts...)
			var s GetGetCachedServerNameExecuteResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetGetCachedServerNameExecuteResultOutput)
}

type GetGetCachedServerNameExecuteOutputArgs struct {
	LocationName      pulumi.StringInput `pulumi:"locationName"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	Sku               SkuInput           `pulumi:"sku"`
	Storage           StorageInput       `pulumi:"storage"`
	Version           pulumi.StringInput `pulumi:"version"`
}

func (GetGetCachedServerNameExecuteOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetGetCachedServerNameExecuteArgs)(nil)).Elem()
}


type GetGetCachedServerNameExecuteResultOutput struct{ *pulumi.OutputState }

func (GetGetCachedServerNameExecuteResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetGetCachedServerNameExecuteResult)(nil)).Elem()
}

func (o GetGetCachedServerNameExecuteResultOutput) ToGetGetCachedServerNameExecuteResultOutput() GetGetCachedServerNameExecuteResultOutput {
	return o
}

func (o GetGetCachedServerNameExecuteResultOutput) ToGetGetCachedServerNameExecuteResultOutputWithContext(ctx context.Context) GetGetCachedServerNameExecuteResultOutput {
	return o
}

func (o GetGetCachedServerNameExecuteResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v GetGetCachedServerNameExecuteResult) string { return v.Name }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(GetGetCachedServerNameExecuteResultOutput{})
}
