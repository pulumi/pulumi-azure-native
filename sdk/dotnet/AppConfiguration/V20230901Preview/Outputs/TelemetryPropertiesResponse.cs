// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.AppConfiguration.V20230901Preview.Outputs
{

    /// <summary>
    /// Telemetry settings
    /// </summary>
    [OutputType]
    public sealed class TelemetryPropertiesResponse
    {
        /// <summary>
        /// Resource ID of a resource enabling telemetry collection
        /// </summary>
        public readonly string? ResourceId;

        [OutputConstructor]
        private TelemetryPropertiesResponse(string? resourceId)
        {
            ResourceId = resourceId;
        }
    }
}
