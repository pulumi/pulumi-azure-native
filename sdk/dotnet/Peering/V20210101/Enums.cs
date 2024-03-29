// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.ComponentModel;
using Pulumi;

namespace Pulumi.AzureNative.Peering.V20210101
{
    /// <summary>
    /// The role of the contact.
    /// </summary>
    [EnumType]
    public readonly struct Role : IEquatable<Role>
    {
        private readonly string _value;

        private Role(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static Role Noc { get; } = new Role("Noc");
        public static Role Policy { get; } = new Role("Policy");
        public static Role Technical { get; } = new Role("Technical");
        public static Role Service { get; } = new Role("Service");
        public static Role Escalation { get; } = new Role("Escalation");
        public static Role Other { get; } = new Role("Other");

        public static bool operator ==(Role left, Role right) => left.Equals(right);
        public static bool operator !=(Role left, Role right) => !left.Equals(right);

        public static explicit operator string(Role value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is Role other && Equals(other);
        public bool Equals(Role other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    /// <summary>
    /// The validation state of the ASN associated with the peer.
    /// </summary>
    [EnumType]
    public readonly struct ValidationState : IEquatable<ValidationState>
    {
        private readonly string _value;

        private ValidationState(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static ValidationState None { get; } = new ValidationState("None");
        public static ValidationState Pending { get; } = new ValidationState("Pending");
        public static ValidationState Approved { get; } = new ValidationState("Approved");
        public static ValidationState Failed { get; } = new ValidationState("Failed");

        public static bool operator ==(ValidationState left, ValidationState right) => left.Equals(right);
        public static bool operator !=(ValidationState left, ValidationState right) => !left.Equals(right);

        public static explicit operator string(ValidationState value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is ValidationState other && Equals(other);
        public bool Equals(ValidationState other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }
}
