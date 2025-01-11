# hmatch

hmatch is a command line that finds the closest drawbar positions to a recording on the Hammond organ.
It does so by matching the frequencies of the drawbars to the frequency spectrum of the recording. The recording must be provided in wav format.

```
Finds best drawbar registration to a recording

        Example:
        hmatch fit audio.wav

Usage:
  hmatch fit <audiofile> [flags]

Flags:
  -h, --help                 help for fit
      --windowLength int32   Length of the Hann window used when calculating the frequency spectrum (default 1024)
```

## Installation

```sh
go install github.com/davidkleiven/hmatch@latest
```

## Acknowledgements

The program was written after reading the article [The Science of Hammond Organ Drawbar Registration](http://www.stefanv.com/electronics/hammond_drawbar_science.html) by Stefan Vorkoetter.

The standard registatrations available through the `hmatch reg` command are found at [b3world.com](https://b3world.com/hammond-drawbar-settings.html)
