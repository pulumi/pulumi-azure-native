// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.SignalRService.Outputs
{

    /// <summary>
    /// Resource log category configuration of a Microsoft.SignalRService resource.
    /// </summary>
    [OutputType]
    public sealed class ResourceLogCategoryResponse
    {
        /// <summary>
        /// Indicates whether or the resource log category is enabled.
        /// Available values: true, false.
        /// Case insensitive.
        /// </summary>
        public readonly string? Enabled;
        /// <summary>
        /// Gets or sets the resource log category's name.
        /// Available values: ConnectivityLogs, MessagingLogs.
        /// Case insensitive.
        /// </summary>
        public readonly string? Name;

        [OutputConstructor]
        private ResourceLogCategoryResponse(
            string? enabled,

            string? name)
        {
            Enabled = enabled;
            Name = name;
        }
    }
}
