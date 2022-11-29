


package operationalinsights

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type Table struct {
	pulumi.CustomResourceState

	ArchiveRetentionInDays pulumi.IntOutput                  `pulumi:"archiveRetentionInDays"`
	LastPlanModifiedDate   pulumi.StringOutput               `pulumi:"lastPlanModifiedDate"`
	Name                   pulumi.StringOutput               `pulumi:"name"`
	Plan                   pulumi.StringPtrOutput            `pulumi:"plan"`
	ProvisioningState      pulumi.StringOutput               `pulumi:"provisioningState"`
	RestoredLogs           RestoredLogsResponsePtrOutput     `pulumi:"restoredLogs"`
	ResultStatistics       ResultStatisticsResponsePtrOutput `pulumi:"resultStatistics"`
	RetentionInDays        pulumi.IntPtrOutput               `pulumi:"retentionInDays"`
	Schema                 SchemaResponsePtrOutput           `pulumi:"schema"`
	SearchResults          SearchResultsResponsePtrOutput    `pulumi:"searchResults"`
	SystemData             SystemDataResponseOutput          `pulumi:"systemData"`
	TotalRetentionInDays   pulumi.IntPtrOutput               `pulumi:"totalRetentionInDays"`
	Type                   pulumi.StringOutput               `pulumi:"type"`
}


func NewTable(ctx *pulumi.Context,
	name string, args *TableArgs, opts ...pulumi.ResourceOption) (*Table, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.WorkspaceName == nil {
		return nil, errors.New("invalid value for required argument 'WorkspaceName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:operationalinsights/v20211201preview:Table"),
		},
		{
			Type: pulumi.String("azure-native:operationalinsights/v20221001:Table"),
		},
	})
	opts = append(opts, aliases)
	var resource Table
	err := ctx.RegisterResource("azure-native:operationalinsights:Table", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetTable(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *TableState, opts ...pulumi.ResourceOption) (*Table, error) {
	var resource Table
	err := ctx.ReadResource("azure-native:operationalinsights:Table", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type tableState struct {
}

type TableState struct {
}

func (TableState) ElementType() reflect.Type {
	return reflect.TypeOf((*tableState)(nil)).Elem()
}

type tableArgs struct {
	Plan                 *string        `pulumi:"plan"`
	ResourceGroupName    string         `pulumi:"resourceGroupName"`
	RestoredLogs         *RestoredLogs  `pulumi:"restoredLogs"`
	RetentionInDays      *int           `pulumi:"retentionInDays"`
	Schema               *Schema        `pulumi:"schema"`
	SearchResults        *SearchResults `pulumi:"searchResults"`
	TableName            *string        `pulumi:"tableName"`
	TotalRetentionInDays *int           `pulumi:"totalRetentionInDays"`
	WorkspaceName        string         `pulumi:"workspaceName"`
}


type TableArgs struct {
	Plan                 pulumi.StringPtrInput
	ResourceGroupName    pulumi.StringInput
	RestoredLogs         RestoredLogsPtrInput
	RetentionInDays      pulumi.IntPtrInput
	Schema               SchemaPtrInput
	SearchResults        SearchResultsPtrInput
	TableName            pulumi.StringPtrInput
	TotalRetentionInDays pulumi.IntPtrInput
	WorkspaceName        pulumi.StringInput
}

func (TableArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*tableArgs)(nil)).Elem()
}

type TableInput interface {
	pulumi.Input

	ToTableOutput() TableOutput
	ToTableOutputWithContext(ctx context.Context) TableOutput
}

func (*Table) ElementType() reflect.Type {
	return reflect.TypeOf((**Table)(nil)).Elem()
}

func (i *Table) ToTableOutput() TableOutput {
	return i.ToTableOutputWithContext(context.Background())
}

func (i *Table) ToTableOutputWithContext(ctx context.Context) TableOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TableOutput)
}

type TableOutput struct{ *pulumi.OutputState }

func (TableOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Table)(nil)).Elem()
}

func (o TableOutput) ToTableOutput() TableOutput {
	return o
}

func (o TableOutput) ToTableOutputWithContext(ctx context.Context) TableOutput {
	return o
}

func (o TableOutput) ArchiveRetentionInDays() pulumi.IntOutput {
	return o.ApplyT(func(v *Table) pulumi.IntOutput { return v.ArchiveRetentionInDays }).(pulumi.IntOutput)
}

func (o TableOutput) LastPlanModifiedDate() pulumi.StringOutput {
	return o.ApplyT(func(v *Table) pulumi.StringOutput { return v.LastPlanModifiedDate }).(pulumi.StringOutput)
}

func (o TableOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Table) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o TableOutput) Plan() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Table) pulumi.StringPtrOutput { return v.Plan }).(pulumi.StringPtrOutput)
}

func (o TableOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *Table) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o TableOutput) RestoredLogs() RestoredLogsResponsePtrOutput {
	return o.ApplyT(func(v *Table) RestoredLogsResponsePtrOutput { return v.RestoredLogs }).(RestoredLogsResponsePtrOutput)
}

func (o TableOutput) ResultStatistics() ResultStatisticsResponsePtrOutput {
	return o.ApplyT(func(v *Table) ResultStatisticsResponsePtrOutput { return v.ResultStatistics }).(ResultStatisticsResponsePtrOutput)
}

func (o TableOutput) RetentionInDays() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *Table) pulumi.IntPtrOutput { return v.RetentionInDays }).(pulumi.IntPtrOutput)
}

func (o TableOutput) Schema() SchemaResponsePtrOutput {
	return o.ApplyT(func(v *Table) SchemaResponsePtrOutput { return v.Schema }).(SchemaResponsePtrOutput)
}

func (o TableOutput) SearchResults() SearchResultsResponsePtrOutput {
	return o.ApplyT(func(v *Table) SearchResultsResponsePtrOutput { return v.SearchResults }).(SearchResultsResponsePtrOutput)
}

func (o TableOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *Table) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

func (o TableOutput) TotalRetentionInDays() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *Table) pulumi.IntPtrOutput { return v.TotalRetentionInDays }).(pulumi.IntPtrOutput)
}

func (o TableOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *Table) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(TableOutput{})
}
