// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.SecurityInsights.Inputs
{

    /// <summary>
    /// Instruction step details, to be displayed in the Instructions steps section in the connector's page in Sentinel Portal.
    /// </summary>
    public sealed class InstructionStepDetailsArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Gets or sets the instruction type parameters settings.
        /// </summary>
        [Input("parameters", required: true)]
        public Input<object> Parameters { get; set; } = null!;

        /// <summary>
        /// Gets or sets the instruction type name.
        /// </summary>
        [Input("type", required: true)]
        public Input<string> Type { get; set; } = null!;

        public InstructionStepDetailsArgs()
        {
        }
        public static new InstructionStepDetailsArgs Empty => new InstructionStepDetailsArgs();
    }
}
