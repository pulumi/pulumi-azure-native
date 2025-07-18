// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * IoT Connector FHIR destination definition.
 *
 * Uses Azure REST API version 2024-03-31. In version 2.x of the Azure Native provider, it used API version 2023-02-28.
 *
 * Other available API versions: 2022-10-01-preview, 2022-12-01, 2023-02-28, 2023-09-06, 2023-11-01, 2023-12-01, 2024-03-01, 2025-03-01-preview, 2025-04-01-preview. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native healthcareapis [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
 */
export class IotConnectorFhirDestination extends pulumi.CustomResource {
    /**
     * Get an existing IotConnectorFhirDestination resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): IotConnectorFhirDestination {
        return new IotConnectorFhirDestination(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azure-native:healthcareapis:IotConnectorFhirDestination';

    /**
     * Returns true if the given object is an instance of IotConnectorFhirDestination.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is IotConnectorFhirDestination {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === IotConnectorFhirDestination.__pulumiType;
    }

    /**
     * The Azure API version of the resource.
     */
    public /*out*/ readonly azureApiVersion!: pulumi.Output<string>;
    /**
     * An etag associated with the resource, used for optimistic concurrency when editing it.
     */
    public /*out*/ readonly etag!: pulumi.Output<string | undefined>;
    /**
     * FHIR Mappings
     */
    public readonly fhirMapping!: pulumi.Output<outputs.healthcareapis.IotMappingPropertiesResponse>;
    /**
     * Fully qualified resource id of the FHIR service to connect to.
     */
    public readonly fhirServiceResourceId!: pulumi.Output<string>;
    /**
     * The resource location.
     */
    public readonly location!: pulumi.Output<string | undefined>;
    /**
     * The resource name.
     */
    public /*out*/ readonly name!: pulumi.Output<string>;
    /**
     * Determines how resource identity is resolved on the destination.
     */
    public readonly resourceIdentityResolutionType!: pulumi.Output<string>;
    /**
     * Metadata pertaining to creation and last modification of the resource.
     */
    public /*out*/ readonly systemData!: pulumi.Output<outputs.healthcareapis.SystemDataResponse>;
    /**
     * The resource type.
     */
    public /*out*/ readonly type!: pulumi.Output<string>;

    /**
     * Create a IotConnectorFhirDestination resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: IotConnectorFhirDestinationArgs, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.fhirMapping === undefined) && !opts.urn) {
                throw new Error("Missing required property 'fhirMapping'");
            }
            if ((!args || args.fhirServiceResourceId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'fhirServiceResourceId'");
            }
            if ((!args || args.iotConnectorName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'iotConnectorName'");
            }
            if ((!args || args.resourceGroupName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'resourceGroupName'");
            }
            if ((!args || args.resourceIdentityResolutionType === undefined) && !opts.urn) {
                throw new Error("Missing required property 'resourceIdentityResolutionType'");
            }
            if ((!args || args.workspaceName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'workspaceName'");
            }
            resourceInputs["fhirDestinationName"] = args ? args.fhirDestinationName : undefined;
            resourceInputs["fhirMapping"] = args ? args.fhirMapping : undefined;
            resourceInputs["fhirServiceResourceId"] = args ? args.fhirServiceResourceId : undefined;
            resourceInputs["iotConnectorName"] = args ? args.iotConnectorName : undefined;
            resourceInputs["location"] = args ? args.location : undefined;
            resourceInputs["resourceGroupName"] = args ? args.resourceGroupName : undefined;
            resourceInputs["resourceIdentityResolutionType"] = args ? args.resourceIdentityResolutionType : undefined;
            resourceInputs["workspaceName"] = args ? args.workspaceName : undefined;
            resourceInputs["azureApiVersion"] = undefined /*out*/;
            resourceInputs["etag"] = undefined /*out*/;
            resourceInputs["name"] = undefined /*out*/;
            resourceInputs["systemData"] = undefined /*out*/;
            resourceInputs["type"] = undefined /*out*/;
        } else {
            resourceInputs["azureApiVersion"] = undefined /*out*/;
            resourceInputs["etag"] = undefined /*out*/;
            resourceInputs["fhirMapping"] = undefined /*out*/;
            resourceInputs["fhirServiceResourceId"] = undefined /*out*/;
            resourceInputs["location"] = undefined /*out*/;
            resourceInputs["name"] = undefined /*out*/;
            resourceInputs["resourceIdentityResolutionType"] = undefined /*out*/;
            resourceInputs["systemData"] = undefined /*out*/;
            resourceInputs["type"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        const aliasOpts = { aliases: [{ type: "azure-native:healthcareapis/v20210601preview:IotConnectorFhirDestination" }, { type: "azure-native:healthcareapis/v20211101:IotConnectorFhirDestination" }, { type: "azure-native:healthcareapis/v20220131preview:IotConnectorFhirDestination" }, { type: "azure-native:healthcareapis/v20220515:IotConnectorFhirDestination" }, { type: "azure-native:healthcareapis/v20220601:IotConnectorFhirDestination" }, { type: "azure-native:healthcareapis/v20221001preview:IotConnectorFhirDestination" }, { type: "azure-native:healthcareapis/v20221201:IotConnectorFhirDestination" }, { type: "azure-native:healthcareapis/v20230228:IotConnectorFhirDestination" }, { type: "azure-native:healthcareapis/v20230906:IotConnectorFhirDestination" }, { type: "azure-native:healthcareapis/v20231101:IotConnectorFhirDestination" }, { type: "azure-native:healthcareapis/v20231201:IotConnectorFhirDestination" }, { type: "azure-native:healthcareapis/v20240301:IotConnectorFhirDestination" }, { type: "azure-native:healthcareapis/v20240331:IotConnectorFhirDestination" }, { type: "azure-native:healthcareapis/v20250301preview:IotConnectorFhirDestination" }, { type: "azure-native:healthcareapis/v20250401preview:IotConnectorFhirDestination" }] };
        opts = pulumi.mergeOptions(opts, aliasOpts);
        super(IotConnectorFhirDestination.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * The set of arguments for constructing a IotConnectorFhirDestination resource.
 */
export interface IotConnectorFhirDestinationArgs {
    /**
     * The name of IoT Connector FHIR destination resource.
     */
    fhirDestinationName?: pulumi.Input<string>;
    /**
     * FHIR Mappings
     */
    fhirMapping: pulumi.Input<inputs.healthcareapis.IotMappingPropertiesArgs>;
    /**
     * Fully qualified resource id of the FHIR service to connect to.
     */
    fhirServiceResourceId: pulumi.Input<string>;
    /**
     * The name of IoT Connector resource.
     */
    iotConnectorName: pulumi.Input<string>;
    /**
     * The resource location.
     */
    location?: pulumi.Input<string>;
    /**
     * The name of the resource group that contains the service instance.
     */
    resourceGroupName: pulumi.Input<string>;
    /**
     * Determines how resource identity is resolved on the destination.
     */
    resourceIdentityResolutionType: pulumi.Input<string | enums.healthcareapis.IotIdentityResolutionType>;
    /**
     * The name of workspace resource.
     */
    workspaceName: pulumi.Input<string>;
}
