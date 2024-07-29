package render

import (
	"html/template"
	"log"
	"net/http"
	"path/filepath"
)

func RenderTemplate(w http.ResponseWriter, tmp string, data map[string]interface{}) {
	templateCache, err := createTemplateCache()

	if err != nil {
		log.Printf("An error occurred while creating the template cache: %v", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	t, ok := templateCache[tmp]
	if !ok {
		log.Printf("Tamplate not found: %s", tmp)
		http.Error(w, "Template not found", http.StatusInternalServerError)
		return
	}

	err = t.ExecuteTemplate(w, tmp[:len(tmp)-10], data)

	if err != nil {
		log.Printf("Error writing template output: %v", err)
		http.Error(w, "Error writing template output", http.StatusInternalServerError)
	}
}

func createTemplateCache() (map[string]*template.Template, error) {
	myCache := map[string]*template.Template{}
	pages, err := filepath.Glob("./templates/views/*-page.html")

	if err != nil {
		return myCache, err
	}

	for _, page := range pages {
		name := filepath.Base(page)
		ts, err := template.New(name).ParseFiles(page)

		if err != nil {
			return myCache, err
		}

		matches, err := filepath.Glob("./tempalates/layouts/*-layout.html")

		if err != nil {
			return myCache, err
		}

		if len(matches) > 0 {
			ts, err = ts.ParseGlob("./templates/layouts/*-layout.html")

			if err != nil {
				return myCache, err
			}
		}
		myCache[name] = ts
	}
	return myCache, nil
}
