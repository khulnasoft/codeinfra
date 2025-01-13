(providers)=
# Providers

The term "provider" can mean different things in different contexts. When we
talk about Codeinfra programs, we often talk about *provider resources* such as
that provided by the `aws.Provider` class in the `@codeinfra/aws` NodeJS/TypeScript
package. Or, we might simply mean a cloud provider, such as AWS or GCP. In the
context of the wider Codeinfra architecture though, a provider (specifically, a
*resource provider*) is a Codeinfra [plugin](plugins) that implements [a
standardized gRPC interface](codeinfrarpc.ResourceProvider) for handling
communication with a third-party service (usually a cloud service, such as AWS,
GCP, or Azure):

* *Configuration* methods are designed to allow a consumer to configure a
  provider instance in some way. The [](codeinfrarpc.ResourceProvider.Configure)
  call is the most common example of this, allowing a caller to e.g. specify the
  AWS region that a provider should use operate in.
  [](codeinfrarpc.ResourceProvider.Parameterize) is a similar method that operates
  at a higher level, allowing a caller to influence more deeply how a provider
  works (see [the section on parameterized providers](parameterized-providers)
  for more).
* *Schema* endpoints allow a caller to interrogate the resources and functions
  that a provider exposes. The [](codeinfrarpc.ResourceProvider.GetSchema) method
  returns a provider's schema, which includes the set of resources and functions
  that the provider supports, as well as the properties and inputs that each
  resource and function expects. This is the primary driver for the various code
  generation processes that Codeinfra uses, such as that underpinning SDK
  generation.
* *Lifecycle* methods expose the typical [](codeinfrarpc.ResourceProvider.Create),
  [](codeinfrarpc.ResourceProvider.Read), [](codeinfrarpc.ResourceProvider.Update),
  and [](codeinfrarpc.ResourceProvider.Delete) (CRUD) operations that allow clients
  to manage provider resources. The [](codeinfrarpc.ResourceProvider.Check),
  [](codeinfrarpc.ResourceProvider.Diff), and
  [](codeinfrarpc.ResourceProvider.Construct) methods also fall into this category,
  as discussed in [resource registration](resource-registration).
* *Functions* can be invoked on a provider through the
  [](codeinfrarpc.ResourceProvider.Invoke) call, or on specific resources by using
  the [](codeinfrarpc.ResourceProvider.Call) operation. Functions are typically
  used to perform operations that don't fit into the CRUD model, such as
  retrieving a list of availability zones, or available regions, etc.

While any program which implements the [](codeinfrarpc.ResourceProvider) interface
can be interfaced with by the Codeinfra engine, in practice most Codeinfra providers
are built in a handful of ways:

* *Bridged providers* wrap a Terraform provider using the [Codeinfra Terraform
  bridge](https://github.com/khulnasoft/codeinfra-terraform-bridge). The majority of
  Codeinfra providers are built in this way.
* The [`codeinfra-go-provider`](https://github.com/khulnasoft/codeinfra-go-provider)
  library provides a high-level API for writing a provider in Go, without having
  to worry about low-level gRPC details.
* *Dynamic providers* can be written as part of a Codeinfra program. They are
  discussed in more detail [in their own section of this
  documentation](dynamic-providers).
* Each language SDK provides a `provider` package that offers low-level
  primitives for writing a provider in that language (e.g.
  <gh-file:codeinfra#sdk/nodejs/provider> for NodeJS,
  <gh-file:codeinfra#sdk/python/lib/codeinfra/provider> for Python, and
  <gh-file:codeinfra#sdk/go/codeinfra/provider> for Go). These packages are in varying
  states of completeness and neatness and are generally only used for building
  [component providers](component-providers).

A provider binary is typically named `codeinfra-resource-<provider-name>`;
`codeinfra-resource-aws` is one example. The following pages provide more
information on the various types of providers, their modes of operation, and
their implementation:

:::{toctree}
:maxdepth: 1
:titlesonly:

/docs/architecture/providers/default
/docs/architecture/providers/components
/docs/architecture/providers/dynamic
/docs/architecture/providers/parameterized
:::
