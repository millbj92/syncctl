package models

type TaskArgs struct {
	Path        string     `yaml:"path"`
	Glob        string     `yaml:"glob"`
	Regex       string     `yaml:"regex"`
	Ignore      []string   `yaml:"ignore`
	Recursive   bool       `yaml:"recursive"`
	Content     string     `yaml:"content"`
	Destination string     `yaml:"destination"`
	Replacement string     `yaml:"replacement"`
	Force 		bool       `yaml:"force"`
}

type Task struct {
	Name        string `yaml:"name"`
	Description string `yaml:"description"`
	Action 	    string `yaml:"action"`
	Tags		[]string `yaml:"tags"`
	Args		TaskArgs `yaml:"args"`
 }
