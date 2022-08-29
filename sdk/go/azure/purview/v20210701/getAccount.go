


package v20210701

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupAccount(ctx *pulumi.Context, args *LookupAccountArgs, opts ...pulumi.InvokeOption) (*LookupAccountResult, error) {
	var rv LookupAccountResult
	err := ctx.Invoke("azure-native:purview/v20210701:getAccount", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return rv.Defaults(), nil
}

type LookupAccountArgs struct {
	AccountName       string `pulumi:"accountName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
}


type LookupAccountResult struct {
	CloudConnectors            *CloudConnectorsResponse                  `pulumi:"cloudConnectors"`
	CreatedAt                  string                                    `pulumi:"createdAt"`
	CreatedBy                  string                                    `pulumi:"createdBy"`
	CreatedByObjectId          string                                    `pulumi:"createdByObjectId"`
	Endpoints                  AccountPropertiesResponseEndpoints        `pulumi:"endpoints"`
	FriendlyName               string                                    `pulumi:"friendlyName"`
	Id                         string                                    `pulumi:"id"`
	Identity                   *IdentityResponse                         `pulumi:"identity"`
	Location                   *string                                   `pulumi:"location"`
	ManagedResourceGroupName   *string                                   `pulumi:"managedResourceGroupName"`
	ManagedResources           AccountPropertiesResponseManagedResources `pulumi:"managedResources"`
	Name                       string                                    `pulumi:"name"`
	PrivateEndpointConnections []PrivateEndpointConnectionResponse       `pulumi:"privateEndpointConnections"`
	ProvisioningState          string                                    `pulumi:"provisioningState"`
	PublicNetworkAccess        *string                                   `pulumi:"publicNetworkAccess"`
	Sku                        AccountResponseSku                        `pulumi:"sku"`
	SystemData                 TrackedResourceResponseSystemData         `pulumi:"systemData"`
	Tags                       map[string]string                         `pulumi:"tags"`
	Type                       string                                    `pulumi:"type"`
}


func (val *LookupAccountResult) Defaults() *LookupAccountResult {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.PublicNetworkAccess) {
		publicNetworkAccess_ := "Enabled"
		tmp.PublicNetworkAccess = &publicNetworkAccess_
	}
	return &tmp
}

func LookupAccountOutput(ctx *pulumi.Context, args LookupAccountOutputArgs, opts ...pulumi.InvokeOption) LookupAccountResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupAccountResult, error) {
			args := v.(LookupAccountArgs)
			r, err := LookupAccount(ctx, &args, opts...)
			var s LookupAccountResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupAccountResultOutput)
}

type LookupAccountOutputArgs struct {
	AccountName       pulumi.StringInput `pulumi:"accountName"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupAccountOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupAccountArgs)(nil)).Elem()
}


type LookupAccountResultOutput struct{ *pulumi.OutputState }

func (LookupAccountResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupAccountResult)(nil)).Elem()
}

func (o LookupAccountResultOutput) ToLookupAccountResultOutput() LookupAccountResultOutput {
	return o
}

func (o LookupAccountResultOutput) ToLookupAccountResultOutputWithContext(ctx context.Context) LookupAccountResultOutput {
	return o
}

func (o LookupAccountResultOutput) CloudConnectors() CloudConnectorsResponsePtrOutput {
	return o.ApplyT(func(v LookupAccountResult) *CloudConnectorsResponse { return v.CloudConnectors }).(CloudConnectorsResponsePtrOutput)
}

func (o LookupAccountResultOutput) CreatedAt() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAccountResult) string { return v.CreatedAt }).(pulumi.StringOutput)
}

func (o LookupAccountResultOutput) CreatedBy() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAccountResult) string { return v.CreatedBy }).(pulumi.StringOutput)
}

func (o LookupAccountResultOutput) CreatedByObjectId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAccountResult) string { return v.CreatedByObjectId }).(pulumi.StringOutput)
}

func (o LookupAccountResultOutput) Endpoints() AccountPropertiesResponseEndpointsOutput {
	return o.ApplyT(func(v LookupAccountResult) AccountPropertiesResponseEndpoints { return v.Endpoints }).(AccountPropertiesResponseEndpointsOutput)
}

func (o LookupAccountResultOutput) FriendlyName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAccountResult) string { return v.FriendlyName }).(pulumi.StringOutput)
}

func (o LookupAccountResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAccountResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupAccountResultOutput) Identity() IdentityResponsePtrOutput {
	return o.ApplyT(func(v LookupAccountResult) *IdentityResponse { return v.Identity }).(IdentityResponsePtrOutput)
}

func (o LookupAccountResultOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupAccountResult) *string { return v.Location }).(pulumi.StringPtrOutput)
}

func (o LookupAccountResultOutput) ManagedResourceGroupName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupAccountResult) *string { return v.ManagedResourceGroupName }).(pulumi.StringPtrOutput)
}

func (o LookupAccountResultOutput) ManagedResources() AccountPropertiesResponseManagedResourcesOutput {
	return o.ApplyT(func(v LookupAccountResult) AccountPropertiesResponseManagedResources { return v.ManagedResources }).(AccountPropertiesResponseManagedResourcesOutput)
}

func (o LookupAccountResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAccountResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupAccountResultOutput) PrivateEndpointConnections() PrivateEndpointConnectionResponseArrayOutput {
	return o.ApplyT(func(v LookupAccountResult) []PrivateEndpointConnectionResponse { return v.PrivateEndpointConnections }).(PrivateEndpointConnectionResponseArrayOutput)
}

func (o LookupAccountResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAccountResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o LookupAccountResultOutput) PublicNetworkAccess() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupAccountResult) *string { return v.PublicNetworkAccess }).(pulumi.StringPtrOutput)
}

func (o LookupAccountResultOutput) Sku() AccountResponseSkuOutput {
	return o.ApplyT(func(v LookupAccountResult) AccountResponseSku { return v.Sku }).(AccountResponseSkuOutput)
}

func (o LookupAccountResultOutput) SystemData() TrackedResourceResponseSystemDataOutput {
	return o.ApplyT(func(v LookupAccountResult) TrackedResourceResponseSystemData { return v.SystemData }).(TrackedResourceResponseSystemDataOutput)
}

func (o LookupAccountResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupAccountResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o LookupAccountResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAccountResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupAccountResultOutput{})
}
