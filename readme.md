# Ball Clock Simulator
For requirements refer to the [PDF](Rakuten-BallClockAssignment.pdf). I had never heard of a Ball Clock before this, I found [this video](https://youtu.be/F7K6GIBWPQw) helpful to understanding what a Ball Clock was and how it worked.

## Executing the main application
There are two things that can be done.
- Unit Tests can be executed
- The Console application can be run

To do either of these things, respectfully, you can run the convenience scripts provided:
- `test-tdd.sh`
- `run-cli-simulation.sh`

## Executing as an AWS Lambda
Since IaaS and FaaS are relatively simple for small projects, and [Go Build supports build tags](https://dave.cheney.net/2013/10/12/how-to-use-conditional-compilation-with-the-go-build-tool) we can use the same codebase to run our CLI application as a REST endpoint. Compilation branching isn't usually advisable unless you're specifically doing it to support multiple platforms. In this case it's a simple way to include a secondary main method for AWS.

This requires:
- [aws-sam-cli](https://github.com/awslabs/aws-sam-cli) to be installed so we can run as an AWS Lambda
- [jq](https://stedolan.github.io/jq/) to be installed so we can pretty-print the output to the console

Once these requirements are met on your local machine:
In one terminal run `run-aws-simulation.sh` to start SAM and reserve the default port (3000). You can now open a new terminal (or move this job to the background with `CTRL+Z` and `bg`, but the output gets hard to read) and run the curls against the endpoint with `test-aws-curls.sh`.

If you wish to deploy this to an AWS environment, you must include a `CodeUri` as part of the cloud-formation template. This must point to a valid artifact in S3.