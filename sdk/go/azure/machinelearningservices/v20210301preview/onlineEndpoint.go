


package v20210301preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type OnlineEndpoint struct {
	pulumi.CustomResourceState

	Identity   ResourceIdentityResponsePtrOutput `pulumi:"identity"`
	Kind       pulumi.StringPtrOutput            `pulumi:"kind"`
	Location   pulumi.StringOutput               `pulumi:"location"`
	Name       pulumi.StringOutput               `pulumi:"name"`
	Properties OnlineEndpointResponseOutput      `pulumi:"properties"`
	SystemData SystemDataResponseOutput          `pulumi:"systemData"`
	Tags       pulumi.StringMapOutput            `pulumi:"tags"`
	Type       pulumi.StringOutput               `pulumi:"type"`
}


func NewOnlineEndpoint(ctx *pulumi.Context,
	name string, args *OnlineEndpointArgs, opts ...pulumi.ResourceOption) (*OnlineEndpoint, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Properties == nil {
		return nil, errors.New("invalid value for required argument 'Properties'")
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
			Type: pulumi.String("azure-native:machinelearningservices/v20220201preview:OnlineEndpoint"),
		},
		{
			Type: pulumi.String("azure-native:machinelearningservices/v20220501:OnlineEndpoint"),
		},
		{
			Type: pulumi.String("azure-native:machinelearningservices/v20220601preview:OnlineEndpoint"),
		},
		{
			Type: pulumi.String("azure-native:machinelearningservices/v20221001:OnlineEndpoint"),
		},
	})
	opts = append(opts, aliases)
	var resource OnlineEndpoint
	err := ctx.RegisterResource("azure-native:machinelearningservices/v20210301preview:OnlineEndpoint", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetOnlineEndpoint(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *OnlineEndpointState, opts ...pulumi.ResourceOption) (*OnlineEndpoint, error) {
	var resource OnlineEndpoint
	err := ctx.ReadResource("azure-native:machinelearningservices/v20210301preview:OnlineEndpoint", name, id, state, &resource, opts...)
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
	EndpointName      *string            `pulumi:"endpointName"`
	Identity          *ResourceIdentity  `pulumi:"identity"`
	Kind              *string            `pulumi:"kind"`
	Location          *string            `pulumi:"location"`
	Properties        OnlineEndpointType `pulumi:"properties"`
	ResourceGroupName string             `pulumi:"resourceGroupName"`
	Tags              map[string]string  `pulumi:"tags"`
	WorkspaceName     string             `pulumi:"workspaceName"`
}


type OnlineEndpointArgs struct {
	EndpointName      pulumi.StringPtrInput
	Identity          ResourceIdentityPtrInput
	Kind              pulumi.StringPtrInput
	Location          pulumi.StringPtrInput
	Properties        OnlineEndpointTypeInput
	ResourceGroupName pulumi.StringInput
	Tags              pulumi.StringMapInput
	WorkspaceName     pulumi.StringInput
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

func (o OnlineEndpointOutput) Identity() ResourceIdentityResponsePtrOutput {
	return o.ApplyT(func(v *OnlineEndpoint) ResourceIdentityResponsePtrOutput { return v.Identity }).(ResourceIdentityResponsePtrOutput)
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

func (o OnlineEndpointOutput) Properties() OnlineEndpointResponseOutput {
	return o.ApplyT(func(v *OnlineEndpoint) OnlineEndpointResponseOutput { return v.Properties }).(OnlineEndpointResponseOutput)
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
