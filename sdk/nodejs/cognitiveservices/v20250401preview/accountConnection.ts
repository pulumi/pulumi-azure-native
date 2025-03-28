// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as enums from "../../types/enums";
import * as utilities from "../../utilities";

/**
 * Connection base resource schema.
 */
export class AccountConnection extends pulumi.CustomResource {
    /**
     * Get an existing AccountConnection resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): AccountConnection {
        return new AccountConnection(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azure-native:cognitiveservices/v20250401preview:AccountConnection';

    /**
     * Returns true if the given object is an instance of AccountConnection.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is AccountConnection {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === AccountConnection.__pulumiType;
    }

    /**
     * The name of the resource
     */
    public /*out*/ readonly name!: pulumi.Output<string>;
    /**
     * Connection property base schema.
     */
    public readonly properties!: pulumi.Output<outputs.cognitiveservices.v20250401preview.AADAuthTypeConnectionPropertiesResponse | outputs.cognitiveservices.v20250401preview.AccessKeyAuthTypeConnectionPropertiesResponse | outputs.cognitiveservices.v20250401preview.AccountKeyAuthTypeConnectionPropertiesResponse | outputs.cognitiveservices.v20250401preview.ApiKeyAuthConnectionPropertiesResponse | outputs.cognitiveservices.v20250401preview.CustomKeysConnectionPropertiesResponse | outputs.cognitiveservices.v20250401preview.ManagedIdentityAuthTypeConnectionPropertiesResponse | outputs.cognitiveservices.v20250401preview.NoneAuthTypeConnectionPropertiesResponse | outputs.cognitiveservices.v20250401preview.OAuth2AuthTypeConnectionPropertiesResponse | outputs.cognitiveservices.v20250401preview.PATAuthTypeConnectionPropertiesResponse | outputs.cognitiveservices.v20250401preview.SASAuthTypeConnectionPropertiesResponse | outputs.cognitiveservices.v20250401preview.ServicePrincipalAuthTypeConnectionPropertiesResponse | outputs.cognitiveservices.v20250401preview.UsernamePasswordAuthTypeConnectionPropertiesResponse>;
    /**
     * The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
     */
    public /*out*/ readonly type!: pulumi.Output<string>;

    /**
     * Create a AccountConnection resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: AccountConnectionArgs, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.accountName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'accountName'");
            }
            if ((!args || args.properties === undefined) && !opts.urn) {
                throw new Error("Missing required property 'properties'");
            }
            if ((!args || args.resourceGroupName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'resourceGroupName'");
            }
            resourceInputs["accountName"] = args ? args.accountName : undefined;
            resourceInputs["connectionName"] = args ? args.connectionName : undefined;
            resourceInputs["properties"] = args ? args.properties : undefined;
            resourceInputs["resourceGroupName"] = args ? args.resourceGroupName : undefined;
            resourceInputs["name"] = undefined /*out*/;
            resourceInputs["type"] = undefined /*out*/;
        } else {
            resourceInputs["name"] = undefined /*out*/;
            resourceInputs["properties"] = undefined /*out*/;
            resourceInputs["type"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(AccountConnection.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * The set of arguments for constructing a AccountConnection resource.
 */
export interface AccountConnectionArgs {
    /**
     * The name of Cognitive Services account.
     */
    accountName: pulumi.Input<string>;
    /**
     * Friendly name of the connection
     */
    connectionName?: pulumi.Input<string>;
    /**
     * Connection property base schema.
     */
    properties: pulumi.Input<inputs.cognitiveservices.v20250401preview.AADAuthTypeConnectionPropertiesArgs | inputs.cognitiveservices.v20250401preview.AccessKeyAuthTypeConnectionPropertiesArgs | inputs.cognitiveservices.v20250401preview.AccountKeyAuthTypeConnectionPropertiesArgs | inputs.cognitiveservices.v20250401preview.ApiKeyAuthConnectionPropertiesArgs | inputs.cognitiveservices.v20250401preview.CustomKeysConnectionPropertiesArgs | inputs.cognitiveservices.v20250401preview.ManagedIdentityAuthTypeConnectionPropertiesArgs | inputs.cognitiveservices.v20250401preview.NoneAuthTypeConnectionPropertiesArgs | inputs.cognitiveservices.v20250401preview.OAuth2AuthTypeConnectionPropertiesArgs | inputs.cognitiveservices.v20250401preview.PATAuthTypeConnectionPropertiesArgs | inputs.cognitiveservices.v20250401preview.SASAuthTypeConnectionPropertiesArgs | inputs.cognitiveservices.v20250401preview.ServicePrincipalAuthTypeConnectionPropertiesArgs | inputs.cognitiveservices.v20250401preview.UsernamePasswordAuthTypeConnectionPropertiesArgs>;
    /**
     * The name of the resource group. The name is case insensitive.
     */
    resourceGroupName: pulumi.Input<string>;
}
