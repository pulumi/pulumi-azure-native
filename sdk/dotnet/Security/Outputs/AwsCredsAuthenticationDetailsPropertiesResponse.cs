// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.Security.Outputs
{

    /// <summary>
    /// AWS cloud account connector based credentials, the credentials is composed of access key ID and secret key, for more details, refer to &lt;a href="https://docs.aws.amazon.com/IAM/latest/UserGuide/id_users_create.html"&gt;Creating an IAM User in Your AWS Account (write only)&lt;/a&gt;
    /// </summary>
    [OutputType]
    public sealed class AwsCredsAuthenticationDetailsPropertiesResponse
    {
        /// <summary>
        /// The ID of the cloud account
        /// </summary>
        public readonly string AccountId;
        /// <summary>
        /// State of the multi-cloud connector
        /// </summary>
        public readonly string AuthenticationProvisioningState;
        /// <summary>
        /// Connect to your cloud account, for AWS use either account credentials or role-based authentication. For GCP use account organization credentials.
        /// Expected value is 'awsCreds'.
        /// </summary>
        public readonly string AuthenticationType;
        /// <summary>
        /// Public key element of the AWS credential object (write only)
        /// </summary>
        public readonly string AwsAccessKeyId;
        /// <summary>
        /// Secret key element of the AWS credential object (write only)
        /// </summary>
        public readonly string AwsSecretAccessKey;
        /// <summary>
        /// The permissions detected in the cloud account.
        /// </summary>
        public readonly ImmutableArray<string> GrantedPermissions;

        [OutputConstructor]
        private AwsCredsAuthenticationDetailsPropertiesResponse(
            string accountId,

            string authenticationProvisioningState,

            string authenticationType,

            string awsAccessKeyId,

            string awsSecretAccessKey,

            ImmutableArray<string> grantedPermissions)
        {
            AccountId = accountId;
            AuthenticationProvisioningState = authenticationProvisioningState;
            AuthenticationType = authenticationType;
            AwsAccessKeyId = awsAccessKeyId;
            AwsSecretAccessKey = awsSecretAccessKey;
            GrantedPermissions = grantedPermissions;
        }
    }
}
