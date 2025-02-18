package interfaces

type ConfigInterface interface {
	GetConfig(value any) (err error)
	GetConfigByteData() []byte
}
