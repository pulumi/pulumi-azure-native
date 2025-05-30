// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.IntegrationSpaces.Outputs
{

    /// <summary>
    /// The tracking event definition.
    /// </summary>
    [OutputType]
    public sealed class TrackingEventDefinitionResponse
    {
        /// <summary>
        /// The operation name.
        /// </summary>
        public readonly string? OperationName;
        /// <summary>
        /// The operation type.
        /// </summary>
        public readonly string? OperationType;
        /// <summary>
        /// The properties to be collected for event.
        /// </summary>
        public readonly ImmutableDictionary<string, object>? Properties;

        [OutputConstructor]
        private TrackingEventDefinitionResponse(
            string? operationName,

            string? operationType,

            ImmutableDictionary<string, object>? properties)
        {
            OperationName = operationName;
            OperationType = operationType;
            Properties = properties;
        }
    }
}
