# Note Tool

Personal note taking app for personal use with different use cases.
Mainly built to play with Golang and automate some daily note taking stuff.

You can create quick notes, TODO notes, meeting notes, weekly meeting notes,
parse and organize lastly created notes by their tags, etc.

## TODOs

- [x] Template for quick notes
  [Docs](https://pkg.go.dev/text/template#example-Template-Block)
- [x] Tags
- [ ] Links
- [ ] Meeting notes
- [x] TODOs notes
- [ ] Organizer
- [ ] Finder
- [ ] Find by tags
  - for this functionality, find all files with those tags and move them to a temporary folder
  - then open the files with tea, and once selected open the file with that name from the original folder
  - finally, delete the temporary folder
- [ ] Add viper so user can configure custom templates

## Usage

```bash
Usage:
  nt [flags]
  nt [command]

Available Commands:
  completion  Generate the autocompletion script for the specified shell
  help        Help about any command
  todo        Create quick note for TODOs

Flags:
  -h, --help               help for nt
  -t, --tags stringArray   Tags for the new note
  -v, --verbose            Enable verbose mode

Use "nt [command] --help" for more information about a command.

```
