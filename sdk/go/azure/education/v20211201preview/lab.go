


package v20211201preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type Lab struct {
	pulumi.CustomResourceState

	BudgetPerStudent AmountResponseOutput     `pulumi:"budgetPerStudent"`
	Currency         pulumi.StringPtrOutput   `pulumi:"currency"`
	Description      pulumi.StringOutput      `pulumi:"description"`
	DisplayName      pulumi.StringOutput      `pulumi:"displayName"`
	EffectiveDate    pulumi.StringOutput      `pulumi:"effectiveDate"`
	ExpirationDate   pulumi.StringOutput      `pulumi:"expirationDate"`
	InvitationCode   pulumi.StringOutput      `pulumi:"invitationCode"`
	MaxStudentCount  pulumi.Float64Output     `pulumi:"maxStudentCount"`
	Name             pulumi.StringOutput      `pulumi:"name"`
	Status           pulumi.StringOutput      `pulumi:"status"`
	SystemData       SystemDataResponseOutput `pulumi:"systemData"`
	Type             pulumi.StringOutput      `pulumi:"type"`
	Value            pulumi.Float64PtrOutput  `pulumi:"value"`
}


func NewLab(ctx *pulumi.Context,
	name string, args *LabArgs, opts ...pulumi.ResourceOption) (*Lab, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.BillingAccountName == nil {
		return nil, errors.New("invalid value for required argument 'BillingAccountName'")
	}
	if args.BillingProfileName == nil {
		return nil, errors.New("invalid value for required argument 'BillingProfileName'")
	}
	if args.BudgetPerStudent == nil {
		return nil, errors.New("invalid value for required argument 'BudgetPerStudent'")
	}
	if args.Description == nil {
		return nil, errors.New("invalid value for required argument 'Description'")
	}
	if args.DisplayName == nil {
		return nil, errors.New("invalid value for required argument 'DisplayName'")
	}
	if args.ExpirationDate == nil {
		return nil, errors.New("invalid value for required argument 'ExpirationDate'")
	}
	if args.InvoiceSectionName == nil {
		return nil, errors.New("invalid value for required argument 'InvoiceSectionName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:education:Lab"),
		},
	})
	opts = append(opts, aliases)
	var resource Lab
	err := ctx.RegisterResource("azure-native:education/v20211201preview:Lab", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetLab(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *LabState, opts ...pulumi.ResourceOption) (*Lab, error) {
	var resource Lab
	err := ctx.ReadResource("azure-native:education/v20211201preview:Lab", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type labState struct {
}

type LabState struct {
}

func (LabState) ElementType() reflect.Type {
	return reflect.TypeOf((*labState)(nil)).Elem()
}

type labArgs struct {
	BillingAccountName string   `pulumi:"billingAccountName"`
	BillingProfileName string   `pulumi:"billingProfileName"`
	BudgetPerStudent   Amount   `pulumi:"budgetPerStudent"`
	Currency           *string  `pulumi:"currency"`
	Description        string   `pulumi:"description"`
	DisplayName        string   `pulumi:"displayName"`
	ExpirationDate     string   `pulumi:"expirationDate"`
	InvoiceSectionName string   `pulumi:"invoiceSectionName"`
	Value              *float64 `pulumi:"value"`
}


type LabArgs struct {
	BillingAccountName pulumi.StringInput
	BillingProfileName pulumi.StringInput
	BudgetPerStudent   AmountInput
	Currency           pulumi.StringPtrInput
	Description        pulumi.StringInput
	DisplayName        pulumi.StringInput
	ExpirationDate     pulumi.StringInput
	InvoiceSectionName pulumi.StringInput
	Value              pulumi.Float64PtrInput
}

func (LabArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*labArgs)(nil)).Elem()
}

type LabInput interface {
	pulumi.Input

	ToLabOutput() LabOutput
	ToLabOutputWithContext(ctx context.Context) LabOutput
}

func (*Lab) ElementType() reflect.Type {
	return reflect.TypeOf((**Lab)(nil)).Elem()
}

func (i *Lab) ToLabOutput() LabOutput {
	return i.ToLabOutputWithContext(context.Background())
}

func (i *Lab) ToLabOutputWithContext(ctx context.Context) LabOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LabOutput)
}

type LabOutput struct{ *pulumi.OutputState }

func (LabOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Lab)(nil)).Elem()
}

func (o LabOutput) ToLabOutput() LabOutput {
	return o
}

func (o LabOutput) ToLabOutputWithContext(ctx context.Context) LabOutput {
	return o
}

func (o LabOutput) BudgetPerStudent() AmountResponseOutput {
	return o.ApplyT(func(v *Lab) AmountResponseOutput { return v.BudgetPerStudent }).(AmountResponseOutput)
}

func (o LabOutput) Currency() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Lab) pulumi.StringPtrOutput { return v.Currency }).(pulumi.StringPtrOutput)
}

func (o LabOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v *Lab) pulumi.StringOutput { return v.Description }).(pulumi.StringOutput)
}

func (o LabOutput) DisplayName() pulumi.StringOutput {
	return o.ApplyT(func(v *Lab) pulumi.StringOutput { return v.DisplayName }).(pulumi.StringOutput)
}

func (o LabOutput) EffectiveDate() pulumi.StringOutput {
	return o.ApplyT(func(v *Lab) pulumi.StringOutput { return v.EffectiveDate }).(pulumi.StringOutput)
}

func (o LabOutput) ExpirationDate() pulumi.StringOutput {
	return o.ApplyT(func(v *Lab) pulumi.StringOutput { return v.ExpirationDate }).(pulumi.StringOutput)
}

func (o LabOutput) InvitationCode() pulumi.StringOutput {
	return o.ApplyT(func(v *Lab) pulumi.StringOutput { return v.InvitationCode }).(pulumi.StringOutput)
}

func (o LabOutput) MaxStudentCount() pulumi.Float64Output {
	return o.ApplyT(func(v *Lab) pulumi.Float64Output { return v.MaxStudentCount }).(pulumi.Float64Output)
}

func (o LabOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Lab) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o LabOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v *Lab) pulumi.StringOutput { return v.Status }).(pulumi.StringOutput)
}

func (o LabOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *Lab) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

func (o LabOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *Lab) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func (o LabOutput) Value() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v *Lab) pulumi.Float64PtrOutput { return v.Value }).(pulumi.Float64PtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LabOutput{})
}
