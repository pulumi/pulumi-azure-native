


package insights

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type DataCollectionRuleAssociation struct {
	pulumi.CustomResourceState

	DataCollectionRuleId pulumi.StringPtrOutput `pulumi:"dataCollectionRuleId"`
	Description          pulumi.StringPtrOutput `pulumi:"description"`
	Etag                 pulumi.StringOutput    `pulumi:"etag"`
	Name                 pulumi.StringOutput    `pulumi:"name"`
	ProvisioningState    pulumi.StringOutput    `pulumi:"provisioningState"`
	Type                 pulumi.StringOutput    `pulumi:"type"`
}


func NewDataCollectionRuleAssociation(ctx *pulumi.Context,
	name string, args *DataCollectionRuleAssociationArgs, opts ...pulumi.ResourceOption) (*DataCollectionRuleAssociation, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceUri == nil {
		return nil, errors.New("invalid value for required argument 'ResourceUri'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:insights/v20191101preview:DataCollectionRuleAssociation"),
		},
		{
			Type: pulumi.String("azure-native:insights/v20210401:DataCollectionRuleAssociation"),
		},
		{
			Type: pulumi.String("azure-native:insights/v20210901preview:DataCollectionRuleAssociation"),
		},
	})
	opts = append(opts, aliases)
	var resource DataCollectionRuleAssociation
	err := ctx.RegisterResource("azure-native:insights:DataCollectionRuleAssociation", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetDataCollectionRuleAssociation(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DataCollectionRuleAssociationState, opts ...pulumi.ResourceOption) (*DataCollectionRuleAssociation, error) {
	var resource DataCollectionRuleAssociation
	err := ctx.ReadResource("azure-native:insights:DataCollectionRuleAssociation", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type dataCollectionRuleAssociationState struct {
}

type DataCollectionRuleAssociationState struct {
}

func (DataCollectionRuleAssociationState) ElementType() reflect.Type {
	return reflect.TypeOf((*dataCollectionRuleAssociationState)(nil)).Elem()
}

type dataCollectionRuleAssociationArgs struct {
	AssociationName      *string `pulumi:"associationName"`
	DataCollectionRuleId *string `pulumi:"dataCollectionRuleId"`
	Description          *string `pulumi:"description"`
	ResourceUri          string  `pulumi:"resourceUri"`
}


type DataCollectionRuleAssociationArgs struct {
	AssociationName      pulumi.StringPtrInput
	DataCollectionRuleId pulumi.StringPtrInput
	Description          pulumi.StringPtrInput
	ResourceUri          pulumi.StringInput
}

func (DataCollectionRuleAssociationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*dataCollectionRuleAssociationArgs)(nil)).Elem()
}

type DataCollectionRuleAssociationInput interface {
	pulumi.Input

	ToDataCollectionRuleAssociationOutput() DataCollectionRuleAssociationOutput
	ToDataCollectionRuleAssociationOutputWithContext(ctx context.Context) DataCollectionRuleAssociationOutput
}

func (*DataCollectionRuleAssociation) ElementType() reflect.Type {
	return reflect.TypeOf((**DataCollectionRuleAssociation)(nil)).Elem()
}

func (i *DataCollectionRuleAssociation) ToDataCollectionRuleAssociationOutput() DataCollectionRuleAssociationOutput {
	return i.ToDataCollectionRuleAssociationOutputWithContext(context.Background())
}

func (i *DataCollectionRuleAssociation) ToDataCollectionRuleAssociationOutputWithContext(ctx context.Context) DataCollectionRuleAssociationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DataCollectionRuleAssociationOutput)
}

type DataCollectionRuleAssociationOutput struct{ *pulumi.OutputState }

func (DataCollectionRuleAssociationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DataCollectionRuleAssociation)(nil)).Elem()
}

func (o DataCollectionRuleAssociationOutput) ToDataCollectionRuleAssociationOutput() DataCollectionRuleAssociationOutput {
	return o
}

func (o DataCollectionRuleAssociationOutput) ToDataCollectionRuleAssociationOutputWithContext(ctx context.Context) DataCollectionRuleAssociationOutput {
	return o
}

func (o DataCollectionRuleAssociationOutput) DataCollectionRuleId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DataCollectionRuleAssociation) pulumi.StringPtrOutput { return v.DataCollectionRuleId }).(pulumi.StringPtrOutput)
}

func (o DataCollectionRuleAssociationOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DataCollectionRuleAssociation) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

func (o DataCollectionRuleAssociationOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v *DataCollectionRuleAssociation) pulumi.StringOutput { return v.Etag }).(pulumi.StringOutput)
}

func (o DataCollectionRuleAssociationOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *DataCollectionRuleAssociation) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o DataCollectionRuleAssociationOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *DataCollectionRuleAssociation) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o DataCollectionRuleAssociationOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *DataCollectionRuleAssociation) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(DataCollectionRuleAssociationOutput{})
}
