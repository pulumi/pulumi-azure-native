


package v20211201preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type Student struct {
	pulumi.CustomResourceState

	Budget                         AmountResponseOutput     `pulumi:"budget"`
	EffectiveDate                  pulumi.StringOutput      `pulumi:"effectiveDate"`
	Email                          pulumi.StringOutput      `pulumi:"email"`
	ExpirationDate                 pulumi.StringOutput      `pulumi:"expirationDate"`
	FirstName                      pulumi.StringOutput      `pulumi:"firstName"`
	LastName                       pulumi.StringOutput      `pulumi:"lastName"`
	Name                           pulumi.StringOutput      `pulumi:"name"`
	Role                           pulumi.StringOutput      `pulumi:"role"`
	Status                         pulumi.StringOutput      `pulumi:"status"`
	SubscriptionAlias              pulumi.StringPtrOutput   `pulumi:"subscriptionAlias"`
	SubscriptionId                 pulumi.StringOutput      `pulumi:"subscriptionId"`
	SubscriptionInviteLastSentDate pulumi.StringPtrOutput   `pulumi:"subscriptionInviteLastSentDate"`
	SystemData                     SystemDataResponseOutput `pulumi:"systemData"`
	Type                           pulumi.StringOutput      `pulumi:"type"`
}


func NewStudent(ctx *pulumi.Context,
	name string, args *StudentArgs, opts ...pulumi.ResourceOption) (*Student, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.BillingAccountName == nil {
		return nil, errors.New("invalid value for required argument 'BillingAccountName'")
	}
	if args.BillingProfileName == nil {
		return nil, errors.New("invalid value for required argument 'BillingProfileName'")
	}
	if args.Budget == nil {
		return nil, errors.New("invalid value for required argument 'Budget'")
	}
	if args.Email == nil {
		return nil, errors.New("invalid value for required argument 'Email'")
	}
	if args.ExpirationDate == nil {
		return nil, errors.New("invalid value for required argument 'ExpirationDate'")
	}
	if args.FirstName == nil {
		return nil, errors.New("invalid value for required argument 'FirstName'")
	}
	if args.InvoiceSectionName == nil {
		return nil, errors.New("invalid value for required argument 'InvoiceSectionName'")
	}
	if args.LastName == nil {
		return nil, errors.New("invalid value for required argument 'LastName'")
	}
	if args.Role == nil {
		return nil, errors.New("invalid value for required argument 'Role'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:education:Student"),
		},
	})
	opts = append(opts, aliases)
	var resource Student
	err := ctx.RegisterResource("azure-native:education/v20211201preview:Student", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetStudent(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *StudentState, opts ...pulumi.ResourceOption) (*Student, error) {
	var resource Student
	err := ctx.ReadResource("azure-native:education/v20211201preview:Student", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type studentState struct {
}

type StudentState struct {
}

func (StudentState) ElementType() reflect.Type {
	return reflect.TypeOf((*studentState)(nil)).Elem()
}

type studentArgs struct {
	BillingAccountName             string  `pulumi:"billingAccountName"`
	BillingProfileName             string  `pulumi:"billingProfileName"`
	Budget                         Amount  `pulumi:"budget"`
	Email                          string  `pulumi:"email"`
	ExpirationDate                 string  `pulumi:"expirationDate"`
	FirstName                      string  `pulumi:"firstName"`
	InvoiceSectionName             string  `pulumi:"invoiceSectionName"`
	LastName                       string  `pulumi:"lastName"`
	Role                           string  `pulumi:"role"`
	StudentAlias                   *string `pulumi:"studentAlias"`
	SubscriptionAlias              *string `pulumi:"subscriptionAlias"`
	SubscriptionInviteLastSentDate *string `pulumi:"subscriptionInviteLastSentDate"`
}


type StudentArgs struct {
	BillingAccountName             pulumi.StringInput
	BillingProfileName             pulumi.StringInput
	Budget                         AmountInput
	Email                          pulumi.StringInput
	ExpirationDate                 pulumi.StringInput
	FirstName                      pulumi.StringInput
	InvoiceSectionName             pulumi.StringInput
	LastName                       pulumi.StringInput
	Role                           pulumi.StringInput
	StudentAlias                   pulumi.StringPtrInput
	SubscriptionAlias              pulumi.StringPtrInput
	SubscriptionInviteLastSentDate pulumi.StringPtrInput
}

func (StudentArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*studentArgs)(nil)).Elem()
}

type StudentInput interface {
	pulumi.Input

	ToStudentOutput() StudentOutput
	ToStudentOutputWithContext(ctx context.Context) StudentOutput
}

func (*Student) ElementType() reflect.Type {
	return reflect.TypeOf((**Student)(nil)).Elem()
}

func (i *Student) ToStudentOutput() StudentOutput {
	return i.ToStudentOutputWithContext(context.Background())
}

func (i *Student) ToStudentOutputWithContext(ctx context.Context) StudentOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StudentOutput)
}

type StudentOutput struct{ *pulumi.OutputState }

func (StudentOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Student)(nil)).Elem()
}

func (o StudentOutput) ToStudentOutput() StudentOutput {
	return o
}

func (o StudentOutput) ToStudentOutputWithContext(ctx context.Context) StudentOutput {
	return o
}

func (o StudentOutput) Budget() AmountResponseOutput {
	return o.ApplyT(func(v *Student) AmountResponseOutput { return v.Budget }).(AmountResponseOutput)
}

func (o StudentOutput) EffectiveDate() pulumi.StringOutput {
	return o.ApplyT(func(v *Student) pulumi.StringOutput { return v.EffectiveDate }).(pulumi.StringOutput)
}

func (o StudentOutput) Email() pulumi.StringOutput {
	return o.ApplyT(func(v *Student) pulumi.StringOutput { return v.Email }).(pulumi.StringOutput)
}

func (o StudentOutput) ExpirationDate() pulumi.StringOutput {
	return o.ApplyT(func(v *Student) pulumi.StringOutput { return v.ExpirationDate }).(pulumi.StringOutput)
}

func (o StudentOutput) FirstName() pulumi.StringOutput {
	return o.ApplyT(func(v *Student) pulumi.StringOutput { return v.FirstName }).(pulumi.StringOutput)
}

func (o StudentOutput) LastName() pulumi.StringOutput {
	return o.ApplyT(func(v *Student) pulumi.StringOutput { return v.LastName }).(pulumi.StringOutput)
}

func (o StudentOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Student) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o StudentOutput) Role() pulumi.StringOutput {
	return o.ApplyT(func(v *Student) pulumi.StringOutput { return v.Role }).(pulumi.StringOutput)
}

func (o StudentOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v *Student) pulumi.StringOutput { return v.Status }).(pulumi.StringOutput)
}

func (o StudentOutput) SubscriptionAlias() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Student) pulumi.StringPtrOutput { return v.SubscriptionAlias }).(pulumi.StringPtrOutput)
}

func (o StudentOutput) SubscriptionId() pulumi.StringOutput {
	return o.ApplyT(func(v *Student) pulumi.StringOutput { return v.SubscriptionId }).(pulumi.StringOutput)
}

func (o StudentOutput) SubscriptionInviteLastSentDate() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Student) pulumi.StringPtrOutput { return v.SubscriptionInviteLastSentDate }).(pulumi.StringPtrOutput)
}

func (o StudentOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *Student) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

func (o StudentOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *Student) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(StudentOutput{})
}
