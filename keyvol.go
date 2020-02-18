package main

import (
  "fmt"
  "strings"
  "os/exec"
  term "github.com/nsf/termbox-go"
)

// Actions

const (
  noop = 0
  exit = 1
  volume_up = 2
  volume_down = 3
  mute = 4
)


func runCommand(cmd string, args []string) {
  _, err := exec.Command(cmd, args...).Output()
  if err != nil {
    panic(err)
  }
}

// Modify these functions that map to enum
// to do change functionality

func volumeUpFunc() {
  args := []string{"set-sink-volume", "@DEFAULT_SINK@", "+2%"}
  runCommand("pactl", args)
}

func volumeDownFunc() {
  args := []string{"set-sink-volume", "@DEFAULT_SINK@", "-2%"}
  runCommand("pactl", args)
}

func volumeMuteFunc() {
  args := []string{"set-sink-mute", "@DEFAULT_SINK@", "toggle"}
  runCommand("pactl", args)
}

func printHelp() {
  fmt.Println(`Instructions:
-------------
* Esc/Ctrl+C/q to exit.
* Up/k to increase volume
* Down/j to decrease volume
* m to toggle mute

`)
}

func isMuted() bool {
  args := []string{"--get-mute"}
  out, _ := exec.Command("pamixer", args...).Output()
  return string(out) == "true\n"
}

// print the current volume (or muted, if its muted)
func printVolume() {
  if !isMuted() {
    vol_args := []string{"--get-volume"}
    vol_out, _ := exec.Command("pamixer", vol_args...).Output()
    fmt.Printf("Volume: %s\n", strings.Trim(string(vol_out), "\n"))
  } else {
    fmt.Println("Volume: (Muted)")
  }
}

func redraw_screen() {
  term.Sync()
  printHelp()
  printVolume()
}

func main() {

  // initialize screen
  err := term.Init()
  if err != nil {
    panic(err)
  }
  defer term.Close()
  redraw_screen()

keyEventLoop:
  // get input from user
  for {
    action := noop
    switch ev := term.PollEvent(); ev.Type {
    case term.EventKey:
      switch ev.Key {
      case term.KeyEsc, term.KeyCtrlC:
        action = exit
      case term.KeyArrowUp:
        action = volume_up
      case term.KeyArrowDown:
        action = volume_down
      default:
        // read single char
        switch ev.Ch {
        case 'q':
          action = exit
        case 'k':
          action = volume_up
        case 'j':
          action = volume_down
        case 'm':
          action = mute
        }
      }
    case term.EventError:
      panic(ev.Err)
    }
    // perform action
    switch action {
    case noop:
    case exit:
      break keyEventLoop
    case volume_up:
        volumeUpFunc()
    case volume_down:
        volumeDownFunc()
    case mute:
        volumeMuteFunc()
    }
    redraw_screen()
  }
}
