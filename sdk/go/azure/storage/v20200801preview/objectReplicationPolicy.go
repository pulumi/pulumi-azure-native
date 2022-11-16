


package v20200801preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)


type ObjectReplicationPolicy struct {
	pulumi.CustomResourceState

	DestinationAccount pulumi.StringOutput                            `pulumi:"destinationAccount"`
	EnabledTime        pulumi.StringOutput                            `pulumi:"enabledTime"`
	Name               pulumi.StringOutput                            `pulumi:"name"`
	PolicyId           pulumi.StringOutput                            `pulumi:"policyId"`
	Rules              ObjectReplicationPolicyRuleResponseArrayOutput `pulumi:"rules"`
	SourceAccount      pulumi.StringOutput                            `pulumi:"sourceAccount"`
	Type               pulumi.StringOutput                            `pulumi:"type"`
}


func NewObjectReplicationPolicy(ctx *pulumi.Context,
	name string, args *ObjectReplicationPolicyArgs, opts ...pulumi.ResourceOption) (*ObjectReplicationPolicy, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AccountName == nil {
		return nil, errors.New("invalid value for required argument 'AccountName'")
	}
	if args.DestinationAccount == nil {
		return nil, errors.New("invalid value for required argument 'DestinationAccount'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.SourceAccount == nil {
		return nil, errors.New("invalid value for required argument 'SourceAccount'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:storage:ObjectReplicationPolicy"),
		},
		{
			Type: pulumi.String("azure-native:storage/v20190601:ObjectReplicationPolicy"),
		},
		{
			Type: pulumi.String("azure-native:storage/v20210101:ObjectReplicationPolicy"),
		},
		{
			Type: pulumi.String("azure-native:storage/v20210201:ObjectReplicationPolicy"),
		},
		{
			Type: pulumi.String("azure-native:storage/v20210401:ObjectReplicationPolicy"),
		},
		{
			Type: pulumi.String("azure-native:storage/v20210601:ObjectReplicationPolicy"),
		},
		{
			Type: pulumi.String("azure-native:storage/v20210801:ObjectReplicationPolicy"),
		},
		{
			Type: pulumi.String("azure-native:storage/v20210901:ObjectReplicationPolicy"),
		},
		{
			Type: pulumi.String("azure-native:storage/v20220501:ObjectReplicationPolicy"),
		},
		{
			Type: pulumi.String("azure-native:storage/v20220901:ObjectReplicationPolicy"),
		},
	})
	opts = append(opts, aliases)
	var resource ObjectReplicationPolicy
	err := ctx.RegisterResource("azure-native:storage/v20200801preview:ObjectReplicationPolicy", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetObjectReplicationPolicy(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ObjectReplicationPolicyState, opts ...pulumi.ResourceOption) (*ObjectReplicationPolicy, error) {
	var resource ObjectReplicationPolicy
	err := ctx.ReadResource("azure-native:storage/v20200801preview:ObjectReplicationPolicy", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type objectReplicationPolicyState struct {
}

type ObjectReplicationPolicyState struct {
}

func (ObjectReplicationPolicyState) ElementType() reflect.Type {
	return reflect.TypeOf((*objectReplicationPolicyState)(nil)).Elem()
}

type objectReplicationPolicyArgs struct {
	AccountName               string                        `pulumi:"accountName"`
	DestinationAccount        string                        `pulumi:"destinationAccount"`
	ObjectReplicationPolicyId *string                       `pulumi:"objectReplicationPolicyId"`
	ResourceGroupName         string                        `pulumi:"resourceGroupName"`
	Rules                     []ObjectReplicationPolicyRule `pulumi:"rules"`
	SourceAccount             string                        `pulumi:"sourceAccount"`
}


type ObjectReplicationPolicyArgs struct {
	AccountName               pulumi.StringInput
	DestinationAccount        pulumi.StringInput
	ObjectReplicationPolicyId pulumi.StringPtrInput
	ResourceGroupName         pulumi.StringInput
	Rules                     ObjectReplicationPolicyRuleArrayInput
	SourceAccount             pulumi.StringInput
}

func (ObjectReplicationPolicyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*objectReplicationPolicyArgs)(nil)).Elem()
}

type ObjectReplicationPolicyInput interface {
	pulumi.Input

	ToObjectReplicationPolicyOutput() ObjectReplicationPolicyOutput
	ToObjectReplicationPolicyOutputWithContext(ctx context.Context) ObjectReplicationPolicyOutput
}

func (*ObjectReplicationPolicy) ElementType() reflect.Type {
	return reflect.TypeOf((**ObjectReplicationPolicy)(nil)).Elem()
}

func (i *ObjectReplicationPolicy) ToObjectReplicationPolicyOutput() ObjectReplicationPolicyOutput {
	return i.ToObjectReplicationPolicyOutputWithContext(context.Background())
}

func (i *ObjectReplicationPolicy) ToObjectReplicationPolicyOutputWithContext(ctx context.Context) ObjectReplicationPolicyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ObjectReplicationPolicyOutput)
}

type ObjectReplicationPolicyOutput struct{ *pulumi.OutputState }

func (ObjectReplicationPolicyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ObjectReplicationPolicy)(nil)).Elem()
}

func (o ObjectReplicationPolicyOutput) ToObjectReplicationPolicyOutput() ObjectReplicationPolicyOutput {
	return o
}

func (o ObjectReplicationPolicyOutput) ToObjectReplicationPolicyOutputWithContext(ctx context.Context) ObjectReplicationPolicyOutput {
	return o
}

func (o ObjectReplicationPolicyOutput) DestinationAccount() pulumi.StringOutput {
	return o.ApplyT(func(v *ObjectReplicationPolicy) pulumi.StringOutput { return v.DestinationAccount }).(pulumi.StringOutput)
}

func (o ObjectReplicationPolicyOutput) EnabledTime() pulumi.StringOutput {
	return o.ApplyT(func(v *ObjectReplicationPolicy) pulumi.StringOutput { return v.EnabledTime }).(pulumi.StringOutput)
}

func (o ObjectReplicationPolicyOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *ObjectReplicationPolicy) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o ObjectReplicationPolicyOutput) PolicyId() pulumi.StringOutput {
	return o.ApplyT(func(v *ObjectReplicationPolicy) pulumi.StringOutput { return v.PolicyId }).(pulumi.StringOutput)
}

func (o ObjectReplicationPolicyOutput) Rules() ObjectReplicationPolicyRuleResponseArrayOutput {
	return o.ApplyT(func(v *ObjectReplicationPolicy) ObjectReplicationPolicyRuleResponseArrayOutput { return v.Rules }).(ObjectReplicationPolicyRuleResponseArrayOutput)
}

func (o ObjectReplicationPolicyOutput) SourceAccount() pulumi.StringOutput {
	return o.ApplyT(func(v *ObjectReplicationPolicy) pulumi.StringOutput { return v.SourceAccount }).(pulumi.StringOutput)
}

func (o ObjectReplicationPolicyOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *ObjectReplicationPolicy) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(ObjectReplicationPolicyOutput{})
}
