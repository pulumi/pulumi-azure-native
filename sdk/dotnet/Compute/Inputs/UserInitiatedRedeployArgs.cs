// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.Compute.Inputs
{

    /// <summary>
    /// Specifies Redeploy related Scheduled Event related configurations.
    /// </summary>
    public sealed class UserInitiatedRedeployArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Specifies Redeploy Scheduled Event related configurations.
        /// </summary>
        [Input("automaticallyApprove")]
        public Input<bool>? AutomaticallyApprove { get; set; }

        public UserInitiatedRedeployArgs()
        {
        }
        public static new UserInitiatedRedeployArgs Empty => new UserInitiatedRedeployArgs();
    }
}
