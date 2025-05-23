// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.HybridNetwork.Inputs
{

    /// <summary>
    /// Azure arc kubernetes deploy mapping rule profile.
    /// </summary>
    public sealed class AzureArcKubernetesDeployMappingRuleProfileArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The application enablement.
        /// </summary>
        [Input("applicationEnablement")]
        public InputUnion<string, Pulumi.AzureNative.HybridNetwork.ApplicationEnablement>? ApplicationEnablement { get; set; }

        /// <summary>
        /// The helm mapping rule profile.
        /// </summary>
        [Input("helmMappingRuleProfile")]
        public Input<Inputs.HelmMappingRuleProfileArgs>? HelmMappingRuleProfile { get; set; }

        public AzureArcKubernetesDeployMappingRuleProfileArgs()
        {
        }
        public static new AzureArcKubernetesDeployMappingRuleProfileArgs Empty => new AzureArcKubernetesDeployMappingRuleProfileArgs();
    }
}
