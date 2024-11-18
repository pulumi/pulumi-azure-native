package squeeze

import (
	"fmt"
	"sort"
	"strings"
	"time"

	mapset "github.com/deckarep/golang-set/v2"

	"github.com/pulumi/pulumi/pkg/v3/codegen"
	"github.com/pulumi/pulumi/pkg/v3/codegen/schema"
)

// CompareAll returns a map of resource tokens to be removed with an optional token to a resource it should be replaced with.
func CompareAll(sch *schema.PackageSpec) (map[string]string, error) {
	resourceMap := map[string]mapset.Set[string]{}
	for name := range sch.Resources {
		if !isVersionedName(name) {
			continue
		}

		vls := versionlessName(name)
		if existing, ok := resourceMap[vls]; ok {
			existing.Add(name)
		} else {
			resourceMap[vls] = mapset.NewSet(name)
		}
	}

	sortedKeys := codegen.SortedKeys(resourceMap)
	replacements := map[string]string{}
	for _, name := range sortedKeys {
		group := resourceMap[name]
		unique := calculateUniqueVersions(sch, group)
		reduced := group.Difference(unique)
		for r := range reduced.Iter() {
			fmt.Println(r)
		}
		for k := range reduced.Iter() {
			for _, a := range mapset.Sorted(unique) {
				if a > k {
					replacements[k] = a
					break
				}
			}
		}
	}

	return replacements, nil
}

func compareResources(sch *schema.PackageSpec, oldName string, newName string) ([]string, error) {
	var violations []string
	oldRes, ok := sch.Resources[oldName]
	if !ok {
		return nil, fmt.Errorf("resource %q missing", oldName)
	}
	newRes, ok := sch.Resources[newName]
	if !ok {
		return nil, fmt.Errorf("resource %q missing", newName)
	}

	for propName, prop := range oldRes.InputProperties {
		newProp, ok := newRes.InputProperties[propName]
		if !ok {
			violations = append(violations, fmt.Sprintf("Resource %q missing input %q", newName, propName))
			continue
		}

		vs := validateTypesDeep(sch, &prop.TypeSpec, &newProp.TypeSpec, fmt.Sprintf("Resource %q input %q", newName, propName), true)
		violations = append(violations, vs...)
	}

	for propName, prop := range oldRes.Properties {
		newProp, ok := newRes.Properties[propName]
		if !ok {
			violations = append(violations, fmt.Sprintf("Resource %q missing output %q", newName, propName))
			continue
		}

		vs := validateTypesDeep(sch, &prop.TypeSpec, &newProp.TypeSpec, fmt.Sprintf("Resource %q output %q", newName, propName), false)
		violations = append(violations, vs...)
	}

	oldRequiredSet := mapset.NewSet(oldRes.RequiredInputs...)
	for _, propName := range newRes.RequiredInputs {
		if !oldRequiredSet.Contains(propName) {
			violations = append(violations, fmt.Sprintf("Resource %q has a new required input %q", newName, propName))
		}
	}

	newRequiredSet := mapset.NewSet(newRes.Required...)
	for _, propName := range oldRes.Required {
		if !newRequiredSet.Contains(propName) {
			violations = append(violations, fmt.Sprintf("Resource %q has output %q that is not required anymore", newName, propName))
		}
	}

	return violations, nil
}

func calculateUniqueVersions(sch *schema.PackageSpec, resVersions mapset.Set[string]) mapset.Set[string] {
	uniqueVersions := mapset.NewSet[string]()

	sortedVersions := mapset.Sorted(resVersions)
	sortApiVersions(sortedVersions)

outer:
	for _, oldName := range sortedVersions {
		for _, newName := range sortedVersions {
			if oldName >= newName {
				continue
			}
			violations, err := compareResources(sch, oldName, newName)
			if err == nil && len(violations) == 0 {
				continue outer
			}
		}
		uniqueVersions.Add(oldName)
	}
	return uniqueVersions
}

func apiVersionToDate(apiVersion string) (time.Time, error) {
	if len(apiVersion) < 9 {
		return time.Time{}, fmt.Errorf("invalid API version %q", apiVersion)
	}
	// The API version is in the format YYYY-MM-DD - ignore suffixes like "-preview".
	return time.Parse("20060102", apiVersion[1:9])
}

