


package v20200801preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type WorkloadGroup struct {
	pulumi.CustomResourceState

	Importance                   pulumi.StringPtrOutput  `pulumi:"importance"`
	MaxResourcePercent           pulumi.IntOutput        `pulumi:"maxResourcePercent"`
	MaxResourcePercentPerRequest pulumi.Float64PtrOutput `pulumi:"maxResourcePercentPerRequest"`
	MinResourcePercent           pulumi.IntOutput        `pulumi:"minResourcePercent"`
	MinResourcePercentPerRequest pulumi.Float64Output    `pulumi:"minResourcePercentPerRequest"`
	Name                         pulumi.StringOutput     `pulumi:"name"`
	QueryExecutionTimeout        pulumi.IntPtrOutput     `pulumi:"queryExecutionTimeout"`
	Type                         pulumi.StringOutput     `pulumi:"type"`
}


func NewWorkloadGroup(ctx *pulumi.Context,
	name string, args *WorkloadGroupArgs, opts ...pulumi.ResourceOption) (*WorkloadGroup, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.DatabaseName == nil {
		return nil, errors.New("invalid value for required argument 'DatabaseName'")
	}
	if args.MaxResourcePercent == nil {
		return nil, errors.New("invalid value for required argument 'MaxResourcePercent'")
	}
	if args.MinResourcePercent == nil {
		return nil, errors.New("invalid value for required argument 'MinResourcePercent'")
	}
	if args.MinResourcePercentPerRequest == nil {
		return nil, errors.New("invalid value for required argument 'MinResourcePercentPerRequest'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.ServerName == nil {
		return nil, errors.New("invalid value for required argument 'ServerName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:sql:WorkloadGroup"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20190601preview:WorkloadGroup"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20200202preview:WorkloadGroup"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20201101preview:WorkloadGroup"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20210201preview:WorkloadGroup"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20210501preview:WorkloadGroup"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20210801preview:WorkloadGroup"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20211101:WorkloadGroup"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20211101preview:WorkloadGroup"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20220201preview:WorkloadGroup"),
		},
	})
	opts = append(opts, aliases)
	var resource WorkloadGroup
	err := ctx.RegisterResource("azure-native:sql/v20200801preview:WorkloadGroup", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetWorkloadGroup(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *WorkloadGroupState, opts ...pulumi.ResourceOption) (*WorkloadGroup, error) {
	var resource WorkloadGroup
	err := ctx.ReadResource("azure-native:sql/v20200801preview:WorkloadGroup", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type workloadGroupState struct {
}

type WorkloadGroupState struct {
}

func (WorkloadGroupState) ElementType() reflect.Type {
	return reflect.TypeOf((*workloadGroupState)(nil)).Elem()
}

type workloadGroupArgs struct {
	DatabaseName                 string   `pulumi:"databaseName"`
	Importance                   *string  `pulumi:"importance"`
	MaxResourcePercent           int      `pulumi:"maxResourcePercent"`
	MaxResourcePercentPerRequest *float64 `pulumi:"maxResourcePercentPerRequest"`
	MinResourcePercent           int      `pulumi:"minResourcePercent"`
	MinResourcePercentPerRequest float64  `pulumi:"minResourcePercentPerRequest"`
	QueryExecutionTimeout        *int     `pulumi:"queryExecutionTimeout"`
	ResourceGroupName            string   `pulumi:"resourceGroupName"`
	ServerName                   string   `pulumi:"serverName"`
	WorkloadGroupName            *string  `pulumi:"workloadGroupName"`
}


type WorkloadGroupArgs struct {
	DatabaseName                 pulumi.StringInput
	Importance                   pulumi.StringPtrInput
	MaxResourcePercent           pulumi.IntInput
	MaxResourcePercentPerRequest pulumi.Float64PtrInput
	MinResourcePercent           pulumi.IntInput
	MinResourcePercentPerRequest pulumi.Float64Input
	QueryExecutionTimeout        pulumi.IntPtrInput
	ResourceGroupName            pulumi.StringInput
	ServerName                   pulumi.StringInput
	WorkloadGroupName            pulumi.StringPtrInput
}

func (WorkloadGroupArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*workloadGroupArgs)(nil)).Elem()
}

type WorkloadGroupInput interface {
	pulumi.Input

	ToWorkloadGroupOutput() WorkloadGroupOutput
	ToWorkloadGroupOutputWithContext(ctx context.Context) WorkloadGroupOutput
}

func (*WorkloadGroup) ElementType() reflect.Type {
	return reflect.TypeOf((**WorkloadGroup)(nil)).Elem()
}

func (i *WorkloadGroup) ToWorkloadGroupOutput() WorkloadGroupOutput {
	return i.ToWorkloadGroupOutputWithContext(context.Background())
}

func (i *WorkloadGroup) ToWorkloadGroupOutputWithContext(ctx context.Context) WorkloadGroupOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WorkloadGroupOutput)
}

type WorkloadGroupOutput struct{ *pulumi.OutputState }

func (WorkloadGroupOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**WorkloadGroup)(nil)).Elem()
}

func (o WorkloadGroupOutput) ToWorkloadGroupOutput() WorkloadGroupOutput {
	return o
}

func (o WorkloadGroupOutput) ToWorkloadGroupOutputWithContext(ctx context.Context) WorkloadGroupOutput {
	return o
}

func (o WorkloadGroupOutput) Importance() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WorkloadGroup) pulumi.StringPtrOutput { return v.Importance }).(pulumi.StringPtrOutput)
}

func (o WorkloadGroupOutput) MaxResourcePercent() pulumi.IntOutput {
	return o.ApplyT(func(v *WorkloadGroup) pulumi.IntOutput { return v.MaxResourcePercent }).(pulumi.IntOutput)
}

func (o WorkloadGroupOutput) MaxResourcePercentPerRequest() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v *WorkloadGroup) pulumi.Float64PtrOutput { return v.MaxResourcePercentPerRequest }).(pulumi.Float64PtrOutput)
}

func (o WorkloadGroupOutput) MinResourcePercent() pulumi.IntOutput {
	return o.ApplyT(func(v *WorkloadGroup) pulumi.IntOutput { return v.MinResourcePercent }).(pulumi.IntOutput)
}

func (o WorkloadGroupOutput) MinResourcePercentPerRequest() pulumi.Float64Output {
	return o.ApplyT(func(v *WorkloadGroup) pulumi.Float64Output { return v.MinResourcePercentPerRequest }).(pulumi.Float64Output)
}

func (o WorkloadGroupOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *WorkloadGroup) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o WorkloadGroupOutput) QueryExecutionTimeout() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *WorkloadGroup) pulumi.IntPtrOutput { return v.QueryExecutionTimeout }).(pulumi.IntPtrOutput)
}

func (o WorkloadGroupOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *WorkloadGroup) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(WorkloadGroupOutput{})
}
