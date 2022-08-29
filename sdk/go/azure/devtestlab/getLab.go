


package devtestlab

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupLab(ctx *pulumi.Context, args *LookupLabArgs, opts ...pulumi.InvokeOption) (*LookupLabResult, error) {
	var rv LookupLabResult
	err := ctx.Invoke("azure-native:devtestlab:getLab", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return rv.Defaults(), nil
}

type LookupLabArgs struct {
	Expand            *string `pulumi:"expand"`
	Name              string  `pulumi:"name"`
	ResourceGroupName string  `pulumi:"resourceGroupName"`
}


type LookupLabResult struct {
	Announcement                         *LabAnnouncementPropertiesResponse `pulumi:"announcement"`
	ArtifactsStorageAccount              string                             `pulumi:"artifactsStorageAccount"`
	CreatedDate                          string                             `pulumi:"createdDate"`
	DefaultPremiumStorageAccount         string                             `pulumi:"defaultPremiumStorageAccount"`
	DefaultStorageAccount                string                             `pulumi:"defaultStorageAccount"`
	EnvironmentPermission                *string                            `pulumi:"environmentPermission"`
	ExtendedProperties                   map[string]string                  `pulumi:"extendedProperties"`
	Id                                   string                             `pulumi:"id"`
	LabStorageType                       *string                            `pulumi:"labStorageType"`
	LoadBalancerId                       string                             `pulumi:"loadBalancerId"`
	Location                             *string                            `pulumi:"location"`
	MandatoryArtifactsResourceIdsLinux   []string                           `pulumi:"mandatoryArtifactsResourceIdsLinux"`
	MandatoryArtifactsResourceIdsWindows []string                           `pulumi:"mandatoryArtifactsResourceIdsWindows"`
	Name                                 string                             `pulumi:"name"`
	NetworkSecurityGroupId               string                             `pulumi:"networkSecurityGroupId"`
	PremiumDataDiskStorageAccount        string                             `pulumi:"premiumDataDiskStorageAccount"`
	PremiumDataDisks                     *string                            `pulumi:"premiumDataDisks"`
	ProvisioningState                    string                             `pulumi:"provisioningState"`
	PublicIpId                           string                             `pulumi:"publicIpId"`
	Support                              *LabSupportPropertiesResponse      `pulumi:"support"`
	Tags                                 map[string]string                  `pulumi:"tags"`
	Type                                 string                             `pulumi:"type"`
	UniqueIdentifier                     string                             `pulumi:"uniqueIdentifier"`
	VaultName                            string                             `pulumi:"vaultName"`
	VmCreationResourceGroup              string                             `pulumi:"vmCreationResourceGroup"`
}


func (val *LookupLabResult) Defaults() *LookupLabResult {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.LabStorageType) {
		labStorageType_ := "Premium"
		tmp.LabStorageType = &labStorageType_
	}
	return &tmp
}

func LookupLabOutput(ctx *pulumi.Context, args LookupLabOutputArgs, opts ...pulumi.InvokeOption) LookupLabResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupLabResult, error) {
			args := v.(LookupLabArgs)
			r, err := LookupLab(ctx, &args, opts...)
			var s LookupLabResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupLabResultOutput)
}

type LookupLabOutputArgs struct {
	Expand            pulumi.StringPtrInput `pulumi:"expand"`
	Name              pulumi.StringInput    `pulumi:"name"`
	ResourceGroupName pulumi.StringInput    `pulumi:"resourceGroupName"`
}

func (LookupLabOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupLabArgs)(nil)).Elem()
}


type LookupLabResultOutput struct{ *pulumi.OutputState }

func (LookupLabResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupLabResult)(nil)).Elem()
}

func (o LookupLabResultOutput) ToLookupLabResultOutput() LookupLabResultOutput {
	return o
}

func (o LookupLabResultOutput) ToLookupLabResultOutputWithContext(ctx context.Context) LookupLabResultOutput {
	return o
}

