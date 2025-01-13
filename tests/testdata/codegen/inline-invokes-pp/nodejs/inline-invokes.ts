import * as codeinfra from "@codeinfra/codeinfra";
import * as aws from "@codeinfra/aws";

const webSecurityGroup = new aws.ec2.SecurityGroup("webSecurityGroup", {vpcId: aws.ec2.getVpc({
    "default": true,
}).then(invoke => invoke.id)});
