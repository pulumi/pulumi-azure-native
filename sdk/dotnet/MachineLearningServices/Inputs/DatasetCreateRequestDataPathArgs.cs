// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.MachineLearningServices.Inputs
{

    public sealed class DatasetCreateRequestDataPathArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The datastore name.
        /// </summary>
        [Input("datastoreName")]
        public Input<string>? DatastoreName { get; set; }

        /// <summary>
        /// Path within the datastore.
        /// </summary>
        [Input("relativePath")]
        public Input<string>? RelativePath { get; set; }

        public DatasetCreateRequestDataPathArgs()
        {
        }
        public static new DatasetCreateRequestDataPathArgs Empty => new DatasetCreateRequestDataPathArgs();
    }
}
