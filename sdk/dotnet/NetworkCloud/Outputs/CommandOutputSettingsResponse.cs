// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.NetworkCloud.Outputs
{

    [OutputType]
    public sealed class CommandOutputSettingsResponse
    {
        /// <summary>
        /// The selection of the managed identity to use with this storage account container. The identity type must be either system assigned or user assigned.
        /// </summary>
        public readonly Outputs.IdentitySelectorResponse? AssociatedIdentity;
        /// <summary>
        /// The URL of the storage account container that is to be used by the specified identities.
        /// </summary>
        public readonly string? ContainerUrl;

        [OutputConstructor]
        private CommandOutputSettingsResponse(
            Outputs.IdentitySelectorResponse? associatedIdentity,

            string? containerUrl)
        {
            AssociatedIdentity = associatedIdentity;
            ContainerUrl = containerUrl;
        }
    }
}
