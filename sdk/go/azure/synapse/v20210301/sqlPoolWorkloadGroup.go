


package v20210301

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type SqlPoolWorkloadGroup struct {
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


func NewSqlPoolWorkloadGroup(ctx *pulumi.Context,
	name string, args *SqlPoolWorkloadGroupArgs, opts ...pulumi.ResourceOption) (*SqlPoolWorkloadGroup, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
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
	if args.SqlPoolName == nil {
		return nil, errors.New("invalid value for required argument 'SqlPoolName'")
	}
	if args.WorkspaceName == nil {
		return nil, errors.New("invalid value for required argument 'WorkspaceName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:synapse:SqlPoolWorkloadGroup"),
		},
		{
			Type: pulumi.String("azure-native:synapse/v20190601preview:SqlPoolWorkloadGroup"),
		},
		{
			Type: pulumi.String("azure-native:synapse/v20201201:SqlPoolWorkloadGroup"),
		},
		{
			Type: pulumi.String("azure-native:synapse/v20210401preview:SqlPoolWorkloadGroup"),
		},
		{
			Type: pulumi.String("azure-native:synapse/v20210501:SqlPoolWorkloadGroup"),
		},
		{
			Type: pulumi.String("azure-native:synapse/v20210601:SqlPoolWorkloadGroup"),
		},
		{
			Type: pulumi.String("azure-native:synapse/v20210601preview:SqlPoolWorkloadGroup"),
		},
	})
	opts = append(opts, aliases)
	var resource SqlPoolWorkloadGroup
	err := ctx.RegisterResource("azure-native:synapse/v20210301:SqlPoolWorkloadGroup", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetSqlPoolWorkloadGroup(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SqlPoolWorkloadGroupState, opts ...pulumi.ResourceOption) (*SqlPoolWorkloadGroup, error) {
	var resource SqlPoolWorkloadGroup
	err := ctx.ReadResource("azure-native:synapse/v20210301:SqlPoolWorkloadGroup", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type sqlPoolWorkloadGroupState struct {
}

type SqlPoolWorkloadGroupState struct {
}

func (SqlPoolWorkloadGroupState) ElementType() reflect.Type {
	return reflect.TypeOf((*sqlPoolWorkloadGroupState)(nil)).Elem()
}

type sqlPoolWorkloadGroupArgs struct {
	Importance                   *string  `pulumi:"importance"`
	MaxResourcePercent           int      `pulumi:"maxResourcePercent"`
	MaxResourcePercentPerRequest *float64 `pulumi:"maxResourcePercentPerRequest"`
	MinResourcePercent           int      `pulumi:"minResourcePercent"`
	MinResourcePercentPerRequest float64  `pulumi:"minResourcePercentPerRequest"`
	QueryExecutionTimeout        *int     `pulumi:"queryExecutionTimeout"`
	ResourceGroupName            string   `pulumi:"resourceGroupName"`
	SqlPoolName                  string   `pulumi:"sqlPoolName"`
	WorkloadGroupName            *string  `pulumi:"workloadGroupName"`
	WorkspaceName                string   `pulumi:"workspaceName"`
}


type SqlPoolWorkloadGroupArgs struct {
	Importance                   pulumi.StringPtrInput
	MaxResourcePercent           pulumi.IntInput
	MaxResourcePercentPerRequest pulumi.Float64PtrInput
	MinResourcePercent           pulumi.IntInput
	MinResourcePercentPerRequest pulumi.Float64Input
	QueryExecutionTimeout        pulumi.IntPtrInput
	ResourceGroupName            pulumi.StringInput
	SqlPoolName                  pulumi.StringInput
	WorkloadGroupName            pulumi.StringPtrInput
	WorkspaceName                pulumi.StringInput
}

func (SqlPoolWorkloadGroupArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*sqlPoolWorkloadGroupArgs)(nil)).Elem()
}

type SqlPoolWorkloadGroupInput interface {
	pulumi.Input

	ToSqlPoolWorkloadGroupOutput() SqlPoolWorkloadGroupOutput
	ToSqlPoolWorkloadGroupOutputWithContext(ctx context.Context) SqlPoolWorkloadGroupOutput
}

func (*SqlPoolWorkloadGroup) ElementType() reflect.Type {
	return reflect.TypeOf((**SqlPoolWorkloadGroup)(nil)).Elem()
}

func (i *SqlPoolWorkloadGroup) ToSqlPoolWorkloadGroupOutput() SqlPoolWorkloadGroupOutput {
	return i.ToSqlPoolWorkloadGroupOutputWithContext(context.Background())
}

func (i *SqlPoolWorkloadGroup) ToSqlPoolWorkloadGroupOutputWithContext(ctx context.Context) SqlPoolWorkloadGroupOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SqlPoolWorkloadGroupOutput)
}

type SqlPoolWorkloadGroupOutput struct{ *pulumi.OutputState }

func (SqlPoolWorkloadGroupOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SqlPoolWorkloadGroup)(nil)).Elem()
}

func (o SqlPoolWorkloadGroupOutput) ToSqlPoolWorkloadGroupOutput() SqlPoolWorkloadGroupOutput {
	return o
}

func (o SqlPoolWorkloadGroupOutput) ToSqlPoolWorkloadGroupOutputWithContext(ctx context.Context) SqlPoolWorkloadGroupOutput {
	return o
}

func (o SqlPoolWorkloadGroupOutput) Importance() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SqlPoolWorkloadGroup) pulumi.StringPtrOutput { return v.Importance }).(pulumi.StringPtrOutput)
}

func (o SqlPoolWorkloadGroupOutput) MaxResourcePercent() pulumi.IntOutput {
	return o.ApplyT(func(v *SqlPoolWorkloadGroup) pulumi.IntOutput { return v.MaxResourcePercent }).(pulumi.IntOutput)
}

func (o SqlPoolWorkloadGroupOutput) MaxResourcePercentPerRequest() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v *SqlPoolWorkloadGroup) pulumi.Float64PtrOutput { return v.MaxResourcePercentPerRequest }).(pulumi.Float64PtrOutput)
}

func (o SqlPoolWorkloadGroupOutput) MinResourcePercent() pulumi.IntOutput {
	return o.ApplyT(func(v *SqlPoolWorkloadGroup) pulumi.IntOutput { return v.MinResourcePercent }).(pulumi.IntOutput)
}

func (o SqlPoolWorkloadGroupOutput) MinResourcePercentPerRequest() pulumi.Float64Output {
	return o.ApplyT(func(v *SqlPoolWorkloadGroup) pulumi.Float64Output { return v.MinResourcePercentPerRequest }).(pulumi.Float64Output)
}

func (o SqlPoolWorkloadGroupOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *SqlPoolWorkloadGroup) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o SqlPoolWorkloadGroupOutput) QueryExecutionTimeout() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *SqlPoolWorkloadGroup) pulumi.IntPtrOutput { return v.QueryExecutionTimeout }).(pulumi.IntPtrOutput)
}

func (o SqlPoolWorkloadGroupOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *SqlPoolWorkloadGroup) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(SqlPoolWorkloadGroupOutput{})
}
