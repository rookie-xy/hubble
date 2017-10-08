package paths

import (
    "fmt"
    "os"
    "sync"
    "path/filepath"

    "github.com/rookie-xy/hubble/types"
)

type Path struct {
    Home   string
    Config string
    Data   string
    Logs   string
}

var instance *Path
var once sync.Once

func GetInstance() *Path {
    once.Do(func() {
        instance = &Path{}
    })

    return instance
}

func (paths *Path) Init(v types.Value) error {
    err := paths.init(v)
    if err != nil {
        return err
    }

    // make sure the data path exists
    err = os.MkdirAll(paths.Data, 0750)
    if err != nil {
        return fmt.Errorf("Failed to create data path %s: %v", paths.Data, err)
    }

    return nil
}

func Init(v types.Value) error {
    return GetInstance().Init(v)
}

func (paths *Path) init(v types.Value) error {
    paths.Home = v.GetString()

    // default for config path
    if paths.Config == "" {
        paths.Config = paths.Home
    }

    // default for data path
    if paths.Data == "" {
        paths.Data = filepath.Join(paths.Home, "data")
    }

    // default for logs path
    if paths.Logs == "" {
        paths.Logs = filepath.Join(paths.Home, "logs")
    }

    return nil
}

func (paths *Path) Resolve(fileType FileType, path string) string {
    // absolute paths are not changed
    if filepath.IsAbs(path) {
        return path
    }

    switch fileType {

    case home:
        return filepath.Join(paths.Home, path)
    case Config:
        return filepath.Join(paths.Config, path)
   	case Data:
        return filepath.Join(paths.Data, path)
   	case Logs:
        return filepath.Join(paths.Logs, path)
   	default:
        panic(fmt.Sprintf("Unknown file type: %s", fileType))
    }
}

func Resolve(fileType FileType, path string) string {
    return GetInstance().Resolve(fileType, path)
}

func (paths *Path) String() string {
    	return fmt.Sprintf("Home path: [%s] Config path: [%s] Data path: [%s] Logs path: [%s]",
                            paths.Home, paths.Config, paths.Data, paths.Logs)
}

func Home() types.Object {
    home, err := filepath.Abs(filepath.Dir(os.Args[0]))
    if err != nil {
        fmt.Errorf("The absolute path to %s could not be obtained. %v", os.Args[0], err)
        return nil
    }

    return home
}

