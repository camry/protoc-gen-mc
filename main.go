package main

import (
    "flag"
    "fmt"
    "os"
    "path/filepath"

    "google.golang.org/protobuf/compiler/protogen"

    "protoc-gen-mc/internal/gengo"
    "protoc-gen-mc/internal/version"
)

const genMcDocURL = "https://github.com/camry/proto-gen-mc"

func main() {
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
