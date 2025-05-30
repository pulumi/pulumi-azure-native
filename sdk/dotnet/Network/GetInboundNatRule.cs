// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.Network
{
    public static class GetInboundNatRule
    {
        /// <summary>
        /// Gets the specified load balancer inbound NAT rule.
        /// 
        /// Uses Azure REST API version 2024-05-01.
        /// 
        /// Other available API versions: 2018-06-01, 2018-07-01, 2018-08-01, 2018-10-01, 2018-11-01, 2018-12-01, 2019-02-01, 2019-04-01, 2019-06-01, 2019-07-01, 2019-08-01, 2019-09-01, 2019-11-01, 2019-12-01, 2020-03-01, 2020-04-01, 2020-05-01, 2020-06-01, 2020-07-01, 2020-08-01, 2020-11-01, 2021-02-01, 2021-03-01, 2021-05-01, 2021-08-01, 2022-01-01, 2022-05-01, 2022-07-01, 2022-09-01, 2022-11-01, 2023-02-01, 2023-04-01, 2023-05-01, 2023-06-01, 2023-09-01, 2023-11-01, 2024-01-01, 2024-03-01, 2024-07-01. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native network [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
        /// </summary>
        public static Task<GetInboundNatRuleResult> InvokeAsync(GetInboundNatRuleArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetInboundNatRuleResult>("azure-native:network:getInboundNatRule", args ?? new GetInboundNatRuleArgs(), options.WithDefaults());

        /// <summary>
        /// Gets the specified load balancer inbound NAT rule.
        /// 
        /// Uses Azure REST API version 2024-05-01.
        /// 
        /// Other available API versions: 2018-06-01, 2018-07-01, 2018-08-01, 2018-10-01, 2018-11-01, 2018-12-01, 2019-02-01, 2019-04-01, 2019-06-01, 2019-07-01, 2019-08-01, 2019-09-01, 2019-11-01, 2019-12-01, 2020-03-01, 2020-04-01, 2020-05-01, 2020-06-01, 2020-07-01, 2020-08-01, 2020-11-01, 2021-02-01, 2021-03-01, 2021-05-01, 2021-08-01, 2022-01-01, 2022-05-01, 2022-07-01, 2022-09-01, 2022-11-01, 2023-02-01, 2023-04-01, 2023-05-01, 2023-06-01, 2023-09-01, 2023-11-01, 2024-01-01, 2024-03-01, 2024-07-01. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native network [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
        /// </summary>
        public static Output<GetInboundNatRuleResult> Invoke(GetInboundNatRuleInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetInboundNatRuleResult>("azure-native:network:getInboundNatRule", args ?? new GetInboundNatRuleInvokeArgs(), options.WithDefaults());

        /// <summary>
        /// Gets the specified load balancer inbound NAT rule.
        /// 
        /// Uses Azure REST API version 2024-05-01.
        /// 
        /// Other available API versions: 2018-06-01, 2018-07-01, 2018-08-01, 2018-10-01, 2018-11-01, 2018-12-01, 2019-02-01, 2019-04-01, 2019-06-01, 2019-07-01, 2019-08-01, 2019-09-01, 2019-11-01, 2019-12-01, 2020-03-01, 2020-04-01, 2020-05-01, 2020-06-01, 2020-07-01, 2020-08-01, 2020-11-01, 2021-02-01, 2021-03-01, 2021-05-01, 2021-08-01, 2022-01-01, 2022-05-01, 2022-07-01, 2022-09-01, 2022-11-01, 2023-02-01, 2023-04-01, 2023-05-01, 2023-06-01, 2023-09-01, 2023-11-01, 2024-01-01, 2024-03-01, 2024-07-01. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native network [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
        /// </summary>
        public static Output<GetInboundNatRuleResult> Invoke(GetInboundNatRuleInvokeArgs args, InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<GetInboundNatRuleResult>("azure-native:network:getInboundNatRule", args ?? new GetInboundNatRuleInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetInboundNatRuleArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Expands referenced resources.
        /// </summary>
        [Input("expand")]
        public string? Expand { get; set; }

        /// <summary>
        /// The name of the inbound NAT rule.
        /// </summary>
        [Input("inboundNatRuleName", required: true)]
        public string InboundNatRuleName { get; set; } = null!;

        /// <summary>
        /// The name of the load balancer.
        /// </summary>
        [Input("loadBalancerName", required: true)]
        public string LoadBalancerName { get; set; } = null!;

        /// <summary>
        /// The name of the resource group.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public string ResourceGroupName { get; set; } = null!;

        public GetInboundNatRuleArgs()
        {
        }
        public static new GetInboundNatRuleArgs Empty => new GetInboundNatRuleArgs();
    }

    public sealed class GetInboundNatRuleInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Expands referenced resources.
        /// </summary>
        [Input("expand")]
        public Input<string>? Expand { get; set; }

        /// <summary>
        /// The name of the inbound NAT rule.
        /// </summary>
        [Input("inboundNatRuleName", required: true)]
        public Input<string> InboundNatRuleName { get; set; } = null!;

        /// <summary>
        /// The name of the load balancer.
        /// </summary>
        [Input("loadBalancerName", required: true)]
        public Input<string> LoadBalancerName { get; set; } = null!;

        /// <summary>
        /// The name of the resource group.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        public GetInboundNatRuleInvokeArgs()
        {
        }
        public static new GetInboundNatRuleInvokeArgs Empty => new GetInboundNatRuleInvokeArgs();
    }


    [OutputType]
    public sealed class GetInboundNatRuleResult
    {
        /// <summary>
        /// The Azure API version of the resource.
        /// </summary>
        public readonly string AzureApiVersion;
        /// <summary>
        /// A reference to backendAddressPool resource.
        /// </summary>
        public readonly Outputs.SubResourceResponse? BackendAddressPool;
        /// <summary>
        /// A reference to a private IP address defined on a network interface of a VM. Traffic sent to the frontend port of each of the frontend IP configurations is forwarded to the backend IP.
        /// </summary>
        public readonly Outputs.NetworkInterfaceIPConfigurationResponse BackendIPConfiguration;
        /// <summary>
        /// The port used for the internal endpoint. Acceptable values range from 1 to 65535.
        /// </summary>
        public readonly int? BackendPort;
        /// <summary>
        /// Configures a virtual machine's endpoint for the floating IP capability required to configure a SQL AlwaysOn Availability Group. This setting is required when using the SQL AlwaysOn Availability Groups in SQL server. This setting can't be changed after you create the endpoint.
        /// </summary>
        public readonly bool? EnableFloatingIP;
        /// <summary>
        /// Receive bidirectional TCP Reset on TCP flow idle timeout or unexpected connection termination. This element is only used when the protocol is set to TCP.
        /// </summary>
        public readonly bool? EnableTcpReset;
        /// <summary>
        /// A unique read-only string that changes whenever the resource is updated.
        /// </summary>
        public readonly string Etag;
        /// <summary>
        /// A reference to frontend IP addresses.
        /// </summary>
        public readonly Outputs.SubResourceResponse? FrontendIPConfiguration;
        /// <summary>
        /// The port for the external endpoint. Port numbers for each rule must be unique within the Load Balancer. Acceptable values range from 1 to 65534.
        /// </summary>
        public readonly int? FrontendPort;
        /// <summary>
        /// The port range end for the external endpoint. This property is used together with BackendAddressPool and FrontendPortRangeStart. Individual inbound NAT rule port mappings will be created for each backend address from BackendAddressPool. Acceptable values range from 1 to 65534.
        /// </summary>
        public readonly int? FrontendPortRangeEnd;
        /// <summary>
        /// The port range start for the external endpoint. This property is used together with BackendAddressPool and FrontendPortRangeEnd. Individual inbound NAT rule port mappings will be created for each backend address from BackendAddressPool. Acceptable values range from 1 to 65534.
        /// </summary>
        public readonly int? FrontendPortRangeStart;
        /// <summary>
        /// Resource ID.
        /// </summary>
        public readonly string? Id;
        /// <summary>
        /// The timeout for the TCP idle connection. The value can be set between 4 and 30 minutes. The default value is 4 minutes. This element is only used when the protocol is set to TCP.
        /// </summary>
        public readonly int? IdleTimeoutInMinutes;
        /// <summary>
        /// The name of the resource that is unique within the set of inbound NAT rules used by the load balancer. This name can be used to access the resource.
        /// </summary>
        public readonly string? Name;
        /// <summary>
        /// The reference to the transport protocol used by the load balancing rule.
        /// </summary>
        public readonly string? Protocol;
        /// <summary>
        /// The provisioning state of the inbound NAT rule resource.
        /// </summary>
        public readonly string ProvisioningState;
        /// <summary>
        /// Type of the resource.
        /// </summary>
        public readonly string Type;

        [OutputConstructor]
        private GetInboundNatRuleResult(
            string azureApiVersion,

            Outputs.SubResourceResponse? backendAddressPool,

            Outputs.NetworkInterfaceIPConfigurationResponse backendIPConfiguration,

            int? backendPort,

            bool? enableFloatingIP,

            bool? enableTcpReset,

            string etag,

            Outputs.SubResourceResponse? frontendIPConfiguration,

            int? frontendPort,

            int? frontendPortRangeEnd,

            int? frontendPortRangeStart,

            string? id,

            int? idleTimeoutInMinutes,

            string? name,

            string? protocol,

            string provisioningState,

            string type)
        {
            AzureApiVersion = azureApiVersion;
            BackendAddressPool = backendAddressPool;
            BackendIPConfiguration = backendIPConfiguration;
            BackendPort = backendPort;
            EnableFloatingIP = enableFloatingIP;
            EnableTcpReset = enableTcpReset;
            Etag = etag;
            FrontendIPConfiguration = frontendIPConfiguration;
            FrontendPort = frontendPort;
            FrontendPortRangeEnd = frontendPortRangeEnd;
            FrontendPortRangeStart = frontendPortRangeStart;
            Id = id;
            IdleTimeoutInMinutes = idleTimeoutInMinutes;
            Name = name;
            Protocol = protocol;
            ProvisioningState = provisioningState;
            Type = type;
        }
    }
}
