# Codeinfra Python SDK

The Codeinfra Python SDK (codeinfra) is the core package used when writing Codeinfra programs in Python. It contains everything that you’ll need in order to interact with Codeinfra resource providers and express infrastructure using Python code. Codeinfra resource providers all depend on this library and express their resources in terms of the types defined in this module.

The Codeinfra Python SDK requires a [supported version](https://devguide.python.org/versions/#versions) of Python.

note:
pip is required to install dependencies. If you installed Python from source, with an installer from [python.org](https://python.org/), or via [Homebrew](https://brew.sh/) you should already have pip. If Python is installed using your OS package manager, you may have to install pip separately, see [Installing pip/setuptools/wheel with Linux Package Managers](https://packaging.python.org/guides/installing-using-linux-tools/). For example, on Debian/Ubuntu you must run sudo apt install python3-venv python3-pip.

## Getting Started

The fastest way to get up and running is to choose from one of the following Getting Started guides:
-[aws](https://www.codeinfra.com/docs/get-started/aws/?language=python)
-[microsoft azure](https://www.codeinfra.com/docs/get-started/azure/?language=python)
-[google cloud](https://www.codeinfra.com/docs/get-started/gcp/?language=python)
-[kubernetes](https://www.codeinfra.com/docs/get-started/kubernetes/?language=python)

## Codeinfra Programming Model

The Codeinfra programming model defines the core concepts you will use when creating infrastructure as code programs using Codeinfra. Architecture & Concepts describes these concepts with examples available in Python. These concepts are made available to you in the Codeinfra SDK.

The Codeinfra SDK is available to Python developers as a Pip package distributed on [PyPI](https://www.codeinfra.com/docs/intro/languages/python/#pypi-packages) . To learn more, [refer to the Codeinfra SDK Reference Guide](https://www.codeinfra.com/docs/reference/pkg/python/codeinfra/).

The Codeinfra programming model includes a core concept of Input and Output values, which are used to track how outputs of one resource flow in as inputs to another resource. This concept is important to understand when getting started with Python and Codeinfra, and the [Inputs and Outputs] (https://www.codeinfra.com/docs/intro/concepts/inputs-outputs/)documentation is recommended to get a feel for how to work with this core part of Codeinfra in common cases.


## [The Codeinfra Python Resource Model](https://www.codeinfra.com/docs/reference/pkg/python/codeinfra/#the-codeinfra-python-resource-model-1)

Like most languages usable with Codeinfra, Codeinfra represents cloud resources as classes and Python programs can instantiate those classes. All classes that can be instantiated to produce actual resources derive from the codeinfra.Resource class.