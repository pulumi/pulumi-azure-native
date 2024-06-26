// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.ComponentModel;
using Pulumi;

namespace Pulumi.AzureNative.Kusto.V20180907Preview
{
    /// <summary>
    /// The data format of the message. Optionally the data format can be added to each message.
    /// </summary>
    [EnumType]
    public readonly struct DataFormat : IEquatable<DataFormat>
    {
        private readonly string _value;

        private DataFormat(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static DataFormat MULTIJSON { get; } = new DataFormat("MULTIJSON");
        public static DataFormat JSON { get; } = new DataFormat("JSON");
        public static DataFormat CSV { get; } = new DataFormat("CSV");

        public static bool operator ==(DataFormat left, DataFormat right) => left.Equals(right);
        public static bool operator !=(DataFormat left, DataFormat right) => !left.Equals(right);

        public static explicit operator string(DataFormat value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is DataFormat other && Equals(other);
        public bool Equals(DataFormat other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }
}
