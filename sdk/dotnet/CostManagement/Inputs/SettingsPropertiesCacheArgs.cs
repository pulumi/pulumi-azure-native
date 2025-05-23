// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.CostManagement.Inputs
{

    public sealed class SettingsPropertiesCacheArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Indicates the account type. Allowed values include: EA, PAYG, Modern, Internal, Unknown.
        /// </summary>
        [Input("channel", required: true)]
        public Input<string> Channel { get; set; } = null!;

        /// <summary>
        /// Resource ID used by Resource Manager to uniquely identify the scope.
        /// </summary>
        [Input("id", required: true)]
        public Input<string> Id { get; set; } = null!;

        /// <summary>
        /// Display name for the scope.
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        /// <summary>
        /// Resource ID of the parent scope. For instance, subscription's resource ID for a resource group or a management group resource ID for a subscription.
        /// </summary>
        [Input("parent")]
        public Input<string>? Parent { get; set; }

        /// <summary>
        /// Indicates the status of the scope. Status only applies to subscriptions and billing accounts.
        /// </summary>
        [Input("status")]
        public Input<string>? Status { get; set; }

        /// <summary>
        /// Indicates the type of modern account. Allowed values include: Individual, Enterprise, Partner, Indirect, NotApplicable
        /// </summary>
        [Input("subchannel", required: true)]
        public Input<string> Subchannel { get; set; } = null!;

        public SettingsPropertiesCacheArgs()
        {
        }
        public static new SettingsPropertiesCacheArgs Empty => new SettingsPropertiesCacheArgs();
    }
}
