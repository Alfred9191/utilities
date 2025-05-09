package main

// Code generated by mkparamfilefunc; DO NOT EDIT.
import (
	"path/filepath"

	"github.com/nickwells/filecheck.mod/filecheck"
	"github.com/nickwells/param.mod/v6/param"
	"github.com/nickwells/xdg.mod/xdg"
)

/*
SetConfigFile adds a config file to the set which the param parser will process
before checking the command line parameters.

This function is one of a pair which add the global and personal config files.
It is generally best practice to add the global config file before adding the
personal one. This allows any system-wide defaults to be overridden by personal
choices. Also any parameters which can only be set once can be set in the global
config file, thereby enforcing a global policy.
*/
func SetConfigFile(ps *param.PSet) error {
	baseDir := xdg.ConfigHome()

	ps.AddConfigFileStrict(
		filepath.Join(baseDir,
			"github.com",
			"nickwells",
			"utilities",
			"mkdoc",
			"common.cfg"),
		filecheck.Optional)

	return nil
}

/*
SetGlobalConfigFile adds a config file to the set which the param parser will
process before checking the command line parameters.

This function is one of a pair which add the global and personal config files.
It is generally best practice to add the global config file before adding the
personal one. This allows any system-wide defaults to be overridden by personal
choices. Also any parameters which can only be set once can be set in the global
config file, thereby enforcing a global policy.
*/
func SetGlobalConfigFile(ps *param.PSet) error {
	dirs := xdg.ConfigDirs()
	if len(dirs) == 0 {
		return nil
	}

	baseDir := dirs[0]

	ps.AddConfigFileStrict(
		filepath.Join(baseDir,
			"github.com",
			"nickwells",
			"utilities",
			"mkdoc",
			"common.cfg"),
		filecheck.Optional)

	return nil
}
