


package v20220501

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type OnlineEndpoint struct {
	pulumi.CustomResourceState

	Identity                 ManagedServiceIdentityResponsePtrOutput `pulumi:"identity"`
	Kind                     pulumi.StringPtrOutput                  `pulumi:"kind"`
	Location                 pulumi.StringOutput                     `pulumi:"location"`
	Name                     pulumi.StringOutput                     `pulumi:"name"`
	OnlineEndpointProperties OnlineEndpointResponseOutput            `pulumi:"onlineEndpointProperties"`
	Sku                      SkuResponsePtrOutput                    `pulumi:"sku"`
	SystemData               SystemDataResponseOutput                `pulumi:"systemData"`
	Tags                     pulumi.StringMapOutput                  `pulumi:"tags"`
	Type                     pulumi.StringOutput                     `pulumi:"type"`
}


func NewOnlineEndpoint(ctx *pulumi.Context,
	name string, args *OnlineEndpointArgs, opts ...pulumi.ResourceOption) (*OnlineEndpoint, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.OnlineEndpointProperties == nil {
		return nil, errors.New("invalid value for required argument 'OnlineEndpointProperties'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.WorkspaceName == nil {
		return nil, errors.New("invalid value for required argument 'WorkspaceName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:machinelearningservices:OnlineEndpoint"),
		},
		{
			Type: pulumi.String("azure-native:machinelearningservices/v20210301preview:OnlineEndpoint"),
		},
		{
			Type: pulumi.String("azure-native:machinelearningservices/v20220201preview:OnlineEndpoint"),
		},
		{
			Type: pulumi.String("azure-native:machinelearningservices/v20220601preview:OnlineEndpoint"),
		},
		{
			Type: pulumi.String("azure-native:machinelearningservices/v20221001:OnlineEndpoint"),
		},
		{
			Type: pulumi.String("azure-native:machinelearningservices/v20221001preview:OnlineEndpoint"),
		},
	})
	opts = append(opts, aliases)
	var resource OnlineEndpoint
	err := ctx.RegisterResource("azure-native:machinelearningservices/v20220501:OnlineEndpoint", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetOnlineEndpoint(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *OnlineEndpointState, opts ...pulumi.ResourceOption) (*OnlineEndpoint, error) {
	var resource OnlineEndpoint
	err := ctx.ReadResource("azure-native:machinelearningservices/v20220501:OnlineEndpoint", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type onlineEndpointState struct {
}

type OnlineEndpointState struct {
}

func (OnlineEndpointState) ElementType() reflect.Type {
	return reflect.TypeOf((*onlineEndpointState)(nil)).Elem()
}

type onlineEndpointArgs struct {
	EndpointName             *string                 `pulumi:"endpointName"`
	Identity                 *ManagedServiceIdentity `pulumi:"identity"`
	Kind                     *string                 `pulumi:"kind"`
	Location                 *string                 `pulumi:"location"`
	OnlineEndpointProperties OnlineEndpointType      `pulumi:"onlineEndpointProperties"`
	ResourceGroupName        string                  `pulumi:"resourceGroupName"`
	Sku                      *Sku                    `pulumi:"sku"`
	Tags                     map[string]string       `pulumi:"tags"`
	WorkspaceName            string                  `pulumi:"workspaceName"`
}


type OnlineEndpointArgs struct {
	EndpointName             pulumi.StringPtrInput
	Identity                 ManagedServiceIdentityPtrInput
	Kind                     pulumi.StringPtrInput
	Location                 pulumi.StringPtrInput
	OnlineEndpointProperties OnlineEndpointTypeInput
	ResourceGroupName        pulumi.StringInput
	Sku                      SkuPtrInput
	Tags                     pulumi.StringMapInput
	WorkspaceName            pulumi.StringInput
}

func (OnlineEndpointArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*onlineEndpointArgs)(nil)).Elem()
}

type OnlineEndpointInput interface {
	pulumi.Input

	ToOnlineEndpointOutput() OnlineEndpointOutput
	ToOnlineEndpointOutputWithContext(ctx context.Context) OnlineEndpointOutput
}

func (*OnlineEndpoint) ElementType() reflect.Type {
	return reflect.TypeOf((**OnlineEndpoint)(nil)).Elem()
}

func (i *OnlineEndpoint) ToOnlineEndpointOutput() OnlineEndpointOutput {
	return i.ToOnlineEndpointOutputWithContext(context.Background())
}

func (i *OnlineEndpoint) ToOnlineEndpointOutputWithContext(ctx context.Context) OnlineEndpointOutput {
	return pulumi.ToOutputWithContext(ctx, i).(OnlineEndpointOutput)
}

type OnlineEndpointOutput struct{ *pulumi.OutputState }

func (OnlineEndpointOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**OnlineEndpoint)(nil)).Elem()
}

func (o OnlineEndpointOutput) ToOnlineEndpointOutput() OnlineEndpointOutput {
	return o
}

func (o OnlineEndpointOutput) ToOnlineEndpointOutputWithContext(ctx context.Context) OnlineEndpointOutput {
	return o
}

func (o OnlineEndpointOutput) Identity() ManagedServiceIdentityResponsePtrOutput {
	return o.ApplyT(func(v *OnlineEndpoint) ManagedServiceIdentityResponsePtrOutput { return v.Identity }).(ManagedServiceIdentityResponsePtrOutput)
}

func (o OnlineEndpointOutput) Kind() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *OnlineEndpoint) pulumi.StringPtrOutput { return v.Kind }).(pulumi.StringPtrOutput)
}

func (o OnlineEndpointOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *OnlineEndpoint) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

func (o OnlineEndpointOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *OnlineEndpoint) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o OnlineEndpointOutput) OnlineEndpointProperties() OnlineEndpointResponseOutput {
	return o.ApplyT(func(v *OnlineEndpoint) OnlineEndpointResponseOutput { return v.OnlineEndpointProperties }).(OnlineEndpointResponseOutput)
}

func (o OnlineEndpointOutput) Sku() SkuResponsePtrOutput {
	return o.ApplyT(func(v *OnlineEndpoint) SkuResponsePtrOutput { return v.Sku }).(SkuResponsePtrOutput)
}

func (o OnlineEndpointOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *OnlineEndpoint) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

func (o OnlineEndpointOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *OnlineEndpoint) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

func (o OnlineEndpointOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *OnlineEndpoint) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(OnlineEndpointOutput{})
}
