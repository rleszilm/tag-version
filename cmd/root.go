/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"os"

	"github.com/rleszilm/tag-version/internal/version"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "tag-version",
	Short: "tool that returns a semver based off the repos history.",
	Long:  `tag-version inspects the repository and uses the most recent tag to return a semantic version in the format "v<major>.<minor>.<patch>".`,
	Run:   root,
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().BoolP("branch", "b", false, "When set the branch will be included in the version.")
	rootCmd.PersistentFlags().BoolP("increment", "i", false, "When set the incremented version is returned.")
	rootCmd.PersistentFlags().BoolP("full", "f", false, "When set the full semver version is returned. IE v1.2.0 instead of v1.2 for a minor version.")
	rootCmd.PersistentFlags().BoolP("revision", "r", false, "When set the revision will be included in the version.")
	rootCmd.PersistentFlags().BoolP("semver", "s", false, "When set the semver will be included in the version regardless of whether the version is for the master branch.")
	rootCmd.PersistentFlags().StringP("master", "m", "master", "Overrides the \"master\" branch of the repo.")
}

func root(cmd *cobra.Command, args []string) {
	patchCmd.Run(cmd, args)
}

func buildVersion(cmd *cobra.Command, args []string) (*version.Version, error) {
	dir := "."
	if len(args) > 0 {
		dir = args[0]
	}

	git, err := version.NewGit(dir)
	if err != nil {
		return nil, fmt.Errorf("could not open repo: %w", err)

	}

	vopts := &version.VersionOption{}

	br, err := cmd.Flags().GetBool("branch")
	if err != nil {
		return nil, fmt.Errorf("invalid branch: %w", err)

	}
	vopts.SetBranch(br)

	f, err := cmd.Flags().GetBool("full")
	if err != nil {
		return nil, fmt.Errorf("invalid full: %w", err)

	}
	vopts.SetFull(f)

	rev, err := cmd.Flags().GetBool("revision")
	if err != nil {
		return nil, fmt.Errorf("invalid revision: %w", err)

	}
	vopts.SetRevision(rev)

	sem, err := cmd.Flags().GetBool("semver")
	if err != nil {
		return nil, fmt.Errorf("invalid semver: %w", err)

	}
	vopts.SetSemver(sem)

	m, err := cmd.Flags().GetString("master")
	if err != nil {
		return nil, fmt.Errorf("invalid master: %w", err)

	}
	vopts.SetMaster(m)

	ver, err := version.NewVersion(git, vopts)
	if err != nil {
		return nil, fmt.Errorf("problem with repo: %w", err)

	}

	return ver, nil
}
