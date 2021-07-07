---
layout: page
title: Command line
hide_hero: true
show_sidebar: false
menubar: docs-menu
---

# Command line

Metro2 has a command line interface to manage Metro 2 files and launch a web service.

```
metro2 --help

Usage:
   [command]

Available Commands:
  convert     Convert metro file format
  help        Help about any command
  print       Print metro file
  validator   Validate metro file
  web         Launches web server

Flags:
  -h, --help           help for this command
      --input string   input file (default is $PWD/metro.json)

Use " [command] --help" for more information about a command.
```

Each interaction that the library supports is exposed in a command-line option:

 Command | Info
 ------- | -------
`convert` | The convert command allows users to convert a metro file to a specified file format (json, metro). The result will create a new file.
`print` | The print command allows users to print a metro file in a specified file format (json, metro).
`validator` | The validator command allows users to validate a metro file.
`web` | The web command will launch a web server with endpoints to manage metro files.

### file convert

```
metro2 convert --help

Usage:
   convert [output] [flags]

Flags:
      --format string   format of metro file(required) (default "json")
  -g, --generate        generate trailer record
  -h, --help            help for convert

Global Flags:
      --input string   input file (default is $PWD/metro.json)
```

- The `output` parameter represents the full path name for the new metro2 file.
- The `format` parameter determines the output file format and supports "json" or "metro".
- The `generate` parameter will create a newly generated trailer record in the file.
- The `input` parameter is the source metro2 file to be converted, and can be raw or json.

Example:
```
metro2 convert output/output.json --input testdata/packed_file.json --format json
```

### file print

```
metro2 print --help

Usage:
   print [flags]

Flags:
      --format string   print format (default "json")
  -h, --help            help for print

Global Flags:
      --input string   input file (default is $PWD/metro.json)
```

- The `format` parameter determines the output format and supports "json" or "metro".
- The `input` parameter is the source metro2 file to be printed, and can be raw or json.

Example:
```
metro2 print --input testdata/packed_file.dat --format json
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
   validator [flags]

Flags:
  -h, --help   help for validator

Global Flags:
      --input string   input file (default is $PWD/metro.json)
```

- The `input` parameter is the source metro2 file to be validated, and can be raw or json.

Example:
```
metro2 validator --input testdata/packed_file.dat
Error: is an invalid value of TotalConsumerSegmentsJ1
```

### Web server

```
metro2 web --help

Usage:
   web [flags]

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
 `GET` | `/health` | text/plain | Check web server status.
 `POST` | `/print` | multipart/form-data | Print metro file.
 `POST` | `/validator` | multipart/form-data | Validate metro file.

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
