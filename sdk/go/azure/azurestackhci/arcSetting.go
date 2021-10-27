


package azurestackhci

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type ArcSetting struct {
	pulumi.CustomResourceState

	AggregateState           pulumi.StringOutput             `pulumi:"aggregateState"`
	ArcInstanceResourceGroup pulumi.StringOutput             `pulumi:"arcInstanceResourceGroup"`
	CreatedAt                pulumi.StringPtrOutput          `pulumi:"createdAt"`
	CreatedBy                pulumi.StringPtrOutput          `pulumi:"createdBy"`
	CreatedByType            pulumi.StringPtrOutput          `pulumi:"createdByType"`
	LastModifiedAt           pulumi.StringPtrOutput          `pulumi:"lastModifiedAt"`
	LastModifiedBy           pulumi.StringPtrOutput          `pulumi:"lastModifiedBy"`
	LastModifiedByType       pulumi.StringPtrOutput          `pulumi:"lastModifiedByType"`
	Name                     pulumi.StringOutput             `pulumi:"name"`
	PerNodeDetails           PerNodeStateResponseArrayOutput `pulumi:"perNodeDetails"`
	ProvisioningState        pulumi.StringOutput             `pulumi:"provisioningState"`
	Type                     pulumi.StringOutput             `pulumi:"type"`
}


func NewArcSetting(ctx *pulumi.Context,
	name string, args *ArcSettingArgs, opts ...pulumi.ResourceOption) (*ArcSetting, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ClusterName == nil {
		return nil, errors.New("invalid value for required argument 'ClusterName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-nextgen:azurestackhci:ArcSetting"),
		},
		{
			Type: pulumi.String("azure-native:azurestackhci/v20210101preview:ArcSetting"),
		},
		{
			Type: pulumi.String("azure-nextgen:azurestackhci/v20210101preview:ArcSetting"),
		},
		{
			Type: pulumi.String("azure-native:azurestackhci/v20210901:ArcSetting"),
		},
		{
			Type: pulumi.String("azure-nextgen:azurestackhci/v20210901:ArcSetting"),
		},
	})
	opts = append(opts, aliases)
	var resource ArcSetting
	err := ctx.RegisterResource("azure-native:azurestackhci:ArcSetting", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetArcSetting(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ArcSettingState, opts ...pulumi.ResourceOption) (*ArcSetting, error) {
	var resource ArcSetting
	err := ctx.ReadResource("azure-native:azurestackhci:ArcSetting", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type arcSettingState struct {
}

type ArcSettingState struct {
}

func (ArcSettingState) ElementType() reflect.Type {
	return reflect.TypeOf((*arcSettingState)(nil)).Elem()
}

type arcSettingArgs struct {
	ArcSettingName     *string `pulumi:"arcSettingName"`
	ClusterName        string  `pulumi:"clusterName"`
	CreatedAt          *string `pulumi:"createdAt"`
	CreatedBy          *string `pulumi:"createdBy"`
	CreatedByType      *string `pulumi:"createdByType"`
	LastModifiedAt     *string `pulumi:"lastModifiedAt"`
	LastModifiedBy     *string `pulumi:"lastModifiedBy"`
	LastModifiedByType *string `pulumi:"lastModifiedByType"`
	ResourceGroupName  string  `pulumi:"resourceGroupName"`
}


type ArcSettingArgs struct {
	ArcSettingName     pulumi.StringPtrInput
	ClusterName        pulumi.StringInput
	CreatedAt          pulumi.StringPtrInput
	CreatedBy          pulumi.StringPtrInput
	CreatedByType      pulumi.StringPtrInput
	LastModifiedAt     pulumi.StringPtrInput
	LastModifiedBy     pulumi.StringPtrInput
	LastModifiedByType pulumi.StringPtrInput
	ResourceGroupName  pulumi.StringInput
}

func (ArcSettingArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*arcSettingArgs)(nil)).Elem()
}

type ArcSettingInput interface {
	pulumi.Input

	ToArcSettingOutput() ArcSettingOutput
	ToArcSettingOutputWithContext(ctx context.Context) ArcSettingOutput
}

func (*ArcSetting) ElementType() reflect.Type {
	return reflect.TypeOf((*ArcSetting)(nil))
}

func (i *ArcSetting) ToArcSettingOutput() ArcSettingOutput {
	return i.ToArcSettingOutputWithContext(context.Background())
}

func (i *ArcSetting) ToArcSettingOutputWithContext(ctx context.Context) ArcSettingOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ArcSettingOutput)
}

type ArcSettingOutput struct{ *pulumi.OutputState }

func (ArcSettingOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ArcSetting)(nil))
}

func (o ArcSettingOutput) ToArcSettingOutput() ArcSettingOutput {
	return o
}

func (o ArcSettingOutput) ToArcSettingOutputWithContext(ctx context.Context) ArcSettingOutput {
	return o
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ArcSettingInput)(nil)).Elem(), &ArcSetting{})
	pulumi.RegisterOutputType(ArcSettingOutput{})
}
