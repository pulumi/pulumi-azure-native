// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * Represents a catalog.
 *
 * Uses Azure REST API version 2024-02-01. In version 2.x of the Azure Native provider, it used API version 2023-04-01.
 *
 * Other available API versions: 2023-04-01, 2023-08-01-preview, 2023-10-01-preview, 2024-05-01-preview, 2024-06-01-preview, 2024-07-01-preview, 2024-08-01-preview, 2024-10-01-preview, 2025-02-01, 2025-04-01-preview, 2025-07-01-preview. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native devcenter [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
 */
export class Catalog extends pulumi.CustomResource {
    /**
     * Get an existing Catalog resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): Catalog {
        return new Catalog(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azure-native:devcenter:Catalog';

    /**
     * Returns true if the given object is an instance of Catalog.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Catalog {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Catalog.__pulumiType;
    }

    /**
     * Properties for an Azure DevOps catalog type.
     */
    public readonly adoGit!: pulumi.Output<outputs.devcenter.GitCatalogResponse | undefined>;
    /**
     * The Azure API version of the resource.
     */
    public /*out*/ readonly azureApiVersion!: pulumi.Output<string>;
    /**
     * The connection state of the catalog.
     */
    public /*out*/ readonly connectionState!: pulumi.Output<string>;
    /**
     * Properties for a GitHub catalog type.
     */
    public readonly gitHub!: pulumi.Output<outputs.devcenter.GitCatalogResponse | undefined>;
    /**
     * When the catalog was last connected.
     */
    public /*out*/ readonly lastConnectionTime!: pulumi.Output<string>;
    /**
     * Stats of the latest synchronization.
     */
    public /*out*/ readonly lastSyncStats!: pulumi.Output<outputs.devcenter.SyncStatsResponse>;
    /**
     * When the catalog was last synced.
     */
    public /*out*/ readonly lastSyncTime!: pulumi.Output<string>;
    /**
     * The name of the resource
     */
    public /*out*/ readonly name!: pulumi.Output<string>;
    /**
     * The provisioning state of the resource.
     */
    public /*out*/ readonly provisioningState!: pulumi.Output<string>;
    /**
     * The synchronization state of the catalog.
     */
    public /*out*/ readonly syncState!: pulumi.Output<string>;
    /**
     * Indicates the type of sync that is configured for the catalog.
     */
    public readonly syncType!: pulumi.Output<string | undefined>;
    /**
     * Azure Resource Manager metadata containing createdBy and modifiedBy information.
     */
    public /*out*/ readonly systemData!: pulumi.Output<outputs.devcenter.SystemDataResponse>;
    /**
     * Resource tags.
     */
    public readonly tags!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
     */
    public /*out*/ readonly type!: pulumi.Output<string>;

    /**
     * Create a Catalog resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: CatalogArgs, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.devCenterName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'devCenterName'");
            }
            if ((!args || args.resourceGroupName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'resourceGroupName'");
            }
            resourceInputs["adoGit"] = args ? args.adoGit : undefined;
            resourceInputs["catalogName"] = args ? args.catalogName : undefined;
            resourceInputs["devCenterName"] = args ? args.devCenterName : undefined;
            resourceInputs["gitHub"] = args ? args.gitHub : undefined;
            resourceInputs["resourceGroupName"] = args ? args.resourceGroupName : undefined;
            resourceInputs["syncType"] = args ? args.syncType : undefined;
            resourceInputs["tags"] = args ? args.tags : undefined;
            resourceInputs["azureApiVersion"] = undefined /*out*/;
            resourceInputs["connectionState"] = undefined /*out*/;
            resourceInputs["lastConnectionTime"] = undefined /*out*/;
            resourceInputs["lastSyncStats"] = undefined /*out*/;
            resourceInputs["lastSyncTime"] = undefined /*out*/;
            resourceInputs["name"] = undefined /*out*/;
            resourceInputs["provisioningState"] = undefined /*out*/;
            resourceInputs["syncState"] = undefined /*out*/;
            resourceInputs["systemData"] = undefined /*out*/;
            resourceInputs["type"] = undefined /*out*/;
        } else {
            resourceInputs["adoGit"] = undefined /*out*/;
            resourceInputs["azureApiVersion"] = undefined /*out*/;
            resourceInputs["connectionState"] = undefined /*out*/;
            resourceInputs["gitHub"] = undefined /*out*/;
            resourceInputs["lastConnectionTime"] = undefined /*out*/;
            resourceInputs["lastSyncStats"] = undefined /*out*/;
            resourceInputs["lastSyncTime"] = undefined /*out*/;
            resourceInputs["name"] = undefined /*out*/;
            resourceInputs["provisioningState"] = undefined /*out*/;
            resourceInputs["syncState"] = undefined /*out*/;
            resourceInputs["syncType"] = undefined /*out*/;
            resourceInputs["systemData"] = undefined /*out*/;
            resourceInputs["tags"] = undefined /*out*/;
            resourceInputs["type"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        const aliasOpts = { aliases: [{ type: "azure-native:devcenter/v20220801preview:Catalog" }, { type: "azure-native:devcenter/v20220901preview:Catalog" }, { type: "azure-native:devcenter/v20221012preview:Catalog" }, { type: "azure-native:devcenter/v20221111preview:Catalog" }, { type: "azure-native:devcenter/v20230101preview:Catalog" }, { type: "azure-native:devcenter/v20230401:Catalog" }, { type: "azure-native:devcenter/v20230801preview:Catalog" }, { type: "azure-native:devcenter/v20231001preview:Catalog" }, { type: "azure-native:devcenter/v20240201:Catalog" }, { type: "azure-native:devcenter/v20240501preview:Catalog" }, { type: "azure-native:devcenter/v20240601preview:Catalog" }, { type: "azure-native:devcenter/v20240701preview:Catalog" }, { type: "azure-native:devcenter/v20240801preview:Catalog" }, { type: "azure-native:devcenter/v20241001preview:Catalog" }, { type: "azure-native:devcenter/v20250201:Catalog" }, { type: "azure-native:devcenter/v20250401preview:Catalog" }, { type: "azure-native:devcenter/v20250701preview:Catalog" }] };
        opts = pulumi.mergeOptions(opts, aliasOpts);
        super(Catalog.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * The set of arguments for constructing a Catalog resource.
 */
export interface CatalogArgs {
    /**
     * Properties for an Azure DevOps catalog type.
     */
    adoGit?: pulumi.Input<inputs.devcenter.GitCatalogArgs>;
    /**
     * The name of the Catalog.
     */
    catalogName?: pulumi.Input<string>;
    /**
     * The name of the devcenter.
     */
    devCenterName: pulumi.Input<string>;
    /**
     * Properties for a GitHub catalog type.
     */
    gitHub?: pulumi.Input<inputs.devcenter.GitCatalogArgs>;
    /**
     * The name of the resource group. The name is case insensitive.
     */
    resourceGroupName: pulumi.Input<string>;
    /**
     * Indicates the type of sync that is configured for the catalog.
     */
    syncType?: pulumi.Input<string | enums.devcenter.CatalogSyncType>;
    /**
     * Resource tags.
     */
    tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
}
