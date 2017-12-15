package paths

type FileType string

const (
    home   FileType = "home"
    Config FileType = "config"
    Data   FileType = "data"
    Logs   FileType = "logs"
)
