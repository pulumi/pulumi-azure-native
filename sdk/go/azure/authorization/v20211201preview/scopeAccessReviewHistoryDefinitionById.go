


package v20211201preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type ScopeAccessReviewHistoryDefinitionById struct {
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


func NewScopeAccessReviewHistoryDefinitionById(ctx *pulumi.Context,
	name string, args *ScopeAccessReviewHistoryDefinitionByIdArgs, opts ...pulumi.ResourceOption) (*ScopeAccessReviewHistoryDefinitionById, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Scope == nil {
		return nil, errors.New("invalid value for required argument 'Scope'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:authorization:ScopeAccessReviewHistoryDefinitionById"),
		},
	})
	opts = append(opts, aliases)
	var resource ScopeAccessReviewHistoryDefinitionById
	err := ctx.RegisterResource("azure-native:authorization/v20211201preview:ScopeAccessReviewHistoryDefinitionById", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetScopeAccessReviewHistoryDefinitionById(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ScopeAccessReviewHistoryDefinitionByIdState, opts ...pulumi.ResourceOption) (*ScopeAccessReviewHistoryDefinitionById, error) {
	var resource ScopeAccessReviewHistoryDefinitionById
	err := ctx.ReadResource("azure-native:authorization/v20211201preview:ScopeAccessReviewHistoryDefinitionById", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type scopeAccessReviewHistoryDefinitionByIdState struct {
}

type ScopeAccessReviewHistoryDefinitionByIdState struct {
}

func (ScopeAccessReviewHistoryDefinitionByIdState) ElementType() reflect.Type {
	return reflect.TypeOf((*scopeAccessReviewHistoryDefinitionByIdState)(nil)).Elem()
}

type scopeAccessReviewHistoryDefinitionByIdArgs struct {
	Decisions           []string                      `pulumi:"decisions"`
	DisplayName         *string                       `pulumi:"displayName"`
	EndDate             *string                       `pulumi:"endDate"`
	HistoryDefinitionId *string                       `pulumi:"historyDefinitionId"`
	Instances           []AccessReviewHistoryInstance `pulumi:"instances"`
	Interval            *int                          `pulumi:"interval"`
	NumberOfOccurrences *int                          `pulumi:"numberOfOccurrences"`
	Scope               string                        `pulumi:"scope"`
	Scopes              []AccessReviewScope           `pulumi:"scopes"`
	StartDate           *string                       `pulumi:"startDate"`
	Type                *string                       `pulumi:"type"`
}


type ScopeAccessReviewHistoryDefinitionByIdArgs struct {
	Decisions           pulumi.StringArrayInput
	DisplayName         pulumi.StringPtrInput
	EndDate             pulumi.StringPtrInput
	HistoryDefinitionId pulumi.StringPtrInput
	Instances           AccessReviewHistoryInstanceArrayInput
	Interval            pulumi.IntPtrInput
	NumberOfOccurrences pulumi.IntPtrInput
	Scope               pulumi.StringInput
	Scopes              AccessReviewScopeArrayInput
	StartDate           pulumi.StringPtrInput
	Type                pulumi.StringPtrInput
}

func (ScopeAccessReviewHistoryDefinitionByIdArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*scopeAccessReviewHistoryDefinitionByIdArgs)(nil)).Elem()
}

type ScopeAccessReviewHistoryDefinitionByIdInput interface {
	pulumi.Input

	ToScopeAccessReviewHistoryDefinitionByIdOutput() ScopeAccessReviewHistoryDefinitionByIdOutput
	ToScopeAccessReviewHistoryDefinitionByIdOutputWithContext(ctx context.Context) ScopeAccessReviewHistoryDefinitionByIdOutput
}

func (*ScopeAccessReviewHistoryDefinitionById) ElementType() reflect.Type {
	return reflect.TypeOf((**ScopeAccessReviewHistoryDefinitionById)(nil)).Elem()
}

func (i *ScopeAccessReviewHistoryDefinitionById) ToScopeAccessReviewHistoryDefinitionByIdOutput() ScopeAccessReviewHistoryDefinitionByIdOutput {
	return i.ToScopeAccessReviewHistoryDefinitionByIdOutputWithContext(context.Background())
}

func (i *ScopeAccessReviewHistoryDefinitionById) ToScopeAccessReviewHistoryDefinitionByIdOutputWithContext(ctx context.Context) ScopeAccessReviewHistoryDefinitionByIdOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ScopeAccessReviewHistoryDefinitionByIdOutput)
}

type ScopeAccessReviewHistoryDefinitionByIdOutput struct{ *pulumi.OutputState }

func (ScopeAccessReviewHistoryDefinitionByIdOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ScopeAccessReviewHistoryDefinitionById)(nil)).Elem()
}

func (o ScopeAccessReviewHistoryDefinitionByIdOutput) ToScopeAccessReviewHistoryDefinitionByIdOutput() ScopeAccessReviewHistoryDefinitionByIdOutput {
	return o
}

func (o ScopeAccessReviewHistoryDefinitionByIdOutput) ToScopeAccessReviewHistoryDefinitionByIdOutputWithContext(ctx context.Context) ScopeAccessReviewHistoryDefinitionByIdOutput {
	return o
}

