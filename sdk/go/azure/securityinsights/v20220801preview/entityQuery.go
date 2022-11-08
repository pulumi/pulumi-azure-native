


package v20220801preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)


type EntityQuery struct {
	pulumi.CustomResourceState

	Etag       pulumi.StringPtrOutput   `pulumi:"etag"`
	Kind       pulumi.StringOutput      `pulumi:"kind"`
	Name       pulumi.StringOutput      `pulumi:"name"`
	SystemData SystemDataResponseOutput `pulumi:"systemData"`
	Type       pulumi.StringOutput      `pulumi:"type"`
}


func NewEntityQuery(ctx *pulumi.Context,
	name string, args *EntityQueryArgs, opts ...pulumi.ResourceOption) (*EntityQuery, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Kind == nil {
		return nil, errors.New("invalid value for required argument 'Kind'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.WorkspaceName == nil {
		return nil, errors.New("invalid value for required argument 'WorkspaceName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:securityinsights:EntityQuery"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20210301preview:EntityQuery"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20210901preview:EntityQuery"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20211001preview:EntityQuery"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20220101preview:EntityQuery"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20220401preview:EntityQuery"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20220501preview:EntityQuery"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20220601preview:EntityQuery"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20220701preview:EntityQuery"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20220901preview:EntityQuery"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20221001preview:EntityQuery"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20221101preview:EntityQuery"),
		},
	})
	opts = append(opts, aliases)
	var resource EntityQuery
	err := ctx.RegisterResource("azure-native:securityinsights/v20220801preview:EntityQuery", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetEntityQuery(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *EntityQueryState, opts ...pulumi.ResourceOption) (*EntityQuery, error) {
	var resource EntityQuery
	err := ctx.ReadResource("azure-native:securityinsights/v20220801preview:EntityQuery", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type entityQueryState struct {
}

type EntityQueryState struct {
}

func (EntityQueryState) ElementType() reflect.Type {
	return reflect.TypeOf((*entityQueryState)(nil)).Elem()
}

type entityQueryArgs struct {
	EntityQueryId     *string `pulumi:"entityQueryId"`
	Kind              string  `pulumi:"kind"`
	ResourceGroupName string  `pulumi:"resourceGroupName"`
	WorkspaceName     string  `pulumi:"workspaceName"`
}


type EntityQueryArgs struct {
	EntityQueryId     pulumi.StringPtrInput
	Kind              pulumi.StringInput
	ResourceGroupName pulumi.StringInput
	WorkspaceName     pulumi.StringInput
}

func (EntityQueryArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*entityQueryArgs)(nil)).Elem()
}

type EntityQueryInput interface {
	pulumi.Input

	ToEntityQueryOutput() EntityQueryOutput
	ToEntityQueryOutputWithContext(ctx context.Context) EntityQueryOutput
}

func (*EntityQuery) ElementType() reflect.Type {
	return reflect.TypeOf((**EntityQuery)(nil)).Elem()
}

func (i *EntityQuery) ToEntityQueryOutput() EntityQueryOutput {
	return i.ToEntityQueryOutputWithContext(context.Background())
}

func (i *EntityQuery) ToEntityQueryOutputWithContext(ctx context.Context) EntityQueryOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EntityQueryOutput)
}

type EntityQueryOutput struct{ *pulumi.OutputState }

func (EntityQueryOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**EntityQuery)(nil)).Elem()
}

func (o EntityQueryOutput) ToEntityQueryOutput() EntityQueryOutput {
	return o
}

func (o EntityQueryOutput) ToEntityQueryOutputWithContext(ctx context.Context) EntityQueryOutput {
	return o
}

func (o EntityQueryOutput) Etag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *EntityQuery) pulumi.StringPtrOutput { return v.Etag }).(pulumi.StringPtrOutput)
}

func (o EntityQueryOutput) Kind() pulumi.StringOutput {
	return o.ApplyT(func(v *EntityQuery) pulumi.StringOutput { return v.Kind }).(pulumi.StringOutput)
}

func (o EntityQueryOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *EntityQuery) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o EntityQueryOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *EntityQuery) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

func (o EntityQueryOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *EntityQuery) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(EntityQueryOutput{})
}
