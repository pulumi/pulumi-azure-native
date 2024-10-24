// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as enums from "../../types/enums";
import * as utilities from "../../utilities";

/**
 * The data product resource.
 */
export class DataProduct extends pulumi.CustomResource {
    /**
     * Get an existing DataProduct resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): DataProduct {
        return new DataProduct(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azure-native:networkanalytics/v20231115:DataProduct';

    /**
     * Returns true if the given object is an instance of DataProduct.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is DataProduct {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === DataProduct.__pulumiType;
    }

    /**
     * List of available minor versions of the data product resource.
     */
    public /*out*/ readonly availableMinorVersions!: pulumi.Output<string[]>;
    /**
     * Resource links which exposed to the customer to query the data.
     */
    public /*out*/ readonly consumptionEndpoints!: pulumi.Output<outputs.networkanalytics.v20231115.ConsumptionEndpointsPropertiesResponse>;
    /**
     * Current configured minor version of the data product resource.
     */
    public readonly currentMinorVersion!: pulumi.Output<string | undefined>;
    /**
     * Customer managed encryption key details for data product.
     */
    public readonly customerEncryptionKey!: pulumi.Output<outputs.networkanalytics.v20231115.EncryptionKeyDetailsResponse | undefined>;
    /**
     * Flag to enable customer managed key encryption for data product.
     */
    public readonly customerManagedKeyEncryptionEnabled!: pulumi.Output<string | undefined>;
    /**
     * Documentation link for the data product based on definition file.
     */
    public /*out*/ readonly documentation!: pulumi.Output<string>;
    /**
     * The managed service identities assigned to this resource.
     */
    public readonly identity!: pulumi.Output<outputs.networkanalytics.v20231115.ManagedServiceIdentityResponse | undefined>;
    /**
     * Key vault url.
     */
    public /*out*/ readonly keyVaultUrl!: pulumi.Output<string>;
    /**
     * The geo-location where the resource lives
     */
    public readonly location!: pulumi.Output<string>;
    /**
     * Major version of data product.
     */
    public readonly majorVersion!: pulumi.Output<string>;
    /**
     * The name of the resource
     */
    public /*out*/ readonly name!: pulumi.Output<string>;
    /**
     * Network rule set for data product.
     */
    public readonly networkacls!: pulumi.Output<outputs.networkanalytics.v20231115.DataProductNetworkAclsResponse | undefined>;
    /**
     * List of name or email associated with data product resource deployment.
     */
    public readonly owners!: pulumi.Output<string[] | undefined>;
    /**
     * Flag to enable or disable private link for data product resource.
     */
    public readonly privateLinksEnabled!: pulumi.Output<string | undefined>;
    /**
     * Product name of data product.
     */
    public readonly product!: pulumi.Output<string>;
    /**
     * Latest provisioning state  of data product.
     */
    public /*out*/ readonly provisioningState!: pulumi.Output<string>;
    /**
     * Flag to enable or disable public access of data product resource.
     */
    public readonly publicNetworkAccess!: pulumi.Output<string | undefined>;
    /**
     * Data product publisher name.
     */
    public readonly publisher!: pulumi.Output<string>;
    /**
     * Purview account url for data product to connect to.
     */
    public readonly purviewAccount!: pulumi.Output<string | undefined>;
    /**
     * Purview collection url for data product to connect to.
     */
    public readonly purviewCollection!: pulumi.Output<string | undefined>;
    /**
     * Flag to enable or disable redundancy for data product.
     */
    public readonly redundancy!: pulumi.Output<string | undefined>;
    /**
     * The resource GUID property of the data product resource.
     */
    public /*out*/ readonly resourceGuid!: pulumi.Output<string>;
    /**
     * Azure Resource Manager metadata containing createdBy and modifiedBy information.
     */
    public /*out*/ readonly systemData!: pulumi.Output<outputs.networkanalytics.v20231115.SystemDataResponse>;
    /**
     * Resource tags.
     */
    public readonly tags!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
     */
    public /*out*/ readonly type!: pulumi.Output<string>;

