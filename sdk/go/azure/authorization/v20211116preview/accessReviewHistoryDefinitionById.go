


package v20211116preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type AccessReviewHistoryDefinitionById struct {
	pulumi.CustomResourceState

	CreatedDateTime                  pulumi.StringOutput                            `pulumi:"createdDateTime"`
	Decisions                        pulumi.StringArrayOutput                       `pulumi:"decisions"`
	DisplayName                      pulumi.StringPtrOutput                         `pulumi:"displayName"`
	EndDate                          pulumi.StringPtrOutput                         `pulumi:"endDate"`
	Instances                        AccessReviewHistoryInstanceResponseArrayOutput `pulumi:"instances"`
	Interval                         pulumi.IntPtrOutput                            `pulumi:"interval"`
	Name                             pulumi.StringOutput                            `pulumi:"name"`
	NumberOfOccurrences              pulumi.IntPtrOutput                            `pulumi:"numberOfOccurrences"`
	PrincipalId                      pulumi.StringOutput                            `pulumi:"principalId"`
	PrincipalName                    pulumi.StringOutput                            `pulumi:"principalName"`
	PrincipalType                    pulumi.StringOutput                            `pulumi:"principalType"`
	ReviewHistoryPeriodEndDateTime   pulumi.StringOutput                            `pulumi:"reviewHistoryPeriodEndDateTime"`
	ReviewHistoryPeriodStartDateTime pulumi.StringOutput                            `pulumi:"reviewHistoryPeriodStartDateTime"`
	Scopes                           AccessReviewScopeResponseArrayOutput           `pulumi:"scopes"`
	StartDate                        pulumi.StringPtrOutput                         `pulumi:"startDate"`
	Status                           pulumi.StringOutput                            `pulumi:"status"`
	Type                             pulumi.StringOutput                            `pulumi:"type"`
	UserPrincipalName                pulumi.StringOutput                            `pulumi:"userPrincipalName"`
}


func NewAccessReviewHistoryDefinitionById(ctx *pulumi.Context,
	name string, args *AccessReviewHistoryDefinitionByIdArgs, opts ...pulumi.ResourceOption) (*AccessReviewHistoryDefinitionById, error) {
	if args == nil {
		args = &AccessReviewHistoryDefinitionByIdArgs{}
	}

	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:authorization:AccessReviewHistoryDefinitionById"),
		},
		{
			Type: pulumi.String("azure-native:authorization/v20211201preview:AccessReviewHistoryDefinitionById"),
		},
	})
	opts = append(opts, aliases)
	var resource AccessReviewHistoryDefinitionById
	err := ctx.RegisterResource("azure-native:authorization/v20211116preview:AccessReviewHistoryDefinitionById", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetAccessReviewHistoryDefinitionById(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AccessReviewHistoryDefinitionByIdState, opts ...pulumi.ResourceOption) (*AccessReviewHistoryDefinitionById, error) {
	var resource AccessReviewHistoryDefinitionById
	err := ctx.ReadResource("azure-native:authorization/v20211116preview:AccessReviewHistoryDefinitionById", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type accessReviewHistoryDefinitionByIdState struct {
}

type AccessReviewHistoryDefinitionByIdState struct {
}

func (AccessReviewHistoryDefinitionByIdState) ElementType() reflect.Type {
	return reflect.TypeOf((*accessReviewHistoryDefinitionByIdState)(nil)).Elem()
}

type accessReviewHistoryDefinitionByIdArgs struct {
	Decisions           []string                      `pulumi:"decisions"`
	DisplayName         *string                       `pulumi:"displayName"`
	EndDate             *string                       `pulumi:"endDate"`
	HistoryDefinitionId *string                       `pulumi:"historyDefinitionId"`
	Instances           []AccessReviewHistoryInstance `pulumi:"instances"`
	Interval            *int                          `pulumi:"interval"`
	NumberOfOccurrences *int                          `pulumi:"numberOfOccurrences"`
	Scopes              []AccessReviewScope           `pulumi:"scopes"`
	StartDate           *string                       `pulumi:"startDate"`
	Type                *string                       `pulumi:"type"`
}


type AccessReviewHistoryDefinitionByIdArgs struct {
	Decisions           pulumi.StringArrayInput
	DisplayName         pulumi.StringPtrInput
	EndDate             pulumi.StringPtrInput
	HistoryDefinitionId pulumi.StringPtrInput
	Instances           AccessReviewHistoryInstanceArrayInput
	Interval            pulumi.IntPtrInput
	NumberOfOccurrences pulumi.IntPtrInput
	Scopes              AccessReviewScopeArrayInput
	StartDate           pulumi.StringPtrInput
	Type                pulumi.StringPtrInput
}

func (AccessReviewHistoryDefinitionByIdArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*accessReviewHistoryDefinitionByIdArgs)(nil)).Elem()
}

type AccessReviewHistoryDefinitionByIdInput interface {
	pulumi.Input

	ToAccessReviewHistoryDefinitionByIdOutput() AccessReviewHistoryDefinitionByIdOutput
	ToAccessReviewHistoryDefinitionByIdOutputWithContext(ctx context.Context) AccessReviewHistoryDefinitionByIdOutput
}

func (*AccessReviewHistoryDefinitionById) ElementType() reflect.Type {
	return reflect.TypeOf((**AccessReviewHistoryDefinitionById)(nil)).Elem()
}

func (i *AccessReviewHistoryDefinitionById) ToAccessReviewHistoryDefinitionByIdOutput() AccessReviewHistoryDefinitionByIdOutput {
	return i.ToAccessReviewHistoryDefinitionByIdOutputWithContext(context.Background())
}

func (i *AccessReviewHistoryDefinitionById) ToAccessReviewHistoryDefinitionByIdOutputWithContext(ctx context.Context) AccessReviewHistoryDefinitionByIdOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AccessReviewHistoryDefinitionByIdOutput)
}

