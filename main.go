package main

import (
    "flag"
    "fmt"
    "os"
    "path/filepath"

    "github.com/dromara/carbon/v2"
    "google.golang.org/protobuf/compiler/protogen"

    "gitea.meigushi.go/meigushi/protoc-gen-mc/internal/gengo"
    "gitea.meigushi.go/meigushi/protoc-gen-mc/internal/version"
)

const genMcDocURL = "https://github.com/camry/proto-gen-mc"

func main() {
    if carbon.Now().Gt(carbon.CreateFromDateTime(2027, 1, 1, 0, 0, 0)) {
        fmt.Fprintf(os.Stderr, "Please update using the latest release from %v.\n", genMcDocURL)
        os.Exit(1)
    }
    if len(os.Args) == 2 && os.Args[1] == "--version" {
        fmt.Fprintf(os.Stdout, "%v %v\n", filepath.Base(os.Args[0]), version.String())
        os.Exit(0)
    }
    if len(os.Args) == 2 && os.Args[1] == "--help" {
        fmt.Fprintf(os.Stdout, "See "+genMcDocURL+" for usage information.\n")
        os.Exit(0)
    }

    var (
        flags flag.FlagSet
    )
    protogen.Options{
        ParamFunc: flags.Set,
    }.Run(func(gen *protogen.Plugin) error {
        for _, f := range gen.Files {
            if f.Generate {
                gengo.GenerateFile(gen, f)
            }
        }
        return nil
    })
}
