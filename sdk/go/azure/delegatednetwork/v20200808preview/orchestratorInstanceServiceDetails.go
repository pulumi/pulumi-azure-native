


package v20200808preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)


type OrchestratorInstanceServiceDetails struct {
	pulumi.CustomResourceState

	ApiServerEndpoint    pulumi.StringPtrOutput                `pulumi:"apiServerEndpoint"`
	ClusterRootCA        pulumi.StringPtrOutput                `pulumi:"clusterRootCA"`
	ControllerDetails    ControllerDetailsResponseOutput       `pulumi:"controllerDetails"`
	Identity             OrchestratorIdentityResponsePtrOutput `pulumi:"identity"`
	Kind                 pulumi.StringOutput                   `pulumi:"kind"`
	Location             pulumi.StringPtrOutput                `pulumi:"location"`
	Name                 pulumi.StringOutput                   `pulumi:"name"`
	OrchestratorAppId    pulumi.StringPtrOutput                `pulumi:"orchestratorAppId"`
	OrchestratorTenantId pulumi.StringPtrOutput                `pulumi:"orchestratorTenantId"`
	ProvisioningState    pulumi.StringOutput                   `pulumi:"provisioningState"`
	ResourceGuid         pulumi.StringOutput                   `pulumi:"resourceGuid"`
	Tags                 pulumi.StringMapOutput                `pulumi:"tags"`
	Type                 pulumi.StringOutput                   `pulumi:"type"`
}


func NewOrchestratorInstanceServiceDetails(ctx *pulumi.Context,
	name string, args *OrchestratorInstanceServiceDetailsArgs, opts ...pulumi.ResourceOption) (*OrchestratorInstanceServiceDetails, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ControllerDetails == nil {
		return nil, errors.New("invalid value for required argument 'ControllerDetails'")
	}
	if args.Kind == nil {
		return nil, errors.New("invalid value for required argument 'Kind'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:delegatednetwork:OrchestratorInstanceServiceDetails"),
		},
		{
			Type: pulumi.String("azure-native:delegatednetwork/v20210315:OrchestratorInstanceServiceDetails"),
		},
	})
	opts = append(opts, aliases)
	var resource OrchestratorInstanceServiceDetails
	err := ctx.RegisterResource("azure-native:delegatednetwork/v20200808preview:OrchestratorInstanceServiceDetails", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetOrchestratorInstanceServiceDetails(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *OrchestratorInstanceServiceDetailsState, opts ...pulumi.ResourceOption) (*OrchestratorInstanceServiceDetails, error) {
	var resource OrchestratorInstanceServiceDetails
	err := ctx.ReadResource("azure-native:delegatednetwork/v20200808preview:OrchestratorInstanceServiceDetails", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type orchestratorInstanceServiceDetailsState struct {
}

type OrchestratorInstanceServiceDetailsState struct {
}

func (OrchestratorInstanceServiceDetailsState) ElementType() reflect.Type {
	return reflect.TypeOf((*orchestratorInstanceServiceDetailsState)(nil)).Elem()
}

type orchestratorInstanceServiceDetailsArgs struct {
	ApiServerEndpoint    *string               `pulumi:"apiServerEndpoint"`
	ClusterRootCA        *string               `pulumi:"clusterRootCA"`
	ControllerDetails    ControllerDetailsType `pulumi:"controllerDetails"`
	Identity             *OrchestratorIdentity `pulumi:"identity"`
	Kind                 string                `pulumi:"kind"`
	Location             *string               `pulumi:"location"`
	OrchestratorAppId    *string               `pulumi:"orchestratorAppId"`
	OrchestratorTenantId *string               `pulumi:"orchestratorTenantId"`
	ResourceGroupName    string                `pulumi:"resourceGroupName"`
	ResourceName         *string               `pulumi:"resourceName"`
	Tags                 map[string]string     `pulumi:"tags"`
}


type OrchestratorInstanceServiceDetailsArgs struct {
	ApiServerEndpoint    pulumi.StringPtrInput
	ClusterRootCA        pulumi.StringPtrInput
	ControllerDetails    ControllerDetailsTypeInput
	Identity             OrchestratorIdentityPtrInput
	Kind                 pulumi.StringInput
	Location             pulumi.StringPtrInput
	OrchestratorAppId    pulumi.StringPtrInput
	OrchestratorTenantId pulumi.StringPtrInput
	ResourceGroupName    pulumi.StringInput
	ResourceName         pulumi.StringPtrInput
	Tags                 pulumi.StringMapInput
}

func (OrchestratorInstanceServiceDetailsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*orchestratorInstanceServiceDetailsArgs)(nil)).Elem()
}

type OrchestratorInstanceServiceDetailsInput interface {
	pulumi.Input

	ToOrchestratorInstanceServiceDetailsOutput() OrchestratorInstanceServiceDetailsOutput
	ToOrchestratorInstanceServiceDetailsOutputWithContext(ctx context.Context) OrchestratorInstanceServiceDetailsOutput
}

func (*OrchestratorInstanceServiceDetails) ElementType() reflect.Type {
	return reflect.TypeOf((**OrchestratorInstanceServiceDetails)(nil)).Elem()
}

func (i *OrchestratorInstanceServiceDetails) ToOrchestratorInstanceServiceDetailsOutput() OrchestratorInstanceServiceDetailsOutput {
	return i.ToOrchestratorInstanceServiceDetailsOutputWithContext(context.Background())
}

func (i *OrchestratorInstanceServiceDetails) ToOrchestratorInstanceServiceDetailsOutputWithContext(ctx context.Context) OrchestratorInstanceServiceDetailsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(OrchestratorInstanceServiceDetailsOutput)
}

