


package v20160515

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type VirtualMachine struct {
	pulumi.CustomResourceState

	AllowClaim                   pulumi.BoolPtrOutput                                `pulumi:"allowClaim"`
	ApplicableSchedule           ApplicableScheduleResponsePtrOutput                 `pulumi:"applicableSchedule"`
	ArtifactDeploymentStatus     ArtifactDeploymentStatusPropertiesResponsePtrOutput `pulumi:"artifactDeploymentStatus"`
	Artifacts                    ArtifactInstallPropertiesResponseArrayOutput        `pulumi:"artifacts"`
	ComputeId                    pulumi.StringOutput                                 `pulumi:"computeId"`
	ComputeVm                    ComputeVmPropertiesResponsePtrOutput                `pulumi:"computeVm"`
	CreatedByUser                pulumi.StringPtrOutput                              `pulumi:"createdByUser"`
	CreatedByUserId              pulumi.StringPtrOutput                              `pulumi:"createdByUserId"`
	CreatedDate                  pulumi.StringPtrOutput                              `pulumi:"createdDate"`
	CustomImageId                pulumi.StringPtrOutput                              `pulumi:"customImageId"`
	DisallowPublicIpAddress      pulumi.BoolPtrOutput                                `pulumi:"disallowPublicIpAddress"`
	EnvironmentId                pulumi.StringPtrOutput                              `pulumi:"environmentId"`
	ExpirationDate               pulumi.StringPtrOutput                              `pulumi:"expirationDate"`
	Fqdn                         pulumi.StringPtrOutput                              `pulumi:"fqdn"`
	GalleryImageReference        GalleryImageReferenceResponsePtrOutput              `pulumi:"galleryImageReference"`
	IsAuthenticationWithSshKey   pulumi.BoolPtrOutput                                `pulumi:"isAuthenticationWithSshKey"`
	LabSubnetName                pulumi.StringPtrOutput                              `pulumi:"labSubnetName"`
	LabVirtualNetworkId          pulumi.StringPtrOutput                              `pulumi:"labVirtualNetworkId"`
	Location                     pulumi.StringPtrOutput                              `pulumi:"location"`
	Name                         pulumi.StringOutput                                 `pulumi:"name"`
	NetworkInterface             NetworkInterfacePropertiesResponsePtrOutput         `pulumi:"networkInterface"`
	Notes                        pulumi.StringPtrOutput                              `pulumi:"notes"`
	OsType                       pulumi.StringPtrOutput                              `pulumi:"osType"`
	OwnerObjectId                pulumi.StringPtrOutput                              `pulumi:"ownerObjectId"`
	OwnerUserPrincipalName       pulumi.StringPtrOutput                              `pulumi:"ownerUserPrincipalName"`
	Password                     pulumi.StringPtrOutput                              `pulumi:"password"`
	ProvisioningState            pulumi.StringPtrOutput                              `pulumi:"provisioningState"`
	Size                         pulumi.StringPtrOutput                              `pulumi:"size"`
	SshKey                       pulumi.StringPtrOutput                              `pulumi:"sshKey"`
	StorageType                  pulumi.StringPtrOutput                              `pulumi:"storageType"`
	Tags                         pulumi.StringMapOutput                              `pulumi:"tags"`
	Type                         pulumi.StringOutput                                 `pulumi:"type"`
	UniqueIdentifier             pulumi.StringPtrOutput                              `pulumi:"uniqueIdentifier"`
	UserName                     pulumi.StringPtrOutput                              `pulumi:"userName"`
	VirtualMachineCreationSource pulumi.StringPtrOutput                              `pulumi:"virtualMachineCreationSource"`
}


