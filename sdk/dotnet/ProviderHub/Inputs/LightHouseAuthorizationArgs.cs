// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.ProviderHub.Inputs
{

    public sealed class LightHouseAuthorizationArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The principal id.
        /// </summary>
        [Input("principalId", required: true)]
        public Input<string> PrincipalId { get; set; } = null!;

        /// <summary>
        /// The role definition id.
        /// </summary>
        [Input("roleDefinitionId", required: true)]
        public Input<string> RoleDefinitionId { get; set; } = null!;

        public LightHouseAuthorizationArgs()
        {
        }
        public static new LightHouseAuthorizationArgs Empty => new LightHouseAuthorizationArgs();
    }
}
