


package v20191109

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupAttachedDatabaseConfiguration(ctx *pulumi.Context, args *LookupAttachedDatabaseConfigurationArgs, opts ...pulumi.InvokeOption) (*LookupAttachedDatabaseConfigurationResult, error) {
	var rv LookupAttachedDatabaseConfigurationResult
	err := ctx.Invoke("azure-native:kusto/v20191109:getAttachedDatabaseConfiguration", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupAttachedDatabaseConfigurationArgs struct {
	AttachedDatabaseConfigurationName string `pulumi:"attachedDatabaseConfigurationName"`
	ClusterName                       string `pulumi:"clusterName"`
	ResourceGroupName                 string `pulumi:"resourceGroupName"`
}


type LookupAttachedDatabaseConfigurationResult struct {
	AttachedDatabaseNames             []string `pulumi:"attachedDatabaseNames"`
	ClusterResourceId                 string   `pulumi:"clusterResourceId"`
	DatabaseName                      string   `pulumi:"databaseName"`
	DefaultPrincipalsModificationKind string   `pulumi:"defaultPrincipalsModificationKind"`
	Id                                string   `pulumi:"id"`
	Location                          *string  `pulumi:"location"`
	Name                              string   `pulumi:"name"`
	ProvisioningState                 string   `pulumi:"provisioningState"`
	Type                              string   `pulumi:"type"`
}

func LookupAttachedDatabaseConfigurationOutput(ctx *pulumi.Context, args LookupAttachedDatabaseConfigurationOutputArgs, opts ...pulumi.InvokeOption) LookupAttachedDatabaseConfigurationResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupAttachedDatabaseConfigurationResult, error) {
			args := v.(LookupAttachedDatabaseConfigurationArgs)
			r, err := LookupAttachedDatabaseConfiguration(ctx, &args, opts...)
			var s LookupAttachedDatabaseConfigurationResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupAttachedDatabaseConfigurationResultOutput)
}

type LookupAttachedDatabaseConfigurationOutputArgs struct {
	AttachedDatabaseConfigurationName pulumi.StringInput `pulumi:"attachedDatabaseConfigurationName"`
	ClusterName                       pulumi.StringInput `pulumi:"clusterName"`
	ResourceGroupName                 pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupAttachedDatabaseConfigurationOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupAttachedDatabaseConfigurationArgs)(nil)).Elem()
}


type LookupAttachedDatabaseConfigurationResultOutput struct{ *pulumi.OutputState }

func (LookupAttachedDatabaseConfigurationResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupAttachedDatabaseConfigurationResult)(nil)).Elem()
}

func (o LookupAttachedDatabaseConfigurationResultOutput) ToLookupAttachedDatabaseConfigurationResultOutput() LookupAttachedDatabaseConfigurationResultOutput {
	return o
}

func (o LookupAttachedDatabaseConfigurationResultOutput) ToLookupAttachedDatabaseConfigurationResultOutputWithContext(ctx context.Context) LookupAttachedDatabaseConfigurationResultOutput {
	return o
}

func (o LookupAttachedDatabaseConfigurationResultOutput) AttachedDatabaseNames() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupAttachedDatabaseConfigurationResult) []string { return v.AttachedDatabaseNames }).(pulumi.StringArrayOutput)
}

func (o LookupAttachedDatabaseConfigurationResultOutput) ClusterResourceId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAttachedDatabaseConfigurationResult) string { return v.ClusterResourceId }).(pulumi.StringOutput)
}

func (o LookupAttachedDatabaseConfigurationResultOutput) DatabaseName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAttachedDatabaseConfigurationResult) string { return v.DatabaseName }).(pulumi.StringOutput)
}

func (o LookupAttachedDatabaseConfigurationResultOutput) DefaultPrincipalsModificationKind() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAttachedDatabaseConfigurationResult) string { return v.DefaultPrincipalsModificationKind }).(pulumi.StringOutput)
}

func (o LookupAttachedDatabaseConfigurationResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAttachedDatabaseConfigurationResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupAttachedDatabaseConfigurationResultOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupAttachedDatabaseConfigurationResult) *string { return v.Location }).(pulumi.StringPtrOutput)
}

func (o LookupAttachedDatabaseConfigurationResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAttachedDatabaseConfigurationResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupAttachedDatabaseConfigurationResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAttachedDatabaseConfigurationResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o LookupAttachedDatabaseConfigurationResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAttachedDatabaseConfigurationResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupAttachedDatabaseConfigurationResultOutput{})
}
