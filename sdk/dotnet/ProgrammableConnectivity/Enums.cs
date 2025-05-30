// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.ComponentModel;
using Pulumi;

namespace Pulumi.AzureNative.ProgrammableConnectivity
{
    /// <summary>
    /// Type of the account the user has with the Operator's Network API infrastructure. AzureManaged | UserManaged.
    /// </summary>
    [EnumType]
    public readonly struct AccountType : IEquatable<AccountType>
    {
        private readonly string _value;

        private AccountType(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        /// <summary>
        /// Managed by Azure on-behalf-of the user.
        /// </summary>
        public static AccountType AzureManaged { get; } = new AccountType("AzureManaged");
        /// <summary>
        /// Managed by the User themselves on the Operator end.
        /// </summary>
        public static AccountType UserManaged { get; } = new AccountType("UserManaged");

        public static bool operator ==(AccountType left, AccountType right) => left.Equals(right);
        public static bool operator !=(AccountType left, AccountType right) => !left.Equals(right);

        public static explicit operator string(AccountType value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is AccountType other && Equals(other);
        public bool Equals(AccountType other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }
}