func (o ScopeAccessReviewHistoryDefinitionByIdOutput) CreatedDateTime() pulumi.StringOutput {
	return o.ApplyT(func(v *ScopeAccessReviewHistoryDefinitionById) pulumi.StringOutput { return v.CreatedDateTime }).(pulumi.StringOutput)
}

func (o ScopeAccessReviewHistoryDefinitionByIdOutput) Decisions() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *ScopeAccessReviewHistoryDefinitionById) pulumi.StringArrayOutput { return v.Decisions }).(pulumi.StringArrayOutput)
}

func (o ScopeAccessReviewHistoryDefinitionByIdOutput) DisplayName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ScopeAccessReviewHistoryDefinitionById) pulumi.StringPtrOutput { return v.DisplayName }).(pulumi.StringPtrOutput)
}

func (o ScopeAccessReviewHistoryDefinitionByIdOutput) EndDate() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ScopeAccessReviewHistoryDefinitionById) pulumi.StringPtrOutput { return v.EndDate }).(pulumi.StringPtrOutput)
}

func (o ScopeAccessReviewHistoryDefinitionByIdOutput) Instances() AccessReviewHistoryInstanceResponseArrayOutput {
	return o.ApplyT(func(v *ScopeAccessReviewHistoryDefinitionById) AccessReviewHistoryInstanceResponseArrayOutput {
		return v.Instances
	}).(AccessReviewHistoryInstanceResponseArrayOutput)
}

func (o ScopeAccessReviewHistoryDefinitionByIdOutput) Interval() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *ScopeAccessReviewHistoryDefinitionById) pulumi.IntPtrOutput { return v.Interval }).(pulumi.IntPtrOutput)
}

func (o ScopeAccessReviewHistoryDefinitionByIdOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *ScopeAccessReviewHistoryDefinitionById) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o ScopeAccessReviewHistoryDefinitionByIdOutput) NumberOfOccurrences() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *ScopeAccessReviewHistoryDefinitionById) pulumi.IntPtrOutput { return v.NumberOfOccurrences }).(pulumi.IntPtrOutput)
}

func (o ScopeAccessReviewHistoryDefinitionByIdOutput) PrincipalId() pulumi.StringOutput {
	return o.ApplyT(func(v *ScopeAccessReviewHistoryDefinitionById) pulumi.StringOutput { return v.PrincipalId }).(pulumi.StringOutput)
}

func (o ScopeAccessReviewHistoryDefinitionByIdOutput) PrincipalName() pulumi.StringOutput {
	return o.ApplyT(func(v *ScopeAccessReviewHistoryDefinitionById) pulumi.StringOutput { return v.PrincipalName }).(pulumi.StringOutput)
}

func (o ScopeAccessReviewHistoryDefinitionByIdOutput) PrincipalType() pulumi.StringOutput {
	return o.ApplyT(func(v *ScopeAccessReviewHistoryDefinitionById) pulumi.StringOutput { return v.PrincipalType }).(pulumi.StringOutput)
}

func (o ScopeAccessReviewHistoryDefinitionByIdOutput) ReviewHistoryPeriodEndDateTime() pulumi.StringOutput {
	return o.ApplyT(func(v *ScopeAccessReviewHistoryDefinitionById) pulumi.StringOutput {
		return v.ReviewHistoryPeriodEndDateTime
	}).(pulumi.StringOutput)
}

func (o ScopeAccessReviewHistoryDefinitionByIdOutput) ReviewHistoryPeriodStartDateTime() pulumi.StringOutput {
	return o.ApplyT(func(v *ScopeAccessReviewHistoryDefinitionById) pulumi.StringOutput {
		return v.ReviewHistoryPeriodStartDateTime
	}).(pulumi.StringOutput)
}

func (o ScopeAccessReviewHistoryDefinitionByIdOutput) Scopes() AccessReviewScopeResponseArrayOutput {
	return o.ApplyT(func(v *ScopeAccessReviewHistoryDefinitionById) AccessReviewScopeResponseArrayOutput { return v.Scopes }).(AccessReviewScopeResponseArrayOutput)
}

func (o ScopeAccessReviewHistoryDefinitionByIdOutput) StartDate() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ScopeAccessReviewHistoryDefinitionById) pulumi.StringPtrOutput { return v.StartDate }).(pulumi.StringPtrOutput)
}

func (o ScopeAccessReviewHistoryDefinitionByIdOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v *ScopeAccessReviewHistoryDefinitionById) pulumi.StringOutput { return v.Status }).(pulumi.StringOutput)
}

func (o ScopeAccessReviewHistoryDefinitionByIdOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *ScopeAccessReviewHistoryDefinitionById) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func (o ScopeAccessReviewHistoryDefinitionByIdOutput) UserPrincipalName() pulumi.StringOutput {
	return o.ApplyT(func(v *ScopeAccessReviewHistoryDefinitionById) pulumi.StringOutput { return v.UserPrincipalName }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(ScopeAccessReviewHistoryDefinitionByIdOutput{})
}
