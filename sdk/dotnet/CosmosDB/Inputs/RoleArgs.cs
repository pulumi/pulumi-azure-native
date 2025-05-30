// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.CosmosDB.Inputs
{

    /// <summary>
    /// The set of roles permitted through this Role Definition.
    /// </summary>
    public sealed class RoleArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The database name the role is applied.
        /// </summary>
        [Input("db")]
        public Input<string>? Db { get; set; }

        /// <summary>
        /// The role name.
        /// </summary>
        [Input("role")]
        public Input<string>? Role { get; set; }

        public RoleArgs()
        {
        }
        public static new RoleArgs Empty => new RoleArgs();
    }
}
