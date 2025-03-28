// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.ContainerService.V20241002Preview.Inputs
{

    /// <summary>
    /// Istio egress gateway configuration.
    /// </summary>
    public sealed class IstioEgressGatewayArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Whether to enable the egress gateway.
        /// </summary>
        [Input("enabled", required: true)]
        public Input<bool> Enabled { get; set; } = null!;

        /// <summary>
        /// Name of the gateway configuration custom resource for the Istio add-on egress gateway. Must be specified when enabling the Istio egress gateway. Must be deployed in the same namespace that the Istio egress gateway will be deployed in.
        /// </summary>
        [Input("gatewayConfigurationName")]
        public Input<string>? GatewayConfigurationName { get; set; }

        /// <summary>
        /// Name of the Istio add-on egress gateway.
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        /// <summary>
        /// Namespace that the Istio add-on egress gateway should be deployed in. If unspecified, the default is aks-istio-egress.
        /// </summary>
        [Input("namespace")]
        public Input<string>? Namespace { get; set; }

        public IstioEgressGatewayArgs()
        {
        }
        public static new IstioEgressGatewayArgs Empty => new IstioEgressGatewayArgs();
    }
}
