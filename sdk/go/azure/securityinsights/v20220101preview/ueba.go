


package v20220101preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type Ueba struct {
	pulumi.CustomResourceState

	DataSources pulumi.StringArrayOutput `pulumi:"dataSources"`
	Etag        pulumi.StringPtrOutput   `pulumi:"etag"`
	Kind        pulumi.StringOutput      `pulumi:"kind"`
	Name        pulumi.StringOutput      `pulumi:"name"`
	SystemData  SystemDataResponseOutput `pulumi:"systemData"`
	Type        pulumi.StringOutput      `pulumi:"type"`
}


func NewUeba(ctx *pulumi.Context,
	name string, args *UebaArgs, opts ...pulumi.ResourceOption) (*Ueba, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Kind == nil {
		return nil, errors.New("invalid value for required argument 'Kind'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.WorkspaceName == nil {
		return nil, errors.New("invalid value for required argument 'WorkspaceName'")
	}
	args.Kind = pulumi.String("Ueba")
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:securityinsights:Ueba"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20190101preview:Ueba"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20210301preview:Ueba"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20210901preview:Ueba"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20211001preview:Ueba"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20220401preview:Ueba"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20220501preview:Ueba"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20220601preview:Ueba"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20220701preview:Ueba"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20220801preview:Ueba"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20220901preview:Ueba"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20221001preview:Ueba"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20221101preview:Ueba"),
		},
	})
	opts = append(opts, aliases)
	var resource Ueba
	err := ctx.RegisterResource("azure-native:securityinsights/v20220101preview:Ueba", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetUeba(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *UebaState, opts ...pulumi.ResourceOption) (*Ueba, error) {
	var resource Ueba
	err := ctx.ReadResource("azure-native:securityinsights/v20220101preview:Ueba", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type uebaState struct {
}

type UebaState struct {
}

func (UebaState) ElementType() reflect.Type {
	return reflect.TypeOf((*uebaState)(nil)).Elem()
}

type uebaArgs struct {
	DataSources       []string `pulumi:"dataSources"`
	Kind              string   `pulumi:"kind"`
	ResourceGroupName string   `pulumi:"resourceGroupName"`
	SettingsName      *string  `pulumi:"settingsName"`
	WorkspaceName     string   `pulumi:"workspaceName"`
}


type UebaArgs struct {
	DataSources       pulumi.StringArrayInput
	Kind              pulumi.StringInput
	ResourceGroupName pulumi.StringInput
	SettingsName      pulumi.StringPtrInput
	WorkspaceName     pulumi.StringInput
}

func (UebaArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*uebaArgs)(nil)).Elem()
}

type UebaInput interface {
	pulumi.Input

	ToUebaOutput() UebaOutput
	ToUebaOutputWithContext(ctx context.Context) UebaOutput
}

func (*Ueba) ElementType() reflect.Type {
	return reflect.TypeOf((**Ueba)(nil)).Elem()
}

func (i *Ueba) ToUebaOutput() UebaOutput {
	return i.ToUebaOutputWithContext(context.Background())
}

func (i *Ueba) ToUebaOutputWithContext(ctx context.Context) UebaOutput {
	return pulumi.ToOutputWithContext(ctx, i).(UebaOutput)
}

type UebaOutput struct{ *pulumi.OutputState }

func (UebaOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Ueba)(nil)).Elem()
}

func (o UebaOutput) ToUebaOutput() UebaOutput {
	return o
}

func (o UebaOutput) ToUebaOutputWithContext(ctx context.Context) UebaOutput {
	return o
}

func (o UebaOutput) DataSources() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *Ueba) pulumi.StringArrayOutput { return v.DataSources }).(pulumi.StringArrayOutput)
}

func (o UebaOutput) Etag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Ueba) pulumi.StringPtrOutput { return v.Etag }).(pulumi.StringPtrOutput)
}

func (o UebaOutput) Kind() pulumi.StringOutput {
	return o.ApplyT(func(v *Ueba) pulumi.StringOutput { return v.Kind }).(pulumi.StringOutput)
}

func (o UebaOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Ueba) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o UebaOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *Ueba) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

func (o UebaOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *Ueba) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(UebaOutput{})
}
