aws-credentials
===

Retrieve AWS credentials from the AWS CLI configuration and execute a command using these credentials via environment variables.

It simplifies launching a command that does not support .aws/config or AWS SSO.

# Usage

```
Usage: aws-credentials [options] [argv...]
  -export
    	Output AWS credentials as export variables format
  -profile string
    	profile
  -region string
    	region
  -version
    	Show version
```

```
# Default credential resolution by AWS SDK
$ aws-credentials aws sts get-caller-identity

# Specify profile
$ aws-credentials --profile SOME_PROFILE aws sts get-caller-identity
```

If any `argv` is not provided, it will output the credentials in JSON format.

```
# Output JSON format
$ aws-credentials
{
  "AccessKeyID": "...",
  "SecretAccessKey": "...",
  "SessionToken": "...",
  "Source": "SSOProvider",
  "CanExpire": true,
  "Expires": "2024-05-15T14:01:15Z"
}

# Output export variables format
$ aws-credentials --export
export AWS_ACCESS_KEY_ID="..."
export AWS_SECRET_ACCESS_KEY="..."
export AWS_SESSION_TOKEN="..."
```

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
