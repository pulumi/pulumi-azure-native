// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.Network.Outputs
{

    /// <summary>
    /// Reference to container resource in remote resource provider.
    /// </summary>
    [OutputType]
    public sealed class ContainerResponse
    {
        /// <summary>
        /// Resource ID.
        /// </summary>
        public readonly string? Id;

        [OutputConstructor]
        private ContainerResponse(string? id)
        {
            Id = id;
        }
    }
}
