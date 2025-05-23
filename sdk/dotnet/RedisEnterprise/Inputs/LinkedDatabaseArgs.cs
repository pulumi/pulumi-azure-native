// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.RedisEnterprise.Inputs
{

    /// <summary>
    /// Specifies details of a linked database resource.
    /// </summary>
    public sealed class LinkedDatabaseArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Resource ID of a database resource to link with this database.
        /// </summary>
        [Input("id")]
        public Input<string>? Id { get; set; }

        public LinkedDatabaseArgs()
        {
        }
        public static new LinkedDatabaseArgs Empty => new LinkedDatabaseArgs();
    }
}
