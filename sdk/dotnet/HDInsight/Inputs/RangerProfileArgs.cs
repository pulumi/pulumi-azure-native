// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.HDInsight.Inputs
{

    /// <summary>
    /// The ranger cluster profile.
    /// </summary>
    public sealed class RangerProfileArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Specification for the Ranger Admin service.
        /// </summary>
        [Input("rangerAdmin", required: true)]
        public Input<Inputs.RangerAdminSpecArgs> RangerAdmin { get; set; } = null!;

        /// <summary>
        /// Properties required to describe audit log storage.
        /// </summary>
        [Input("rangerAudit")]
        public Input<Inputs.RangerAuditSpecArgs>? RangerAudit { get; set; }

        /// <summary>
        /// Specification for the Ranger Usersync service
        /// </summary>
        [Input("rangerUsersync", required: true)]
        public Input<Inputs.RangerUsersyncSpecArgs> RangerUsersync { get; set; } = null!;

        public RangerProfileArgs()
        {
        }
        public static new RangerProfileArgs Empty => new RangerProfileArgs();
    }
}
