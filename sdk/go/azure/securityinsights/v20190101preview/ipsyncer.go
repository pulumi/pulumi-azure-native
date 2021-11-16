


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
	Etag                                *string `pulumi:"etag"`
	Kind                                string  `pulumi:"kind"`
	OperationalInsightsResourceProvider string  `pulumi:"operationalInsightsResourceProvider"`
	ResourceGroupName                   string  `pulumi:"resourceGroupName"`
	SettingsName                        *string `pulumi:"settingsName"`
	WorkspaceName                       string  `pulumi:"workspaceName"`
}


type IPSyncerArgs struct {
	Etag                                pulumi.StringPtrInput
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
	return reflect.TypeOf((*IPSyncer)(nil))
}

func (i *IPSyncer) ToIPSyncerOutput() IPSyncerOutput {
	return i.ToIPSyncerOutputWithContext(context.Background())
}

func (i *IPSyncer) ToIPSyncerOutputWithContext(ctx context.Context) IPSyncerOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IPSyncerOutput)
}

type IPSyncerOutput struct{ *pulumi.OutputState }

func (IPSyncerOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*IPSyncer)(nil))
}

func (o IPSyncerOutput) ToIPSyncerOutput() IPSyncerOutput {
	return o
}

func (o IPSyncerOutput) ToIPSyncerOutputWithContext(ctx context.Context) IPSyncerOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(IPSyncerOutput{})
}
