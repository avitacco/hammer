package module

type Severity int

const (
	Info Severity = iota
	Warning
	Error
)

func (s Severity) String() string {
	switch s {
	case Info:
		return "info"
	case Warning:
		return "warning"
	case Error:
		return "error"
	default:
		return "unknown"
	}
}

type ValidationResult struct {
	Level   Severity
	Field   string
	Message string
}

func (m Metadata) Validate() []ValidationResult {
	return nil
}
