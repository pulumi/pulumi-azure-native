// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.DataReplication.Inputs
{

    /// <summary>
    /// Private endpoint connection details at member level.
    /// </summary>
    public sealed class ConnectionDetailsArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Gets or sets group id.
        /// </summary>
        [Input("groupId")]
        public Input<string>? GroupId { get; set; }

        /// <summary>
        /// Gets or sets id.
        /// </summary>
        [Input("id")]
        public Input<string>? Id { get; set; }

        /// <summary>
        /// Gets or sets link identifier.
        /// </summary>
        [Input("linkIdentifier")]
        public Input<string>? LinkIdentifier { get; set; }

        /// <summary>
        /// Gets or sets member name.
        /// </summary>
        [Input("memberName")]
        public Input<string>? MemberName { get; set; }

        /// <summary>
        /// Gets or sets private IP address.
        /// </summary>
        [Input("privateIpAddress")]
        public Input<string>? PrivateIpAddress { get; set; }

        public ConnectionDetailsArgs()
        {
        }
        public static new ConnectionDetailsArgs Empty => new ConnectionDetailsArgs();
    }
}
