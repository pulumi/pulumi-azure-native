// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * Gets the properties of the specified Iot Connector FHIR destination.
 *
 * Uses Azure REST API version 2024-03-31.
 *
 * Other available API versions: 2022-10-01-preview, 2022-12-01, 2023-02-28, 2023-09-06, 2023-11-01, 2023-12-01, 2024-03-01, 2025-03-01-preview, 2025-04-01-preview. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native healthcareapis [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
 */
export function getIotConnectorFhirDestination(args: GetIotConnectorFhirDestinationArgs, opts?: pulumi.InvokeOptions): Promise<GetIotConnectorFhirDestinationResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("azure-native:healthcareapis:getIotConnectorFhirDestination", {
        "fhirDestinationName": args.fhirDestinationName,
        "iotConnectorName": args.iotConnectorName,
        "resourceGroupName": args.resourceGroupName,
        "workspaceName": args.workspaceName,
    }, opts);
}

export interface GetIotConnectorFhirDestinationArgs {
    /**
     * The name of IoT Connector FHIR destination resource.
     */
    fhirDestinationName: string;
    /**
     * The name of IoT Connector resource.
     */
    iotConnectorName: string;
    /**
     * The name of the resource group that contains the service instance.
     */
    resourceGroupName: string;
    /**
     * The name of workspace resource.
     */
    workspaceName: string;
}

/**
 * IoT Connector FHIR destination definition.
 */
export interface GetIotConnectorFhirDestinationResult {
    /**
     * The Azure API version of the resource.
     */
    readonly azureApiVersion: string;
    /**
     * An etag associated with the resource, used for optimistic concurrency when editing it.
     */
    readonly etag?: string;
    /**
     * FHIR Mappings
     */
    readonly fhirMapping: outputs.healthcareapis.IotMappingPropertiesResponse;
    /**
     * Fully qualified resource id of the FHIR service to connect to.
     */
    readonly fhirServiceResourceId: string;
    /**
     * The resource identifier.
     */
    readonly id: string;
    /**
     * The resource location.
     */
    readonly location?: string;
    /**
     * The resource name.
     */
    readonly name: string;
    /**
     * Determines how resource identity is resolved on the destination.
     */
    readonly resourceIdentityResolutionType: string;
    /**
     * Metadata pertaining to creation and last modification of the resource.
     */
    readonly systemData: outputs.healthcareapis.SystemDataResponse;
    /**
     * The resource type.
     */
    readonly type: string;
}
/**
 * Gets the properties of the specified Iot Connector FHIR destination.
 *
 * Uses Azure REST API version 2024-03-31.
 *
 * Other available API versions: 2022-10-01-preview, 2022-12-01, 2023-02-28, 2023-09-06, 2023-11-01, 2023-12-01, 2024-03-01, 2025-03-01-preview, 2025-04-01-preview. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native healthcareapis [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
 */
export function getIotConnectorFhirDestinationOutput(args: GetIotConnectorFhirDestinationOutputArgs, opts?: pulumi.InvokeOutputOptions): pulumi.Output<GetIotConnectorFhirDestinationResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invokeOutput("azure-native:healthcareapis:getIotConnectorFhirDestination", {
        "fhirDestinationName": args.fhirDestinationName,
        "iotConnectorName": args.iotConnectorName,
        "resourceGroupName": args.resourceGroupName,
        "workspaceName": args.workspaceName,
    }, opts);
}

export interface GetIotConnectorFhirDestinationOutputArgs {
    /**
     * The name of IoT Connector FHIR destination resource.
     */
    fhirDestinationName: pulumi.Input<string>;
    /**
     * The name of IoT Connector resource.
     */
    iotConnectorName: pulumi.Input<string>;
    /**
     * The name of the resource group that contains the service instance.
     */
    resourceGroupName: pulumi.Input<string>;
    /**
     * The name of workspace resource.
     */
    workspaceName: pulumi.Input<string>;
}
