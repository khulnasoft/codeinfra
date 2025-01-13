import codeinfra
import codeinfra_aws as aws

db_cluster = aws.rds.Cluster("dbCluster", master_password=codeinfra.Output.secret("foobar"))
