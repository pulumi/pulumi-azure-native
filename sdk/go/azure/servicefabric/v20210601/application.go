


package v20210601

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type Application struct {
	pulumi.CustomResourceState

	Etag                      pulumi.StringOutput                                `pulumi:"etag"`
	Identity                  ManagedIdentityResponsePtrOutput                   `pulumi:"identity"`
	Location                  pulumi.StringPtrOutput                             `pulumi:"location"`
	ManagedIdentities         ApplicationUserAssignedIdentityResponseArrayOutput `pulumi:"managedIdentities"`
	MaximumNodes              pulumi.Float64PtrOutput                            `pulumi:"maximumNodes"`
	Metrics                   ApplicationMetricDescriptionResponseArrayOutput    `pulumi:"metrics"`
	MinimumNodes              pulumi.Float64PtrOutput                            `pulumi:"minimumNodes"`
	Name                      pulumi.StringOutput                                `pulumi:"name"`
	Parameters                pulumi.StringMapOutput                             `pulumi:"parameters"`
	ProvisioningState         pulumi.StringOutput                                `pulumi:"provisioningState"`
	RemoveApplicationCapacity pulumi.BoolPtrOutput                               `pulumi:"removeApplicationCapacity"`
	SystemData                SystemDataResponseOutput                           `pulumi:"systemData"`
	Tags                      pulumi.StringMapOutput                             `pulumi:"tags"`
	Type                      pulumi.StringOutput                                `pulumi:"type"`
	TypeName                  pulumi.StringPtrOutput                             `pulumi:"typeName"`
	TypeVersion               pulumi.StringPtrOutput                             `pulumi:"typeVersion"`
	UpgradePolicy             ApplicationUpgradePolicyResponsePtrOutput          `pulumi:"upgradePolicy"`
}


func NewApplication(ctx *pulumi.Context,
	name string, args *ApplicationArgs, opts ...pulumi.ResourceOption) (*Application, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ClusterName == nil {
		return nil, errors.New("invalid value for required argument 'ClusterName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.MaximumNodes == nil {
		args.MaximumNodes = pulumi.Float64Ptr(0.0)
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:servicefabric:Application"),
		},
		{
			Type: pulumi.String("azure-native:servicefabric/v20170701preview:Application"),
		},
		{
			Type: pulumi.String("azure-native:servicefabric/v20190301:Application"),
		},
		{
			Type: pulumi.String("azure-native:servicefabric/v20190301preview:Application"),
		},
		{
			Type: pulumi.String("azure-native:servicefabric/v20190601preview:Application"),
		},
		{
			Type: pulumi.String("azure-native:servicefabric/v20191101preview:Application"),
		},
		{
			Type: pulumi.String("azure-native:servicefabric/v20200301:Application"),
		},
		{
			Type: pulumi.String("azure-native:servicefabric/v20201201preview:Application"),
		},
	})
	opts = append(opts, aliases)
	var resource Application
	err := ctx.RegisterResource("azure-native:servicefabric/v20210601:Application", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetApplication(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ApplicationState, opts ...pulumi.ResourceOption) (*Application, error) {
	var resource Application
	err := ctx.ReadResource("azure-native:servicefabric/v20210601:Application", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type applicationState struct {
}

type ApplicationState struct {
}

func (ApplicationState) ElementType() reflect.Type {
	return reflect.TypeOf((*applicationState)(nil)).Elem()
}

type applicationArgs struct {
	ApplicationName           *string                           `pulumi:"applicationName"`
	ClusterName               string                            `pulumi:"clusterName"`
	Identity                  *ManagedIdentity                  `pulumi:"identity"`
	Location                  *string                           `pulumi:"location"`
	ManagedIdentities         []ApplicationUserAssignedIdentity `pulumi:"managedIdentities"`
	MaximumNodes              *float64                          `pulumi:"maximumNodes"`
	Metrics                   []ApplicationMetricDescription    `pulumi:"metrics"`
	MinimumNodes              *float64                          `pulumi:"minimumNodes"`
	Parameters                map[string]string                 `pulumi:"parameters"`
	RemoveApplicationCapacity *bool                             `pulumi:"removeApplicationCapacity"`
	ResourceGroupName         string                            `pulumi:"resourceGroupName"`
	Tags                      map[string]string                 `pulumi:"tags"`
	TypeName                  *string                           `pulumi:"typeName"`
	TypeVersion               *string                           `pulumi:"typeVersion"`
	UpgradePolicy             *ApplicationUpgradePolicy         `pulumi:"upgradePolicy"`
}


type ApplicationArgs struct {
	ApplicationName           pulumi.StringPtrInput
	ClusterName               pulumi.StringInput
	Identity                  ManagedIdentityPtrInput
	Location                  pulumi.StringPtrInput
	ManagedIdentities         ApplicationUserAssignedIdentityArrayInput
	MaximumNodes              pulumi.Float64PtrInput
	Metrics                   ApplicationMetricDescriptionArrayInput
	MinimumNodes              pulumi.Float64PtrInput
	Parameters                pulumi.StringMapInput
	RemoveApplicationCapacity pulumi.BoolPtrInput
	ResourceGroupName         pulumi.StringInput
	Tags                      pulumi.StringMapInput
	TypeName                  pulumi.StringPtrInput
	TypeVersion               pulumi.StringPtrInput
	UpgradePolicy             ApplicationUpgradePolicyPtrInput
}

func (ApplicationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*applicationArgs)(nil)).Elem()
}

type ApplicationInput interface {
	pulumi.Input

	ToApplicationOutput() ApplicationOutput
	ToApplicationOutputWithContext(ctx context.Context) ApplicationOutput
}

func (*Application) ElementType() reflect.Type {
	return reflect.TypeOf((*Application)(nil))
}

func (i *Application) ToApplicationOutput() ApplicationOutput {
	return i.ToApplicationOutputWithContext(context.Background())
}

func (i *Application) ToApplicationOutputWithContext(ctx context.Context) ApplicationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApplicationOutput)
}

type ApplicationOutput struct{ *pulumi.OutputState }

func (ApplicationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Application)(nil))
}

func (o ApplicationOutput) ToApplicationOutput() ApplicationOutput {
	return o
}

func (o ApplicationOutput) ToApplicationOutputWithContext(ctx context.Context) ApplicationOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(ApplicationOutput{})
}
