


package workloads

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupPhpWorkload(ctx *pulumi.Context, args *LookupPhpWorkloadArgs, opts ...pulumi.InvokeOption) (*LookupPhpWorkloadResult, error) {
	var rv LookupPhpWorkloadResult
	err := ctx.Invoke("azure-native:workloads:getPhpWorkload", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupPhpWorkloadArgs struct {
	PhpWorkloadName   string `pulumi:"phpWorkloadName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
}


type LookupPhpWorkloadResult struct {
	AdminUserProfile                  UserProfileResponse                  `pulumi:"adminUserProfile"`
	AppLocation                       string                               `pulumi:"appLocation"`
	BackupProfile                     *BackupProfileResponse               `pulumi:"backupProfile"`
	CacheProfile                      *CacheProfileResponse                `pulumi:"cacheProfile"`
	ControllerProfile                 NodeProfileResponse                  `pulumi:"controllerProfile"`
	DatabaseProfile                   DatabaseProfileResponse              `pulumi:"databaseProfile"`
	FileshareProfile                  *FileshareProfileResponse            `pulumi:"fileshareProfile"`
	Id                                string                               `pulumi:"id"`
	Identity                          *PhpWorkloadResourceResponseIdentity `pulumi:"identity"`
	Kind                              string                               `pulumi:"kind"`
	Location                          string                               `pulumi:"location"`
	ManagedResourceGroupConfiguration *ManagedRGConfigurationResponse      `pulumi:"managedResourceGroupConfiguration"`
	Name                              string                               `pulumi:"name"`
	NetworkProfile                    *NetworkProfileResponse              `pulumi:"networkProfile"`
	PhpProfile                        *PhpProfileResponse                  `pulumi:"phpProfile"`
	ProvisioningState                 string                               `pulumi:"provisioningState"`
	SearchProfile                     *SearchProfileResponse               `pulumi:"searchProfile"`
	SiteProfile                       *SiteProfileResponse                 `pulumi:"siteProfile"`
	Sku                               *SkuResponse                         `pulumi:"sku"`
	SystemData                        SystemDataResponse                   `pulumi:"systemData"`
	Tags                              map[string]string                    `pulumi:"tags"`
	Type                              string                               `pulumi:"type"`
	WebNodesProfile                   VmssNodesProfileResponse             `pulumi:"webNodesProfile"`
}

func LookupPhpWorkloadOutput(ctx *pulumi.Context, args LookupPhpWorkloadOutputArgs, opts ...pulumi.InvokeOption) LookupPhpWorkloadResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupPhpWorkloadResult, error) {
			args := v.(LookupPhpWorkloadArgs)
			r, err := LookupPhpWorkload(ctx, &args, opts...)
			var s LookupPhpWorkloadResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupPhpWorkloadResultOutput)
}

type LookupPhpWorkloadOutputArgs struct {
	PhpWorkloadName   pulumi.StringInput `pulumi:"phpWorkloadName"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupPhpWorkloadOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupPhpWorkloadArgs)(nil)).Elem()
}


type LookupPhpWorkloadResultOutput struct{ *pulumi.OutputState }

func (LookupPhpWorkloadResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupPhpWorkloadResult)(nil)).Elem()
}

func (o LookupPhpWorkloadResultOutput) ToLookupPhpWorkloadResultOutput() LookupPhpWorkloadResultOutput {
	return o
}

func (o LookupPhpWorkloadResultOutput) ToLookupPhpWorkloadResultOutputWithContext(ctx context.Context) LookupPhpWorkloadResultOutput {
	return o
}

func (o LookupPhpWorkloadResultOutput) AdminUserProfile() UserProfileResponseOutput {
	return o.ApplyT(func(v LookupPhpWorkloadResult) UserProfileResponse { return v.AdminUserProfile }).(UserProfileResponseOutput)
}

func (o LookupPhpWorkloadResultOutput) AppLocation() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPhpWorkloadResult) string { return v.AppLocation }).(pulumi.StringOutput)
}

func (o LookupPhpWorkloadResultOutput) BackupProfile() BackupProfileResponsePtrOutput {
	return o.ApplyT(func(v LookupPhpWorkloadResult) *BackupProfileResponse { return v.BackupProfile }).(BackupProfileResponsePtrOutput)
}

