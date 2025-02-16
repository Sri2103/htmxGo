package web

import (
	"bytes"
	"errors"
	"fmt"
	"html/template"
	"io"
	"log"
	"path/filepath"

	sprig "github.com/go-task/slim-sprig/v3"
	"github.com/labstack/echo/v4"
)

type Template struct {
	templates *template.Template
}

func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

type RenderPageWrapper struct {
	TemplateDir string
	Functions   template.FuncMap
	UseCache    bool
	TemplateMap map[string]*template.Template
	Partials    []string
	Debug       bool
}

type Data struct {
	Data map[string]any
}

func NewRender() *RenderPageWrapper {
	return &RenderPageWrapper{
		TemplateDir: "./templates",
		Functions:   sprig.FuncMap(),
		UseCache:    true,
		TemplateMap: make(map[string]*template.Template),
		Partials:    []string{},
		Debug:       true,
	}
}

func (t *RenderPageWrapper) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	td := data.(*Data)
	return t.Show(w, name, td)
}

// Show generates a page of html from our template file(s).
func (ren *RenderPageWrapper) Show(w io.Writer, t string, td *Data) error {
	// Call buildTemplate to get the template, either from the cache or by building it
	// from disk.
	tmpl, err := ren.buildTemplate(t)
	if err != nil {
		return err
	}

	// if we don't have template data, just use an empty struct.
	if td == nil {
		td = &Data{}
	}

	// execute the template
	if err := tmpl.ExecuteTemplate(w, t, td); err != nil {
		return err
	}
	return nil
}

// String renders a template and returns it as a string.
func (ren *RenderPageWrapper) String(w io.Writer, t string, td *Data) (string, error) {
	// Call buildTemplate to get the template, either from the cache or by building it
	// from disk.
	tmpl, err := ren.buildTemplate(t)
	if err != nil {
		return "", err
	}

	// if we don't have template data, just use an empty struct.
	if td == nil {
		td = &Data{}
	}

	// Execute the template, storing the result in a bytes.Buffer variable.
	var tpl bytes.Buffer
	if err := tmpl.Execute(&tpl, td); err != nil {
		return "", err
	}

	// Return a string from the bytes.Buffer.
	result := tpl.String()
	return result, nil
}

// buildTemplate is a utility function that creates a template, either from the cache, or from
// disk. The template is ready to accept functions & data, and then get rendered.
func (ren *RenderPageWrapper) buildTemplate(t string) (*template.Template, error) {
	// tmpl is the variable that will hold our template.
	var tmpl *template.Template

	// If we are using the cache, get try to get the pre-compiled template from our
	// map templateMap, stored in the receiver.
	if ren.UseCache {
		if templateFromMap, ok := ren.TemplateMap[t]; ok {
			if ren.Debug {
				log.Println("Reading template", t, "from cache")
			}
			tmpl = templateFromMap
		}
	}

	// At this point, tmpl will be nil if we do not have a value in the map (our template
	// cache). In this case, we build the template from disk.
	if tmpl == nil {
		newTemplate, err := ren.buildTemplateFromDisk(t)
		if err != nil {
			return nil, err
		}
		tmpl = newTemplate
	}

	return tmpl, nil
}

// buildTemplateFromDisk builds a template from disk.
func (ren *RenderPageWrapper) buildTemplateFromDisk(t string) (*template.Template, error) {
	// templateSlice will hold all the templates necessary to build our finished template.
	var templateSlice []string

	// get filenames of all the components from a path

	components, err := filepath.Glob(fmt.Sprintf("%s/%s", ren.TemplateDir, "components/*.html"))
	if err != nil {
		return nil, errors.New("error getting component files: " + err.Error())
	}
	layoutFile := fmt.Sprintf("layout/%s.html", "base")

	// Append the components to the slice.
	templateSlice = append(templateSlice, components...)
	ren.Partials = append(ren.Partials, layoutFile)

	// Read in the partials, if any.
	for _, x := range ren.Partials {
		// We use filepath.Join to make this OS-agnostic.
		path := filepath.Join(ren.TemplateDir, x)
		templateSlice = append(templateSlice, path)
	}

	// Append the template we want to render to the slice.
	// templateSlice = append(templateSlice, fmt.Sprintf("%s/%s", ren.TemplateDir, t))
	templateSlice = append(templateSlice, filepath.Join(ren.TemplateDir, t))

	// Create a new template by parsing all the files in the slice.
	tmpl, err := template.New(t).Funcs(ren.Functions).ParseFiles(templateSlice...)
	if err != nil {
		return nil, err
	}

	// Add the template to the template map stored in our receiver.
	// Note that this is ignored in development, but does not hurt anything.
	ren.TemplateMap[t] = tmpl

	if ren.Debug {
		log.Println("Reading template", t, "from disk")
	}

	return tmpl, nil
}
