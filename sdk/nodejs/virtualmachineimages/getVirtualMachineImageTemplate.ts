// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * Get information about a virtual machine image template
 *
 * Uses Azure REST API version 2024-02-01.
 *
 * Other available API versions: 2022-07-01, 2023-07-01. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native virtualmachineimages [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
 */
export function getVirtualMachineImageTemplate(args: GetVirtualMachineImageTemplateArgs, opts?: pulumi.InvokeOptions): Promise<GetVirtualMachineImageTemplateResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("azure-native:virtualmachineimages:getVirtualMachineImageTemplate", {
        "imageTemplateName": args.imageTemplateName,
        "resourceGroupName": args.resourceGroupName,
    }, opts);
}

export interface GetVirtualMachineImageTemplateArgs {
    /**
     * The name of the image Template
     */
    imageTemplateName: string;
    /**
     * The name of the resource group.
     */
    resourceGroupName: string;
}

/**
 * Image template is an ARM resource managed by Microsoft.VirtualMachineImages provider
 */
export interface GetVirtualMachineImageTemplateResult {
    /**
     * Indicates whether or not to automatically run the image template build on template creation or update.
     */
    readonly autoRun?: outputs.virtualmachineimages.ImageTemplateAutoRunResponse;
    /**
     * The Azure API version of the resource.
     */
    readonly azureApiVersion: string;
    /**
     * Maximum duration to wait while building the image template (includes all customizations, optimization, validations, and distributions). Omit or specify 0 to use the default (4 hours).
     */
    readonly buildTimeoutInMinutes?: number;
    /**
     * Specifies the properties used to describe the customization steps of the image, like Image source etc
     */
    readonly customize?: (outputs.virtualmachineimages.ImageTemplateFileCustomizerResponse | outputs.virtualmachineimages.ImageTemplatePowerShellCustomizerResponse | outputs.virtualmachineimages.ImageTemplateRestartCustomizerResponse | outputs.virtualmachineimages.ImageTemplateShellCustomizerResponse | outputs.virtualmachineimages.ImageTemplateWindowsUpdateCustomizerResponse)[];
    /**
     * The distribution targets where the image output needs to go to.
     */
    readonly distribute: (outputs.virtualmachineimages.ImageTemplateManagedImageDistributorResponse | outputs.virtualmachineimages.ImageTemplateSharedImageDistributorResponse | outputs.virtualmachineimages.ImageTemplateVhdDistributorResponse)[];
    /**
     * Error handling options upon a build failure
     */
    readonly errorHandling?: outputs.virtualmachineimages.ImageTemplatePropertiesResponseErrorHandling;
    /**
     * The staging resource group id in the same subscription as the image template that will be used to build the image. This read-only field differs from 'stagingResourceGroup' only if the value specified in the 'stagingResourceGroup' field is empty.
     */
    readonly exactStagingResourceGroup: string;
    /**
     * Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
     */
    readonly id: string;
    /**
     * The identity of the image template, if configured.
     */
    readonly identity: outputs.virtualmachineimages.ImageTemplateIdentityResponse;
    /**
     * State of 'run' that is currently executing or was last executed.
     */
    readonly lastRunStatus: outputs.virtualmachineimages.ImageTemplateLastRunStatusResponse;
    /**
     * The geo-location where the resource lives
     */
    readonly location: string;
    /**
     * Tags that will be applied to the resource group and/or resources created by the service.
     */
    readonly managedResourceTags?: {[key: string]: string};
    /**
     * The name of the resource
     */
    readonly name: string;
    /**
     * Specifies optimization to be performed on image.
     */
    readonly optimize?: outputs.virtualmachineimages.ImageTemplatePropertiesResponseOptimize;
    /**
     * Provisioning error, if any
     */
    readonly provisioningError: outputs.virtualmachineimages.ProvisioningErrorResponse;
    /**
     * Provisioning state of the resource
     */
    readonly provisioningState: string;
    /**
     * Specifies the properties used to describe the source image.
     */
    readonly source: outputs.virtualmachineimages.ImageTemplateManagedImageSourceResponse | outputs.virtualmachineimages.ImageTemplatePlatformImageSourceResponse | outputs.virtualmachineimages.ImageTemplateSharedImageVersionSourceResponse;
    /**
     * The staging resource group id in the same subscription as the image template that will be used to build the image. If this field is empty, a resource group with a random name will be created. If the resource group specified in this field doesn't exist, it will be created with the same name. If the resource group specified exists, it must be empty and in the same region as the image template. The resource group created will be deleted during template deletion if this field is empty or the resource group specified doesn't exist, but if the resource group specified exists the resources created in the resource group will be deleted during template deletion and the resource group itself will remain.
     */
    readonly stagingResourceGroup?: string;
    /**
     * Azure Resource Manager metadata containing createdBy and modifiedBy information.
     */
    readonly systemData: outputs.virtualmachineimages.SystemDataResponse;
    /**
     * Resource tags.
     */
    readonly tags?: {[key: string]: string};
    /**
     * The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
     */
    readonly type: string;
    /**
     * Configuration options and list of validations to be performed on the resulting image.
     */
    readonly validate?: outputs.virtualmachineimages.ImageTemplatePropertiesResponseValidate;
    /**
     * Describes how virtual machine is set up to build images
     */
    readonly vmProfile?: outputs.virtualmachineimages.ImageTemplateVmProfileResponse;
}
/**
 * Get information about a virtual machine image template
 *
 * Uses Azure REST API version 2024-02-01.
 *
 * Other available API versions: 2022-07-01, 2023-07-01. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native virtualmachineimages [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
 */
export function getVirtualMachineImageTemplateOutput(args: GetVirtualMachineImageTemplateOutputArgs, opts?: pulumi.InvokeOutputOptions): pulumi.Output<GetVirtualMachineImageTemplateResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invokeOutput("azure-native:virtualmachineimages:getVirtualMachineImageTemplate", {
        "imageTemplateName": args.imageTemplateName,
        "resourceGroupName": args.resourceGroupName,
    }, opts);
}

export interface GetVirtualMachineImageTemplateOutputArgs {
    /**
     * The name of the image Template
     */
    imageTemplateName: pulumi.Input<string>;
    /**
     * The name of the resource group.
     */
    resourceGroupName: pulumi.Input<string>;
}
