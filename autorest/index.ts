import { glob } from "fast-glob";
import { writeFile, readFile } from "fs/promises";
import * as path from "path";
import { fileURLToPath } from "url";
import { ConfigurationLoader } from "@autorest/configuration";
import {
  ConsoleLogger,
  FilterLogger,
  AutorestSyncLogger,
} from "@autorest/common";
import { createFolderUri, createFileOrFolderUri } from "@azure-tools/uri";
import yaml from "yaml";

(async () => {
  const logger = new AutorestSyncLogger({
    sinks: [new ConsoleLogger()],
    processors: [new FilterLogger({ level: "error" })],
  });
  const defaultConfigUri = createFileOrFolderUri(
    require.resolve(
      "@autorest/configuration/resources/default-configuration.md"
    )
  );

  const specificationRootDir = path.resolve("..", "azure-rest-api-specs");
  const readmeFiles = await glob("specification/*/resource-manager/readme.md", {
    cwd: specificationRootDir,
  });

  const result: string[] = [];
  for (const readmeFile of readmeFiles) {
    const configFolderUri = createFolderUri(
      path.dirname(path.resolve(specificationRootDir, readmeFile))
    );
    const configLoader = new ConfigurationLoader(
      logger,
      defaultConfigUri,
      configFolderUri
    );
    const config = await configLoader.load([], true);
    const inputFiles = config.config.inputFileUris.map((uri) =>
      path.relative(specificationRootDir, fileURLToPath(uri))
    );
    result.push(...inputFiles);
  }
  result.sort();
  writeFile(
    "../reports/autorest-input-files.json",
    JSON.stringify(result, null, 2)
  );

  const autorestSpecFilePathsLookup = new Set<string>(result);

  const allResourcesByVersionContent = await readFile(
    "../reports/allResourcesByVersion.json",
    "utf-8"
  );
  const allResourcesByVersion: Record<
    ProviderModule,
    Record<ApiVersion, Record<ResourceName, VersionEntry>>
  > = JSON.parse(allResourcesByVersionContent);
  const specVersions: Record<
    ProviderModule,
    Record<ResourceName, VersionEntry>
  > = {};
  for (const [providerModule, versions] of Object.entries(
    allResourcesByVersion
  )) {
    const module: Record<ResourceName, VersionEntry> = {};
    specVersions[providerModule] = module;
    for (const resources of Object.values(versions)) {
      for (const [resourceName, entry] of Object.entries(resources)) {
        if (autorestSpecFilePathsLookup.has(entry.SpecFilePath)) {
          module[resourceName] = entry;
        }
      }
    }
  }
  writeFile("../reports/autorest.yaml", yaml.stringify(specVersions, null, 2));
})().catch(console.error);

type ProviderModule = string;
type ApiVersion = string;
type ResourceName = string;
type VersionEntry = {
  ApiVersion: ApiVersion;
  SpecFilePath: string;
  ResourceUri: string;
  RpNamespace: string;
};
