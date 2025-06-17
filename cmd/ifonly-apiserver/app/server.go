package app

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/xiahuaxiahua0616/ifonly/cmd/ifonly-apiserver/app/options"
	"github.com/xiahuaxiahua0616/ifonly/internal/pkg/log"
	"github.com/xiahuaxiahua0616/ifonly/pkg/version"
)

var configFile string // 配置文件路径

func NewIfOnlyCommand() *cobra.Command {
	opts := options.NewServerOptions()
	cmd := &cobra.Command{
		// 指定命令的名字，该名字会出现在帮助信息中
		Use: "ifonly-apiserver",
		// 命令的简短描述
		Short: "ifonly",
		// 命令的详细描述
		Long: ``,
		// 命令出错时，不打印帮助信息。设置为 true 可以确保命令出错时一眼就能看到错误信息
		SilenceUsage: true,
		// 指定调用 cmd.Execute() 时，执行的 Run 函数
		RunE: func(cmd *cobra.Command, args []string) error {
			return run(opts)
		},
		// 设置命令运行时的参数检查，不需要指定命令行参数。例如：./miniblog param1 param2
		Args: cobra.NoArgs,
	}

	// 初始化配置函数，在每个命令运行时调用
	// 这里代码是异步的，主要是用来读配置文件的
	cobra.OnInitialize(onInitialize)

	// cobra 支持持久性标志(PersistentFlag)，该标志可用于它所分配的命令以及该命令下的每个子命令
	// 推荐使用配置文件来配置应用，便于管理配置项
	cmd.PersistentFlags().StringVarP(&configFile, "config", "c", filePath(), "Path to the miniblog configuration file.")

	// 将 ServerOptions 中的选项绑定到命令标志
	opts.AddFlags(cmd.PersistentFlags())

	// 添加 --version 标志
	version.AddFlags(cmd.PersistentFlags())

	return cmd
}

func run(opts *options.ServerOptions) error {
	// 如果传入 --version，则打印版本信息并退出
	version.PrintAndExitIfRequested()

	// 初始化日志
	log.Init(logOptions())
	defer log.Sync() // 确保日志在退出时被刷新到磁盘
	// 将 viper 中的配置解析到 opts.
	if err := viper.Unmarshal(opts); err != nil {
		return err
	}

	// 校验命令行选项
	if err := opts.Validate(); err != nil {
		return err
	}

	// 获取应用配置.
	// 将命令行选项和应用配置分开，可以更加灵活的处理 2 种不同类型的配置.
	cfg, err := opts.Config()
	if err != nil {
		return err
	}

	// 创建服务器实例.
	// 注意这里是联合服务器，因为可能同时启动多个不同类型的服务器.
	server, err := cfg.NewUnionServer()
	if err != nil {
		return err
	}

	// 启动服务器
	server.Run()
	return nil
}

// logOptions 从 viper 中读取日志配置，构建 *log.Options 并返回.
// 注意：viper.Get<Type>() 中 key 的名字需要使用 . 分割，以跟 YAML 中保持相同的缩进.
func logOptions() *log.Options {
	opts := log.NewOptions()
	if viper.IsSet("log.disable-caller") {
		opts.DisableCaller = viper.GetBool("log.disable-caller")
	}
	if viper.IsSet("log.disable-stacktrace") {
		opts.DisableStacktrace = viper.GetBool("log.disable-stacktrace")
	}
	if viper.IsSet("log.level") {
		opts.Level = viper.GetString("log.level")
	}
	if viper.IsSet("log.format") {
		opts.Format = viper.GetString("log.format")
	}
	if viper.IsSet("log.output-paths") {
		opts.OutputPaths = viper.GetStringSlice("log.output-paths")
	}
	return opts
}

// 使用cobra就是为了他提供的命令提示，用这个构建后，在系统中就可以使用 ifony-apisever 就像 git和docker这类似的命令构建
// 因为我们用了pflag添加了config配置，但是我们仍然得提供 --help的方法，教别人使用。
// 理论上是可以自己实现，不过有优秀的包就用它，不用自己造轮子。
// 使用pflag知道了用什么配置文件，再使用viper解析配置文件
// 根据配置文件生成gin服务还是grpc服务
