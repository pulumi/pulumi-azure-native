// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.VideoAnalyzer.Inputs
{

    /// <summary>
    /// Required validation properties for tokens generated with RSA algorithm.
    /// </summary>
    public sealed class RsaTokenKeyArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// RSA algorithm to be used: RS256, RS384 or RS512.
        /// </summary>
        [Input("alg", required: true)]
        public InputUnion<string, Pulumi.AzureNative.VideoAnalyzer.AccessPolicyRsaAlgo> Alg { get; set; } = null!;

        /// <summary>
        /// RSA public key exponent.
        /// </summary>
        [Input("e", required: true)]
        public Input<string> E { get; set; } = null!;

        /// <summary>
        /// JWT token key id. Validation keys are looked up based on the key id present on the JWT token header.
        /// </summary>
        [Input("kid", required: true)]
        public Input<string> Kid { get; set; } = null!;

        /// <summary>
        /// RSA public key modulus.
        /// </summary>
        [Input("n", required: true)]
        public Input<string> N { get; set; } = null!;

        /// <summary>
        /// The discriminator for derived types.
        /// Expected value is '#Microsoft.VideoAnalyzer.RsaTokenKey'.
        /// </summary>
        [Input("type", required: true)]
        public Input<string> Type { get; set; } = null!;

        public RsaTokenKeyArgs()
        {
        }
        public static new RsaTokenKeyArgs Empty => new RsaTokenKeyArgs();
    }
}
