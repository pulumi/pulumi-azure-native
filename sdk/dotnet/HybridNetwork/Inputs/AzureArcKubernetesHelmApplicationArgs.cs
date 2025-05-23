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
    /// Azure arc kubernetes helm application configurations.
    /// </summary>
    public sealed class AzureArcKubernetesHelmApplicationArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Azure arc kubernetes artifact profile.
        /// </summary>
        [Input("artifactProfile")]
        public Input<Inputs.AzureArcKubernetesArtifactProfileArgs>? ArtifactProfile { get; set; }

        /// <summary>
        /// The artifact type.
        /// Expected value is 'HelmPackage'.
        /// </summary>
        [Input("artifactType", required: true)]
        public Input<string> ArtifactType { get; set; } = null!;

        /// <summary>
        /// Depends on profile definition.
        /// </summary>
        [Input("dependsOnProfile")]
        public Input<Inputs.DependsOnProfileArgs>? DependsOnProfile { get; set; }

        /// <summary>
        /// Deploy mapping rule profile.
        /// </summary>
        [Input("deployParametersMappingRuleProfile")]
        public Input<Inputs.AzureArcKubernetesDeployMappingRuleProfileArgs>? DeployParametersMappingRuleProfile { get; set; }

        /// <summary>
        /// The name of the network function application.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        public AzureArcKubernetesHelmApplicationArgs()
        {
        }
        public static new AzureArcKubernetesHelmApplicationArgs Empty => new AzureArcKubernetesHelmApplicationArgs();
    }
}
