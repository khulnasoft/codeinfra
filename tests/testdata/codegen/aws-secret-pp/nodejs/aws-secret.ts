import * as codeinfra from "@codeinfra/codeinfra";
import * as aws from "@codeinfra/aws";

const dbCluster = new aws.rds.Cluster("dbCluster", {masterPassword: codeinfra.secret("foobar")});
