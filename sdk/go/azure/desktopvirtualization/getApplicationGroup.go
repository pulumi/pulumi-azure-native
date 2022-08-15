


package desktopvirtualization

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupApplicationGroup(ctx *pulumi.Context, args *LookupApplicationGroupArgs, opts ...pulumi.InvokeOption) (*LookupApplicationGroupResult, error) {
	var rv LookupApplicationGroupResult
	err := ctx.Invoke("azure-native:desktopvirtualization:getApplicationGroup", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupApplicationGroupArgs struct {
	ApplicationGroupName string `pulumi:"applicationGroupName"`
	ResourceGroupName    string `pulumi:"resourceGroupName"`
}


type LookupApplicationGroupResult struct {
	ApplicationGroupType string                                               `pulumi:"applicationGroupType"`
	CloudPcResource      bool                                                 `pulumi:"cloudPcResource"`
	Description          *string                                              `pulumi:"description"`
	Etag                 string                                               `pulumi:"etag"`
	FriendlyName         *string                                              `pulumi:"friendlyName"`
	HostPoolArmPath      string                                               `pulumi:"hostPoolArmPath"`
	Id                   string                                               `pulumi:"id"`
	Identity             *ResourceModelWithAllowedPropertySetResponseIdentity `pulumi:"identity"`
	Kind                 *string                                              `pulumi:"kind"`
	Location             *string                                              `pulumi:"location"`
	ManagedBy            *string                                              `pulumi:"managedBy"`
	MigrationRequest     *MigrationRequestPropertiesResponse                  `pulumi:"migrationRequest"`
	Name                 string                                               `pulumi:"name"`
	ObjectId             string                                               `pulumi:"objectId"`
	Plan                 *ResourceModelWithAllowedPropertySetResponsePlan     `pulumi:"plan"`
	Sku                  *ResourceModelWithAllowedPropertySetResponseSku      `pulumi:"sku"`
	Tags                 map[string]string                                    `pulumi:"tags"`
	Type                 string                                               `pulumi:"type"`
	WorkspaceArmPath     string                                               `pulumi:"workspaceArmPath"`
}

func LookupApplicationGroupOutput(ctx *pulumi.Context, args LookupApplicationGroupOutputArgs, opts ...pulumi.InvokeOption) LookupApplicationGroupResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupApplicationGroupResult, error) {
			args := v.(LookupApplicationGroupArgs)
			r, err := LookupApplicationGroup(ctx, &args, opts...)
			var s LookupApplicationGroupResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupApplicationGroupResultOutput)
}

type LookupApplicationGroupOutputArgs struct {
	ApplicationGroupName pulumi.StringInput `pulumi:"applicationGroupName"`
	ResourceGroupName    pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupApplicationGroupOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupApplicationGroupArgs)(nil)).Elem()
}


type LookupApplicationGroupResultOutput struct{ *pulumi.OutputState }

func (LookupApplicationGroupResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupApplicationGroupResult)(nil)).Elem()
}

func (o LookupApplicationGroupResultOutput) ToLookupApplicationGroupResultOutput() LookupApplicationGroupResultOutput {
	return o
}

func (o LookupApplicationGroupResultOutput) ToLookupApplicationGroupResultOutputWithContext(ctx context.Context) LookupApplicationGroupResultOutput {
	return o
}

func (o LookupApplicationGroupResultOutput) ApplicationGroupType() pulumi.StringOutput {
	return o.ApplyT(func(v LookupApplicationGroupResult) string { return v.ApplicationGroupType }).(pulumi.StringOutput)
}

func (o LookupApplicationGroupResultOutput) CloudPcResource() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupApplicationGroupResult) bool { return v.CloudPcResource }).(pulumi.BoolOutput)
}

func (o LookupApplicationGroupResultOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupApplicationGroupResult) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o LookupApplicationGroupResultOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v LookupApplicationGroupResult) string { return v.Etag }).(pulumi.StringOutput)
}

func (o LookupApplicationGroupResultOutput) FriendlyName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupApplicationGroupResult) *string { return v.FriendlyName }).(pulumi.StringPtrOutput)
}

func (o LookupApplicationGroupResultOutput) HostPoolArmPath() pulumi.StringOutput {
	return o.ApplyT(func(v LookupApplicationGroupResult) string { return v.HostPoolArmPath }).(pulumi.StringOutput)
}

func (o LookupApplicationGroupResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupApplicationGroupResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupApplicationGroupResultOutput) Identity() ResourceModelWithAllowedPropertySetResponseIdentityPtrOutput {
	return o.ApplyT(func(v LookupApplicationGroupResult) *ResourceModelWithAllowedPropertySetResponseIdentity {
		return v.Identity
	}).(ResourceModelWithAllowedPropertySetResponseIdentityPtrOutput)
}

func (o LookupApplicationGroupResultOutput) Kind() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupApplicationGroupResult) *string { return v.Kind }).(pulumi.StringPtrOutput)
}

func (o LookupApplicationGroupResultOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupApplicationGroupResult) *string { return v.Location }).(pulumi.StringPtrOutput)
}

func (o LookupApplicationGroupResultOutput) ManagedBy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupApplicationGroupResult) *string { return v.ManagedBy }).(pulumi.StringPtrOutput)
}

func (o LookupApplicationGroupResultOutput) MigrationRequest() MigrationRequestPropertiesResponsePtrOutput {
	return o.ApplyT(func(v LookupApplicationGroupResult) *MigrationRequestPropertiesResponse { return v.MigrationRequest }).(MigrationRequestPropertiesResponsePtrOutput)
}

func (o LookupApplicationGroupResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupApplicationGroupResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupApplicationGroupResultOutput) ObjectId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupApplicationGroupResult) string { return v.ObjectId }).(pulumi.StringOutput)
}

func (o LookupApplicationGroupResultOutput) Plan() ResourceModelWithAllowedPropertySetResponsePlanPtrOutput {
	return o.ApplyT(func(v LookupApplicationGroupResult) *ResourceModelWithAllowedPropertySetResponsePlan { return v.Plan }).(ResourceModelWithAllowedPropertySetResponsePlanPtrOutput)
}

func (o LookupApplicationGroupResultOutput) Sku() ResourceModelWithAllowedPropertySetResponseSkuPtrOutput {
	return o.ApplyT(func(v LookupApplicationGroupResult) *ResourceModelWithAllowedPropertySetResponseSku { return v.Sku }).(ResourceModelWithAllowedPropertySetResponseSkuPtrOutput)
}

func (o LookupApplicationGroupResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupApplicationGroupResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o LookupApplicationGroupResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupApplicationGroupResult) string { return v.Type }).(pulumi.StringOutput)
}

func (o LookupApplicationGroupResultOutput) WorkspaceArmPath() pulumi.StringOutput {
	return o.ApplyT(func(v LookupApplicationGroupResult) string { return v.WorkspaceArmPath }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupApplicationGroupResultOutput{})
}
