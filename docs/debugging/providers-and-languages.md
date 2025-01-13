# Providers and Languages

Languages and providers run as their own executables, so debugging isn't as
simple as attaching to the engine.

Instead you'll have to launch them in a debugger and provide the information
for connecting to them to the engine via environmental variables.

## Languages

Languages work very similary to [Providers](#Providers).

Instead of CODEINFRA_DEBUG_PROVIDERS environmental variable, you use CODEINFRA_DEBUG_LANGUAGES like so:

```sh
CODEINFRA_DEBUG_LANGUAGES="go:12345" codeinfra preview
```

After opening your language executable (eg `codeinfra-language-go`) in your debugger of choice.
