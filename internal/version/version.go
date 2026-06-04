package version

import (
    "fmt"
)

const (
    Major      = 2
    Minor      = 0
    Patch      = 8
    PreRelease = ""
)

// String 以 semver 格式化此模块的版本字符串。
//
// 示例：
//
//	v1.20.1
//	v1.21.0-rc.1
func String() string {
    v := fmt.Sprintf("v%d.%d.%d", Major, Minor, Patch)
    if PreRelease != "" {
        v += "-" + PreRelease
    }
    return v
}
