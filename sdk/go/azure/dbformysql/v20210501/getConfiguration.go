


package v20210501

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupConfiguration(ctx *pulumi.Context, args *LookupConfigurationArgs, opts ...pulumi.InvokeOption) (*LookupConfigurationResult, error) {
	var rv LookupConfigurationResult
	err := ctx.Invoke("azure-native:dbformysql/v20210501:getConfiguration", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupConfigurationArgs struct {
	ConfigurationName string `pulumi:"configurationName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
	ServerName        string `pulumi:"serverName"`
}


type LookupConfigurationResult struct {
	AllowedValues          string             `pulumi:"allowedValues"`
	DataType               string             `pulumi:"dataType"`
	DefaultValue           string             `pulumi:"defaultValue"`
	Description            string             `pulumi:"description"`
	Id                     string             `pulumi:"id"`
	IsConfigPendingRestart string             `pulumi:"isConfigPendingRestart"`
	IsDynamicConfig        string             `pulumi:"isDynamicConfig"`
	IsReadOnly             string             `pulumi:"isReadOnly"`
	Name                   string             `pulumi:"name"`
	Source                 *string            `pulumi:"source"`
	SystemData             SystemDataResponse `pulumi:"systemData"`
	Type                   string             `pulumi:"type"`
	Value                  *string            `pulumi:"value"`
}

func LookupConfigurationOutput(ctx *pulumi.Context, args LookupConfigurationOutputArgs, opts ...pulumi.InvokeOption) LookupConfigurationResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupConfigurationResult, error) {
			args := v.(LookupConfigurationArgs)
			r, err := LookupConfiguration(ctx, &args, opts...)
			var s LookupConfigurationResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupConfigurationResultOutput)
}

type LookupConfigurationOutputArgs struct {
	ConfigurationName pulumi.StringInput `pulumi:"configurationName"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	ServerName        pulumi.StringInput `pulumi:"serverName"`
}

func (LookupConfigurationOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupConfigurationArgs)(nil)).Elem()
}


type LookupConfigurationResultOutput struct{ *pulumi.OutputState }

func (LookupConfigurationResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupConfigurationResult)(nil)).Elem()
}

func (o LookupConfigurationResultOutput) ToLookupConfigurationResultOutput() LookupConfigurationResultOutput {
	return o
}

func (o LookupConfigurationResultOutput) ToLookupConfigurationResultOutputWithContext(ctx context.Context) LookupConfigurationResultOutput {
	return o
}

func (o LookupConfigurationResultOutput) AllowedValues() pulumi.StringOutput {
	return o.ApplyT(func(v LookupConfigurationResult) string { return v.AllowedValues }).(pulumi.StringOutput)
}

func (o LookupConfigurationResultOutput) DataType() pulumi.StringOutput {
	return o.ApplyT(func(v LookupConfigurationResult) string { return v.DataType }).(pulumi.StringOutput)
}

func (o LookupConfigurationResultOutput) DefaultValue() pulumi.StringOutput {
	return o.ApplyT(func(v LookupConfigurationResult) string { return v.DefaultValue }).(pulumi.StringOutput)
}

func (o LookupConfigurationResultOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v LookupConfigurationResult) string { return v.Description }).(pulumi.StringOutput)
}

func (o LookupConfigurationResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupConfigurationResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupConfigurationResultOutput) IsConfigPendingRestart() pulumi.StringOutput {
	return o.ApplyT(func(v LookupConfigurationResult) string { return v.IsConfigPendingRestart }).(pulumi.StringOutput)
}

func (o LookupConfigurationResultOutput) IsDynamicConfig() pulumi.StringOutput {
	return o.ApplyT(func(v LookupConfigurationResult) string { return v.IsDynamicConfig }).(pulumi.StringOutput)
}

func (o LookupConfigurationResultOutput) IsReadOnly() pulumi.StringOutput {
	return o.ApplyT(func(v LookupConfigurationResult) string { return v.IsReadOnly }).(pulumi.StringOutput)
}

func (o LookupConfigurationResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupConfigurationResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupConfigurationResultOutput) Source() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupConfigurationResult) *string { return v.Source }).(pulumi.StringPtrOutput)
}

func (o LookupConfigurationResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupConfigurationResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

func (o LookupConfigurationResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupConfigurationResult) string { return v.Type }).(pulumi.StringOutput)
}

func (o LookupConfigurationResultOutput) Value() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupConfigurationResult) *string { return v.Value }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupConfigurationResultOutput{})
}
