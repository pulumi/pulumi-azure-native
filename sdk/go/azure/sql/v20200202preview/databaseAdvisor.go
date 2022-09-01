


package v20200202preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type DatabaseAdvisor struct {
	pulumi.CustomResourceState

	AdvisorStatus                  pulumi.StringOutput                  `pulumi:"advisorStatus"`
	AutoExecuteStatus              pulumi.StringOutput                  `pulumi:"autoExecuteStatus"`
	AutoExecuteStatusInheritedFrom pulumi.StringOutput                  `pulumi:"autoExecuteStatusInheritedFrom"`
	Kind                           pulumi.StringOutput                  `pulumi:"kind"`
	LastChecked                    pulumi.StringOutput                  `pulumi:"lastChecked"`
	Location                       pulumi.StringOutput                  `pulumi:"location"`
	Name                           pulumi.StringOutput                  `pulumi:"name"`
	RecommendationsStatus          pulumi.StringOutput                  `pulumi:"recommendationsStatus"`
	RecommendedActions             RecommendedActionResponseArrayOutput `pulumi:"recommendedActions"`
	Type                           pulumi.StringOutput                  `pulumi:"type"`
}


func NewDatabaseAdvisor(ctx *pulumi.Context,
	name string, args *DatabaseAdvisorArgs, opts ...pulumi.ResourceOption) (*DatabaseAdvisor, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AutoExecuteStatus == nil {
		return nil, errors.New("invalid value for required argument 'AutoExecuteStatus'")
	}
	if args.DatabaseName == nil {
		return nil, errors.New("invalid value for required argument 'DatabaseName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.ServerName == nil {
		return nil, errors.New("invalid value for required argument 'ServerName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:sql:DatabaseAdvisor"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20140401:DatabaseAdvisor"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20150501preview:DatabaseAdvisor"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20200801preview:DatabaseAdvisor"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20201101preview:DatabaseAdvisor"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20210201preview:DatabaseAdvisor"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20210501preview:DatabaseAdvisor"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20210801preview:DatabaseAdvisor"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20211101:DatabaseAdvisor"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20211101preview:DatabaseAdvisor"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20220201preview:DatabaseAdvisor"),
		},
	})
	opts = append(opts, aliases)
	var resource DatabaseAdvisor
	err := ctx.RegisterResource("azure-native:sql/v20200202preview:DatabaseAdvisor", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetDatabaseAdvisor(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DatabaseAdvisorState, opts ...pulumi.ResourceOption) (*DatabaseAdvisor, error) {
	var resource DatabaseAdvisor
	err := ctx.ReadResource("azure-native:sql/v20200202preview:DatabaseAdvisor", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type databaseAdvisorState struct {
}

type DatabaseAdvisorState struct {
}

func (DatabaseAdvisorState) ElementType() reflect.Type {
	return reflect.TypeOf((*databaseAdvisorState)(nil)).Elem()
}

type databaseAdvisorArgs struct {
	AdvisorName       *string           `pulumi:"advisorName"`
	AutoExecuteStatus AutoExecuteStatus `pulumi:"autoExecuteStatus"`
	DatabaseName      string            `pulumi:"databaseName"`
	ResourceGroupName string            `pulumi:"resourceGroupName"`
	ServerName        string            `pulumi:"serverName"`
}


type DatabaseAdvisorArgs struct {
	AdvisorName       pulumi.StringPtrInput
	AutoExecuteStatus AutoExecuteStatusInput
	DatabaseName      pulumi.StringInput
	ResourceGroupName pulumi.StringInput
	ServerName        pulumi.StringInput
}

func (DatabaseAdvisorArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*databaseAdvisorArgs)(nil)).Elem()
}

type DatabaseAdvisorInput interface {
	pulumi.Input

	ToDatabaseAdvisorOutput() DatabaseAdvisorOutput
	ToDatabaseAdvisorOutputWithContext(ctx context.Context) DatabaseAdvisorOutput
}

func (*DatabaseAdvisor) ElementType() reflect.Type {
	return reflect.TypeOf((**DatabaseAdvisor)(nil)).Elem()
}

func (i *DatabaseAdvisor) ToDatabaseAdvisorOutput() DatabaseAdvisorOutput {
	return i.ToDatabaseAdvisorOutputWithContext(context.Background())
}

func (i *DatabaseAdvisor) ToDatabaseAdvisorOutputWithContext(ctx context.Context) DatabaseAdvisorOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DatabaseAdvisorOutput)
}

type DatabaseAdvisorOutput struct{ *pulumi.OutputState }

func (DatabaseAdvisorOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DatabaseAdvisor)(nil)).Elem()
}

func (o DatabaseAdvisorOutput) ToDatabaseAdvisorOutput() DatabaseAdvisorOutput {
	return o
}

func (o DatabaseAdvisorOutput) ToDatabaseAdvisorOutputWithContext(ctx context.Context) DatabaseAdvisorOutput {
	return o
}

func (o DatabaseAdvisorOutput) AdvisorStatus() pulumi.StringOutput {
	return o.ApplyT(func(v *DatabaseAdvisor) pulumi.StringOutput { return v.AdvisorStatus }).(pulumi.StringOutput)
}

func (o DatabaseAdvisorOutput) AutoExecuteStatus() pulumi.StringOutput {
	return o.ApplyT(func(v *DatabaseAdvisor) pulumi.StringOutput { return v.AutoExecuteStatus }).(pulumi.StringOutput)
}

func (o DatabaseAdvisorOutput) AutoExecuteStatusInheritedFrom() pulumi.StringOutput {
	return o.ApplyT(func(v *DatabaseAdvisor) pulumi.StringOutput { return v.AutoExecuteStatusInheritedFrom }).(pulumi.StringOutput)
}

func (o DatabaseAdvisorOutput) Kind() pulumi.StringOutput {
	return o.ApplyT(func(v *DatabaseAdvisor) pulumi.StringOutput { return v.Kind }).(pulumi.StringOutput)
}

func (o DatabaseAdvisorOutput) LastChecked() pulumi.StringOutput {
	return o.ApplyT(func(v *DatabaseAdvisor) pulumi.StringOutput { return v.LastChecked }).(pulumi.StringOutput)
}

func (o DatabaseAdvisorOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *DatabaseAdvisor) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

func (o DatabaseAdvisorOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *DatabaseAdvisor) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o DatabaseAdvisorOutput) RecommendationsStatus() pulumi.StringOutput {
	return o.ApplyT(func(v *DatabaseAdvisor) pulumi.StringOutput { return v.RecommendationsStatus }).(pulumi.StringOutput)
}

func (o DatabaseAdvisorOutput) RecommendedActions() RecommendedActionResponseArrayOutput {
	return o.ApplyT(func(v *DatabaseAdvisor) RecommendedActionResponseArrayOutput { return v.RecommendedActions }).(RecommendedActionResponseArrayOutput)
}

func (o DatabaseAdvisorOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *DatabaseAdvisor) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(DatabaseAdvisorOutput{})
}
