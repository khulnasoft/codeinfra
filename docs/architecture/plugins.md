(plugins)=
# Plugins

Plugins are Codeinfra's core extensibility mechanism, allowing the Codeinfra engine to
communicate in a uniform manner with various languages, resource providers, and
other tools. Generally speaking, plugins are run as separate processes and often
(though not always) communicated with over gRPC. Presently, Codeinfra supports the
following kinds of plugins:

* *Resource plugins* (or *resource providers*/*providers*, see
  [providers](providers) for more) expose a [standardized gRPC
  interface](codeinfrarpc.ResourceProvider) for managing resources (such as those
  in an AWS or GCP cloud).
* *Language plugins* host programs written in a particular language, allowing
  the engine to invoke Codeinfra programs without having to understand the
  specifics of their implementation.
* *Analyzer plugins* are used to analyze Codeinfra programs for potential issues
  before they are executed. Analyzers underpin [CrossGuard, Codeinfra's Policy as
  Code](https://www.codeinfra.com/docs/using-codeinfra/crossguard/) product.
* *Converter plugins* support the conversion of existing infrastructure as code
  (e.g. [Terraform](https://github.com/khulnasoft/codeinfra-converter-terraform)) to
  Codeinfra programs.
* *Tool plugins* allow integrating Codeinfra with arbitrary tools.

(plugin-loading-execution)=
(shimless)=
## Loading and execution

Plugins may be provided in one of two ways:

* As a binary that can be directly executed. Binaries are named
  `codeinfra-<kind>-<name>`, where `<kind>` is one of `resource`, `language`,
  `analyzer`, `converter`, or `tool`, and `<name>` is a unique name for the
  plugin. For example, the AWS resource provider plugin is named
  `codeinfra-resource-aws`.
* As a directory containing a `CodeinfraPlugin.yaml` file and a set of files that
  implement the plugin's functionality. The `CodeinfraPlugin.yaml` file specifies
  the `runtime` to be used to execute the provided files. The engine reads this
  and spawns the runtime's [language](language-hosts) plugin to run the plugin.
  The language plugin's [](codeinfrarpc.LanguageRuntime.RunPlugin) method is then
  used to execute the plugin. This method of running plugins is sometimes
  referred to as *shimless*, since prior to its introduction one would always
  have to provide a "shim" executable (such as a shell script or batch file)
  that did nothing but spawn the relevant interpreter over the provided files.
