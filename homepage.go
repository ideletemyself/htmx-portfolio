package main

import (
    "fmt"
    "html/template"
    "net/http"
)

func HomePageHandler(w http.ResponseWriter, r *http.Request) {
    // Assuming IsHTMXRequest checks if the request is an HTMX request
    isHTMX := IsHTMXRequest(r)

    // Load all necessary templates. Since we're using PageType to control
    // which content is rendered, we only need to parse 'base.html' once.
    tmpl, err := template.New("base").Funcs(template.FuncMap{
        "safeHTML": func(htmlContent []byte) template.HTML {
            return template.HTML(htmlContent)
        },
    }).ParseFiles(
        "templates/base.html", // base.html should now include conditional logic for rendering content
        "templates/header.html", // Ensure these are included in base.html as {{template "header.html" .}}
        "templates/homepage.html", // Now defined as {{define "homepageContent"}}...{{end}}
        "templates/footer.html", // Ensure these are included in base.html as {{template "footer.html" .}}
        "templates/hero.html",
        "templates/bloglist.html",
        "templates/blogpost.html", // This can be included in homepage.html or base.html based on your design
    )
    if err != nil {
        fmt.Printf("Error loading templates: %v\n", err) // More detailed error logging
        http.Error(w, "Error loading templates", http.StatusInternalServerError)
        return
    }

    // Prepare the data for template execution, including a PageType field
    data := map[string]interface{}{
        "Title":   "Home Page",
        "PageType": "homepage", // Used to control which content section to render
        "IsHTMX":  isHTMX, // Pass the IsHTMX flag to the template if needed
    }

    // Execute the base template, which includes logic to render the appropriate content section
    err = tmpl.ExecuteTemplate(w, "base.html", data)
    if err != nil {
        fmt.Printf("Error rendering page: %v\n", err) // More detailed error logging
        http.Error(w, "Error rendering page", http.StatusInternalServerError)
    }
}

