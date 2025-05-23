// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.CustomerInsights.Inputs
{

    /// <summary>
    /// The participant property reference.
    /// </summary>
    public sealed class ParticipantPropertyReferenceArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The source property that maps to the target property.
        /// </summary>
        [Input("sourcePropertyName", required: true)]
        public Input<string> SourcePropertyName { get; set; } = null!;

        /// <summary>
        /// The target property that maps to the source property.
        /// </summary>
        [Input("targetPropertyName", required: true)]
        public Input<string> TargetPropertyName { get; set; } = null!;

        public ParticipantPropertyReferenceArgs()
        {
        }
        public static new ParticipantPropertyReferenceArgs Empty => new ParticipantPropertyReferenceArgs();
    }
}
