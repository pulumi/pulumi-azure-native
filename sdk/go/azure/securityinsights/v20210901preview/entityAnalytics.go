


package v20210901preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type EntityAnalytics struct {
	pulumi.CustomResourceState

	Etag       pulumi.StringPtrOutput   `pulumi:"etag"`
	IsEnabled  pulumi.BoolOutput        `pulumi:"isEnabled"`
	Kind       pulumi.StringOutput      `pulumi:"kind"`
	Name       pulumi.StringOutput      `pulumi:"name"`
	SystemData SystemDataResponseOutput `pulumi:"systemData"`
	Type       pulumi.StringOutput      `pulumi:"type"`
}


func NewEntityAnalytics(ctx *pulumi.Context,
	name string, args *EntityAnalyticsArgs, opts ...pulumi.ResourceOption) (*EntityAnalytics, error) {
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
	args.Kind = pulumi.String("EntityAnalytics")
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:securityinsights:EntityAnalytics"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20190101preview:EntityAnalytics"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20210301preview:EntityAnalytics"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20211001preview:EntityAnalytics"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20220101preview:EntityAnalytics"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20220401preview:EntityAnalytics"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20220501preview:EntityAnalytics"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20220601preview:EntityAnalytics"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20220701preview:EntityAnalytics"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20220801preview:EntityAnalytics"),
		},
	})
	opts = append(opts, aliases)
	var resource EntityAnalytics
	err := ctx.RegisterResource("azure-native:securityinsights/v20210901preview:EntityAnalytics", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetEntityAnalytics(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *EntityAnalyticsState, opts ...pulumi.ResourceOption) (*EntityAnalytics, error) {
	var resource EntityAnalytics
	err := ctx.ReadResource("azure-native:securityinsights/v20210901preview:EntityAnalytics", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type entityAnalyticsState struct {
}

type EntityAnalyticsState struct {
}

func (EntityAnalyticsState) ElementType() reflect.Type {
	return reflect.TypeOf((*entityAnalyticsState)(nil)).Elem()
}

type entityAnalyticsArgs struct {
	Kind              string  `pulumi:"kind"`
	ResourceGroupName string  `pulumi:"resourceGroupName"`
	SettingsName      *string `pulumi:"settingsName"`
	WorkspaceName     string  `pulumi:"workspaceName"`
}


type EntityAnalyticsArgs struct {
	Kind              pulumi.StringInput
	ResourceGroupName pulumi.StringInput
	SettingsName      pulumi.StringPtrInput
	WorkspaceName     pulumi.StringInput
}

func (EntityAnalyticsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*entityAnalyticsArgs)(nil)).Elem()
}

type EntityAnalyticsInput interface {
	pulumi.Input

	ToEntityAnalyticsOutput() EntityAnalyticsOutput
	ToEntityAnalyticsOutputWithContext(ctx context.Context) EntityAnalyticsOutput
}

func (*EntityAnalytics) ElementType() reflect.Type {
	return reflect.TypeOf((**EntityAnalytics)(nil)).Elem()
}

func (i *EntityAnalytics) ToEntityAnalyticsOutput() EntityAnalyticsOutput {
	return i.ToEntityAnalyticsOutputWithContext(context.Background())
}

func (i *EntityAnalytics) ToEntityAnalyticsOutputWithContext(ctx context.Context) EntityAnalyticsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EntityAnalyticsOutput)
}

type EntityAnalyticsOutput struct{ *pulumi.OutputState }

func (EntityAnalyticsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**EntityAnalytics)(nil)).Elem()
}

func (o EntityAnalyticsOutput) ToEntityAnalyticsOutput() EntityAnalyticsOutput {
	return o
}

func (o EntityAnalyticsOutput) ToEntityAnalyticsOutputWithContext(ctx context.Context) EntityAnalyticsOutput {
	return o
}

func (o EntityAnalyticsOutput) Etag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *EntityAnalytics) pulumi.StringPtrOutput { return v.Etag }).(pulumi.StringPtrOutput)
}

func (o EntityAnalyticsOutput) IsEnabled() pulumi.BoolOutput {
	return o.ApplyT(func(v *EntityAnalytics) pulumi.BoolOutput { return v.IsEnabled }).(pulumi.BoolOutput)
}

func (o EntityAnalyticsOutput) Kind() pulumi.StringOutput {
	return o.ApplyT(func(v *EntityAnalytics) pulumi.StringOutput { return v.Kind }).(pulumi.StringOutput)
}

func (o EntityAnalyticsOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *EntityAnalytics) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o EntityAnalyticsOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *EntityAnalytics) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

func (o EntityAnalyticsOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *EntityAnalytics) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(EntityAnalyticsOutput{})
}
