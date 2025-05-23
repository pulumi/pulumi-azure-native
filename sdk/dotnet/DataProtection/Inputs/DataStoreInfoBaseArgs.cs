// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.DataProtection.Inputs
{

    /// <summary>
    /// DataStoreInfo base
    /// </summary>
    public sealed class DataStoreInfoBaseArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// type of datastore; Operational/Vault/Archive
        /// </summary>
        [Input("dataStoreType", required: true)]
        public InputUnion<string, Pulumi.AzureNative.DataProtection.DataStoreTypes> DataStoreType { get; set; } = null!;

        /// <summary>
        /// Type of Datasource object, used to initialize the right inherited type
        /// </summary>
        [Input("objectType", required: true)]
        public Input<string> ObjectType { get; set; } = null!;

        public DataStoreInfoBaseArgs()
        {
        }
        public static new DataStoreInfoBaseArgs Empty => new DataStoreInfoBaseArgs();
    }
}
