// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.MachineLearningServices.Inputs
{

    /// <summary>
    /// Empty/none datastore credentials.
    /// </summary>
    public sealed class NoneDatastoreCredentialsArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Enum to determine the datastore credentials type.
        /// Expected value is 'None'.
        /// </summary>
        [Input("credentialsType", required: true)]
        public Input<string> CredentialsType { get; set; } = null!;

        public NoneDatastoreCredentialsArgs()
        {
        }
        public static new NoneDatastoreCredentialsArgs Empty => new NoneDatastoreCredentialsArgs();
    }
}
