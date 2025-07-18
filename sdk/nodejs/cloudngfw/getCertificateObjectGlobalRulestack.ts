// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * Get a CertificateObjectGlobalRulestackResource
 *
 * Uses Azure REST API version 2025-02-06-preview.
 *
 * Other available API versions: 2023-09-01, 2023-10-10-preview, 2024-01-19-preview, 2024-02-07-preview, 2025-05-23. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native cloudngfw [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
 */
export function getCertificateObjectGlobalRulestack(args: GetCertificateObjectGlobalRulestackArgs, opts?: pulumi.InvokeOptions): Promise<GetCertificateObjectGlobalRulestackResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("azure-native:cloudngfw:getCertificateObjectGlobalRulestack", {
        "globalRulestackName": args.globalRulestackName,
        "name": args.name,
    }, opts);
}

export interface GetCertificateObjectGlobalRulestackArgs {
    /**
     * GlobalRulestack resource name
     */
    globalRulestackName: string;
    /**
     * certificate name
     */
    name: string;
}

/**
 * GlobalRulestack Certificate Object
 */
export interface GetCertificateObjectGlobalRulestackResult {
    /**
     * comment for this object
     */
    readonly auditComment?: string;
    /**
     * The Azure API version of the resource.
     */
    readonly azureApiVersion: string;
    /**
     * use certificate self signed
     */
    readonly certificateSelfSigned: string;
    /**
     * Resource Id of certificate signer, to be populated only when certificateSelfSigned is false
     */
    readonly certificateSignerResourceId?: string;
    /**
     * user description for this object
     */
    readonly description?: string;
    /**
     * read only string representing last create or update
     */
    readonly etag?: string;
    /**
     * Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
     */
    readonly id: string;
    /**
     * The name of the resource
     */
    readonly name: string;
    /**
     * Provisioning state of the resource.
     */
    readonly provisioningState: string;
    /**
     * Azure Resource Manager metadata containing createdBy and modifiedBy information.
     */
    readonly systemData: outputs.cloudngfw.SystemDataResponse;
    /**
     * The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
     */
    readonly type: string;
}
/**
 * Get a CertificateObjectGlobalRulestackResource
 *
 * Uses Azure REST API version 2025-02-06-preview.
 *
 * Other available API versions: 2023-09-01, 2023-10-10-preview, 2024-01-19-preview, 2024-02-07-preview, 2025-05-23. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native cloudngfw [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
 */
export function getCertificateObjectGlobalRulestackOutput(args: GetCertificateObjectGlobalRulestackOutputArgs, opts?: pulumi.InvokeOutputOptions): pulumi.Output<GetCertificateObjectGlobalRulestackResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invokeOutput("azure-native:cloudngfw:getCertificateObjectGlobalRulestack", {
        "globalRulestackName": args.globalRulestackName,
        "name": args.name,
    }, opts);
}

export interface GetCertificateObjectGlobalRulestackOutputArgs {
    /**
     * GlobalRulestack resource name
     */
    globalRulestackName: pulumi.Input<string>;
    /**
     * certificate name
     */
    name: pulumi.Input<string>;
}