func (o LookupLabResultOutput) Announcement() LabAnnouncementPropertiesResponsePtrOutput {
	return o.ApplyT(func(v LookupLabResult) *LabAnnouncementPropertiesResponse { return v.Announcement }).(LabAnnouncementPropertiesResponsePtrOutput)
}

func (o LookupLabResultOutput) ArtifactsStorageAccount() pulumi.StringOutput {
	return o.ApplyT(func(v LookupLabResult) string { return v.ArtifactsStorageAccount }).(pulumi.StringOutput)
}

func (o LookupLabResultOutput) CreatedDate() pulumi.StringOutput {
	return o.ApplyT(func(v LookupLabResult) string { return v.CreatedDate }).(pulumi.StringOutput)
}

func (o LookupLabResultOutput) DefaultPremiumStorageAccount() pulumi.StringOutput {
	return o.ApplyT(func(v LookupLabResult) string { return v.DefaultPremiumStorageAccount }).(pulumi.StringOutput)
}

func (o LookupLabResultOutput) DefaultStorageAccount() pulumi.StringOutput {
	return o.ApplyT(func(v LookupLabResult) string { return v.DefaultStorageAccount }).(pulumi.StringOutput)
}

func (o LookupLabResultOutput) EnvironmentPermission() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupLabResult) *string { return v.EnvironmentPermission }).(pulumi.StringPtrOutput)
}

func (o LookupLabResultOutput) ExtendedProperties() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupLabResult) map[string]string { return v.ExtendedProperties }).(pulumi.StringMapOutput)
}

func (o LookupLabResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupLabResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupLabResultOutput) LabStorageType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupLabResult) *string { return v.LabStorageType }).(pulumi.StringPtrOutput)
}

func (o LookupLabResultOutput) LoadBalancerId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupLabResult) string { return v.LoadBalancerId }).(pulumi.StringOutput)
}

func (o LookupLabResultOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupLabResult) *string { return v.Location }).(pulumi.StringPtrOutput)
}

func (o LookupLabResultOutput) MandatoryArtifactsResourceIdsLinux() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupLabResult) []string { return v.MandatoryArtifactsResourceIdsLinux }).(pulumi.StringArrayOutput)
}

func (o LookupLabResultOutput) MandatoryArtifactsResourceIdsWindows() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupLabResult) []string { return v.MandatoryArtifactsResourceIdsWindows }).(pulumi.StringArrayOutput)
}

func (o LookupLabResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupLabResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupLabResultOutput) NetworkSecurityGroupId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupLabResult) string { return v.NetworkSecurityGroupId }).(pulumi.StringOutput)
}

func (o LookupLabResultOutput) PremiumDataDiskStorageAccount() pulumi.StringOutput {
	return o.ApplyT(func(v LookupLabResult) string { return v.PremiumDataDiskStorageAccount }).(pulumi.StringOutput)
}

func (o LookupLabResultOutput) PremiumDataDisks() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupLabResult) *string { return v.PremiumDataDisks }).(pulumi.StringPtrOutput)
}

func (o LookupLabResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupLabResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o LookupLabResultOutput) PublicIpId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupLabResult) string { return v.PublicIpId }).(pulumi.StringOutput)
}

func (o LookupLabResultOutput) Support() LabSupportPropertiesResponsePtrOutput {
	return o.ApplyT(func(v LookupLabResult) *LabSupportPropertiesResponse { return v.Support }).(LabSupportPropertiesResponsePtrOutput)
}

func (o LookupLabResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupLabResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o LookupLabResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupLabResult) string { return v.Type }).(pulumi.StringOutput)
}

func (o LookupLabResultOutput) UniqueIdentifier() pulumi.StringOutput {
	return o.ApplyT(func(v LookupLabResult) string { return v.UniqueIdentifier }).(pulumi.StringOutput)
}

func (o LookupLabResultOutput) VaultName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupLabResult) string { return v.VaultName }).(pulumi.StringOutput)
}

func (o LookupLabResultOutput) VmCreationResourceGroup() pulumi.StringOutput {
	return o.ApplyT(func(v LookupLabResult) string { return v.VmCreationResourceGroup }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupLabResultOutput{})
}
