


package confidentialledger

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type Ledger struct {
	pulumi.CustomResourceState

	Location   pulumi.StringPtrOutput         `pulumi:"location"`
	Name       pulumi.StringOutput            `pulumi:"name"`
	Properties LedgerPropertiesResponseOutput `pulumi:"properties"`
	SystemData SystemDataResponseOutput       `pulumi:"systemData"`
	Tags       pulumi.StringMapOutput         `pulumi:"tags"`
	Type       pulumi.StringOutput            `pulumi:"type"`
}


func NewLedger(ctx *pulumi.Context,
	name string, args *LedgerArgs, opts ...pulumi.ResourceOption) (*Ledger, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:confidentialledger/v20201201preview:Ledger"),
		},
		{
			Type: pulumi.String("azure-native:confidentialledger/v20210513preview:Ledger"),
		},
		{
			Type: pulumi.String("azure-native:confidentialledger/v20220513:Ledger"),
		},
	})
	opts = append(opts, aliases)
	var resource Ledger
	err := ctx.RegisterResource("azure-native:confidentialledger:Ledger", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetLedger(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *LedgerState, opts ...pulumi.ResourceOption) (*Ledger, error) {
	var resource Ledger
	err := ctx.ReadResource("azure-native:confidentialledger:Ledger", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type ledgerState struct {
}

type LedgerState struct {
}

func (LedgerState) ElementType() reflect.Type {
	return reflect.TypeOf((*ledgerState)(nil)).Elem()
}

type ledgerArgs struct {
	LedgerName        *string           `pulumi:"ledgerName"`
	Location          *string           `pulumi:"location"`
	Properties        *LedgerProperties `pulumi:"properties"`
	ResourceGroupName string            `pulumi:"resourceGroupName"`
	Tags              map[string]string `pulumi:"tags"`
}


type LedgerArgs struct {
	LedgerName        pulumi.StringPtrInput
	Location          pulumi.StringPtrInput
	Properties        LedgerPropertiesPtrInput
	ResourceGroupName pulumi.StringInput
	Tags              pulumi.StringMapInput
}

func (LedgerArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ledgerArgs)(nil)).Elem()
}

type LedgerInput interface {
	pulumi.Input

	ToLedgerOutput() LedgerOutput
	ToLedgerOutputWithContext(ctx context.Context) LedgerOutput
}

func (*Ledger) ElementType() reflect.Type {
	return reflect.TypeOf((**Ledger)(nil)).Elem()
}

func (i *Ledger) ToLedgerOutput() LedgerOutput {
	return i.ToLedgerOutputWithContext(context.Background())
}

func (i *Ledger) ToLedgerOutputWithContext(ctx context.Context) LedgerOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LedgerOutput)
}

type LedgerOutput struct{ *pulumi.OutputState }

func (LedgerOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Ledger)(nil)).Elem()
}

func (o LedgerOutput) ToLedgerOutput() LedgerOutput {
	return o
}

func (o LedgerOutput) ToLedgerOutputWithContext(ctx context.Context) LedgerOutput {
	return o
}

func (o LedgerOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Ledger) pulumi.StringPtrOutput { return v.Location }).(pulumi.StringPtrOutput)
}

func (o LedgerOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Ledger) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o LedgerOutput) Properties() LedgerPropertiesResponseOutput {
	return o.ApplyT(func(v *Ledger) LedgerPropertiesResponseOutput { return v.Properties }).(LedgerPropertiesResponseOutput)
}

func (o LedgerOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *Ledger) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

func (o LedgerOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *Ledger) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

func (o LedgerOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *Ledger) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LedgerOutput{})
}
