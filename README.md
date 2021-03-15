# rasterm
Encodes images to iTerm / Kitty / SIXEL (terminal) inline graphics protocols.

[![GoDoc](https://godoc.org/github.com/BourgeoisBear/rasterm?status.png)](http://godoc.org/github.com/BourgeoisBear/rasterm)

## Supported Image Encodings
- *WezTerm & iTerm2*: https://iterm2.com/documentation-images.html
- *Kitty*: https://sw.kovidgoyal.net/kitty/graphics-protocol.html
- *Sixel*: https://saitoha.github.io/libsixel/

## TODO
- godoc link
- improve terminal identification
	4:sixel graphics
	ESC[0c = 63;1;2;4;6;9;15;22c
	19:VT340
	ESC[>0c = 19;344:0c
	https://invisible-island.net/xterm/ctlseqs/ctlseqs-contents.html

- check that mintty supports iterm/wezterm format, get mintty identifier
- unit tests
- conditionally enable/disable tmux passthrough

perhaps query tmux directly
TMUX=/tmp/tmux-1000/default,3218,4

## TESTING
- test sixel with
	- https://github.com/liamg/aminal
	- https://domterm.org/
	- https://www.macterm.net/
- test wez/iterm img with
	- https://www.macterm.net/
  - mintty
