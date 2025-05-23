// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.ServiceFabric.Inputs
{

    public sealed class ApplicationUserAssignedIdentityArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The friendly name of user assigned identity.
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        /// <summary>
        /// The principal id of user assigned identity.
        /// </summary>
        [Input("principalId", required: true)]
        public Input<string> PrincipalId { get; set; } = null!;

        public ApplicationUserAssignedIdentityArgs()
        {
        }
        public static new ApplicationUserAssignedIdentityArgs Empty => new ApplicationUserAssignedIdentityArgs();
    }
}
