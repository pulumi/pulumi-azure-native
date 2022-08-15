


package v20200701preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupConfigurationStore(ctx *pulumi.Context, args *LookupConfigurationStoreArgs, opts ...pulumi.InvokeOption) (*LookupConfigurationStoreResult, error) {
	var rv LookupConfigurationStoreResult
	err := ctx.Invoke("azure-native:appconfiguration/v20200701preview:getConfigurationStore", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupConfigurationStoreArgs struct {
	ConfigStoreName   string `pulumi:"configStoreName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
}


type LookupConfigurationStoreResult struct {
	CreationDate               string                                       `pulumi:"creationDate"`
	Encryption                 *EncryptionPropertiesResponse                `pulumi:"encryption"`
	Endpoint                   string                                       `pulumi:"endpoint"`
	Id                         string                                       `pulumi:"id"`
	Identity                   *ResourceIdentityResponse                    `pulumi:"identity"`
	Location                   string                                       `pulumi:"location"`
	Name                       string                                       `pulumi:"name"`
	PrivateEndpointConnections []PrivateEndpointConnectionReferenceResponse `pulumi:"privateEndpointConnections"`
	ProvisioningState          string                                       `pulumi:"provisioningState"`
	PublicNetworkAccess        *string                                      `pulumi:"publicNetworkAccess"`
	Sku                        SkuResponse                                  `pulumi:"sku"`
	Tags                       map[string]string                            `pulumi:"tags"`
	Type                       string                                       `pulumi:"type"`
}

func LookupConfigurationStoreOutput(ctx *pulumi.Context, args LookupConfigurationStoreOutputArgs, opts ...pulumi.InvokeOption) LookupConfigurationStoreResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupConfigurationStoreResult, error) {
			args := v.(LookupConfigurationStoreArgs)
			r, err := LookupConfigurationStore(ctx, &args, opts...)
			var s LookupConfigurationStoreResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupConfigurationStoreResultOutput)
}

type LookupConfigurationStoreOutputArgs struct {
	ConfigStoreName   pulumi.StringInput `pulumi:"configStoreName"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupConfigurationStoreOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupConfigurationStoreArgs)(nil)).Elem()
}


type LookupConfigurationStoreResultOutput struct{ *pulumi.OutputState }

func (LookupConfigurationStoreResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupConfigurationStoreResult)(nil)).Elem()
}

func (o LookupConfigurationStoreResultOutput) ToLookupConfigurationStoreResultOutput() LookupConfigurationStoreResultOutput {
	return o
}

func (o LookupConfigurationStoreResultOutput) ToLookupConfigurationStoreResultOutputWithContext(ctx context.Context) LookupConfigurationStoreResultOutput {
	return o
}

func (o LookupConfigurationStoreResultOutput) CreationDate() pulumi.StringOutput {
	return o.ApplyT(func(v LookupConfigurationStoreResult) string { return v.CreationDate }).(pulumi.StringOutput)
}

func (o LookupConfigurationStoreResultOutput) Encryption() EncryptionPropertiesResponsePtrOutput {
	return o.ApplyT(func(v LookupConfigurationStoreResult) *EncryptionPropertiesResponse { return v.Encryption }).(EncryptionPropertiesResponsePtrOutput)
}

func (o LookupConfigurationStoreResultOutput) Endpoint() pulumi.StringOutput {
	return o.ApplyT(func(v LookupConfigurationStoreResult) string { return v.Endpoint }).(pulumi.StringOutput)
}

func (o LookupConfigurationStoreResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupConfigurationStoreResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupConfigurationStoreResultOutput) Identity() ResourceIdentityResponsePtrOutput {
	return o.ApplyT(func(v LookupConfigurationStoreResult) *ResourceIdentityResponse { return v.Identity }).(ResourceIdentityResponsePtrOutput)
}

func (o LookupConfigurationStoreResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupConfigurationStoreResult) string { return v.Location }).(pulumi.StringOutput)
}

func (o LookupConfigurationStoreResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupConfigurationStoreResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupConfigurationStoreResultOutput) PrivateEndpointConnections() PrivateEndpointConnectionReferenceResponseArrayOutput {
	return o.ApplyT(func(v LookupConfigurationStoreResult) []PrivateEndpointConnectionReferenceResponse {
		return v.PrivateEndpointConnections
	}).(PrivateEndpointConnectionReferenceResponseArrayOutput)
}

func (o LookupConfigurationStoreResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupConfigurationStoreResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o LookupConfigurationStoreResultOutput) PublicNetworkAccess() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupConfigurationStoreResult) *string { return v.PublicNetworkAccess }).(pulumi.StringPtrOutput)
}

func (o LookupConfigurationStoreResultOutput) Sku() SkuResponseOutput {
	return o.ApplyT(func(v LookupConfigurationStoreResult) SkuResponse { return v.Sku }).(SkuResponseOutput)
}

func (o LookupConfigurationStoreResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupConfigurationStoreResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o LookupConfigurationStoreResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupConfigurationStoreResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupConfigurationStoreResultOutput{})
}
