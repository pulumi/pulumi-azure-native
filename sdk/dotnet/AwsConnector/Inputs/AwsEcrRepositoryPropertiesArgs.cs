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
    /// Definition of awsEcrRepository
    /// </summary>
    public sealed class AwsEcrRepositoryPropertiesArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Property arn
        /// </summary>
        [Input("arn")]
        public Input<string>? Arn { get; set; }

        /// <summary>
        /// If true, deleting the repository force deletes the contents of the repository. If false, the repository must be empty before attempting to delete it. If true, deleting the repository force deletes the contents of the repository. Without a force delete, you can only delete empty repositories.
        /// </summary>
        [Input("emptyOnDelete")]
        public Input<bool>? EmptyOnDelete { get; set; }

        /// <summary>
        /// The encryption configuration for the repository. This determines how the contents of your repository are encrypted at rest. The encryption configuration for the repository. This determines how the contents of your repository are encrypted at rest. By default, when no encryption configuration is set or the ``AES256`` encryption type is used, Amazon ECR uses server-side encryption with Amazon S3-managed encryption keys which encrypts your data at rest using an AES-256 encryption algorithm. This does not require any action on your part. For more control over the encryption of the contents of your repository, you can use server-side encryption with KMSlong key stored in KMSlong (KMS) to encrypt your images. For more information, see [Amazon ECR encryption at rest](https://docs.aws.amazon.com/AmazonECR/latest/userguide/encryption-at-rest.html) in the *Amazon Elastic Container Registry User Guide*.
        /// </summary>
        [Input("encryptionConfiguration")]
        public Input<Inputs.EncryptionConfigurationArgs>? EncryptionConfiguration { get; set; }

        /// <summary>
        /// The image scanning configuration for the repository. This determines whether images are scanned for known vulnerabilities after being pushed to the repository. The image scanning configuration for a repository.
        /// </summary>
        [Input("imageScanningConfiguration")]
        public Input<Inputs.ImageScanningConfigurationArgs>? ImageScanningConfiguration { get; set; }

        /// <summary>
        /// The tag mutability setting for the repository. If this parameter is omitted, the default setting of ``MUTABLE`` will be used which will allow image tags to be overwritten. If ``IMMUTABLE`` is specified, all image tags within the repository will be immutable which will prevent them from being overwritten.
        /// </summary>
        [Input("imageTagMutability")]
        public InputUnion<string, Pulumi.AzureNative.AwsConnector.ImageTagMutability>? ImageTagMutability { get; set; }

        /// <summary>
        /// Creates or updates a lifecycle policy. For information about lifecycle policy syntax, see [Lifecycle policy template](https://docs.aws.amazon.com/AmazonECR/latest/userguide/LifecyclePolicies.html). The ``LifecyclePolicy`` property type specifies a lifecycle policy. For information about lifecycle policy syntax, see [Lifecycle policy template](https://docs.aws.amazon.com/AmazonECR/latest/userguide/LifecyclePolicies.html) in the *Amazon ECR User Guide*.
        /// </summary>
        [Input("lifecyclePolicy")]
        public Input<Inputs.LifecyclePolicyArgs>? LifecyclePolicy { get; set; }

        /// <summary>
        /// The name to use for the repository. The repository name may be specified on its own (such as ``nginx-web-app``) or it can be prepended with a namespace to group the repository into a category (such as ``project-a/nginx-web-app``). If you don't specify a name, CFNlong generates a unique physical ID and uses that ID for the repository name. For more information, see [Name type](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-name.html). The repository name must start with a letter and can only contain lowercase letters, numbers, hyphens, underscores, and forward slashes.  If you specify a name, you cannot perform updates that require replacement of this resource. You can perform updates that require no or some interruption. If you must replace the resource, specify a new name.
        /// </summary>
        [Input("repositoryName")]
        public Input<string>? RepositoryName { get; set; }

        /// <summary>
        /// The JSON repository policy text to apply to the repository. For more information, see [Amazon ECR repository policies](https://docs.aws.amazon.com/AmazonECR/latest/userguide/repository-policy-examples.html) in the *Amazon Elastic Container Registry User Guide*.
        /// </summary>
        [Input("repositoryPolicyText")]
        public Input<object>? RepositoryPolicyText { get; set; }

        /// <summary>
        /// Property repositoryUri
        /// </summary>
        [Input("repositoryUri")]
        public Input<string>? RepositoryUri { get; set; }

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

        public AwsEcrRepositoryPropertiesArgs()
        {
        }
        public static new AwsEcrRepositoryPropertiesArgs Empty => new AwsEcrRepositoryPropertiesArgs();
    }
}
