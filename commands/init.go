package commands

import (
	"fmt"
	models "github.com/JasnRathore/arlo/models"
	tmpl "github.com/JasnRathore/arlo/templates"
	utils "github.com/JasnRathore/arlo/utils"
	tea "github.com/charmbracelet/bubbletea"
	"os"
	"os/exec"
	"regexp"
	"strings"
)

type state1 int

const (
	stateInput state1 = iota
	stateMenu
	stateDone
)

type model struct {
	state     state1
	input     string
	selection int
	options   []string
	done      bool
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch m.state {
	case stateInput:
		switch msg := msg.(type) {
		case tea.KeyMsg:
			switch msg.Type {
			case tea.KeyEnter:
				m.state = stateMenu
			case tea.KeyBackspace:
				if len(m.input) > 0 {
					m.input = m.input[:len(m.input)-1]
				}
			case tea.KeyCtrlC:
				return m, tea.Quit
			default:
				m.input += msg.String()
			}
		}
	case stateMenu:
		switch msg := msg.(type) {
		case tea.KeyMsg:
			switch msg.Type {
			case tea.KeyUp:
				if m.selection > 0 {
					m.selection--
				}
			case tea.KeyDown:
				if m.selection < len(m.options)-1 {
					m.selection++
				}
			case tea.KeyEnter:
				m.state = stateDone
				m.done = true
				return m, tea.Quit
			case tea.KeyCtrlC:
				return m, tea.Quit
			}
		}
	}
	return m, nil
}

func (m model) View() string {
	var b strings.Builder
	switch m.state {
	case stateInput:
		b.WriteString("Enter Project Name: ")
		b.WriteString(m.input)
	case stateMenu:
		b.WriteString(fmt.Sprintf("You typed: %s\n\nChoose an option:\n\n", m.input))
		for i, opt := range m.options {
			cursor := " "
			if i == m.selection {
				cursor = ">"
			}
			b.WriteString(fmt.Sprintf("%s %s\n", cursor, opt))
		}
	case stateDone:
		b.WriteString(fmt.Sprintf("You typed: %s\n", m.input))
		b.WriteString(fmt.Sprintf("You selected: %s\n", m.options[m.selection]))
	}
	return b.String()
}

// State for model2
type state2 int

const (
	state2Menu state2 = iota
	state2Done
)

type model2 struct {
	state     state2
	selection int
	options   []string
	done      bool
}

func (m model2) Init() tea.Cmd {
	return nil
}

func (m model2) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch m.state {
	case state2Menu:
		switch msg := msg.(type) {
		case tea.KeyMsg:
			switch msg.Type {
			case tea.KeyUp:
				if m.selection > 0 {
					m.selection--
				}
			case tea.KeyDown:
				if m.selection < len(m.options)-1 {
					m.selection++
				}
			case tea.KeyEnter:
				m.state = state2Done
				m.done = true
				return m, tea.Quit
			case tea.KeyCtrlC:
				return m, tea.Quit
			}
		}
	case state2Done:
		// No further interaction after selection
	}
	return m, nil
}

func (m model2) View() string {
	var b strings.Builder
	switch m.state {
	case state2Menu:
		b.WriteString("Choose an option:\n\n")
		for i, opt := range m.options {
			cursor := " "
			if i == m.selection {
				cursor = ">"
			}
			b.WriteString(fmt.Sprintf("%s %s\n", cursor, opt))
		}
	case state2Done:
		b.WriteString(fmt.Sprintf("You selected: %s\n", m.options[m.selection]))
	}
	return b.String()
}

func CreateFrontend(packageManager string, name string) {
	switch packageManager {
	case "npm":
		err := utils.RunCommand("npm", "create", "vite@latest", strings.ToLower(name))
		utils.Check(err)
	case "pnpm":
		err := utils.RunCommand("pnpm", "create", "vite", strings.ToLower(name))
		utils.Check(err)
	case "yarn":
		err := utils.RunCommand("yarn", "create", "vite", strings.ToLower(name))
		utils.Check(err)
	case "bun":
		err := utils.RunCommand("bun", "create", "vite", strings.ToLower(name))
		utils.Check(err)
	case "deno":
		err := utils.RunCommand("deno", "init", "--npm", "vite", strings.ToLower(name))
		utils.Check(err)
	default:
		fmt.Println("default")
	}
}

