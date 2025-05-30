// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.AlertsManagement.Outputs
{

    /// <summary>
    /// Indicates if all action groups should be removed.
    /// </summary>
    [OutputType]
    public sealed class RemoveAllActionGroupsResponse
    {
        /// <summary>
        /// Action that should be applied.
        /// Expected value is 'RemoveAllActionGroups'.
        /// </summary>
        public readonly string ActionType;

        [OutputConstructor]
        private RemoveAllActionGroupsResponse(string actionType)
        {
            ActionType = actionType;
        }
    }
}
