


package v20150801preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupConnection(ctx *pulumi.Context, args *LookupConnectionArgs, opts ...pulumi.InvokeOption) (*LookupConnectionResult, error) {
	var rv LookupConnectionResult
	err := ctx.Invoke("azure-native:web/v20150801preview:getConnection", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupConnectionArgs struct {
	ConnectionName    string `pulumi:"connectionName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
}


type LookupConnectionResult struct {
	Api                      *ExpandedParentApiEntityResponse                     `pulumi:"api"`
	ChangedTime              *string                                              `pulumi:"changedTime"`
	CreatedTime              *string                                              `pulumi:"createdTime"`
	CustomParameterValues    map[string]ParameterCustomLoginSettingValuesResponse `pulumi:"customParameterValues"`
	DisplayName              *string                                              `pulumi:"displayName"`
	FirstExpirationTime      *string                                              `pulumi:"firstExpirationTime"`
	Id                       *string                                              `pulumi:"id"`
	Keywords                 []string                                             `pulumi:"keywords"`
	Kind                     *string                                              `pulumi:"kind"`
	Location                 string                                               `pulumi:"location"`
	Metadata                 interface{}                                          `pulumi:"metadata"`
	Name                     *string                                              `pulumi:"name"`
	NonSecretParameterValues map[string]interface{}                               `pulumi:"nonSecretParameterValues"`
	ParameterValues          map[string]interface{}                               `pulumi:"parameterValues"`
	Statuses                 []ConnectionStatusResponse                           `pulumi:"statuses"`
	Tags                     map[string]string                                    `pulumi:"tags"`
	TenantId                 *string                                              `pulumi:"tenantId"`
	Type                     *string                                              `pulumi:"type"`
}

func LookupConnectionOutput(ctx *pulumi.Context, args LookupConnectionOutputArgs, opts ...pulumi.InvokeOption) LookupConnectionResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupConnectionResult, error) {
			args := v.(LookupConnectionArgs)
			r, err := LookupConnection(ctx, &args, opts...)
			var s LookupConnectionResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupConnectionResultOutput)
}

type LookupConnectionOutputArgs struct {
	ConnectionName    pulumi.StringInput `pulumi:"connectionName"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupConnectionOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupConnectionArgs)(nil)).Elem()
}


type LookupConnectionResultOutput struct{ *pulumi.OutputState }

func (LookupConnectionResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupConnectionResult)(nil)).Elem()
}

func (o LookupConnectionResultOutput) ToLookupConnectionResultOutput() LookupConnectionResultOutput {
	return o
}

func (o LookupConnectionResultOutput) ToLookupConnectionResultOutputWithContext(ctx context.Context) LookupConnectionResultOutput {
	return o
}

func (o LookupConnectionResultOutput) Api() ExpandedParentApiEntityResponsePtrOutput {
	return o.ApplyT(func(v LookupConnectionResult) *ExpandedParentApiEntityResponse { return v.Api }).(ExpandedParentApiEntityResponsePtrOutput)
}

func (o LookupConnectionResultOutput) ChangedTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupConnectionResult) *string { return v.ChangedTime }).(pulumi.StringPtrOutput)
}

func (o LookupConnectionResultOutput) CreatedTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupConnectionResult) *string { return v.CreatedTime }).(pulumi.StringPtrOutput)
}

func (o LookupConnectionResultOutput) CustomParameterValues() ParameterCustomLoginSettingValuesResponseMapOutput {
	return o.ApplyT(func(v LookupConnectionResult) map[string]ParameterCustomLoginSettingValuesResponse {
		return v.CustomParameterValues
	}).(ParameterCustomLoginSettingValuesResponseMapOutput)
}

func (o LookupConnectionResultOutput) DisplayName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupConnectionResult) *string { return v.DisplayName }).(pulumi.StringPtrOutput)
}

func (o LookupConnectionResultOutput) FirstExpirationTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupConnectionResult) *string { return v.FirstExpirationTime }).(pulumi.StringPtrOutput)
}

func (o LookupConnectionResultOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupConnectionResult) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o LookupConnectionResultOutput) Keywords() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupConnectionResult) []string { return v.Keywords }).(pulumi.StringArrayOutput)
}

func (o LookupConnectionResultOutput) Kind() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupConnectionResult) *string { return v.Kind }).(pulumi.StringPtrOutput)
}

func (o LookupConnectionResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupConnectionResult) string { return v.Location }).(pulumi.StringOutput)
}

func (o LookupConnectionResultOutput) Metadata() pulumi.AnyOutput {
	return o.ApplyT(func(v LookupConnectionResult) interface{} { return v.Metadata }).(pulumi.AnyOutput)
}

func (o LookupConnectionResultOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupConnectionResult) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o LookupConnectionResultOutput) NonSecretParameterValues() pulumi.MapOutput {
	return o.ApplyT(func(v LookupConnectionResult) map[string]interface{} { return v.NonSecretParameterValues }).(pulumi.MapOutput)
}

func (o LookupConnectionResultOutput) ParameterValues() pulumi.MapOutput {
	return o.ApplyT(func(v LookupConnectionResult) map[string]interface{} { return v.ParameterValues }).(pulumi.MapOutput)
}

func (o LookupConnectionResultOutput) Statuses() ConnectionStatusResponseArrayOutput {
	return o.ApplyT(func(v LookupConnectionResult) []ConnectionStatusResponse { return v.Statuses }).(ConnectionStatusResponseArrayOutput)
}

func (o LookupConnectionResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupConnectionResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o LookupConnectionResultOutput) TenantId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupConnectionResult) *string { return v.TenantId }).(pulumi.StringPtrOutput)
}

func (o LookupConnectionResultOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupConnectionResult) *string { return v.Type }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupConnectionResultOutput{})
}
