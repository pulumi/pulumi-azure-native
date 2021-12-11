


package v20200202preview

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
			Type: pulumi.String("azure-native:sql/v20170301preview:JobAgent"),
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
	})
	opts = append(opts, aliases)
	var resource JobAgent
	err := ctx.RegisterResource("azure-native:sql/v20200202preview:JobAgent", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetJobAgent(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *JobAgentState, opts ...pulumi.ResourceOption) (*JobAgent, error) {
	var resource JobAgent
	err := ctx.ReadResource("azure-native:sql/v20200202preview:JobAgent", name, id, state, &resource, opts...)
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
	return reflect.TypeOf((*JobAgent)(nil))
}

func (i *JobAgent) ToJobAgentOutput() JobAgentOutput {
	return i.ToJobAgentOutputWithContext(context.Background())
}

func (i *JobAgent) ToJobAgentOutputWithContext(ctx context.Context) JobAgentOutput {
	return pulumi.ToOutputWithContext(ctx, i).(JobAgentOutput)
}

type JobAgentOutput struct{ *pulumi.OutputState }

func (JobAgentOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*JobAgent)(nil))
}

func (o JobAgentOutput) ToJobAgentOutput() JobAgentOutput {
	return o
}

func (o JobAgentOutput) ToJobAgentOutputWithContext(ctx context.Context) JobAgentOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(JobAgentOutput{})
}
