


package security

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupIngestionSetting(ctx *pulumi.Context, args *LookupIngestionSettingArgs, opts ...pulumi.InvokeOption) (*LookupIngestionSettingResult, error) {
	var rv LookupIngestionSettingResult
	err := ctx.Invoke("azure-native:security:getIngestionSetting", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupIngestionSettingArgs struct {
	IngestionSettingName string `pulumi:"ingestionSettingName"`
}


type LookupIngestionSettingResult struct {
	Id   string `pulumi:"id"`
	Name string `pulumi:"name"`
	Type string `pulumi:"type"`
}

func LookupIngestionSettingOutput(ctx *pulumi.Context, args LookupIngestionSettingOutputArgs, opts ...pulumi.InvokeOption) LookupIngestionSettingResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupIngestionSettingResult, error) {
			args := v.(LookupIngestionSettingArgs)
			r, err := LookupIngestionSetting(ctx, &args, opts...)
			var s LookupIngestionSettingResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupIngestionSettingResultOutput)
}

type LookupIngestionSettingOutputArgs struct {
	IngestionSettingName pulumi.StringInput `pulumi:"ingestionSettingName"`
}

func (LookupIngestionSettingOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupIngestionSettingArgs)(nil)).Elem()
}


type LookupIngestionSettingResultOutput struct{ *pulumi.OutputState }

func (LookupIngestionSettingResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupIngestionSettingResult)(nil)).Elem()
}

func (o LookupIngestionSettingResultOutput) ToLookupIngestionSettingResultOutput() LookupIngestionSettingResultOutput {
	return o
}

func (o LookupIngestionSettingResultOutput) ToLookupIngestionSettingResultOutputWithContext(ctx context.Context) LookupIngestionSettingResultOutput {
	return o
}

func (o LookupIngestionSettingResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupIngestionSettingResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupIngestionSettingResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupIngestionSettingResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupIngestionSettingResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupIngestionSettingResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupIngestionSettingResultOutput{})
}
