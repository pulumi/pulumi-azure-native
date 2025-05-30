// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.HybridNetwork.Outputs
{

    /// <summary>
    /// Azure Arc kubernetes network function template.
    /// </summary>
    [OutputType]
    public sealed class AzureArcKubernetesNetworkFunctionTemplateResponse
    {
        /// <summary>
        /// Network function applications.
        /// </summary>
        public readonly ImmutableArray<Outputs.AzureArcKubernetesHelmApplicationResponse> NetworkFunctionApplications;
        /// <summary>
        /// The network function type.
        /// Expected value is 'AzureArcKubernetes'.
        /// </summary>
        public readonly string NfviType;

        [OutputConstructor]
        private AzureArcKubernetesNetworkFunctionTemplateResponse(
            ImmutableArray<Outputs.AzureArcKubernetesHelmApplicationResponse> networkFunctionApplications,

            string nfviType)
        {
            NetworkFunctionApplications = networkFunctionApplications;
            NfviType = nfviType;
        }
    }
}
