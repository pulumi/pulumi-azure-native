


package notebooks

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type NotebookProxy struct {
	pulumi.CustomResourceState

	Hostname            pulumi.StringPtrOutput                      `pulumi:"hostname"`
	Name                pulumi.StringOutput                         `pulumi:"name"`
	PublicDns           pulumi.StringPtrOutput                      `pulumi:"publicDns"`
	PublicNetworkAccess pulumi.StringPtrOutput                      `pulumi:"publicNetworkAccess"`
	Region              pulumi.StringPtrOutput                      `pulumi:"region"`
	ResourceId          pulumi.StringOutput                         `pulumi:"resourceId"`
	SecondaryAppId      pulumi.StringPtrOutput                      `pulumi:"secondaryAppId"`
	SystemData          NotebookResourceSystemDataResponsePtrOutput `pulumi:"systemData"`
	Type                pulumi.StringOutput                         `pulumi:"type"`
}


func NewNotebookProxy(ctx *pulumi.Context,
	name string, args *NotebookProxyArgs, opts ...pulumi.ResourceOption) (*NotebookProxy, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-nextgen:notebooks:NotebookProxy"),
		},
		{
			Type: pulumi.String("azure-native:notebooks/v20191011preview:NotebookProxy"),
		},
		{
			Type: pulumi.String("azure-nextgen:notebooks/v20191011preview:NotebookProxy"),
		},
	})
	opts = append(opts, aliases)
	var resource NotebookProxy
	err := ctx.RegisterResource("azure-native:notebooks:NotebookProxy", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetNotebookProxy(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *NotebookProxyState, opts ...pulumi.ResourceOption) (*NotebookProxy, error) {
	var resource NotebookProxy
	err := ctx.ReadResource("azure-native:notebooks:NotebookProxy", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type notebookProxyState struct {
}

type NotebookProxyState struct {
}

func (NotebookProxyState) ElementType() reflect.Type {
	return reflect.TypeOf((*notebookProxyState)(nil)).Elem()
}

type notebookProxyArgs struct {
	Hostname            *string                     `pulumi:"hostname"`
	PublicDns           *string                     `pulumi:"publicDns"`
	PublicNetworkAccess *string                     `pulumi:"publicNetworkAccess"`
	Region              *string                     `pulumi:"region"`
	ResourceGroupName   string                      `pulumi:"resourceGroupName"`
	ResourceName        *string                     `pulumi:"resourceName"`
	SecondaryAppId      *string                     `pulumi:"secondaryAppId"`
	SystemData          *NotebookResourceSystemData `pulumi:"systemData"`
}


type NotebookProxyArgs struct {
	Hostname            pulumi.StringPtrInput
	PublicDns           pulumi.StringPtrInput
	PublicNetworkAccess pulumi.StringPtrInput
	Region              pulumi.StringPtrInput
	ResourceGroupName   pulumi.StringInput
	ResourceName        pulumi.StringPtrInput
	SecondaryAppId      pulumi.StringPtrInput
	SystemData          NotebookResourceSystemDataPtrInput
}

func (NotebookProxyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*notebookProxyArgs)(nil)).Elem()
}

type NotebookProxyInput interface {
	pulumi.Input

	ToNotebookProxyOutput() NotebookProxyOutput
	ToNotebookProxyOutputWithContext(ctx context.Context) NotebookProxyOutput
}

func (*NotebookProxy) ElementType() reflect.Type {
	return reflect.TypeOf((*NotebookProxy)(nil))
}

func (i *NotebookProxy) ToNotebookProxyOutput() NotebookProxyOutput {
	return i.ToNotebookProxyOutputWithContext(context.Background())
}

func (i *NotebookProxy) ToNotebookProxyOutputWithContext(ctx context.Context) NotebookProxyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NotebookProxyOutput)
}

type NotebookProxyOutput struct{ *pulumi.OutputState }

func (NotebookProxyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*NotebookProxy)(nil))
}

func (o NotebookProxyOutput) ToNotebookProxyOutput() NotebookProxyOutput {
	return o
}

func (o NotebookProxyOutput) ToNotebookProxyOutputWithContext(ctx context.Context) NotebookProxyOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(NotebookProxyOutput{})
}
