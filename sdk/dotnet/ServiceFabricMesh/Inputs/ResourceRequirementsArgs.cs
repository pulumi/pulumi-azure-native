// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.ServiceFabricMesh.Inputs
{

    /// <summary>
    /// This type describes the resource requirements for a container or a service.
    /// </summary>
    public sealed class ResourceRequirementsArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Describes the maximum limits on the resources for a given container.
        /// </summary>
        [Input("limits")]
        public Input<Inputs.ResourceLimitsArgs>? Limits { get; set; }

        /// <summary>
        /// Describes the requested resources for a given container.
        /// </summary>
        [Input("requests", required: true)]
        public Input<Inputs.ResourceRequestsArgs> Requests { get; set; } = null!;

        public ResourceRequirementsArgs()
        {
        }
        public static new ResourceRequirementsArgs Empty => new ResourceRequirementsArgs();
    }
}
