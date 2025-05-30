// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.Network.Inputs
{

    /// <summary>
    /// Details the service to which the subnet is delegated.
    /// </summary>
    public sealed class DelegationArgs : global::Pulumi.ResourceArgs
    {
        [Input("actions")]
        private InputList<string>? _actions;

        /// <summary>
        /// Describes the actions permitted to the service upon delegation
        /// </summary>
        public InputList<string> Actions
        {
            get => _actions ?? (_actions = new InputList<string>());
            set => _actions = value;
        }

        /// <summary>
        /// A unique read-only string that changes whenever the resource is updated.
        /// </summary>
        [Input("etag")]
        public Input<string>? Etag { get; set; }

        /// <summary>
        /// Resource ID.
        /// </summary>
        [Input("id")]
        public Input<string>? Id { get; set; }

        /// <summary>
        /// The name of the resource that is unique within a subnet. This name can be used to access the resource.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The name of the service to whom the subnet should be delegated (e.g. Microsoft.Sql/servers).
        /// </summary>
        [Input("serviceName")]
        public Input<string>? ServiceName { get; set; }

        /// <summary>
        /// Resource type.
        /// </summary>
        [Input("type")]
        public Input<string>? Type { get; set; }

        public DelegationArgs()
        {
        }
        public static new DelegationArgs Empty => new DelegationArgs();
    }
}
