// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.Storage.Inputs
{

    public sealed class SshPublicKeyArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Optional. It is used to store the function/usage of the key
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// Ssh public key base64 encoded. The format should be: '&lt;keyType&gt; &lt;keyData&gt;', e.g. ssh-rsa AAAABBBB
        /// </summary>
        [Input("key")]
        public Input<string>? Key { get; set; }

        public SshPublicKeyArgs()
        {
        }
        public static new SshPublicKeyArgs Empty => new SshPublicKeyArgs();
    }
}
