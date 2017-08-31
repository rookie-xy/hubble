package paths

import (
    "fmt"
    "os"
    "sync"
    "path/filepath"

    "github.com/rookie-xy/hubble/src/types"
)

type Path struct {
    Home   string
    Config string
    Data   string
    Logs   string
}

func Home() types.Object {
    home, err := filepath.Abs(filepath.Dir(os.Args[0]))
    if err != nil {
        fmt.Errorf("The absolute path to %s could not be obtained. %v", os.Args[0], err)
        return nil
    }

    return home
}

// FileType is an enumeration type representing the file types.
// Currently existing file types are: Home, Config, Data
type FileType string

const (
    home   FileType = "home"
    Config FileType = "config"
    Data   FileType = "data"
    Logs   FileType = "logs"
)

// Paths is the Path singleton on which the top level functions from this
// package operate.
var Paths = New()


var instance *Path
var once sync.Once
// New creates a new Paths object with all values set to empty values.
func New() *Path {
    once.Do(func() {
        instance = &Path{}
    })

    return instance
}

// InitPaths sets the default paths in the configuration based on CLI flags,
// configuration file and default values. It also tries to create the data
// path with mode 0750 and returns an error on failure.
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

// InitPaths sets the default paths in the configuration based on CLI flags,
// configuration file and default values. It also tries to create the data
// path with mode 0750 and returns an error on failure.
func Init(v types.Value) error {
    return Paths.Init(v)
}

// initPaths sets the default paths in the configuration based on CLI flags,
// configuration file and default values.
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

// Resolve resolves a path to a location in one of the default
// folders. For example, Resolve(Home, "test") returns an absolute
// path for "test" in the home path.
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

// Resolve resolves a path to a location in one of the default
// folders. For example, Resolve(Home, "test") returns an absolute
// path for "test" in the home path.
// In case path is already an absolute path, the path itself is returned.
func Resolve(fileType FileType, path string) string {
    return Paths.Resolve(fileType, path)
}

// String returns a textual representation
func (paths *Path) String() string {
    	return fmt.Sprintf("Home path: [%s] Config path: [%s] Data path: [%s] Logs path: [%s]",
		                       paths.Home, paths.Config, paths.Data, paths.Logs)
}
