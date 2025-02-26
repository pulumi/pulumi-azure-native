// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as enums from "../../types/enums";
import * as utilities from "../../utilities";

/**
 * Details of a specific cloud account connector
 */
export function getConnector(args: GetConnectorArgs, opts?: pulumi.InvokeOptions): Promise<GetConnectorResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("azure-native:security/v20200101preview:getConnector", {
        "connectorName": args.connectorName,
    }, opts);
}

export interface GetConnectorArgs {
    /**
     * Name of the cloud account connector
     */
    connectorName: string;
}

/**
 * The connector setting
 */
export interface GetConnectorResult {
    /**
     * Settings for authentication management, these settings are relevant only for the cloud connector.
     */
    readonly authenticationDetails?: outputs.security.v20200101preview.AwAssumeRoleAuthenticationDetailsPropertiesResponse | outputs.security.v20200101preview.AwsCredsAuthenticationDetailsPropertiesResponse | outputs.security.v20200101preview.GcpCredentialsDetailsPropertiesResponse;
    /**
     * Settings for hybrid compute management. These settings are relevant only for Arc autoProvision (Hybrid Compute).
     */
    readonly hybridComputeSettings?: outputs.security.v20200101preview.HybridComputeSettingsPropertiesResponse;
    /**
     * Resource Id
     */
    readonly id: string;
    /**
     * Resource name
     */
    readonly name: string;
    /**
     * Resource type
     */
    readonly type: string;
}
/**
 * Details of a specific cloud account connector
 */
export function getConnectorOutput(args: GetConnectorOutputArgs, opts?: pulumi.InvokeOutputOptions): pulumi.Output<GetConnectorResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invokeOutput("azure-native:security/v20200101preview:getConnector", {
        "connectorName": args.connectorName,
    }, opts);
}

export interface GetConnectorOutputArgs {
    /**
     * Name of the cloud account connector
     */
    connectorName: pulumi.Input<string>;
}
