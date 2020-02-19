# gomodinit

select remote repository and do go mod init

this tool uses `git remote -v` and `go mod init`.

```bash
$ gomodinit
? repository:
  ▸ github.com/nu50218/gomodinit
```

→

```bash
$ gomodinit
✔ github.com/nu50218/gomodinit
--
$ go mod init github.com/nu50218/gomodinit
go: creating new go.mod: module github.com/nu50218/gomodinit
--
```

## Install

`$ go get -u github.com/nu50218/gomodinit`

## Usage

`$ gomodinit`
