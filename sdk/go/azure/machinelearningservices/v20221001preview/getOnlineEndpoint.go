


package v20221001preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupOnlineEndpoint(ctx *pulumi.Context, args *LookupOnlineEndpointArgs, opts ...pulumi.InvokeOption) (*LookupOnlineEndpointResult, error) {
	var rv LookupOnlineEndpointResult
	err := ctx.Invoke("azure-native:machinelearningservices/v20221001preview:getOnlineEndpoint", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return rv.Defaults(), nil
}

type LookupOnlineEndpointArgs struct {
	EndpointName      string `pulumi:"endpointName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
	WorkspaceName     string `pulumi:"workspaceName"`
}

type LookupOnlineEndpointResult struct {
	Id                       string                          `pulumi:"id"`
	Identity                 *ManagedServiceIdentityResponse `pulumi:"identity"`
	Kind                     *string                         `pulumi:"kind"`
	Location                 string                          `pulumi:"location"`
	Name                     string                          `pulumi:"name"`
	OnlineEndpointProperties OnlineEndpointResponse          `pulumi:"onlineEndpointProperties"`
	Sku                      *SkuResponse                    `pulumi:"sku"`
	SystemData               SystemDataResponse              `pulumi:"systemData"`
	Tags                     map[string]string               `pulumi:"tags"`
	Type                     string                          `pulumi:"type"`
}


func (val *LookupOnlineEndpointResult) Defaults() *LookupOnlineEndpointResult {
	if val == nil {
		return nil
	}
	tmp := *val
	tmp.OnlineEndpointProperties = *tmp.OnlineEndpointProperties.Defaults()

	return &tmp
}

func LookupOnlineEndpointOutput(ctx *pulumi.Context, args LookupOnlineEndpointOutputArgs, opts ...pulumi.InvokeOption) LookupOnlineEndpointResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupOnlineEndpointResult, error) {
			args := v.(LookupOnlineEndpointArgs)
			r, err := LookupOnlineEndpoint(ctx, &args, opts...)
			var s LookupOnlineEndpointResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupOnlineEndpointResultOutput)
}

type LookupOnlineEndpointOutputArgs struct {
	EndpointName      pulumi.StringInput `pulumi:"endpointName"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	WorkspaceName     pulumi.StringInput `pulumi:"workspaceName"`
}

func (LookupOnlineEndpointOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupOnlineEndpointArgs)(nil)).Elem()
}

type LookupOnlineEndpointResultOutput struct{ *pulumi.OutputState }

func (LookupOnlineEndpointResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupOnlineEndpointResult)(nil)).Elem()
}

func (o LookupOnlineEndpointResultOutput) ToLookupOnlineEndpointResultOutput() LookupOnlineEndpointResultOutput {
	return o
}

func (o LookupOnlineEndpointResultOutput) ToLookupOnlineEndpointResultOutputWithContext(ctx context.Context) LookupOnlineEndpointResultOutput {
	return o
}

func (o LookupOnlineEndpointResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupOnlineEndpointResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupOnlineEndpointResultOutput) Identity() ManagedServiceIdentityResponsePtrOutput {
	return o.ApplyT(func(v LookupOnlineEndpointResult) *ManagedServiceIdentityResponse { return v.Identity }).(ManagedServiceIdentityResponsePtrOutput)
}

func (o LookupOnlineEndpointResultOutput) Kind() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupOnlineEndpointResult) *string { return v.Kind }).(pulumi.StringPtrOutput)
}

func (o LookupOnlineEndpointResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupOnlineEndpointResult) string { return v.Location }).(pulumi.StringOutput)
}

func (o LookupOnlineEndpointResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupOnlineEndpointResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupOnlineEndpointResultOutput) OnlineEndpointProperties() OnlineEndpointResponseOutput {
	return o.ApplyT(func(v LookupOnlineEndpointResult) OnlineEndpointResponse { return v.OnlineEndpointProperties }).(OnlineEndpointResponseOutput)
}

func (o LookupOnlineEndpointResultOutput) Sku() SkuResponsePtrOutput {
	return o.ApplyT(func(v LookupOnlineEndpointResult) *SkuResponse { return v.Sku }).(SkuResponsePtrOutput)
}

func (o LookupOnlineEndpointResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupOnlineEndpointResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

func (o LookupOnlineEndpointResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupOnlineEndpointResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o LookupOnlineEndpointResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupOnlineEndpointResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupOnlineEndpointResultOutput{})
}
