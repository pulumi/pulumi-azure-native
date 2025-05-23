// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.ContainerService.Outputs
{

    /// <summary>
    /// The node image upgrade to be applied to the target nodes in update run.
    /// </summary>
    [OutputType]
    public sealed class NodeImageSelectionResponse
    {
        /// <summary>
        /// Custom node image versions to upgrade the nodes to. This field is required if node image selection type is Custom. Otherwise, it must be empty. For each node image family (e.g., 'AKSUbuntu-1804gen2containerd'), this field can contain at most one version (e.g., only one of 'AKSUbuntu-1804gen2containerd-2023.01.12' or 'AKSUbuntu-1804gen2containerd-2023.02.12', not both). If the nodes belong to a family without a matching image version in this field, they are not upgraded.
        /// </summary>
        public readonly ImmutableArray<Outputs.NodeImageVersionResponse> CustomNodeImageVersions;
        /// <summary>
        /// The node image upgrade type.
        /// </summary>
        public readonly string Type;

        [OutputConstructor]
        private NodeImageSelectionResponse(
            ImmutableArray<Outputs.NodeImageVersionResponse> customNodeImageVersions,

            string type)
        {
            CustomNodeImageVersions = customNodeImageVersions;
            Type = type;
        }
    }
}
