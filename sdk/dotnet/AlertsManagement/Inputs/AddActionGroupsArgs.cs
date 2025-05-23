// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.AlertsManagement.Inputs
{

    /// <summary>
    /// Add action groups to alert processing rule.
    /// </summary>
    public sealed class AddActionGroupsArgs : global::Pulumi.ResourceArgs
    {
        [Input("actionGroupIds", required: true)]
        private InputList<string>? _actionGroupIds;

        /// <summary>
        /// List of action group Ids to add to alert processing rule.
        /// </summary>
        public InputList<string> ActionGroupIds
        {
            get => _actionGroupIds ?? (_actionGroupIds = new InputList<string>());
            set => _actionGroupIds = value;
        }

        /// <summary>
        /// Action that should be applied.
        /// Expected value is 'AddActionGroups'.
        /// </summary>
        [Input("actionType", required: true)]
        public Input<string> ActionType { get; set; } = null!;

        public AddActionGroupsArgs()
        {
        }
        public static new AddActionGroupsArgs Empty => new AddActionGroupsArgs();
    }
}
