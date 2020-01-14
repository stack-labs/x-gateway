package cmd

import (
	"fmt"

	ccli "github.com/micro/cli"
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/config/cmd"
	"github.com/micro-in-cn/x-gateway/api"
	"github.com/micro-in-cn/x-gateway/debug"
	"github.com/micro-in-cn/x-gateway/internal/platform"
	"github.com/micro-in-cn/x-gateway/plugin"
	"github.com/micro-in-cn/x-gateway/plugin/build"
	"github.com/micro-in-cn/x-gateway/web"
)

//App Info Vars
var (
	GitCommit string
	GitTag    string
	BuildDate string

	name        = "x-gateway"
	description = "A micro apigateway"
	version     = "0.0.1"
)

func init() {
	// setup the build plugin
	plugin.Register(build.Flags())

	// set platform build date
	platform.Version = BuildDate
}

func setup(app *ccli.App) {
	app.Flags = append(app.Flags,
		ccli.BoolFlag{
			Name:  "local",
			Usage: "Enable local only development",
		},
		ccli.BoolFlag{
			Name:   "enable_acme",
			Usage:  "Enables ACME support via Let's Encrypt. ACME hosts should also be specified.",
			EnvVar: "MICRO_ENABLE_ACME",
		},
		ccli.StringFlag{
			Name:   "acme_hosts",
			Usage:  "Comma separated list of hostnames to manage ACME certs for",
			EnvVar: "MICRO_ACME_HOSTS",
		},
		ccli.StringFlag{
			Name:   "acme_provider",
			Usage:  "The provider that will be used to communicate with Let's Encrypt. Valid options: autocert, certmagic",
			EnvVar: "MICRO_ACME_PROVIDER",
		},
		ccli.BoolFlag{
			Name:   "enable_tls",
			Usage:  "Enable TLS support. Expects cert and key file to be specified",
			EnvVar: "MICRO_ENABLE_TLS",
		},
		ccli.StringFlag{
			Name:   "tls_cert_file",
			Usage:  "Path to the TLS Certificate file",
			EnvVar: "MICRO_TLS_CERT_FILE",
		},
		ccli.StringFlag{
			Name:   "tls_key_file",
			Usage:  "Path to the TLS Key file",
			EnvVar: "MICRO_TLS_KEY_FILE",
		},
		ccli.StringFlag{
			Name:   "tls_client_ca_file",
			Usage:  "Path to the TLS CA file to verify clients against",
			EnvVar: "MICRO_TLS_CLIENT_CA_FILE",
		},
		ccli.StringFlag{
			Name:   "api_address",
			Usage:  "Set the api address e.g 0.0.0.0:8080",
			EnvVar: "MICRO_API_ADDRESS",
		},
		ccli.StringFlag{
			Name:   "gateway_address",
			Usage:  "Set the micro default gateway address e.g. :9094",
			EnvVar: "MICRO_GATEWAY_ADDRESS",
		},
		ccli.StringFlag{
			Name:   "api_handler",
			Usage:  "Specify the request handler to be used for mapping HTTP requests to services; {api, proxy, rpc}",
			EnvVar: "MICRO_API_HANDLER",
		},
		ccli.StringFlag{
			Name:   "api_namespace",
			Usage:  "Set the namespace used by the API e.g. com.example.api",
			EnvVar: "MICRO_API_NAMESPACE",
		},
		ccli.BoolFlag{
			Name:   "auto_update",
			Usage:  "Enable automatic updates",
			EnvVar: "MICRO_AUTO_UPDATE",
		},
		ccli.BoolTFlag{
			Name:   "report_usage",
			Usage:  "Report usage statistics",
			EnvVar: "MICRO_REPORT_USAGE",
		},
		ccli.StringFlag{
			Name:   "namespace",
			Usage:  "Set the micro service namespace",
			EnvVar: "MICRO_NAMESPACE",
			Value:  "go.micro",
		},
	)

	plugins := plugin.Plugins()

	for _, p := range plugins {
		if flags := p.Flags(); len(flags) > 0 {
			app.Flags = append(app.Flags, flags...)
		}

		if cmds := p.Commands(); len(cmds) > 0 {
			app.Commands = append(app.Commands, cmds...)
		}
	}

	before := app.Before

	app.Before = func(ctx *ccli.Context) error {
		if len(ctx.String("api_handler")) > 0 {
			api.Handler = ctx.String("api_handler")
		}
		if len(ctx.String("api_address")) > 0 {
			api.Address = ctx.String("api_address")
		}
		if len(ctx.String("api_namespace")) > 0 {
			api.Namespace = ctx.String("api_namespace")
		}
		for _, p := range plugins {
			if err := p.Init(ctx); err != nil {
				return err
			}
		}

		// now do previous before
		return before(ctx)
	}
}

func buildVersion() string {
	microVersion := version

	if GitTag != "" {
		microVersion = GitTag
	}

	if GitCommit != "" {
		microVersion += fmt.Sprintf("-%s", GitCommit)
	}

	if BuildDate != "" {
		microVersion += fmt.Sprintf("-%s", BuildDate)
	}

	return microVersion
}

// Init initialised the command line
func Init(options ...micro.Option) {
	Setup(cmd.App(), options...)

	regularArguments(cmd.App())

	cmd.Init(
		cmd.Name(name),
		cmd.Description(description),
		cmd.Version(buildVersion()),
	)
}

// Setup sets up a cli.App
func Setup(app *ccli.App, options ...micro.Option) {
	// Add the various commands
	app.Commands = append(app.Commands, api.Commands(options...)...)
	app.Commands = append(app.Commands, debug.Commands(options...)...)
	app.Commands = append(app.Commands, build.Commands()...)
	app.Commands = append(app.Commands, web.Commands(options...)...)

	// add the init command for our internal operator
	app.Commands = append(app.Commands, ccli.Command{
		Name:  "init",
		Usage: "Run the micro operator",
		Action: func(c *ccli.Context) {
			platform.Init(c)
		},
		Flags: []ccli.Flag{},
	})

	// boot micro runtime
	app.Action = platform.Run

	setup(app)
}
