


package v20200701preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type Extension struct {
	pulumi.CustomResourceState

	AutoUpgradeMinorVersion        pulumi.BoolPtrOutput                   `pulumi:"autoUpgradeMinorVersion"`
	ConfigurationProtectedSettings pulumi.StringMapOutput                 `pulumi:"configurationProtectedSettings"`
	ConfigurationSettings          pulumi.StringMapOutput                 `pulumi:"configurationSettings"`
	CreationTime                   pulumi.StringOutput                    `pulumi:"creationTime"`
	ErrorInfo                      ErrorDefinitionResponseOutput          `pulumi:"errorInfo"`
	ExtensionType                  pulumi.StringPtrOutput                 `pulumi:"extensionType"`
	Identity                       ConfigurationIdentityResponsePtrOutput `pulumi:"identity"`
	InstallState                   pulumi.StringOutput                    `pulumi:"installState"`
	LastModifiedTime               pulumi.StringOutput                    `pulumi:"lastModifiedTime"`
	LastStatusTime                 pulumi.StringOutput                    `pulumi:"lastStatusTime"`
	Name                           pulumi.StringOutput                    `pulumi:"name"`
	ReleaseTrain                   pulumi.StringPtrOutput                 `pulumi:"releaseTrain"`
	Scope                          ScopeResponsePtrOutput                 `pulumi:"scope"`
	Statuses                       ExtensionStatusResponseArrayOutput     `pulumi:"statuses"`
	SystemData                     SystemDataResponsePtrOutput            `pulumi:"systemData"`
	Type                           pulumi.StringOutput                    `pulumi:"type"`
	Version                        pulumi.StringPtrOutput                 `pulumi:"version"`
}


func NewExtension(ctx *pulumi.Context,
	name string, args *ExtensionArgs, opts ...pulumi.ResourceOption) (*Extension, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ClusterName == nil {
		return nil, errors.New("invalid value for required argument 'ClusterName'")
	}
	if args.ClusterResourceName == nil {
		return nil, errors.New("invalid value for required argument 'ClusterResourceName'")
	}
	if args.ClusterRp == nil {
		return nil, errors.New("invalid value for required argument 'ClusterRp'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:kubernetesconfiguration:Extension"),
		},
		{
			Type: pulumi.String("azure-native:kubernetesconfiguration/v20210501preview:Extension"),
		},
		{
			Type: pulumi.String("azure-native:kubernetesconfiguration/v20210901:Extension"),
		},
		{
			Type: pulumi.String("azure-native:kubernetesconfiguration/v20211101preview:Extension"),
		},
		{
			Type: pulumi.String("azure-native:kubernetesconfiguration/v20220101preview:Extension"),
		},
	})
	opts = append(opts, aliases)
	var resource Extension
	err := ctx.RegisterResource("azure-native:kubernetesconfiguration/v20200701preview:Extension", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetExtension(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ExtensionState, opts ...pulumi.ResourceOption) (*Extension, error) {
	var resource Extension
	err := ctx.ReadResource("azure-native:kubernetesconfiguration/v20200701preview:Extension", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type extensionState struct {
}

type ExtensionState struct {
}

func (ExtensionState) ElementType() reflect.Type {
	return reflect.TypeOf((*extensionState)(nil)).Elem()
}

type extensionArgs struct {
	AutoUpgradeMinorVersion        *bool                  `pulumi:"autoUpgradeMinorVersion"`
	ClusterName                    string                 `pulumi:"clusterName"`
	ClusterResourceName            string                 `pulumi:"clusterResourceName"`
	ClusterRp                      string                 `pulumi:"clusterRp"`
	ConfigurationProtectedSettings map[string]string      `pulumi:"configurationProtectedSettings"`
	ConfigurationSettings          map[string]string      `pulumi:"configurationSettings"`
	ExtensionInstanceName          *string                `pulumi:"extensionInstanceName"`
	ExtensionType                  *string                `pulumi:"extensionType"`
	Identity                       *ConfigurationIdentity `pulumi:"identity"`
	ReleaseTrain                   *string                `pulumi:"releaseTrain"`
	ResourceGroupName              string                 `pulumi:"resourceGroupName"`
	Scope                          *Scope                 `pulumi:"scope"`
	Statuses                       []ExtensionStatus      `pulumi:"statuses"`
	Version                        *string                `pulumi:"version"`
}


type ExtensionArgs struct {
	AutoUpgradeMinorVersion        pulumi.BoolPtrInput
	ClusterName                    pulumi.StringInput
	ClusterResourceName            pulumi.StringInput
	ClusterRp                      pulumi.StringInput
	ConfigurationProtectedSettings pulumi.StringMapInput
	ConfigurationSettings          pulumi.StringMapInput
	ExtensionInstanceName          pulumi.StringPtrInput
	ExtensionType                  pulumi.StringPtrInput
	Identity                       ConfigurationIdentityPtrInput
	ReleaseTrain                   pulumi.StringPtrInput
	ResourceGroupName              pulumi.StringInput
	Scope                          ScopePtrInput
	Statuses                       ExtensionStatusArrayInput
	Version                        pulumi.StringPtrInput
}

func (ExtensionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*extensionArgs)(nil)).Elem()
}

type ExtensionInput interface {
	pulumi.Input

	ToExtensionOutput() ExtensionOutput
	ToExtensionOutputWithContext(ctx context.Context) ExtensionOutput
}

func (*Extension) ElementType() reflect.Type {
	return reflect.TypeOf((**Extension)(nil)).Elem()
}

func (i *Extension) ToExtensionOutput() ExtensionOutput {
	return i.ToExtensionOutputWithContext(context.Background())
}

func (i *Extension) ToExtensionOutputWithContext(ctx context.Context) ExtensionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ExtensionOutput)
}

type ExtensionOutput struct{ *pulumi.OutputState }

func (ExtensionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Extension)(nil)).Elem()
}

func (o ExtensionOutput) ToExtensionOutput() ExtensionOutput {
	return o
}

func (o ExtensionOutput) ToExtensionOutputWithContext(ctx context.Context) ExtensionOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(ExtensionOutput{})
}
