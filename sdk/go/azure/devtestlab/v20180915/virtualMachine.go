


package v20180915

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type VirtualMachine struct {
	pulumi.CustomResourceState

	AllowClaim                   pulumi.BoolPtrOutput                             `pulumi:"allowClaim"`
	ApplicableSchedule           ApplicableScheduleResponseOutput                 `pulumi:"applicableSchedule"`
	ArtifactDeploymentStatus     ArtifactDeploymentStatusPropertiesResponseOutput `pulumi:"artifactDeploymentStatus"`
	Artifacts                    ArtifactInstallPropertiesResponseArrayOutput     `pulumi:"artifacts"`
	ComputeId                    pulumi.StringOutput                              `pulumi:"computeId"`
	ComputeVm                    ComputeVmPropertiesResponseOutput                `pulumi:"computeVm"`
	CreatedByUser                pulumi.StringOutput                              `pulumi:"createdByUser"`
	CreatedByUserId              pulumi.StringOutput                              `pulumi:"createdByUserId"`
	CreatedDate                  pulumi.StringPtrOutput                           `pulumi:"createdDate"`
	CustomImageId                pulumi.StringPtrOutput                           `pulumi:"customImageId"`
	DataDiskParameters           DataDiskPropertiesResponseArrayOutput            `pulumi:"dataDiskParameters"`
	DisallowPublicIpAddress      pulumi.BoolPtrOutput                             `pulumi:"disallowPublicIpAddress"`
	EnvironmentId                pulumi.StringPtrOutput                           `pulumi:"environmentId"`
	ExpirationDate               pulumi.StringPtrOutput                           `pulumi:"expirationDate"`
	Fqdn                         pulumi.StringOutput                              `pulumi:"fqdn"`
	GalleryImageReference        GalleryImageReferenceResponsePtrOutput           `pulumi:"galleryImageReference"`
	IsAuthenticationWithSshKey   pulumi.BoolPtrOutput                             `pulumi:"isAuthenticationWithSshKey"`
	LabSubnetName                pulumi.StringPtrOutput                           `pulumi:"labSubnetName"`
	LabVirtualNetworkId          pulumi.StringPtrOutput                           `pulumi:"labVirtualNetworkId"`
	LastKnownPowerState          pulumi.StringOutput                              `pulumi:"lastKnownPowerState"`
	Location                     pulumi.StringPtrOutput                           `pulumi:"location"`
	Name                         pulumi.StringOutput                              `pulumi:"name"`
	NetworkInterface             NetworkInterfacePropertiesResponsePtrOutput      `pulumi:"networkInterface"`
	Notes                        pulumi.StringPtrOutput                           `pulumi:"notes"`
	OsType                       pulumi.StringOutput                              `pulumi:"osType"`
	OwnerObjectId                pulumi.StringPtrOutput                           `pulumi:"ownerObjectId"`
	OwnerUserPrincipalName       pulumi.StringPtrOutput                           `pulumi:"ownerUserPrincipalName"`
	Password                     pulumi.StringPtrOutput                           `pulumi:"password"`
	PlanId                       pulumi.StringPtrOutput                           `pulumi:"planId"`
	ProvisioningState            pulumi.StringOutput                              `pulumi:"provisioningState"`
	ScheduleParameters           ScheduleCreationParameterResponseArrayOutput     `pulumi:"scheduleParameters"`
	Size                         pulumi.StringPtrOutput                           `pulumi:"size"`
	SshKey                       pulumi.StringPtrOutput                           `pulumi:"sshKey"`
	StorageType                  pulumi.StringPtrOutput                           `pulumi:"storageType"`
	Tags                         pulumi.StringMapOutput                           `pulumi:"tags"`
	Type                         pulumi.StringOutput                              `pulumi:"type"`
	UniqueIdentifier             pulumi.StringOutput                              `pulumi:"uniqueIdentifier"`
	UserName                     pulumi.StringPtrOutput                           `pulumi:"userName"`
	VirtualMachineCreationSource pulumi.StringOutput                              `pulumi:"virtualMachineCreationSource"`
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
	if args.AllowClaim == nil {
		args.AllowClaim = pulumi.BoolPtr(false)
	}
	if args.DisallowPublicIpAddress == nil {
		args.DisallowPublicIpAddress = pulumi.BoolPtr(false)
	}
	if args.OwnerObjectId == nil {
		args.OwnerObjectId = pulumi.StringPtr("dynamicValue")
	}
	if args.StorageType == nil {
		args.StorageType = pulumi.StringPtr("labStorageType")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-nextgen:devtestlab/v20180915:VirtualMachine"),
		},
		{
			Type: pulumi.String("azure-native:devtestlab:VirtualMachine"),
		},
		{
			Type: pulumi.String("azure-nextgen:devtestlab:VirtualMachine"),
		},
		{
			Type: pulumi.String("azure-native:devtestlab/v20150521preview:VirtualMachine"),
		},
		{
			Type: pulumi.String("azure-nextgen:devtestlab/v20150521preview:VirtualMachine"),
		},
		{
			Type: pulumi.String("azure-native:devtestlab/v20160515:VirtualMachine"),
		},
		{
			Type: pulumi.String("azure-nextgen:devtestlab/v20160515:VirtualMachine"),
		},
	})
	opts = append(opts, aliases)
	var resource VirtualMachine
	err := ctx.RegisterResource("azure-native:devtestlab/v20180915:VirtualMachine", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetVirtualMachine(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *VirtualMachineState, opts ...pulumi.ResourceOption) (*VirtualMachine, error) {
	var resource VirtualMachine
	err := ctx.ReadResource("azure-native:devtestlab/v20180915:VirtualMachine", name, id, state, &resource, opts...)
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
	AllowClaim                 *bool                       `pulumi:"allowClaim"`
	Artifacts                  []ArtifactInstallProperties `pulumi:"artifacts"`
	CreatedDate                *string                     `pulumi:"createdDate"`
	CustomImageId              *string                     `pulumi:"customImageId"`
	DataDiskParameters         []DataDiskProperties        `pulumi:"dataDiskParameters"`
	DisallowPublicIpAddress    *bool                       `pulumi:"disallowPublicIpAddress"`
	EnvironmentId              *string                     `pulumi:"environmentId"`
	ExpirationDate             *string                     `pulumi:"expirationDate"`
	GalleryImageReference      *GalleryImageReference      `pulumi:"galleryImageReference"`
	IsAuthenticationWithSshKey *bool                       `pulumi:"isAuthenticationWithSshKey"`
	LabName                    string                      `pulumi:"labName"`
	LabSubnetName              *string                     `pulumi:"labSubnetName"`
	LabVirtualNetworkId        *string                     `pulumi:"labVirtualNetworkId"`
	Location                   *string                     `pulumi:"location"`
	Name                       *string                     `pulumi:"name"`
	NetworkInterface           *NetworkInterfaceProperties `pulumi:"networkInterface"`
	Notes                      *string                     `pulumi:"notes"`
	OwnerObjectId              *string                     `pulumi:"ownerObjectId"`
	OwnerUserPrincipalName     *string                     `pulumi:"ownerUserPrincipalName"`
	Password                   *string                     `pulumi:"password"`
	PlanId                     *string                     `pulumi:"planId"`
	ResourceGroupName          string                      `pulumi:"resourceGroupName"`
	ScheduleParameters         []ScheduleCreationParameter `pulumi:"scheduleParameters"`
	Size                       *string                     `pulumi:"size"`
	SshKey                     *string                     `pulumi:"sshKey"`
	StorageType                *string                     `pulumi:"storageType"`
	Tags                       map[string]string           `pulumi:"tags"`
	UserName                   *string                     `pulumi:"userName"`
}


type VirtualMachineArgs struct {
	AllowClaim                 pulumi.BoolPtrInput
	Artifacts                  ArtifactInstallPropertiesArrayInput
	CreatedDate                pulumi.StringPtrInput
	CustomImageId              pulumi.StringPtrInput
	DataDiskParameters         DataDiskPropertiesArrayInput
	DisallowPublicIpAddress    pulumi.BoolPtrInput
	EnvironmentId              pulumi.StringPtrInput
	ExpirationDate             pulumi.StringPtrInput
	GalleryImageReference      GalleryImageReferencePtrInput
	IsAuthenticationWithSshKey pulumi.BoolPtrInput
	LabName                    pulumi.StringInput
	LabSubnetName              pulumi.StringPtrInput
	LabVirtualNetworkId        pulumi.StringPtrInput
	Location                   pulumi.StringPtrInput
	Name                       pulumi.StringPtrInput
	NetworkInterface           NetworkInterfacePropertiesPtrInput
	Notes                      pulumi.StringPtrInput
	OwnerObjectId              pulumi.StringPtrInput
	OwnerUserPrincipalName     pulumi.StringPtrInput
	Password                   pulumi.StringPtrInput
	PlanId                     pulumi.StringPtrInput
	ResourceGroupName          pulumi.StringInput
	ScheduleParameters         ScheduleCreationParameterArrayInput
	Size                       pulumi.StringPtrInput
	SshKey                     pulumi.StringPtrInput
	StorageType                pulumi.StringPtrInput
	Tags                       pulumi.StringMapInput
	UserName                   pulumi.StringPtrInput
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
	return reflect.TypeOf((*VirtualMachine)(nil))
}

func (i *VirtualMachine) ToVirtualMachineOutput() VirtualMachineOutput {
	return i.ToVirtualMachineOutputWithContext(context.Background())
}

func (i *VirtualMachine) ToVirtualMachineOutputWithContext(ctx context.Context) VirtualMachineOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualMachineOutput)
}

type VirtualMachineOutput struct{ *pulumi.OutputState }

func (VirtualMachineOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VirtualMachine)(nil))
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