type AccessReviewHistoryDefinitionByIdOutput struct{ *pulumi.OutputState }

func (AccessReviewHistoryDefinitionByIdOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AccessReviewHistoryDefinitionById)(nil)).Elem()
}

func (o AccessReviewHistoryDefinitionByIdOutput) ToAccessReviewHistoryDefinitionByIdOutput() AccessReviewHistoryDefinitionByIdOutput {
	return o
}

func (o AccessReviewHistoryDefinitionByIdOutput) ToAccessReviewHistoryDefinitionByIdOutputWithContext(ctx context.Context) AccessReviewHistoryDefinitionByIdOutput {
	return o
}

func (o AccessReviewHistoryDefinitionByIdOutput) CreatedDateTime() pulumi.StringOutput {
	return o.ApplyT(func(v *AccessReviewHistoryDefinitionById) pulumi.StringOutput { return v.CreatedDateTime }).(pulumi.StringOutput)
}

func (o AccessReviewHistoryDefinitionByIdOutput) Decisions() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *AccessReviewHistoryDefinitionById) pulumi.StringArrayOutput { return v.Decisions }).(pulumi.StringArrayOutput)
}

func (o AccessReviewHistoryDefinitionByIdOutput) DisplayName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AccessReviewHistoryDefinitionById) pulumi.StringPtrOutput { return v.DisplayName }).(pulumi.StringPtrOutput)
}

func (o AccessReviewHistoryDefinitionByIdOutput) EndDate() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AccessReviewHistoryDefinitionById) pulumi.StringPtrOutput { return v.EndDate }).(pulumi.StringPtrOutput)
}

func (o AccessReviewHistoryDefinitionByIdOutput) Instances() AccessReviewHistoryInstanceResponseArrayOutput {
	return o.ApplyT(func(v *AccessReviewHistoryDefinitionById) AccessReviewHistoryInstanceResponseArrayOutput {
		return v.Instances
	}).(AccessReviewHistoryInstanceResponseArrayOutput)
}

func (o AccessReviewHistoryDefinitionByIdOutput) Interval() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *AccessReviewHistoryDefinitionById) pulumi.IntPtrOutput { return v.Interval }).(pulumi.IntPtrOutput)
}

func (o AccessReviewHistoryDefinitionByIdOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *AccessReviewHistoryDefinitionById) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o AccessReviewHistoryDefinitionByIdOutput) NumberOfOccurrences() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *AccessReviewHistoryDefinitionById) pulumi.IntPtrOutput { return v.NumberOfOccurrences }).(pulumi.IntPtrOutput)
}

func (o AccessReviewHistoryDefinitionByIdOutput) PrincipalId() pulumi.StringOutput {
	return o.ApplyT(func(v *AccessReviewHistoryDefinitionById) pulumi.StringOutput { return v.PrincipalId }).(pulumi.StringOutput)
}

func (o AccessReviewHistoryDefinitionByIdOutput) PrincipalName() pulumi.StringOutput {
	return o.ApplyT(func(v *AccessReviewHistoryDefinitionById) pulumi.StringOutput { return v.PrincipalName }).(pulumi.StringOutput)
}

func (o AccessReviewHistoryDefinitionByIdOutput) PrincipalType() pulumi.StringOutput {
	return o.ApplyT(func(v *AccessReviewHistoryDefinitionById) pulumi.StringOutput { return v.PrincipalType }).(pulumi.StringOutput)
}

func (o AccessReviewHistoryDefinitionByIdOutput) ReviewHistoryPeriodEndDateTime() pulumi.StringOutput {
	return o.ApplyT(func(v *AccessReviewHistoryDefinitionById) pulumi.StringOutput {
		return v.ReviewHistoryPeriodEndDateTime
	}).(pulumi.StringOutput)
}

func (o AccessReviewHistoryDefinitionByIdOutput) ReviewHistoryPeriodStartDateTime() pulumi.StringOutput {
	return o.ApplyT(func(v *AccessReviewHistoryDefinitionById) pulumi.StringOutput {
		return v.ReviewHistoryPeriodStartDateTime
	}).(pulumi.StringOutput)
}

func (o AccessReviewHistoryDefinitionByIdOutput) Scopes() AccessReviewScopeResponseArrayOutput {
	return o.ApplyT(func(v *AccessReviewHistoryDefinitionById) AccessReviewScopeResponseArrayOutput { return v.Scopes }).(AccessReviewScopeResponseArrayOutput)
}

func (o AccessReviewHistoryDefinitionByIdOutput) StartDate() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AccessReviewHistoryDefinitionById) pulumi.StringPtrOutput { return v.StartDate }).(pulumi.StringPtrOutput)
}

func (o AccessReviewHistoryDefinitionByIdOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v *AccessReviewHistoryDefinitionById) pulumi.StringOutput { return v.Status }).(pulumi.StringOutput)
}

func (o AccessReviewHistoryDefinitionByIdOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *AccessReviewHistoryDefinitionById) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func (o AccessReviewHistoryDefinitionByIdOutput) UserPrincipalName() pulumi.StringOutput {
	return o.ApplyT(func(v *AccessReviewHistoryDefinitionById) pulumi.StringOutput { return v.UserPrincipalName }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(AccessReviewHistoryDefinitionByIdOutput{})
}
