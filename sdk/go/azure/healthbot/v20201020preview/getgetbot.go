


package v20201020preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)


func Getgetbot(ctx *pulumi.Context, args *GetgetbotArgs, opts ...pulumi.InvokeOption) (*GetgetbotResult, error) {
	var rv GetgetbotResult
	err := ctx.Invoke("azure-native:healthbot/v20201020preview:getgetbot", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type GetgetbotArgs struct {
	ResourceGroupName string `pulumi:"resourceGroupName"`
	ResourceName      string `pulumi:"resourceName"`
}


type GetgetbotResult struct {
	Id         string                      `pulumi:"id"`
	Location   string                      `pulumi:"location"`
	Name       string                      `pulumi:"name"`
	Properties HealthBotPropertiesResponse `pulumi:"properties"`
	Sku        *SkuResponse                `pulumi:"sku"`
	SystemData SystemDataResponse          `pulumi:"systemData"`
	Tags       map[string]string           `pulumi:"tags"`
	Type       string                      `pulumi:"type"`
}

func GetgetbotOutput(ctx *pulumi.Context, args GetgetbotOutputArgs, opts ...pulumi.InvokeOption) GetgetbotResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetgetbotResult, error) {
			args := v.(GetgetbotArgs)
			r, err := Getgetbot(ctx, &args, opts...)
			var s GetgetbotResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetgetbotResultOutput)
}

type GetgetbotOutputArgs struct {
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	ResourceName      pulumi.StringInput `pulumi:"resourceName"`
}

func (GetgetbotOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetgetbotArgs)(nil)).Elem()
}


type GetgetbotResultOutput struct{ *pulumi.OutputState }

func (GetgetbotResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetgetbotResult)(nil)).Elem()
}

func (o GetgetbotResultOutput) ToGetgetbotResultOutput() GetgetbotResultOutput {
	return o
}

func (o GetgetbotResultOutput) ToGetgetbotResultOutputWithContext(ctx context.Context) GetgetbotResultOutput {
	return o
}

func (o GetgetbotResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetgetbotResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o GetgetbotResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v GetgetbotResult) string { return v.Location }).(pulumi.StringOutput)
}

func (o GetgetbotResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v GetgetbotResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o GetgetbotResultOutput) Properties() HealthBotPropertiesResponseOutput {
	return o.ApplyT(func(v GetgetbotResult) HealthBotPropertiesResponse { return v.Properties }).(HealthBotPropertiesResponseOutput)
}

func (o GetgetbotResultOutput) Sku() SkuResponsePtrOutput {
	return o.ApplyT(func(v GetgetbotResult) *SkuResponse { return v.Sku }).(SkuResponsePtrOutput)
}

func (o GetgetbotResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v GetgetbotResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

func (o GetgetbotResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v GetgetbotResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o GetgetbotResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v GetgetbotResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(GetgetbotResultOutput{})
}
