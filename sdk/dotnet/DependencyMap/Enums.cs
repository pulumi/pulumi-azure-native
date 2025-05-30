// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.ComponentModel;
using Pulumi;

namespace Pulumi.AzureNative.DependencyMap
{
    /// <summary>
    /// Operator for process name filter
    /// </summary>
    [EnumType]
    public readonly struct ProcessNameFilterOperator : IEquatable<ProcessNameFilterOperator>
    {
        private readonly string _value;

        private ProcessNameFilterOperator(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        /// <summary>
        /// Operator to include items in the result
        /// </summary>
        public static ProcessNameFilterOperator Contains { get; } = new ProcessNameFilterOperator("contains");
        /// <summary>
        /// Operator to exclude items in the result
        /// </summary>
        public static ProcessNameFilterOperator NotContains { get; } = new ProcessNameFilterOperator("notContains");

        public static bool operator ==(ProcessNameFilterOperator left, ProcessNameFilterOperator right) => left.Equals(right);
        public static bool operator !=(ProcessNameFilterOperator left, ProcessNameFilterOperator right) => !left.Equals(right);

        public static explicit operator string(ProcessNameFilterOperator value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is ProcessNameFilterOperator other && Equals(other);
        public bool Equals(ProcessNameFilterOperator other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    /// <summary>
    /// Source type of Discovery Source resource.
    /// </summary>
    [EnumType]
    public readonly struct SourceType : IEquatable<SourceType>
    {
        private readonly string _value;

        private SourceType(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        /// <summary>
        /// OffAzure source type
        /// </summary>
        public static SourceType OffAzure { get; } = new SourceType("OffAzure");

        public static bool operator ==(SourceType left, SourceType right) => left.Equals(right);
        public static bool operator !=(SourceType left, SourceType right) => !left.Equals(right);

        public static explicit operator string(SourceType value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is SourceType other && Equals(other);
        public bool Equals(SourceType other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }
}
