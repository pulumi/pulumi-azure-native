


package v20200901

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupCloudEndpoint(ctx *pulumi.Context, args *LookupCloudEndpointArgs, opts ...pulumi.InvokeOption) (*LookupCloudEndpointResult, error) {
	var rv LookupCloudEndpointResult
	err := ctx.Invoke("azure-native:storagesync/v20200901:getCloudEndpoint", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupCloudEndpointArgs struct {
	CloudEndpointName      string `pulumi:"cloudEndpointName"`
	ResourceGroupName      string `pulumi:"resourceGroupName"`
	StorageSyncServiceName string `pulumi:"storageSyncServiceName"`
	SyncGroupName          string `pulumi:"syncGroupName"`
}


type LookupCloudEndpointResult struct {
	AzureFileShareName       *string                                      `pulumi:"azureFileShareName"`
	BackupEnabled            string                                       `pulumi:"backupEnabled"`
	ChangeEnumerationStatus  CloudEndpointChangeEnumerationStatusResponse `pulumi:"changeEnumerationStatus"`
	FriendlyName             *string                                      `pulumi:"friendlyName"`
	Id                       string                                       `pulumi:"id"`
	LastOperationName        *string                                      `pulumi:"lastOperationName"`
	LastWorkflowId           *string                                      `pulumi:"lastWorkflowId"`
	Name                     string                                       `pulumi:"name"`
	PartnershipId            *string                                      `pulumi:"partnershipId"`
	ProvisioningState        *string                                      `pulumi:"provisioningState"`
	StorageAccountResourceId *string                                      `pulumi:"storageAccountResourceId"`
	StorageAccountTenantId   *string                                      `pulumi:"storageAccountTenantId"`
	SystemData               SystemDataResponse                           `pulumi:"systemData"`
	Type                     string                                       `pulumi:"type"`
}

func LookupCloudEndpointOutput(ctx *pulumi.Context, args LookupCloudEndpointOutputArgs, opts ...pulumi.InvokeOption) LookupCloudEndpointResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupCloudEndpointResult, error) {
			args := v.(LookupCloudEndpointArgs)
			r, err := LookupCloudEndpoint(ctx, &args, opts...)
			var s LookupCloudEndpointResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupCloudEndpointResultOutput)
}

type LookupCloudEndpointOutputArgs struct {
	CloudEndpointName      pulumi.StringInput `pulumi:"cloudEndpointName"`
	ResourceGroupName      pulumi.StringInput `pulumi:"resourceGroupName"`
	StorageSyncServiceName pulumi.StringInput `pulumi:"storageSyncServiceName"`
	SyncGroupName          pulumi.StringInput `pulumi:"syncGroupName"`
}

func (LookupCloudEndpointOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupCloudEndpointArgs)(nil)).Elem()
}


type LookupCloudEndpointResultOutput struct{ *pulumi.OutputState }

func (LookupCloudEndpointResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupCloudEndpointResult)(nil)).Elem()
}

func (o LookupCloudEndpointResultOutput) ToLookupCloudEndpointResultOutput() LookupCloudEndpointResultOutput {
	return o
}

func (o LookupCloudEndpointResultOutput) ToLookupCloudEndpointResultOutputWithContext(ctx context.Context) LookupCloudEndpointResultOutput {
	return o
}

func (o LookupCloudEndpointResultOutput) AzureFileShareName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupCloudEndpointResult) *string { return v.AzureFileShareName }).(pulumi.StringPtrOutput)
}

func (o LookupCloudEndpointResultOutput) BackupEnabled() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCloudEndpointResult) string { return v.BackupEnabled }).(pulumi.StringOutput)
}

func (o LookupCloudEndpointResultOutput) ChangeEnumerationStatus() CloudEndpointChangeEnumerationStatusResponseOutput {
	return o.ApplyT(func(v LookupCloudEndpointResult) CloudEndpointChangeEnumerationStatusResponse {
		return v.ChangeEnumerationStatus
	}).(CloudEndpointChangeEnumerationStatusResponseOutput)
}

func (o LookupCloudEndpointResultOutput) FriendlyName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupCloudEndpointResult) *string { return v.FriendlyName }).(pulumi.StringPtrOutput)
}

func (o LookupCloudEndpointResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCloudEndpointResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupCloudEndpointResultOutput) LastOperationName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupCloudEndpointResult) *string { return v.LastOperationName }).(pulumi.StringPtrOutput)
}

func (o LookupCloudEndpointResultOutput) LastWorkflowId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupCloudEndpointResult) *string { return v.LastWorkflowId }).(pulumi.StringPtrOutput)
}

func (o LookupCloudEndpointResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCloudEndpointResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupCloudEndpointResultOutput) PartnershipId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupCloudEndpointResult) *string { return v.PartnershipId }).(pulumi.StringPtrOutput)
}

func (o LookupCloudEndpointResultOutput) ProvisioningState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupCloudEndpointResult) *string { return v.ProvisioningState }).(pulumi.StringPtrOutput)
}

func (o LookupCloudEndpointResultOutput) StorageAccountResourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupCloudEndpointResult) *string { return v.StorageAccountResourceId }).(pulumi.StringPtrOutput)
}

func (o LookupCloudEndpointResultOutput) StorageAccountTenantId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupCloudEndpointResult) *string { return v.StorageAccountTenantId }).(pulumi.StringPtrOutput)
}

func (o LookupCloudEndpointResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupCloudEndpointResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

func (o LookupCloudEndpointResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCloudEndpointResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupCloudEndpointResultOutput{})
}
