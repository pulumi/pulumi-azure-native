// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.ComponentModel;
using Pulumi;

namespace Pulumi.AzureNative.DataShare.V20201001Preview
{
    /// <summary>
    /// Kind of data set.
    /// </summary>
    [EnumType]
    public readonly struct DataSetKind : IEquatable<DataSetKind>
    {
        private readonly string _value;

        private DataSetKind(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static DataSetKind Blob { get; } = new DataSetKind("Blob");
        public static DataSetKind Container { get; } = new DataSetKind("Container");
        public static DataSetKind BlobFolder { get; } = new DataSetKind("BlobFolder");
        public static DataSetKind AdlsGen2FileSystem { get; } = new DataSetKind("AdlsGen2FileSystem");
        public static DataSetKind AdlsGen2Folder { get; } = new DataSetKind("AdlsGen2Folder");
        public static DataSetKind AdlsGen2File { get; } = new DataSetKind("AdlsGen2File");
        public static DataSetKind AdlsGen1Folder { get; } = new DataSetKind("AdlsGen1Folder");
        public static DataSetKind AdlsGen1File { get; } = new DataSetKind("AdlsGen1File");
        public static DataSetKind AdlsGen2StorageAccount { get; } = new DataSetKind("AdlsGen2StorageAccount");
        public static DataSetKind KustoCluster { get; } = new DataSetKind("KustoCluster");
        public static DataSetKind KustoDatabase { get; } = new DataSetKind("KustoDatabase");
        public static DataSetKind SqlDBTable { get; } = new DataSetKind("SqlDBTable");
        public static DataSetKind SqlDWTable { get; } = new DataSetKind("SqlDWTable");
        public static DataSetKind SynapseWorkspaceSqlPoolTable { get; } = new DataSetKind("SynapseWorkspaceSqlPoolTable");
        public static DataSetKind BlobStorageAccount { get; } = new DataSetKind("BlobStorageAccount");

        public static bool operator ==(DataSetKind left, DataSetKind right) => left.Equals(right);
        public static bool operator !=(DataSetKind left, DataSetKind right) => !left.Equals(right);

        public static explicit operator string(DataSetKind value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is DataSetKind other && Equals(other);
        public bool Equals(DataSetKind other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    /// <summary>
    /// Kind of data set mapping.
    /// </summary>
    [EnumType]
    public readonly struct DataSetMappingKind : IEquatable<DataSetMappingKind>
    {
        private readonly string _value;

        private DataSetMappingKind(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static DataSetMappingKind Blob { get; } = new DataSetMappingKind("Blob");
        public static DataSetMappingKind Container { get; } = new DataSetMappingKind("Container");
        public static DataSetMappingKind BlobFolder { get; } = new DataSetMappingKind("BlobFolder");
        public static DataSetMappingKind AdlsGen2FileSystem { get; } = new DataSetMappingKind("AdlsGen2FileSystem");
        public static DataSetMappingKind AdlsGen2Folder { get; } = new DataSetMappingKind("AdlsGen2Folder");
        public static DataSetMappingKind AdlsGen2File { get; } = new DataSetMappingKind("AdlsGen2File");
        public static DataSetMappingKind AdlsGen2StorageAccount { get; } = new DataSetMappingKind("AdlsGen2StorageAccount");
        public static DataSetMappingKind KustoCluster { get; } = new DataSetMappingKind("KustoCluster");
        public static DataSetMappingKind KustoDatabase { get; } = new DataSetMappingKind("KustoDatabase");
        public static DataSetMappingKind SqlDBTable { get; } = new DataSetMappingKind("SqlDBTable");
        public static DataSetMappingKind SqlDWTable { get; } = new DataSetMappingKind("SqlDWTable");
        public static DataSetMappingKind SynapseWorkspaceSqlPoolTable { get; } = new DataSetMappingKind("SynapseWorkspaceSqlPoolTable");
        public static DataSetMappingKind BlobStorageAccount { get; } = new DataSetMappingKind("BlobStorageAccount");

        public static bool operator ==(DataSetMappingKind left, DataSetMappingKind right) => left.Equals(right);
        public static bool operator !=(DataSetMappingKind left, DataSetMappingKind right) => !left.Equals(right);

        public static explicit operator string(DataSetMappingKind value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is DataSetMappingKind other && Equals(other);
        public bool Equals(DataSetMappingKind other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }
}
