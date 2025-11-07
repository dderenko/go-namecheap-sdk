# Logger

Logger is a simple wrapper of `zap.Logger` with addition of having `Logs` and `LogsSkipped` metrics.
Logger adds following fields to the log messages automatically:

If encoding is `console`: logger uses `zap.DevelopmentConfig`. Sampling is disabled.

In prod sampling is always **ON**. If you set sampling: **initial: 0, thereafter: 0** - it **disables** logging at all.

```go
field.Service(env.Service()),
field.Env(env.Env()),
field.Version(env.ReleaseVersion()),
field.Tags(config.Tags),
```

You can configure:
- minimum enabled log level via `config.Level`,
- encoding type via `config.Encoding`
- and set tags for every log message via `config.Tags`
- also you can sample logs by setting `config.Sampling`
- annotate logs with caller info with `config.EnableCaller`

### Sampling configuration

Next config

```yaml
sampling:
    initial: 10
    thereafter: 5 
```

means that logger will log **10 first log messages** with same msg and level and then will log only **5 messages** with same msg and level **per second**.

In prod sampling is always **ON**. If you set sampling: **initial: 0, thereafter: 0** - it **disables** logging at all.

## Documentation

To see some doc just run the following command:

```bash
godoc -http=:6060
```

Then open your browser and go to the following URL: [http://localhost:6060/pkg/github.com/propellerads/logger/](http://localhost:6060/pkg/github.com/propellerads/logger/)
