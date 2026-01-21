package gengo

import (
    "strings"

    "google.golang.org/protobuf/compiler/protogen"
)

type trailingComment protogen.Comments

func (c trailingComment) String() string {
    s := strings.TrimSuffix(protogen.Comments(c).String(), "\n")
    if strings.Contains(s, "\n") {
        return ""
    }
    return s
}
