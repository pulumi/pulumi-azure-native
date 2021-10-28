


package v20210701preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type SecurityConnector struct {
	pulumi.CustomResourceState

	CloudName           pulumi.StringPtrOutput                                         `pulumi:"cloudName"`
	Etag                pulumi.StringPtrOutput                                         `pulumi:"etag"`
	HierarchyIdentifier pulumi.StringPtrOutput                                         `pulumi:"hierarchyIdentifier"`
	Kind                pulumi.StringPtrOutput                                         `pulumi:"kind"`
	Location            pulumi.StringPtrOutput                                         `pulumi:"location"`
	Name                pulumi.StringOutput                                            `pulumi:"name"`
	Offerings           pulumi.ArrayOutput                                             `pulumi:"offerings"`
	OrganizationalData  SecurityConnectorPropertiesResponseOrganizationalDataPtrOutput `pulumi:"organizationalData"`
	SystemData          SystemDataResponseOutput                                       `pulumi:"systemData"`
	Tags                pulumi.StringMapOutput                                         `pulumi:"tags"`
	Type                pulumi.StringOutput                                            `pulumi:"type"`
}


func NewSecurityConnector(ctx *pulumi.Context,
	name string, args *SecurityConnectorArgs, opts ...pulumi.ResourceOption) (*SecurityConnector, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-nextgen:security/v20210701preview:SecurityConnector"),
		},
		{
			Type: pulumi.String("azure-native:security:SecurityConnector"),
		},
		{
			Type: pulumi.String("azure-nextgen:security:SecurityConnector"),
		},
	})
	opts = append(opts, aliases)
	var resource SecurityConnector
	err := ctx.RegisterResource("azure-native:security/v20210701preview:SecurityConnector", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetSecurityConnector(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SecurityConnectorState, opts ...pulumi.ResourceOption) (*SecurityConnector, error) {
	var resource SecurityConnector
	err := ctx.ReadResource("azure-native:security/v20210701preview:SecurityConnector", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type securityConnectorState struct {
}

type SecurityConnectorState struct {
}

func (SecurityConnectorState) ElementType() reflect.Type {
	return reflect.TypeOf((*securityConnectorState)(nil)).Elem()
}

type securityConnectorArgs struct {
	CloudName             *string                                        `pulumi:"cloudName"`
	Etag                  *string                                        `pulumi:"etag"`
	HierarchyIdentifier   *string                                        `pulumi:"hierarchyIdentifier"`
	Kind                  *string                                        `pulumi:"kind"`
	Location              *string                                        `pulumi:"location"`
	Offerings             []interface{}                                  `pulumi:"offerings"`
	OrganizationalData    *SecurityConnectorPropertiesOrganizationalData `pulumi:"organizationalData"`
	ResourceGroupName     string                                         `pulumi:"resourceGroupName"`
	SecurityConnectorName *string                                        `pulumi:"securityConnectorName"`
	Tags                  map[string]string                              `pulumi:"tags"`
}


type SecurityConnectorArgs struct {
	CloudName             pulumi.StringPtrInput
	Etag                  pulumi.StringPtrInput
	HierarchyIdentifier   pulumi.StringPtrInput
	Kind                  pulumi.StringPtrInput
	Location              pulumi.StringPtrInput
	Offerings             pulumi.ArrayInput
	OrganizationalData    SecurityConnectorPropertiesOrganizationalDataPtrInput
	ResourceGroupName     pulumi.StringInput
	SecurityConnectorName pulumi.StringPtrInput
	Tags                  pulumi.StringMapInput
}

func (SecurityConnectorArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*securityConnectorArgs)(nil)).Elem()
}

type SecurityConnectorInput interface {
	pulumi.Input

	ToSecurityConnectorOutput() SecurityConnectorOutput
	ToSecurityConnectorOutputWithContext(ctx context.Context) SecurityConnectorOutput
}

func (*SecurityConnector) ElementType() reflect.Type {
	return reflect.TypeOf((*SecurityConnector)(nil))
}

func (i *SecurityConnector) ToSecurityConnectorOutput() SecurityConnectorOutput {
	return i.ToSecurityConnectorOutputWithContext(context.Background())
}

func (i *SecurityConnector) ToSecurityConnectorOutputWithContext(ctx context.Context) SecurityConnectorOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SecurityConnectorOutput)
}

type SecurityConnectorOutput struct{ *pulumi.OutputState }

func (SecurityConnectorOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SecurityConnector)(nil))
}

func (o SecurityConnectorOutput) ToSecurityConnectorOutput() SecurityConnectorOutput {
	return o
}

func (o SecurityConnectorOutput) ToSecurityConnectorOutputWithContext(ctx context.Context) SecurityConnectorOutput {
	return o
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*SecurityConnectorInput)(nil)).Elem(), &SecurityConnector{})
	pulumi.RegisterOutputType(SecurityConnectorOutput{})
}
