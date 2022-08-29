


package v20211201preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type PhpWorkload struct {
	pulumi.CustomResourceState

	AdminUserProfile                  UserProfileResponseOutput                    `pulumi:"adminUserProfile"`
	AppLocation                       pulumi.StringOutput                          `pulumi:"appLocation"`
	BackupProfile                     BackupProfileResponsePtrOutput               `pulumi:"backupProfile"`
	CacheProfile                      CacheProfileResponsePtrOutput                `pulumi:"cacheProfile"`
	ControllerProfile                 NodeProfileResponseOutput                    `pulumi:"controllerProfile"`
	DatabaseProfile                   DatabaseProfileResponseOutput                `pulumi:"databaseProfile"`
	FileshareProfile                  FileshareProfileResponsePtrOutput            `pulumi:"fileshareProfile"`
	Identity                          PhpWorkloadResourceResponseIdentityPtrOutput `pulumi:"identity"`
	Kind                              pulumi.StringOutput                          `pulumi:"kind"`
	Location                          pulumi.StringOutput                          `pulumi:"location"`
	ManagedResourceGroupConfiguration ManagedRGConfigurationResponsePtrOutput      `pulumi:"managedResourceGroupConfiguration"`
	Name                              pulumi.StringOutput                          `pulumi:"name"`
	NetworkProfile                    NetworkProfileResponsePtrOutput              `pulumi:"networkProfile"`
	PhpProfile                        PhpProfileResponsePtrOutput                  `pulumi:"phpProfile"`
	ProvisioningState                 pulumi.StringOutput                          `pulumi:"provisioningState"`
	SearchProfile                     SearchProfileResponsePtrOutput               `pulumi:"searchProfile"`
	SiteProfile                       SiteProfileResponsePtrOutput                 `pulumi:"siteProfile"`
	Sku                               SkuResponsePtrOutput                         `pulumi:"sku"`
	SystemData                        SystemDataResponseOutput                     `pulumi:"systemData"`
	Tags                              pulumi.StringMapOutput                       `pulumi:"tags"`
	Type                              pulumi.StringOutput                          `pulumi:"type"`
	WebNodesProfile                   VmssNodesProfileResponseOutput               `pulumi:"webNodesProfile"`
}


