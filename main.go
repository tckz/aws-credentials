package main

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"runtime/debug"
	"syscall"

	"github.com/aws/aws-sdk-go-v2/config"
)

var version string

var (
	optVersion = flag.Bool("version", false, "Show version")
	optProfile = flag.String("profile", "", "profile")
	optRegion  = flag.String("region", "", "region")
	optExport  = flag.Bool("export", false, "Output AWS credentials as export variables format")
)

const (
	accessKeyID  = "AWS_ACCESS_KEY_ID"
	secretKey    = "AWS_SECRET_ACCESS_KEY"
	sessionToken = "AWS_SESSION_TOKEN"
)

func main() {
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage: %s [options] [argv...]\n", filepath.Base(os.Args[0]))
		flag.PrintDefaults()
	}
	flag.Parse()

	if *optVersion {
		if version != "" {
			fmt.Println(version)
		} else if info, ok := debug.ReadBuildInfo(); ok {
			fmt.Println(info.Main.Version)
		}
		return
	}

	if err := run(); err != nil {
		log.Fatalf("*** %v", err)
	}
}

func run() error {
	ctx := context.Background()

	var opts []func(*config.LoadOptions) error
	if *optProfile != "" {
		opts = append(opts, config.WithSharedConfigProfile(*optProfile))
	}
	if *optRegion != "" {
		opts = append(opts, config.WithRegion(*optRegion))
	}

	cfg, err := config.LoadDefaultConfig(ctx, opts...)
	if err != nil {
		return fmt.Errorf("config.LoadDefaultConfig: %w", err)
	}

	creds, err := cfg.Credentials.Retrieve(ctx)
	if err != nil {
		return fmt.Errorf("Credentials.Retrieve: %w", err)
	}

	if len(flag.Args()) == 0 {
		if *optExport {
			fmt.Fprintf(os.Stdout, "export %s=%q\n", accessKeyID, creds.AccessKeyID)
			fmt.Fprintf(os.Stdout, "export %s=%q\n", secretKey, creds.SecretAccessKey)
			if creds.SessionToken != "" {
				fmt.Fprintf(os.Stdout, "export %s=%q\n", sessionToken, creds.SessionToken)
			}
			return nil
		} else {
			return json.NewEncoder(os.Stdout).Encode(creds)
		}
	}

	// Clear existing AWS credentials from environment
	for _, e := range []string{accessKeyID, secretKey, sessionToken} {
		os.Unsetenv(e)
	}

	argv0 := flag.Arg(0)
	if argv0 == filepath.Base(argv0) {
		lp, err := exec.LookPath(argv0)
		if err != nil {
			return fmt.Errorf("exec.LookPath %s: %w", argv0, err)
		}
		argv0 = lp
	}

	// Create new environment variables array with AWS credentials
	env := os.Environ()
	env = append(env, fmt.Sprintf("%s=%s", accessKeyID, creds.AccessKeyID))
	env = append(env, fmt.Sprintf("%s=%s", secretKey, creds.SecretAccessKey))
	if creds.SessionToken != "" {
		env = append(env, fmt.Sprintf("%s=%s", sessionToken, creds.SessionToken))
	}

	return syscall.Exec(argv0, flag.Args(), env)
}
