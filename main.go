package main

import (
        "fmt"
        "log"
        "os"

        "github.com/shogo82148/androidbinary/apk"
)

func main() {
        if len(os.Args) < 2 {
                log.Fatalf("Usage: %s <path-to-apk>", os.Args[0])
        }

        apkPath := os.Args[1]

        reader, err := apk.OpenFile(apkPath)
        if err != nil {
                log.Fatalf("Failed to open APK: %v", err)
        }
        defer reader.Close()

        manifest := reader.Manifest()

        fmt.Println("Package Name:", manifest.Package.MustString())
        fmt.Println("Version Code:", manifest.VersionCode.MustInt32())
        fmt.Println("Version Name:", manifest.VersionName.MustString())

}
