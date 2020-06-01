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
If you already have [jq](https://github.com/stedolan/jq) and [fzf](https://github.com/junegunn/fzf) or [peco](https://github.com/peco/peco), you can download binary by running the following command.

```sh
curl -Ls https://api.github.com/repos/skmatz/ccat/releases/latest | jq -r ".assets[].browser_download_url" | fzf | wget -i -
```

### Source

```sh
go get github.com/skmatz/ccat
```

## Commands

```console
> ccat --help

Usage:
  ccat [OPTIONS]

Application Options:
  -n, --number           Show contents with line numbers
  -b, --number-nonblank  Show contents with nonempty line numbers
  -E, --show-ends        Show $ at end of lines
  -T, --show-tabs        Show TAB characters as ^T
  -t, --theme=           Overwrite syntax highlighting theme
  -v, --version          Show ccat version

Help Options:
  -h, --help             Show this help message
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
