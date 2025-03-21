// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as enums from "../../types/enums";
import * as utilities from "../../utilities";

/**
 * Gets the details of the Diagnostic for an API specified by its identifier.
 */
export function getApiDiagnostic(args: GetApiDiagnosticArgs, opts?: pulumi.InvokeOptions): Promise<GetApiDiagnosticResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("azure-native:apimanagement/v20190101:getApiDiagnostic", {
        "apiId": args.apiId,
        "diagnosticId": args.diagnosticId,
        "resourceGroupName": args.resourceGroupName,
        "serviceName": args.serviceName,
    }, opts);
}

export interface GetApiDiagnosticArgs {
    /**
     * API identifier. Must be unique in the current API Management service instance.
     */
    apiId: string;
    /**
     * Diagnostic identifier. Must be unique in the current API Management service instance.
     */
    diagnosticId: string;
    /**
     * The name of the resource group.
     */
    resourceGroupName: string;
    /**
     * The name of the API Management service.
     */
    serviceName: string;
}

/**
 * Diagnostic details.
 */
export interface GetApiDiagnosticResult {
    /**
     * Specifies for what type of messages sampling settings should not apply.
     */
    readonly alwaysLog?: string;
    /**
     * Diagnostic settings for incoming/outgoing HTTP messages to the Backend
     */
    readonly backend?: outputs.apimanagement.v20190101.PipelineDiagnosticSettingsResponse;
    /**
     * Whether to process Correlation Headers coming to Api Management Service. Only applicable to Application Insights diagnostics. Default is true.
     */
    readonly enableHttpCorrelationHeaders?: boolean;
    /**
     * Diagnostic settings for incoming/outgoing HTTP messages to the Gateway.
     */
    readonly frontend?: outputs.apimanagement.v20190101.PipelineDiagnosticSettingsResponse;
    /**
     * Sets correlation protocol to use for Application Insights diagnostics.
     */
    readonly httpCorrelationProtocol?: string;
    /**
     * Resource ID.
     */
    readonly id: string;
    /**
     * Resource Id of a target logger.
     */
    readonly loggerId: string;
    /**
     * Resource name.
     */
    readonly name: string;
    /**
     * Sampling settings for Diagnostic.
     */
    readonly sampling?: outputs.apimanagement.v20190101.SamplingSettingsResponse;
    /**
     * Resource type for API Management resource.
     */
    readonly type: string;
    /**
     * The verbosity level applied to traces emitted by trace policies.
     */
    readonly verbosity?: string;
}
/**
 * Gets the details of the Diagnostic for an API specified by its identifier.
 */
export function getApiDiagnosticOutput(args: GetApiDiagnosticOutputArgs, opts?: pulumi.InvokeOutputOptions): pulumi.Output<GetApiDiagnosticResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invokeOutput("azure-native:apimanagement/v20190101:getApiDiagnostic", {
        "apiId": args.apiId,
        "diagnosticId": args.diagnosticId,
        "resourceGroupName": args.resourceGroupName,
        "serviceName": args.serviceName,
    }, opts);
}

export interface GetApiDiagnosticOutputArgs {
    /**
     * API identifier. Must be unique in the current API Management service instance.
     */
    apiId: pulumi.Input<string>;
    /**
     * Diagnostic identifier. Must be unique in the current API Management service instance.
     */
    diagnosticId: pulumi.Input<string>;
    /**
     * The name of the resource group.
     */
    resourceGroupName: pulumi.Input<string>;
    /**
     * The name of the API Management service.
     */
    serviceName: pulumi.Input<string>;
}
