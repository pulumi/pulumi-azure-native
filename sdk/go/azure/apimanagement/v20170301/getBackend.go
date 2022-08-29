


package v20170301

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupBackend(ctx *pulumi.Context, args *LookupBackendArgs, opts ...pulumi.InvokeOption) (*LookupBackendResult, error) {
	var rv LookupBackendResult
	err := ctx.Invoke("azure-native:apimanagement/v20170301:getBackend", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return rv.Defaults(), nil
}

type LookupBackendArgs struct {
	Backendid         string `pulumi:"backendid"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
	ServiceName       string `pulumi:"serviceName"`
}


type LookupBackendResult struct {
	Credentials *BackendCredentialsContractResponse `pulumi:"credentials"`
	Description *string                             `pulumi:"description"`
	Id          string                              `pulumi:"id"`
	Name        string                              `pulumi:"name"`
	Properties  BackendPropertiesResponse           `pulumi:"properties"`
	Protocol    string                              `pulumi:"protocol"`
	Proxy       *BackendProxyContractResponse       `pulumi:"proxy"`
	ResourceId  *string                             `pulumi:"resourceId"`
	Title       *string                             `pulumi:"title"`
	Tls         *BackendTlsPropertiesResponse       `pulumi:"tls"`
	Type        string                              `pulumi:"type"`
	Url         string                              `pulumi:"url"`
}


func (val *LookupBackendResult) Defaults() *LookupBackendResult {
	if val == nil {
		return nil
	}
	tmp := *val
	tmp.Tls = tmp.Tls.Defaults()

	return &tmp
}

func LookupBackendOutput(ctx *pulumi.Context, args LookupBackendOutputArgs, opts ...pulumi.InvokeOption) LookupBackendResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupBackendResult, error) {
			args := v.(LookupBackendArgs)
			r, err := LookupBackend(ctx, &args, opts...)
			var s LookupBackendResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupBackendResultOutput)
}

type LookupBackendOutputArgs struct {
	Backendid         pulumi.StringInput `pulumi:"backendid"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	ServiceName       pulumi.StringInput `pulumi:"serviceName"`
}

func (LookupBackendOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupBackendArgs)(nil)).Elem()
}


type LookupBackendResultOutput struct{ *pulumi.OutputState }

func (LookupBackendResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupBackendResult)(nil)).Elem()
}

func (o LookupBackendResultOutput) ToLookupBackendResultOutput() LookupBackendResultOutput {
	return o
}

func (o LookupBackendResultOutput) ToLookupBackendResultOutputWithContext(ctx context.Context) LookupBackendResultOutput {
	return o
}

func (o LookupBackendResultOutput) Credentials() BackendCredentialsContractResponsePtrOutput {
	return o.ApplyT(func(v LookupBackendResult) *BackendCredentialsContractResponse { return v.Credentials }).(BackendCredentialsContractResponsePtrOutput)
}

func (o LookupBackendResultOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupBackendResult) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o LookupBackendResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBackendResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupBackendResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBackendResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupBackendResultOutput) Properties() BackendPropertiesResponseOutput {
	return o.ApplyT(func(v LookupBackendResult) BackendPropertiesResponse { return v.Properties }).(BackendPropertiesResponseOutput)
}

func (o LookupBackendResultOutput) Protocol() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBackendResult) string { return v.Protocol }).(pulumi.StringOutput)
}

func (o LookupBackendResultOutput) Proxy() BackendProxyContractResponsePtrOutput {
	return o.ApplyT(func(v LookupBackendResult) *BackendProxyContractResponse { return v.Proxy }).(BackendProxyContractResponsePtrOutput)
}

func (o LookupBackendResultOutput) ResourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupBackendResult) *string { return v.ResourceId }).(pulumi.StringPtrOutput)
}

func (o LookupBackendResultOutput) Title() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupBackendResult) *string { return v.Title }).(pulumi.StringPtrOutput)
}

func (o LookupBackendResultOutput) Tls() BackendTlsPropertiesResponsePtrOutput {
	return o.ApplyT(func(v LookupBackendResult) *BackendTlsPropertiesResponse { return v.Tls }).(BackendTlsPropertiesResponsePtrOutput)
}

func (o LookupBackendResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBackendResult) string { return v.Type }).(pulumi.StringOutput)
}

func (o LookupBackendResultOutput) Url() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBackendResult) string { return v.Url }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupBackendResultOutput{})
}
