package main

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"gopkg.in/yaml.v3"
)

// ── Data types ────────────────────────────────────────────────────────────────

type ModuleSpec struct {
	Tracking  *string           `yaml:"tracking"`
	Additions map[string]string `yaml:"additions"`
}

type Spec map[string]ModuleSpec

type ResourceType struct {
	ResourceType      string   `json:"resourceType"`
	ApiVersions       []string `json:"apiVersions"`
	DefaultApiVersion *string  `json:"defaultApiVersion"`
}

type Provider struct {
	Namespace     string         `json:"namespace"`
	ResourceTypes []ResourceType `json:"resourceTypes"`
}

// ── Status ────────────────────────────────────────────────────────────────────

type UpdateStatus int

const (
	StatusNeedsBump UpdateStatus = iota
	StatusNotFound  UpdateStatus = iota
	StatusUpToDate  UpdateStatus = iota
)

// ── Row ───────────────────────────────────────────────────────────────────────

type Row struct {
	Module  string
	Current string
	Latest  string
	Status  UpdateStatus
}

// ── Data loading ──────────────────────────────────────────────────────────────

func loadSpec(path string) (Spec, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	var spec Spec
	if err := yaml.Unmarshal(data, &spec); err != nil {
		return nil, err
	}
	return spec, nil
}

func loadProviderList(path string) ([]Provider, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	var providers []Provider
	if err := json.Unmarshal(data, &providers); err != nil {
		return nil, err
	}
	return providers, nil
}

func isPreview(v string) bool {
	lower := strings.ToLower(v)
	return strings.HasSuffix(lower, "-preview") ||
		strings.HasSuffix(lower, "-beta") ||
		strings.HasSuffix(lower, "-privatepreview")
}

func latestVersion(versions []string) string {
	var stable, preview []string
	for _, v := range versions {
		if isPreview(v) {
			preview = append(preview, v)
		} else {
			stable = append(stable, v)
		}
	}
	if len(stable) > 0 {
		sort.Strings(stable)
		return stable[len(stable)-1]
	}
	if len(preview) > 0 {
		sort.Strings(preview)
		return preview[len(preview)-1]
	}
	return ""
}

func buildNamespaceIndex(providers []Provider) map[string]string {
	index := make(map[string]string)
	for _, p := range providers {
		lower := strings.ToLower(p.Namespace)
		if !strings.HasPrefix(lower, "microsoft.") {
			continue
		}
		seen := map[string]struct{}{}
		var all []string
		for _, rt := range p.ResourceTypes {
			for _, v := range rt.ApiVersions {
				if _, ok := seen[v]; !ok {
					seen[v] = struct{}{}
					all = append(all, v)
				}
			}
		}
		index[lower] = latestVersion(all)
	}
	return index
}

func buildRows(spec Spec, index map[string]string) []Row {
	var rows []Row
	for moduleName, moduleSpec := range spec {
		current := ""
		if moduleSpec.Tracking != nil {
			current = *moduleSpec.Tracking
		}
		nsKey := "microsoft." + strings.ToLower(moduleName)
		latest, found := index[nsKey]

		var status UpdateStatus
		switch {
		case !found:
			status = StatusNotFound
		case current == "":
			status = StatusNotFound
		case current == latest:
			status = StatusUpToDate
		default:
			status = StatusNeedsBump
		}

		rows = append(rows, Row{
			Module:  moduleName,
			Current: current,
			Latest:  latest,
			Status:  status,
		})
	}
	sort.Slice(rows, func(i, j int) bool {
		if rows[i].Status != rows[j].Status {
			return rows[i].Status < rows[j].Status
		}
		return rows[i].Module < rows[j].Module
	})
	return rows
}

// ── Styles ────────────────────────────────────────────────────────────────────

var (
	titleStyle    = lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color("255"))
	headerStyle   = lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color("255"))
	selectedStyle = lipgloss.NewStyle().Background(lipgloss.Color("237"))
	bumpStyle     = lipgloss.NewStyle().Foreground(lipgloss.Color("214"))
	upToDateStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("46"))
	notFoundStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("240"))
	filterStyle   = lipgloss.NewStyle().Foreground(lipgloss.Color("81"))
	statusBar     = lipgloss.NewStyle().Foreground(lipgloss.Color("240"))
	dimStyle      = lipgloss.NewStyle().Foreground(lipgloss.Color("240"))
	menuCursorStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("81")).Bold(true)
)

func padRight(s string, n int) string {
	if len(s) >= n {
		return s[:n-1] + " "
	}
	return s + strings.Repeat(" ", n-len(s))
}

