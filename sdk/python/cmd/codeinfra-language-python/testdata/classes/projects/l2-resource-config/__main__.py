import codeinfra
import codeinfra_config as config

prov = config.Provider("prov",
    name="my config",
    plugin_download_url="not the same as the codeinfra resource option")
# Note this isn't _using_ the explicit provider, it's just grabbing a value from it.
res = config.Resource("res", text=prov.version)
codeinfra.export("pluginDownloadURL", prov.plugin_download_url)
