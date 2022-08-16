


package v20210501preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type SyncAgent struct {
	pulumi.CustomResourceState

	ExpiryTime     pulumi.StringOutput    `pulumi:"expiryTime"`
	IsUpToDate     pulumi.BoolOutput      `pulumi:"isUpToDate"`
	LastAliveTime  pulumi.StringOutput    `pulumi:"lastAliveTime"`
	Name           pulumi.StringOutput    `pulumi:"name"`
	State          pulumi.StringOutput    `pulumi:"state"`
	SyncDatabaseId pulumi.StringPtrOutput `pulumi:"syncDatabaseId"`
	Type           pulumi.StringOutput    `pulumi:"type"`
	Version        pulumi.StringOutput    `pulumi:"version"`
}


func NewSyncAgent(ctx *pulumi.Context,
	name string, args *SyncAgentArgs, opts ...pulumi.ResourceOption) (*SyncAgent, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.ServerName == nil {
		return nil, errors.New("invalid value for required argument 'ServerName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:sql:SyncAgent"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20150501preview:SyncAgent"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20200202preview:SyncAgent"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20200801preview:SyncAgent"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20201101preview:SyncAgent"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20210201preview:SyncAgent"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20210801preview:SyncAgent"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20211101:SyncAgent"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20211101preview:SyncAgent"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20220201preview:SyncAgent"),
		},
	})
	opts = append(opts, aliases)
	var resource SyncAgent
	err := ctx.RegisterResource("azure-native:sql/v20210501preview:SyncAgent", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetSyncAgent(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SyncAgentState, opts ...pulumi.ResourceOption) (*SyncAgent, error) {
	var resource SyncAgent
	err := ctx.ReadResource("azure-native:sql/v20210501preview:SyncAgent", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type syncAgentState struct {
}

type SyncAgentState struct {
}

func (SyncAgentState) ElementType() reflect.Type {
	return reflect.TypeOf((*syncAgentState)(nil)).Elem()
}

type syncAgentArgs struct {
	ResourceGroupName string  `pulumi:"resourceGroupName"`
	ServerName        string  `pulumi:"serverName"`
	SyncAgentName     *string `pulumi:"syncAgentName"`
	SyncDatabaseId    *string `pulumi:"syncDatabaseId"`
}


type SyncAgentArgs struct {
	ResourceGroupName pulumi.StringInput
	ServerName        pulumi.StringInput
	SyncAgentName     pulumi.StringPtrInput
	SyncDatabaseId    pulumi.StringPtrInput
}

func (SyncAgentArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*syncAgentArgs)(nil)).Elem()
}

type SyncAgentInput interface {
	pulumi.Input

	ToSyncAgentOutput() SyncAgentOutput
	ToSyncAgentOutputWithContext(ctx context.Context) SyncAgentOutput
}

func (*SyncAgent) ElementType() reflect.Type {
	return reflect.TypeOf((**SyncAgent)(nil)).Elem()
}

func (i *SyncAgent) ToSyncAgentOutput() SyncAgentOutput {
	return i.ToSyncAgentOutputWithContext(context.Background())
}

func (i *SyncAgent) ToSyncAgentOutputWithContext(ctx context.Context) SyncAgentOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SyncAgentOutput)
}

type SyncAgentOutput struct{ *pulumi.OutputState }

func (SyncAgentOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SyncAgent)(nil)).Elem()
}

func (o SyncAgentOutput) ToSyncAgentOutput() SyncAgentOutput {
	return o
}

func (o SyncAgentOutput) ToSyncAgentOutputWithContext(ctx context.Context) SyncAgentOutput {
	return o
}

func (o SyncAgentOutput) ExpiryTime() pulumi.StringOutput {
	return o.ApplyT(func(v *SyncAgent) pulumi.StringOutput { return v.ExpiryTime }).(pulumi.StringOutput)
}

func (o SyncAgentOutput) IsUpToDate() pulumi.BoolOutput {
	return o.ApplyT(func(v *SyncAgent) pulumi.BoolOutput { return v.IsUpToDate }).(pulumi.BoolOutput)
}

func (o SyncAgentOutput) LastAliveTime() pulumi.StringOutput {
	return o.ApplyT(func(v *SyncAgent) pulumi.StringOutput { return v.LastAliveTime }).(pulumi.StringOutput)
}

func (o SyncAgentOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *SyncAgent) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o SyncAgentOutput) State() pulumi.StringOutput {
	return o.ApplyT(func(v *SyncAgent) pulumi.StringOutput { return v.State }).(pulumi.StringOutput)
}

func (o SyncAgentOutput) SyncDatabaseId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SyncAgent) pulumi.StringPtrOutput { return v.SyncDatabaseId }).(pulumi.StringPtrOutput)
}

func (o SyncAgentOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *SyncAgent) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func (o SyncAgentOutput) Version() pulumi.StringOutput {
	return o.ApplyT(func(v *SyncAgent) pulumi.StringOutput { return v.Version }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(SyncAgentOutput{})
}
