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

If any `argv` is not provided, it will output the credentials in JSON format.

# Requirements

* go 1.22
* Unix

# LICENSE

See LICENCE
