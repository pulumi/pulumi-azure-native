


package hybriddata

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupJobDefinition(ctx *pulumi.Context, args *LookupJobDefinitionArgs, opts ...pulumi.InvokeOption) (*LookupJobDefinitionResult, error) {
	var rv LookupJobDefinitionResult
	err := ctx.Invoke("azure-native:hybriddata:getJobDefinition", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return rv.Defaults(), nil
}

type LookupJobDefinitionArgs struct {
	DataManagerName   string `pulumi:"dataManagerName"`
	DataServiceName   string `pulumi:"dataServiceName"`
	JobDefinitionName string `pulumi:"jobDefinitionName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
}


type LookupJobDefinitionResult struct {
	CustomerSecrets  []CustomerSecretResponse `pulumi:"customerSecrets"`
	DataServiceInput interface{}              `pulumi:"dataServiceInput"`
	DataSinkId       string                   `pulumi:"dataSinkId"`
	DataSourceId     string                   `pulumi:"dataSourceId"`
	Id               string                   `pulumi:"id"`
	LastModifiedTime *string                  `pulumi:"lastModifiedTime"`
	Name             string                   `pulumi:"name"`
	RunLocation      *string                  `pulumi:"runLocation"`
	Schedules        []ScheduleResponse       `pulumi:"schedules"`
	State            string                   `pulumi:"state"`
	Type             string                   `pulumi:"type"`
	UserConfirmation *string                  `pulumi:"userConfirmation"`
}


func (val *LookupJobDefinitionResult) Defaults() *LookupJobDefinitionResult {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.UserConfirmation) {
		userConfirmation_ := "NotRequired"
		tmp.UserConfirmation = &userConfirmation_
	}
	return &tmp
}

func LookupJobDefinitionOutput(ctx *pulumi.Context, args LookupJobDefinitionOutputArgs, opts ...pulumi.InvokeOption) LookupJobDefinitionResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupJobDefinitionResult, error) {
			args := v.(LookupJobDefinitionArgs)
			r, err := LookupJobDefinition(ctx, &args, opts...)
			var s LookupJobDefinitionResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupJobDefinitionResultOutput)
}

type LookupJobDefinitionOutputArgs struct {
	DataManagerName   pulumi.StringInput `pulumi:"dataManagerName"`
	DataServiceName   pulumi.StringInput `pulumi:"dataServiceName"`
	JobDefinitionName pulumi.StringInput `pulumi:"jobDefinitionName"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupJobDefinitionOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupJobDefinitionArgs)(nil)).Elem()
}


type LookupJobDefinitionResultOutput struct{ *pulumi.OutputState }

func (LookupJobDefinitionResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupJobDefinitionResult)(nil)).Elem()
}

func (o LookupJobDefinitionResultOutput) ToLookupJobDefinitionResultOutput() LookupJobDefinitionResultOutput {
	return o
}

func (o LookupJobDefinitionResultOutput) ToLookupJobDefinitionResultOutputWithContext(ctx context.Context) LookupJobDefinitionResultOutput {
	return o
}

func (o LookupJobDefinitionResultOutput) CustomerSecrets() CustomerSecretResponseArrayOutput {
	return o.ApplyT(func(v LookupJobDefinitionResult) []CustomerSecretResponse { return v.CustomerSecrets }).(CustomerSecretResponseArrayOutput)
}

func (o LookupJobDefinitionResultOutput) DataServiceInput() pulumi.AnyOutput {
	return o.ApplyT(func(v LookupJobDefinitionResult) interface{} { return v.DataServiceInput }).(pulumi.AnyOutput)
}

func (o LookupJobDefinitionResultOutput) DataSinkId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupJobDefinitionResult) string { return v.DataSinkId }).(pulumi.StringOutput)
}

func (o LookupJobDefinitionResultOutput) DataSourceId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupJobDefinitionResult) string { return v.DataSourceId }).(pulumi.StringOutput)
}

func (o LookupJobDefinitionResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupJobDefinitionResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupJobDefinitionResultOutput) LastModifiedTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupJobDefinitionResult) *string { return v.LastModifiedTime }).(pulumi.StringPtrOutput)
}

func (o LookupJobDefinitionResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupJobDefinitionResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupJobDefinitionResultOutput) RunLocation() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupJobDefinitionResult) *string { return v.RunLocation }).(pulumi.StringPtrOutput)
}

func (o LookupJobDefinitionResultOutput) Schedules() ScheduleResponseArrayOutput {
	return o.ApplyT(func(v LookupJobDefinitionResult) []ScheduleResponse { return v.Schedules }).(ScheduleResponseArrayOutput)
}

func (o LookupJobDefinitionResultOutput) State() pulumi.StringOutput {
	return o.ApplyT(func(v LookupJobDefinitionResult) string { return v.State }).(pulumi.StringOutput)
}

func (o LookupJobDefinitionResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupJobDefinitionResult) string { return v.Type }).(pulumi.StringOutput)
}

func (o LookupJobDefinitionResultOutput) UserConfirmation() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupJobDefinitionResult) *string { return v.UserConfirmation }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupJobDefinitionResultOutput{})
}
