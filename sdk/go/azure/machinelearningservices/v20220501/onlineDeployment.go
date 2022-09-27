


package v20220501

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type OnlineDeployment struct {
	pulumi.CustomResourceState

	Identity                   ManagedServiceIdentityResponsePtrOutput `pulumi:"identity"`
	Kind                       pulumi.StringPtrOutput                  `pulumi:"kind"`
	Location                   pulumi.StringOutput                     `pulumi:"location"`
	Name                       pulumi.StringOutput                     `pulumi:"name"`
	OnlineDeploymentProperties pulumi.AnyOutput                        `pulumi:"onlineDeploymentProperties"`
	Sku                        SkuResponsePtrOutput                    `pulumi:"sku"`
	SystemData                 SystemDataResponseOutput                `pulumi:"systemData"`
	Tags                       pulumi.StringMapOutput                  `pulumi:"tags"`
	Type                       pulumi.StringOutput                     `pulumi:"type"`
}


func NewOnlineDeployment(ctx *pulumi.Context,
	name string, args *OnlineDeploymentArgs, opts ...pulumi.ResourceOption) (*OnlineDeployment, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.EndpointName == nil {
		return nil, errors.New("invalid value for required argument 'EndpointName'")
	}
	if args.OnlineDeploymentProperties == nil {
		return nil, errors.New("invalid value for required argument 'OnlineDeploymentProperties'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.WorkspaceName == nil {
		return nil, errors.New("invalid value for required argument 'WorkspaceName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:machinelearningservices:OnlineDeployment"),
		},
		{
			Type: pulumi.String("azure-native:machinelearningservices/v20210301preview:OnlineDeployment"),
		},
		{
			Type: pulumi.String("azure-native:machinelearningservices/v20220201preview:OnlineDeployment"),
		},
		{
			Type: pulumi.String("azure-native:machinelearningservices/v20220601preview:OnlineDeployment"),
		},
	})
	opts = append(opts, aliases)
	var resource OnlineDeployment
	err := ctx.RegisterResource("azure-native:machinelearningservices/v20220501:OnlineDeployment", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetOnlineDeployment(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *OnlineDeploymentState, opts ...pulumi.ResourceOption) (*OnlineDeployment, error) {
	var resource OnlineDeployment
	err := ctx.ReadResource("azure-native:machinelearningservices/v20220501:OnlineDeployment", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type onlineDeploymentState struct {
}

type OnlineDeploymentState struct {
}

func (OnlineDeploymentState) ElementType() reflect.Type {
	return reflect.TypeOf((*onlineDeploymentState)(nil)).Elem()
}

type onlineDeploymentArgs struct {
	DeploymentName             *string                 `pulumi:"deploymentName"`
	EndpointName               string                  `pulumi:"endpointName"`
	Identity                   *ManagedServiceIdentity `pulumi:"identity"`
	Kind                       *string                 `pulumi:"kind"`
	Location                   *string                 `pulumi:"location"`
	OnlineDeploymentProperties interface{}             `pulumi:"onlineDeploymentProperties"`
	ResourceGroupName          string                  `pulumi:"resourceGroupName"`
	Sku                        *Sku                    `pulumi:"sku"`
	Tags                       map[string]string       `pulumi:"tags"`
	WorkspaceName              string                  `pulumi:"workspaceName"`
}


type OnlineDeploymentArgs struct {
	DeploymentName             pulumi.StringPtrInput
	EndpointName               pulumi.StringInput
	Identity                   ManagedServiceIdentityPtrInput
	Kind                       pulumi.StringPtrInput
	Location                   pulumi.StringPtrInput
	OnlineDeploymentProperties pulumi.Input
	ResourceGroupName          pulumi.StringInput
	Sku                        SkuPtrInput
	Tags                       pulumi.StringMapInput
	WorkspaceName              pulumi.StringInput
}

func (OnlineDeploymentArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*onlineDeploymentArgs)(nil)).Elem()
}

type OnlineDeploymentInput interface {
	pulumi.Input

	ToOnlineDeploymentOutput() OnlineDeploymentOutput
	ToOnlineDeploymentOutputWithContext(ctx context.Context) OnlineDeploymentOutput
}

func (*OnlineDeployment) ElementType() reflect.Type {
	return reflect.TypeOf((**OnlineDeployment)(nil)).Elem()
}

func (i *OnlineDeployment) ToOnlineDeploymentOutput() OnlineDeploymentOutput {
	return i.ToOnlineDeploymentOutputWithContext(context.Background())
}

func (i *OnlineDeployment) ToOnlineDeploymentOutputWithContext(ctx context.Context) OnlineDeploymentOutput {
	return pulumi.ToOutputWithContext(ctx, i).(OnlineDeploymentOutput)
}

type OnlineDeploymentOutput struct{ *pulumi.OutputState }

func (OnlineDeploymentOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**OnlineDeployment)(nil)).Elem()
}

func (o OnlineDeploymentOutput) ToOnlineDeploymentOutput() OnlineDeploymentOutput {
	return o
}

func (o OnlineDeploymentOutput) ToOnlineDeploymentOutputWithContext(ctx context.Context) OnlineDeploymentOutput {
	return o
}

func (o OnlineDeploymentOutput) Identity() ManagedServiceIdentityResponsePtrOutput {
	return o.ApplyT(func(v *OnlineDeployment) ManagedServiceIdentityResponsePtrOutput { return v.Identity }).(ManagedServiceIdentityResponsePtrOutput)
}

func (o OnlineDeploymentOutput) Kind() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *OnlineDeployment) pulumi.StringPtrOutput { return v.Kind }).(pulumi.StringPtrOutput)
}

func (o OnlineDeploymentOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *OnlineDeployment) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

func (o OnlineDeploymentOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *OnlineDeployment) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o OnlineDeploymentOutput) OnlineDeploymentProperties() pulumi.AnyOutput {
	return o.ApplyT(func(v *OnlineDeployment) pulumi.AnyOutput { return v.OnlineDeploymentProperties }).(pulumi.AnyOutput)
}

func (o OnlineDeploymentOutput) Sku() SkuResponsePtrOutput {
	return o.ApplyT(func(v *OnlineDeployment) SkuResponsePtrOutput { return v.Sku }).(SkuResponsePtrOutput)
}

func (o OnlineDeploymentOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *OnlineDeployment) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

func (o OnlineDeploymentOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *OnlineDeployment) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

func (o OnlineDeploymentOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *OnlineDeployment) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(OnlineDeploymentOutput{})
}
