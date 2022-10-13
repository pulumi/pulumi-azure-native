


package v20210101preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type HcxEnterpriseSite struct {
	pulumi.CustomResourceState

	ActivationKey pulumi.StringOutput `pulumi:"activationKey"`
	Name          pulumi.StringOutput `pulumi:"name"`
	Status        pulumi.StringOutput `pulumi:"status"`
	Type          pulumi.StringOutput `pulumi:"type"`
}


func NewHcxEnterpriseSite(ctx *pulumi.Context,
	name string, args *HcxEnterpriseSiteArgs, opts ...pulumi.ResourceOption) (*HcxEnterpriseSite, error) {
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
			Type: pulumi.String("azure-native:avs:HcxEnterpriseSite"),
		},
		{
			Type: pulumi.String("azure-native:avs/v20200320:HcxEnterpriseSite"),
		},
		{
			Type: pulumi.String("azure-native:avs/v20200717preview:HcxEnterpriseSite"),
		},
		{
			Type: pulumi.String("azure-native:avs/v20210601:HcxEnterpriseSite"),
		},
		{
			Type: pulumi.String("azure-native:avs/v20211201:HcxEnterpriseSite"),
		},
		{
			Type: pulumi.String("azure-native:avs/v20220501:HcxEnterpriseSite"),
		},
	})
	opts = append(opts, aliases)
	var resource HcxEnterpriseSite
	err := ctx.RegisterResource("azure-native:avs/v20210101preview:HcxEnterpriseSite", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetHcxEnterpriseSite(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *HcxEnterpriseSiteState, opts ...pulumi.ResourceOption) (*HcxEnterpriseSite, error) {
	var resource HcxEnterpriseSite
	err := ctx.ReadResource("azure-native:avs/v20210101preview:HcxEnterpriseSite", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type hcxEnterpriseSiteState struct {
}

type HcxEnterpriseSiteState struct {
}

func (HcxEnterpriseSiteState) ElementType() reflect.Type {
	return reflect.TypeOf((*hcxEnterpriseSiteState)(nil)).Elem()
}

type hcxEnterpriseSiteArgs struct {
	HcxEnterpriseSiteName *string `pulumi:"hcxEnterpriseSiteName"`
	PrivateCloudName      string  `pulumi:"privateCloudName"`
	ResourceGroupName     string  `pulumi:"resourceGroupName"`
}


type HcxEnterpriseSiteArgs struct {
	HcxEnterpriseSiteName pulumi.StringPtrInput
	PrivateCloudName      pulumi.StringInput
	ResourceGroupName     pulumi.StringInput
}

func (HcxEnterpriseSiteArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*hcxEnterpriseSiteArgs)(nil)).Elem()
}

type HcxEnterpriseSiteInput interface {
	pulumi.Input

	ToHcxEnterpriseSiteOutput() HcxEnterpriseSiteOutput
	ToHcxEnterpriseSiteOutputWithContext(ctx context.Context) HcxEnterpriseSiteOutput
}

func (*HcxEnterpriseSite) ElementType() reflect.Type {
	return reflect.TypeOf((**HcxEnterpriseSite)(nil)).Elem()
}

func (i *HcxEnterpriseSite) ToHcxEnterpriseSiteOutput() HcxEnterpriseSiteOutput {
	return i.ToHcxEnterpriseSiteOutputWithContext(context.Background())
}

func (i *HcxEnterpriseSite) ToHcxEnterpriseSiteOutputWithContext(ctx context.Context) HcxEnterpriseSiteOutput {
	return pulumi.ToOutputWithContext(ctx, i).(HcxEnterpriseSiteOutput)
}

type HcxEnterpriseSiteOutput struct{ *pulumi.OutputState }

func (HcxEnterpriseSiteOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**HcxEnterpriseSite)(nil)).Elem()
}

func (o HcxEnterpriseSiteOutput) ToHcxEnterpriseSiteOutput() HcxEnterpriseSiteOutput {
	return o
}

func (o HcxEnterpriseSiteOutput) ToHcxEnterpriseSiteOutputWithContext(ctx context.Context) HcxEnterpriseSiteOutput {
	return o
}

func (o HcxEnterpriseSiteOutput) ActivationKey() pulumi.StringOutput {
	return o.ApplyT(func(v *HcxEnterpriseSite) pulumi.StringOutput { return v.ActivationKey }).(pulumi.StringOutput)
}

func (o HcxEnterpriseSiteOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *HcxEnterpriseSite) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o HcxEnterpriseSiteOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v *HcxEnterpriseSite) pulumi.StringOutput { return v.Status }).(pulumi.StringOutput)
}

func (o HcxEnterpriseSiteOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *HcxEnterpriseSite) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(HcxEnterpriseSiteOutput{})
}
