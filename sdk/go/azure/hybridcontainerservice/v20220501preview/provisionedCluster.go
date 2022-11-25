


package v20220501preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type ProvisionedCluster struct {
	pulumi.CustomResourceState

	ExtendedLocation ProvisionedClustersResponseResponseExtendedLocationPtrOutput `pulumi:"extendedLocation"`
	Identity         ProvisionedClusterIdentityResponsePtrOutput                  `pulumi:"identity"`
	Location         pulumi.StringOutput                                          `pulumi:"location"`
	Name             pulumi.StringOutput                                          `pulumi:"name"`
	Properties       ProvisionedClustersResponsePropertiesResponseOutput          `pulumi:"properties"`
	SystemData       SystemDataResponseOutput                                     `pulumi:"systemData"`
	Tags             pulumi.StringMapOutput                                       `pulumi:"tags"`
	Type             pulumi.StringOutput                                          `pulumi:"type"`
}


func NewProvisionedCluster(ctx *pulumi.Context,
	name string, args *ProvisionedClusterArgs, opts ...pulumi.ResourceOption) (*ProvisionedCluster, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.Properties != nil {
		args.Properties = args.Properties.ToProvisionedClustersAllPropertiesPtrOutput().ApplyT(func(v *ProvisionedClustersAllProperties) *ProvisionedClustersAllProperties { return v.Defaults() }).(ProvisionedClustersAllPropertiesPtrOutput)
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:hybridcontainerservice:ProvisionedCluster"),
		},
	})
	opts = append(opts, aliases)
	var resource ProvisionedCluster
	err := ctx.RegisterResource("azure-native:hybridcontainerservice/v20220501preview:ProvisionedCluster", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetProvisionedCluster(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ProvisionedClusterState, opts ...pulumi.ResourceOption) (*ProvisionedCluster, error) {
	var resource ProvisionedCluster
	err := ctx.ReadResource("azure-native:hybridcontainerservice/v20220501preview:ProvisionedCluster", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type provisionedClusterState struct {
}

type ProvisionedClusterState struct {
}

func (ProvisionedClusterState) ElementType() reflect.Type {
	return reflect.TypeOf((*provisionedClusterState)(nil)).Elem()
}

type provisionedClusterArgs struct {
	ExtendedLocation        *ProvisionedClustersExtendedLocation `pulumi:"extendedLocation"`
	Identity                *ProvisionedClusterIdentity          `pulumi:"identity"`
	Location                *string                              `pulumi:"location"`
	Properties              *ProvisionedClustersAllProperties    `pulumi:"properties"`
	ProvisionedClustersName *string                              `pulumi:"provisionedClustersName"`
	ResourceGroupName       string                               `pulumi:"resourceGroupName"`
	Tags                    map[string]string                    `pulumi:"tags"`
}


type ProvisionedClusterArgs struct {
	ExtendedLocation        ProvisionedClustersExtendedLocationPtrInput
	Identity                ProvisionedClusterIdentityPtrInput
	Location                pulumi.StringPtrInput
	Properties              ProvisionedClustersAllPropertiesPtrInput
	ProvisionedClustersName pulumi.StringPtrInput
	ResourceGroupName       pulumi.StringInput
	Tags                    pulumi.StringMapInput
}

func (ProvisionedClusterArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*provisionedClusterArgs)(nil)).Elem()
}

type ProvisionedClusterInput interface {
	pulumi.Input

	ToProvisionedClusterOutput() ProvisionedClusterOutput
	ToProvisionedClusterOutputWithContext(ctx context.Context) ProvisionedClusterOutput
}

func (*ProvisionedCluster) ElementType() reflect.Type {
	return reflect.TypeOf((**ProvisionedCluster)(nil)).Elem()
}

func (i *ProvisionedCluster) ToProvisionedClusterOutput() ProvisionedClusterOutput {
	return i.ToProvisionedClusterOutputWithContext(context.Background())
}

func (i *ProvisionedCluster) ToProvisionedClusterOutputWithContext(ctx context.Context) ProvisionedClusterOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ProvisionedClusterOutput)
}

type ProvisionedClusterOutput struct{ *pulumi.OutputState }

func (ProvisionedClusterOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ProvisionedCluster)(nil)).Elem()
}

func (o ProvisionedClusterOutput) ToProvisionedClusterOutput() ProvisionedClusterOutput {
	return o
}

func (o ProvisionedClusterOutput) ToProvisionedClusterOutputWithContext(ctx context.Context) ProvisionedClusterOutput {
	return o
}

func (o ProvisionedClusterOutput) ExtendedLocation() ProvisionedClustersResponseResponseExtendedLocationPtrOutput {
	return o.ApplyT(func(v *ProvisionedCluster) ProvisionedClustersResponseResponseExtendedLocationPtrOutput {
		return v.ExtendedLocation
	}).(ProvisionedClustersResponseResponseExtendedLocationPtrOutput)
}

func (o ProvisionedClusterOutput) Identity() ProvisionedClusterIdentityResponsePtrOutput {
	return o.ApplyT(func(v *ProvisionedCluster) ProvisionedClusterIdentityResponsePtrOutput { return v.Identity }).(ProvisionedClusterIdentityResponsePtrOutput)
}

func (o ProvisionedClusterOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *ProvisionedCluster) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

func (o ProvisionedClusterOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *ProvisionedCluster) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o ProvisionedClusterOutput) Properties() ProvisionedClustersResponsePropertiesResponseOutput {
	return o.ApplyT(func(v *ProvisionedCluster) ProvisionedClustersResponsePropertiesResponseOutput { return v.Properties }).(ProvisionedClustersResponsePropertiesResponseOutput)
}

func (o ProvisionedClusterOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *ProvisionedCluster) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

func (o ProvisionedClusterOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *ProvisionedCluster) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

func (o ProvisionedClusterOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *ProvisionedCluster) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(ProvisionedClusterOutput{})
}
