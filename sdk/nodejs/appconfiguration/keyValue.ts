// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * The key-value resource along with all resource properties.
 *
 * Uses Azure REST API version 2024-05-01. In version 2.x of the Azure Native provider, it used API version 2023-03-01.
 *
 * Other available API versions: 2023-03-01, 2023-08-01-preview, 2023-09-01-preview, 2024-06-01, 2024-06-15-preview, 2025-02-01-preview. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native appconfiguration [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
 */
export class KeyValue extends pulumi.CustomResource {
    /**
     * Get an existing KeyValue resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): KeyValue {
        return new KeyValue(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azure-native:appconfiguration:KeyValue';

    /**
     * Returns true if the given object is an instance of KeyValue.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is KeyValue {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === KeyValue.__pulumiType;
    }

    /**
     * The Azure API version of the resource.
     */
    public /*out*/ readonly azureApiVersion!: pulumi.Output<string>;
    /**
     * The content type of the key-value's value.
     * Providing a proper content-type can enable transformations of values when they are retrieved by applications.
     */
    public readonly contentType!: pulumi.Output<string | undefined>;
    /**
     * An ETag indicating the state of a key-value within a configuration store.
     */
    public /*out*/ readonly eTag!: pulumi.Output<string>;
    /**
     * The primary identifier of a key-value.
     * The key is used in unison with the label to uniquely identify a key-value.
     */
    public /*out*/ readonly key!: pulumi.Output<string>;
    /**
     * A value used to group key-values.
     * The label is used in unison with the key to uniquely identify a key-value.
     */
    public /*out*/ readonly label!: pulumi.Output<string>;
    /**
     * The last time a modifying operation was performed on the given key-value.
     */
    public /*out*/ readonly lastModified!: pulumi.Output<string>;
    /**
     * A value indicating whether the key-value is locked.
     * A locked key-value may not be modified until it is unlocked.
     */
    public /*out*/ readonly locked!: pulumi.Output<boolean>;
    /**
     * The name of the resource.
     */
    public /*out*/ readonly name!: pulumi.Output<string>;
    /**
     * A dictionary of tags that can help identify what a key-value may be applicable for.
     */
    public readonly tags!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * The type of the resource.
     */
    public /*out*/ readonly type!: pulumi.Output<string>;
    /**
     * The value of the key-value.
     */
    public readonly value!: pulumi.Output<string | undefined>;

    /**
     * Create a KeyValue resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: KeyValueArgs, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.configStoreName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'configStoreName'");
            }
            if ((!args || args.resourceGroupName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'resourceGroupName'");
            }
            resourceInputs["configStoreName"] = args ? args.configStoreName : undefined;
            resourceInputs["contentType"] = args ? args.contentType : undefined;
            resourceInputs["keyValueName"] = args ? args.keyValueName : undefined;
            resourceInputs["resourceGroupName"] = args ? args.resourceGroupName : undefined;
            resourceInputs["tags"] = args ? args.tags : undefined;
            resourceInputs["value"] = args ? args.value : undefined;
            resourceInputs["azureApiVersion"] = undefined /*out*/;
            resourceInputs["eTag"] = undefined /*out*/;
            resourceInputs["key"] = undefined /*out*/;
            resourceInputs["label"] = undefined /*out*/;
            resourceInputs["lastModified"] = undefined /*out*/;
            resourceInputs["locked"] = undefined /*out*/;
            resourceInputs["name"] = undefined /*out*/;
            resourceInputs["type"] = undefined /*out*/;
        } else {
            resourceInputs["azureApiVersion"] = undefined /*out*/;
            resourceInputs["contentType"] = undefined /*out*/;
            resourceInputs["eTag"] = undefined /*out*/;
            resourceInputs["key"] = undefined /*out*/;
            resourceInputs["label"] = undefined /*out*/;
            resourceInputs["lastModified"] = undefined /*out*/;
            resourceInputs["locked"] = undefined /*out*/;
            resourceInputs["name"] = undefined /*out*/;
            resourceInputs["tags"] = undefined /*out*/;
            resourceInputs["type"] = undefined /*out*/;
            resourceInputs["value"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        const aliasOpts = { aliases: [{ type: "azure-native:appconfiguration/v20200701preview:KeyValue" }, { type: "azure-native:appconfiguration/v20210301preview:KeyValue" }, { type: "azure-native:appconfiguration/v20211001preview:KeyValue" }, { type: "azure-native:appconfiguration/v20220301preview:KeyValue" }, { type: "azure-native:appconfiguration/v20220501:KeyValue" }, { type: "azure-native:appconfiguration/v20230301:KeyValue" }, { type: "azure-native:appconfiguration/v20230801preview:KeyValue" }, { type: "azure-native:appconfiguration/v20230901preview:KeyValue" }, { type: "azure-native:appconfiguration/v20240501:KeyValue" }, { type: "azure-native:appconfiguration/v20240601:KeyValue" }, { type: "azure-native:appconfiguration/v20240615preview:KeyValue" }, { type: "azure-native:appconfiguration/v20250201preview:KeyValue" }] };
        opts = pulumi.mergeOptions(opts, aliasOpts);
        super(KeyValue.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * The set of arguments for constructing a KeyValue resource.
 */
export interface KeyValueArgs {
    /**
     * The name of the configuration store.
     */
    configStoreName: pulumi.Input<string>;
    /**
     * The content type of the key-value's value.
     * Providing a proper content-type can enable transformations of values when they are retrieved by applications.
     */
    contentType?: pulumi.Input<string>;
    /**
     * Identifier of key and label combination. Key and label are joined by $ character. Label is optional.
     */
    keyValueName?: pulumi.Input<string>;
    /**
     * The name of the resource group to which the container registry belongs.
     */
    resourceGroupName: pulumi.Input<string>;
    /**
     * A dictionary of tags that can help identify what a key-value may be applicable for.
     */
    tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * The value of the key-value.
     */
    value?: pulumi.Input<string>;
}
