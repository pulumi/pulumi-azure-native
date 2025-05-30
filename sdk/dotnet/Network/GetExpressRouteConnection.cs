// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.Network
{
    public static class GetExpressRouteConnection
    {
        /// <summary>
        /// Gets the specified ExpressRouteConnection.
        /// 
        /// Uses Azure REST API version 2024-05-01.
        /// 
        /// Other available API versions: 2018-08-01, 2018-10-01, 2018-11-01, 2018-12-01, 2019-02-01, 2019-04-01, 2019-06-01, 2019-07-01, 2019-08-01, 2019-09-01, 2019-11-01, 2019-12-01, 2020-03-01, 2020-04-01, 2020-05-01, 2020-06-01, 2020-07-01, 2020-08-01, 2020-11-01, 2021-02-01, 2021-03-01, 2021-05-01, 2021-08-01, 2022-01-01, 2022-05-01, 2022-07-01, 2022-09-01, 2022-11-01, 2023-02-01, 2023-04-01, 2023-05-01, 2023-06-01, 2023-09-01, 2023-11-01, 2024-01-01, 2024-03-01, 2024-07-01. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native network [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
        /// </summary>
        public static Task<GetExpressRouteConnectionResult> InvokeAsync(GetExpressRouteConnectionArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetExpressRouteConnectionResult>("azure-native:network:getExpressRouteConnection", args ?? new GetExpressRouteConnectionArgs(), options.WithDefaults());

        /// <summary>
        /// Gets the specified ExpressRouteConnection.
        /// 
        /// Uses Azure REST API version 2024-05-01.
        /// 
        /// Other available API versions: 2018-08-01, 2018-10-01, 2018-11-01, 2018-12-01, 2019-02-01, 2019-04-01, 2019-06-01, 2019-07-01, 2019-08-01, 2019-09-01, 2019-11-01, 2019-12-01, 2020-03-01, 2020-04-01, 2020-05-01, 2020-06-01, 2020-07-01, 2020-08-01, 2020-11-01, 2021-02-01, 2021-03-01, 2021-05-01, 2021-08-01, 2022-01-01, 2022-05-01, 2022-07-01, 2022-09-01, 2022-11-01, 2023-02-01, 2023-04-01, 2023-05-01, 2023-06-01, 2023-09-01, 2023-11-01, 2024-01-01, 2024-03-01, 2024-07-01. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native network [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
        /// </summary>
        public static Output<GetExpressRouteConnectionResult> Invoke(GetExpressRouteConnectionInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetExpressRouteConnectionResult>("azure-native:network:getExpressRouteConnection", args ?? new GetExpressRouteConnectionInvokeArgs(), options.WithDefaults());

        /// <summary>
        /// Gets the specified ExpressRouteConnection.
        /// 
        /// Uses Azure REST API version 2024-05-01.
        /// 
        /// Other available API versions: 2018-08-01, 2018-10-01, 2018-11-01, 2018-12-01, 2019-02-01, 2019-04-01, 2019-06-01, 2019-07-01, 2019-08-01, 2019-09-01, 2019-11-01, 2019-12-01, 2020-03-01, 2020-04-01, 2020-05-01, 2020-06-01, 2020-07-01, 2020-08-01, 2020-11-01, 2021-02-01, 2021-03-01, 2021-05-01, 2021-08-01, 2022-01-01, 2022-05-01, 2022-07-01, 2022-09-01, 2022-11-01, 2023-02-01, 2023-04-01, 2023-05-01, 2023-06-01, 2023-09-01, 2023-11-01, 2024-01-01, 2024-03-01, 2024-07-01. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native network [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
        /// </summary>
        public static Output<GetExpressRouteConnectionResult> Invoke(GetExpressRouteConnectionInvokeArgs args, InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<GetExpressRouteConnectionResult>("azure-native:network:getExpressRouteConnection", args ?? new GetExpressRouteConnectionInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetExpressRouteConnectionArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the ExpressRoute connection.
        /// </summary>
        [Input("connectionName", required: true)]
        public string ConnectionName { get; set; } = null!;

        /// <summary>
        /// The name of the ExpressRoute gateway.
        /// </summary>
        [Input("expressRouteGatewayName", required: true)]
        public string ExpressRouteGatewayName { get; set; } = null!;

        /// <summary>
        /// The name of the resource group.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public string ResourceGroupName { get; set; } = null!;

        public GetExpressRouteConnectionArgs()
        {
        }
        public static new GetExpressRouteConnectionArgs Empty => new GetExpressRouteConnectionArgs();
    }

    public sealed class GetExpressRouteConnectionInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the ExpressRoute connection.
        /// </summary>
        [Input("connectionName", required: true)]
        public Input<string> ConnectionName { get; set; } = null!;

        /// <summary>
        /// The name of the ExpressRoute gateway.
        /// </summary>
        [Input("expressRouteGatewayName", required: true)]
        public Input<string> ExpressRouteGatewayName { get; set; } = null!;

        /// <summary>
        /// The name of the resource group.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        public GetExpressRouteConnectionInvokeArgs()
        {
        }
        public static new GetExpressRouteConnectionInvokeArgs Empty => new GetExpressRouteConnectionInvokeArgs();
    }


    [OutputType]
    public sealed class GetExpressRouteConnectionResult
    {
        /// <summary>
        /// Authorization key to establish the connection.
        /// </summary>
        public readonly string? AuthorizationKey;
        /// <summary>
        /// The Azure API version of the resource.
        /// </summary>
        public readonly string AzureApiVersion;
        /// <summary>
        /// Enable internet security.
        /// </summary>
        public readonly bool? EnableInternetSecurity;
        /// <summary>
        /// Bypass the ExpressRoute gateway when accessing private-links. ExpressRoute FastPath (expressRouteGatewayBypass) must be enabled.
        /// </summary>
        public readonly bool? EnablePrivateLinkFastPath;
        /// <summary>
        /// The ExpressRoute circuit peering.
        /// </summary>
        public readonly Outputs.ExpressRouteCircuitPeeringIdResponse ExpressRouteCircuitPeering;
        /// <summary>
        /// Enable FastPath to vWan Firewall hub.
        /// </summary>
        public readonly bool? ExpressRouteGatewayBypass;
        /// <summary>
        /// Resource ID.
        /// </summary>
        public readonly string? Id;
        /// <summary>
        /// The name of the resource.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// The provisioning state of the express route connection resource.
        /// </summary>
        public readonly string ProvisioningState;
        /// <summary>
        /// The Routing Configuration indicating the associated and propagated route tables on this connection.
        /// </summary>
        public readonly Outputs.RoutingConfigurationResponse? RoutingConfiguration;
        /// <summary>
        /// The routing weight associated to the connection.
        /// </summary>
        public readonly int? RoutingWeight;

        [OutputConstructor]
        private GetExpressRouteConnectionResult(
            string? authorizationKey,

            string azureApiVersion,

            bool? enableInternetSecurity,

            bool? enablePrivateLinkFastPath,

            Outputs.ExpressRouteCircuitPeeringIdResponse expressRouteCircuitPeering,

            bool? expressRouteGatewayBypass,

            string? id,

            string name,

            string provisioningState,

            Outputs.RoutingConfigurationResponse? routingConfiguration,

            int? routingWeight)
        {
            AuthorizationKey = authorizationKey;
            AzureApiVersion = azureApiVersion;
            EnableInternetSecurity = enableInternetSecurity;
            EnablePrivateLinkFastPath = enablePrivateLinkFastPath;
            ExpressRouteCircuitPeering = expressRouteCircuitPeering;
            ExpressRouteGatewayBypass = expressRouteGatewayBypass;
            Id = id;
            Name = name;
            ProvisioningState = provisioningState;
            RoutingConfiguration = routingConfiguration;
            RoutingWeight = routingWeight;
        }
    }
}
