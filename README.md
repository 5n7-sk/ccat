<p align="center">
  <a href="https://github.com/skmatz/ccat">
    <img src="./assets/images/banner.png" width="1000" alt="banner" />
  </a>
</p>

<p align="center">
  <a href="https://github.com/skmatz/ccat/actions?query=workflow%3Abuild">
    <img
      src="https://github.com/skmatz/ccat/workflows/build/badge.svg"
      alt="build"
    />
  </a>
  <a href="https://github.com/skmatz/ccat/actions?query=workflow%3Arelease">
    <img
      src="https://github.com/skmatz/ccat/workflows/release/badge.svg"
      alt="release"
    />
  </a>
  <a href="./LICENSE">
    <img
      src="https://img.shields.io/github/license/skmatz/ccat"
      alt="license"
    />
  </a>
  <a href="./go.mod">
    <img
      src="https://img.shields.io/github/go-mod/go-version/skmatz/ccat"
      alt="go version"
    />
  </a>
  <a href="https://github.com/skmatz/ccat/releases/latest">
    <img
      src="https://img.shields.io/github/v/release/skmatz/ccat"
      alt="release"
    />
  </a>
</p>

<p align="center">
  <img src="./assets/images/demo.gif" width="640" alt="demo" />
</p>

# ccat

**ccat** is the colored `cat` command.

## Install

### Binary

Get binary from [releases](https://github.com/skmatz/ccat/releases).

### Source

```sh
go get github.com/skmatz/ccat
```

## Available Themes

We can use various themes.

- adap
- dracula
- emacs
- github
- monokai (default)
- pygments
- solarized-dark
- solarized-light
- vim
- ... [more](https://github.com/alecthomas/chroma/tree/master/styles)

<p align="center">
  <img src="./assets/images/themes.gif" width="640" alt="themes" />
</p>

Set your favorite theme in `~/.config/ccat.json`.  
When you run `ccat` for the first time, the default configuration will be downloaded automatically.

## Supported Languages

Thanks to [alecthomas/chroma](https://github.com/alecthomas/chroma), almost all languages are supported.
