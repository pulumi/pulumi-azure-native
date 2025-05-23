// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.Monitor.Inputs
{

    /// <summary>
    /// Cache configurations.
    /// </summary>
    public sealed class CacheConfigurationArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Max storage usage in megabytes.
        /// </summary>
        [Input("maxStorageUsage")]
        public Input<int>? MaxStorageUsage { get; set; }

        /// <summary>
        /// Retention period in minutes.
        /// </summary>
        [Input("retentionPeriod")]
        public Input<int>? RetentionPeriod { get; set; }

        public CacheConfigurationArgs()
        {
        }
        public static new CacheConfigurationArgs Empty => new CacheConfigurationArgs();
    }
}
