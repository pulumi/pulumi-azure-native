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
    /// Definition of awsLightsailBucket
    /// </summary>
    public sealed class AwsLightsailBucketPropertiesArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Indicates whether the bundle that is currently applied to a bucket can be changed to another bundle. You can update a bucket's bundle only one time within a monthly AWS billing cycle.
        /// </summary>
        [Input("ableToUpdateBundle")]
        public Input<bool>? AbleToUpdateBundle { get; set; }

        /// <summary>
        /// An object that sets the public accessibility of objects in the specified bucket.
        /// </summary>
        [Input("accessRules")]
        public Input<Inputs.AccessRulesArgs>? AccessRules { get; set; }

        /// <summary>
        /// Property bucketArn
        /// </summary>
        [Input("bucketArn")]
        public Input<string>? BucketArn { get; set; }

        /// <summary>
        /// The name for the bucket.
        /// </summary>
        [Input("bucketName")]
        public Input<string>? BucketName { get; set; }

        /// <summary>
        /// The ID of the bundle to use for the bucket.
        /// </summary>
        [Input("bundleId")]
        public Input<string>? BundleId { get; set; }

        /// <summary>
        /// Specifies whether to enable or disable versioning of objects in the bucket.
        /// </summary>
        [Input("objectVersioning")]
        public Input<bool>? ObjectVersioning { get; set; }

        [Input("readOnlyAccessAccounts")]
        private InputList<string>? _readOnlyAccessAccounts;

        /// <summary>
        /// An array of strings to specify the AWS account IDs that can access the bucket.
        /// </summary>
        public InputList<string> ReadOnlyAccessAccounts
        {
            get => _readOnlyAccessAccounts ?? (_readOnlyAccessAccounts = new InputList<string>());
            set => _readOnlyAccessAccounts = value;
        }

        [Input("resourcesReceivingAccess")]
        private InputList<string>? _resourcesReceivingAccess;

        /// <summary>
        /// The names of the Lightsail resources for which to set bucket access.
        /// </summary>
        public InputList<string> ResourcesReceivingAccess
        {
            get => _resourcesReceivingAccess ?? (_resourcesReceivingAccess = new InputList<string>());
            set => _resourcesReceivingAccess = value;
        }

        [Input("tags")]
        private InputList<Inputs.TagArgs>? _tags;

        /// <summary>
        /// An array of key-value pairs to apply to this resource.
        /// </summary>
        public InputList<Inputs.TagArgs> Tags
        {
            get => _tags ?? (_tags = new InputList<Inputs.TagArgs>());
            set => _tags = value;
        }

        /// <summary>
        /// The URL of the bucket.
        /// </summary>
        [Input("url")]
        public Input<string>? Url { get; set; }

        public AwsLightsailBucketPropertiesArgs()
        {
        }
        public static new AwsLightsailBucketPropertiesArgs Empty => new AwsLightsailBucketPropertiesArgs();
    }
}
