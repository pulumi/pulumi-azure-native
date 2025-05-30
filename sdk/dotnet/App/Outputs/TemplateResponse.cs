// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.App.Outputs
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
        /// List of specialized containers that run before app containers.
        /// </summary>
        public readonly ImmutableArray<Outputs.InitContainerResponse> InitContainers;
        /// <summary>
        /// User friendly suffix that is appended to the revision name
        /// </summary>
        public readonly string? RevisionSuffix;
        /// <summary>
        /// Scaling properties for the Container App.
        /// </summary>
        public readonly Outputs.ScaleResponse? Scale;
        /// <summary>
        /// List of container app services bound to the app
        /// </summary>
        public readonly ImmutableArray<Outputs.ServiceBindResponse> ServiceBinds;
        /// <summary>
        /// Optional duration in seconds the Container App Instance needs to terminate gracefully. Value must be non-negative integer. The value zero indicates stop immediately via the kill signal (no opportunity to shut down). If this value is nil, the default grace period will be used instead. Set this value longer than the expected cleanup time for your process. Defaults to 30 seconds.
        /// </summary>
        public readonly double? TerminationGracePeriodSeconds;
        /// <summary>
        /// List of volume definitions for the Container App.
        /// </summary>
        public readonly ImmutableArray<Outputs.VolumeResponse> Volumes;

        [OutputConstructor]
        private TemplateResponse(
            ImmutableArray<Outputs.ContainerResponse> containers,

            ImmutableArray<Outputs.InitContainerResponse> initContainers,

            string? revisionSuffix,

            Outputs.ScaleResponse? scale,

            ImmutableArray<Outputs.ServiceBindResponse> serviceBinds,

            double? terminationGracePeriodSeconds,

            ImmutableArray<Outputs.VolumeResponse> volumes)
        {
            Containers = containers;
            InitContainers = initContainers;
            RevisionSuffix = revisionSuffix;
            Scale = scale;
            ServiceBinds = serviceBinds;
            TerminationGracePeriodSeconds = terminationGracePeriodSeconds;
            Volumes = volumes;
        }
    }
}
