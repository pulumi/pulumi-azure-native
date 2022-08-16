


package v20201001

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupCluster(ctx *pulumi.Context, args *LookupClusterArgs, opts ...pulumi.InvokeOption) (*LookupClusterResult, error) {
	var rv LookupClusterResult
	err := ctx.Invoke("azure-native:operationalinsights/v20201001:getCluster", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupClusterArgs struct {
	ClusterName       string `pulumi:"clusterName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
}


type LookupClusterResult struct {
	AssociatedWorkspaces          []AssociatedWorkspaceResponse          `pulumi:"associatedWorkspaces"`
	BillingType                   *string                                `pulumi:"billingType"`
	CapacityReservationProperties *CapacityReservationPropertiesResponse `pulumi:"capacityReservationProperties"`
	ClusterId                     string                                 `pulumi:"clusterId"`
	CreatedDate                   string                                 `pulumi:"createdDate"`
	Id                            string                                 `pulumi:"id"`
	Identity                      *IdentityResponse                      `pulumi:"identity"`
	IsAvailabilityZonesEnabled    *bool                                  `pulumi:"isAvailabilityZonesEnabled"`
	KeyVaultProperties            *KeyVaultPropertiesResponse            `pulumi:"keyVaultProperties"`
	LastModifiedDate              string                                 `pulumi:"lastModifiedDate"`
	Location                      string                                 `pulumi:"location"`
	Name                          string                                 `pulumi:"name"`
	ProvisioningState             string                                 `pulumi:"provisioningState"`
	Sku                           *ClusterSkuResponse                    `pulumi:"sku"`
	Tags                          map[string]string                      `pulumi:"tags"`
	Type                          string                                 `pulumi:"type"`
}

func LookupClusterOutput(ctx *pulumi.Context, args LookupClusterOutputArgs, opts ...pulumi.InvokeOption) LookupClusterResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupClusterResult, error) {
			args := v.(LookupClusterArgs)
			r, err := LookupCluster(ctx, &args, opts...)
			var s LookupClusterResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupClusterResultOutput)
}

type LookupClusterOutputArgs struct {
	ClusterName       pulumi.StringInput `pulumi:"clusterName"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupClusterOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupClusterArgs)(nil)).Elem()
}


type LookupClusterResultOutput struct{ *pulumi.OutputState }

func (LookupClusterResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupClusterResult)(nil)).Elem()
}

func (o LookupClusterResultOutput) ToLookupClusterResultOutput() LookupClusterResultOutput {
	return o
}

func (o LookupClusterResultOutput) ToLookupClusterResultOutputWithContext(ctx context.Context) LookupClusterResultOutput {
	return o
}

func (o LookupClusterResultOutput) AssociatedWorkspaces() AssociatedWorkspaceResponseArrayOutput {
	return o.ApplyT(func(v LookupClusterResult) []AssociatedWorkspaceResponse { return v.AssociatedWorkspaces }).(AssociatedWorkspaceResponseArrayOutput)
}

func (o LookupClusterResultOutput) BillingType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupClusterResult) *string { return v.BillingType }).(pulumi.StringPtrOutput)
}

func (o LookupClusterResultOutput) CapacityReservationProperties() CapacityReservationPropertiesResponsePtrOutput {
	return o.ApplyT(func(v LookupClusterResult) *CapacityReservationPropertiesResponse {
		return v.CapacityReservationProperties
	}).(CapacityReservationPropertiesResponsePtrOutput)
}

func (o LookupClusterResultOutput) ClusterId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupClusterResult) string { return v.ClusterId }).(pulumi.StringOutput)
}

func (o LookupClusterResultOutput) CreatedDate() pulumi.StringOutput {
	return o.ApplyT(func(v LookupClusterResult) string { return v.CreatedDate }).(pulumi.StringOutput)
}

func (o LookupClusterResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupClusterResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupClusterResultOutput) Identity() IdentityResponsePtrOutput {
	return o.ApplyT(func(v LookupClusterResult) *IdentityResponse { return v.Identity }).(IdentityResponsePtrOutput)
}

func (o LookupClusterResultOutput) IsAvailabilityZonesEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupClusterResult) *bool { return v.IsAvailabilityZonesEnabled }).(pulumi.BoolPtrOutput)
}

func (o LookupClusterResultOutput) KeyVaultProperties() KeyVaultPropertiesResponsePtrOutput {
	return o.ApplyT(func(v LookupClusterResult) *KeyVaultPropertiesResponse { return v.KeyVaultProperties }).(KeyVaultPropertiesResponsePtrOutput)
}

func (o LookupClusterResultOutput) LastModifiedDate() pulumi.StringOutput {
	return o.ApplyT(func(v LookupClusterResult) string { return v.LastModifiedDate }).(pulumi.StringOutput)
}

func (o LookupClusterResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupClusterResult) string { return v.Location }).(pulumi.StringOutput)
}

func (o LookupClusterResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupClusterResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupClusterResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupClusterResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o LookupClusterResultOutput) Sku() ClusterSkuResponsePtrOutput {
	return o.ApplyT(func(v LookupClusterResult) *ClusterSkuResponse { return v.Sku }).(ClusterSkuResponsePtrOutput)
}

func (o LookupClusterResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupClusterResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o LookupClusterResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupClusterResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupClusterResultOutput{})
}
