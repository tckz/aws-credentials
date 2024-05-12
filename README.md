aws-credentials
===

Retrieve AWS credentials from the AWS CLI configuration and execute a command using these credentials via environment variables.

It simplifies launching a command that does not support .aws/config or AWS SSO.

# Usage

```
Usage: aws-credentials [options] [argv...]
  -profile string
    	profile
  -region string
    	region
  -version
    	show version
```

```
# Default credential resolution by AWS SDK
$ aws-credentials aws sts get-caller-identity

# Specify profile
$ aws-credentials --profile SOME_PROFILE aws sts get-caller-identity
```

If any `argv` is not provided, it will output the credentials in JSON format.


# Installation

https://github.com/tckz/aws-credentials/releases or
```
go install github.com/tckz/aws-credentials@latest
```


# Requirements

* go 1.22
* Unix


# LICENSE

See LICENCE