func maxInt(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// ── Navigation message ────────────────────────────────────────────────────────

type navigateMsg struct{ index int }
type backMsg struct{}

// ── Menu model ────────────────────────────────────────────────────────────────

type menuItem struct {
	label       string
	description string
}

type menuModel struct {
	items  []menuItem
	cursor int
	height int
}

func newMenuModel() menuModel {
	return menuModel{
		items: []menuItem{
			{
				label:       "Version Tracking Updates",
				description: "Compare v3-spec.yaml tracking versions against latest available API versions",
			},
		},
	}
}

func (m menuModel) Init() tea.Cmd { return nil }

func (m menuModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.height = msg.Height
	case tea.KeyMsg:
		switch msg.String() {
		case "q", "ctrl+c":
			return m, tea.Quit
		case "j", "down":
			if m.cursor < len(m.items)-1 {
				m.cursor++
			}
		case "k", "up":
			if m.cursor > 0 {
				m.cursor--
			}
		case "enter", " ":
			return m, func() tea.Msg { return navigateMsg{index: m.cursor} }
		}
	}
	return m, nil
}

func (m menuModel) View() string {
	var sb strings.Builder

	sb.WriteString(titleStyle.Render("Azure Native Provider Tools") + "\n")
	sb.WriteString(dimStyle.Render(strings.Repeat("─", 50)) + "\n\n")

	for i, item := range m.items {
		cursor := "  "
		label := item.label
		desc := dimStyle.Render("  " + item.description)

		if i == m.cursor {
			cursor = menuCursorStyle.Render("▶ ")
			label = menuCursorStyle.Render(label)
		}

		sb.WriteString(cursor + label + "\n")
		sb.WriteString(desc + "\n\n")
	}

	sb.WriteString(dimStyle.Render("↑/↓ navigate  enter select  q quit"))
	return sb.String()
}

// ── Tracker model (version tracking sub-view) ─────────────────────────────────

const (
	colModule  = 30
	colCurrent = 22
	colLatest  = 22
	colStatus  = 8
)

type trackerModel struct {
	allRows   []Row
	rows      []Row
	cursor    int
	filter    string
	filtering bool
	height    int
}

func newTrackerModel(rows []Row) trackerModel {
	return trackerModel{allRows: rows, rows: rows}
}

func (m trackerModel) Init() tea.Cmd { return nil }

func (m trackerModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.height = msg.Height

	case tea.KeyMsg:
		if m.filtering {
			switch msg.Type {
			case tea.KeyEscape:
				m.filtering = false
				m.filter = ""
				m.applyFilter()
			case tea.KeyEnter:
				m.filtering = false
			case tea.KeyBackspace:
				if len(m.filter) > 0 {
					m.filter = m.filter[:len(m.filter)-1]
					m.applyFilter()
				}
			default:
				if msg.Type == tea.KeyRunes {
					m.filter += string(msg.Runes)
					m.applyFilter()
				}
			}
			return m, nil
		}

		switch msg.String() {
		case "ctrl+c":
			return m, tea.Quit
		case "q", "esc":
			return m, func() tea.Msg { return backMsg{} }
		case "j", "down":
			if m.cursor < len(m.rows)-1 {
				m.cursor++
			}
		case "k", "up":
			if m.cursor > 0 {
				m.cursor--
			}
		case "g":
			m.cursor = 0
		case "G":
			m.cursor = len(m.rows) - 1
		case "/":
			m.filtering = true
		}
	}
	return m, nil
}

func (m *trackerModel) applyFilter() {
	lower := strings.ToLower(m.filter)
	m.rows = nil
	for _, r := range m.allRows {
		if strings.Contains(strings.ToLower(r.Module), lower) {
			m.rows = append(m.rows, r)
		}
	}
	if m.cursor >= len(m.rows) {
		m.cursor = maxInt(0, len(m.rows)-1)
	}
}

