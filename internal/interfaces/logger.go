package interfaces

type LogInterface interface {
	Trace(...string)
	Info(...string)
	Debug(...string)
	Warn(...string)
	Error(...string)
	Fatal(...string)
	Panic(...string)

	// RequestLogger(ctx context.Context, messages map[string]string)

	// PanicRecovery(context.Context, ...string)
}

type MultyLogInterface interface {
	GetLogger() LogInterface
	GetInstituteLogger(institute string) LogInterface
}
