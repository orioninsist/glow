package main

import (
	"errors"
	"fmt"
	"io/fs"
	"os"
	"path"
	"path/filepath"

	"github.com/charmbracelet/x/editor"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

const defaultConfig = `# style name or JSON path (default "auto")
style: "auto"
# use pager to display markdown
pager: false
# use the interactive TUI instead of plain CLI output
tui: false
# word-wrap at width; set 0 to auto-detect the terminal width
width: 0
# mouse support (TUI-mode only)
mouse: false
# show all files, including hidden and ignored (TUI-mode only)
all: false
# show line numbers (TUI-mode only)
showLineNumbers: false
# preserve authored newlines while rendering
preserveNewLines: false
# render markdown image lines inline with kitten icat in Kitty (CLI output only)
kittyImages: false
`

var configCmd = &cobra.Command{
	Use:     "config",
	Hidden:  false,
	Short:   "Edit the glow config file",
	Long:    paragraph(fmt.Sprintf("\n%s the glow config file. We’ll use EDITOR to determine which editor to use. If the config file doesn't exist, it will be created.", keyword("Edit"))),
	Example: paragraph("glow config\nglow config --config path/to/config.yml"),
	Args:    cobra.NoArgs,
	RunE: func(*cobra.Command, []string) error {
		if err := ensureConfigFile(); err != nil {
			return err
		}

		c, err := editor.Cmd("Glow", configFile)
		if err != nil {
			return fmt.Errorf("unable to set config file: %w", err)
		}
		c.Stdin = os.Stdin
		c.Stdout = os.Stdout
		c.Stderr = os.Stderr
		if err := c.Run(); err != nil {
			return fmt.Errorf("unable to run command: %w", err)
		}

		fmt.Println("Wrote config file to:", configFile)
		return nil
	},
}

func ensureConfigFile() error {
	if configFile == "" {
		configFile = viper.GetViper().ConfigFileUsed()
		if err := os.MkdirAll(filepath.Dir(configFile), 0o755); err != nil { //nolint:gosec
			return fmt.Errorf("could not write configuration file: %w", err)
		}
	}

	if ext := path.Ext(configFile); ext != ".yaml" && ext != ".yml" {
		return fmt.Errorf("'%s' is not a supported configuration type: use '%s' or '%s'", ext, ".yaml", ".yml")
	}

	if _, err := os.Stat(configFile); errors.Is(err, fs.ErrNotExist) {
		// File doesn't exist yet, create all necessary directories and
		// write the default config file
		if err := os.MkdirAll(filepath.Dir(configFile), 0o700); err != nil {
			return fmt.Errorf("unable create directory: %w", err)
		}

		f, err := os.Create(configFile)
		if err != nil {
			return fmt.Errorf("unable to create config file: %w", err)
		}
		defer func() { _ = f.Close() }()

		if _, err := f.WriteString(defaultConfig); err != nil {
			return fmt.Errorf("unable to write config file: %w", err)
		}
	} else if err != nil { // some other error occurred
		return fmt.Errorf("unable to stat config file: %w", err)
	}
	return nil
}
