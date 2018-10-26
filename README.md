# binary-clock

simple binary clock implemented in Go, using `termbox` as the interface.

Here's what it looks like at 3:52:36 pm:
```
what's the current time?
  ┌─┐     ┌─┐     ┌─┐
  │ │     │ │     │ │
  ├─┤   ┌─┼─┤   ┌─┼─┤
  │x│   │x│ │   │ │x│
┌─┼─┤   ├─┼─┤   ├─┼─┤
│ │ │   │ │x│   │x│x│
├─┼─┤   ├─┼─┤   ├─┼─┤
│x│x│   │x│ │   │x│ │
└─┴─┘   └─┴─┘   └─┴─┘
```