func InstallFrontendDependencies(packageManager string) {
	switch packageManager {
	case "npm":
		err := utils.RunCommand("npm", "install")
		utils.Check(err)
	case "pnpm":
		err := utils.RunCommand("pnpm", "install")
		utils.Check(err)
	case "yarn":
		err := utils.RunCommand("yarn")
		utils.Check(err)
	case "bun":
		err := utils.RunCommand("bun", "install")
		utils.Check(err)
	case "deno":
		err := utils.RunCommand("deno", "install")
		utils.Check(err)
	default:
		fmt.Println("default")
	}
}

func GetFrontendDependenciesCommand(packageManager string) string {
	switch packageManager {
	case "npm":
		return "npm install"
	case "pnpm":
		return "pnpm install"
	case "yarn":
		return "yarn"
	case "bun":
		return "bun install"
	case "deno":
		return "deno install"
	default:
		return "default"
	}
}

func InstallNodeTypes(packageManager string) {
	var err error
	switch packageManager {
	case "npm":
		err = utils.RunCommand("npm", "i", "--save-dev", "@types/node")
	case "pnpm":
		err = utils.RunCommand("pnpm", "add", "-D", "@types/node")
	case "yarn":
		err = utils.RunCommand("yarn", "add", "--dev", "@types/node")
	case "bun":
		err = utils.RunCommand("bun", "add", "-d", "@types/node")
	case "deno":
		// Deno does not use npm packages like @types/node, so you might want to handle this differently.
		fmt.Println("Deno does not require @types/node.")
		return
	default:
		fmt.Println("Unsupported package manager:", packageManager)
		return
	}
	utils.Check(err)
}

func ui() (models.ProjectDetails, error) {
	m := model{
		state:   stateInput,
		options: []string{"npm", "yarn", "pnpm", "deno", "bun"},
	}

	prog := tea.NewProgram(m)
	finalModel, err := prog.Run()
	utils.Check(err)

	m = finalModel.(model) // type assert to get final state

	if m.done {
		return models.ProjectDetails{
			Name:           m.input,
			PackageManager: m.options[m.selection],
		}, nil
	}
	return models.ProjectDetails{}, err
}

func ui2() (string, error) {
	m := model2{
		state:   state2Menu,
		options: []string{"Standard", "Gin"},
	}

	prog := tea.NewProgram(m)
	finalModel, err := prog.Run()
	utils.Check(err)

	m = finalModel.(model2) // type assert to get final state

	if m.done {
		return m.options[m.selection], nil
	}
	return "", err
}

// checkCommand returns true if the command is available in the system PATH
func checkCommand(name string) bool {
	_, err := exec.LookPath(name)
	return err == nil
}

// CheckDependencies verifies all required tools and returns true if all are installed
func CheckDependencies(jsTool string) bool {
	coreTools := []string{"go", "node", "air"}
	jsOptions := map[string]bool{
		"npm": true, "deno": true, "pnpm": true, "bun": true, "yarn": true,
	}

	jsTool = strings.ToLower(jsTool)
	if !jsOptions[jsTool] {
		fmt.Printf("⚠ '%s' is not a supported JS tool\n", jsTool)
		return false
	}

	allInstalled := true

	// Check core tools
	for _, tool := range coreTools {
		if checkCommand(tool) {
			fmt.Printf("✔ %s is installed\n", tool)
		} else {
			fmt.Printf("✘ %s is NOT installed\n", tool)
			allInstalled = false
		}
	}

	// Check selected JS tool
	if checkCommand(jsTool) {
		fmt.Printf("✔ %s is installed\n", jsTool)
	} else {
		fmt.Printf("✘ %s is NOT installed\n", jsTool)
		allInstalled = false
	}

	return allInstalled
}

