package logger

type Param struct {
	Level          string
	Path           string
	MaxSize        int
	MaxBackupNum   int
	MackupDuration int
}

func NewParam(level string, path string, maxSize int, maxBackups int, maxAge int) *Param {
	return &Param{
		Level:          level,
		Path:           path,
		MaxSize:        maxSize,
		MaxBackupNum:   maxBackups,
		MackupDuration: maxBackups,
	}
}
