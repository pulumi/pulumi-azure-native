


package v20150501preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupManagedInstance(ctx *pulumi.Context, args *LookupManagedInstanceArgs, opts ...pulumi.InvokeOption) (*LookupManagedInstanceResult, error) {
	var rv LookupManagedInstanceResult
	err := ctx.Invoke("azure-native:sql/v20150501preview:getManagedInstance", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupManagedInstanceArgs struct {
	ManagedInstanceName string `pulumi:"managedInstanceName"`
	ResourceGroupName   string `pulumi:"resourceGroupName"`
}


type LookupManagedInstanceResult struct {
	AdministratorLogin         *string                   `pulumi:"administratorLogin"`
	Collation                  *string                   `pulumi:"collation"`
	DnsZone                    string                    `pulumi:"dnsZone"`
	FullyQualifiedDomainName   string                    `pulumi:"fullyQualifiedDomainName"`
	Id                         string                    `pulumi:"id"`
	Identity                   *ResourceIdentityResponse `pulumi:"identity"`
	InstancePoolId             *string                   `pulumi:"instancePoolId"`
	LicenseType                *string                   `pulumi:"licenseType"`
	Location                   string                    `pulumi:"location"`
	MaintenanceConfigurationId *string                   `pulumi:"maintenanceConfigurationId"`
	MinimalTlsVersion          *string                   `pulumi:"minimalTlsVersion"`
	Name                       string                    `pulumi:"name"`
	ProxyOverride              *string                   `pulumi:"proxyOverride"`
	PublicDataEndpointEnabled  *bool                     `pulumi:"publicDataEndpointEnabled"`
	Sku                        *SkuResponse              `pulumi:"sku"`
	State                      string                    `pulumi:"state"`
	StorageSizeInGB            *int                      `pulumi:"storageSizeInGB"`
	SubnetId                   *string                   `pulumi:"subnetId"`
	Tags                       map[string]string         `pulumi:"tags"`
	TimezoneId                 *string                   `pulumi:"timezoneId"`
	Type                       string                    `pulumi:"type"`
	VCores                     *int                      `pulumi:"vCores"`
}

func LookupManagedInstanceOutput(ctx *pulumi.Context, args LookupManagedInstanceOutputArgs, opts ...pulumi.InvokeOption) LookupManagedInstanceResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupManagedInstanceResult, error) {
			args := v.(LookupManagedInstanceArgs)
			r, err := LookupManagedInstance(ctx, &args, opts...)
			var s LookupManagedInstanceResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupManagedInstanceResultOutput)
}

type LookupManagedInstanceOutputArgs struct {
	ManagedInstanceName pulumi.StringInput `pulumi:"managedInstanceName"`
	ResourceGroupName   pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupManagedInstanceOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupManagedInstanceArgs)(nil)).Elem()
}


type LookupManagedInstanceResultOutput struct{ *pulumi.OutputState }

func (LookupManagedInstanceResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupManagedInstanceResult)(nil)).Elem()
}

func (o LookupManagedInstanceResultOutput) ToLookupManagedInstanceResultOutput() LookupManagedInstanceResultOutput {
	return o
}

func (o LookupManagedInstanceResultOutput) ToLookupManagedInstanceResultOutputWithContext(ctx context.Context) LookupManagedInstanceResultOutput {
	return o
}

func (o LookupManagedInstanceResultOutput) AdministratorLogin() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupManagedInstanceResult) *string { return v.AdministratorLogin }).(pulumi.StringPtrOutput)
}

func (o LookupManagedInstanceResultOutput) Collation() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupManagedInstanceResult) *string { return v.Collation }).(pulumi.StringPtrOutput)
}

func (o LookupManagedInstanceResultOutput) DnsZone() pulumi.StringOutput {
	return o.ApplyT(func(v LookupManagedInstanceResult) string { return v.DnsZone }).(pulumi.StringOutput)
}

func (o LookupManagedInstanceResultOutput) FullyQualifiedDomainName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupManagedInstanceResult) string { return v.FullyQualifiedDomainName }).(pulumi.StringOutput)
}

func (o LookupManagedInstanceResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupManagedInstanceResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupManagedInstanceResultOutput) Identity() ResourceIdentityResponsePtrOutput {
	return o.ApplyT(func(v LookupManagedInstanceResult) *ResourceIdentityResponse { return v.Identity }).(ResourceIdentityResponsePtrOutput)
}

func (o LookupManagedInstanceResultOutput) InstancePoolId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupManagedInstanceResult) *string { return v.InstancePoolId }).(pulumi.StringPtrOutput)
}

func (o LookupManagedInstanceResultOutput) LicenseType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupManagedInstanceResult) *string { return v.LicenseType }).(pulumi.StringPtrOutput)
}

func (o LookupManagedInstanceResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupManagedInstanceResult) string { return v.Location }).(pulumi.StringOutput)
}

func (o LookupManagedInstanceResultOutput) MaintenanceConfigurationId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupManagedInstanceResult) *string { return v.MaintenanceConfigurationId }).(pulumi.StringPtrOutput)
}

func (o LookupManagedInstanceResultOutput) MinimalTlsVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupManagedInstanceResult) *string { return v.MinimalTlsVersion }).(pulumi.StringPtrOutput)
}

func (o LookupManagedInstanceResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupManagedInstanceResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupManagedInstanceResultOutput) ProxyOverride() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupManagedInstanceResult) *string { return v.ProxyOverride }).(pulumi.StringPtrOutput)
}

func (o LookupManagedInstanceResultOutput) PublicDataEndpointEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupManagedInstanceResult) *bool { return v.PublicDataEndpointEnabled }).(pulumi.BoolPtrOutput)
}

func (o LookupManagedInstanceResultOutput) Sku() SkuResponsePtrOutput {
	return o.ApplyT(func(v LookupManagedInstanceResult) *SkuResponse { return v.Sku }).(SkuResponsePtrOutput)
}

func (o LookupManagedInstanceResultOutput) State() pulumi.StringOutput {
	return o.ApplyT(func(v LookupManagedInstanceResult) string { return v.State }).(pulumi.StringOutput)
}

func (o LookupManagedInstanceResultOutput) StorageSizeInGB() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupManagedInstanceResult) *int { return v.StorageSizeInGB }).(pulumi.IntPtrOutput)
}

func (o LookupManagedInstanceResultOutput) SubnetId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupManagedInstanceResult) *string { return v.SubnetId }).(pulumi.StringPtrOutput)
}

func (o LookupManagedInstanceResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupManagedInstanceResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o LookupManagedInstanceResultOutput) TimezoneId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupManagedInstanceResult) *string { return v.TimezoneId }).(pulumi.StringPtrOutput)
}

func (o LookupManagedInstanceResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupManagedInstanceResult) string { return v.Type }).(pulumi.StringOutput)
}

func (o LookupManagedInstanceResultOutput) VCores() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupManagedInstanceResult) *int { return v.VCores }).(pulumi.IntPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupManagedInstanceResultOutput{})
}
