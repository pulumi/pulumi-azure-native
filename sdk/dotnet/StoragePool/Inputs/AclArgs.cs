// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.StoragePool.Inputs
{

    /// <summary>
    /// Access Control List (ACL) for an iSCSI Target; defines LUN masking policy
    /// </summary>
    public sealed class AclArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// iSCSI initiator IQN (iSCSI Qualified Name); example: "iqn.2005-03.org.iscsi:client".
        /// </summary>
        [Input("initiatorIqn", required: true)]
        public Input<string> InitiatorIqn { get; set; } = null!;

        [Input("mappedLuns", required: true)]
        private InputList<string>? _mappedLuns;

        /// <summary>
        /// List of LUN names mapped to the ACL.
        /// </summary>
        public InputList<string> MappedLuns
        {
            get => _mappedLuns ?? (_mappedLuns = new InputList<string>());
            set => _mappedLuns = value;
        }

        public AclArgs()
        {
        }
        public static new AclArgs Empty => new AclArgs();
    }
}
