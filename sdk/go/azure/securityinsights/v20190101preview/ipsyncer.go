


package v20190101preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type IPSyncer struct {
	pulumi.CustomResourceState

	Etag      pulumi.StringPtrOutput `pulumi:"etag"`
	IsEnabled pulumi.BoolOutput      `pulumi:"isEnabled"`
	Kind      pulumi.StringOutput    `pulumi:"kind"`
	Name      pulumi.StringOutput    `pulumi:"name"`
	Type      pulumi.StringOutput    `pulumi:"type"`
}


func NewIPSyncer(ctx *pulumi.Context,
	name string, args *IPSyncerArgs, opts ...pulumi.ResourceOption) (*IPSyncer, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Kind == nil {
		return nil, errors.New("invalid value for required argument 'Kind'")
	}
	if args.OperationalInsightsResourceProvider == nil {
		return nil, errors.New("invalid value for required argument 'OperationalInsightsResourceProvider'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.WorkspaceName == nil {
		return nil, errors.New("invalid value for required argument 'WorkspaceName'")
	}
	args.Kind = pulumi.String("IPSyncer")
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:securityinsights:IPSyncer"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20210301preview:IPSyncer"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20210901preview:IPSyncer"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20211001preview:IPSyncer"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20220101preview:IPSyncer"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20220401preview:IPSyncer"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20220501preview:IPSyncer"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20220601preview:IPSyncer"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20220701preview:IPSyncer"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20220801preview:IPSyncer"),
		},
	})
	opts = append(opts, aliases)
	var resource IPSyncer
	err := ctx.RegisterResource("azure-native:securityinsights/v20190101preview:IPSyncer", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetIPSyncer(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *IPSyncerState, opts ...pulumi.ResourceOption) (*IPSyncer, error) {
	var resource IPSyncer
	err := ctx.ReadResource("azure-native:securityinsights/v20190101preview:IPSyncer", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type ipsyncerState struct {
}

type IPSyncerState struct {
}

func (IPSyncerState) ElementType() reflect.Type {
	return reflect.TypeOf((*ipsyncerState)(nil)).Elem()
}

type ipsyncerArgs struct {
	Kind                                string  `pulumi:"kind"`
	OperationalInsightsResourceProvider string  `pulumi:"operationalInsightsResourceProvider"`
	ResourceGroupName                   string  `pulumi:"resourceGroupName"`
	SettingsName                        *string `pulumi:"settingsName"`
	WorkspaceName                       string  `pulumi:"workspaceName"`
}


type IPSyncerArgs struct {
	Kind                                pulumi.StringInput
	OperationalInsightsResourceProvider pulumi.StringInput
	ResourceGroupName                   pulumi.StringInput
	SettingsName                        pulumi.StringPtrInput
	WorkspaceName                       pulumi.StringInput
}

func (IPSyncerArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ipsyncerArgs)(nil)).Elem()
}

type IPSyncerInput interface {
	pulumi.Input

	ToIPSyncerOutput() IPSyncerOutput
	ToIPSyncerOutputWithContext(ctx context.Context) IPSyncerOutput
}

func (*IPSyncer) ElementType() reflect.Type {
	return reflect.TypeOf((**IPSyncer)(nil)).Elem()
}

func (i *IPSyncer) ToIPSyncerOutput() IPSyncerOutput {
	return i.ToIPSyncerOutputWithContext(context.Background())
}

func (i *IPSyncer) ToIPSyncerOutputWithContext(ctx context.Context) IPSyncerOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IPSyncerOutput)
}

type IPSyncerOutput struct{ *pulumi.OutputState }

func (IPSyncerOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**IPSyncer)(nil)).Elem()
}

func (o IPSyncerOutput) ToIPSyncerOutput() IPSyncerOutput {
	return o
}

func (o IPSyncerOutput) ToIPSyncerOutputWithContext(ctx context.Context) IPSyncerOutput {
	return o
}

func (o IPSyncerOutput) Etag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *IPSyncer) pulumi.StringPtrOutput { return v.Etag }).(pulumi.StringPtrOutput)
}

func (o IPSyncerOutput) IsEnabled() pulumi.BoolOutput {
	return o.ApplyT(func(v *IPSyncer) pulumi.BoolOutput { return v.IsEnabled }).(pulumi.BoolOutput)
}

func (o IPSyncerOutput) Kind() pulumi.StringOutput {
	return o.ApplyT(func(v *IPSyncer) pulumi.StringOutput { return v.Kind }).(pulumi.StringOutput)
}

func (o IPSyncerOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *IPSyncer) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o IPSyncerOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *IPSyncer) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(IPSyncerOutput{})
}
