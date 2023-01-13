


package v20220101

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupCluster(ctx *pulumi.Context, args *LookupClusterArgs, opts ...pulumi.InvokeOption) (*LookupClusterResult, error) {
	var rv LookupClusterResult
	err := ctx.Invoke("azure-native:azurestackhci/v20220101:getCluster", args, &rv, opts...)
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
	AadClientId             string                            `pulumi:"aadClientId"`
	AadTenantId             string                            `pulumi:"aadTenantId"`
	BillingModel            string                            `pulumi:"billingModel"`
	CloudId                 string                            `pulumi:"cloudId"`
	CloudManagementEndpoint *string                           `pulumi:"cloudManagementEndpoint"`
	CreatedAt               *string                           `pulumi:"createdAt"`
	CreatedBy               *string                           `pulumi:"createdBy"`
	CreatedByType           *string                           `pulumi:"createdByType"`
	DesiredProperties       *ClusterDesiredPropertiesResponse `pulumi:"desiredProperties"`
	Id                      string                            `pulumi:"id"`
	LastBillingTimestamp    string                            `pulumi:"lastBillingTimestamp"`
	LastModifiedAt          *string                           `pulumi:"lastModifiedAt"`
	LastModifiedBy          *string                           `pulumi:"lastModifiedBy"`
	LastModifiedByType      *string                           `pulumi:"lastModifiedByType"`
	LastSyncTimestamp       string                            `pulumi:"lastSyncTimestamp"`
	Location                string                            `pulumi:"location"`
	Name                    string                            `pulumi:"name"`
	ProvisioningState       string                            `pulumi:"provisioningState"`
	RegistrationTimestamp   string                            `pulumi:"registrationTimestamp"`
	ReportedProperties      ClusterReportedPropertiesResponse `pulumi:"reportedProperties"`
	Status                  string                            `pulumi:"status"`
	Tags                    map[string]string                 `pulumi:"tags"`
	TrialDaysRemaining      float64                           `pulumi:"trialDaysRemaining"`
	Type                    string                            `pulumi:"type"`
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

func (o LookupClusterResultOutput) AadClientId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupClusterResult) string { return v.AadClientId }).(pulumi.StringOutput)
}

func (o LookupClusterResultOutput) AadTenantId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupClusterResult) string { return v.AadTenantId }).(pulumi.StringOutput)
}

func (o LookupClusterResultOutput) BillingModel() pulumi.StringOutput {
	return o.ApplyT(func(v LookupClusterResult) string { return v.BillingModel }).(pulumi.StringOutput)
}

func (o LookupClusterResultOutput) CloudId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupClusterResult) string { return v.CloudId }).(pulumi.StringOutput)
}

func (o LookupClusterResultOutput) CloudManagementEndpoint() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupClusterResult) *string { return v.CloudManagementEndpoint }).(pulumi.StringPtrOutput)
}

func (o LookupClusterResultOutput) CreatedAt() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupClusterResult) *string { return v.CreatedAt }).(pulumi.StringPtrOutput)
}

func (o LookupClusterResultOutput) CreatedBy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupClusterResult) *string { return v.CreatedBy }).(pulumi.StringPtrOutput)
}

func (o LookupClusterResultOutput) CreatedByType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupClusterResult) *string { return v.CreatedByType }).(pulumi.StringPtrOutput)
}

func (o LookupClusterResultOutput) DesiredProperties() ClusterDesiredPropertiesResponsePtrOutput {
	return o.ApplyT(func(v LookupClusterResult) *ClusterDesiredPropertiesResponse { return v.DesiredProperties }).(ClusterDesiredPropertiesResponsePtrOutput)
}

func (o LookupClusterResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupClusterResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupClusterResultOutput) LastBillingTimestamp() pulumi.StringOutput {
	return o.ApplyT(func(v LookupClusterResult) string { return v.LastBillingTimestamp }).(pulumi.StringOutput)
}

func (o LookupClusterResultOutput) LastModifiedAt() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupClusterResult) *string { return v.LastModifiedAt }).(pulumi.StringPtrOutput)
}

func (o LookupClusterResultOutput) LastModifiedBy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupClusterResult) *string { return v.LastModifiedBy }).(pulumi.StringPtrOutput)
}

func (o LookupClusterResultOutput) LastModifiedByType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupClusterResult) *string { return v.LastModifiedByType }).(pulumi.StringPtrOutput)
}

func (o LookupClusterResultOutput) LastSyncTimestamp() pulumi.StringOutput {
	return o.ApplyT(func(v LookupClusterResult) string { return v.LastSyncTimestamp }).(pulumi.StringOutput)
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

func (o LookupClusterResultOutput) RegistrationTimestamp() pulumi.StringOutput {
	return o.ApplyT(func(v LookupClusterResult) string { return v.RegistrationTimestamp }).(pulumi.StringOutput)
}

func (o LookupClusterResultOutput) ReportedProperties() ClusterReportedPropertiesResponseOutput {
	return o.ApplyT(func(v LookupClusterResult) ClusterReportedPropertiesResponse { return v.ReportedProperties }).(ClusterReportedPropertiesResponseOutput)
}

func (o LookupClusterResultOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v LookupClusterResult) string { return v.Status }).(pulumi.StringOutput)
}

func (o LookupClusterResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupClusterResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o LookupClusterResultOutput) TrialDaysRemaining() pulumi.Float64Output {
	return o.ApplyT(func(v LookupClusterResult) float64 { return v.TrialDaysRemaining }).(pulumi.Float64Output)
}

func (o LookupClusterResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupClusterResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupClusterResultOutput{})
}
