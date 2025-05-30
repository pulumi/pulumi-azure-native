// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.DataFactory.Outputs
{

    /// <summary>
    /// Azure Data Explorer command activity.
    /// </summary>
    [OutputType]
    public sealed class AzureDataExplorerCommandActivityResponse
    {
        /// <summary>
        /// A control command, according to the Azure Data Explorer command syntax. Type: string (or Expression with resultType string).
        /// </summary>
        public readonly object Command;
        /// <summary>
        /// Control command timeout. Type: string (or Expression with resultType string), pattern: ((\d+)\.)?(\d\d):(60|([0-5][0-9])):(60|([0-5][0-9]))..)
        /// </summary>
        public readonly object? CommandTimeout;
        /// <summary>
        /// Activity depends on condition.
        /// </summary>
        public readonly ImmutableArray<Outputs.ActivityDependencyResponse> DependsOn;
        /// <summary>
        /// Activity description.
        /// </summary>
        public readonly string? Description;
        /// <summary>
        /// Linked service reference.
        /// </summary>
        public readonly Outputs.LinkedServiceReferenceResponse? LinkedServiceName;
        /// <summary>
        /// Activity name.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// Status result of the activity when the state is set to Inactive. This is an optional property and if not provided when the activity is inactive, the status will be Succeeded by default.
        /// </summary>
        public readonly string? OnInactiveMarkAs;
        /// <summary>
        /// Activity policy.
        /// </summary>
        public readonly Outputs.ActivityPolicyResponse? Policy;
        /// <summary>
        /// Activity state. This is an optional property and if not provided, the state will be Active by default.
        /// </summary>
        public readonly string? State;
        /// <summary>
        /// Type of activity.
        /// Expected value is 'AzureDataExplorerCommand'.
        /// </summary>
        public readonly string Type;
        /// <summary>
        /// Activity user properties.
        /// </summary>
        public readonly ImmutableArray<Outputs.UserPropertyResponse> UserProperties;

        [OutputConstructor]
        private AzureDataExplorerCommandActivityResponse(
            object command,

            object? commandTimeout,

            ImmutableArray<Outputs.ActivityDependencyResponse> dependsOn,

            string? description,

            Outputs.LinkedServiceReferenceResponse? linkedServiceName,

            string name,

            string? onInactiveMarkAs,

            Outputs.ActivityPolicyResponse? policy,

            string? state,

            string type,

            ImmutableArray<Outputs.UserPropertyResponse> userProperties)
        {
            Command = command;
            CommandTimeout = commandTimeout;
            DependsOn = dependsOn;
            Description = description;
            LinkedServiceName = linkedServiceName;
            Name = name;
            OnInactiveMarkAs = onInactiveMarkAs;
            Policy = policy;
            State = state;
            Type = type;
            UserProperties = userProperties;
        }
    }
}
