


package v20181015

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type Lab struct {
	pulumi.CustomResourceState

	CreatedByObjectId          pulumi.StringOutput                 `pulumi:"createdByObjectId"`
	CreatedByUserPrincipalName pulumi.StringOutput                 `pulumi:"createdByUserPrincipalName"`
	CreatedDate                pulumi.StringOutput                 `pulumi:"createdDate"`
	InvitationCode             pulumi.StringOutput                 `pulumi:"invitationCode"`
	LatestOperationResult      LatestOperationResultResponseOutput `pulumi:"latestOperationResult"`
	Location                   pulumi.StringPtrOutput              `pulumi:"location"`
	MaxUsersInLab              pulumi.IntPtrOutput                 `pulumi:"maxUsersInLab"`
	Name                       pulumi.StringOutput                 `pulumi:"name"`
	ProvisioningState          pulumi.StringPtrOutput              `pulumi:"provisioningState"`
	Tags                       pulumi.StringMapOutput              `pulumi:"tags"`
	Type                       pulumi.StringOutput                 `pulumi:"type"`
	UniqueIdentifier           pulumi.StringPtrOutput              `pulumi:"uniqueIdentifier"`
	UsageQuota                 pulumi.StringPtrOutput              `pulumi:"usageQuota"`
	UserAccessMode             pulumi.StringPtrOutput              `pulumi:"userAccessMode"`
	UserQuota                  pulumi.IntOutput                    `pulumi:"userQuota"`
}


func NewLab(ctx *pulumi.Context,
	name string, args *LabArgs, opts ...pulumi.ResourceOption) (*Lab, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.LabAccountName == nil {
		return nil, errors.New("invalid value for required argument 'LabAccountName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:labservices:Lab"),
		},
	})
	opts = append(opts, aliases)
	var resource Lab
	err := ctx.RegisterResource("azure-native:labservices/v20181015:Lab", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetLab(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *LabState, opts ...pulumi.ResourceOption) (*Lab, error) {
	var resource Lab
	err := ctx.ReadResource("azure-native:labservices/v20181015:Lab", name, id, state, &resource, opts...)
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
	LabAccountName    string            `pulumi:"labAccountName"`
	LabName           *string           `pulumi:"labName"`
	Location          *string           `pulumi:"location"`
	MaxUsersInLab     *int              `pulumi:"maxUsersInLab"`
	ProvisioningState *string           `pulumi:"provisioningState"`
	ResourceGroupName string            `pulumi:"resourceGroupName"`
	Tags              map[string]string `pulumi:"tags"`
	UniqueIdentifier  *string           `pulumi:"uniqueIdentifier"`
	UsageQuota        *string           `pulumi:"usageQuota"`
	UserAccessMode    *string           `pulumi:"userAccessMode"`
}


type LabArgs struct {
	LabAccountName    pulumi.StringInput
	LabName           pulumi.StringPtrInput
	Location          pulumi.StringPtrInput
	MaxUsersInLab     pulumi.IntPtrInput
	ProvisioningState pulumi.StringPtrInput
	ResourceGroupName pulumi.StringInput
	Tags              pulumi.StringMapInput
	UniqueIdentifier  pulumi.StringPtrInput
	UsageQuota        pulumi.StringPtrInput
	UserAccessMode    pulumi.StringPtrInput
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

func (o LabOutput) CreatedByObjectId() pulumi.StringOutput {
	return o.ApplyT(func(v *Lab) pulumi.StringOutput { return v.CreatedByObjectId }).(pulumi.StringOutput)
}

func (o LabOutput) CreatedByUserPrincipalName() pulumi.StringOutput {
	return o.ApplyT(func(v *Lab) pulumi.StringOutput { return v.CreatedByUserPrincipalName }).(pulumi.StringOutput)
}

func (o LabOutput) CreatedDate() pulumi.StringOutput {
	return o.ApplyT(func(v *Lab) pulumi.StringOutput { return v.CreatedDate }).(pulumi.StringOutput)
}

func (o LabOutput) InvitationCode() pulumi.StringOutput {
	return o.ApplyT(func(v *Lab) pulumi.StringOutput { return v.InvitationCode }).(pulumi.StringOutput)
}

func (o LabOutput) LatestOperationResult() LatestOperationResultResponseOutput {
	return o.ApplyT(func(v *Lab) LatestOperationResultResponseOutput { return v.LatestOperationResult }).(LatestOperationResultResponseOutput)
}

func (o LabOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Lab) pulumi.StringPtrOutput { return v.Location }).(pulumi.StringPtrOutput)
}

func (o LabOutput) MaxUsersInLab() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *Lab) pulumi.IntPtrOutput { return v.MaxUsersInLab }).(pulumi.IntPtrOutput)
}

func (o LabOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Lab) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o LabOutput) ProvisioningState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Lab) pulumi.StringPtrOutput { return v.ProvisioningState }).(pulumi.StringPtrOutput)
}

func (o LabOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *Lab) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

func (o LabOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *Lab) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func (o LabOutput) UniqueIdentifier() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Lab) pulumi.StringPtrOutput { return v.UniqueIdentifier }).(pulumi.StringPtrOutput)
}

func (o LabOutput) UsageQuota() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Lab) pulumi.StringPtrOutput { return v.UsageQuota }).(pulumi.StringPtrOutput)
}

func (o LabOutput) UserAccessMode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Lab) pulumi.StringPtrOutput { return v.UserAccessMode }).(pulumi.StringPtrOutput)
}

func (o LabOutput) UserQuota() pulumi.IntOutput {
	return o.ApplyT(func(v *Lab) pulumi.IntOutput { return v.UserQuota }).(pulumi.IntOutput)
}

func init() {
	pulumi.RegisterOutputType(LabOutput{})
}
