// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.AzureStackHCI.Inputs
{

    /// <summary>
    /// The Storage config of AzureStackHCI Cluster.
    /// </summary>
    public sealed class StorageArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// By default, this mode is set to Express and your storage is configured as per best practices based on the number of nodes in the cluster. Allowed values are 'Express','InfraOnly', 'KeepStorage'
        /// </summary>
        [Input("configurationMode")]
        public Input<string>? ConfigurationMode { get; set; }

        public StorageArgs()
        {
            ConfigurationMode = "Express";
        }
        public static new StorageArgs Empty => new StorageArgs();
    }
}
