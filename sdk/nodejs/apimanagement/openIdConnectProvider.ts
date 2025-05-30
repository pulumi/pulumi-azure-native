// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * OpenId Connect Provider details.
 *
 * Uses Azure REST API version 2022-09-01-preview. In version 2.x of the Azure Native provider, it used API version 2022-08-01.
 *
 * Other available API versions: 2021-04-01-preview, 2021-08-01, 2021-12-01-preview, 2022-04-01-preview, 2022-08-01, 2023-03-01-preview, 2023-05-01-preview, 2023-09-01-preview, 2024-05-01, 2024-06-01-preview. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native apimanagement [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
 */
export class OpenIdConnectProvider extends pulumi.CustomResource {
    /**
     * Get an existing OpenIdConnectProvider resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): OpenIdConnectProvider {
        return new OpenIdConnectProvider(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azure-native:apimanagement:OpenIdConnectProvider';

    /**
     * Returns true if the given object is an instance of OpenIdConnectProvider.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is OpenIdConnectProvider {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === OpenIdConnectProvider.__pulumiType;
    }

    /**
     * The Azure API version of the resource.
     */
    public /*out*/ readonly azureApiVersion!: pulumi.Output<string>;
    /**
     * Client ID of developer console which is the client application.
     */
    public readonly clientId!: pulumi.Output<string>;
    /**
     * Client Secret of developer console which is the client application.
     */
    public readonly clientSecret!: pulumi.Output<string | undefined>;
    /**
     * User-friendly description of OpenID Connect Provider.
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * User-friendly OpenID Connect Provider name.
     */
    public readonly displayName!: pulumi.Output<string>;
    /**
     * Metadata endpoint URI.
     */
    public readonly metadataEndpoint!: pulumi.Output<string>;
    /**
     * The name of the resource
     */
    public /*out*/ readonly name!: pulumi.Output<string>;
    /**
     * The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
     */
    public /*out*/ readonly type!: pulumi.Output<string>;
    /**
     * If true, the Open ID Connect provider will be used in the API documentation in the developer portal. False by default if no value is provided.
     */
    public readonly useInApiDocumentation!: pulumi.Output<boolean | undefined>;
    /**
     * If true, the Open ID Connect provider may be used in the developer portal test console. True by default if no value is provided.
     */
    public readonly useInTestConsole!: pulumi.Output<boolean | undefined>;

    /**
     * Create a OpenIdConnectProvider resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: OpenIdConnectProviderArgs, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.clientId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'clientId'");
            }
            if ((!args || args.displayName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'displayName'");
            }
            if ((!args || args.metadataEndpoint === undefined) && !opts.urn) {
                throw new Error("Missing required property 'metadataEndpoint'");
            }
            if ((!args || args.resourceGroupName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'resourceGroupName'");
            }
            if ((!args || args.serviceName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'serviceName'");
            }
            resourceInputs["clientId"] = args ? args.clientId : undefined;
            resourceInputs["clientSecret"] = args ? args.clientSecret : undefined;
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["displayName"] = args ? args.displayName : undefined;
            resourceInputs["metadataEndpoint"] = args ? args.metadataEndpoint : undefined;
            resourceInputs["opid"] = args ? args.opid : undefined;
            resourceInputs["resourceGroupName"] = args ? args.resourceGroupName : undefined;
            resourceInputs["serviceName"] = args ? args.serviceName : undefined;
            resourceInputs["useInApiDocumentation"] = args ? args.useInApiDocumentation : undefined;
            resourceInputs["useInTestConsole"] = args ? args.useInTestConsole : undefined;
            resourceInputs["azureApiVersion"] = undefined /*out*/;
            resourceInputs["name"] = undefined /*out*/;
            resourceInputs["type"] = undefined /*out*/;
        } else {
            resourceInputs["azureApiVersion"] = undefined /*out*/;
            resourceInputs["clientId"] = undefined /*out*/;
            resourceInputs["clientSecret"] = undefined /*out*/;
            resourceInputs["description"] = undefined /*out*/;
            resourceInputs["displayName"] = undefined /*out*/;
            resourceInputs["metadataEndpoint"] = undefined /*out*/;
            resourceInputs["name"] = undefined /*out*/;
            resourceInputs["type"] = undefined /*out*/;
            resourceInputs["useInApiDocumentation"] = undefined /*out*/;
            resourceInputs["useInTestConsole"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        const aliasOpts = { aliases: [{ type: "azure-native:apimanagement/v20160707:OpenIdConnectProvider" }, { type: "azure-native:apimanagement/v20161010:OpenIdConnectProvider" }, { type: "azure-native:apimanagement/v20170301:OpenIdConnectProvider" }, { type: "azure-native:apimanagement/v20180101:OpenIdConnectProvider" }, { type: "azure-native:apimanagement/v20180601preview:OpenIdConnectProvider" }, { type: "azure-native:apimanagement/v20190101:OpenIdConnectProvider" }, { type: "azure-native:apimanagement/v20191201:OpenIdConnectProvider" }, { type: "azure-native:apimanagement/v20191201preview:OpenIdConnectProvider" }, { type: "azure-native:apimanagement/v20200601preview:OpenIdConnectProvider" }, { type: "azure-native:apimanagement/v20201201:OpenIdConnectProvider" }, { type: "azure-native:apimanagement/v20210101preview:OpenIdConnectProvider" }, { type: "azure-native:apimanagement/v20210401preview:OpenIdConnectProvider" }, { type: "azure-native:apimanagement/v20210801:OpenIdConnectProvider" }, { type: "azure-native:apimanagement/v20211201preview:OpenIdConnectProvider" }, { type: "azure-native:apimanagement/v20220401preview:OpenIdConnectProvider" }, { type: "azure-native:apimanagement/v20220801:OpenIdConnectProvider" }, { type: "azure-native:apimanagement/v20220901preview:OpenIdConnectProvider" }, { type: "azure-native:apimanagement/v20230301preview:OpenIdConnectProvider" }, { type: "azure-native:apimanagement/v20230501preview:OpenIdConnectProvider" }, { type: "azure-native:apimanagement/v20230901preview:OpenIdConnectProvider" }, { type: "azure-native:apimanagement/v20240501:OpenIdConnectProvider" }, { type: "azure-native:apimanagement/v20240601preview:OpenIdConnectProvider" }] };
        opts = pulumi.mergeOptions(opts, aliasOpts);
        super(OpenIdConnectProvider.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * The set of arguments for constructing a OpenIdConnectProvider resource.
 */
export interface OpenIdConnectProviderArgs {
    /**
     * Client ID of developer console which is the client application.
     */
    clientId: pulumi.Input<string>;
    /**
     * Client Secret of developer console which is the client application.
     */
    clientSecret?: pulumi.Input<string>;
    /**
     * User-friendly description of OpenID Connect Provider.
     */
    description?: pulumi.Input<string>;
    /**
     * User-friendly OpenID Connect Provider name.
     */
    displayName: pulumi.Input<string>;
    /**
     * Metadata endpoint URI.
     */
    metadataEndpoint: pulumi.Input<string>;
    /**
     * Identifier of the OpenID Connect Provider.
     */
    opid?: pulumi.Input<string>;
    /**
     * The name of the resource group. The name is case insensitive.
     */
    resourceGroupName: pulumi.Input<string>;
    /**
     * The name of the API Management service.
     */
    serviceName: pulumi.Input<string>;
    /**
     * If true, the Open ID Connect provider will be used in the API documentation in the developer portal. False by default if no value is provided.
     */
    useInApiDocumentation?: pulumi.Input<boolean>;
    /**
     * If true, the Open ID Connect provider may be used in the developer portal test console. True by default if no value is provided.
     */
    useInTestConsole?: pulumi.Input<boolean>;
}
