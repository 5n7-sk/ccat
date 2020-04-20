package main

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/user"
	"path"
	"strings"

	"github.com/acarl005/stripansi"
	"github.com/alecthomas/chroma"
	"github.com/alecthomas/chroma/formatters"
	"github.com/alecthomas/chroma/lexers"
	"github.com/alecthomas/chroma/styles"
	"github.com/jessevdk/go-flags"
	"github.com/spf13/viper"
)

var usr = func() *user.User {
	usr, err := user.Current()
	if err != nil {
		log.Fatal(err)
	}
	return usr
}()

// These variables are set in build step
var (
	Version = "unset"
)

// Option represents application options
type Option struct {
	Number         bool   `short:"n" long:"number" description:"Show contents with line numbers"`
	NumberNonblank bool   `short:"b" long:"number-nonblank" description:"Show contents with nonempty line numbers"`
	ShowEnds       bool   `short:"E" long:"show-ends" description:"Show $ at end of lines"`
	ShowTabs       bool   `short:"T" long:"show-tabs" description:"Show TAB characters as ^T"`
	Theme          string `short:"t" long:"theme" description:"Overwrite syntax highlighting theme"`
	Version        bool   `short:"v" long:"version" description:"Show ccat version"`
}

// Config represents the settings for this application
type Config struct {
	Theme string `json:"theme"`
}

// CLI represents this application itself
type CLI struct {
	Config Config
}

func download(url string, path string) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	out, err := os.Create(path)
	if err != nil {
		return err
	}
	defer out.Close()

	_, err = io.Copy(out, resp.Body)
	return err
}

// Cat formats file with syntax highlighting
func (c CLI) Cat(opt Option, path string) (string, error) {
	lexer := lexers.Match(path)
	if lexer == nil {
		lexer = lexers.Fallback
	}

	var style *chroma.Style
	if opt.Theme != "" {
		style = styles.Get(opt.Theme)
	} else {
		style = styles.Get((c.Config.Theme))
	}
	if style == nil {
		style = styles.Fallback
	}

	formatter := formatters.Get("terminal256")
	if formatter == nil {
		formatter = formatters.Fallback
	}

	r, err := ioutil.ReadFile(path)
	if err != nil {
		return "", err
	}

	iterator, err := lexer.Tokenise(nil, string(r))
	if err != nil {
		return "", err
	}

	w := new(bytes.Buffer)
	if err := formatter.Format(w, style, iterator); err != nil {
		return "", err
	}

	s := w.String()
	ss := strings.Split(s, "\n")

	i := 0
	contents := ""
	for _, s := range ss {
		if opt.NumberNonblank {
			if stripansi.Strip(s) != "" {
				contents += fmt.Sprintf("%6d  ", i+1)
				i++
			}
		} else if opt.Number {
			contents += fmt.Sprintf("%6d  ", i+1)
			i++
		}

		contents += s

		if opt.ShowEnds {
			contents += "$"
		}

		contents += "\n"
	}

	if opt.ShowTabs {
		contents = strings.Replace(contents, "\t", "^I", -1)
	}

	return contents, nil
}

func run(args []string) int {
	var opt Option
	args, err := flags.ParseArgs(&opt, args)
	if err != nil {
		return 2
	}

	if opt.Version {
		fmt.Printf("ccat v%s\n", Version)
		return 0
	}

	p := path.Join(usr.HomeDir, ".config/ccat.json")

	if _, err = os.Stat(p); os.IsNotExist(err) {
		url := "https://raw.githubusercontent.com/skmatz/ccat/master/ccat.json"
		if err := download(url, p); err != nil {
			fmt.Println(err)
			return 1
		}
	}

	if len(args) == 0 {
		return 2
	}

	viper.SetConfigName("ccat")
	viper.SetConfigType("json")
	viper.AddConfigPath("$HOME/.config")

	err = viper.ReadInConfig()
	if err != nil {
		fmt.Println(err)
		return 1
	}

	var cli CLI
	if err := viper.Unmarshal(&cli); err != nil {
		fmt.Println(err)
		return 1
	}

	for _, arg := range args {
		contents, err := cli.Cat(opt, arg)
		if err != nil {
			return 1
		}
		fmt.Println(contents)
	}

	return 0
}

func main() {
	os.Exit(run(os.Args[1:]))
}