func NewPhpWorkload(ctx *pulumi.Context,
	name string, args *PhpWorkloadArgs, opts ...pulumi.ResourceOption) (*PhpWorkload, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AdminUserProfile == nil {
		return nil, errors.New("invalid value for required argument 'AdminUserProfile'")
	}
	if args.AppLocation == nil {
		return nil, errors.New("invalid value for required argument 'AppLocation'")
	}
	if args.ControllerProfile == nil {
		return nil, errors.New("invalid value for required argument 'ControllerProfile'")
	}
	if args.DatabaseProfile == nil {
		return nil, errors.New("invalid value for required argument 'DatabaseProfile'")
	}
	if args.Kind == nil {
		return nil, errors.New("invalid value for required argument 'Kind'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.WebNodesProfile == nil {
		return nil, errors.New("invalid value for required argument 'WebNodesProfile'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:workloads:PhpWorkload"),
		},
	})
	opts = append(opts, aliases)
	var resource PhpWorkload
	err := ctx.RegisterResource("azure-native:workloads/v20211201preview:PhpWorkload", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetPhpWorkload(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *PhpWorkloadState, opts ...pulumi.ResourceOption) (*PhpWorkload, error) {
	var resource PhpWorkload
	err := ctx.ReadResource("azure-native:workloads/v20211201preview:PhpWorkload", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type phpWorkloadState struct {
}

type PhpWorkloadState struct {
}

func (PhpWorkloadState) ElementType() reflect.Type {
	return reflect.TypeOf((*phpWorkloadState)(nil)).Elem()
}

type phpWorkloadArgs struct {
	AdminUserProfile                  UserProfile                  `pulumi:"adminUserProfile"`
	AppLocation                       string                       `pulumi:"appLocation"`
	BackupProfile                     *BackupProfile               `pulumi:"backupProfile"`
	CacheProfile                      *CacheProfile                `pulumi:"cacheProfile"`
	ControllerProfile                 NodeProfile                  `pulumi:"controllerProfile"`
	DatabaseProfile                   DatabaseProfile              `pulumi:"databaseProfile"`
	FileshareProfile                  *FileshareProfile            `pulumi:"fileshareProfile"`
	Identity                          *PhpWorkloadResourceIdentity `pulumi:"identity"`
	Kind                              string                       `pulumi:"kind"`
	Location                          *string                      `pulumi:"location"`
	ManagedResourceGroupConfiguration *ManagedRGConfiguration      `pulumi:"managedResourceGroupConfiguration"`
	NetworkProfile                    *NetworkProfile              `pulumi:"networkProfile"`
	PhpProfile                        *PhpProfile                  `pulumi:"phpProfile"`
	PhpWorkloadName                   *string                      `pulumi:"phpWorkloadName"`
	ResourceGroupName                 string                       `pulumi:"resourceGroupName"`
	SearchProfile                     *SearchProfile               `pulumi:"searchProfile"`
	SiteProfile                       *SiteProfile                 `pulumi:"siteProfile"`
	Sku                               *Sku                         `pulumi:"sku"`
	Tags                              map[string]string            `pulumi:"tags"`
	WebNodesProfile                   VmssNodesProfile             `pulumi:"webNodesProfile"`
}


type PhpWorkloadArgs struct {
	AdminUserProfile                  UserProfileInput
	AppLocation                       pulumi.StringInput
	BackupProfile                     BackupProfilePtrInput
	CacheProfile                      CacheProfilePtrInput
	ControllerProfile                 NodeProfileInput
	DatabaseProfile                   DatabaseProfileInput
	FileshareProfile                  FileshareProfilePtrInput
	Identity                          PhpWorkloadResourceIdentityPtrInput
	Kind                              pulumi.StringInput
	Location                          pulumi.StringPtrInput
	ManagedResourceGroupConfiguration ManagedRGConfigurationPtrInput
	NetworkProfile                    NetworkProfilePtrInput
	PhpProfile                        PhpProfilePtrInput
	PhpWorkloadName                   pulumi.StringPtrInput
	ResourceGroupName                 pulumi.StringInput
	SearchProfile                     SearchProfilePtrInput
	SiteProfile                       SiteProfilePtrInput
	Sku                               SkuPtrInput
	Tags                              pulumi.StringMapInput
	WebNodesProfile                   VmssNodesProfileInput
}

func (PhpWorkloadArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*phpWorkloadArgs)(nil)).Elem()
}

type PhpWorkloadInput interface {
	pulumi.Input

	ToPhpWorkloadOutput() PhpWorkloadOutput
	ToPhpWorkloadOutputWithContext(ctx context.Context) PhpWorkloadOutput
}

func (*PhpWorkload) ElementType() reflect.Type {
	return reflect.TypeOf((**PhpWorkload)(nil)).Elem()
}

func (i *PhpWorkload) ToPhpWorkloadOutput() PhpWorkloadOutput {
	return i.ToPhpWorkloadOutputWithContext(context.Background())
}

func (i *PhpWorkload) ToPhpWorkloadOutputWithContext(ctx context.Context) PhpWorkloadOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PhpWorkloadOutput)
}

type PhpWorkloadOutput struct{ *pulumi.OutputState }

func (PhpWorkloadOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PhpWorkload)(nil)).Elem()
}

func (o PhpWorkloadOutput) ToPhpWorkloadOutput() PhpWorkloadOutput {
	return o
}

func (o PhpWorkloadOutput) ToPhpWorkloadOutputWithContext(ctx context.Context) PhpWorkloadOutput {
	return o
}

func (o PhpWorkloadOutput) AdminUserProfile() UserProfileResponseOutput {
	return o.ApplyT(func(v *PhpWorkload) UserProfileResponseOutput { return v.AdminUserProfile }).(UserProfileResponseOutput)
}

func (o PhpWorkloadOutput) AppLocation() pulumi.StringOutput {
	return o.ApplyT(func(v *PhpWorkload) pulumi.StringOutput { return v.AppLocation }).(pulumi.StringOutput)
}

