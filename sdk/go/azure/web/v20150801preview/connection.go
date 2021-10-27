


package v20150801preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type Connection struct {
	pulumi.CustomResourceState

	Api                      ExpandedParentApiEntityResponsePtrOutput           `pulumi:"api"`
	ChangedTime              pulumi.StringPtrOutput                             `pulumi:"changedTime"`
	CreatedTime              pulumi.StringPtrOutput                             `pulumi:"createdTime"`
	CustomParameterValues    ParameterCustomLoginSettingValuesResponseMapOutput `pulumi:"customParameterValues"`
	DisplayName              pulumi.StringPtrOutput                             `pulumi:"displayName"`
	FirstExpirationTime      pulumi.StringPtrOutput                             `pulumi:"firstExpirationTime"`
	Keywords                 pulumi.StringArrayOutput                           `pulumi:"keywords"`
	Kind                     pulumi.StringPtrOutput                             `pulumi:"kind"`
	Location                 pulumi.StringOutput                                `pulumi:"location"`
	Metadata                 pulumi.AnyOutput                                   `pulumi:"metadata"`
	Name                     pulumi.StringPtrOutput                             `pulumi:"name"`
	NonSecretParameterValues pulumi.MapOutput                                   `pulumi:"nonSecretParameterValues"`
	ParameterValues          pulumi.MapOutput                                   `pulumi:"parameterValues"`
	Statuses                 ConnectionStatusResponseArrayOutput                `pulumi:"statuses"`
	Tags                     pulumi.StringMapOutput                             `pulumi:"tags"`
	TenantId                 pulumi.StringPtrOutput                             `pulumi:"tenantId"`
	Type                     pulumi.StringPtrOutput                             `pulumi:"type"`
}


func NewConnection(ctx *pulumi.Context,
	name string, args *ConnectionArgs, opts ...pulumi.ResourceOption) (*Connection, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-nextgen:web/v20150801preview:Connection"),
		},
		{
			Type: pulumi.String("azure-native:web:Connection"),
		},
		{
			Type: pulumi.String("azure-nextgen:web:Connection"),
		},
		{
			Type: pulumi.String("azure-native:web/v20160601:Connection"),
		},
		{
			Type: pulumi.String("azure-nextgen:web/v20160601:Connection"),
		},
	})
	opts = append(opts, aliases)
	var resource Connection
	err := ctx.RegisterResource("azure-native:web/v20150801preview:Connection", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetConnection(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ConnectionState, opts ...pulumi.ResourceOption) (*Connection, error) {
	var resource Connection
	err := ctx.ReadResource("azure-native:web/v20150801preview:Connection", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type connectionState struct {
}

type ConnectionState struct {
}

func (ConnectionState) ElementType() reflect.Type {
	return reflect.TypeOf((*connectionState)(nil)).Elem()
}

type connectionArgs struct {
	Api                      *ExpandedParentApiEntity                     `pulumi:"api"`
	ChangedTime              *string                                      `pulumi:"changedTime"`
	ConnectionName           *string                                      `pulumi:"connectionName"`
	CreatedTime              *string                                      `pulumi:"createdTime"`
	CustomParameterValues    map[string]ParameterCustomLoginSettingValues `pulumi:"customParameterValues"`
	DisplayName              *string                                      `pulumi:"displayName"`
	FirstExpirationTime      *string                                      `pulumi:"firstExpirationTime"`
	Id                       *string                                      `pulumi:"id"`
	Keywords                 []string                                     `pulumi:"keywords"`
	Kind                     *string                                      `pulumi:"kind"`
	Location                 *string                                      `pulumi:"location"`
	Metadata                 interface{}                                  `pulumi:"metadata"`
	Name                     *string                                      `pulumi:"name"`
	NonSecretParameterValues map[string]interface{}                       `pulumi:"nonSecretParameterValues"`
	ParameterValues          map[string]interface{}                       `pulumi:"parameterValues"`
	ResourceGroupName        string                                       `pulumi:"resourceGroupName"`
	Statuses                 []ConnectionStatus                           `pulumi:"statuses"`
	Tags                     map[string]string                            `pulumi:"tags"`
	TenantId                 *string                                      `pulumi:"tenantId"`
	Type                     *string                                      `pulumi:"type"`
}


type ConnectionArgs struct {
	Api                      ExpandedParentApiEntityPtrInput
	ChangedTime              pulumi.StringPtrInput
	ConnectionName           pulumi.StringPtrInput
	CreatedTime              pulumi.StringPtrInput
	CustomParameterValues    ParameterCustomLoginSettingValuesMapInput
	DisplayName              pulumi.StringPtrInput
	FirstExpirationTime      pulumi.StringPtrInput
	Id                       pulumi.StringPtrInput
	Keywords                 pulumi.StringArrayInput
	Kind                     pulumi.StringPtrInput
	Location                 pulumi.StringPtrInput
	Metadata                 pulumi.Input
	Name                     pulumi.StringPtrInput
	NonSecretParameterValues pulumi.MapInput
	ParameterValues          pulumi.MapInput
	ResourceGroupName        pulumi.StringInput
	Statuses                 ConnectionStatusArrayInput
	Tags                     pulumi.StringMapInput
	TenantId                 pulumi.StringPtrInput
	Type                     pulumi.StringPtrInput
}

func (ConnectionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*connectionArgs)(nil)).Elem()
}

type ConnectionInput interface {
	pulumi.Input

	ToConnectionOutput() ConnectionOutput
	ToConnectionOutputWithContext(ctx context.Context) ConnectionOutput
}

func (*Connection) ElementType() reflect.Type {
	return reflect.TypeOf((*Connection)(nil))
}

func (i *Connection) ToConnectionOutput() ConnectionOutput {
	return i.ToConnectionOutputWithContext(context.Background())
}

func (i *Connection) ToConnectionOutputWithContext(ctx context.Context) ConnectionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ConnectionOutput)
}

type ConnectionOutput struct{ *pulumi.OutputState }

func (ConnectionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Connection)(nil))
}

func (o ConnectionOutput) ToConnectionOutput() ConnectionOutput {
	return o
}

func (o ConnectionOutput) ToConnectionOutputWithContext(ctx context.Context) ConnectionOutput {
	return o
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ConnectionInput)(nil)).Elem(), &Connection{})
	pulumi.RegisterOutputType(ConnectionOutput{})
}
