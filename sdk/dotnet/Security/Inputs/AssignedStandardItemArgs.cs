// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.Security.Inputs
{

    /// <summary>
    /// Describe the properties of a of a standard assignments object reference
    /// </summary>
    public sealed class AssignedStandardItemArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Full resourceId of the Microsoft.Security/standard object
        /// </summary>
        [Input("id")]
        public Input<string>? Id { get; set; }

        public AssignedStandardItemArgs()
        {
        }
        public static new AssignedStandardItemArgs Empty => new AssignedStandardItemArgs();
    }
}
