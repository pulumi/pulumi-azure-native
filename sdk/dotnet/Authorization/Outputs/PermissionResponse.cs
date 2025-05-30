// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.Authorization.Outputs
{

    /// <summary>
    /// Role definition permissions.
    /// </summary>
    [OutputType]
    public sealed class PermissionResponse
    {
        /// <summary>
        /// Allowed actions.
        /// </summary>
        public readonly ImmutableArray<string> Actions;
        /// <summary>
        /// The conditions on the role definition. This limits the resources it can be assigned to. e.g.: @Resource[Microsoft.Storage/storageAccounts/blobServices/containers:ContainerName] StringEqualsIgnoreCase 'foo_storage_container'
        /// </summary>
        public readonly string Condition;
        /// <summary>
        /// Version of the condition. Currently the only accepted value is '2.0'
        /// </summary>
        public readonly string ConditionVersion;
        /// <summary>
        /// Allowed Data actions.
        /// </summary>
        public readonly ImmutableArray<string> DataActions;
        /// <summary>
        /// Denied actions.
        /// </summary>
        public readonly ImmutableArray<string> NotActions;
        /// <summary>
        /// Denied Data actions.
        /// </summary>
        public readonly ImmutableArray<string> NotDataActions;

        [OutputConstructor]
        private PermissionResponse(
            ImmutableArray<string> actions,

            string condition,

            string conditionVersion,

            ImmutableArray<string> dataActions,

            ImmutableArray<string> notActions,

            ImmutableArray<string> notDataActions)
        {
            Actions = actions;
            Condition = condition;
            ConditionVersion = conditionVersion;
            DataActions = dataActions;
            NotActions = notActions;
            NotDataActions = notDataActions;
        }
    }
}
