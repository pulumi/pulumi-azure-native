// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.Workloads.Outputs
{

    /// <summary>
    /// Gets or sets the storage configuration.
    /// </summary>
    [OutputType]
    public sealed class StorageConfigurationResponse
    {
        /// <summary>
        /// The properties of the transport directory attached to the VIS. The default for transportFileShareConfiguration is the createAndMount flow if storage configuration is missing.
        /// </summary>
        public readonly object? TransportFileShareConfiguration;

        [OutputConstructor]
        private StorageConfigurationResponse(object? transportFileShareConfiguration)
        {
            TransportFileShareConfiguration = transportFileShareConfiguration;
        }
    }
}
