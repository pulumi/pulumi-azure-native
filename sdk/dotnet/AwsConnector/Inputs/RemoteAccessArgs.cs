// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.AwsConnector.Inputs
{

    /// <summary>
    /// Definition of RemoteAccess
    /// </summary>
    public sealed class RemoteAccessArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Property ec2SshKey
        /// </summary>
        [Input("ec2SshKey")]
        public Input<string>? Ec2SshKey { get; set; }

        [Input("sourceSecurityGroups")]
        private InputList<string>? _sourceSecurityGroups;

        /// <summary>
        /// Property sourceSecurityGroups
        /// </summary>
        public InputList<string> SourceSecurityGroups
        {
            get => _sourceSecurityGroups ?? (_sourceSecurityGroups = new InputList<string>());
            set => _sourceSecurityGroups = value;
        }

        public RemoteAccessArgs()
        {
        }
        public static new RemoteAccessArgs Empty => new RemoteAccessArgs();
    }
}
