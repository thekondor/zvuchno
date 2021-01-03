# 🎵 Zvuchno

Yet another simple PulseAudio volume notification service over `libnotify`. The service provides notification for:

- Volume up/down
- Volume mute/unmute

![In action](zvuchno-demo.gif)

Change of volume's level is visualized using a textual progress bar. No graphics, no other extra complex use-cases is an intention.
The text **is customizable**: ASCII and Emojis.

The app was written as tiny entertainment to close a gap in my existing [i3wm](https://i3wm.org/) setup. My alternative to `volnoti` daemon.

## Build & Configuration

Build:
- `$ git clone https://git.thekondor.net/zvuchno.git` or (`$ git clone https://github.com/thekondor/zvuchno`)
- `$ cd zvuchno && go build`

Run w/o any configuration file to use default settings. Or use `config.sample.yml` as a foundation for personal one.
The configuration file could be stored either as `$HOME/.zvuchno.yml` or `$XDG_CONFIG_HOME/zvuchno.yml`.

The application heavily relies on running PulseAudio daemon as well as available DBus session.

### Format

The final volume's level representation is defined via `appearance.format.full` key. The value's format is simply a Go's template.
`{{ .Percent }}` is expanded to current volume's level in percent.
`{{ .Bar}}` is expanded to a textual representation of current volume's level.

Bar's representation is defined via `appearance.format.bar` key. The value is 5-character string. Example: `[=>-]`, where

1. `[` is a left border of the bar.
2. `=` is a floating part of the bar.
3. `>` is a pointer of the floating part showing current volume's level.
4. `-` is a remaining space for the bar to float to.
5. `]` is a right border of the bar.

Any character here could be a regular ASCII one as well as a fancy Emoji (should be supported by active font)

## Usage

As simple as:

```shell
$ zvuchno
```

or to be added to `~/.xinitrc` (or `i3wm/config`... whatever!) for a launch on WM startup:

```
# Mind '&' at the end!
/path/to/zvuchno &
```