func (o LookupPhpWorkloadResultOutput) CacheProfile() CacheProfileResponsePtrOutput {
	return o.ApplyT(func(v LookupPhpWorkloadResult) *CacheProfileResponse { return v.CacheProfile }).(CacheProfileResponsePtrOutput)
}

func (o LookupPhpWorkloadResultOutput) ControllerProfile() NodeProfileResponseOutput {
	return o.ApplyT(func(v LookupPhpWorkloadResult) NodeProfileResponse { return v.ControllerProfile }).(NodeProfileResponseOutput)
}

func (o LookupPhpWorkloadResultOutput) DatabaseProfile() DatabaseProfileResponseOutput {
	return o.ApplyT(func(v LookupPhpWorkloadResult) DatabaseProfileResponse { return v.DatabaseProfile }).(DatabaseProfileResponseOutput)
}

func (o LookupPhpWorkloadResultOutput) FileshareProfile() FileshareProfileResponsePtrOutput {
	return o.ApplyT(func(v LookupPhpWorkloadResult) *FileshareProfileResponse { return v.FileshareProfile }).(FileshareProfileResponsePtrOutput)
}

func (o LookupPhpWorkloadResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPhpWorkloadResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupPhpWorkloadResultOutput) Identity() PhpWorkloadResourceResponseIdentityPtrOutput {
	return o.ApplyT(func(v LookupPhpWorkloadResult) *PhpWorkloadResourceResponseIdentity { return v.Identity }).(PhpWorkloadResourceResponseIdentityPtrOutput)
}

func (o LookupPhpWorkloadResultOutput) Kind() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPhpWorkloadResult) string { return v.Kind }).(pulumi.StringOutput)
}

func (o LookupPhpWorkloadResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPhpWorkloadResult) string { return v.Location }).(pulumi.StringOutput)
}

func (o LookupPhpWorkloadResultOutput) ManagedResourceGroupConfiguration() ManagedRGConfigurationResponsePtrOutput {
	return o.ApplyT(func(v LookupPhpWorkloadResult) *ManagedRGConfigurationResponse {
		return v.ManagedResourceGroupConfiguration
	}).(ManagedRGConfigurationResponsePtrOutput)
}

func (o LookupPhpWorkloadResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPhpWorkloadResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupPhpWorkloadResultOutput) NetworkProfile() NetworkProfileResponsePtrOutput {
	return o.ApplyT(func(v LookupPhpWorkloadResult) *NetworkProfileResponse { return v.NetworkProfile }).(NetworkProfileResponsePtrOutput)
}

func (o LookupPhpWorkloadResultOutput) PhpProfile() PhpProfileResponsePtrOutput {
	return o.ApplyT(func(v LookupPhpWorkloadResult) *PhpProfileResponse { return v.PhpProfile }).(PhpProfileResponsePtrOutput)
}

func (o LookupPhpWorkloadResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPhpWorkloadResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o LookupPhpWorkloadResultOutput) SearchProfile() SearchProfileResponsePtrOutput {
	return o.ApplyT(func(v LookupPhpWorkloadResult) *SearchProfileResponse { return v.SearchProfile }).(SearchProfileResponsePtrOutput)
}

func (o LookupPhpWorkloadResultOutput) SiteProfile() SiteProfileResponsePtrOutput {
	return o.ApplyT(func(v LookupPhpWorkloadResult) *SiteProfileResponse { return v.SiteProfile }).(SiteProfileResponsePtrOutput)
}

func (o LookupPhpWorkloadResultOutput) Sku() SkuResponsePtrOutput {
	return o.ApplyT(func(v LookupPhpWorkloadResult) *SkuResponse { return v.Sku }).(SkuResponsePtrOutput)
}

func (o LookupPhpWorkloadResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupPhpWorkloadResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

func (o LookupPhpWorkloadResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupPhpWorkloadResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o LookupPhpWorkloadResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPhpWorkloadResult) string { return v.Type }).(pulumi.StringOutput)
}

func (o LookupPhpWorkloadResultOutput) WebNodesProfile() VmssNodesProfileResponseOutput {
	return o.ApplyT(func(v LookupPhpWorkloadResult) VmssNodesProfileResponse { return v.WebNodesProfile }).(VmssNodesProfileResponseOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupPhpWorkloadResultOutput{})
}
