package app

// Config yaml
type Config struct {
	Info      `yaml:",inline"`
	Overrides map[string]Overridables `yaml:"overrides,omitempty"`
}

// Info contains information about a single package
type Info struct {
	Overridables `yaml:",inline"`
	Name         string `yaml:"name,omitempty"`
	Arch         string `yaml:"arch,omitempty"`
	Platform     string `yaml:"platform,omitempty"`
	Version      string `yaml:"version,omitempty"`
	Section      string `yaml:"section,omitempty"`
	Priority     string `yaml:"priority,omitempty"`
	Maintainer   string `yaml:"maintainer,omitempty"`
	Description  string `yaml:"description,omitempty"`
	Vendor       string `yaml:"vendor,omitempty"`
	Homepage     string `yaml:"homepage,omitempty"`
	License      string `yaml:"license,omitempty"`
}

// Overridables contain the field which are overridable in a package
type Overridables struct {
	Replaces     []string          `yaml:"replaces,omitempty"`
	Provides     []string          `yaml:"provides,omitempty"`
	Depends      []string          `yaml:"depends,omitempty"`
	Recommends   []string          `yaml:"recommends,omitempty"`
	Suggests     []string          `yaml:"suggests,omitempty"`
	Conflicts    []string          `yaml:"conflicts,omitempty"`
	Files        map[string]string `yaml:"files,omitempty"`
	ConfigFiles  map[string]string `yaml:"config_files,omitempty"`
	EmptyFolders []string          `yaml:"empty_folders,omitempty"`
	Bindir       string            `yaml:"bindir,omitempty"`
	Changelog    string            `yaml:"changelog,omitempty"`
	Scripts      Scripts           `yaml:"scripts,omitempty"`
}

// Scripts contains information about maintainer scripts for packages
type Scripts struct {
	PreInstall  string `yaml:"preinstall,omitempty"`
	PostInstall string `yaml:"postinstall,omitempty"`
	PreRemove   string `yaml:"preremove,omitempty"`
	PostRemove  string `yaml:"postremove,omitempty"`
}
