# keyvol

An interactive terminal interface for `pamixer`/`pactl` to control system volume.

<img src="https://raw.githubusercontent.com/seanbreckenridge/keyvol/master/.github/demo.gif" alt="keyvol demo gif">

Run: `keyvol`

```
Instructions:
-------------
* Esc/Ctrl+C/q to exit.
* Up/k to increase volume
* Down/j to decrease volume
* m to toggle mute
```

This could be repurposed to wrap another command by changing the command/args [here](https://github.com/seanbreckenridge/keyvol/blob/master/keyvol.go#L29-L59).

#### Install

Dependencies: `go`, `pamixer`, `pactl`

```
go get -v -u "github.com/seanbreckenridge/keyvol"
```
