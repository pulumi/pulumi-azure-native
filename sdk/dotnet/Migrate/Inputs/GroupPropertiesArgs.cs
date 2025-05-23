// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.Migrate.Inputs
{

    /// <summary>
    /// Properties of group resource.
    /// </summary>
    public sealed class GroupPropertiesArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The type of group.
        /// </summary>
        [Input("groupType")]
        public Input<string>? GroupType { get; set; }

        public GroupPropertiesArgs()
        {
        }
        public static new GroupPropertiesArgs Empty => new GroupPropertiesArgs();
    }
}
