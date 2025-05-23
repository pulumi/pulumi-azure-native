// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.HDInsight.Outputs
{

    /// <summary>
    /// Contains the IpTag associated with the public IP address
    /// </summary>
    [OutputType]
    public sealed class IpTagResponse
    {
        /// <summary>
        /// Gets or sets the ipTag type: Example FirstPartyUsage.
        /// </summary>
        public readonly string IpTagType;
        /// <summary>
        /// Gets or sets value of the IpTag associated with the public IP. Example HDInsight, SQL, Storage etc
        /// </summary>
        public readonly string Tag;

        [OutputConstructor]
        private IpTagResponse(
            string ipTagType,

            string tag)
        {
            IpTagType = ipTagType;
            Tag = tag;
        }
    }
}
