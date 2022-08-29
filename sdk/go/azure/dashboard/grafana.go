


package dashboard

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type Grafana struct {
	pulumi.CustomResourceState

	Identity   ManagedServiceIdentityResponsePtrOutput `pulumi:"identity"`
	Location   pulumi.StringPtrOutput                  `pulumi:"location"`
	Name       pulumi.StringOutput                     `pulumi:"name"`
	Properties ManagedGrafanaPropertiesResponseOutput  `pulumi:"properties"`
	Sku        ResourceSkuResponsePtrOutput            `pulumi:"sku"`
	SystemData SystemDataResponseOutput                `pulumi:"systemData"`
	Tags       pulumi.StringMapOutput                  `pulumi:"tags"`
	Type       pulumi.StringOutput                     `pulumi:"type"`
}


func NewGrafana(ctx *pulumi.Context,
	name string, args *GrafanaArgs, opts ...pulumi.ResourceOption) (*Grafana, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:dashboard/v20210901preview:Grafana"),
		},
		{
			Type: pulumi.String("azure-native:dashboard/v20220501preview:Grafana"),
		},
		{
			Type: pulumi.String("azure-native:dashboard/v20220801:Grafana"),
		},
	})
	opts = append(opts, aliases)
	var resource Grafana
	err := ctx.RegisterResource("azure-native:dashboard:Grafana", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetGrafana(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *GrafanaState, opts ...pulumi.ResourceOption) (*Grafana, error) {
	var resource Grafana
	err := ctx.ReadResource("azure-native:dashboard:Grafana", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type grafanaState struct {
}

type GrafanaState struct {
}

func (GrafanaState) ElementType() reflect.Type {
	return reflect.TypeOf((*grafanaState)(nil)).Elem()
}

type grafanaArgs struct {
	Identity          *ManagedServiceIdentity   `pulumi:"identity"`
	Location          *string                   `pulumi:"location"`
	Properties        *ManagedGrafanaProperties `pulumi:"properties"`
	ResourceGroupName string                    `pulumi:"resourceGroupName"`
	Sku               *ResourceSku              `pulumi:"sku"`
	Tags              map[string]string         `pulumi:"tags"`
	WorkspaceName     *string                   `pulumi:"workspaceName"`
}


type GrafanaArgs struct {
	Identity          ManagedServiceIdentityPtrInput
	Location          pulumi.StringPtrInput
	Properties        ManagedGrafanaPropertiesPtrInput
	ResourceGroupName pulumi.StringInput
	Sku               ResourceSkuPtrInput
	Tags              pulumi.StringMapInput
	WorkspaceName     pulumi.StringPtrInput
}

func (GrafanaArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*grafanaArgs)(nil)).Elem()
}

type GrafanaInput interface {
	pulumi.Input

	ToGrafanaOutput() GrafanaOutput
	ToGrafanaOutputWithContext(ctx context.Context) GrafanaOutput
}

func (*Grafana) ElementType() reflect.Type {
	return reflect.TypeOf((**Grafana)(nil)).Elem()
}

func (i *Grafana) ToGrafanaOutput() GrafanaOutput {
	return i.ToGrafanaOutputWithContext(context.Background())
}

func (i *Grafana) ToGrafanaOutputWithContext(ctx context.Context) GrafanaOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GrafanaOutput)
}

type GrafanaOutput struct{ *pulumi.OutputState }

func (GrafanaOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Grafana)(nil)).Elem()
}

func (o GrafanaOutput) ToGrafanaOutput() GrafanaOutput {
	return o
}

func (o GrafanaOutput) ToGrafanaOutputWithContext(ctx context.Context) GrafanaOutput {
	return o
}

func (o GrafanaOutput) Identity() ManagedServiceIdentityResponsePtrOutput {
	return o.ApplyT(func(v *Grafana) ManagedServiceIdentityResponsePtrOutput { return v.Identity }).(ManagedServiceIdentityResponsePtrOutput)
}

func (o GrafanaOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Grafana) pulumi.StringPtrOutput { return v.Location }).(pulumi.StringPtrOutput)
}

func (o GrafanaOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Grafana) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o GrafanaOutput) Properties() ManagedGrafanaPropertiesResponseOutput {
	return o.ApplyT(func(v *Grafana) ManagedGrafanaPropertiesResponseOutput { return v.Properties }).(ManagedGrafanaPropertiesResponseOutput)
}

func (o GrafanaOutput) Sku() ResourceSkuResponsePtrOutput {
	return o.ApplyT(func(v *Grafana) ResourceSkuResponsePtrOutput { return v.Sku }).(ResourceSkuResponsePtrOutput)
}

func (o GrafanaOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *Grafana) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

func (o GrafanaOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *Grafana) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

func (o GrafanaOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *Grafana) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(GrafanaOutput{})
}