func compareApiVersions(a, b string) int {
	timeA, err := apiVersionToDate(a)
	if err != nil {
		return strings.Compare(a, b)
	}
	timeB, err := apiVersionToDate(b)
	if err != nil {
		return strings.Compare(a, b)
	}
	timeDiff := timeA.Compare(timeB)
	if timeDiff != 0 {
		return timeDiff
	}

	// Sort private first, preview second, stable last.
	aPrivate := isPrivate(a)
	bPrivate := isPrivate(b)
	if aPrivate != bPrivate {
		if aPrivate {
			return -1
		}
		return 1
	}
	aPreview := isPreview(a)
	bPreview := isPreview(b)
	if aPreview != bPreview {
		if aPreview {
			return -1
		}
		return 1
	}
	return 0
}

func isPreview(apiVersion string) bool {
	lower := strings.ToLower(apiVersion)
	return strings.Contains(lower, "preview") || strings.Contains(lower, "beta")
}

func isPrivate(apiVersion string) bool {
	lower := strings.ToLower(apiVersion)
	return strings.Contains(lower, "private")
}

func sortApiVersions(versions []string) {
	sort.SliceStable(versions, func(i, j int) bool {
		return compareApiVersions(versions[i], versions[j]) < 0
	})
}

func validateTypesDeep(sch *schema.PackageSpec, old *schema.TypeSpec, new *schema.TypeSpec, prefix string, input bool) (violations []string) {
	switch {
	case old == nil && new == nil:
		return
	case old != nil && new == nil:
		violations = append(violations, fmt.Sprintf("had %+v but now has no type", old))
		return
	case old == nil && new != nil:
		violations = append(violations, fmt.Sprintf("had no type but now has %+v", new))
		return
	}

	oldType := old.Type
	if old.Ref != "" {
		oldType = old.Ref
	}
	newType := new.Type
	if new.Ref != "" {
		newType = new.Ref
	}
	if oldType != newType {
		if strings.HasPrefix(oldType, "#/types/azure-native") && //azure-native:resources/v20210101:MyType
			strings.HasPrefix(newType, "#/types/azure-native") &&
			versionlessName(oldType) == versionlessName(newType) { // resources:MyType
			// Both are reference types, let's do a deep comparison
			oldTypeRef := sch.Types[oldType]
			newTypeRef := sch.Types[newType]
			for propName, prop := range oldTypeRef.Properties {
				newProp, ok := newTypeRef.Properties[propName]
				if !ok {
					violations = append(violations, fmt.Sprintf("Type %q missing input %q", newType, propName))
					continue
				}

				vs := validateTypesDeep(sch, &prop.TypeSpec, &newProp.TypeSpec, fmt.Sprintf("Type %q input %q", newType, propName), input)
				violations = append(violations, vs...)
			}

			if input {
				oldRequiredSet := mapset.NewSet(oldTypeRef.Required...)
				for _, propName := range newTypeRef.Required {
					if !oldRequiredSet.Contains(propName) {
						violations = append(violations, fmt.Sprintf("Type %q has a new required input %q", newType, propName))
					}
				}
			} else {
				newRequiredSet := mapset.NewSet(newTypeRef.Required...)
				for _, propName := range oldTypeRef.Required {
					if !newRequiredSet.Contains(propName) {
						violations = append(violations, fmt.Sprintf("Type %q has output %q that is not required anymore", newType, propName))
					}
				}
			}
		} else {
			violations = append(violations, fmt.Sprintf("%s type changed from %q to %q", prefix, oldType, newType))
		}
	}
	violations = append(violations, validateTypesDeep(sch, old.Items, new.Items, prefix+" items", input)...)
	violations = append(violations, validateTypesDeep(sch, old.AdditionalProperties, new.AdditionalProperties, prefix+" additional properties", input)...)
	return
}

// Is it of the form "azure-native:appplatform/v20230101preview" or just "azure-native:appplatform"?
func isVersionedName(name string) bool {
	return strings.Contains(name, "/v")
}

// "azure-native:appplatform/v20230101preview" -> "appplatform"
func versionlessName(name string) string {
	parts := strings.Split(name, ":")
	mod := parts[1]
	modParts := strings.Split(mod, "/")
	if len(modParts) == 2 {
		mod = modParts[0]
	}
	return fmt.Sprintf("%s:%s", mod, parts[2])
}
