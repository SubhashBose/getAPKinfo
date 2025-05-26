package main

import (
        "encoding/json"
        "flag"
        "fmt"
        "log"
        "os"

        "github.com/shogo82148/androidbinary/apk"
)

func main() {
        // Flags
        showAppName := flag.Bool("app-name", false, "Show app name")
        showPackage := flag.Bool("package", false, "Show package name")
        showVersionCode := flag.Bool("version-code", false, "Show version code")
        showVersionName := flag.Bool("version-name", false, "Show version name")
        noLabels := flag.Bool("no-labels", false, "Print values only, without labels")
        jsonOutput := flag.Bool("json", false, "Output as JSON")

		// Rearrange os.Args: move non-flag arguments to the end
		rearranged := []string{os.Args[0]} // keep program name
			flagsFound := []string{}
			nonFlags := []string{}

			for _, arg := range os.Args[1:] {
				if len(arg) > 0 && arg[0] == '-' {
					flagsFound = append(flagsFound, arg)
				} else {
					nonFlags = append(nonFlags, arg)
				}
			}
			rearranged = append(rearranged, flagsFound...)
			rearranged = append(rearranged, nonFlags...)
			os.Args = rearranged

        // Parse flags and get APK path
        flag.Parse()
        args := flag.Args()

        if len(args) < 1 {
                log.Fatalf("Usage: %s [flags] <path-to-apk>", os.Args[0])
        }
        apkPath := args[0]

        // Open APK
        reader, err := apk.OpenFile(apkPath)
        if err != nil {
                log.Fatalf("Failed to open APK: %v", err)
        }
        defer reader.Close()

        manifest := reader.Manifest()

        // Extract values
        label, err := reader.Label(nil)
        if err != nil {
                log.Printf("Failed to get app name: %v", err)
                label = "(unknown)"
        }
        packageName := manifest.Package.MustString()
        versionCode := manifest.VersionCode.MustInt32()
        versionName := manifest.VersionName.MustString()

        // Determine if no output flags were set
        outputAll := !*showAppName && !*showPackage && !*showVersionCode && !*showVersionName

        // JSON output
        if *jsonOutput {
                data := map[string]interface{}{
                        "app_name":     label,
                        "package_name": packageName,
                        "version_code": versionCode,
                        "version_name": versionName,
                }
                enc := json.NewEncoder(os.Stdout)
                enc.SetIndent("", "  ")
                if err := enc.Encode(data); err != nil {
                        log.Fatalf("Failed to encode JSON: %v", err)
                }
                return
        }

        // Text output
        if *showAppName || outputAll {
                printLine("App Name", label, *noLabels)
        }
        if *showPackage || outputAll {
                printLine("Package Name", packageName, *noLabels)
        }
        if *showVersionCode || outputAll {
                printLine("Version Code", versionCode, *noLabels)
        }
        if *showVersionName || outputAll {
                printLine("Version Name", versionName, *noLabels)
        }
}

func printLine(label string, value interface{}, noLabel bool) {
        if noLabel {
                fmt.Println(value)
        } else {
                fmt.Printf("%-13s: %v\n", label, value)
        }
}
