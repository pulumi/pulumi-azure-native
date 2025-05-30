// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.Network
{
    /// <summary>
    /// ExpressRouteConnection resource.
    /// 
    /// Uses Azure REST API version 2024-05-01. In version 2.x of the Azure Native provider, it used API version 2023-02-01.
    /// 
    /// Other available API versions: 2018-08-01, 2018-10-01, 2018-11-01, 2018-12-01, 2019-02-01, 2019-04-01, 2019-06-01, 2019-07-01, 2019-08-01, 2019-09-01, 2019-11-01, 2019-12-01, 2020-03-01, 2020-04-01, 2020-05-01, 2020-06-01, 2020-07-01, 2020-08-01, 2020-11-01, 2021-02-01, 2021-03-01, 2021-05-01, 2021-08-01, 2022-01-01, 2022-05-01, 2022-07-01, 2022-09-01, 2022-11-01, 2023-02-01, 2023-04-01, 2023-05-01, 2023-06-01, 2023-09-01, 2023-11-01, 2024-01-01, 2024-03-01, 2024-07-01. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native network [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
    /// </summary>
    [AzureNativeResourceType("azure-native:network:ExpressRouteConnection")]
    public partial class ExpressRouteConnection : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Authorization key to establish the connection.
        /// </summary>
        [Output("authorizationKey")]
        public Output<string?> AuthorizationKey { get; private set; } = null!;

        /// <summary>
        /// The Azure API version of the resource.
        /// </summary>
        [Output("azureApiVersion")]
        public Output<string> AzureApiVersion { get; private set; } = null!;

        /// <summary>
        /// Enable internet security.
        /// </summary>
        [Output("enableInternetSecurity")]
        public Output<bool?> EnableInternetSecurity { get; private set; } = null!;

        /// <summary>
        /// Bypass the ExpressRoute gateway when accessing private-links. ExpressRoute FastPath (expressRouteGatewayBypass) must be enabled.
        /// </summary>
        [Output("enablePrivateLinkFastPath")]
        public Output<bool?> EnablePrivateLinkFastPath { get; private set; } = null!;

        /// <summary>
        /// The ExpressRoute circuit peering.
        /// </summary>
        [Output("expressRouteCircuitPeering")]
        public Output<Outputs.ExpressRouteCircuitPeeringIdResponse> ExpressRouteCircuitPeering { get; private set; } = null!;

        /// <summary>
        /// Enable FastPath to vWan Firewall hub.
        /// </summary>
        [Output("expressRouteGatewayBypass")]
        public Output<bool?> ExpressRouteGatewayBypass { get; private set; } = null!;

        /// <summary>
        /// The name of the resource.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The provisioning state of the express route connection resource.
        /// </summary>
        [Output("provisioningState")]
        public Output<string> ProvisioningState { get; private set; } = null!;

        /// <summary>
        /// The Routing Configuration indicating the associated and propagated route tables on this connection.
        /// </summary>
        [Output("routingConfiguration")]
        public Output<Outputs.RoutingConfigurationResponse?> RoutingConfiguration { get; private set; } = null!;

        /// <summary>
        /// The routing weight associated to the connection.
        /// </summary>
        [Output("routingWeight")]
        public Output<int?> RoutingWeight { get; private set; } = null!;


        /// <summary>
        /// Create a ExpressRouteConnection resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public ExpressRouteConnection(string name, ExpressRouteConnectionArgs args, CustomResourceOptions? options = null)
            : base("azure-native:network:ExpressRouteConnection", name, args ?? new ExpressRouteConnectionArgs(), MakeResourceOptions(options, ""))
        {
        }

        private ExpressRouteConnection(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("azure-native:network:ExpressRouteConnection", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                Aliases =
                {
                    new global::Pulumi.Alias { Type = "azure-native:network/v20180801:ExpressRouteConnection" },
                    new global::Pulumi.Alias { Type = "azure-native:network/v20181001:ExpressRouteConnection" },
                    new global::Pulumi.Alias { Type = "azure-native:network/v20181101:ExpressRouteConnection" },
                    new global::Pulumi.Alias { Type = "azure-native:network/v20181201:ExpressRouteConnection" },
                    new global::Pulumi.Alias { Type = "azure-native:network/v20190201:ExpressRouteConnection" },
                    new global::Pulumi.Alias { Type = "azure-native:network/v20190401:ExpressRouteConnection" },
                    new global::Pulumi.Alias { Type = "azure-native:network/v20190601:ExpressRouteConnection" },
                    new global::Pulumi.Alias { Type = "azure-native:network/v20190701:ExpressRouteConnection" },
                    new global::Pulumi.Alias { Type = "azure-native:network/v20190801:ExpressRouteConnection" },
                    new global::Pulumi.Alias { Type = "azure-native:network/v20190901:ExpressRouteConnection" },
                    new global::Pulumi.Alias { Type = "azure-native:network/v20191101:ExpressRouteConnection" },
                    new global::Pulumi.Alias { Type = "azure-native:network/v20191201:ExpressRouteConnection" },
                    new global::Pulumi.Alias { Type = "azure-native:network/v20200301:ExpressRouteConnection" },
                    new global::Pulumi.Alias { Type = "azure-native:network/v20200401:ExpressRouteConnection" },
                    new global::Pulumi.Alias { Type = "azure-native:network/v20200501:ExpressRouteConnection" },
                    new global::Pulumi.Alias { Type = "azure-native:network/v20200601:ExpressRouteConnection" },
                    new global::Pulumi.Alias { Type = "azure-native:network/v20200701:ExpressRouteConnection" },
                    new global::Pulumi.Alias { Type = "azure-native:network/v20200801:ExpressRouteConnection" },
                    new global::Pulumi.Alias { Type = "azure-native:network/v20201101:ExpressRouteConnection" },
                    new global::Pulumi.Alias { Type = "azure-native:network/v20210201:ExpressRouteConnection" },
                    new global::Pulumi.Alias { Type = "azure-native:network/v20210301:ExpressRouteConnection" },
                    new global::Pulumi.Alias { Type = "azure-native:network/v20210501:ExpressRouteConnection" },
                    new global::Pulumi.Alias { Type = "azure-native:network/v20210801:ExpressRouteConnection" },
                    new global::Pulumi.Alias { Type = "azure-native:network/v20220101:ExpressRouteConnection" },
                    new global::Pulumi.Alias { Type = "azure-native:network/v20220501:ExpressRouteConnection" },
                    new global::Pulumi.Alias { Type = "azure-native:network/v20220701:ExpressRouteConnection" },
                    new global::Pulumi.Alias { Type = "azure-native:network/v20220901:ExpressRouteConnection" },
                    new global::Pulumi.Alias { Type = "azure-native:network/v20221101:ExpressRouteConnection" },
                    new global::Pulumi.Alias { Type = "azure-native:network/v20230201:ExpressRouteConnection" },
                    new global::Pulumi.Alias { Type = "azure-native:network/v20230401:ExpressRouteConnection" },
                    new global::Pulumi.Alias { Type = "azure-native:network/v20230501:ExpressRouteConnection" },
                    new global::Pulumi.Alias { Type = "azure-native:network/v20230601:ExpressRouteConnection" },
                    new global::Pulumi.Alias { Type = "azure-native:network/v20230901:ExpressRouteConnection" },
                    new global::Pulumi.Alias { Type = "azure-native:network/v20231101:ExpressRouteConnection" },
                    new global::Pulumi.Alias { Type = "azure-native:network/v20240101:ExpressRouteConnection" },
                    new global::Pulumi.Alias { Type = "azure-native:network/v20240301:ExpressRouteConnection" },
                    new global::Pulumi.Alias { Type = "azure-native:network/v20240501:ExpressRouteConnection" },
                    new global::Pulumi.Alias { Type = "azure-native:network/v20240701:ExpressRouteConnection" },
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing ExpressRouteConnection resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static ExpressRouteConnection Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new ExpressRouteConnection(name, id, options);
        }
    }

    public sealed class ExpressRouteConnectionArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Authorization key to establish the connection.
        /// </summary>
        [Input("authorizationKey")]
        public Input<string>? AuthorizationKey { get; set; }

        /// <summary>
        /// The name of the connection subresource.
        /// </summary>
        [Input("connectionName")]
        public Input<string>? ConnectionName { get; set; }

        /// <summary>
        /// Enable internet security.
        /// </summary>
        [Input("enableInternetSecurity")]
        public Input<bool>? EnableInternetSecurity { get; set; }

        /// <summary>
        /// Bypass the ExpressRoute gateway when accessing private-links. ExpressRoute FastPath (expressRouteGatewayBypass) must be enabled.
        /// </summary>
        [Input("enablePrivateLinkFastPath")]
        public Input<bool>? EnablePrivateLinkFastPath { get; set; }

        /// <summary>
        /// The ExpressRoute circuit peering.
        /// </summary>
        [Input("expressRouteCircuitPeering", required: true)]
        public Input<Inputs.ExpressRouteCircuitPeeringIdArgs> ExpressRouteCircuitPeering { get; set; } = null!;

        /// <summary>
        /// Enable FastPath to vWan Firewall hub.
        /// </summary>
        [Input("expressRouteGatewayBypass")]
        public Input<bool>? ExpressRouteGatewayBypass { get; set; }

        /// <summary>
        /// The name of the ExpressRoute gateway.
        /// </summary>
        [Input("expressRouteGatewayName", required: true)]
        public Input<string> ExpressRouteGatewayName { get; set; } = null!;

        /// <summary>
        /// Resource ID.
        /// </summary>
        [Input("id")]
        public Input<string>? Id { get; set; }

        /// <summary>
        /// The name of the resource.
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        /// <summary>
        /// The name of the resource group.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        /// <summary>
        /// The Routing Configuration indicating the associated and propagated route tables on this connection.
        /// </summary>
        [Input("routingConfiguration")]
        public Input<Inputs.RoutingConfigurationArgs>? RoutingConfiguration { get; set; }

        /// <summary>
        /// The routing weight associated to the connection.
        /// </summary>
        [Input("routingWeight")]
        public Input<int>? RoutingWeight { get; set; }

        public ExpressRouteConnectionArgs()
        {
        }
        public static new ExpressRouteConnectionArgs Empty => new ExpressRouteConnectionArgs();
    }
}
