


package v20150501preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)


type RouteTable struct {
	pulumi.CustomResourceState

	Etag              pulumi.StringPtrOutput         `pulumi:"etag"`
	Location          pulumi.StringOutput            `pulumi:"location"`
	Name              pulumi.StringOutput            `pulumi:"name"`
	ProvisioningState pulumi.StringPtrOutput         `pulumi:"provisioningState"`
	Routes            RouteResponseArrayOutput       `pulumi:"routes"`
	Subnets           SubResourceResponseArrayOutput `pulumi:"subnets"`
	Tags              pulumi.StringMapOutput         `pulumi:"tags"`
	Type              pulumi.StringOutput            `pulumi:"type"`
}


func NewRouteTable(ctx *pulumi.Context,
	name string, args *RouteTableArgs, opts ...pulumi.ResourceOption) (*RouteTable, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:network:RouteTable"),
		},
		{
			Type: pulumi.String("azure-native:network/v20150615:RouteTable"),
		},
		{
			Type: pulumi.String("azure-native:network/v20160330:RouteTable"),
		},
		{
			Type: pulumi.String("azure-native:network/v20160601:RouteTable"),
		},
		{
			Type: pulumi.String("azure-native:network/v20160901:RouteTable"),
		},
		{
			Type: pulumi.String("azure-native:network/v20161201:RouteTable"),
		},
		{
			Type: pulumi.String("azure-native:network/v20170301:RouteTable"),
		},
		{
			Type: pulumi.String("azure-native:network/v20170601:RouteTable"),
		},
		{
			Type: pulumi.String("azure-native:network/v20170801:RouteTable"),
		},
		{
			Type: pulumi.String("azure-native:network/v20170901:RouteTable"),
		},
		{
			Type: pulumi.String("azure-native:network/v20171001:RouteTable"),
		},
		{
			Type: pulumi.String("azure-native:network/v20171101:RouteTable"),
		},
		{
			Type: pulumi.String("azure-native:network/v20180101:RouteTable"),
		},
		{
			Type: pulumi.String("azure-native:network/v20180201:RouteTable"),
		},
		{
			Type: pulumi.String("azure-native:network/v20180401:RouteTable"),
		},
		{
			Type: pulumi.String("azure-native:network/v20180601:RouteTable"),
		},
		{
			Type: pulumi.String("azure-native:network/v20180701:RouteTable"),
		},
		{
			Type: pulumi.String("azure-native:network/v20180801:RouteTable"),
		},
		{
			Type: pulumi.String("azure-native:network/v20181001:RouteTable"),
		},
		{
			Type: pulumi.String("azure-native:network/v20181101:RouteTable"),
		},
		{
			Type: pulumi.String("azure-native:network/v20181201:RouteTable"),
		},
		{
			Type: pulumi.String("azure-native:network/v20190201:RouteTable"),
		},
		{
			Type: pulumi.String("azure-native:network/v20190401:RouteTable"),
		},
		{
			Type: pulumi.String("azure-native:network/v20190601:RouteTable"),
		},
		{
			Type: pulumi.String("azure-native:network/v20190701:RouteTable"),
		},
		{
			Type: pulumi.String("azure-native:network/v20190801:RouteTable"),
		},
		{
			Type: pulumi.String("azure-native:network/v20190901:RouteTable"),
		},
		{
			Type: pulumi.String("azure-native:network/v20191101:RouteTable"),
		},
		{
			Type: pulumi.String("azure-native:network/v20191201:RouteTable"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200301:RouteTable"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200401:RouteTable"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200501:RouteTable"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200601:RouteTable"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200701:RouteTable"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200801:RouteTable"),
		},
		{
			Type: pulumi.String("azure-native:network/v20201101:RouteTable"),
		},
		{
			Type: pulumi.String("azure-native:network/v20210201:RouteTable"),
		},
		{
			Type: pulumi.String("azure-native:network/v20210301:RouteTable"),
		},
		{
			Type: pulumi.String("azure-native:network/v20210501:RouteTable"),
		},
		{
			Type: pulumi.String("azure-native:network/v20210801:RouteTable"),
		},
		{
			Type: pulumi.String("azure-native:network/v20220101:RouteTable"),
		},
	})
	opts = append(opts, aliases)
	var resource RouteTable
	err := ctx.RegisterResource("azure-native:network/v20150501preview:RouteTable", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetRouteTable(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *RouteTableState, opts ...pulumi.ResourceOption) (*RouteTable, error) {
	var resource RouteTable
	err := ctx.ReadResource("azure-native:network/v20150501preview:RouteTable", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type routeTableState struct {
}

type RouteTableState struct {
}

func (RouteTableState) ElementType() reflect.Type {
	return reflect.TypeOf((*routeTableState)(nil)).Elem()
}

type routeTableArgs struct {
	Location          *string           `pulumi:"location"`
	ProvisioningState *string           `pulumi:"provisioningState"`
	ResourceGroupName string            `pulumi:"resourceGroupName"`
	RouteTableName    *string           `pulumi:"routeTableName"`
	Routes            []RouteType       `pulumi:"routes"`
	Subnets           []SubResource     `pulumi:"subnets"`
	Tags              map[string]string `pulumi:"tags"`
}


type RouteTableArgs struct {
	Location          pulumi.StringPtrInput
	ProvisioningState pulumi.StringPtrInput
	ResourceGroupName pulumi.StringInput
	RouteTableName    pulumi.StringPtrInput
	Routes            RouteTypeArrayInput
	Subnets           SubResourceArrayInput
	Tags              pulumi.StringMapInput
}

func (RouteTableArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*routeTableArgs)(nil)).Elem()
}

type RouteTableInput interface {
	pulumi.Input

	ToRouteTableOutput() RouteTableOutput
	ToRouteTableOutputWithContext(ctx context.Context) RouteTableOutput
}

func (*RouteTable) ElementType() reflect.Type {
	return reflect.TypeOf((**RouteTable)(nil)).Elem()
}

func (i *RouteTable) ToRouteTableOutput() RouteTableOutput {
	return i.ToRouteTableOutputWithContext(context.Background())
}

func (i *RouteTable) ToRouteTableOutputWithContext(ctx context.Context) RouteTableOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RouteTableOutput)
}

type RouteTableOutput struct{ *pulumi.OutputState }

func (RouteTableOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**RouteTable)(nil)).Elem()
}

func (o RouteTableOutput) ToRouteTableOutput() RouteTableOutput {
	return o
}

func (o RouteTableOutput) ToRouteTableOutputWithContext(ctx context.Context) RouteTableOutput {
	return o
}

func (o RouteTableOutput) Etag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *RouteTable) pulumi.StringPtrOutput { return v.Etag }).(pulumi.StringPtrOutput)
}

func (o RouteTableOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *RouteTable) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

func (o RouteTableOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *RouteTable) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o RouteTableOutput) ProvisioningState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *RouteTable) pulumi.StringPtrOutput { return v.ProvisioningState }).(pulumi.StringPtrOutput)
}

func (o RouteTableOutput) Routes() RouteResponseArrayOutput {
	return o.ApplyT(func(v *RouteTable) RouteResponseArrayOutput { return v.Routes }).(RouteResponseArrayOutput)
}

func (o RouteTableOutput) Subnets() SubResourceResponseArrayOutput {
	return o.ApplyT(func(v *RouteTable) SubResourceResponseArrayOutput { return v.Subnets }).(SubResourceResponseArrayOutput)
}

func (o RouteTableOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *RouteTable) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

func (o RouteTableOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *RouteTable) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(RouteTableOutput{})
}
