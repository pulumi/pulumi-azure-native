


package synapse

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupKustoPoolAttachedDatabaseConfiguration(ctx *pulumi.Context, args *LookupKustoPoolAttachedDatabaseConfigurationArgs, opts ...pulumi.InvokeOption) (*LookupKustoPoolAttachedDatabaseConfigurationResult, error) {
	var rv LookupKustoPoolAttachedDatabaseConfigurationResult
	err := ctx.Invoke("azure-native:synapse:getKustoPoolAttachedDatabaseConfiguration", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupKustoPoolAttachedDatabaseConfigurationArgs struct {
	AttachedDatabaseConfigurationName string `pulumi:"attachedDatabaseConfigurationName"`
	KustoPoolName                     string `pulumi:"kustoPoolName"`
	ResourceGroupName                 string `pulumi:"resourceGroupName"`
	WorkspaceName                     string `pulumi:"workspaceName"`
}


type LookupKustoPoolAttachedDatabaseConfigurationResult struct {
	AttachedDatabaseNames             []string                             `pulumi:"attachedDatabaseNames"`
	DatabaseName                      string                               `pulumi:"databaseName"`
	DefaultPrincipalsModificationKind string                               `pulumi:"defaultPrincipalsModificationKind"`
	Id                                string                               `pulumi:"id"`
	KustoPoolResourceId               string                               `pulumi:"kustoPoolResourceId"`
	Location                          *string                              `pulumi:"location"`
	Name                              string                               `pulumi:"name"`
	ProvisioningState                 string                               `pulumi:"provisioningState"`
	SystemData                        SystemDataResponse                   `pulumi:"systemData"`
	TableLevelSharingProperties       *TableLevelSharingPropertiesResponse `pulumi:"tableLevelSharingProperties"`
	Type                              string                               `pulumi:"type"`
}

func LookupKustoPoolAttachedDatabaseConfigurationOutput(ctx *pulumi.Context, args LookupKustoPoolAttachedDatabaseConfigurationOutputArgs, opts ...pulumi.InvokeOption) LookupKustoPoolAttachedDatabaseConfigurationResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupKustoPoolAttachedDatabaseConfigurationResult, error) {
			args := v.(LookupKustoPoolAttachedDatabaseConfigurationArgs)
			r, err := LookupKustoPoolAttachedDatabaseConfiguration(ctx, &args, opts...)
			var s LookupKustoPoolAttachedDatabaseConfigurationResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupKustoPoolAttachedDatabaseConfigurationResultOutput)
}

type LookupKustoPoolAttachedDatabaseConfigurationOutputArgs struct {
	AttachedDatabaseConfigurationName pulumi.StringInput `pulumi:"attachedDatabaseConfigurationName"`
	KustoPoolName                     pulumi.StringInput `pulumi:"kustoPoolName"`
	ResourceGroupName                 pulumi.StringInput `pulumi:"resourceGroupName"`
	WorkspaceName                     pulumi.StringInput `pulumi:"workspaceName"`
}

func (LookupKustoPoolAttachedDatabaseConfigurationOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupKustoPoolAttachedDatabaseConfigurationArgs)(nil)).Elem()
}


type LookupKustoPoolAttachedDatabaseConfigurationResultOutput struct{ *pulumi.OutputState }

func (LookupKustoPoolAttachedDatabaseConfigurationResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupKustoPoolAttachedDatabaseConfigurationResult)(nil)).Elem()
}

func (o LookupKustoPoolAttachedDatabaseConfigurationResultOutput) ToLookupKustoPoolAttachedDatabaseConfigurationResultOutput() LookupKustoPoolAttachedDatabaseConfigurationResultOutput {
	return o
}

func (o LookupKustoPoolAttachedDatabaseConfigurationResultOutput) ToLookupKustoPoolAttachedDatabaseConfigurationResultOutputWithContext(ctx context.Context) LookupKustoPoolAttachedDatabaseConfigurationResultOutput {
	return o
}

func (o LookupKustoPoolAttachedDatabaseConfigurationResultOutput) AttachedDatabaseNames() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupKustoPoolAttachedDatabaseConfigurationResult) []string { return v.AttachedDatabaseNames }).(pulumi.StringArrayOutput)
}

func (o LookupKustoPoolAttachedDatabaseConfigurationResultOutput) DatabaseName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupKustoPoolAttachedDatabaseConfigurationResult) string { return v.DatabaseName }).(pulumi.StringOutput)
}

func (o LookupKustoPoolAttachedDatabaseConfigurationResultOutput) DefaultPrincipalsModificationKind() pulumi.StringOutput {
	return o.ApplyT(func(v LookupKustoPoolAttachedDatabaseConfigurationResult) string {
		return v.DefaultPrincipalsModificationKind
	}).(pulumi.StringOutput)
}

func (o LookupKustoPoolAttachedDatabaseConfigurationResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupKustoPoolAttachedDatabaseConfigurationResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupKustoPoolAttachedDatabaseConfigurationResultOutput) KustoPoolResourceId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupKustoPoolAttachedDatabaseConfigurationResult) string { return v.KustoPoolResourceId }).(pulumi.StringOutput)
}

func (o LookupKustoPoolAttachedDatabaseConfigurationResultOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupKustoPoolAttachedDatabaseConfigurationResult) *string { return v.Location }).(pulumi.StringPtrOutput)
}

func (o LookupKustoPoolAttachedDatabaseConfigurationResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupKustoPoolAttachedDatabaseConfigurationResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupKustoPoolAttachedDatabaseConfigurationResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupKustoPoolAttachedDatabaseConfigurationResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o LookupKustoPoolAttachedDatabaseConfigurationResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupKustoPoolAttachedDatabaseConfigurationResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

func (o LookupKustoPoolAttachedDatabaseConfigurationResultOutput) TableLevelSharingProperties() TableLevelSharingPropertiesResponsePtrOutput {
	return o.ApplyT(func(v LookupKustoPoolAttachedDatabaseConfigurationResult) *TableLevelSharingPropertiesResponse {
		return v.TableLevelSharingProperties
	}).(TableLevelSharingPropertiesResponsePtrOutput)
}

func (o LookupKustoPoolAttachedDatabaseConfigurationResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupKustoPoolAttachedDatabaseConfigurationResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupKustoPoolAttachedDatabaseConfigurationResultOutput{})
}
