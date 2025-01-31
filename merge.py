import json

# Load the JSON files
with open('reports/oldApiVersions.json', 'r') as old_file:
    old_versions = json.load(old_file)

with open('versions/v3-removed.json', 'r') as v3_file:
    v3_versions = json.load(v3_file)

# Merge the dictionaries
for key, versions in old_versions.items():
    if key in v3_versions:
        # Merge and sort the version arrays
        v3_versions[key] = sorted(set(v3_versions[key] + versions))
    else:
        # Add new key and versions
        v3_versions[key] = sorted(versions)

# Write the merged dictionary back to v3-removed.json
with open('v3-removed.json', 'w') as v3_file:
    json.dump(v3_versions, v3_file, indent=2)