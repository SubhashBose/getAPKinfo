
# 📦 APK Info Extractor

A simple Go command-line tool to extract metadata from Android APK files.

It supports extracting:

- App Name
- Package Name
- Version Code
- Version Name

Supports plain text and JSON output, with flags to control which fields are displayed.

---

## 🔧 Download

Go to release of this repo and download binary for relvant OS and Arch.

---

## 🚀 Usage

```bash
./APKinfo [flags] <path-to-apk>
```

The `<path-to-apk>` can be placed anywhere in the command-line arguments.

---

## 🏷️ Flags

| Flag             | Description                      |
| ---------------- | -------------------------------- |
| `--app-name`     | Show app name                    |
| `--package`      | Show package name                |
| `--version-code` | Show version code                |
| `--version-name` | Show version name                |
| `--no-labels`    | Print values only (no labels)    |
| `--json`         | Output all values in JSON format |

If no output flags are provided (`--app-name`, `--package`, `--version-code`, or `--version-name`), all fields are shown by default.

---

## 📦 Examples

### Show All Fields (Default)

```bash
./APKinfo Telegram.apk
```

```
App Name     : Telegram
Package Name : org.telegram.messenger.web
Version Code : 59309
Version Name : 11.11.1
```

---

### Show Specific Fields

```bash
./APKinfo --version-code --version-name Telegram.apk
```

```
Version Code : 59309
Version Name : 11.11.1
```

---

### Show Value Only (No Labels)

```bash
./APKinfo --version-name --no-labels Telegram.apk
```

```
11.11.1
```

---

### JSON Output

```bash
./APKinfo --json Telegram.apk
```

```json
{
  "app_name": "Telegram",
  "package_name": "org.telegram.messenger.web",
  "version_code": 59309,
  "version_name": "11.11.1"
}
```

