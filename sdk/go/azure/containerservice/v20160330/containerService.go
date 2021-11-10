


package v20160330

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type ContainerService struct {
	pulumi.CustomResourceState

	AgentPoolProfiles   ContainerServiceAgentPoolProfileResponseArrayOutput  `pulumi:"agentPoolProfiles"`
	DiagnosticsProfile  ContainerServiceDiagnosticsProfileResponsePtrOutput  `pulumi:"diagnosticsProfile"`
	LinuxProfile        ContainerServiceLinuxProfileResponseOutput           `pulumi:"linuxProfile"`
	Location            pulumi.StringOutput                                  `pulumi:"location"`
	MasterProfile       ContainerServiceMasterProfileResponseOutput          `pulumi:"masterProfile"`
	Name                pulumi.StringOutput                                  `pulumi:"name"`
	OrchestratorProfile ContainerServiceOrchestratorProfileResponsePtrOutput `pulumi:"orchestratorProfile"`
	ProvisioningState   pulumi.StringOutput                                  `pulumi:"provisioningState"`
	Tags                pulumi.StringMapOutput                               `pulumi:"tags"`
	Type                pulumi.StringOutput                                  `pulumi:"type"`
	WindowsProfile      ContainerServiceWindowsProfileResponsePtrOutput      `pulumi:"windowsProfile"`
}


func NewContainerService(ctx *pulumi.Context,
	name string, args *ContainerServiceArgs, opts ...pulumi.ResourceOption) (*ContainerService, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AgentPoolProfiles == nil {
		return nil, errors.New("invalid value for required argument 'AgentPoolProfiles'")
	}
	if args.LinuxProfile == nil {
		return nil, errors.New("invalid value for required argument 'LinuxProfile'")
	}
	if args.MasterProfile == nil {
		return nil, errors.New("invalid value for required argument 'MasterProfile'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:containerservice/v20151101preview:ContainerService"),
		},
		{
			Type: pulumi.String("azure-native:containerservice/v20160930:ContainerService"),
		},
		{
			Type: pulumi.String("azure-native:containerservice/v20170131:ContainerService"),
		},
	})
	opts = append(opts, aliases)
	var resource ContainerService
	err := ctx.RegisterResource("azure-native:containerservice/v20160330:ContainerService", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetContainerService(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ContainerServiceState, opts ...pulumi.ResourceOption) (*ContainerService, error) {
	var resource ContainerService
	err := ctx.ReadResource("azure-native:containerservice/v20160330:ContainerService", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type containerServiceState struct {
}

type ContainerServiceState struct {
}

func (ContainerServiceState) ElementType() reflect.Type {
	return reflect.TypeOf((*containerServiceState)(nil)).Elem()
}

type containerServiceArgs struct {
	AgentPoolProfiles    []ContainerServiceAgentPoolProfile   `pulumi:"agentPoolProfiles"`
	ContainerServiceName *string                              `pulumi:"containerServiceName"`
	DiagnosticsProfile   *ContainerServiceDiagnosticsProfile  `pulumi:"diagnosticsProfile"`
	LinuxProfile         ContainerServiceLinuxProfile         `pulumi:"linuxProfile"`
	Location             *string                              `pulumi:"location"`
	MasterProfile        ContainerServiceMasterProfile        `pulumi:"masterProfile"`
	OrchestratorProfile  *ContainerServiceOrchestratorProfile `pulumi:"orchestratorProfile"`
	ResourceGroupName    string                               `pulumi:"resourceGroupName"`
	Tags                 map[string]string                    `pulumi:"tags"`
	WindowsProfile       *ContainerServiceWindowsProfile      `pulumi:"windowsProfile"`
}


type ContainerServiceArgs struct {
	AgentPoolProfiles    ContainerServiceAgentPoolProfileArrayInput
	ContainerServiceName pulumi.StringPtrInput
	DiagnosticsProfile   ContainerServiceDiagnosticsProfilePtrInput
	LinuxProfile         ContainerServiceLinuxProfileInput
	Location             pulumi.StringPtrInput
	MasterProfile        ContainerServiceMasterProfileInput
	OrchestratorProfile  ContainerServiceOrchestratorProfilePtrInput
	ResourceGroupName    pulumi.StringInput
	Tags                 pulumi.StringMapInput
	WindowsProfile       ContainerServiceWindowsProfilePtrInput
}

func (ContainerServiceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*containerServiceArgs)(nil)).Elem()
}

type ContainerServiceInput interface {
	pulumi.Input

	ToContainerServiceOutput() ContainerServiceOutput
	ToContainerServiceOutputWithContext(ctx context.Context) ContainerServiceOutput
}

func (*ContainerService) ElementType() reflect.Type {
	return reflect.TypeOf((*ContainerService)(nil))
}

func (i *ContainerService) ToContainerServiceOutput() ContainerServiceOutput {
	return i.ToContainerServiceOutputWithContext(context.Background())
}

func (i *ContainerService) ToContainerServiceOutputWithContext(ctx context.Context) ContainerServiceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ContainerServiceOutput)
}

type ContainerServiceOutput struct{ *pulumi.OutputState }

func (ContainerServiceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ContainerService)(nil))
}

func (o ContainerServiceOutput) ToContainerServiceOutput() ContainerServiceOutput {
	return o
}

func (o ContainerServiceOutput) ToContainerServiceOutputWithContext(ctx context.Context) ContainerServiceOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(ContainerServiceOutput{})
}
