


package v20220401preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupLogger(ctx *pulumi.Context, args *LookupLoggerArgs, opts ...pulumi.InvokeOption) (*LookupLoggerResult, error) {
	var rv LookupLoggerResult
	err := ctx.Invoke("azure-native:apimanagement/v20220401preview:getLogger", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupLoggerArgs struct {
	LoggerId          string `pulumi:"loggerId"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
	ServiceName       string `pulumi:"serviceName"`
}


type LookupLoggerResult struct {
	Credentials map[string]string `pulumi:"credentials"`
	Description *string           `pulumi:"description"`
	Id          string            `pulumi:"id"`
	IsBuffered  *bool             `pulumi:"isBuffered"`
	LoggerType  string            `pulumi:"loggerType"`
	Name        string            `pulumi:"name"`
	ResourceId  *string           `pulumi:"resourceId"`
	Type        string            `pulumi:"type"`
}

func LookupLoggerOutput(ctx *pulumi.Context, args LookupLoggerOutputArgs, opts ...pulumi.InvokeOption) LookupLoggerResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupLoggerResult, error) {
			args := v.(LookupLoggerArgs)
			r, err := LookupLogger(ctx, &args, opts...)
			var s LookupLoggerResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupLoggerResultOutput)
}

type LookupLoggerOutputArgs struct {
	LoggerId          pulumi.StringInput `pulumi:"loggerId"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	ServiceName       pulumi.StringInput `pulumi:"serviceName"`
}

func (LookupLoggerOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupLoggerArgs)(nil)).Elem()
}


type LookupLoggerResultOutput struct{ *pulumi.OutputState }

func (LookupLoggerResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupLoggerResult)(nil)).Elem()
}

func (o LookupLoggerResultOutput) ToLookupLoggerResultOutput() LookupLoggerResultOutput {
	return o
}

func (o LookupLoggerResultOutput) ToLookupLoggerResultOutputWithContext(ctx context.Context) LookupLoggerResultOutput {
	return o
}

func (o LookupLoggerResultOutput) Credentials() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupLoggerResult) map[string]string { return v.Credentials }).(pulumi.StringMapOutput)
}

func (o LookupLoggerResultOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupLoggerResult) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o LookupLoggerResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupLoggerResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupLoggerResultOutput) IsBuffered() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupLoggerResult) *bool { return v.IsBuffered }).(pulumi.BoolPtrOutput)
}

func (o LookupLoggerResultOutput) LoggerType() pulumi.StringOutput {
	return o.ApplyT(func(v LookupLoggerResult) string { return v.LoggerType }).(pulumi.StringOutput)
}

func (o LookupLoggerResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupLoggerResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupLoggerResultOutput) ResourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupLoggerResult) *string { return v.ResourceId }).(pulumi.StringPtrOutput)
}

func (o LookupLoggerResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupLoggerResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupLoggerResultOutput{})
}