func NewVirtualMachine(ctx *pulumi.Context,
	name string, args *VirtualMachineArgs, opts ...pulumi.ResourceOption) (*VirtualMachine, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.LabName == nil {
		return nil, errors.New("invalid value for required argument 'LabName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:devtestlab:VirtualMachine"),
		},
		{
			Type: pulumi.String("azure-native:devtestlab/v20150521preview:VirtualMachine"),
		},
		{
			Type: pulumi.String("azure-native:devtestlab/v20180915:VirtualMachine"),
		},
	})
	opts = append(opts, aliases)
	var resource VirtualMachine
	err := ctx.RegisterResource("azure-native:devtestlab/v20160515:VirtualMachine", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetVirtualMachine(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *VirtualMachineState, opts ...pulumi.ResourceOption) (*VirtualMachine, error) {
	var resource VirtualMachine
	err := ctx.ReadResource("azure-native:devtestlab/v20160515:VirtualMachine", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type virtualMachineState struct {
}

type VirtualMachineState struct {
}

func (VirtualMachineState) ElementType() reflect.Type {
	return reflect.TypeOf((*virtualMachineState)(nil)).Elem()
}

type virtualMachineArgs struct {
	AllowClaim                   *bool                               `pulumi:"allowClaim"`
	ApplicableSchedule           *ApplicableSchedule                 `pulumi:"applicableSchedule"`
	ArtifactDeploymentStatus     *ArtifactDeploymentStatusProperties `pulumi:"artifactDeploymentStatus"`
	Artifacts                    []ArtifactInstallProperties         `pulumi:"artifacts"`
	ComputeVm                    *ComputeVmProperties                `pulumi:"computeVm"`
	CreatedByUser                *string                             `pulumi:"createdByUser"`
	CreatedByUserId              *string                             `pulumi:"createdByUserId"`
	CreatedDate                  *string                             `pulumi:"createdDate"`
	CustomImageId                *string                             `pulumi:"customImageId"`
	DisallowPublicIpAddress      *bool                               `pulumi:"disallowPublicIpAddress"`
	EnvironmentId                *string                             `pulumi:"environmentId"`
	ExpirationDate               *string                             `pulumi:"expirationDate"`
	Fqdn                         *string                             `pulumi:"fqdn"`
	GalleryImageReference        *GalleryImageReference              `pulumi:"galleryImageReference"`
	IsAuthenticationWithSshKey   *bool                               `pulumi:"isAuthenticationWithSshKey"`
	LabName                      string                              `pulumi:"labName"`
	LabSubnetName                *string                             `pulumi:"labSubnetName"`
	LabVirtualNetworkId          *string                             `pulumi:"labVirtualNetworkId"`
	Location                     *string                             `pulumi:"location"`
	Name                         *string                             `pulumi:"name"`
	NetworkInterface             *NetworkInterfaceProperties         `pulumi:"networkInterface"`
	Notes                        *string                             `pulumi:"notes"`
	OsType                       *string                             `pulumi:"osType"`
	OwnerObjectId                *string                             `pulumi:"ownerObjectId"`
	OwnerUserPrincipalName       *string                             `pulumi:"ownerUserPrincipalName"`
	Password                     *string                             `pulumi:"password"`
	ProvisioningState            *string                             `pulumi:"provisioningState"`
	ResourceGroupName            string                              `pulumi:"resourceGroupName"`
	Size                         *string                             `pulumi:"size"`
	SshKey                       *string                             `pulumi:"sshKey"`
	StorageType                  *string                             `pulumi:"storageType"`
	Tags                         map[string]string                   `pulumi:"tags"`
	UniqueIdentifier             *string                             `pulumi:"uniqueIdentifier"`
	UserName                     *string                             `pulumi:"userName"`
	VirtualMachineCreationSource *string                             `pulumi:"virtualMachineCreationSource"`
}


type VirtualMachineArgs struct {
	AllowClaim                   pulumi.BoolPtrInput
	ApplicableSchedule           ApplicableSchedulePtrInput
	ArtifactDeploymentStatus     ArtifactDeploymentStatusPropertiesPtrInput
	Artifacts                    ArtifactInstallPropertiesArrayInput
	ComputeVm                    ComputeVmPropertiesPtrInput
	CreatedByUser                pulumi.StringPtrInput
	CreatedByUserId              pulumi.StringPtrInput
	CreatedDate                  pulumi.StringPtrInput
	CustomImageId                pulumi.StringPtrInput
	DisallowPublicIpAddress      pulumi.BoolPtrInput
	EnvironmentId                pulumi.StringPtrInput
	ExpirationDate               pulumi.StringPtrInput
	Fqdn                         pulumi.StringPtrInput
	GalleryImageReference        GalleryImageReferencePtrInput
	IsAuthenticationWithSshKey   pulumi.BoolPtrInput
	LabName                      pulumi.StringInput
	LabSubnetName                pulumi.StringPtrInput
	LabVirtualNetworkId          pulumi.StringPtrInput
	Location                     pulumi.StringPtrInput
	Name                         pulumi.StringPtrInput
	NetworkInterface             NetworkInterfacePropertiesPtrInput
	Notes                        pulumi.StringPtrInput
	OsType                       pulumi.StringPtrInput
	OwnerObjectId                pulumi.StringPtrInput
	OwnerUserPrincipalName       pulumi.StringPtrInput
	Password                     pulumi.StringPtrInput
	ProvisioningState            pulumi.StringPtrInput
	ResourceGroupName            pulumi.StringInput
	Size                         pulumi.StringPtrInput
	SshKey                       pulumi.StringPtrInput
	StorageType                  pulumi.StringPtrInput
	Tags                         pulumi.StringMapInput
	UniqueIdentifier             pulumi.StringPtrInput
	UserName                     pulumi.StringPtrInput
	VirtualMachineCreationSource pulumi.StringPtrInput
}

func (VirtualMachineArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*virtualMachineArgs)(nil)).Elem()
}

type VirtualMachineInput interface {
	pulumi.Input

	ToVirtualMachineOutput() VirtualMachineOutput
	ToVirtualMachineOutputWithContext(ctx context.Context) VirtualMachineOutput
}

func (*VirtualMachine) ElementType() reflect.Type {
	return reflect.TypeOf((**VirtualMachine)(nil)).Elem()
}

func (i *VirtualMachine) ToVirtualMachineOutput() VirtualMachineOutput {
	return i.ToVirtualMachineOutputWithContext(context.Background())
}

func (i *VirtualMachine) ToVirtualMachineOutputWithContext(ctx context.Context) VirtualMachineOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualMachineOutput)
}

type VirtualMachineOutput struct{ *pulumi.OutputState }

func (VirtualMachineOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**VirtualMachine)(nil)).Elem()
}

func (o VirtualMachineOutput) ToVirtualMachineOutput() VirtualMachineOutput {
	return o
}

func (o VirtualMachineOutput) ToVirtualMachineOutputWithContext(ctx context.Context) VirtualMachineOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(VirtualMachineOutput{})
}
