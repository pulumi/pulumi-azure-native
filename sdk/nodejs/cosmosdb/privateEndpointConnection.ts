// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * A private endpoint connection
 *
 * Uses Azure REST API version 2024-11-15.
 *
 * Other available API versions: 2019-08-01-preview, 2021-01-15, 2021-03-01-preview, 2021-03-15, 2021-04-01-preview, 2021-04-15, 2021-05-15, 2021-06-15, 2021-07-01-preview, 2021-10-15, 2021-10-15-preview, 2021-11-15-preview, 2022-02-15-preview, 2022-05-15, 2022-05-15-preview, 2022-08-15, 2022-08-15-preview, 2022-11-15, 2022-11-15-preview, 2023-03-01-preview, 2023-03-15, 2023-03-15-preview, 2023-04-15, 2023-09-15, 2023-09-15-preview, 2023-11-15, 2023-11-15-preview, 2024-02-15-preview, 2024-05-15, 2024-05-15-preview, 2024-08-15, 2024-09-01-preview, 2024-12-01-preview, 2025-04-15, 2025-05-01-preview. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native cosmosdb [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
 */
export class PrivateEndpointConnection extends pulumi.CustomResource {
    /**
     * Get an existing PrivateEndpointConnection resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): PrivateEndpointConnection {
        return new PrivateEndpointConnection(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azure-native:cosmosdb:PrivateEndpointConnection';

    /**
     * Returns true if the given object is an instance of PrivateEndpointConnection.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is PrivateEndpointConnection {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === PrivateEndpointConnection.__pulumiType;
    }

    /**
     * The Azure API version of the resource.
     */
    public /*out*/ readonly azureApiVersion!: pulumi.Output<string>;
    /**
     * Group id of the private endpoint.
     */
    public readonly groupId!: pulumi.Output<string | undefined>;
    /**
     * The name of the resource
     */
    public /*out*/ readonly name!: pulumi.Output<string>;
    /**
     * Private endpoint which the connection belongs to.
     */
    public readonly privateEndpoint!: pulumi.Output<outputs.cosmosdb.PrivateEndpointPropertyResponse | undefined>;
    /**
     * Connection State of the Private Endpoint Connection.
     */
    public readonly privateLinkServiceConnectionState!: pulumi.Output<outputs.cosmosdb.PrivateLinkServiceConnectionStatePropertyResponse | undefined>;
    /**
     * Provisioning state of the private endpoint.
     */
    public readonly provisioningState!: pulumi.Output<string | undefined>;
    /**
     * The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
     */
    public /*out*/ readonly type!: pulumi.Output<string>;

    /**
     * Create a PrivateEndpointConnection resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: PrivateEndpointConnectionArgs, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.accountName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'accountName'");
            }
            if ((!args || args.resourceGroupName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'resourceGroupName'");
            }
            resourceInputs["accountName"] = args ? args.accountName : undefined;
            resourceInputs["groupId"] = args ? args.groupId : undefined;
            resourceInputs["privateEndpoint"] = args ? args.privateEndpoint : undefined;
            resourceInputs["privateEndpointConnectionName"] = args ? args.privateEndpointConnectionName : undefined;
            resourceInputs["privateLinkServiceConnectionState"] = args ? args.privateLinkServiceConnectionState : undefined;
            resourceInputs["provisioningState"] = args ? args.provisioningState : undefined;
            resourceInputs["resourceGroupName"] = args ? args.resourceGroupName : undefined;
            resourceInputs["azureApiVersion"] = undefined /*out*/;
            resourceInputs["name"] = undefined /*out*/;
            resourceInputs["type"] = undefined /*out*/;
        } else {
            resourceInputs["azureApiVersion"] = undefined /*out*/;
            resourceInputs["groupId"] = undefined /*out*/;
            resourceInputs["name"] = undefined /*out*/;
            resourceInputs["privateEndpoint"] = undefined /*out*/;
            resourceInputs["privateLinkServiceConnectionState"] = undefined /*out*/;
            resourceInputs["provisioningState"] = undefined /*out*/;
            resourceInputs["type"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        const aliasOpts = { aliases: [{ type: "azure-native:cosmosdb/v20190801preview:PrivateEndpointConnection" }, { type: "azure-native:cosmosdb/v20210115:PrivateEndpointConnection" }, { type: "azure-native:cosmosdb/v20210301preview:PrivateEndpointConnection" }, { type: "azure-native:cosmosdb/v20210315:PrivateEndpointConnection" }, { type: "azure-native:cosmosdb/v20210401preview:PrivateEndpointConnection" }, { type: "azure-native:cosmosdb/v20210415:PrivateEndpointConnection" }, { type: "azure-native:cosmosdb/v20210515:PrivateEndpointConnection" }, { type: "azure-native:cosmosdb/v20210615:PrivateEndpointConnection" }, { type: "azure-native:cosmosdb/v20210701preview:PrivateEndpointConnection" }, { type: "azure-native:cosmosdb/v20211015:PrivateEndpointConnection" }, { type: "azure-native:cosmosdb/v20211015preview:PrivateEndpointConnection" }, { type: "azure-native:cosmosdb/v20211115preview:PrivateEndpointConnection" }, { type: "azure-native:cosmosdb/v20220215preview:PrivateEndpointConnection" }, { type: "azure-native:cosmosdb/v20220515:PrivateEndpointConnection" }, { type: "azure-native:cosmosdb/v20220515preview:PrivateEndpointConnection" }, { type: "azure-native:cosmosdb/v20220815:PrivateEndpointConnection" }, { type: "azure-native:cosmosdb/v20220815preview:PrivateEndpointConnection" }, { type: "azure-native:cosmosdb/v20221115:PrivateEndpointConnection" }, { type: "azure-native:cosmosdb/v20221115preview:PrivateEndpointConnection" }, { type: "azure-native:cosmosdb/v20230301preview:PrivateEndpointConnection" }, { type: "azure-native:cosmosdb/v20230315:PrivateEndpointConnection" }, { type: "azure-native:cosmosdb/v20230315preview:PrivateEndpointConnection" }, { type: "azure-native:cosmosdb/v20230415:PrivateEndpointConnection" }, { type: "azure-native:cosmosdb/v20230915:PrivateEndpointConnection" }, { type: "azure-native:cosmosdb/v20230915preview:PrivateEndpointConnection" }, { type: "azure-native:cosmosdb/v20231115:PrivateEndpointConnection" }, { type: "azure-native:cosmosdb/v20231115preview:PrivateEndpointConnection" }, { type: "azure-native:cosmosdb/v20240215preview:PrivateEndpointConnection" }, { type: "azure-native:cosmosdb/v20240515:PrivateEndpointConnection" }, { type: "azure-native:cosmosdb/v20240515preview:PrivateEndpointConnection" }, { type: "azure-native:cosmosdb/v20240815:PrivateEndpointConnection" }, { type: "azure-native:cosmosdb/v20240901preview:PrivateEndpointConnection" }, { type: "azure-native:cosmosdb/v20241115:PrivateEndpointConnection" }, { type: "azure-native:cosmosdb/v20241201preview:PrivateEndpointConnection" }, { type: "azure-native:cosmosdb/v20250415:PrivateEndpointConnection" }, { type: "azure-native:cosmosdb/v20250501preview:PrivateEndpointConnection" }, { type: "azure-native:documentdb/v20230415:PrivateEndpointConnection" }, { type: "azure-native:documentdb/v20230915:PrivateEndpointConnection" }, { type: "azure-native:documentdb/v20230915preview:PrivateEndpointConnection" }, { type: "azure-native:documentdb/v20231115:PrivateEndpointConnection" }, { type: "azure-native:documentdb/v20231115preview:PrivateEndpointConnection" }, { type: "azure-native:documentdb/v20240215preview:PrivateEndpointConnection" }, { type: "azure-native:documentdb/v20240515:PrivateEndpointConnection" }, { type: "azure-native:documentdb/v20240515preview:PrivateEndpointConnection" }, { type: "azure-native:documentdb/v20240815:PrivateEndpointConnection" }, { type: "azure-native:documentdb/v20240901preview:PrivateEndpointConnection" }, { type: "azure-native:documentdb/v20241115:PrivateEndpointConnection" }, { type: "azure-native:documentdb/v20241201preview:PrivateEndpointConnection" }, { type: "azure-native:documentdb:PrivateEndpointConnection" }] };
        opts = pulumi.mergeOptions(opts, aliasOpts);
        super(PrivateEndpointConnection.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * The set of arguments for constructing a PrivateEndpointConnection resource.
 */
export interface PrivateEndpointConnectionArgs {
    /**
     * Cosmos DB database account name.
     */
    accountName: pulumi.Input<string>;
    /**
     * Group id of the private endpoint.
     */
    groupId?: pulumi.Input<string>;
    /**
     * Private endpoint which the connection belongs to.
     */
    privateEndpoint?: pulumi.Input<inputs.cosmosdb.PrivateEndpointPropertyArgs>;
    /**
     * The name of the private endpoint connection.
     */
    privateEndpointConnectionName?: pulumi.Input<string>;
    /**
     * Connection State of the Private Endpoint Connection.
     */
    privateLinkServiceConnectionState?: pulumi.Input<inputs.cosmosdb.PrivateLinkServiceConnectionStatePropertyArgs>;
    /**
     * Provisioning state of the private endpoint.
     */
    provisioningState?: pulumi.Input<string>;
    /**
     * The name of the resource group. The name is case insensitive.
     */
    resourceGroupName: pulumi.Input<string>;
}
