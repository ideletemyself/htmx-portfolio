package main

import (
    "fmt"
    "os"
    "time"
    "sort"
    "html/template"
    "net/http"
    "io/ioutil"
    "path/filepath"
    "strings"
    "gopkg.in/yaml.v2"
    "github.com/russross/blackfriday/v2"
)

type BlogPost struct {
    Title       string
    Date        time.Time
    Description string
    Image       string
    FileName    string 
    Content template.HTML
}

// IsHTMXRequest checks if the request is coming from HTMX
func IsHTMXRequest(r *http.Request) bool {
    return r.Header.Get("HX-Request") == "true"
}

func main() {
    // Serve static files
    http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./static"))))

    // Homepage handler (assuming HomePageHandler is defined in homepage.go)
    http.HandleFunc("/", HomePageHandler)

    // Blog list handler
    http.HandleFunc("/blog", BlogListHandler)

    // Individual blog content handler
    http.HandleFunc("/content/", BlogContentHandler)

    fmt.Println("Listening on http://localhost:8080/")
    http.ListenAndServe(":8080", nil)
}

// BlogContentHandler serves the markdown content as HTML
func BlogContentHandler(w http.ResponseWriter, r *http.Request) {
    postName := strings.TrimPrefix(r.URL.Path, "/content/")
    if postName == "" {
        http.NotFound(w, r)
        return
    }

    filePath := filepath.Join("posts", postName+".md")
    fileContent, err := ioutil.ReadFile(filePath)
    if err != nil {
        fmt.Printf("Failed to read file: %v\n", err)
        http.Error(w, "File not found.", http.StatusNotFound)
        return
    }

    parts := strings.SplitN(string(fileContent), "---", 3)
    if len(parts) < 3 {
        http.Error(w, "Error processing blog post.", http.StatusInternalServerError)
        return
    }

    var metaData struct {
        Title string `yaml:"title"`
        Description string `yaml:"description"`
        Image string `yaml:"image"` 
    }
    err = yaml.Unmarshal([]byte(parts[1]), &metaData)
    if err != nil {
        fmt.Printf("Failed to parse YAML: %v\n", err)
        http.Error(w, "Error processing metadata", http.StatusInternalServerError)
        return
    }

    markdownBody := parts[2]
    htmlContent := blackfriday.Run([]byte(markdownBody))

    tmpl, err := template.New("").Funcs(template.FuncMap{
        "safeHTML": func(htmlContent template.HTML) template.HTML {
            return htmlContent
        },
    }).ParseFiles(
        "templates/base.html", "templates/header.html", "templates/blogpost.html", 
        "templates/footer.html", "templates/hero.html", "templates/homepage.html",
        "templates/bloglist.html",
    )
    if err != nil {
        http.Error(w, "Error loading template", http.StatusInternalServerError)
        return
    }

    data := map[string]interface{}{
        "Title":   metaData.Title,
        "PageType": "blogpost",
        "Content": template.HTML(htmlContent), // Directly use template.HTML without safeHTML
    }

    err = tmpl.ExecuteTemplate(w, "base.html", data)
    if err != nil {
        fmt.Printf("Error rendering template: %v\n", err)
        http.Error(w, "Error rendering template", http.StatusInternalServerError)
    }
}

func LoadBlogPosts(postsDir string) ([]BlogPost, error) {
    var posts []BlogPost

    err := filepath.Walk(postsDir, func(path string, info os.FileInfo, err error) error {
        if err != nil {
            return err
        }
        if !info.IsDir() && strings.HasSuffix(info.Name(), ".md") {
            content, err := ioutil.ReadFile(path)
            if err != nil {
                return err
            }

            // Split the content to extract the YAML front matter and Markdown
            parts := strings.SplitN(string(content), "---", 3)
            if len(parts) < 3 {
                return nil // Not a valid post format
            }

            var post BlogPost
            err = yaml.Unmarshal([]byte(parts[1]), &post)
            if err != nil {
                return err
            }

            // Convert the Markdown content to HTML
            markdownBody := parts[2]
            htmlContent := blackfriday.Run([]byte(markdownBody))
            post.Content = template.HTML(htmlContent)

            post.FileName = strings.TrimSuffix(info.Name(), ".md")

            posts = append(posts, post)
        }
        return nil
    })

    return posts, err
}

func BlogListHandler(w http.ResponseWriter, r *http.Request) {
    // Determine if the request is coming from HTMX (assuming IsHTMXRequest is correctly implemented)
    isHTMX := IsHTMXRequest(r)

    // Load blog posts
    posts, err := LoadBlogPosts("posts")
    if err != nil {
        fmt.Printf("Error loading blog posts: %v\n", err) // For debug purposes
        http.Error(w, "Error loading blog posts", http.StatusInternalServerError)
        return
    }

    // Sort the posts by date, most recent first
    sort.Slice(posts, func(i, j int) bool {
        return posts[i].Date.After(posts[j].Date)
    })

    tmpl, err := template.New("base").Funcs(template.FuncMap{
        "safeHTML": func(htmlContent []byte) template.HTML {
            return template.HTML(htmlContent)
        },
    }).ParseFiles(
        "templates/base.html",
        "templates/header.html",
        "templates/bloglist.html",
        "templates/footer.html",
        "templates/homepage.html",
        "templates/hero.html",
        "templates/blogpost.html",
    )
    if err != nil {
        fmt.Printf("Error loading templates: %v\n", err) // For debug purposes
        http.Error(w, "Error loading template", http.StatusInternalServerError)
        return
    }

    // Prepare data for the template, including a flag or identifier for the page type
    data := map[string]interface{}{
        "Title":    "Blog Posts",
        "Posts":    posts,
        "PageType": "bloglist", // This is used to conditionally render the blog list content in base.html
        "IsHTMX":   isHTMX,     // Pass the IsHTMX flag to the template, if needed for further conditional rendering
    }

    // Execute the base template, which now includes the logic to conditionally render content based on PageType
    // Note: Ensure your base.html template correctly handles this logic as described previously
    err = tmpl.ExecuteTemplate(w, "base.html", data)
    if err != nil {
        fmt.Printf("Error rendering template: %v\n", err) // For debug purposes
        http.Error(w, "Error rendering template", http.StatusInternalServerError)
    }
}

