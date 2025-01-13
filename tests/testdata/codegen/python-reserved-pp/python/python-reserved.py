import codeinfra
import codeinfra_lambda as lambda_

assert_ = lambda_.lambda_.Lambda("assert", lambda_="dns")
codeinfra.export("global", assert_.lambda_)
