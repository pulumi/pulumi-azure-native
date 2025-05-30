// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * Contact details and configurations for notifications coming from Microsoft Defender for Cloud.
 *
 * Uses Azure REST API version 2023-12-01-preview. In version 2.x of the Azure Native provider, it used API version 2020-01-01-preview.
 *
 * Other available API versions: 2017-08-01-preview, 2020-01-01-preview. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native security [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
 */
export class SecurityContact extends pulumi.CustomResource {
    /**
     * Get an existing SecurityContact resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): SecurityContact {
        return new SecurityContact(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azure-native:security:SecurityContact';

    /**
     * Returns true if the given object is an instance of SecurityContact.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is SecurityContact {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === SecurityContact.__pulumiType;
    }

    /**
     * The Azure API version of the resource.
     */
    public /*out*/ readonly azureApiVersion!: pulumi.Output<string>;
    /**
     * List of email addresses which will get notifications from Microsoft Defender for Cloud by the configurations defined in this security contact.
     */
    public readonly emails!: pulumi.Output<string | undefined>;
    /**
     * Indicates whether the security contact is enabled.
     */
    public readonly isEnabled!: pulumi.Output<boolean | undefined>;
    /**
     * Resource name
     */
    public /*out*/ readonly name!: pulumi.Output<string>;
    /**
     * Defines whether to send email notifications from Microsoft Defender for Cloud to persons with specific RBAC roles on the subscription.
     */
    public readonly notificationsByRole!: pulumi.Output<outputs.security.SecurityContactPropertiesResponseNotificationsByRole | undefined>;
    /**
     * A collection of sources types which evaluate the email notification.
     */
    public readonly notificationsSources!: pulumi.Output<(outputs.security.NotificationsSourceAlertResponse | outputs.security.NotificationsSourceAttackPathResponse)[] | undefined>;
    /**
     * The security contact's phone number
     */
    public readonly phone!: pulumi.Output<string | undefined>;
    /**
     * Resource type
     */
    public /*out*/ readonly type!: pulumi.Output<string>;

    /**
     * Create a SecurityContact resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: SecurityContactArgs, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            resourceInputs["emails"] = args ? args.emails : undefined;
            resourceInputs["isEnabled"] = args ? args.isEnabled : undefined;
            resourceInputs["notificationsByRole"] = args ? args.notificationsByRole : undefined;
            resourceInputs["notificationsSources"] = args ? args.notificationsSources : undefined;
            resourceInputs["phone"] = args ? args.phone : undefined;
            resourceInputs["securityContactName"] = args ? args.securityContactName : undefined;
            resourceInputs["azureApiVersion"] = undefined /*out*/;
            resourceInputs["name"] = undefined /*out*/;
            resourceInputs["type"] = undefined /*out*/;
        } else {
            resourceInputs["azureApiVersion"] = undefined /*out*/;
            resourceInputs["emails"] = undefined /*out*/;
            resourceInputs["isEnabled"] = undefined /*out*/;
            resourceInputs["name"] = undefined /*out*/;
            resourceInputs["notificationsByRole"] = undefined /*out*/;
            resourceInputs["notificationsSources"] = undefined /*out*/;
            resourceInputs["phone"] = undefined /*out*/;
            resourceInputs["type"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        const aliasOpts = { aliases: [{ type: "azure-native:security/v20170801preview:SecurityContact" }, { type: "azure-native:security/v20200101preview:SecurityContact" }, { type: "azure-native:security/v20231201preview:SecurityContact" }] };
        opts = pulumi.mergeOptions(opts, aliasOpts);
        super(SecurityContact.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * The set of arguments for constructing a SecurityContact resource.
 */
export interface SecurityContactArgs {
    /**
     * List of email addresses which will get notifications from Microsoft Defender for Cloud by the configurations defined in this security contact.
     */
    emails?: pulumi.Input<string>;
    /**
     * Indicates whether the security contact is enabled.
     */
    isEnabled?: pulumi.Input<boolean>;
    /**
     * Defines whether to send email notifications from Microsoft Defender for Cloud to persons with specific RBAC roles on the subscription.
     */
    notificationsByRole?: pulumi.Input<inputs.security.SecurityContactPropertiesNotificationsByRoleArgs>;
    /**
     * A collection of sources types which evaluate the email notification.
     */
    notificationsSources?: pulumi.Input<pulumi.Input<inputs.security.NotificationsSourceAlertArgs | inputs.security.NotificationsSourceAttackPathArgs>[]>;
    /**
     * The security contact's phone number
     */
    phone?: pulumi.Input<string>;
    /**
     * Name of the security contact object
     */
    securityContactName?: pulumi.Input<string>;
}
