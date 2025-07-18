// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * The configuration store along with all resource properties. The Configuration Store will have all information to begin utilizing it.
 *
 * Uses Azure REST API version 2024-05-01. In version 2.x of the Azure Native provider, it used API version 2023-03-01.
 *
 * Other available API versions: 2023-03-01, 2023-08-01-preview, 2023-09-01-preview, 2024-06-01, 2024-06-15-preview, 2025-02-01-preview. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native appconfiguration [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
 */
export class ConfigurationStore extends pulumi.CustomResource {
    /**
     * Get an existing ConfigurationStore resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): ConfigurationStore {
        return new ConfigurationStore(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azure-native:appconfiguration:ConfigurationStore';

    /**
     * Returns true if the given object is an instance of ConfigurationStore.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is ConfigurationStore {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === ConfigurationStore.__pulumiType;
    }

    /**
     * The Azure API version of the resource.
     */
    public /*out*/ readonly azureApiVersion!: pulumi.Output<string>;
    /**
     * The creation date of configuration store.
     */
    public /*out*/ readonly creationDate!: pulumi.Output<string>;
    /**
     * Property specifying the configuration of data plane proxy for Azure Resource Manager (ARM).
     */
    public readonly dataPlaneProxy!: pulumi.Output<outputs.appconfiguration.DataPlaneProxyPropertiesResponse | undefined>;
    /**
     * Disables all authentication methods other than AAD authentication.
     */
    public readonly disableLocalAuth!: pulumi.Output<boolean | undefined>;
    /**
     * Property specifying whether protection against purge is enabled for this configuration store.
     */
    public readonly enablePurgeProtection!: pulumi.Output<boolean | undefined>;
    /**
     * The encryption settings of the configuration store.
     */
    public readonly encryption!: pulumi.Output<outputs.appconfiguration.EncryptionPropertiesResponse | undefined>;
    /**
     * The DNS endpoint where the configuration store API will be available.
     */
    public /*out*/ readonly endpoint!: pulumi.Output<string>;
    /**
     * The managed identity information, if configured.
     */
    public readonly identity!: pulumi.Output<outputs.appconfiguration.ResourceIdentityResponse | undefined>;
    /**
     * The geo-location where the resource lives
     */
    public readonly location!: pulumi.Output<string>;
    /**
     * The name of the resource
     */
    public /*out*/ readonly name!: pulumi.Output<string>;
    /**
     * The list of private endpoint connections that are set up for this resource.
     */
    public /*out*/ readonly privateEndpointConnections!: pulumi.Output<outputs.appconfiguration.PrivateEndpointConnectionReferenceResponse[]>;
    /**
     * The provisioning state of the configuration store.
     */
    public /*out*/ readonly provisioningState!: pulumi.Output<string>;
    /**
     * Control permission for data plane traffic coming from public networks while private endpoint is enabled.
     */
    public readonly publicNetworkAccess!: pulumi.Output<string | undefined>;
    /**
     * The sku of the configuration store.
     */
    public readonly sku!: pulumi.Output<outputs.appconfiguration.SkuResponse>;
    /**
     * The amount of time in days that the configuration store will be retained when it is soft deleted.
     */
    public readonly softDeleteRetentionInDays!: pulumi.Output<number | undefined>;
    /**
     * Resource system metadata.
     */
    public /*out*/ readonly systemData!: pulumi.Output<outputs.appconfiguration.SystemDataResponse>;
    /**
     * Resource tags.
     */
    public readonly tags!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
     */
    public /*out*/ readonly type!: pulumi.Output<string>;

    /**
     * Create a ConfigurationStore resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ConfigurationStoreArgs, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.resourceGroupName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'resourceGroupName'");
            }
            if ((!args || args.sku === undefined) && !opts.urn) {
                throw new Error("Missing required property 'sku'");
            }
            resourceInputs["configStoreName"] = args ? args.configStoreName : undefined;
            resourceInputs["createMode"] = args ? args.createMode : undefined;
            resourceInputs["dataPlaneProxy"] = args ? args.dataPlaneProxy : undefined;
            resourceInputs["disableLocalAuth"] = (args ? args.disableLocalAuth : undefined) ?? false;
            resourceInputs["enablePurgeProtection"] = (args ? args.enablePurgeProtection : undefined) ?? false;
            resourceInputs["encryption"] = args ? args.encryption : undefined;
            resourceInputs["identity"] = args ? args.identity : undefined;
            resourceInputs["location"] = args ? args.location : undefined;
            resourceInputs["publicNetworkAccess"] = args ? args.publicNetworkAccess : undefined;
            resourceInputs["resourceGroupName"] = args ? args.resourceGroupName : undefined;
            resourceInputs["sku"] = args ? args.sku : undefined;
            resourceInputs["softDeleteRetentionInDays"] = (args ? args.softDeleteRetentionInDays : undefined) ?? 7;
            resourceInputs["tags"] = args ? args.tags : undefined;
            resourceInputs["azureApiVersion"] = undefined /*out*/;
            resourceInputs["creationDate"] = undefined /*out*/;
            resourceInputs["endpoint"] = undefined /*out*/;
            resourceInputs["name"] = undefined /*out*/;
            resourceInputs["privateEndpointConnections"] = undefined /*out*/;
            resourceInputs["provisioningState"] = undefined /*out*/;
            resourceInputs["systemData"] = undefined /*out*/;
            resourceInputs["type"] = undefined /*out*/;
        } else {
            resourceInputs["azureApiVersion"] = undefined /*out*/;
            resourceInputs["creationDate"] = undefined /*out*/;
            resourceInputs["dataPlaneProxy"] = undefined /*out*/;
            resourceInputs["disableLocalAuth"] = undefined /*out*/;
            resourceInputs["enablePurgeProtection"] = undefined /*out*/;
            resourceInputs["encryption"] = undefined /*out*/;
            resourceInputs["endpoint"] = undefined /*out*/;
            resourceInputs["identity"] = undefined /*out*/;
            resourceInputs["location"] = undefined /*out*/;
            resourceInputs["name"] = undefined /*out*/;
            resourceInputs["privateEndpointConnections"] = undefined /*out*/;
            resourceInputs["provisioningState"] = undefined /*out*/;
            resourceInputs["publicNetworkAccess"] = undefined /*out*/;
            resourceInputs["sku"] = undefined /*out*/;
            resourceInputs["softDeleteRetentionInDays"] = undefined /*out*/;
            resourceInputs["systemData"] = undefined /*out*/;
            resourceInputs["tags"] = undefined /*out*/;
            resourceInputs["type"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        const aliasOpts = { aliases: [{ type: "azure-native:appconfiguration/v20190201preview:ConfigurationStore" }, { type: "azure-native:appconfiguration/v20191001:ConfigurationStore" }, { type: "azure-native:appconfiguration/v20191101preview:ConfigurationStore" }, { type: "azure-native:appconfiguration/v20200601:ConfigurationStore" }, { type: "azure-native:appconfiguration/v20200701preview:ConfigurationStore" }, { type: "azure-native:appconfiguration/v20210301preview:ConfigurationStore" }, { type: "azure-native:appconfiguration/v20211001preview:ConfigurationStore" }, { type: "azure-native:appconfiguration/v20220301preview:ConfigurationStore" }, { type: "azure-native:appconfiguration/v20220501:ConfigurationStore" }, { type: "azure-native:appconfiguration/v20230301:ConfigurationStore" }, { type: "azure-native:appconfiguration/v20230801preview:ConfigurationStore" }, { type: "azure-native:appconfiguration/v20230901preview:ConfigurationStore" }, { type: "azure-native:appconfiguration/v20240501:ConfigurationStore" }, { type: "azure-native:appconfiguration/v20240601:ConfigurationStore" }, { type: "azure-native:appconfiguration/v20240615preview:ConfigurationStore" }, { type: "azure-native:appconfiguration/v20250201preview:ConfigurationStore" }] };
        opts = pulumi.mergeOptions(opts, aliasOpts);
        super(ConfigurationStore.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * The set of arguments for constructing a ConfigurationStore resource.
 */
export interface ConfigurationStoreArgs {
    /**
     * The name of the configuration store.
     */
    configStoreName?: pulumi.Input<string>;
    /**
     * Indicates whether the configuration store need to be recovered.
     */
    createMode?: pulumi.Input<enums.appconfiguration.CreateMode>;
    /**
     * Property specifying the configuration of data plane proxy for Azure Resource Manager (ARM).
     */
    dataPlaneProxy?: pulumi.Input<inputs.appconfiguration.DataPlaneProxyPropertiesArgs>;
    /**
     * Disables all authentication methods other than AAD authentication.
     */
    disableLocalAuth?: pulumi.Input<boolean>;
    /**
     * Property specifying whether protection against purge is enabled for this configuration store.
     */
    enablePurgeProtection?: pulumi.Input<boolean>;
    /**
     * The encryption settings of the configuration store.
     */
    encryption?: pulumi.Input<inputs.appconfiguration.EncryptionPropertiesArgs>;
    /**
     * The managed identity information, if configured.
     */
    identity?: pulumi.Input<inputs.appconfiguration.ResourceIdentityArgs>;
    /**
     * The geo-location where the resource lives
     */
    location?: pulumi.Input<string>;
    /**
     * Control permission for data plane traffic coming from public networks while private endpoint is enabled.
     */
    publicNetworkAccess?: pulumi.Input<string | enums.appconfiguration.PublicNetworkAccess>;
    /**
     * The name of the resource group to which the container registry belongs.
     */
    resourceGroupName: pulumi.Input<string>;
    /**
     * The sku of the configuration store.
     */
    sku: pulumi.Input<inputs.appconfiguration.SkuArgs>;
    /**
     * The amount of time in days that the configuration store will be retained when it is soft deleted.
     */
    softDeleteRetentionInDays?: pulumi.Input<number>;
    /**
     * Resource tags.
     */
    tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
}
