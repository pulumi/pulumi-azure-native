


package v20170301preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type JobAgent struct {
	pulumi.CustomResourceState

	DatabaseId pulumi.StringOutput    `pulumi:"databaseId"`
	Location   pulumi.StringOutput    `pulumi:"location"`
	Name       pulumi.StringOutput    `pulumi:"name"`
	Sku        SkuResponsePtrOutput   `pulumi:"sku"`
	State      pulumi.StringOutput    `pulumi:"state"`
	Tags       pulumi.StringMapOutput `pulumi:"tags"`
	Type       pulumi.StringOutput    `pulumi:"type"`
}


func NewJobAgent(ctx *pulumi.Context,
	name string, args *JobAgentArgs, opts ...pulumi.ResourceOption) (*JobAgent, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.DatabaseId == nil {
		return nil, errors.New("invalid value for required argument 'DatabaseId'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.ServerName == nil {
		return nil, errors.New("invalid value for required argument 'ServerName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:sql:JobAgent"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20200202preview:JobAgent"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20200801preview:JobAgent"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20201101preview:JobAgent"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20210201preview:JobAgent"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20210501preview:JobAgent"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20210801preview:JobAgent"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20211101:JobAgent"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20211101preview:JobAgent"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20220201preview:JobAgent"),
		},
	})
	opts = append(opts, aliases)
	var resource JobAgent
	err := ctx.RegisterResource("azure-native:sql/v20170301preview:JobAgent", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetJobAgent(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *JobAgentState, opts ...pulumi.ResourceOption) (*JobAgent, error) {
	var resource JobAgent
	err := ctx.ReadResource("azure-native:sql/v20170301preview:JobAgent", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type jobAgentState struct {
}

type JobAgentState struct {
}

func (JobAgentState) ElementType() reflect.Type {
	return reflect.TypeOf((*jobAgentState)(nil)).Elem()
}

type jobAgentArgs struct {
	DatabaseId        string            `pulumi:"databaseId"`
	JobAgentName      *string           `pulumi:"jobAgentName"`
	Location          *string           `pulumi:"location"`
	ResourceGroupName string            `pulumi:"resourceGroupName"`
	ServerName        string            `pulumi:"serverName"`
	Sku               *Sku              `pulumi:"sku"`
	Tags              map[string]string `pulumi:"tags"`
}


type JobAgentArgs struct {
	DatabaseId        pulumi.StringInput
	JobAgentName      pulumi.StringPtrInput
	Location          pulumi.StringPtrInput
	ResourceGroupName pulumi.StringInput
	ServerName        pulumi.StringInput
	Sku               SkuPtrInput
	Tags              pulumi.StringMapInput
}

func (JobAgentArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*jobAgentArgs)(nil)).Elem()
}

type JobAgentInput interface {
	pulumi.Input

	ToJobAgentOutput() JobAgentOutput
	ToJobAgentOutputWithContext(ctx context.Context) JobAgentOutput
}

func (*JobAgent) ElementType() reflect.Type {
	return reflect.TypeOf((**JobAgent)(nil)).Elem()
}

func (i *JobAgent) ToJobAgentOutput() JobAgentOutput {
	return i.ToJobAgentOutputWithContext(context.Background())
}

func (i *JobAgent) ToJobAgentOutputWithContext(ctx context.Context) JobAgentOutput {
	return pulumi.ToOutputWithContext(ctx, i).(JobAgentOutput)
}

type JobAgentOutput struct{ *pulumi.OutputState }

func (JobAgentOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**JobAgent)(nil)).Elem()
}

func (o JobAgentOutput) ToJobAgentOutput() JobAgentOutput {
	return o
}

func (o JobAgentOutput) ToJobAgentOutputWithContext(ctx context.Context) JobAgentOutput {
	return o
}

func (o JobAgentOutput) DatabaseId() pulumi.StringOutput {
	return o.ApplyT(func(v *JobAgent) pulumi.StringOutput { return v.DatabaseId }).(pulumi.StringOutput)
}

func (o JobAgentOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *JobAgent) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

func (o JobAgentOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *JobAgent) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o JobAgentOutput) Sku() SkuResponsePtrOutput {
	return o.ApplyT(func(v *JobAgent) SkuResponsePtrOutput { return v.Sku }).(SkuResponsePtrOutput)
}

func (o JobAgentOutput) State() pulumi.StringOutput {
	return o.ApplyT(func(v *JobAgent) pulumi.StringOutput { return v.State }).(pulumi.StringOutput)
}

func (o JobAgentOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *JobAgent) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

func (o JobAgentOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *JobAgent) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(JobAgentOutput{})
}