    /**
     * Create a DataProduct resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: DataProductArgs, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.majorVersion === undefined) && !opts.urn) {
                throw new Error("Missing required property 'majorVersion'");
            }
            if ((!args || args.product === undefined) && !opts.urn) {
                throw new Error("Missing required property 'product'");
            }
            if ((!args || args.publisher === undefined) && !opts.urn) {
                throw new Error("Missing required property 'publisher'");
            }
            if ((!args || args.resourceGroupName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'resourceGroupName'");
            }
            resourceInputs["currentMinorVersion"] = args ? args.currentMinorVersion : undefined;
            resourceInputs["customerEncryptionKey"] = args ? args.customerEncryptionKey : undefined;
            resourceInputs["customerManagedKeyEncryptionEnabled"] = args ? args.customerManagedKeyEncryptionEnabled : undefined;
            resourceInputs["dataProductName"] = args ? args.dataProductName : undefined;
            resourceInputs["identity"] = args ? args.identity : undefined;
            resourceInputs["location"] = args ? args.location : undefined;
            resourceInputs["majorVersion"] = args ? args.majorVersion : undefined;
            resourceInputs["managedResourceGroupConfiguration"] = args ? args.managedResourceGroupConfiguration : undefined;
            resourceInputs["networkacls"] = args ? args.networkacls : undefined;
            resourceInputs["owners"] = args ? args.owners : undefined;
            resourceInputs["privateLinksEnabled"] = args ? args.privateLinksEnabled : undefined;
            resourceInputs["product"] = args ? args.product : undefined;
            resourceInputs["publicNetworkAccess"] = args ? args.publicNetworkAccess : undefined;
            resourceInputs["publisher"] = args ? args.publisher : undefined;
            resourceInputs["purviewAccount"] = args ? args.purviewAccount : undefined;
            resourceInputs["purviewCollection"] = args ? args.purviewCollection : undefined;
            resourceInputs["redundancy"] = args ? args.redundancy : undefined;
            resourceInputs["resourceGroupName"] = args ? args.resourceGroupName : undefined;
            resourceInputs["tags"] = args ? args.tags : undefined;
            resourceInputs["availableMinorVersions"] = undefined /*out*/;
            resourceInputs["consumptionEndpoints"] = undefined /*out*/;
            resourceInputs["documentation"] = undefined /*out*/;
            resourceInputs["keyVaultUrl"] = undefined /*out*/;
            resourceInputs["name"] = undefined /*out*/;
            resourceInputs["provisioningState"] = undefined /*out*/;
            resourceInputs["resourceGuid"] = undefined /*out*/;
            resourceInputs["systemData"] = undefined /*out*/;
            resourceInputs["type"] = undefined /*out*/;
        } else {
            resourceInputs["availableMinorVersions"] = undefined /*out*/;
            resourceInputs["consumptionEndpoints"] = undefined /*out*/;
            resourceInputs["currentMinorVersion"] = undefined /*out*/;
            resourceInputs["customerEncryptionKey"] = undefined /*out*/;
            resourceInputs["customerManagedKeyEncryptionEnabled"] = undefined /*out*/;
            resourceInputs["documentation"] = undefined /*out*/;
            resourceInputs["identity"] = undefined /*out*/;
            resourceInputs["keyVaultUrl"] = undefined /*out*/;
            resourceInputs["location"] = undefined /*out*/;
            resourceInputs["majorVersion"] = undefined /*out*/;
            resourceInputs["name"] = undefined /*out*/;
            resourceInputs["networkacls"] = undefined /*out*/;
            resourceInputs["owners"] = undefined /*out*/;
            resourceInputs["privateLinksEnabled"] = undefined /*out*/;
            resourceInputs["product"] = undefined /*out*/;
            resourceInputs["provisioningState"] = undefined /*out*/;
            resourceInputs["publicNetworkAccess"] = undefined /*out*/;
            resourceInputs["publisher"] = undefined /*out*/;
            resourceInputs["purviewAccount"] = undefined /*out*/;
            resourceInputs["purviewCollection"] = undefined /*out*/;
            resourceInputs["redundancy"] = undefined /*out*/;
            resourceInputs["resourceGuid"] = undefined /*out*/;
            resourceInputs["systemData"] = undefined /*out*/;
            resourceInputs["tags"] = undefined /*out*/;
            resourceInputs["type"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        const aliasOpts = { aliases: [{ type: "azure-native:networkanalytics:DataProduct" }] };
        opts = pulumi.mergeOptions(opts, aliasOpts);
        super(DataProduct.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * The set of arguments for constructing a DataProduct resource.
 */
export interface DataProductArgs {
    /**
     * Current configured minor version of the data product resource.
     */
    currentMinorVersion?: pulumi.Input<string>;
    /**
     * Customer managed encryption key details for data product.
     */
    customerEncryptionKey?: pulumi.Input<inputs.networkanalytics.v20231115.EncryptionKeyDetailsArgs>;
    /**
     * Flag to enable customer managed key encryption for data product.
     */
    customerManagedKeyEncryptionEnabled?: pulumi.Input<string | enums.networkanalytics.v20231115.ControlState>;
    /**
     * The data product resource name
     */
    dataProductName?: pulumi.Input<string>;
    /**
     * The managed service identities assigned to this resource.
     */
    identity?: pulumi.Input<inputs.networkanalytics.v20231115.ManagedServiceIdentityArgs>;
    /**
     * The geo-location where the resource lives
     */
    location?: pulumi.Input<string>;
    /**
     * Major version of data product.
     */
    majorVersion: pulumi.Input<string>;
    /**
     * Managed resource group configuration.
     */
    managedResourceGroupConfiguration?: pulumi.Input<inputs.networkanalytics.v20231115.ManagedResourceGroupConfigurationArgs>;
    /**
     * Network rule set for data product.
     */
    networkacls?: pulumi.Input<inputs.networkanalytics.v20231115.DataProductNetworkAclsArgs>;
    /**
     * List of name or email associated with data product resource deployment.
     */
    owners?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Flag to enable or disable private link for data product resource.
     */
    privateLinksEnabled?: pulumi.Input<string | enums.networkanalytics.v20231115.ControlState>;
    /**
     * Product name of data product.
     */
    product: pulumi.Input<string>;
    /**
     * Flag to enable or disable public access of data product resource.
     */
    publicNetworkAccess?: pulumi.Input<string | enums.networkanalytics.v20231115.ControlState>;
    /**
     * Data product publisher name.
     */
    publisher: pulumi.Input<string>;
    /**
     * Purview account url for data product to connect to.
     */
    purviewAccount?: pulumi.Input<string>;
    /**
     * Purview collection url for data product to connect to.
     */
    purviewCollection?: pulumi.Input<string>;
    /**
     * Flag to enable or disable redundancy for data product.
     */
    redundancy?: pulumi.Input<string | enums.networkanalytics.v20231115.ControlState>;
    /**
     * The name of the resource group. The name is case insensitive.
     */
    resourceGroupName: pulumi.Input<string>;
    /**
     * Resource tags.
     */
    tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
}
