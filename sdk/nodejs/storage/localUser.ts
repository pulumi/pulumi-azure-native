// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * The local user associated with the storage accounts.
 *
 * Uses Azure REST API version 2024-01-01. In version 2.x of the Azure Native provider, it used API version 2022-09-01.
 *
 * Other available API versions: 2022-09-01, 2023-01-01, 2023-04-01, 2023-05-01, 2025-01-01. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native storage [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
 */
export class LocalUser extends pulumi.CustomResource {
    /**
     * Get an existing LocalUser resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): LocalUser {
        return new LocalUser(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azure-native:storage:LocalUser';

    /**
     * Returns true if the given object is an instance of LocalUser.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is LocalUser {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === LocalUser.__pulumiType;
    }

    /**
     * Indicates whether ACL authorization is allowed for this user. Set it to false to disallow using ACL authorization.
     */
    public readonly allowAclAuthorization!: pulumi.Output<boolean | undefined>;
    /**
     * The Azure API version of the resource.
     */
    public /*out*/ readonly azureApiVersion!: pulumi.Output<string>;
    /**
     * Supplementary group membership. Only applicable for local users enabled for NFSv3 access.
     */
    public readonly extendedGroups!: pulumi.Output<number[] | undefined>;
    /**
     * An identifier for associating a group of users.
     */
    public readonly groupId!: pulumi.Output<number | undefined>;
    /**
     * Indicates whether shared key exists. Set it to false to remove existing shared key.
     */
    public readonly hasSharedKey!: pulumi.Output<boolean | undefined>;
    /**
     * Indicates whether ssh key exists. Set it to false to remove existing SSH key.
     */
    public readonly hasSshKey!: pulumi.Output<boolean | undefined>;
    /**
     * Indicates whether ssh password exists. Set it to false to remove existing SSH password.
     */
    public readonly hasSshPassword!: pulumi.Output<boolean | undefined>;
    /**
     * Optional, local user home directory.
     */
    public readonly homeDirectory!: pulumi.Output<string | undefined>;
    /**
     * Indicates if the local user is enabled for access with NFSv3 protocol.
     */
    public readonly isNFSv3Enabled!: pulumi.Output<boolean | undefined>;
    /**
     * The name of the resource
     */
    public /*out*/ readonly name!: pulumi.Output<string>;
    /**
     * The permission scopes of the local user.
     */
    public readonly permissionScopes!: pulumi.Output<outputs.storage.PermissionScopeResponse[] | undefined>;
    /**
     * A unique Security Identifier that is generated by the server.
     */
    public /*out*/ readonly sid!: pulumi.Output<string>;
    /**
     * Optional, local user ssh authorized keys for SFTP.
     */
    public readonly sshAuthorizedKeys!: pulumi.Output<outputs.storage.SshPublicKeyResponse[] | undefined>;
    /**
     * Metadata pertaining to creation and last modification of the resource.
     */
    public /*out*/ readonly systemData!: pulumi.Output<outputs.storage.SystemDataResponse>;
    /**
     * The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
     */
    public /*out*/ readonly type!: pulumi.Output<string>;
    /**
     * A unique Identifier that is generated by the server.
     */
    public /*out*/ readonly userId!: pulumi.Output<number>;

    /**
     * Create a LocalUser resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: LocalUserArgs, opts?: pulumi.CustomResourceOptions) {
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
            resourceInputs["allowAclAuthorization"] = args ? args.allowAclAuthorization : undefined;
            resourceInputs["extendedGroups"] = args ? args.extendedGroups : undefined;
            resourceInputs["groupId"] = args ? args.groupId : undefined;
            resourceInputs["hasSharedKey"] = args ? args.hasSharedKey : undefined;
            resourceInputs["hasSshKey"] = args ? args.hasSshKey : undefined;
            resourceInputs["hasSshPassword"] = args ? args.hasSshPassword : undefined;
            resourceInputs["homeDirectory"] = args ? args.homeDirectory : undefined;
            resourceInputs["isNFSv3Enabled"] = args ? args.isNFSv3Enabled : undefined;
            resourceInputs["permissionScopes"] = args ? args.permissionScopes : undefined;
            resourceInputs["resourceGroupName"] = args ? args.resourceGroupName : undefined;
            resourceInputs["sshAuthorizedKeys"] = args ? args.sshAuthorizedKeys : undefined;
            resourceInputs["username"] = args ? args.username : undefined;
            resourceInputs["azureApiVersion"] = undefined /*out*/;
            resourceInputs["name"] = undefined /*out*/;
            resourceInputs["sid"] = undefined /*out*/;
            resourceInputs["systemData"] = undefined /*out*/;
            resourceInputs["type"] = undefined /*out*/;
            resourceInputs["userId"] = undefined /*out*/;
        } else {
            resourceInputs["allowAclAuthorization"] = undefined /*out*/;
            resourceInputs["azureApiVersion"] = undefined /*out*/;
            resourceInputs["extendedGroups"] = undefined /*out*/;
            resourceInputs["groupId"] = undefined /*out*/;
            resourceInputs["hasSharedKey"] = undefined /*out*/;
            resourceInputs["hasSshKey"] = undefined /*out*/;
            resourceInputs["hasSshPassword"] = undefined /*out*/;
            resourceInputs["homeDirectory"] = undefined /*out*/;
            resourceInputs["isNFSv3Enabled"] = undefined /*out*/;
            resourceInputs["name"] = undefined /*out*/;
            resourceInputs["permissionScopes"] = undefined /*out*/;
            resourceInputs["sid"] = undefined /*out*/;
            resourceInputs["sshAuthorizedKeys"] = undefined /*out*/;
            resourceInputs["systemData"] = undefined /*out*/;
            resourceInputs["type"] = undefined /*out*/;
            resourceInputs["userId"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        const aliasOpts = { aliases: [{ type: "azure-native:storage/v20210801:LocalUser" }, { type: "azure-native:storage/v20210901:LocalUser" }, { type: "azure-native:storage/v20220501:LocalUser" }, { type: "azure-native:storage/v20220901:LocalUser" }, { type: "azure-native:storage/v20230101:LocalUser" }, { type: "azure-native:storage/v20230401:LocalUser" }, { type: "azure-native:storage/v20230501:LocalUser" }, { type: "azure-native:storage/v20240101:LocalUser" }, { type: "azure-native:storage/v20250101:LocalUser" }] };
        opts = pulumi.mergeOptions(opts, aliasOpts);
        super(LocalUser.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * The set of arguments for constructing a LocalUser resource.
 */
export interface LocalUserArgs {
    /**
     * The name of the storage account within the specified resource group. Storage account names must be between 3 and 24 characters in length and use numbers and lower-case letters only.
     */
    accountName: pulumi.Input<string>;
    /**
     * Indicates whether ACL authorization is allowed for this user. Set it to false to disallow using ACL authorization.
     */
    allowAclAuthorization?: pulumi.Input<boolean>;
    /**
     * Supplementary group membership. Only applicable for local users enabled for NFSv3 access.
     */
    extendedGroups?: pulumi.Input<pulumi.Input<number>[]>;
    /**
     * An identifier for associating a group of users.
     */
    groupId?: pulumi.Input<number>;
    /**
     * Indicates whether shared key exists. Set it to false to remove existing shared key.
     */
    hasSharedKey?: pulumi.Input<boolean>;
    /**
     * Indicates whether ssh key exists. Set it to false to remove existing SSH key.
     */
    hasSshKey?: pulumi.Input<boolean>;
    /**
     * Indicates whether ssh password exists. Set it to false to remove existing SSH password.
     */
    hasSshPassword?: pulumi.Input<boolean>;
    /**
     * Optional, local user home directory.
     */
    homeDirectory?: pulumi.Input<string>;
    /**
     * Indicates if the local user is enabled for access with NFSv3 protocol.
     */
    isNFSv3Enabled?: pulumi.Input<boolean>;
    /**
     * The permission scopes of the local user.
     */
    permissionScopes?: pulumi.Input<pulumi.Input<inputs.storage.PermissionScopeArgs>[]>;
    /**
     * The name of the resource group within the user's subscription. The name is case insensitive.
     */
    resourceGroupName: pulumi.Input<string>;
    /**
     * Optional, local user ssh authorized keys for SFTP.
     */
    sshAuthorizedKeys?: pulumi.Input<pulumi.Input<inputs.storage.SshPublicKeyArgs>[]>;
    /**
     * The name of local user. The username must contain lowercase letters and numbers only. It must be unique only within the storage account.
     */
    username?: pulumi.Input<string>;
}
