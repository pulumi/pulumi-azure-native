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
    /// An arm role receiver.
    /// </summary>
    public sealed class ArmRoleReceiverArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The name of the arm role receiver. Names must be unique across all receivers within an action group.
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        /// <summary>
        /// The arm role id.
        /// </summary>
        [Input("roleId", required: true)]
        public Input<string> RoleId { get; set; } = null!;

        /// <summary>
        /// Indicates whether to use common alert schema.
        /// </summary>
        [Input("useCommonAlertSchema")]
        public Input<bool>? UseCommonAlertSchema { get; set; }

        public ArmRoleReceiverArgs()
        {
            UseCommonAlertSchema = false;
        }
        public static new ArmRoleReceiverArgs Empty => new ArmRoleReceiverArgs();
    }
}
