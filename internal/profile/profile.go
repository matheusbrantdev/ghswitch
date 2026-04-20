package profile

type Profile struct {
	Name     string `yaml:"name"`
	GitName  string `yaml:"git_name"`
	GitEmail string `yaml:"git_email"`
	SSHKey   string `yaml:"ssh_key"`
}
