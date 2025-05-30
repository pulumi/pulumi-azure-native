// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.DataBoxEdge.Inputs
{

    /// <summary>
    /// Kubernetes role resources
    /// </summary>
    public sealed class KubernetesRoleResourcesArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Kubernetes role compute resource
        /// </summary>
        [Input("compute", required: true)]
        public Input<Inputs.KubernetesRoleComputeArgs> Compute { get; set; } = null!;

        /// <summary>
        /// Kubernetes role storage resource
        /// </summary>
        [Input("storage")]
        public Input<Inputs.KubernetesRoleStorageArgs>? Storage { get; set; }

        public KubernetesRoleResourcesArgs()
        {
        }
        public static new KubernetesRoleResourcesArgs Empty => new KubernetesRoleResourcesArgs();
    }
}