func (m trackerModel) View() string {
	if m.height == 0 {
		m.height = 40
	}

	var sb strings.Builder

	// Title + back hint
	sb.WriteString(titleStyle.Render("Version Tracking Updates") +
		dimStyle.Render("  esc/q back") + "\n")

	// Column headers
	sb.WriteString(headerStyle.Render(
		padRight("MODULE", colModule)+
			padRight("CURRENT", colCurrent)+
			padRight("LATEST", colLatest)+
			"STATUS",
	) + "\n")
	sb.WriteString(strings.Repeat("─", colModule+colCurrent+colLatest+colStatus) + "\n")

	// Reserve: title(1) + header(1) + separator(1) + statusbar(1) + filter(1) = 5
	visibleLines := m.height - 5
	if visibleLines < 1 {
		visibleLines = 1
	}

	start := 0
	if m.cursor >= visibleLines {
		start = m.cursor - visibleLines + 1
	}
	end := start + visibleLines
	if end > len(m.rows) {
		end = len(m.rows)
	}

	for i := start; i < end; i++ {
		r := m.rows[i]
		line := padRight(r.Module, colModule) +
			padRight(r.Current, colCurrent) +
			padRight(r.Latest, colLatest) +
			statusLabel(r.Status)

		if i == m.cursor {
			line = selectedStyle.Render(line)
		} else {
			line = styleRow(r, line)
		}
		sb.WriteString(line + "\n")
	}

	for i := end - start; i < visibleLines; i++ {
		sb.WriteString("\n")
	}

	// Stats bar
	bump, upToDate, notFound := 0, 0, 0
	for _, r := range m.allRows {
		switch r.Status {
		case StatusNeedsBump:
			bump++
		case StatusUpToDate:
			upToDate++
		case StatusNotFound:
			notFound++
		}
	}
	sb.WriteString(strings.Repeat("─", colModule+colCurrent+colLatest+colStatus) + "\n")
	sb.WriteString(statusBar.Render(fmt.Sprintf(
		" %s needs bump  %s up to date  %s not tracked  |  %d/%d shown  |  /filter",
		bumpStyle.Render(fmt.Sprintf("%d", bump)),
		upToDateStyle.Render(fmt.Sprintf("%d", upToDate)),
		notFoundStyle.Render(fmt.Sprintf("%d", notFound)),
		len(m.rows), len(m.allRows),
	)) + "\n")

	if m.filtering {
		sb.WriteString(filterStyle.Render("/ " + m.filter + "█"))
	} else if m.filter != "" {
		sb.WriteString(filterStyle.Render("filter: " + m.filter + "  (esc to clear)"))
	}

	return sb.String()
}

func statusLabel(s UpdateStatus) string {
	switch s {
	case StatusNeedsBump:
		return bumpStyle.Render("↑ bump")
	case StatusUpToDate:
		return upToDateStyle.Render("✓")
	default:
		return notFoundStyle.Render("?")
	}
}

func styleRow(r Row, line string) string {
	switch r.Status {
	case StatusNeedsBump:
		return bumpStyle.Render(line)
	case StatusUpToDate:
		return upToDateStyle.Render(line)
	default:
		return notFoundStyle.Render(line)
	}
}

// ── App model (top-level router) ──────────────────────────────────────────────

type screen int

const (
	screenMenu    screen = iota
	screenTracker screen = iota
)

type appModel struct {
	screen  screen
	menu    menuModel
	tracker trackerModel
}

func newAppModel(rows []Row) appModel {
	return appModel{
		screen:  screenMenu,
		menu:    newMenuModel(),
		tracker: newTrackerModel(rows),
	}
}

func (a appModel) Init() tea.Cmd { return nil }

func (a appModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	// Propagate window size to both sub-models.
	if ws, ok := msg.(tea.WindowSizeMsg); ok {
		a.menu.height = ws.Height
		a.tracker.height = ws.Height
	}

	switch a.screen {
	case screenMenu:
		updated, cmd := a.menu.Update(msg)
		a.menu = updated.(menuModel)
		if nav, ok := msg.(navigateMsg); ok {
			switch nav.index {
			case 0:
				a.screen = screenTracker
			}
		}
		return a, cmd

	case screenTracker:
		updated, cmd := a.tracker.Update(msg)
		a.tracker = updated.(trackerModel)
		if _, ok := msg.(backMsg); ok {
			a.screen = screenMenu
		}
		return a, cmd
	}

	return a, nil
}

func (a appModel) View() string {
	switch a.screen {
	case screenTracker:
		return a.tracker.View()
	default:
		return a.menu.View()
	}
}

// ── Entry point ───────────────────────────────────────────────────────────────

func repoRoot() string {
	dir, _ := os.Getwd()
	for {
		if _, err := os.Stat(filepath.Join(dir, "versions")); err == nil {
			return dir
		}
		parent := filepath.Dir(dir)
		if parent == dir {
			break
		}
		dir = parent
	}
	return "."
}

func main() {
	root := repoRoot()
	specPath := filepath.Join(root, "versions", "v3-spec.yaml")
	providerListPath := filepath.Join(root, "versions", "az-provider-list.json")

	spec, err := loadSpec(specPath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error reading spec: %v\n", err)
		os.Exit(1)
	}

	providers, err := loadProviderList(providerListPath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error reading provider list: %v\n", err)
		os.Exit(1)
	}

	index := buildNamespaceIndex(providers)
	rows := buildRows(spec, index)

	p := tea.NewProgram(newAppModel(rows), tea.WithAltScreen())
	if _, err := p.Run(); err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
}
