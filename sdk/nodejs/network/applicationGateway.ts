// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * Application gateway resource.
 *
 * Uses Azure REST API version 2024-05-01. In version 2.x of the Azure Native provider, it used API version 2023-02-01.
 *
 * Other available API versions: 2018-06-01, 2018-07-01, 2018-08-01, 2018-10-01, 2018-11-01, 2018-12-01, 2019-02-01, 2019-04-01, 2019-06-01, 2019-07-01, 2019-08-01, 2019-09-01, 2019-11-01, 2019-12-01, 2020-03-01, 2020-04-01, 2020-05-01, 2020-06-01, 2020-07-01, 2020-08-01, 2020-11-01, 2021-02-01, 2021-03-01, 2021-05-01, 2021-08-01, 2022-01-01, 2022-05-01, 2022-07-01, 2022-09-01, 2022-11-01, 2023-02-01, 2023-04-01, 2023-05-01, 2023-06-01, 2023-09-01, 2023-11-01, 2024-01-01, 2024-03-01, 2024-07-01. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native network [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
 */
export class ApplicationGateway extends pulumi.CustomResource {
    /**
     * Get an existing ApplicationGateway resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): ApplicationGateway {
        return new ApplicationGateway(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azure-native:network:ApplicationGateway';

    /**
     * Returns true if the given object is an instance of ApplicationGateway.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is ApplicationGateway {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === ApplicationGateway.__pulumiType;
    }

    /**
     * Authentication certificates of the application gateway resource. For default limits, see [Application Gateway limits](https://docs.microsoft.com/azure/azure-subscription-service-limits#application-gateway-limits).
     */
    public readonly authenticationCertificates!: pulumi.Output<outputs.network.ApplicationGatewayAuthenticationCertificateResponse[] | undefined>;
    /**
     * Autoscale Configuration.
     */
    public readonly autoscaleConfiguration!: pulumi.Output<outputs.network.ApplicationGatewayAutoscaleConfigurationResponse | undefined>;
    /**
     * The Azure API version of the resource.
     */
    public /*out*/ readonly azureApiVersion!: pulumi.Output<string>;
    /**
     * Backend address pool of the application gateway resource. For default limits, see [Application Gateway limits](https://docs.microsoft.com/azure/azure-subscription-service-limits#application-gateway-limits).
     */
    public readonly backendAddressPools!: pulumi.Output<outputs.network.ApplicationGatewayBackendAddressPoolResponse[] | undefined>;
    /**
     * Backend http settings of the application gateway resource. For default limits, see [Application Gateway limits](https://docs.microsoft.com/azure/azure-subscription-service-limits#application-gateway-limits).
     */
    public readonly backendHttpSettingsCollection!: pulumi.Output<outputs.network.ApplicationGatewayBackendHttpSettingsResponse[] | undefined>;
    /**
     * Backend settings of the application gateway resource. For default limits, see [Application Gateway limits](https://docs.microsoft.com/azure/azure-subscription-service-limits#application-gateway-limits).
     */
    public readonly backendSettingsCollection!: pulumi.Output<outputs.network.ApplicationGatewayBackendSettingsResponse[] | undefined>;
    /**
     * Custom error configurations of the application gateway resource.
     */
    public readonly customErrorConfigurations!: pulumi.Output<outputs.network.ApplicationGatewayCustomErrorResponse[] | undefined>;
    /**
     * The default predefined SSL Policy applied on the application gateway resource.
     */
    public /*out*/ readonly defaultPredefinedSslPolicy!: pulumi.Output<string>;
    /**
     * Whether FIPS is enabled on the application gateway resource.
     */
    public readonly enableFips!: pulumi.Output<boolean | undefined>;
    /**
     * Whether HTTP2 is enabled on the application gateway resource.
     */
    public readonly enableHttp2!: pulumi.Output<boolean | undefined>;
    /**
     * A unique read-only string that changes whenever the resource is updated.
     */
    public /*out*/ readonly etag!: pulumi.Output<string>;
    /**
     * Reference to the FirewallPolicy resource.
     */
    public readonly firewallPolicy!: pulumi.Output<outputs.network.SubResourceResponse | undefined>;
    /**
     * If true, associates a firewall policy with an application gateway regardless whether the policy differs from the WAF Config.
     */
    public readonly forceFirewallPolicyAssociation!: pulumi.Output<boolean | undefined>;
    /**
     * Frontend IP addresses of the application gateway resource. For default limits, see [Application Gateway limits](https://docs.microsoft.com/azure/azure-subscription-service-limits#application-gateway-limits).
     */
    public readonly frontendIPConfigurations!: pulumi.Output<outputs.network.ApplicationGatewayFrontendIPConfigurationResponse[] | undefined>;
    /**
     * Frontend ports of the application gateway resource. For default limits, see [Application Gateway limits](https://docs.microsoft.com/azure/azure-subscription-service-limits#application-gateway-limits).
     */
    public readonly frontendPorts!: pulumi.Output<outputs.network.ApplicationGatewayFrontendPortResponse[] | undefined>;
    /**
     * Subnets of the application gateway resource. For default limits, see [Application Gateway limits](https://docs.microsoft.com/azure/azure-subscription-service-limits#application-gateway-limits).
     */
    public readonly gatewayIPConfigurations!: pulumi.Output<outputs.network.ApplicationGatewayIPConfigurationResponse[] | undefined>;
    /**
     * Global Configuration.
     */
    public readonly globalConfiguration!: pulumi.Output<outputs.network.ApplicationGatewayGlobalConfigurationResponse | undefined>;
    /**
     * Http listeners of the application gateway resource. For default limits, see [Application Gateway limits](https://docs.microsoft.com/azure/azure-subscription-service-limits#application-gateway-limits).
     */
    public readonly httpListeners!: pulumi.Output<outputs.network.ApplicationGatewayHttpListenerResponse[] | undefined>;
    /**
     * The identity of the application gateway, if configured.
     */
    public readonly identity!: pulumi.Output<outputs.network.ManagedServiceIdentityResponse | undefined>;
    /**
     * Listeners of the application gateway resource. For default limits, see [Application Gateway limits](https://docs.microsoft.com/azure/azure-subscription-service-limits#application-gateway-limits).
     */
    public readonly listeners!: pulumi.Output<outputs.network.ApplicationGatewayListenerResponse[] | undefined>;
    /**
     * Load distribution policies of the application gateway resource.
     */
    public readonly loadDistributionPolicies!: pulumi.Output<outputs.network.ApplicationGatewayLoadDistributionPolicyResponse[] | undefined>;
    /**
     * Resource location.
     */
    public readonly location!: pulumi.Output<string | undefined>;
    /**
     * Resource name.
     */
    public /*out*/ readonly name!: pulumi.Output<string>;
    /**
     * Operational state of the application gateway resource.
     */
    public /*out*/ readonly operationalState!: pulumi.Output<string>;
    /**
     * Private Endpoint connections on application gateway.
     */
    public /*out*/ readonly privateEndpointConnections!: pulumi.Output<outputs.network.ApplicationGatewayPrivateEndpointConnectionResponse[]>;
    /**
     * PrivateLink configurations on application gateway.
     */
    public readonly privateLinkConfigurations!: pulumi.Output<outputs.network.ApplicationGatewayPrivateLinkConfigurationResponse[] | undefined>;
    /**
     * Probes of the application gateway resource.
     */
    public readonly probes!: pulumi.Output<outputs.network.ApplicationGatewayProbeResponse[] | undefined>;
    /**
     * The provisioning state of the application gateway resource.
     */
    public /*out*/ readonly provisioningState!: pulumi.Output<string>;
    /**
     * Redirect configurations of the application gateway resource. For default limits, see [Application Gateway limits](https://docs.microsoft.com/azure/azure-subscription-service-limits#application-gateway-limits).
     */
    public readonly redirectConfigurations!: pulumi.Output<outputs.network.ApplicationGatewayRedirectConfigurationResponse[] | undefined>;
    /**
     * Request routing rules of the application gateway resource.
     */
    public readonly requestRoutingRules!: pulumi.Output<outputs.network.ApplicationGatewayRequestRoutingRuleResponse[] | undefined>;
    /**
     * The resource GUID property of the application gateway resource.
     */
    public /*out*/ readonly resourceGuid!: pulumi.Output<string>;
    /**
     * Rewrite rules for the application gateway resource.
     */
    public readonly rewriteRuleSets!: pulumi.Output<outputs.network.ApplicationGatewayRewriteRuleSetResponse[] | undefined>;
    /**
     * Routing rules of the application gateway resource.
     */
    public readonly routingRules!: pulumi.Output<outputs.network.ApplicationGatewayRoutingRuleResponse[] | undefined>;
    /**
     * SKU of the application gateway resource.
     */
    public readonly sku!: pulumi.Output<outputs.network.ApplicationGatewaySkuResponse | undefined>;
    /**
     * SSL certificates of the application gateway resource. For default limits, see [Application Gateway limits](https://docs.microsoft.com/azure/azure-subscription-service-limits#application-gateway-limits).
     */
    public readonly sslCertificates!: pulumi.Output<outputs.network.ApplicationGatewaySslCertificateResponse[] | undefined>;
    /**
     * SSL policy of the application gateway resource.
     */
    public readonly sslPolicy!: pulumi.Output<outputs.network.ApplicationGatewaySslPolicyResponse | undefined>;
    /**
     * SSL profiles of the application gateway resource. For default limits, see [Application Gateway limits](https://docs.microsoft.com/azure/azure-subscription-service-limits#application-gateway-limits).
     */
    public readonly sslProfiles!: pulumi.Output<outputs.network.ApplicationGatewaySslProfileResponse[] | undefined>;
    /**
     * Resource tags.
     */
    public readonly tags!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * Trusted client certificates of the application gateway resource. For default limits, see [Application Gateway limits](https://docs.microsoft.com/azure/azure-subscription-service-limits#application-gateway-limits).
     */
    public readonly trustedClientCertificates!: pulumi.Output<outputs.network.ApplicationGatewayTrustedClientCertificateResponse[] | undefined>;
    /**
     * Trusted Root certificates of the application gateway resource. For default limits, see [Application Gateway limits](https://docs.microsoft.com/azure/azure-subscription-service-limits#application-gateway-limits).
     */
    public readonly trustedRootCertificates!: pulumi.Output<outputs.network.ApplicationGatewayTrustedRootCertificateResponse[] | undefined>;
    /**
     * Resource type.
     */
    public /*out*/ readonly type!: pulumi.Output<string>;
    /**
     * URL path map of the application gateway resource. For default limits, see [Application Gateway limits](https://docs.microsoft.com/azure/azure-subscription-service-limits#application-gateway-limits).
     */
    public readonly urlPathMaps!: pulumi.Output<outputs.network.ApplicationGatewayUrlPathMapResponse[] | undefined>;
    /**
     * Web application firewall configuration.
     */
    public readonly webApplicationFirewallConfiguration!: pulumi.Output<outputs.network.ApplicationGatewayWebApplicationFirewallConfigurationResponse | undefined>;
    /**
     * A list of availability zones denoting where the resource needs to come from.
     */
    public readonly zones!: pulumi.Output<string[] | undefined>;

    /**
     * Create a ApplicationGateway resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ApplicationGatewayArgs, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.resourceGroupName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'resourceGroupName'");
            }
            resourceInputs["applicationGatewayName"] = args ? args.applicationGatewayName : undefined;
            resourceInputs["authenticationCertificates"] = args ? args.authenticationCertificates : undefined;
            resourceInputs["autoscaleConfiguration"] = args ? args.autoscaleConfiguration : undefined;
            resourceInputs["backendAddressPools"] = args ? args.backendAddressPools : undefined;
            resourceInputs["backendHttpSettingsCollection"] = args ? args.backendHttpSettingsCollection : undefined;
            resourceInputs["backendSettingsCollection"] = args ? args.backendSettingsCollection : undefined;
            resourceInputs["customErrorConfigurations"] = args ? args.customErrorConfigurations : undefined;
            resourceInputs["enableFips"] = args ? args.enableFips : undefined;
            resourceInputs["enableHttp2"] = args ? args.enableHttp2 : undefined;
            resourceInputs["firewallPolicy"] = args ? args.firewallPolicy : undefined;
            resourceInputs["forceFirewallPolicyAssociation"] = args ? args.forceFirewallPolicyAssociation : undefined;
            resourceInputs["frontendIPConfigurations"] = args ? args.frontendIPConfigurations : undefined;
            resourceInputs["frontendPorts"] = args ? args.frontendPorts : undefined;
            resourceInputs["gatewayIPConfigurations"] = args ? args.gatewayIPConfigurations : undefined;
            resourceInputs["globalConfiguration"] = args ? args.globalConfiguration : undefined;
            resourceInputs["httpListeners"] = args ? args.httpListeners : undefined;
            resourceInputs["id"] = args ? args.id : undefined;
            resourceInputs["identity"] = args ? args.identity : undefined;
            resourceInputs["listeners"] = args ? args.listeners : undefined;
            resourceInputs["loadDistributionPolicies"] = args ? args.loadDistributionPolicies : undefined;
            resourceInputs["location"] = args ? args.location : undefined;
            resourceInputs["privateLinkConfigurations"] = args ? args.privateLinkConfigurations : undefined;
            resourceInputs["probes"] = args ? args.probes : undefined;
            resourceInputs["redirectConfigurations"] = args ? args.redirectConfigurations : undefined;
            resourceInputs["requestRoutingRules"] = args ? args.requestRoutingRules : undefined;
            resourceInputs["resourceGroupName"] = args ? args.resourceGroupName : undefined;
            resourceInputs["rewriteRuleSets"] = args ? args.rewriteRuleSets : undefined;
            resourceInputs["routingRules"] = args ? args.routingRules : undefined;
            resourceInputs["sku"] = args ? args.sku : undefined;
            resourceInputs["sslCertificates"] = args ? args.sslCertificates : undefined;
            resourceInputs["sslPolicy"] = args ? args.sslPolicy : undefined;
            resourceInputs["sslProfiles"] = args ? args.sslProfiles : undefined;
            resourceInputs["tags"] = args ? args.tags : undefined;
            resourceInputs["trustedClientCertificates"] = args ? args.trustedClientCertificates : undefined;
            resourceInputs["trustedRootCertificates"] = args ? args.trustedRootCertificates : undefined;
            resourceInputs["urlPathMaps"] = args ? args.urlPathMaps : undefined;
            resourceInputs["webApplicationFirewallConfiguration"] = args ? args.webApplicationFirewallConfiguration : undefined;
            resourceInputs["zones"] = args ? args.zones : undefined;
            resourceInputs["azureApiVersion"] = undefined /*out*/;
            resourceInputs["defaultPredefinedSslPolicy"] = undefined /*out*/;
            resourceInputs["etag"] = undefined /*out*/;
            resourceInputs["name"] = undefined /*out*/;
            resourceInputs["operationalState"] = undefined /*out*/;
            resourceInputs["privateEndpointConnections"] = undefined /*out*/;
            resourceInputs["provisioningState"] = undefined /*out*/;
            resourceInputs["resourceGuid"] = undefined /*out*/;
            resourceInputs["type"] = undefined /*out*/;
        } else {
            resourceInputs["authenticationCertificates"] = undefined /*out*/;
            resourceInputs["autoscaleConfiguration"] = undefined /*out*/;
            resourceInputs["azureApiVersion"] = undefined /*out*/;
            resourceInputs["backendAddressPools"] = undefined /*out*/;
            resourceInputs["backendHttpSettingsCollection"] = undefined /*out*/;
            resourceInputs["backendSettingsCollection"] = undefined /*out*/;
            resourceInputs["customErrorConfigurations"] = undefined /*out*/;
            resourceInputs["defaultPredefinedSslPolicy"] = undefined /*out*/;
            resourceInputs["enableFips"] = undefined /*out*/;
            resourceInputs["enableHttp2"] = undefined /*out*/;
            resourceInputs["etag"] = undefined /*out*/;
            resourceInputs["firewallPolicy"] = undefined /*out*/;
            resourceInputs["forceFirewallPolicyAssociation"] = undefined /*out*/;
            resourceInputs["frontendIPConfigurations"] = undefined /*out*/;
            resourceInputs["frontendPorts"] = undefined /*out*/;
            resourceInputs["gatewayIPConfigurations"] = undefined /*out*/;
            resourceInputs["globalConfiguration"] = undefined /*out*/;
            resourceInputs["httpListeners"] = undefined /*out*/;
            resourceInputs["identity"] = undefined /*out*/;
            resourceInputs["listeners"] = undefined /*out*/;
            resourceInputs["loadDistributionPolicies"] = undefined /*out*/;
            resourceInputs["location"] = undefined /*out*/;
            resourceInputs["name"] = undefined /*out*/;
            resourceInputs["operationalState"] = undefined /*out*/;
            resourceInputs["privateEndpointConnections"] = undefined /*out*/;
            resourceInputs["privateLinkConfigurations"] = undefined /*out*/;
            resourceInputs["probes"] = undefined /*out*/;
            resourceInputs["provisioningState"] = undefined /*out*/;
            resourceInputs["redirectConfigurations"] = undefined /*out*/;
            resourceInputs["requestRoutingRules"] = undefined /*out*/;
            resourceInputs["resourceGuid"] = undefined /*out*/;
            resourceInputs["rewriteRuleSets"] = undefined /*out*/;
            resourceInputs["routingRules"] = undefined /*out*/;
            resourceInputs["sku"] = undefined /*out*/;
            resourceInputs["sslCertificates"] = undefined /*out*/;
            resourceInputs["sslPolicy"] = undefined /*out*/;
            resourceInputs["sslProfiles"] = undefined /*out*/;
            resourceInputs["tags"] = undefined /*out*/;
            resourceInputs["trustedClientCertificates"] = undefined /*out*/;
            resourceInputs["trustedRootCertificates"] = undefined /*out*/;
            resourceInputs["type"] = undefined /*out*/;
            resourceInputs["urlPathMaps"] = undefined /*out*/;
            resourceInputs["webApplicationFirewallConfiguration"] = undefined /*out*/;
            resourceInputs["zones"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        const aliasOpts = { aliases: [{ type: "azure-native:network/v20150501preview:ApplicationGateway" }, { type: "azure-native:network/v20150615:ApplicationGateway" }, { type: "azure-native:network/v20160330:ApplicationGateway" }, { type: "azure-native:network/v20160601:ApplicationGateway" }, { type: "azure-native:network/v20160901:ApplicationGateway" }, { type: "azure-native:network/v20161201:ApplicationGateway" }, { type: "azure-native:network/v20170301:ApplicationGateway" }, { type: "azure-native:network/v20170601:ApplicationGateway" }, { type: "azure-native:network/v20170801:ApplicationGateway" }, { type: "azure-native:network/v20170901:ApplicationGateway" }, { type: "azure-native:network/v20171001:ApplicationGateway" }, { type: "azure-native:network/v20171101:ApplicationGateway" }, { type: "azure-native:network/v20180101:ApplicationGateway" }, { type: "azure-native:network/v20180201:ApplicationGateway" }, { type: "azure-native:network/v20180401:ApplicationGateway" }, { type: "azure-native:network/v20180601:ApplicationGateway" }, { type: "azure-native:network/v20180701:ApplicationGateway" }, { type: "azure-native:network/v20180801:ApplicationGateway" }, { type: "azure-native:network/v20181001:ApplicationGateway" }, { type: "azure-native:network/v20181101:ApplicationGateway" }, { type: "azure-native:network/v20181201:ApplicationGateway" }, { type: "azure-native:network/v20190201:ApplicationGateway" }, { type: "azure-native:network/v20190401:ApplicationGateway" }, { type: "azure-native:network/v20190601:ApplicationGateway" }, { type: "azure-native:network/v20190701:ApplicationGateway" }, { type: "azure-native:network/v20190801:ApplicationGateway" }, { type: "azure-native:network/v20190901:ApplicationGateway" }, { type: "azure-native:network/v20191101:ApplicationGateway" }, { type: "azure-native:network/v20191201:ApplicationGateway" }, { type: "azure-native:network/v20200301:ApplicationGateway" }, { type: "azure-native:network/v20200401:ApplicationGateway" }, { type: "azure-native:network/v20200501:ApplicationGateway" }, { type: "azure-native:network/v20200601:ApplicationGateway" }, { type: "azure-native:network/v20200701:ApplicationGateway" }, { type: "azure-native:network/v20200801:ApplicationGateway" }, { type: "azure-native:network/v20201101:ApplicationGateway" }, { type: "azure-native:network/v20210201:ApplicationGateway" }, { type: "azure-native:network/v20210301:ApplicationGateway" }, { type: "azure-native:network/v20210501:ApplicationGateway" }, { type: "azure-native:network/v20210801:ApplicationGateway" }, { type: "azure-native:network/v20220101:ApplicationGateway" }, { type: "azure-native:network/v20220501:ApplicationGateway" }, { type: "azure-native:network/v20220701:ApplicationGateway" }, { type: "azure-native:network/v20220901:ApplicationGateway" }, { type: "azure-native:network/v20221101:ApplicationGateway" }, { type: "azure-native:network/v20230201:ApplicationGateway" }, { type: "azure-native:network/v20230401:ApplicationGateway" }, { type: "azure-native:network/v20230501:ApplicationGateway" }, { type: "azure-native:network/v20230601:ApplicationGateway" }, { type: "azure-native:network/v20230901:ApplicationGateway" }, { type: "azure-native:network/v20231101:ApplicationGateway" }, { type: "azure-native:network/v20240101:ApplicationGateway" }, { type: "azure-native:network/v20240301:ApplicationGateway" }, { type: "azure-native:network/v20240501:ApplicationGateway" }, { type: "azure-native:network/v20240701:ApplicationGateway" }] };
        opts = pulumi.mergeOptions(opts, aliasOpts);
        super(ApplicationGateway.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * The set of arguments for constructing a ApplicationGateway resource.
 */
export interface ApplicationGatewayArgs {
    /**
     * The name of the application gateway.
     */
    applicationGatewayName?: pulumi.Input<string>;
    /**
     * Authentication certificates of the application gateway resource. For default limits, see [Application Gateway limits](https://docs.microsoft.com/azure/azure-subscription-service-limits#application-gateway-limits).
     */
    authenticationCertificates?: pulumi.Input<pulumi.Input<inputs.network.ApplicationGatewayAuthenticationCertificateArgs>[]>;
    /**
     * Autoscale Configuration.
     */
    autoscaleConfiguration?: pulumi.Input<inputs.network.ApplicationGatewayAutoscaleConfigurationArgs>;
    /**
     * Backend address pool of the application gateway resource. For default limits, see [Application Gateway limits](https://docs.microsoft.com/azure/azure-subscription-service-limits#application-gateway-limits).
     */
    backendAddressPools?: pulumi.Input<pulumi.Input<inputs.network.ApplicationGatewayBackendAddressPoolArgs>[]>;
    /**
     * Backend http settings of the application gateway resource. For default limits, see [Application Gateway limits](https://docs.microsoft.com/azure/azure-subscription-service-limits#application-gateway-limits).
     */
    backendHttpSettingsCollection?: pulumi.Input<pulumi.Input<inputs.network.ApplicationGatewayBackendHttpSettingsArgs>[]>;
    /**
     * Backend settings of the application gateway resource. For default limits, see [Application Gateway limits](https://docs.microsoft.com/azure/azure-subscription-service-limits#application-gateway-limits).
     */
    backendSettingsCollection?: pulumi.Input<pulumi.Input<inputs.network.ApplicationGatewayBackendSettingsArgs>[]>;
    /**
     * Custom error configurations of the application gateway resource.
     */
    customErrorConfigurations?: pulumi.Input<pulumi.Input<inputs.network.ApplicationGatewayCustomErrorArgs>[]>;
    /**
     * Whether FIPS is enabled on the application gateway resource.
     */
    enableFips?: pulumi.Input<boolean>;
    /**
     * Whether HTTP2 is enabled on the application gateway resource.
     */
    enableHttp2?: pulumi.Input<boolean>;
    /**
     * Reference to the FirewallPolicy resource.
     */
    firewallPolicy?: pulumi.Input<inputs.network.SubResourceArgs>;
    /**
     * If true, associates a firewall policy with an application gateway regardless whether the policy differs from the WAF Config.
     */
    forceFirewallPolicyAssociation?: pulumi.Input<boolean>;
    /**
     * Frontend IP addresses of the application gateway resource. For default limits, see [Application Gateway limits](https://docs.microsoft.com/azure/azure-subscription-service-limits#application-gateway-limits).
     */
    frontendIPConfigurations?: pulumi.Input<pulumi.Input<inputs.network.ApplicationGatewayFrontendIPConfigurationArgs>[]>;
    /**
     * Frontend ports of the application gateway resource. For default limits, see [Application Gateway limits](https://docs.microsoft.com/azure/azure-subscription-service-limits#application-gateway-limits).
     */
    frontendPorts?: pulumi.Input<pulumi.Input<inputs.network.ApplicationGatewayFrontendPortArgs>[]>;
    /**
     * Subnets of the application gateway resource. For default limits, see [Application Gateway limits](https://docs.microsoft.com/azure/azure-subscription-service-limits#application-gateway-limits).
     */
    gatewayIPConfigurations?: pulumi.Input<pulumi.Input<inputs.network.ApplicationGatewayIPConfigurationArgs>[]>;
    /**
     * Global Configuration.
     */
    globalConfiguration?: pulumi.Input<inputs.network.ApplicationGatewayGlobalConfigurationArgs>;
    /**
     * Http listeners of the application gateway resource. For default limits, see [Application Gateway limits](https://docs.microsoft.com/azure/azure-subscription-service-limits#application-gateway-limits).
     */
    httpListeners?: pulumi.Input<pulumi.Input<inputs.network.ApplicationGatewayHttpListenerArgs>[]>;
    /**
     * Resource ID.
     */
    id?: pulumi.Input<string>;
    /**
     * The identity of the application gateway, if configured.
     */
    identity?: pulumi.Input<inputs.network.ManagedServiceIdentityArgs>;
    /**
     * Listeners of the application gateway resource. For default limits, see [Application Gateway limits](https://docs.microsoft.com/azure/azure-subscription-service-limits#application-gateway-limits).
     */
    listeners?: pulumi.Input<pulumi.Input<inputs.network.ApplicationGatewayListenerArgs>[]>;
    /**
     * Load distribution policies of the application gateway resource.
     */
    loadDistributionPolicies?: pulumi.Input<pulumi.Input<inputs.network.ApplicationGatewayLoadDistributionPolicyArgs>[]>;
    /**
     * Resource location.
     */
    location?: pulumi.Input<string>;
    /**
     * PrivateLink configurations on application gateway.
     */
    privateLinkConfigurations?: pulumi.Input<pulumi.Input<inputs.network.ApplicationGatewayPrivateLinkConfigurationArgs>[]>;
    /**
     * Probes of the application gateway resource.
     */
    probes?: pulumi.Input<pulumi.Input<inputs.network.ApplicationGatewayProbeArgs>[]>;
    /**
     * Redirect configurations of the application gateway resource. For default limits, see [Application Gateway limits](https://docs.microsoft.com/azure/azure-subscription-service-limits#application-gateway-limits).
     */
    redirectConfigurations?: pulumi.Input<pulumi.Input<inputs.network.ApplicationGatewayRedirectConfigurationArgs>[]>;
    /**
     * Request routing rules of the application gateway resource.
     */
    requestRoutingRules?: pulumi.Input<pulumi.Input<inputs.network.ApplicationGatewayRequestRoutingRuleArgs>[]>;
    /**
     * The name of the resource group.
     */
    resourceGroupName: pulumi.Input<string>;
    /**
     * Rewrite rules for the application gateway resource.
     */
    rewriteRuleSets?: pulumi.Input<pulumi.Input<inputs.network.ApplicationGatewayRewriteRuleSetArgs>[]>;
    /**
     * Routing rules of the application gateway resource.
     */
    routingRules?: pulumi.Input<pulumi.Input<inputs.network.ApplicationGatewayRoutingRuleArgs>[]>;
    /**
     * SKU of the application gateway resource.
     */
    sku?: pulumi.Input<inputs.network.ApplicationGatewaySkuArgs>;
    /**
     * SSL certificates of the application gateway resource. For default limits, see [Application Gateway limits](https://docs.microsoft.com/azure/azure-subscription-service-limits#application-gateway-limits).
     */
    sslCertificates?: pulumi.Input<pulumi.Input<inputs.network.ApplicationGatewaySslCertificateArgs>[]>;
    /**
     * SSL policy of the application gateway resource.
     */
    sslPolicy?: pulumi.Input<inputs.network.ApplicationGatewaySslPolicyArgs>;
    /**
     * SSL profiles of the application gateway resource. For default limits, see [Application Gateway limits](https://docs.microsoft.com/azure/azure-subscription-service-limits#application-gateway-limits).
     */
    sslProfiles?: pulumi.Input<pulumi.Input<inputs.network.ApplicationGatewaySslProfileArgs>[]>;
    /**
     * Resource tags.
     */
    tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * Trusted client certificates of the application gateway resource. For default limits, see [Application Gateway limits](https://docs.microsoft.com/azure/azure-subscription-service-limits#application-gateway-limits).
     */
    trustedClientCertificates?: pulumi.Input<pulumi.Input<inputs.network.ApplicationGatewayTrustedClientCertificateArgs>[]>;
    /**
     * Trusted Root certificates of the application gateway resource. For default limits, see [Application Gateway limits](https://docs.microsoft.com/azure/azure-subscription-service-limits#application-gateway-limits).
     */
    trustedRootCertificates?: pulumi.Input<pulumi.Input<inputs.network.ApplicationGatewayTrustedRootCertificateArgs>[]>;
    /**
     * URL path map of the application gateway resource. For default limits, see [Application Gateway limits](https://docs.microsoft.com/azure/azure-subscription-service-limits#application-gateway-limits).
     */
    urlPathMaps?: pulumi.Input<pulumi.Input<inputs.network.ApplicationGatewayUrlPathMapArgs>[]>;
    /**
     * Web application firewall configuration.
     */
    webApplicationFirewallConfiguration?: pulumi.Input<inputs.network.ApplicationGatewayWebApplicationFirewallConfigurationArgs>;
    /**
     * A list of availability zones denoting where the resource needs to come from.
     */
    zones?: pulumi.Input<pulumi.Input<string>[]>;
}
