# Versioning Summary

## Next default version selection process

1. Read full spec
2. Remove currently deprecated APIs (as these will be remove next version)
3. Find candidate versions - prefer stable versions if we have recent stable versions (i.e. within 1 year of a preview)
4. Find minimal set of API versions which can provide all resources
5. Pick latest version for each resource

## Version pipeline

1. Run `az provider list`, cache into `provider_list.json`, summarise in `active.json` report
2. Read all specs from original Azure repository
3. Write `spec.json` and `spec-resources.json` reports
4. Use the existing versioner to generate `v1.json`
5. From `v1.json` we can generate `deprecated.json` and `pending.json`
6. `v1.json` is used by the gen program to build `schema.json` and `metadata.json`

Expressed as equations:

- az provider list => provider_list & active
- azure-rest-api-specs => spec & spec-resources
- provider_list + spec => v1 & deprecated & pending
- spec + v1 => schema/metadata
- spec + active + deprecated => v2
