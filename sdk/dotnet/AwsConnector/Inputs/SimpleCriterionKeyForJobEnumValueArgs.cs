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
    /// Definition of SimpleCriterionKeyForJobEnumValue
    /// </summary>
    public sealed class SimpleCriterionKeyForJobEnumValueArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Property value
        /// </summary>
        [Input("value")]
        public InputUnion<string, Pulumi.AzureNative.AwsConnector.SimpleCriterionKeyForJob>? Value { get; set; }

        public SimpleCriterionKeyForJobEnumValueArgs()
        {
        }
        public static new SimpleCriterionKeyForJobEnumValueArgs Empty => new SimpleCriterionKeyForJobEnumValueArgs();
    }
}
