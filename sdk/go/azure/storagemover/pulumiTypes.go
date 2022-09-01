


package storagemover

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type AgentPropertiesResponseErrorDetails struct {
	Code    *string `pulumi:"code"`
	Message *string `pulumi:"message"`
}

type AgentPropertiesResponseErrorDetailsOutput struct{ *pulumi.OutputState }

func (AgentPropertiesResponseErrorDetailsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AgentPropertiesResponseErrorDetails)(nil)).Elem()
}

func (o AgentPropertiesResponseErrorDetailsOutput) ToAgentPropertiesResponseErrorDetailsOutput() AgentPropertiesResponseErrorDetailsOutput {
	return o
}

func (o AgentPropertiesResponseErrorDetailsOutput) ToAgentPropertiesResponseErrorDetailsOutputWithContext(ctx context.Context) AgentPropertiesResponseErrorDetailsOutput {
	return o
}

func (o AgentPropertiesResponseErrorDetailsOutput) Code() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AgentPropertiesResponseErrorDetails) *string { return v.Code }).(pulumi.StringPtrOutput)
}

func (o AgentPropertiesResponseErrorDetailsOutput) Message() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AgentPropertiesResponseErrorDetails) *string { return v.Message }).(pulumi.StringPtrOutput)
}

type AzureStorageBlobContainerEndpointProperties struct {
	BlobContainerName        string  `pulumi:"blobContainerName"`
	Description              *string `pulumi:"description"`
	EndpointType             string  `pulumi:"endpointType"`
	StorageAccountResourceId string  `pulumi:"storageAccountResourceId"`
}

type AzureStorageBlobContainerEndpointPropertiesResponse struct {
	BlobContainerName        string  `pulumi:"blobContainerName"`
	Description              *string `pulumi:"description"`
	EndpointType             string  `pulumi:"endpointType"`
	ProvisioningState        string  `pulumi:"provisioningState"`
	StorageAccountResourceId string  `pulumi:"storageAccountResourceId"`
}

type NfsMountEndpointProperties struct {
	Description  *string `pulumi:"description"`
	EndpointType string  `pulumi:"endpointType"`
	Export       string  `pulumi:"export"`
	Host         string  `pulumi:"host"`
	NfsVersion   *string `pulumi:"nfsVersion"`
}

type NfsMountEndpointPropertiesResponse struct {
	Description       *string `pulumi:"description"`
	EndpointType      string  `pulumi:"endpointType"`
	Export            string  `pulumi:"export"`
	Host              string  `pulumi:"host"`
	NfsVersion        *string `pulumi:"nfsVersion"`
	ProvisioningState string  `pulumi:"provisioningState"`
}

type SystemDataResponse struct {
	CreatedAt          *string `pulumi:"createdAt"`
	CreatedBy          *string `pulumi:"createdBy"`
	CreatedByType      *string `pulumi:"createdByType"`
	LastModifiedAt     *string `pulumi:"lastModifiedAt"`
	LastModifiedBy     *string `pulumi:"lastModifiedBy"`
	LastModifiedByType *string `pulumi:"lastModifiedByType"`
}

type SystemDataResponseOutput struct{ *pulumi.OutputState }

func (SystemDataResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SystemDataResponse)(nil)).Elem()
}

func (o SystemDataResponseOutput) ToSystemDataResponseOutput() SystemDataResponseOutput {
	return o
}

func (o SystemDataResponseOutput) ToSystemDataResponseOutputWithContext(ctx context.Context) SystemDataResponseOutput {
	return o
}

func (o SystemDataResponseOutput) CreatedAt() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SystemDataResponse) *string { return v.CreatedAt }).(pulumi.StringPtrOutput)
}

func (o SystemDataResponseOutput) CreatedBy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SystemDataResponse) *string { return v.CreatedBy }).(pulumi.StringPtrOutput)
}

func (o SystemDataResponseOutput) CreatedByType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SystemDataResponse) *string { return v.CreatedByType }).(pulumi.StringPtrOutput)
}

func (o SystemDataResponseOutput) LastModifiedAt() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SystemDataResponse) *string { return v.LastModifiedAt }).(pulumi.StringPtrOutput)
}

func (o SystemDataResponseOutput) LastModifiedBy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SystemDataResponse) *string { return v.LastModifiedBy }).(pulumi.StringPtrOutput)
}

func (o SystemDataResponseOutput) LastModifiedByType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SystemDataResponse) *string { return v.LastModifiedByType }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(AgentPropertiesResponseErrorDetailsOutput{})
	pulumi.RegisterOutputType(SystemDataResponseOutput{})
}
