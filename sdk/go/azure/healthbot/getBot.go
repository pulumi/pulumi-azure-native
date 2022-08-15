


package healthbot

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupBot(ctx *pulumi.Context, args *LookupBotArgs, opts ...pulumi.InvokeOption) (*LookupBotResult, error) {
	var rv LookupBotResult
	err := ctx.Invoke("azure-native:healthbot:getBot", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupBotArgs struct {
	BotName           string `pulumi:"botName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
}


type LookupBotResult struct {
	Id         string                      `pulumi:"id"`
	Location   string                      `pulumi:"location"`
	Name       string                      `pulumi:"name"`
	Properties HealthBotPropertiesResponse `pulumi:"properties"`
	Sku        SkuResponse                 `pulumi:"sku"`
	SystemData SystemDataResponse          `pulumi:"systemData"`
	Tags       map[string]string           `pulumi:"tags"`
	Type       string                      `pulumi:"type"`
}

func LookupBotOutput(ctx *pulumi.Context, args LookupBotOutputArgs, opts ...pulumi.InvokeOption) LookupBotResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupBotResult, error) {
			args := v.(LookupBotArgs)
			r, err := LookupBot(ctx, &args, opts...)
			var s LookupBotResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupBotResultOutput)
}

type LookupBotOutputArgs struct {
	BotName           pulumi.StringInput `pulumi:"botName"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupBotOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupBotArgs)(nil)).Elem()
}


type LookupBotResultOutput struct{ *pulumi.OutputState }

func (LookupBotResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupBotResult)(nil)).Elem()
}

func (o LookupBotResultOutput) ToLookupBotResultOutput() LookupBotResultOutput {
	return o
}

func (o LookupBotResultOutput) ToLookupBotResultOutputWithContext(ctx context.Context) LookupBotResultOutput {
	return o
}

func (o LookupBotResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBotResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupBotResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBotResult) string { return v.Location }).(pulumi.StringOutput)
}

func (o LookupBotResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBotResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupBotResultOutput) Properties() HealthBotPropertiesResponseOutput {
	return o.ApplyT(func(v LookupBotResult) HealthBotPropertiesResponse { return v.Properties }).(HealthBotPropertiesResponseOutput)
}

func (o LookupBotResultOutput) Sku() SkuResponseOutput {
	return o.ApplyT(func(v LookupBotResult) SkuResponse { return v.Sku }).(SkuResponseOutput)
}

func (o LookupBotResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupBotResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

func (o LookupBotResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupBotResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o LookupBotResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBotResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupBotResultOutput{})
}
