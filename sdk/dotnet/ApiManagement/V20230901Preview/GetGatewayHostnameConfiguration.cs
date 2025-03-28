// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.ApiManagement.V20230901Preview
{
    public static class GetGatewayHostnameConfiguration
    {
        /// <summary>
        /// Get details of a hostname configuration
        /// </summary>
        public static Task<GetGatewayHostnameConfigurationResult> InvokeAsync(GetGatewayHostnameConfigurationArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetGatewayHostnameConfigurationResult>("azure-native:apimanagement/v20230901preview:getGatewayHostnameConfiguration", args ?? new GetGatewayHostnameConfigurationArgs(), options.WithDefaults());

        /// <summary>
        /// Get details of a hostname configuration
        /// </summary>
        public static Output<GetGatewayHostnameConfigurationResult> Invoke(GetGatewayHostnameConfigurationInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetGatewayHostnameConfigurationResult>("azure-native:apimanagement/v20230901preview:getGatewayHostnameConfiguration", args ?? new GetGatewayHostnameConfigurationInvokeArgs(), options.WithDefaults());

        /// <summary>
        /// Get details of a hostname configuration
        /// </summary>
        public static Output<GetGatewayHostnameConfigurationResult> Invoke(GetGatewayHostnameConfigurationInvokeArgs args, InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<GetGatewayHostnameConfigurationResult>("azure-native:apimanagement/v20230901preview:getGatewayHostnameConfiguration", args ?? new GetGatewayHostnameConfigurationInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetGatewayHostnameConfigurationArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Gateway entity identifier. Must be unique in the current API Management service instance. Must not have value 'managed'
        /// </summary>
        [Input("gatewayId", required: true)]
        public string GatewayId { get; set; } = null!;

        /// <summary>
        /// Gateway hostname configuration identifier. Must be unique in the scope of parent Gateway entity.
        /// </summary>
        [Input("hcId", required: true)]
        public string HcId { get; set; } = null!;

        /// <summary>
        /// The name of the resource group. The name is case insensitive.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public string ResourceGroupName { get; set; } = null!;

        /// <summary>
        /// The name of the API Management service.
        /// </summary>
        [Input("serviceName", required: true)]
        public string ServiceName { get; set; } = null!;

        public GetGatewayHostnameConfigurationArgs()
        {
        }
        public static new GetGatewayHostnameConfigurationArgs Empty => new GetGatewayHostnameConfigurationArgs();
    }

    public sealed class GetGatewayHostnameConfigurationInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Gateway entity identifier. Must be unique in the current API Management service instance. Must not have value 'managed'
        /// </summary>
        [Input("gatewayId", required: true)]
        public Input<string> GatewayId { get; set; } = null!;

        /// <summary>
        /// Gateway hostname configuration identifier. Must be unique in the scope of parent Gateway entity.
        /// </summary>
        [Input("hcId", required: true)]
        public Input<string> HcId { get; set; } = null!;

        /// <summary>
        /// The name of the resource group. The name is case insensitive.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        /// <summary>
        /// The name of the API Management service.
        /// </summary>
        [Input("serviceName", required: true)]
        public Input<string> ServiceName { get; set; } = null!;

        public GetGatewayHostnameConfigurationInvokeArgs()
        {
        }
        public static new GetGatewayHostnameConfigurationInvokeArgs Empty => new GetGatewayHostnameConfigurationInvokeArgs();
    }


    [OutputType]
    public sealed class GetGatewayHostnameConfigurationResult
    {
        /// <summary>
        /// Identifier of Certificate entity that will be used for TLS connection establishment
        /// </summary>
        public readonly string? CertificateId;
        /// <summary>
        /// Hostname value. Supports valid domain name, partial or full wildcard
        /// </summary>
        public readonly string? Hostname;
        /// <summary>
        /// Specifies if HTTP/2.0 is supported
        /// </summary>
        public readonly bool? Http2Enabled;
        /// <summary>
        /// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// The name of the resource
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// Determines whether gateway requests client certificate
        /// </summary>
        public readonly bool? NegotiateClientCertificate;
        /// <summary>
        /// Specifies if TLS 1.0 is supported
        /// </summary>
        public readonly bool? Tls10Enabled;
        /// <summary>
        /// Specifies if TLS 1.1 is supported
        /// </summary>
        public readonly bool? Tls11Enabled;
        /// <summary>
        /// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
        /// </summary>
        public readonly string Type;

        [OutputConstructor]
        private GetGatewayHostnameConfigurationResult(
            string? certificateId,

            string? hostname,

            bool? http2Enabled,

            string id,

            string name,

            bool? negotiateClientCertificate,

            bool? tls10Enabled,

            bool? tls11Enabled,

            string type)
        {
            CertificateId = certificateId;
            Hostname = hostname;
            Http2Enabled = http2Enabled;
            Id = id;
            Name = name;
            NegotiateClientCertificate = negotiateClientCertificate;
            Tls10Enabled = tls10Enabled;
            Tls11Enabled = tls11Enabled;
            Type = type;
        }
    }
}
