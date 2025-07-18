// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Full view of the custom domain suffix configuration for ASEv3.
 *
 * Uses Azure REST API version 2024-04-01. In version 2.x of the Azure Native provider, it used API version 2022-09-01.
 *
 * Other available API versions: 2022-03-01, 2022-09-01, 2023-01-01, 2023-12-01, 2024-11-01. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native web [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
 */
export class AppServiceEnvironmentAseCustomDnsSuffixConfiguration extends pulumi.CustomResource {
    /**
     * Get an existing AppServiceEnvironmentAseCustomDnsSuffixConfiguration resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): AppServiceEnvironmentAseCustomDnsSuffixConfiguration {
        return new AppServiceEnvironmentAseCustomDnsSuffixConfiguration(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azure-native:web:AppServiceEnvironmentAseCustomDnsSuffixConfiguration';

    /**
     * Returns true if the given object is an instance of AppServiceEnvironmentAseCustomDnsSuffixConfiguration.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is AppServiceEnvironmentAseCustomDnsSuffixConfiguration {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === AppServiceEnvironmentAseCustomDnsSuffixConfiguration.__pulumiType;
    }

    /**
     * The Azure API version of the resource.
     */
    public /*out*/ readonly azureApiVersion!: pulumi.Output<string>;
    /**
     * The URL referencing the Azure Key Vault certificate secret that should be used as the default SSL/TLS certificate for sites with the custom domain suffix.
     */
    public readonly certificateUrl!: pulumi.Output<string | undefined>;
    /**
     * The default custom domain suffix to use for all sites deployed on the ASE.
     */
    public readonly dnsSuffix!: pulumi.Output<string | undefined>;
    /**
     * The user-assigned identity to use for resolving the key vault certificate reference. If not specified, the system-assigned ASE identity will be used if available.
     */
    public readonly keyVaultReferenceIdentity!: pulumi.Output<string | undefined>;
    /**
     * Kind of resource.
     */
    public readonly kind!: pulumi.Output<string | undefined>;
    /**
     * Resource Name.
     */
    public readonly name!: pulumi.Output<string>;
    public /*out*/ readonly provisioningDetails!: pulumi.Output<string>;
    public /*out*/ readonly provisioningState!: pulumi.Output<string>;
    /**
     * Resource type.
     */
    public /*out*/ readonly type!: pulumi.Output<string>;

    /**
     * Create a AppServiceEnvironmentAseCustomDnsSuffixConfiguration resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: AppServiceEnvironmentAseCustomDnsSuffixConfigurationArgs, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.name === undefined) && !opts.urn) {
                throw new Error("Missing required property 'name'");
            }
            if ((!args || args.resourceGroupName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'resourceGroupName'");
            }
            resourceInputs["certificateUrl"] = args ? args.certificateUrl : undefined;
            resourceInputs["dnsSuffix"] = args ? args.dnsSuffix : undefined;
            resourceInputs["keyVaultReferenceIdentity"] = args ? args.keyVaultReferenceIdentity : undefined;
            resourceInputs["kind"] = args ? args.kind : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["resourceGroupName"] = args ? args.resourceGroupName : undefined;
            resourceInputs["azureApiVersion"] = undefined /*out*/;
            resourceInputs["provisioningDetails"] = undefined /*out*/;
            resourceInputs["provisioningState"] = undefined /*out*/;
            resourceInputs["type"] = undefined /*out*/;
        } else {
            resourceInputs["azureApiVersion"] = undefined /*out*/;
            resourceInputs["certificateUrl"] = undefined /*out*/;
            resourceInputs["dnsSuffix"] = undefined /*out*/;
            resourceInputs["keyVaultReferenceIdentity"] = undefined /*out*/;
            resourceInputs["kind"] = undefined /*out*/;
            resourceInputs["name"] = undefined /*out*/;
            resourceInputs["provisioningDetails"] = undefined /*out*/;
            resourceInputs["provisioningState"] = undefined /*out*/;
            resourceInputs["type"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        const aliasOpts = { aliases: [{ type: "azure-native:web/v20220301:AppServiceEnvironmentAseCustomDnsSuffixConfiguration" }, { type: "azure-native:web/v20220901:AppServiceEnvironmentAseCustomDnsSuffixConfiguration" }, { type: "azure-native:web/v20230101:AppServiceEnvironmentAseCustomDnsSuffixConfiguration" }, { type: "azure-native:web/v20231201:AppServiceEnvironmentAseCustomDnsSuffixConfiguration" }, { type: "azure-native:web/v20240401:AppServiceEnvironmentAseCustomDnsSuffixConfiguration" }, { type: "azure-native:web/v20241101:AppServiceEnvironmentAseCustomDnsSuffixConfiguration" }] };
        opts = pulumi.mergeOptions(opts, aliasOpts);
        super(AppServiceEnvironmentAseCustomDnsSuffixConfiguration.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * The set of arguments for constructing a AppServiceEnvironmentAseCustomDnsSuffixConfiguration resource.
 */
export interface AppServiceEnvironmentAseCustomDnsSuffixConfigurationArgs {
    /**
     * The URL referencing the Azure Key Vault certificate secret that should be used as the default SSL/TLS certificate for sites with the custom domain suffix.
     */
    certificateUrl?: pulumi.Input<string>;
    /**
     * The default custom domain suffix to use for all sites deployed on the ASE.
     */
    dnsSuffix?: pulumi.Input<string>;
    /**
     * The user-assigned identity to use for resolving the key vault certificate reference. If not specified, the system-assigned ASE identity will be used if available.
     */
    keyVaultReferenceIdentity?: pulumi.Input<string>;
    /**
     * Kind of resource.
     */
    kind?: pulumi.Input<string>;
    /**
     * Name of the App Service Environment.
     */
    name: pulumi.Input<string>;
    /**
     * Name of the resource group to which the resource belongs.
     */
    resourceGroupName: pulumi.Input<string>;
}