func PatchViteConfig() error {
	files := []string{"vite.config.ts", "vite.config.js"}
	var configFile string
	for _, f := range files {
		if _, err := os.Stat(f); err == nil {
			configFile = f
			break
		}
	}
	if configFile == "" {
		return fmt.Errorf("No vite.config.js or vite.config.ts found")
	}

	data, err := os.ReadFile(configFile)
	if err != nil {
		return err
	}
	content := string(data)

	// Ensure loadEnv is imported
	importRegex := regexp.MustCompile(`import\s+\{\s*defineConfig\s*\}\s+from\s+['"]vite['"]`)
	content = importRegex.ReplaceAllString(content, "import { defineConfig, loadEnv } from 'vite'")

	// If already function-based, just patch the proxy target
	if strings.Contains(content, "defineConfig(({ mode })") {
		content = strings.ReplaceAll(content, "target: process.env.VITE_API_URL", "target: env.VITE_API_URL")
		return os.WriteFile(configFile, []byte(content), 0644)
	}

	// Convert object-based config to function-based
	defineConfigRegex := regexp.MustCompile(`export\s+default\s+defineConfig\s*\(\s*\{`)
	loc := defineConfigRegex.FindStringIndex(content)
	if loc == nil {
		return fmt.Errorf("Could not find export default defineConfig({ in config file")
	}

	// Find the closing }) at the end
	lastClosing := strings.LastIndex(content, "})")
	if lastClosing == -1 {
		return fmt.Errorf("Could not find closing }) in config file")
	}

	// Extract the config object
	configBody := content[loc[1]:lastClosing]
	configBody = strings.TrimSpace(configBody)
	if strings.HasPrefix(configBody, "{") {
		configBody = configBody[1:]
	}
	if strings.HasSuffix(configBody, "}") {
		configBody = configBody[:len(configBody)-1]
	}

	serverProxy := `
    server: {
      proxy: {
        '/api': {
          target: env.VITE_API_URL,
          changeOrigin: true,
          rewrite: path => path.replace(/^\/api/, ''),
        }
      }
    }
  }
	`
	// Build the new function-based config
	newConfig := `
export default defineConfig(({ mode }) => {
  const env = loadEnv(mode, process.cwd(), '')
  return {
	` + configBody + serverProxy + `
})
`
	// Replace the old config with the new one
	content = content[:loc[0]] + newConfig

	return os.WriteFile(configFile, []byte(content), 0644)
}

func CreateBackend(framework string, data models.TemplateData) {
	switch framework {
	case "Standard":
		tmpl.GenerateTemplate("std.main.go.tmpl", "main.go", data)
		err := os.Mkdir("app", 0755)
		utils.Check(err)
		tmpl.GenerateTemplate("std.app.go.tmpl", "app/app.go", data)
		tmpl.GenerateTemplate("std.build.go.tmpl", "build.go", data)
		return
	case "Gin":
		tmpl.GenerateTemplate("gin.main.go.tmpl", "main.go", data)
		err := os.Mkdir("app", 0755)
		utils.Check(err)
		tmpl.GenerateTemplate("gin.app.go.tmpl", "app/app.go", data)
		tmpl.GenerateTemplate("gin.build.go.tmpl", "build.go", data)
		err = utils.RunCommand("go", "get", "github.com/gin-gonic/gin")
		utils.Check(err)
		return
	default:
		fmt.Println("Invalid Option")

	}
}

func InitProject() {
	//createing the webapp

	project, err := ui()
	utils.Check(err)

	fmt.Println()
	installed := CheckDependencies(project.PackageManager)
	if installed {
		fmt.Println("\n✅ All dependencies are installed.")
	} else {
		fmt.Println("\n❌ Some dependencies are missing.")
		return
	}
	fmt.Println()
	fmt.Println("Select Go API framework:")
	goFramework, err := ui2()

	CreateFrontend(project.PackageManager, project.Name)

	//go into projfolder
	err = os.Chdir(strings.ToLower(project.Name))
	utils.Check(err)

	//making the arlo config file
	jsonData, err := utils.StructToJSON(project)
	utils.Check(err)
	err = utils.WriteJSONToFile("arlo.config.json", jsonData)
	utils.Check(err)
	InstallNodeTypes(project.PackageManager)

	//tmpl.CopyTemplate("vite.config.ts.tmpl", "vite.config.ts")
	err = PatchViteConfig()
	utils.Check(err)

	//making the src-arlo dir
	dirName := "src-backend"
	err = os.Mkdir(dirName, 0755)
	utils.Check(err)
	err = os.Chdir(dirName)
	utils.Check(err)

	//init go proj
	utils.RunCommand("go", "mod", "init", strings.ToLower(project.Name))

	//init air {hotrealoading}
	tmpl.CopyTemplate("air.toml.tmpl", ".air.toml")

	//generating src-backend files
	data := models.TemplateData{
		Title: strings.ToLower(project.Name),
	}

	CreateBackend(goFramework, data)

	fmt.Println("cd", strings.ToLower(project.Name))
	fmt.Println(GetFrontendDependenciesCommand(project.PackageManager))
	fmt.Println("arlo dev")

}
