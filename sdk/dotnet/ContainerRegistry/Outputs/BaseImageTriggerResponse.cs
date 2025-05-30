// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.ContainerRegistry.Outputs
{

    /// <summary>
    /// The trigger based on base image dependency.
    /// </summary>
    [OutputType]
    public sealed class BaseImageTriggerResponse
    {
        /// <summary>
        /// The type of the auto trigger for base image dependency updates.
        /// </summary>
        public readonly string BaseImageTriggerType;
        /// <summary>
        /// The name of the trigger.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// The current status of trigger.
        /// </summary>
        public readonly string? Status;
        /// <summary>
        /// The endpoint URL for receiving update triggers.
        /// </summary>
        public readonly string? UpdateTriggerEndpoint;
        /// <summary>
        /// Type of Payload body for Base image update triggers.
        /// </summary>
        public readonly string? UpdateTriggerPayloadType;

        [OutputConstructor]
        private BaseImageTriggerResponse(
            string baseImageTriggerType,

            string name,

            string? status,

            string? updateTriggerEndpoint,

            string? updateTriggerPayloadType)
        {
            BaseImageTriggerType = baseImageTriggerType;
            Name = name;
            Status = status;
            UpdateTriggerEndpoint = updateTriggerEndpoint;
            UpdateTriggerPayloadType = updateTriggerPayloadType;
        }
    }
}
