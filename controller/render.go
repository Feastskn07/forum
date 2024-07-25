package render

import (
	"log"
	"net/http"
	"path/filepath"
	"text/template"
)

func RenderTemplate(w http.ResponseWriter, tmp string, data map[string]interface{}) {
	templateCache, err := createTemplateCache()
	if err != nil {
		log.Printf("Şablon önbelleği oluşturulurken hata oluştu: %v", err)
		http.Error(w, "İç sunucu hatası", http.StatusInternalServerError)
		return
	}
	t, ok := templateCache[tmp]
	if !ok {
		log.Printf("Şablon bulunamadı: %s", tmp)
		http.Error(w, "Şablon bulunamadı", http.StatusInternalServerError)
		return
	}
	err = t.ExecuteTemplate(w, tmp[:len(tmp)-5], data)
	if err != nil {
		log.Printf("Şablon çıktısı yazma hatası: %v", err)
		http.Error(w, "Şablon çıktısı yazma hatası", http.StatusInternalServerError)
	}
}

func createTemplateCache() (map[string]*template.Template, error) {
	myCache := map[string]*template.Template{}
	pages, err := filepath.Glob("./templates/views/*.html")
	if err != nil {
		return myCache, err
	}
	for _, page := range pages {
		name := filepath.Base(page)
		ts, err := template.New(name).ParseFiles(page)
		if err != nil {
			return myCache, err
		}
		matches, err := filepath.Glob("./templates/layouts/*.html")
		if err != nil {
			return myCache, err
		}
		if len(matches) > 0 {
			ts, err = ts.ParseGlob("./templates/layouts/*.html")
			if err != nil {
				return myCache, err
			}
		}
		myCache[name] = ts
	}
	return myCache, nil
}
