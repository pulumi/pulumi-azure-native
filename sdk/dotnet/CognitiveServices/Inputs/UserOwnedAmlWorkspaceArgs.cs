// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.CognitiveServices.Inputs
{

    /// <summary>
    /// The user owned AML workspace for Cognitive Services account.
    /// </summary>
    public sealed class UserOwnedAmlWorkspaceArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Identity Client id of a AML workspace resource.
        /// </summary>
        [Input("identityClientId")]
        public Input<string>? IdentityClientId { get; set; }

        /// <summary>
        /// Full resource id of a AML workspace resource.
        /// </summary>
        [Input("resourceId")]
        public Input<string>? ResourceId { get; set; }

        public UserOwnedAmlWorkspaceArgs()
        {
        }
        public static new UserOwnedAmlWorkspaceArgs Empty => new UserOwnedAmlWorkspaceArgs();
    }
}
