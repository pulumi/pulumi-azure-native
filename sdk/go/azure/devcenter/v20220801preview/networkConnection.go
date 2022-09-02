


package v20220801preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type NetworkConnection struct {
	pulumi.CustomResourceState

	DomainJoinType              pulumi.StringOutput      `pulumi:"domainJoinType"`
	DomainName                  pulumi.StringPtrOutput   `pulumi:"domainName"`
	DomainPassword              pulumi.StringPtrOutput   `pulumi:"domainPassword"`
	DomainUsername              pulumi.StringPtrOutput   `pulumi:"domainUsername"`
	HealthCheckStatus           pulumi.StringOutput      `pulumi:"healthCheckStatus"`
	Location                    pulumi.StringOutput      `pulumi:"location"`
	Name                        pulumi.StringOutput      `pulumi:"name"`
	NetworkingResourceGroupName pulumi.StringPtrOutput   `pulumi:"networkingResourceGroupName"`
	OrganizationUnit            pulumi.StringPtrOutput   `pulumi:"organizationUnit"`
	ProvisioningState           pulumi.StringOutput      `pulumi:"provisioningState"`
	SubnetId                    pulumi.StringOutput      `pulumi:"subnetId"`
	SystemData                  SystemDataResponseOutput `pulumi:"systemData"`
	Tags                        pulumi.StringMapOutput   `pulumi:"tags"`
	Type                        pulumi.StringOutput      `pulumi:"type"`
}


func NewNetworkConnection(ctx *pulumi.Context,
	name string, args *NetworkConnectionArgs, opts ...pulumi.ResourceOption) (*NetworkConnection, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.DomainJoinType == nil {
		return nil, errors.New("invalid value for required argument 'DomainJoinType'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.SubnetId == nil {
		return nil, errors.New("invalid value for required argument 'SubnetId'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:devcenter:NetworkConnection"),
		},
		{
			Type: pulumi.String("azure-native:devcenter/v20220901preview:NetworkConnection"),
		},
	})
	opts = append(opts, aliases)
	var resource NetworkConnection
	err := ctx.RegisterResource("azure-native:devcenter/v20220801preview:NetworkConnection", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetNetworkConnection(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *NetworkConnectionState, opts ...pulumi.ResourceOption) (*NetworkConnection, error) {
	var resource NetworkConnection
	err := ctx.ReadResource("azure-native:devcenter/v20220801preview:NetworkConnection", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type networkConnectionState struct {
}

type NetworkConnectionState struct {
}

func (NetworkConnectionState) ElementType() reflect.Type {
	return reflect.TypeOf((*networkConnectionState)(nil)).Elem()
}

type networkConnectionArgs struct {
	DomainJoinType              string            `pulumi:"domainJoinType"`
	DomainName                  *string           `pulumi:"domainName"`
	DomainPassword              *string           `pulumi:"domainPassword"`
	DomainUsername              *string           `pulumi:"domainUsername"`
	Location                    *string           `pulumi:"location"`
	NetworkConnectionName       *string           `pulumi:"networkConnectionName"`
	NetworkingResourceGroupName *string           `pulumi:"networkingResourceGroupName"`
	OrganizationUnit            *string           `pulumi:"organizationUnit"`
	ResourceGroupName           string            `pulumi:"resourceGroupName"`
	SubnetId                    string            `pulumi:"subnetId"`
	Tags                        map[string]string `pulumi:"tags"`
}


type NetworkConnectionArgs struct {
	DomainJoinType              pulumi.StringInput
	DomainName                  pulumi.StringPtrInput
	DomainPassword              pulumi.StringPtrInput
	DomainUsername              pulumi.StringPtrInput
	Location                    pulumi.StringPtrInput
	NetworkConnectionName       pulumi.StringPtrInput
	NetworkingResourceGroupName pulumi.StringPtrInput
	OrganizationUnit            pulumi.StringPtrInput
	ResourceGroupName           pulumi.StringInput
	SubnetId                    pulumi.StringInput
	Tags                        pulumi.StringMapInput
}

func (NetworkConnectionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*networkConnectionArgs)(nil)).Elem()
}

type NetworkConnectionInput interface {
	pulumi.Input

	ToNetworkConnectionOutput() NetworkConnectionOutput
	ToNetworkConnectionOutputWithContext(ctx context.Context) NetworkConnectionOutput
}

func (*NetworkConnection) ElementType() reflect.Type {
	return reflect.TypeOf((**NetworkConnection)(nil)).Elem()
}

func (i *NetworkConnection) ToNetworkConnectionOutput() NetworkConnectionOutput {
	return i.ToNetworkConnectionOutputWithContext(context.Background())
}

func (i *NetworkConnection) ToNetworkConnectionOutputWithContext(ctx context.Context) NetworkConnectionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NetworkConnectionOutput)
}

type NetworkConnectionOutput struct{ *pulumi.OutputState }

func (NetworkConnectionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**NetworkConnection)(nil)).Elem()
}

func (o NetworkConnectionOutput) ToNetworkConnectionOutput() NetworkConnectionOutput {
	return o
}

func (o NetworkConnectionOutput) ToNetworkConnectionOutputWithContext(ctx context.Context) NetworkConnectionOutput {
	return o
}

func (o NetworkConnectionOutput) DomainJoinType() pulumi.StringOutput {
	return o.ApplyT(func(v *NetworkConnection) pulumi.StringOutput { return v.DomainJoinType }).(pulumi.StringOutput)
}

func (o NetworkConnectionOutput) DomainName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *NetworkConnection) pulumi.StringPtrOutput { return v.DomainName }).(pulumi.StringPtrOutput)
}

func (o NetworkConnectionOutput) DomainPassword() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *NetworkConnection) pulumi.StringPtrOutput { return v.DomainPassword }).(pulumi.StringPtrOutput)
}

func (o NetworkConnectionOutput) DomainUsername() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *NetworkConnection) pulumi.StringPtrOutput { return v.DomainUsername }).(pulumi.StringPtrOutput)
}

func (o NetworkConnectionOutput) HealthCheckStatus() pulumi.StringOutput {
	return o.ApplyT(func(v *NetworkConnection) pulumi.StringOutput { return v.HealthCheckStatus }).(pulumi.StringOutput)
}

func (o NetworkConnectionOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *NetworkConnection) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

func (o NetworkConnectionOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *NetworkConnection) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o NetworkConnectionOutput) NetworkingResourceGroupName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *NetworkConnection) pulumi.StringPtrOutput { return v.NetworkingResourceGroupName }).(pulumi.StringPtrOutput)
}

func (o NetworkConnectionOutput) OrganizationUnit() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *NetworkConnection) pulumi.StringPtrOutput { return v.OrganizationUnit }).(pulumi.StringPtrOutput)
}

func (o NetworkConnectionOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *NetworkConnection) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o NetworkConnectionOutput) SubnetId() pulumi.StringOutput {
	return o.ApplyT(func(v *NetworkConnection) pulumi.StringOutput { return v.SubnetId }).(pulumi.StringOutput)
}

func (o NetworkConnectionOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *NetworkConnection) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

func (o NetworkConnectionOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *NetworkConnection) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

func (o NetworkConnectionOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *NetworkConnection) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(NetworkConnectionOutput{})
}
