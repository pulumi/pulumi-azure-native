// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.AwsConnector.Outputs
{

    /// <summary>
    /// Definition of awsIamInstanceProfile
    /// </summary>
    [OutputType]
    public sealed class AwsIamInstanceProfilePropertiesResponse
    {
        /// <summary>
        /// Property arn
        /// </summary>
        public readonly string? Arn;
        /// <summary>
        /// The name of the instance profile to create. This parameter allows (through its [regex pattern](https://docs.aws.amazon.com/http://wikipedia.org/wiki/regex)) a string of characters consisting of upper and lowercase alphanumeric characters with no spaces. You can also include any of the following characters: _+=,.@-
        /// </summary>
        public readonly string? InstanceProfileName;
        /// <summary>
        /// The path to the instance profile. For more information about paths, see [IAM Identifiers](https://docs.aws.amazon.com/IAM/latest/UserGuide/Using_Identifiers.html) in the *IAM User Guide*. This parameter is optional. If it is not included, it defaults to a slash (/). This parameter allows (through its [regex pattern](https://docs.aws.amazon.com/http://wikipedia.org/wiki/regex)) a string of characters consisting of either a forward slash (/) by itself or a string that must begin and end with forward slashes. In addition, it can contain any ASCII character from the ! (``\u0021``) through the DEL character (``\u007F``), including most punctuation characters, digits, and upper and lowercased letters.
        /// </summary>
        public readonly string? Path;
        /// <summary>
        /// The name of the role to associate with the instance profile. Only one role can be assigned to an EC2 instance at a time, and all applications on the instance share the same role and permissions.
        /// </summary>
        public readonly ImmutableArray<string> Roles;

        [OutputConstructor]
        private AwsIamInstanceProfilePropertiesResponse(
            string? arn,

            string? instanceProfileName,

            string? path,

            ImmutableArray<string> roles)
        {
            Arn = arn;
            InstanceProfileName = instanceProfileName;
            Path = path;
            Roles = roles;
        }
    }
}
