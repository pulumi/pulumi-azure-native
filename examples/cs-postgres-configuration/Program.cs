// Copyright 2024, Pulumi Corporation.  All rights reserved.

using System.Threading.Tasks;
using Pulumi;

partial class Program
{
    static Task<int> Main() => Deployment.RunAsync<MyStack>();
}
