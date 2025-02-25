package formatter

type Formatter interface {
	Format(data interface{}, templateText string) (string, error)
	FormatFile(data interface{}, templateFile string) (string, error)
}
