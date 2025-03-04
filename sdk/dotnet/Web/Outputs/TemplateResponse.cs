// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.Web.Outputs
{

    /// <summary>
    /// Container App versioned application definition.
    /// Defines the desired state of an immutable revision.
    /// Any changes to this section Will result in a new revision being created
    /// </summary>
    [OutputType]
    public sealed class TemplateResponse
    {
        /// <summary>
        /// List of container definitions for the Container App.
        /// </summary>
        public readonly ImmutableArray<Outputs.ContainerResponse> Containers;
        /// <summary>
        /// Dapr configuration for the Container App.
        /// </summary>
        public readonly Outputs.DaprResponse? Dapr;
        /// <summary>
        /// User friendly suffix that is appended to the revision name
        /// </summary>
        public readonly string? RevisionSuffix;
        /// <summary>
        /// Scaling properties for the Container App.
        /// </summary>
        public readonly Outputs.ScaleResponse? Scale;

        [OutputConstructor]
        private TemplateResponse(
            ImmutableArray<Outputs.ContainerResponse> containers,

            Outputs.DaprResponse? dapr,

            string? revisionSuffix,

            Outputs.ScaleResponse? scale)
        {
            Containers = containers;
            Dapr = dapr;
            RevisionSuffix = revisionSuffix;
            Scale = scale;
        }
    }
}
