


package v20210601

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type CloudLink struct {
	pulumi.CustomResourceState

	LinkedCloud pulumi.StringPtrOutput `pulumi:"linkedCloud"`
	Name        pulumi.StringOutput    `pulumi:"name"`
	Status      pulumi.StringOutput    `pulumi:"status"`
	Type        pulumi.StringOutput    `pulumi:"type"`
}


func NewCloudLink(ctx *pulumi.Context,
	name string, args *CloudLinkArgs, opts ...pulumi.ResourceOption) (*CloudLink, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.PrivateCloudName == nil {
		return nil, errors.New("invalid value for required argument 'PrivateCloudName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:avs:CloudLink"),
		},
		{
			Type: pulumi.String("azure-native:avs/v20211201:CloudLink"),
		},
	})
	opts = append(opts, aliases)
	var resource CloudLink
	err := ctx.RegisterResource("azure-native:avs/v20210601:CloudLink", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetCloudLink(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *CloudLinkState, opts ...pulumi.ResourceOption) (*CloudLink, error) {
	var resource CloudLink
	err := ctx.ReadResource("azure-native:avs/v20210601:CloudLink", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type cloudLinkState struct {
}

type CloudLinkState struct {
}

func (CloudLinkState) ElementType() reflect.Type {
	return reflect.TypeOf((*cloudLinkState)(nil)).Elem()
}

type cloudLinkArgs struct {
	CloudLinkName     *string `pulumi:"cloudLinkName"`
	LinkedCloud       *string `pulumi:"linkedCloud"`
	PrivateCloudName  string  `pulumi:"privateCloudName"`
	ResourceGroupName string  `pulumi:"resourceGroupName"`
}


type CloudLinkArgs struct {
	CloudLinkName     pulumi.StringPtrInput
	LinkedCloud       pulumi.StringPtrInput
	PrivateCloudName  pulumi.StringInput
	ResourceGroupName pulumi.StringInput
}

func (CloudLinkArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*cloudLinkArgs)(nil)).Elem()
}

type CloudLinkInput interface {
	pulumi.Input

	ToCloudLinkOutput() CloudLinkOutput
	ToCloudLinkOutputWithContext(ctx context.Context) CloudLinkOutput
}

func (*CloudLink) ElementType() reflect.Type {
	return reflect.TypeOf((**CloudLink)(nil)).Elem()
}

func (i *CloudLink) ToCloudLinkOutput() CloudLinkOutput {
	return i.ToCloudLinkOutputWithContext(context.Background())
}

func (i *CloudLink) ToCloudLinkOutputWithContext(ctx context.Context) CloudLinkOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CloudLinkOutput)
}

type CloudLinkOutput struct{ *pulumi.OutputState }

func (CloudLinkOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**CloudLink)(nil)).Elem()
}

func (o CloudLinkOutput) ToCloudLinkOutput() CloudLinkOutput {
	return o
}

func (o CloudLinkOutput) ToCloudLinkOutputWithContext(ctx context.Context) CloudLinkOutput {
	return o
}

func (o CloudLinkOutput) LinkedCloud() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CloudLink) pulumi.StringPtrOutput { return v.LinkedCloud }).(pulumi.StringPtrOutput)
}

func (o CloudLinkOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *CloudLink) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o CloudLinkOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v *CloudLink) pulumi.StringOutput { return v.Status }).(pulumi.StringOutput)
}

func (o CloudLinkOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *CloudLink) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(CloudLinkOutput{})
}