func (o PhpWorkloadOutput) BackupProfile() BackupProfileResponsePtrOutput {
	return o.ApplyT(func(v *PhpWorkload) BackupProfileResponsePtrOutput { return v.BackupProfile }).(BackupProfileResponsePtrOutput)
}

func (o PhpWorkloadOutput) CacheProfile() CacheProfileResponsePtrOutput {
	return o.ApplyT(func(v *PhpWorkload) CacheProfileResponsePtrOutput { return v.CacheProfile }).(CacheProfileResponsePtrOutput)
}

func (o PhpWorkloadOutput) ControllerProfile() NodeProfileResponseOutput {
	return o.ApplyT(func(v *PhpWorkload) NodeProfileResponseOutput { return v.ControllerProfile }).(NodeProfileResponseOutput)
}

func (o PhpWorkloadOutput) DatabaseProfile() DatabaseProfileResponseOutput {
	return o.ApplyT(func(v *PhpWorkload) DatabaseProfileResponseOutput { return v.DatabaseProfile }).(DatabaseProfileResponseOutput)
}

func (o PhpWorkloadOutput) FileshareProfile() FileshareProfileResponsePtrOutput {
	return o.ApplyT(func(v *PhpWorkload) FileshareProfileResponsePtrOutput { return v.FileshareProfile }).(FileshareProfileResponsePtrOutput)
}

func (o PhpWorkloadOutput) Identity() PhpWorkloadResourceResponseIdentityPtrOutput {
	return o.ApplyT(func(v *PhpWorkload) PhpWorkloadResourceResponseIdentityPtrOutput { return v.Identity }).(PhpWorkloadResourceResponseIdentityPtrOutput)
}

func (o PhpWorkloadOutput) Kind() pulumi.StringOutput {
	return o.ApplyT(func(v *PhpWorkload) pulumi.StringOutput { return v.Kind }).(pulumi.StringOutput)
}

func (o PhpWorkloadOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *PhpWorkload) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

func (o PhpWorkloadOutput) ManagedResourceGroupConfiguration() ManagedRGConfigurationResponsePtrOutput {
	return o.ApplyT(func(v *PhpWorkload) ManagedRGConfigurationResponsePtrOutput {
		return v.ManagedResourceGroupConfiguration
	}).(ManagedRGConfigurationResponsePtrOutput)
}

func (o PhpWorkloadOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *PhpWorkload) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o PhpWorkloadOutput) NetworkProfile() NetworkProfileResponsePtrOutput {
	return o.ApplyT(func(v *PhpWorkload) NetworkProfileResponsePtrOutput { return v.NetworkProfile }).(NetworkProfileResponsePtrOutput)
}

func (o PhpWorkloadOutput) PhpProfile() PhpProfileResponsePtrOutput {
	return o.ApplyT(func(v *PhpWorkload) PhpProfileResponsePtrOutput { return v.PhpProfile }).(PhpProfileResponsePtrOutput)
}

func (o PhpWorkloadOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *PhpWorkload) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o PhpWorkloadOutput) SearchProfile() SearchProfileResponsePtrOutput {
	return o.ApplyT(func(v *PhpWorkload) SearchProfileResponsePtrOutput { return v.SearchProfile }).(SearchProfileResponsePtrOutput)
}

func (o PhpWorkloadOutput) SiteProfile() SiteProfileResponsePtrOutput {
	return o.ApplyT(func(v *PhpWorkload) SiteProfileResponsePtrOutput { return v.SiteProfile }).(SiteProfileResponsePtrOutput)
}

func (o PhpWorkloadOutput) Sku() SkuResponsePtrOutput {
	return o.ApplyT(func(v *PhpWorkload) SkuResponsePtrOutput { return v.Sku }).(SkuResponsePtrOutput)
}

func (o PhpWorkloadOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *PhpWorkload) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

func (o PhpWorkloadOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *PhpWorkload) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

func (o PhpWorkloadOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *PhpWorkload) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func (o PhpWorkloadOutput) WebNodesProfile() VmssNodesProfileResponseOutput {
	return o.ApplyT(func(v *PhpWorkload) VmssNodesProfileResponseOutput { return v.WebNodesProfile }).(VmssNodesProfileResponseOutput)
}

func init() {
	pulumi.RegisterOutputType(PhpWorkloadOutput{})
}
