


package v20220101preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type Anomalies struct {
	pulumi.CustomResourceState

	Etag       pulumi.StringPtrOutput   `pulumi:"etag"`
	IsEnabled  pulumi.BoolOutput        `pulumi:"isEnabled"`
	Kind       pulumi.StringOutput      `pulumi:"kind"`
	Name       pulumi.StringOutput      `pulumi:"name"`
	SystemData SystemDataResponseOutput `pulumi:"systemData"`
	Type       pulumi.StringOutput      `pulumi:"type"`
}


func NewAnomalies(ctx *pulumi.Context,
	name string, args *AnomaliesArgs, opts ...pulumi.ResourceOption) (*Anomalies, error) {
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
	args.Kind = pulumi.String("Anomalies")
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:securityinsights:Anomalies"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20190101preview:Anomalies"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20210301preview:Anomalies"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20210901preview:Anomalies"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20211001preview:Anomalies"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20220401preview:Anomalies"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20220501preview:Anomalies"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20220601preview:Anomalies"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20220701preview:Anomalies"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20220801preview:Anomalies"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20220901preview:Anomalies"),
		},
	})
	opts = append(opts, aliases)
	var resource Anomalies
	err := ctx.RegisterResource("azure-native:securityinsights/v20220101preview:Anomalies", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetAnomalies(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AnomaliesState, opts ...pulumi.ResourceOption) (*Anomalies, error) {
	var resource Anomalies
	err := ctx.ReadResource("azure-native:securityinsights/v20220101preview:Anomalies", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type anomaliesState struct {
}

type AnomaliesState struct {
}

func (AnomaliesState) ElementType() reflect.Type {
	return reflect.TypeOf((*anomaliesState)(nil)).Elem()
}

type anomaliesArgs struct {
	Kind              string  `pulumi:"kind"`
	ResourceGroupName string  `pulumi:"resourceGroupName"`
	SettingsName      *string `pulumi:"settingsName"`
	WorkspaceName     string  `pulumi:"workspaceName"`
}


type AnomaliesArgs struct {
	Kind              pulumi.StringInput
	ResourceGroupName pulumi.StringInput
	SettingsName      pulumi.StringPtrInput
	WorkspaceName     pulumi.StringInput
}

func (AnomaliesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*anomaliesArgs)(nil)).Elem()
}

type AnomaliesInput interface {
	pulumi.Input

	ToAnomaliesOutput() AnomaliesOutput
	ToAnomaliesOutputWithContext(ctx context.Context) AnomaliesOutput
}

func (*Anomalies) ElementType() reflect.Type {
	return reflect.TypeOf((**Anomalies)(nil)).Elem()
}

func (i *Anomalies) ToAnomaliesOutput() AnomaliesOutput {
	return i.ToAnomaliesOutputWithContext(context.Background())
}

func (i *Anomalies) ToAnomaliesOutputWithContext(ctx context.Context) AnomaliesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AnomaliesOutput)
}

type AnomaliesOutput struct{ *pulumi.OutputState }

func (AnomaliesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Anomalies)(nil)).Elem()
}

func (o AnomaliesOutput) ToAnomaliesOutput() AnomaliesOutput {
	return o
}

func (o AnomaliesOutput) ToAnomaliesOutputWithContext(ctx context.Context) AnomaliesOutput {
	return o
}

func (o AnomaliesOutput) Etag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Anomalies) pulumi.StringPtrOutput { return v.Etag }).(pulumi.StringPtrOutput)
}

func (o AnomaliesOutput) IsEnabled() pulumi.BoolOutput {
	return o.ApplyT(func(v *Anomalies) pulumi.BoolOutput { return v.IsEnabled }).(pulumi.BoolOutput)
}

func (o AnomaliesOutput) Kind() pulumi.StringOutput {
	return o.ApplyT(func(v *Anomalies) pulumi.StringOutput { return v.Kind }).(pulumi.StringOutput)
}

func (o AnomaliesOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Anomalies) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o AnomaliesOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *Anomalies) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

func (o AnomaliesOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *Anomalies) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(AnomaliesOutput{})
}
