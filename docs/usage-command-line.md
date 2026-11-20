---
layout: page
title: Command line
hide_hero: true
show_sidebar: false
menubar: docs-menu
---

# Command line

Metro2 has a command line interface to manage Metro 2 files and launch a web service.

## Install

Download a release binary from the [latest GitHub release](https://github.com/moov-io/metro2/releases/latest). Assets are named `metro2-linux-amd64`, `metro2-darwin-amd64`, and `metro2.exe` (Windows). For example, on Linux:

```
curl -sSL -o metro2 "https://github.com/moov-io/metro2/releases/download/v0.10.0/metro2-linux-amd64"
chmod +x metro2
```

Or install with Go (requires [Go v1.26 or newer](https://golang.org/doc/install)):

```
go install github.com/moov-io/metro2/cmd/metro2@latest
```

From a source checkout you can also build with `make build`, which writes `./bin/metro2`.

## Help

```
metro2 --help

Moov Metro 2 file reader, writer, and validator

Usage:
  metro2 [command]

Available Commands:
  completion  Generate the autocompletion script for the specified shell
  convert     Convert metro file format
  help        Help about any command
  print       Print metro file
  validator   Validate metro file
  version     Print the version of metro2
  web         Launches web server

Flags:
  -h, --help           help for metro2
      --input string   input file (default is $PWD/metro.json)

Use "metro2 [command] --help" for more information about a command.
```

```
metro2 version
v0.10.0
```

### Exit codes

| Code | Meaning |
| ---- | ------- |
| `0` | The command succeeded. |
| `1` | The command failed. This includes a missing or unreadable `--input` file, a validation error, an unsupported `--format`, or a convert output path whose parent directory does not exist. |

Each interaction that the library supports is exposed in a command-line option:

 Command | Info
 ------- | -------
`convert` | The convert command allows users to convert a metro file to a specified file format (json, metro). The result will create a new file.
`print` | The print command allows users to print a metro file in a specified file format (json, metro).
`validator` | The validator command allows users to validate a metro file.
`version` | Print the library version (from `metro2.Version`).
`web` | The web command will launch a web server with endpoints to manage metro files.
`completion` | Generate a shell autocompletion script.

### File formats

Metro 2 files are either **character** (unpacked) or **packed**:

- **Character** format uses 426-byte header, base, and trailer records encoded as ASCII. Fixed character files omit the block descriptor word and start with a 4-digit record descriptor word (for example `0426HEADER`). Variable-blocked character files include both a 4-digit block descriptor word and a record descriptor word (for example `04960426HEADER` in `test/testdata/unpacked_variable_file.dat`). Character format is preferred for electronic transmission.
- **Packed** format uses 366-byte records. The block and record descriptor words are 4-byte binary values (a big-endian uint16 in the first two bytes, then two zero bytes), so a packed Metro 2 file is not plain text. `test/testdata/packed_file.json` is a packed file (`"blockDescriptorWord": 370`, `"recordDescriptorWord": 366`).

`convert` and `print` accept `--format json` or `--format metro`. `--format metro` writes the raw Metro 2 records. On `convert`, `--type character` or `--type packed` changes the file type before writing. `--type character` recasts packed records to character encoding but keeps the original descriptor-word values, so converting `packed_file.json` with `--type character` starts with `03700366HEADER` rather than `0426HEADER`. Use `-n` / `--newline` to insert a newline after each record, and `-g` / `--generate` to replace the trailer with one computed from the data records (`GeneratorTrailer`).

Follow each credit bureau's Metro 2 specification for whether they expect packed or character format and fixed or variable blocking.

### file convert

```
metro2 convert --help

Usage:
  metro2 convert [output] [flags]

Flags:
      --format string   format of metro file(required) (default "json")
  -g, --generate        generate trailer record
  -h, --help            help for convert
  -n, --newline         has new line
      --type string     file type (character or packed)

Global Flags:
      --input string   input file (default is $PWD/metro.json)
```

- The `output` parameter represents the full path name for the new metro2 file. The parent directory must already exist.
- The `format` parameter determines the output file format and supports "json" or "metro".
- The `generate` parameter will create a newly generated trailer record in the file.
- The `type` parameter optionally sets the file type to `character` or `packed`.
- The `newline` parameter inserts a newline after each record when writing Metro format.
- The `input` parameter is the source metro2 file to be converted, and can be raw or json.

Examples (run from the repository root):

```
mkdir -p output
metro2 convert output/output.json --input test/testdata/packed_file.json --format json
```

Converting packed JSON to Metro format writes binary block descriptor words. The first 8 bytes of the result are `01 72 00 00 01 6e 00 00` (BDW 370 and RDW 366), followed by `HEADER`:

```
metro2 convert output/packed.metro --input test/testdata/packed_file.json --format metro
```

To emit character-encoded records and separate them with newlines:

```
metro2 convert output/character.metro --input test/testdata/packed_file.json --format metro --type character -n
```

The resulting file starts with `03700366HEADER`.

### file print

```
metro2 print --help

Usage:
  metro2 print [flags]

Flags:
      --format string   print format (default "json")
  -h, --help            help for print
  -n, --newline         has new line

Global Flags:
      --input string   input file (default is $PWD/metro.json)
```

- The `format` parameter determines the output format and supports "json" or "metro".
- The `newline` parameter inserts a newline after each record when printing Metro format.
- The `input` parameter is the source metro2 file to be printed, and can be raw or json.

Example:

```
metro2 print --input test/testdata/packed_file.dat --format json
{
  "header": {
    "blockDescriptorWord": 370,
    "recordDescriptorWord": 366,
    "recordIdentifier": "HEADER",
    "transUnionProgramIdentifier": "5555555555",
    "activityDate": "2002-08-20T00:00:00Z",
    "dateCreated": "1999-05-10T00:00:00Z",
    "programDate": "1999-05-10T00:00:00Z",
    "programRevisionDate": "1999-05-10T00:00:00Z",
    "reporterName": "YOUR BUSINESS NAME HERE",
    "reporterAddress": "LINE ONE OF YOUR ADDRESS LINE TWO OF YOUR ADDRESS LINE THERE OF YOUR ADDRESS",
    "reporterTelephoneNumber": 1234567890
  },
  ...
}
```

### file validate

```
metro2 validator --help

Usage:
  metro2 validator [flags]

Flags:
  -h, --help   help for validator

Global Flags:
      --input string   input file (default is $PWD/metro.json)
```

- The `input` parameter is the source metro2 file to be validated, and can be raw or json.

Example:

```
metro2 validator --input test/testdata/packed_file.dat
the file is valid
```

### Web server

```
metro2 web --help

Usage:
  metro2 web [flags]

Flags:
  -h, --help          help for web
      --port string   port of the web server (default "8080")
  -t, --test          test server

Global Flags:
      --input string   input file (default is $PWD/metro.json)
```

- The `port` parameter is the port number for the web service.

Example:
```
metro2 web
```

The web server has some endpoints to manage metro2 files:

Method | Endpoint | Content-Type | Info
 ------- | ------- | ------- | -------
 `POST` | `/convert` | multipart/form-data | Convert metro file, will download new file.
 `GET` | `/health` | application/json | Check web server status. Returns `{"status":"alive"}`.
 `POST` | `/print` | multipart/form-data | Print metro file. JSON responses use `Content-Type: application/json`.
 `POST` | `/validator` | multipart/form-data | Validate metro file. A valid file returns `{"status":"valid file"}`.

Upload the file in a form field named `file`. Requests that omit that field return HTTP 400 with `missing form file field "file"`.

Web page example of the metro2 web server:

```
<!doctype html>
<html lang="en">
<head>
    <meta charset="utf-8">
    <title>Single file upload</title>
</head>
<body>
<h1>Upload single file with fields</h1>

<form action="http://localhost:8080/convert" method="post" enctype="multipart/form-data">
    Format: <input type="text" name="format"><br>
    Files: <input type="file" name="file"><br><br>
    <input type="submit" value="Submit">
</form>
</body>
</html>
```
