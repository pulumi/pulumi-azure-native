


package v20211101

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type HybridConnection struct {
	pulumi.CustomResourceState

	CreatedAt                   pulumi.StringOutput      `pulumi:"createdAt"`
	ListenerCount               pulumi.IntOutput         `pulumi:"listenerCount"`
	Location                    pulumi.StringOutput      `pulumi:"location"`
	Name                        pulumi.StringOutput      `pulumi:"name"`
	RequiresClientAuthorization pulumi.BoolPtrOutput     `pulumi:"requiresClientAuthorization"`
	SystemData                  SystemDataResponseOutput `pulumi:"systemData"`
	Type                        pulumi.StringOutput      `pulumi:"type"`
	UpdatedAt                   pulumi.StringOutput      `pulumi:"updatedAt"`
	UserMetadata                pulumi.StringPtrOutput   `pulumi:"userMetadata"`
}


func NewHybridConnection(ctx *pulumi.Context,
	name string, args *HybridConnectionArgs, opts ...pulumi.ResourceOption) (*HybridConnection, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.NamespaceName == nil {
		return nil, errors.New("invalid value for required argument 'NamespaceName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:relay:HybridConnection"),
		},
		{
			Type: pulumi.String("azure-native:relay/v20160701:HybridConnection"),
		},
		{
			Type: pulumi.String("azure-native:relay/v20170401:HybridConnection"),
		},
	})
	opts = append(opts, aliases)
	var resource HybridConnection
	err := ctx.RegisterResource("azure-native:relay/v20211101:HybridConnection", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetHybridConnection(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *HybridConnectionState, opts ...pulumi.ResourceOption) (*HybridConnection, error) {
	var resource HybridConnection
	err := ctx.ReadResource("azure-native:relay/v20211101:HybridConnection", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type hybridConnectionState struct {
}

type HybridConnectionState struct {
}

func (HybridConnectionState) ElementType() reflect.Type {
	return reflect.TypeOf((*hybridConnectionState)(nil)).Elem()
}

type hybridConnectionArgs struct {
	HybridConnectionName        *string `pulumi:"hybridConnectionName"`
	NamespaceName               string  `pulumi:"namespaceName"`
	RequiresClientAuthorization *bool   `pulumi:"requiresClientAuthorization"`
	ResourceGroupName           string  `pulumi:"resourceGroupName"`
	UserMetadata                *string `pulumi:"userMetadata"`
}


type HybridConnectionArgs struct {
	HybridConnectionName        pulumi.StringPtrInput
	NamespaceName               pulumi.StringInput
	RequiresClientAuthorization pulumi.BoolPtrInput
	ResourceGroupName           pulumi.StringInput
	UserMetadata                pulumi.StringPtrInput
}

func (HybridConnectionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*hybridConnectionArgs)(nil)).Elem()
}

type HybridConnectionInput interface {
	pulumi.Input

	ToHybridConnectionOutput() HybridConnectionOutput
	ToHybridConnectionOutputWithContext(ctx context.Context) HybridConnectionOutput
}

func (*HybridConnection) ElementType() reflect.Type {
	return reflect.TypeOf((**HybridConnection)(nil)).Elem()
}

func (i *HybridConnection) ToHybridConnectionOutput() HybridConnectionOutput {
	return i.ToHybridConnectionOutputWithContext(context.Background())
}

func (i *HybridConnection) ToHybridConnectionOutputWithContext(ctx context.Context) HybridConnectionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(HybridConnectionOutput)
}

type HybridConnectionOutput struct{ *pulumi.OutputState }

func (HybridConnectionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**HybridConnection)(nil)).Elem()
}

func (o HybridConnectionOutput) ToHybridConnectionOutput() HybridConnectionOutput {
	return o
}

func (o HybridConnectionOutput) ToHybridConnectionOutputWithContext(ctx context.Context) HybridConnectionOutput {
	return o
}

func (o HybridConnectionOutput) CreatedAt() pulumi.StringOutput {
	return o.ApplyT(func(v *HybridConnection) pulumi.StringOutput { return v.CreatedAt }).(pulumi.StringOutput)
}

func (o HybridConnectionOutput) ListenerCount() pulumi.IntOutput {
	return o.ApplyT(func(v *HybridConnection) pulumi.IntOutput { return v.ListenerCount }).(pulumi.IntOutput)
}

func (o HybridConnectionOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *HybridConnection) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

func (o HybridConnectionOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *HybridConnection) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o HybridConnectionOutput) RequiresClientAuthorization() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *HybridConnection) pulumi.BoolPtrOutput { return v.RequiresClientAuthorization }).(pulumi.BoolPtrOutput)
}

func (o HybridConnectionOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *HybridConnection) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

func (o HybridConnectionOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *HybridConnection) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func (o HybridConnectionOutput) UpdatedAt() pulumi.StringOutput {
	return o.ApplyT(func(v *HybridConnection) pulumi.StringOutput { return v.UpdatedAt }).(pulumi.StringOutput)
}

func (o HybridConnectionOutput) UserMetadata() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *HybridConnection) pulumi.StringPtrOutput { return v.UserMetadata }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(HybridConnectionOutput{})
}
