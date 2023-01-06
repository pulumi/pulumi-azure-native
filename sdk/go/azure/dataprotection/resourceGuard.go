


package dataprotection

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type ResourceGuard struct {
	pulumi.CustomResourceState

	ETag       pulumi.StringPtrOutput              `pulumi:"eTag"`
	Identity   DppIdentityDetailsResponsePtrOutput `pulumi:"identity"`
	Location   pulumi.StringPtrOutput              `pulumi:"location"`
	Name       pulumi.StringOutput                 `pulumi:"name"`
	Properties ResourceGuardResponseOutput         `pulumi:"properties"`
	SystemData SystemDataResponseOutput            `pulumi:"systemData"`
	Tags       pulumi.StringMapOutput              `pulumi:"tags"`
	Type       pulumi.StringOutput                 `pulumi:"type"`
}


func NewResourceGuard(ctx *pulumi.Context,
	name string, args *ResourceGuardArgs, opts ...pulumi.ResourceOption) (*ResourceGuard, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:dataprotection/v20210701:ResourceGuard"),
		},
		{
			Type: pulumi.String("azure-native:dataprotection/v20211001preview:ResourceGuard"),
		},
		{
			Type: pulumi.String("azure-native:dataprotection/v20211201preview:ResourceGuard"),
		},
		{
			Type: pulumi.String("azure-native:dataprotection/v20220101:ResourceGuard"),
		},
		{
			Type: pulumi.String("azure-native:dataprotection/v20220201preview:ResourceGuard"),
		},
		{
			Type: pulumi.String("azure-native:dataprotection/v20220301:ResourceGuard"),
		},
		{
			Type: pulumi.String("azure-native:dataprotection/v20220331preview:ResourceGuard"),
		},
		{
			Type: pulumi.String("azure-native:dataprotection/v20220401:ResourceGuard"),
		},
		{
			Type: pulumi.String("azure-native:dataprotection/v20220501:ResourceGuard"),
		},
		{
			Type: pulumi.String("azure-native:dataprotection/v20220901preview:ResourceGuard"),
		},
		{
			Type: pulumi.String("azure-native:dataprotection/v20221001preview:ResourceGuard"),
		},
		{
			Type: pulumi.String("azure-native:dataprotection/v20221101preview:ResourceGuard"),
		},
		{
			Type: pulumi.String("azure-native:dataprotection/v20221201:ResourceGuard"),
		},
	})
	opts = append(opts, aliases)
	var resource ResourceGuard
	err := ctx.RegisterResource("azure-native:dataprotection:ResourceGuard", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetResourceGuard(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ResourceGuardState, opts ...pulumi.ResourceOption) (*ResourceGuard, error) {
	var resource ResourceGuard
	err := ctx.ReadResource("azure-native:dataprotection:ResourceGuard", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type resourceGuardState struct {
}

type ResourceGuardState struct {
}

func (ResourceGuardState) ElementType() reflect.Type {
	return reflect.TypeOf((*resourceGuardState)(nil)).Elem()
}

type resourceGuardArgs struct {
	ETag               *string             `pulumi:"eTag"`
	Identity           *DppIdentityDetails `pulumi:"identity"`
	Location           *string             `pulumi:"location"`
	ResourceGroupName  string              `pulumi:"resourceGroupName"`
	ResourceGuardsName *string             `pulumi:"resourceGuardsName"`
	Tags               map[string]string   `pulumi:"tags"`
}


type ResourceGuardArgs struct {
	ETag               pulumi.StringPtrInput
	Identity           DppIdentityDetailsPtrInput
	Location           pulumi.StringPtrInput
	ResourceGroupName  pulumi.StringInput
	ResourceGuardsName pulumi.StringPtrInput
	Tags               pulumi.StringMapInput
}

func (ResourceGuardArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*resourceGuardArgs)(nil)).Elem()
}

type ResourceGuardInput interface {
	pulumi.Input

	ToResourceGuardOutput() ResourceGuardOutput
	ToResourceGuardOutputWithContext(ctx context.Context) ResourceGuardOutput
}

func (*ResourceGuard) ElementType() reflect.Type {
	return reflect.TypeOf((**ResourceGuard)(nil)).Elem()
}

func (i *ResourceGuard) ToResourceGuardOutput() ResourceGuardOutput {
	return i.ToResourceGuardOutputWithContext(context.Background())
}

func (i *ResourceGuard) ToResourceGuardOutputWithContext(ctx context.Context) ResourceGuardOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ResourceGuardOutput)
}

type ResourceGuardOutput struct{ *pulumi.OutputState }

func (ResourceGuardOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ResourceGuard)(nil)).Elem()
}

func (o ResourceGuardOutput) ToResourceGuardOutput() ResourceGuardOutput {
	return o
}

func (o ResourceGuardOutput) ToResourceGuardOutputWithContext(ctx context.Context) ResourceGuardOutput {
	return o
}

func (o ResourceGuardOutput) ETag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ResourceGuard) pulumi.StringPtrOutput { return v.ETag }).(pulumi.StringPtrOutput)
}

func (o ResourceGuardOutput) Identity() DppIdentityDetailsResponsePtrOutput {
	return o.ApplyT(func(v *ResourceGuard) DppIdentityDetailsResponsePtrOutput { return v.Identity }).(DppIdentityDetailsResponsePtrOutput)
}

func (o ResourceGuardOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ResourceGuard) pulumi.StringPtrOutput { return v.Location }).(pulumi.StringPtrOutput)
}

func (o ResourceGuardOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *ResourceGuard) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o ResourceGuardOutput) Properties() ResourceGuardResponseOutput {
	return o.ApplyT(func(v *ResourceGuard) ResourceGuardResponseOutput { return v.Properties }).(ResourceGuardResponseOutput)
}

func (o ResourceGuardOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *ResourceGuard) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

func (o ResourceGuardOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *ResourceGuard) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

func (o ResourceGuardOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *ResourceGuard) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(ResourceGuardOutput{})
}
