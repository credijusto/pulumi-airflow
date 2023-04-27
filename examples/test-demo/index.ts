import * as pulumi from "@pulumi/pulumi";
import * as airflow from "@pulumi/airflow";

const example = new airflow.Connection("example", {
    connType: "example",
    connectionId: "example",
});
