


package v20220101

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type Application struct {
	pulumi.CustomResourceState

	Identity          ManagedIdentityResponsePtrOutput                   `pulumi:"identity"`
	Location          pulumi.StringPtrOutput                             `pulumi:"location"`
	ManagedIdentities ApplicationUserAssignedIdentityResponseArrayOutput `pulumi:"managedIdentities"`
	Name              pulumi.StringOutput                                `pulumi:"name"`
	Parameters        pulumi.StringMapOutput                             `pulumi:"parameters"`
	ProvisioningState pulumi.StringOutput                                `pulumi:"provisioningState"`
	SystemData        SystemDataResponseOutput                           `pulumi:"systemData"`
	Tags              pulumi.StringMapOutput                             `pulumi:"tags"`
	Type              pulumi.StringOutput                                `pulumi:"type"`
	UpgradePolicy     ApplicationUpgradePolicyResponsePtrOutput          `pulumi:"upgradePolicy"`
	Version           pulumi.StringPtrOutput                             `pulumi:"version"`
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
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:servicefabric/v20210101preview:Application"),
		},
		{
			Type: pulumi.String("azure-native:servicefabric/v20210501:Application"),
		},
		{
			Type: pulumi.String("azure-native:servicefabric/v20210701preview:Application"),
		},
		{
			Type: pulumi.String("azure-native:servicefabric/v20210901privatepreview:Application"),
		},
		{
			Type: pulumi.String("azure-native:servicefabric/v20211101preview:Application"),
		},
		{
			Type: pulumi.String("azure-native:servicefabric/v20220201preview:Application"),
		},
		{
			Type: pulumi.String("azure-native:servicefabric/v20220601preview:Application"),
		},
		{
			Type: pulumi.String("azure-native:servicefabric/v20220801preview:Application"),
		},
	})
	opts = append(opts, aliases)
	var resource Application
	err := ctx.RegisterResource("azure-native:servicefabric/v20220101:Application", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetApplication(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ApplicationState, opts ...pulumi.ResourceOption) (*Application, error) {
	var resource Application
	err := ctx.ReadResource("azure-native:servicefabric/v20220101:Application", name, id, state, &resource, opts...)
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
	ApplicationName   *string                           `pulumi:"applicationName"`
	ClusterName       string                            `pulumi:"clusterName"`
	Identity          *ManagedIdentity                  `pulumi:"identity"`
	Location          *string                           `pulumi:"location"`
	ManagedIdentities []ApplicationUserAssignedIdentity `pulumi:"managedIdentities"`
	Parameters        map[string]string                 `pulumi:"parameters"`
	ResourceGroupName string                            `pulumi:"resourceGroupName"`
	Tags              map[string]string                 `pulumi:"tags"`
	UpgradePolicy     *ApplicationUpgradePolicy         `pulumi:"upgradePolicy"`
	Version           *string                           `pulumi:"version"`
}


type ApplicationArgs struct {
	ApplicationName   pulumi.StringPtrInput
	ClusterName       pulumi.StringInput
	Identity          ManagedIdentityPtrInput
	Location          pulumi.StringPtrInput
	ManagedIdentities ApplicationUserAssignedIdentityArrayInput
	Parameters        pulumi.StringMapInput
	ResourceGroupName pulumi.StringInput
	Tags              pulumi.StringMapInput
	UpgradePolicy     ApplicationUpgradePolicyPtrInput
	Version           pulumi.StringPtrInput
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
	return reflect.TypeOf((**Application)(nil)).Elem()
}

func (i *Application) ToApplicationOutput() ApplicationOutput {
	return i.ToApplicationOutputWithContext(context.Background())
}

func (i *Application) ToApplicationOutputWithContext(ctx context.Context) ApplicationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApplicationOutput)
}

type ApplicationOutput struct{ *pulumi.OutputState }

func (ApplicationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Application)(nil)).Elem()
}

func (o ApplicationOutput) ToApplicationOutput() ApplicationOutput {
	return o
}

func (o ApplicationOutput) ToApplicationOutputWithContext(ctx context.Context) ApplicationOutput {
	return o
}

func (o ApplicationOutput) Identity() ManagedIdentityResponsePtrOutput {
	return o.ApplyT(func(v *Application) ManagedIdentityResponsePtrOutput { return v.Identity }).(ManagedIdentityResponsePtrOutput)
}

func (o ApplicationOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Application) pulumi.StringPtrOutput { return v.Location }).(pulumi.StringPtrOutput)
}

func (o ApplicationOutput) ManagedIdentities() ApplicationUserAssignedIdentityResponseArrayOutput {
	return o.ApplyT(func(v *Application) ApplicationUserAssignedIdentityResponseArrayOutput { return v.ManagedIdentities }).(ApplicationUserAssignedIdentityResponseArrayOutput)
}

func (o ApplicationOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Application) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o ApplicationOutput) Parameters() pulumi.StringMapOutput {
	return o.ApplyT(func(v *Application) pulumi.StringMapOutput { return v.Parameters }).(pulumi.StringMapOutput)
}

func (o ApplicationOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *Application) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o ApplicationOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *Application) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

func (o ApplicationOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *Application) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

func (o ApplicationOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *Application) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func (o ApplicationOutput) UpgradePolicy() ApplicationUpgradePolicyResponsePtrOutput {
	return o.ApplyT(func(v *Application) ApplicationUpgradePolicyResponsePtrOutput { return v.UpgradePolicy }).(ApplicationUpgradePolicyResponsePtrOutput)
}

func (o ApplicationOutput) Version() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Application) pulumi.StringPtrOutput { return v.Version }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(ApplicationOutput{})
}
