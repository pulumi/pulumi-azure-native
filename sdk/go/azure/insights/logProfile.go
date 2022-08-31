


package insights

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type LogProfile struct {
	pulumi.CustomResourceState

	Categories       pulumi.StringArrayOutput      `pulumi:"categories"`
	Location         pulumi.StringOutput           `pulumi:"location"`
	Locations        pulumi.StringArrayOutput      `pulumi:"locations"`
	Name             pulumi.StringOutput           `pulumi:"name"`
	RetentionPolicy  RetentionPolicyResponseOutput `pulumi:"retentionPolicy"`
	ServiceBusRuleId pulumi.StringPtrOutput        `pulumi:"serviceBusRuleId"`
	StorageAccountId pulumi.StringPtrOutput        `pulumi:"storageAccountId"`
	Tags             pulumi.StringMapOutput        `pulumi:"tags"`
	Type             pulumi.StringOutput           `pulumi:"type"`
}


func NewLogProfile(ctx *pulumi.Context,
	name string, args *LogProfileArgs, opts ...pulumi.ResourceOption) (*LogProfile, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Categories == nil {
		return nil, errors.New("invalid value for required argument 'Categories'")
	}
	if args.Locations == nil {
		return nil, errors.New("invalid value for required argument 'Locations'")
	}
	if args.RetentionPolicy == nil {
		return nil, errors.New("invalid value for required argument 'RetentionPolicy'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:insights/v20160301:LogProfile"),
		},
	})
	opts = append(opts, aliases)
	var resource LogProfile
	err := ctx.RegisterResource("azure-native:insights:LogProfile", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetLogProfile(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *LogProfileState, opts ...pulumi.ResourceOption) (*LogProfile, error) {
	var resource LogProfile
	err := ctx.ReadResource("azure-native:insights:LogProfile", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type logProfileState struct {
}

type LogProfileState struct {
}

func (LogProfileState) ElementType() reflect.Type {
	return reflect.TypeOf((*logProfileState)(nil)).Elem()
}

type logProfileArgs struct {
	Categories       []string          `pulumi:"categories"`
	Location         *string           `pulumi:"location"`
	Locations        []string          `pulumi:"locations"`
	LogProfileName   *string           `pulumi:"logProfileName"`
	RetentionPolicy  RetentionPolicy   `pulumi:"retentionPolicy"`
	ServiceBusRuleId *string           `pulumi:"serviceBusRuleId"`
	StorageAccountId *string           `pulumi:"storageAccountId"`
	Tags             map[string]string `pulumi:"tags"`
}


type LogProfileArgs struct {
	Categories       pulumi.StringArrayInput
	Location         pulumi.StringPtrInput
	Locations        pulumi.StringArrayInput
	LogProfileName   pulumi.StringPtrInput
	RetentionPolicy  RetentionPolicyInput
	ServiceBusRuleId pulumi.StringPtrInput
	StorageAccountId pulumi.StringPtrInput
	Tags             pulumi.StringMapInput
}

func (LogProfileArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*logProfileArgs)(nil)).Elem()
}

type LogProfileInput interface {
	pulumi.Input

	ToLogProfileOutput() LogProfileOutput
	ToLogProfileOutputWithContext(ctx context.Context) LogProfileOutput
}

func (*LogProfile) ElementType() reflect.Type {
	return reflect.TypeOf((**LogProfile)(nil)).Elem()
}

func (i *LogProfile) ToLogProfileOutput() LogProfileOutput {
	return i.ToLogProfileOutputWithContext(context.Background())
}

func (i *LogProfile) ToLogProfileOutputWithContext(ctx context.Context) LogProfileOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LogProfileOutput)
}

type LogProfileOutput struct{ *pulumi.OutputState }

func (LogProfileOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**LogProfile)(nil)).Elem()
}

func (o LogProfileOutput) ToLogProfileOutput() LogProfileOutput {
	return o
}

func (o LogProfileOutput) ToLogProfileOutputWithContext(ctx context.Context) LogProfileOutput {
	return o
}

func (o LogProfileOutput) Categories() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *LogProfile) pulumi.StringArrayOutput { return v.Categories }).(pulumi.StringArrayOutput)
}

func (o LogProfileOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *LogProfile) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

func (o LogProfileOutput) Locations() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *LogProfile) pulumi.StringArrayOutput { return v.Locations }).(pulumi.StringArrayOutput)
}

func (o LogProfileOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *LogProfile) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o LogProfileOutput) RetentionPolicy() RetentionPolicyResponseOutput {
	return o.ApplyT(func(v *LogProfile) RetentionPolicyResponseOutput { return v.RetentionPolicy }).(RetentionPolicyResponseOutput)
}

func (o LogProfileOutput) ServiceBusRuleId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *LogProfile) pulumi.StringPtrOutput { return v.ServiceBusRuleId }).(pulumi.StringPtrOutput)
}

func (o LogProfileOutput) StorageAccountId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *LogProfile) pulumi.StringPtrOutput { return v.StorageAccountId }).(pulumi.StringPtrOutput)
}

func (o LogProfileOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *LogProfile) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

func (o LogProfileOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *LogProfile) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LogProfileOutput{})
}
