


package v20201201

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type KubernetesRole struct {
	pulumi.CustomResourceState

	HostPlatform            pulumi.StringOutput                   `pulumi:"hostPlatform"`
	HostPlatformType        pulumi.StringOutput                   `pulumi:"hostPlatformType"`
	Kind                    pulumi.StringOutput                   `pulumi:"kind"`
	KubernetesClusterInfo   KubernetesClusterInfoResponseOutput   `pulumi:"kubernetesClusterInfo"`
	KubernetesRoleResources KubernetesRoleResourcesResponseOutput `pulumi:"kubernetesRoleResources"`
	Name                    pulumi.StringOutput                   `pulumi:"name"`
	ProvisioningState       pulumi.StringOutput                   `pulumi:"provisioningState"`
	RoleStatus              pulumi.StringOutput                   `pulumi:"roleStatus"`
	SystemData              SystemDataResponseOutput              `pulumi:"systemData"`
	Type                    pulumi.StringOutput                   `pulumi:"type"`
}


func NewKubernetesRole(ctx *pulumi.Context,
	name string, args *KubernetesRoleArgs, opts ...pulumi.ResourceOption) (*KubernetesRole, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.DeviceName == nil {
		return nil, errors.New("invalid value for required argument 'DeviceName'")
	}
	if args.HostPlatform == nil {
		return nil, errors.New("invalid value for required argument 'HostPlatform'")
	}
	if args.Kind == nil {
		return nil, errors.New("invalid value for required argument 'Kind'")
	}
	if args.KubernetesClusterInfo == nil {
		return nil, errors.New("invalid value for required argument 'KubernetesClusterInfo'")
	}
	if args.KubernetesRoleResources == nil {
		return nil, errors.New("invalid value for required argument 'KubernetesRoleResources'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.RoleStatus == nil {
		return nil, errors.New("invalid value for required argument 'RoleStatus'")
	}
	args.Kind = pulumi.String("Kubernetes")
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:databoxedge:KubernetesRole"),
		},
		{
			Type: pulumi.String("azure-native:databoxedge/v20190301:KubernetesRole"),
		},
		{
			Type: pulumi.String("azure-native:databoxedge/v20190701:KubernetesRole"),
		},
		{
			Type: pulumi.String("azure-native:databoxedge/v20190801:KubernetesRole"),
		},
		{
			Type: pulumi.String("azure-native:databoxedge/v20200501preview:KubernetesRole"),
		},
		{
			Type: pulumi.String("azure-native:databoxedge/v20200901:KubernetesRole"),
		},
		{
			Type: pulumi.String("azure-native:databoxedge/v20200901preview:KubernetesRole"),
		},
		{
			Type: pulumi.String("azure-native:databoxedge/v20210201:KubernetesRole"),
		},
		{
			Type: pulumi.String("azure-native:databoxedge/v20210201preview:KubernetesRole"),
		},
		{
			Type: pulumi.String("azure-native:databoxedge/v20210601:KubernetesRole"),
		},
		{
			Type: pulumi.String("azure-native:databoxedge/v20210601preview:KubernetesRole"),
		},
	})
	opts = append(opts, aliases)
	var resource KubernetesRole
	err := ctx.RegisterResource("azure-native:databoxedge/v20201201:KubernetesRole", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetKubernetesRole(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *KubernetesRoleState, opts ...pulumi.ResourceOption) (*KubernetesRole, error) {
	var resource KubernetesRole
	err := ctx.ReadResource("azure-native:databoxedge/v20201201:KubernetesRole", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type kubernetesRoleState struct {
}

type KubernetesRoleState struct {
}

func (KubernetesRoleState) ElementType() reflect.Type {
	return reflect.TypeOf((*kubernetesRoleState)(nil)).Elem()
}

type kubernetesRoleArgs struct {
	DeviceName              string                  `pulumi:"deviceName"`
	HostPlatform            string                  `pulumi:"hostPlatform"`
	Kind                    string                  `pulumi:"kind"`
	KubernetesClusterInfo   KubernetesClusterInfo   `pulumi:"kubernetesClusterInfo"`
	KubernetesRoleResources KubernetesRoleResources `pulumi:"kubernetesRoleResources"`
	Name                    *string                 `pulumi:"name"`
	ResourceGroupName       string                  `pulumi:"resourceGroupName"`
	RoleStatus              string                  `pulumi:"roleStatus"`
}


type KubernetesRoleArgs struct {
	DeviceName              pulumi.StringInput
	HostPlatform            pulumi.StringInput
	Kind                    pulumi.StringInput
	KubernetesClusterInfo   KubernetesClusterInfoInput
	KubernetesRoleResources KubernetesRoleResourcesInput
	Name                    pulumi.StringPtrInput
	ResourceGroupName       pulumi.StringInput
	RoleStatus              pulumi.StringInput
}

func (KubernetesRoleArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*kubernetesRoleArgs)(nil)).Elem()
}

type KubernetesRoleInput interface {
	pulumi.Input

	ToKubernetesRoleOutput() KubernetesRoleOutput
	ToKubernetesRoleOutputWithContext(ctx context.Context) KubernetesRoleOutput
}

func (*KubernetesRole) ElementType() reflect.Type {
	return reflect.TypeOf((*KubernetesRole)(nil))
}

func (i *KubernetesRole) ToKubernetesRoleOutput() KubernetesRoleOutput {
	return i.ToKubernetesRoleOutputWithContext(context.Background())
}

func (i *KubernetesRole) ToKubernetesRoleOutputWithContext(ctx context.Context) KubernetesRoleOutput {
	return pulumi.ToOutputWithContext(ctx, i).(KubernetesRoleOutput)
}

type KubernetesRoleOutput struct{ *pulumi.OutputState }

func (KubernetesRoleOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*KubernetesRole)(nil))
}

func (o KubernetesRoleOutput) ToKubernetesRoleOutput() KubernetesRoleOutput {
	return o
}

func (o KubernetesRoleOutput) ToKubernetesRoleOutputWithContext(ctx context.Context) KubernetesRoleOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(KubernetesRoleOutput{})
}
