


package v20200401preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type SqlPoolsV3 struct {
	pulumi.CustomResourceState

	AutoPauseTimer                pulumi.IntPtrOutput      `pulumi:"autoPauseTimer"`
	AutoResume                    pulumi.BoolPtrOutput     `pulumi:"autoResume"`
	CurrentServiceObjectiveName   pulumi.StringOutput      `pulumi:"currentServiceObjectiveName"`
	Kind                          pulumi.StringOutput      `pulumi:"kind"`
	Location                      pulumi.StringOutput      `pulumi:"location"`
	MaxServiceObjectiveName       pulumi.StringPtrOutput   `pulumi:"maxServiceObjectiveName"`
	Name                          pulumi.StringOutput      `pulumi:"name"`
	RequestedServiceObjectiveName pulumi.StringOutput      `pulumi:"requestedServiceObjectiveName"`
	Sku                           SkuV3ResponsePtrOutput   `pulumi:"sku"`
	SqlPoolGuid                   pulumi.StringOutput      `pulumi:"sqlPoolGuid"`
	Status                        pulumi.StringOutput      `pulumi:"status"`
	SystemData                    SystemDataResponseOutput `pulumi:"systemData"`
	Tags                          pulumi.StringMapOutput   `pulumi:"tags"`
	Type                          pulumi.StringOutput      `pulumi:"type"`
}


func NewSqlPoolsV3(ctx *pulumi.Context,
	name string, args *SqlPoolsV3Args, opts ...pulumi.ResourceOption) (*SqlPoolsV3, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.WorkspaceName == nil {
		return nil, errors.New("invalid value for required argument 'WorkspaceName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-nextgen:synapse/v20200401preview:SqlPoolsV3"),
		},
		{
			Type: pulumi.String("azure-native:synapse:SqlPoolsV3"),
		},
		{
			Type: pulumi.String("azure-nextgen:synapse:SqlPoolsV3"),
		},
		{
			Type: pulumi.String("azure-native:synapse/v20190601preview:SqlPoolsV3"),
		},
		{
			Type: pulumi.String("azure-nextgen:synapse/v20190601preview:SqlPoolsV3"),
		},
		{
			Type: pulumi.String("azure-native:synapse/v20201201:SqlPoolsV3"),
		},
		{
			Type: pulumi.String("azure-nextgen:synapse/v20201201:SqlPoolsV3"),
		},
		{
			Type: pulumi.String("azure-native:synapse/v20210301:SqlPoolsV3"),
		},
		{
			Type: pulumi.String("azure-nextgen:synapse/v20210301:SqlPoolsV3"),
		},
		{
			Type: pulumi.String("azure-native:synapse/v20210401preview:SqlPoolsV3"),
		},
		{
			Type: pulumi.String("azure-nextgen:synapse/v20210401preview:SqlPoolsV3"),
		},
		{
			Type: pulumi.String("azure-native:synapse/v20210501:SqlPoolsV3"),
		},
		{
			Type: pulumi.String("azure-nextgen:synapse/v20210501:SqlPoolsV3"),
		},
		{
			Type: pulumi.String("azure-native:synapse/v20210601:SqlPoolsV3"),
		},
		{
			Type: pulumi.String("azure-nextgen:synapse/v20210601:SqlPoolsV3"),
		},
		{
			Type: pulumi.String("azure-native:synapse/v20210601preview:SqlPoolsV3"),
		},
		{
			Type: pulumi.String("azure-nextgen:synapse/v20210601preview:SqlPoolsV3"),
		},
	})
	opts = append(opts, aliases)
	var resource SqlPoolsV3
	err := ctx.RegisterResource("azure-native:synapse/v20200401preview:SqlPoolsV3", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetSqlPoolsV3(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SqlPoolsV3State, opts ...pulumi.ResourceOption) (*SqlPoolsV3, error) {
	var resource SqlPoolsV3
	err := ctx.ReadResource("azure-native:synapse/v20200401preview:SqlPoolsV3", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type sqlPoolsV3State struct {
}

type SqlPoolsV3State struct {
}

func (SqlPoolsV3State) ElementType() reflect.Type {
	return reflect.TypeOf((*sqlPoolsV3State)(nil)).Elem()
}

type sqlPoolsV3Args struct {
	AutoPauseTimer          *int              `pulumi:"autoPauseTimer"`
	AutoResume              *bool             `pulumi:"autoResume"`
	Location                *string           `pulumi:"location"`
	MaxServiceObjectiveName *string           `pulumi:"maxServiceObjectiveName"`
	ResourceGroupName       string            `pulumi:"resourceGroupName"`
	Sku                     *SkuV3            `pulumi:"sku"`
	SqlPoolName             *string           `pulumi:"sqlPoolName"`
	Tags                    map[string]string `pulumi:"tags"`
	WorkspaceName           string            `pulumi:"workspaceName"`
}


type SqlPoolsV3Args struct {
	AutoPauseTimer          pulumi.IntPtrInput
	AutoResume              pulumi.BoolPtrInput
	Location                pulumi.StringPtrInput
	MaxServiceObjectiveName pulumi.StringPtrInput
	ResourceGroupName       pulumi.StringInput
	Sku                     SkuV3PtrInput
	SqlPoolName             pulumi.StringPtrInput
	Tags                    pulumi.StringMapInput
	WorkspaceName           pulumi.StringInput
}

func (SqlPoolsV3Args) ElementType() reflect.Type {
	return reflect.TypeOf((*sqlPoolsV3Args)(nil)).Elem()
}

type SqlPoolsV3Input interface {
	pulumi.Input

	ToSqlPoolsV3Output() SqlPoolsV3Output
	ToSqlPoolsV3OutputWithContext(ctx context.Context) SqlPoolsV3Output
}

func (*SqlPoolsV3) ElementType() reflect.Type {
	return reflect.TypeOf((*SqlPoolsV3)(nil))
}

func (i *SqlPoolsV3) ToSqlPoolsV3Output() SqlPoolsV3Output {
	return i.ToSqlPoolsV3OutputWithContext(context.Background())
}

func (i *SqlPoolsV3) ToSqlPoolsV3OutputWithContext(ctx context.Context) SqlPoolsV3Output {
	return pulumi.ToOutputWithContext(ctx, i).(SqlPoolsV3Output)
}

type SqlPoolsV3Output struct{ *pulumi.OutputState }

func (SqlPoolsV3Output) ElementType() reflect.Type {
	return reflect.TypeOf((*SqlPoolsV3)(nil))
}

func (o SqlPoolsV3Output) ToSqlPoolsV3Output() SqlPoolsV3Output {
	return o
}

func (o SqlPoolsV3Output) ToSqlPoolsV3OutputWithContext(ctx context.Context) SqlPoolsV3Output {
	return o
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*SqlPoolsV3Input)(nil)).Elem(), &SqlPoolsV3{})
	pulumi.RegisterOutputType(SqlPoolsV3Output{})
}
