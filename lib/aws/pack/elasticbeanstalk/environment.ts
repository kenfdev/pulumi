// *** WARNING: this file was generated by the Lumi IDL Compiler (LUMIDL). ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as lumi from "@lumi/lumi";

import {Application} from "./application";
import {ApplicationVersion} from "./applicationVersion";

export let WebServerTier: Tier = "WebServer::Standard::1.0";
export let WorkerTier: Tier = "Worker::SQS/HTTP::1.0";

export class Environment extends lumi.Resource implements EnvironmentArgs {
    public readonly name: string;
    public readonly application: Application;
    public readonly cnamePrefix?: string;
    public description?: string;
    public readonly environmentName?: string;
    public optionSettings?: OptionSetting[];
    public readonly solutionStackName?: string;
    public readonly tags?: Tag[];
    public templateName?: string;
    public readonly tier?: Tier;
    public version?: ApplicationVersion;
    @lumi.out public endpointURL: string;

    constructor(name: string, args: EnvironmentArgs) {
        super();
        if (name === undefined) {
            throw new Error("Missing required resource name");
        }
        this.name = name;
        if (args.application === undefined) {
            throw new Error("Missing required argument 'application'");
        }
        this.application = args.application;
        this.cnamePrefix = args.cnamePrefix;
        this.description = args.description;
        this.environmentName = args.environmentName;
        this.optionSettings = args.optionSettings;
        this.solutionStackName = args.solutionStackName;
        this.tags = args.tags;
        this.templateName = args.templateName;
        this.tier = args.tier;
        this.version = args.version;
    }
}

export interface EnvironmentArgs {
    readonly application: Application;
    readonly cnamePrefix?: string;
    description?: string;
    readonly environmentName?: string;
    optionSettings?: OptionSetting[];
    readonly solutionStackName?: string;
    readonly tags?: Tag[];
    templateName?: string;
    readonly tier?: Tier;
    version?: ApplicationVersion;
}

export interface OptionSetting {
    namespace: string;
    optionName: string;
    value: string;
}

export interface Tag {
    key: string;
    value: string;
}

export type Tier =
    "WebServer::Standard::1.0" |
    "Worker::SQS/HTTP::1.0";


