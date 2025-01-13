import * as codeinfra from "@codeinfra/codeinfra";
import * as aws from "@codeinfra/aws";

const test = new aws.fsx.OpenZfsFileSystem("test", {
    storageCapacity: 64,
    subnetIds: [aws_subnet.test1.id],
    deploymentType: "SINGLE_AZ_1",
    throughputCapacity: 64,
});
