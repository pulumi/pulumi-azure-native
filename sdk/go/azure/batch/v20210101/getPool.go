


package v20210101

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupPool(ctx *pulumi.Context, args *LookupPoolArgs, opts ...pulumi.InvokeOption) (*LookupPoolResult, error) {
	var rv LookupPoolResult
	err := ctx.Invoke("azure-native:batch/v20210101:getPool", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupPoolArgs struct {
	AccountName       string `pulumi:"accountName"`
	PoolName          string `pulumi:"poolName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
}


type LookupPoolResult struct {
	AllocationState                 string                                `pulumi:"allocationState"`
	AllocationStateTransitionTime   string                                `pulumi:"allocationStateTransitionTime"`
	ApplicationLicenses             []string                              `pulumi:"applicationLicenses"`
	ApplicationPackages             []ApplicationPackageReferenceResponse `pulumi:"applicationPackages"`
	AutoScaleRun                    AutoScaleRunResponse                  `pulumi:"autoScaleRun"`
	Certificates                    []CertificateReferenceResponse        `pulumi:"certificates"`
	CreationTime                    string                                `pulumi:"creationTime"`
	CurrentDedicatedNodes           int                                   `pulumi:"currentDedicatedNodes"`
	CurrentLowPriorityNodes         int                                   `pulumi:"currentLowPriorityNodes"`
	DeploymentConfiguration         *DeploymentConfigurationResponse      `pulumi:"deploymentConfiguration"`
	DisplayName                     *string                               `pulumi:"displayName"`
	Etag                            string                                `pulumi:"etag"`
	Id                              string                                `pulumi:"id"`
	Identity                        *BatchPoolIdentityResponse            `pulumi:"identity"`
	InterNodeCommunication          *string                               `pulumi:"interNodeCommunication"`
	LastModified                    string                                `pulumi:"lastModified"`
	Metadata                        []MetadataItemResponse                `pulumi:"metadata"`
	MountConfiguration              []MountConfigurationResponse          `pulumi:"mountConfiguration"`
	Name                            string                                `pulumi:"name"`
	NetworkConfiguration            *NetworkConfigurationResponse         `pulumi:"networkConfiguration"`
	ProvisioningState               string                                `pulumi:"provisioningState"`
	ProvisioningStateTransitionTime string                                `pulumi:"provisioningStateTransitionTime"`
	ResizeOperationStatus           ResizeOperationStatusResponse         `pulumi:"resizeOperationStatus"`
	ScaleSettings                   *ScaleSettingsResponse                `pulumi:"scaleSettings"`
	StartTask                       *StartTaskResponse                    `pulumi:"startTask"`
	TaskSchedulingPolicy            *TaskSchedulingPolicyResponse         `pulumi:"taskSchedulingPolicy"`
	TaskSlotsPerNode                *int                                  `pulumi:"taskSlotsPerNode"`
	Type                            string                                `pulumi:"type"`
	UserAccounts                    []UserAccountResponse                 `pulumi:"userAccounts"`
	VmSize                          *string                               `pulumi:"vmSize"`
}

func LookupPoolOutput(ctx *pulumi.Context, args LookupPoolOutputArgs, opts ...pulumi.InvokeOption) LookupPoolResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupPoolResult, error) {
			args := v.(LookupPoolArgs)
			r, err := LookupPool(ctx, &args, opts...)
			var s LookupPoolResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupPoolResultOutput)
}

type LookupPoolOutputArgs struct {
	AccountName       pulumi.StringInput `pulumi:"accountName"`
	PoolName          pulumi.StringInput `pulumi:"poolName"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupPoolOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupPoolArgs)(nil)).Elem()
}


type LookupPoolResultOutput struct{ *pulumi.OutputState }

func (LookupPoolResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupPoolResult)(nil)).Elem()
}

func (o LookupPoolResultOutput) ToLookupPoolResultOutput() LookupPoolResultOutput {
	return o
}

func (o LookupPoolResultOutput) ToLookupPoolResultOutputWithContext(ctx context.Context) LookupPoolResultOutput {
	return o
}

func (o LookupPoolResultOutput) AllocationState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPoolResult) string { return v.AllocationState }).(pulumi.StringOutput)
}

func (o LookupPoolResultOutput) AllocationStateTransitionTime() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPoolResult) string { return v.AllocationStateTransitionTime }).(pulumi.StringOutput)
}

func (o LookupPoolResultOutput) ApplicationLicenses() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupPoolResult) []string { return v.ApplicationLicenses }).(pulumi.StringArrayOutput)
}

func (o LookupPoolResultOutput) ApplicationPackages() ApplicationPackageReferenceResponseArrayOutput {
	return o.ApplyT(func(v LookupPoolResult) []ApplicationPackageReferenceResponse { return v.ApplicationPackages }).(ApplicationPackageReferenceResponseArrayOutput)
}

func (o LookupPoolResultOutput) AutoScaleRun() AutoScaleRunResponseOutput {
	return o.ApplyT(func(v LookupPoolResult) AutoScaleRunResponse { return v.AutoScaleRun }).(AutoScaleRunResponseOutput)
}

