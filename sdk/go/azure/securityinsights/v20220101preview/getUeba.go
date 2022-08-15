


package v20220101preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupUeba(ctx *pulumi.Context, args *LookupUebaArgs, opts ...pulumi.InvokeOption) (*LookupUebaResult, error) {
	var rv LookupUebaResult
	err := ctx.Invoke("azure-native:securityinsights/v20220101preview:getUeba", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupUebaArgs struct {
	ResourceGroupName string `pulumi:"resourceGroupName"`
	SettingsName      string `pulumi:"settingsName"`
	WorkspaceName     string `pulumi:"workspaceName"`
}


type LookupUebaResult struct {
	DataSources []string           `pulumi:"dataSources"`
	Etag        *string            `pulumi:"etag"`
	Id          string             `pulumi:"id"`
	Kind        string             `pulumi:"kind"`
	Name        string             `pulumi:"name"`
	SystemData  SystemDataResponse `pulumi:"systemData"`
	Type        string             `pulumi:"type"`
}

func LookupUebaOutput(ctx *pulumi.Context, args LookupUebaOutputArgs, opts ...pulumi.InvokeOption) LookupUebaResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupUebaResult, error) {
			args := v.(LookupUebaArgs)
			r, err := LookupUeba(ctx, &args, opts...)
			var s LookupUebaResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupUebaResultOutput)
}

type LookupUebaOutputArgs struct {
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	SettingsName      pulumi.StringInput `pulumi:"settingsName"`
	WorkspaceName     pulumi.StringInput `pulumi:"workspaceName"`
}

func (LookupUebaOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupUebaArgs)(nil)).Elem()
}


type LookupUebaResultOutput struct{ *pulumi.OutputState }

func (LookupUebaResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupUebaResult)(nil)).Elem()
}

func (o LookupUebaResultOutput) ToLookupUebaResultOutput() LookupUebaResultOutput {
	return o
}

func (o LookupUebaResultOutput) ToLookupUebaResultOutputWithContext(ctx context.Context) LookupUebaResultOutput {
	return o
}

func (o LookupUebaResultOutput) DataSources() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupUebaResult) []string { return v.DataSources }).(pulumi.StringArrayOutput)
}

func (o LookupUebaResultOutput) Etag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupUebaResult) *string { return v.Etag }).(pulumi.StringPtrOutput)
}

func (o LookupUebaResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupUebaResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupUebaResultOutput) Kind() pulumi.StringOutput {
	return o.ApplyT(func(v LookupUebaResult) string { return v.Kind }).(pulumi.StringOutput)
}

func (o LookupUebaResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupUebaResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupUebaResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupUebaResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

func (o LookupUebaResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupUebaResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupUebaResultOutput{})
}
