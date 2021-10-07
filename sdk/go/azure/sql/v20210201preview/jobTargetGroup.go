


package v20210201preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type JobTargetGroup struct {
	pulumi.CustomResourceState

	Members JobTargetResponseArrayOutput `pulumi:"members"`
	Name    pulumi.StringOutput          `pulumi:"name"`
	Type    pulumi.StringOutput          `pulumi:"type"`
}


func NewJobTargetGroup(ctx *pulumi.Context,
	name string, args *JobTargetGroupArgs, opts ...pulumi.ResourceOption) (*JobTargetGroup, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.JobAgentName == nil {
		return nil, errors.New("invalid value for required argument 'JobAgentName'")
	}
	if args.Members == nil {
		return nil, errors.New("invalid value for required argument 'Members'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.ServerName == nil {
		return nil, errors.New("invalid value for required argument 'ServerName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-nextgen:sql/v20210201preview:JobTargetGroup"),
		},
		{
			Type: pulumi.String("azure-native:sql:JobTargetGroup"),
		},
		{
			Type: pulumi.String("azure-nextgen:sql:JobTargetGroup"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20170301preview:JobTargetGroup"),
		},
		{
			Type: pulumi.String("azure-nextgen:sql/v20170301preview:JobTargetGroup"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20200202preview:JobTargetGroup"),
		},
		{
			Type: pulumi.String("azure-nextgen:sql/v20200202preview:JobTargetGroup"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20200801preview:JobTargetGroup"),
		},
		{
			Type: pulumi.String("azure-nextgen:sql/v20200801preview:JobTargetGroup"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20201101preview:JobTargetGroup"),
		},
		{
			Type: pulumi.String("azure-nextgen:sql/v20201101preview:JobTargetGroup"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20210501preview:JobTargetGroup"),
		},
		{
			Type: pulumi.String("azure-nextgen:sql/v20210501preview:JobTargetGroup"),
		},
	})
	opts = append(opts, aliases)
	var resource JobTargetGroup
	err := ctx.RegisterResource("azure-native:sql/v20210201preview:JobTargetGroup", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetJobTargetGroup(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *JobTargetGroupState, opts ...pulumi.ResourceOption) (*JobTargetGroup, error) {
	var resource JobTargetGroup
	err := ctx.ReadResource("azure-native:sql/v20210201preview:JobTargetGroup", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type jobTargetGroupState struct {
}

type JobTargetGroupState struct {
}

func (JobTargetGroupState) ElementType() reflect.Type {
	return reflect.TypeOf((*jobTargetGroupState)(nil)).Elem()
}

type jobTargetGroupArgs struct {
	JobAgentName      string      `pulumi:"jobAgentName"`
	Members           []JobTarget `pulumi:"members"`
	ResourceGroupName string      `pulumi:"resourceGroupName"`
	ServerName        string      `pulumi:"serverName"`
	TargetGroupName   *string     `pulumi:"targetGroupName"`
}


type JobTargetGroupArgs struct {
	JobAgentName      pulumi.StringInput
	Members           JobTargetArrayInput
	ResourceGroupName pulumi.StringInput
	ServerName        pulumi.StringInput
	TargetGroupName   pulumi.StringPtrInput
}

func (JobTargetGroupArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*jobTargetGroupArgs)(nil)).Elem()
}

type JobTargetGroupInput interface {
	pulumi.Input

	ToJobTargetGroupOutput() JobTargetGroupOutput
	ToJobTargetGroupOutputWithContext(ctx context.Context) JobTargetGroupOutput
}

func (*JobTargetGroup) ElementType() reflect.Type {
	return reflect.TypeOf((*JobTargetGroup)(nil))
}

func (i *JobTargetGroup) ToJobTargetGroupOutput() JobTargetGroupOutput {
	return i.ToJobTargetGroupOutputWithContext(context.Background())
}

func (i *JobTargetGroup) ToJobTargetGroupOutputWithContext(ctx context.Context) JobTargetGroupOutput {
	return pulumi.ToOutputWithContext(ctx, i).(JobTargetGroupOutput)
}

type JobTargetGroupOutput struct{ *pulumi.OutputState }

func (JobTargetGroupOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*JobTargetGroup)(nil))
}

func (o JobTargetGroupOutput) ToJobTargetGroupOutput() JobTargetGroupOutput {
	return o
}

func (o JobTargetGroupOutput) ToJobTargetGroupOutputWithContext(ctx context.Context) JobTargetGroupOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(JobTargetGroupOutput{})
}