func (o LookupPoolResultOutput) Certificates() CertificateReferenceResponseArrayOutput {
	return o.ApplyT(func(v LookupPoolResult) []CertificateReferenceResponse { return v.Certificates }).(CertificateReferenceResponseArrayOutput)
}

func (o LookupPoolResultOutput) CreationTime() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPoolResult) string { return v.CreationTime }).(pulumi.StringOutput)
}

func (o LookupPoolResultOutput) CurrentDedicatedNodes() pulumi.IntOutput {
	return o.ApplyT(func(v LookupPoolResult) int { return v.CurrentDedicatedNodes }).(pulumi.IntOutput)
}

func (o LookupPoolResultOutput) CurrentLowPriorityNodes() pulumi.IntOutput {
	return o.ApplyT(func(v LookupPoolResult) int { return v.CurrentLowPriorityNodes }).(pulumi.IntOutput)
}

func (o LookupPoolResultOutput) DeploymentConfiguration() DeploymentConfigurationResponsePtrOutput {
	return o.ApplyT(func(v LookupPoolResult) *DeploymentConfigurationResponse { return v.DeploymentConfiguration }).(DeploymentConfigurationResponsePtrOutput)
}

func (o LookupPoolResultOutput) DisplayName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupPoolResult) *string { return v.DisplayName }).(pulumi.StringPtrOutput)
}

func (o LookupPoolResultOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPoolResult) string { return v.Etag }).(pulumi.StringOutput)
}

func (o LookupPoolResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPoolResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupPoolResultOutput) Identity() BatchPoolIdentityResponsePtrOutput {
	return o.ApplyT(func(v LookupPoolResult) *BatchPoolIdentityResponse { return v.Identity }).(BatchPoolIdentityResponsePtrOutput)
}

func (o LookupPoolResultOutput) InterNodeCommunication() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupPoolResult) *string { return v.InterNodeCommunication }).(pulumi.StringPtrOutput)
}

func (o LookupPoolResultOutput) LastModified() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPoolResult) string { return v.LastModified }).(pulumi.StringOutput)
}

func (o LookupPoolResultOutput) Metadata() MetadataItemResponseArrayOutput {
	return o.ApplyT(func(v LookupPoolResult) []MetadataItemResponse { return v.Metadata }).(MetadataItemResponseArrayOutput)
}

func (o LookupPoolResultOutput) MountConfiguration() MountConfigurationResponseArrayOutput {
	return o.ApplyT(func(v LookupPoolResult) []MountConfigurationResponse { return v.MountConfiguration }).(MountConfigurationResponseArrayOutput)
}

func (o LookupPoolResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPoolResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupPoolResultOutput) NetworkConfiguration() NetworkConfigurationResponsePtrOutput {
	return o.ApplyT(func(v LookupPoolResult) *NetworkConfigurationResponse { return v.NetworkConfiguration }).(NetworkConfigurationResponsePtrOutput)
}

func (o LookupPoolResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPoolResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o LookupPoolResultOutput) ProvisioningStateTransitionTime() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPoolResult) string { return v.ProvisioningStateTransitionTime }).(pulumi.StringOutput)
}

func (o LookupPoolResultOutput) ResizeOperationStatus() ResizeOperationStatusResponseOutput {
	return o.ApplyT(func(v LookupPoolResult) ResizeOperationStatusResponse { return v.ResizeOperationStatus }).(ResizeOperationStatusResponseOutput)
}

func (o LookupPoolResultOutput) ScaleSettings() ScaleSettingsResponsePtrOutput {
	return o.ApplyT(func(v LookupPoolResult) *ScaleSettingsResponse { return v.ScaleSettings }).(ScaleSettingsResponsePtrOutput)
}

func (o LookupPoolResultOutput) StartTask() StartTaskResponsePtrOutput {
	return o.ApplyT(func(v LookupPoolResult) *StartTaskResponse { return v.StartTask }).(StartTaskResponsePtrOutput)
}

func (o LookupPoolResultOutput) TaskSchedulingPolicy() TaskSchedulingPolicyResponsePtrOutput {
	return o.ApplyT(func(v LookupPoolResult) *TaskSchedulingPolicyResponse { return v.TaskSchedulingPolicy }).(TaskSchedulingPolicyResponsePtrOutput)
}

func (o LookupPoolResultOutput) TaskSlotsPerNode() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupPoolResult) *int { return v.TaskSlotsPerNode }).(pulumi.IntPtrOutput)
}

func (o LookupPoolResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPoolResult) string { return v.Type }).(pulumi.StringOutput)
}

func (o LookupPoolResultOutput) UserAccounts() UserAccountResponseArrayOutput {
	return o.ApplyT(func(v LookupPoolResult) []UserAccountResponse { return v.UserAccounts }).(UserAccountResponseArrayOutput)
}

func (o LookupPoolResultOutput) VmSize() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupPoolResult) *string { return v.VmSize }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupPoolResultOutput{})
}