type OrchestratorInstanceServiceDetailsOutput struct{ *pulumi.OutputState }

func (OrchestratorInstanceServiceDetailsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**OrchestratorInstanceServiceDetails)(nil)).Elem()
}

func (o OrchestratorInstanceServiceDetailsOutput) ToOrchestratorInstanceServiceDetailsOutput() OrchestratorInstanceServiceDetailsOutput {
	return o
}

func (o OrchestratorInstanceServiceDetailsOutput) ToOrchestratorInstanceServiceDetailsOutputWithContext(ctx context.Context) OrchestratorInstanceServiceDetailsOutput {
	return o
}

func (o OrchestratorInstanceServiceDetailsOutput) ApiServerEndpoint() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *OrchestratorInstanceServiceDetails) pulumi.StringPtrOutput { return v.ApiServerEndpoint }).(pulumi.StringPtrOutput)
}

func (o OrchestratorInstanceServiceDetailsOutput) ClusterRootCA() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *OrchestratorInstanceServiceDetails) pulumi.StringPtrOutput { return v.ClusterRootCA }).(pulumi.StringPtrOutput)
}

func (o OrchestratorInstanceServiceDetailsOutput) ControllerDetails() ControllerDetailsResponseOutput {
	return o.ApplyT(func(v *OrchestratorInstanceServiceDetails) ControllerDetailsResponseOutput {
		return v.ControllerDetails
	}).(ControllerDetailsResponseOutput)
}

func (o OrchestratorInstanceServiceDetailsOutput) Identity() OrchestratorIdentityResponsePtrOutput {
	return o.ApplyT(func(v *OrchestratorInstanceServiceDetails) OrchestratorIdentityResponsePtrOutput { return v.Identity }).(OrchestratorIdentityResponsePtrOutput)
}

func (o OrchestratorInstanceServiceDetailsOutput) Kind() pulumi.StringOutput {
	return o.ApplyT(func(v *OrchestratorInstanceServiceDetails) pulumi.StringOutput { return v.Kind }).(pulumi.StringOutput)
}

func (o OrchestratorInstanceServiceDetailsOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *OrchestratorInstanceServiceDetails) pulumi.StringPtrOutput { return v.Location }).(pulumi.StringPtrOutput)
}

func (o OrchestratorInstanceServiceDetailsOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *OrchestratorInstanceServiceDetails) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o OrchestratorInstanceServiceDetailsOutput) OrchestratorAppId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *OrchestratorInstanceServiceDetails) pulumi.StringPtrOutput { return v.OrchestratorAppId }).(pulumi.StringPtrOutput)
}

func (o OrchestratorInstanceServiceDetailsOutput) OrchestratorTenantId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *OrchestratorInstanceServiceDetails) pulumi.StringPtrOutput { return v.OrchestratorTenantId }).(pulumi.StringPtrOutput)
}

func (o OrchestratorInstanceServiceDetailsOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *OrchestratorInstanceServiceDetails) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o OrchestratorInstanceServiceDetailsOutput) ResourceGuid() pulumi.StringOutput {
	return o.ApplyT(func(v *OrchestratorInstanceServiceDetails) pulumi.StringOutput { return v.ResourceGuid }).(pulumi.StringOutput)
}

func (o OrchestratorInstanceServiceDetailsOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *OrchestratorInstanceServiceDetails) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

func (o OrchestratorInstanceServiceDetailsOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *OrchestratorInstanceServiceDetails) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(OrchestratorInstanceServiceDetailsOutput{})
}
