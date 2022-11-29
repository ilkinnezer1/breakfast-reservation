package render

import (
	"fmt"
	"html/template"
	"net/http"
)

// HandleError handles with errors
func HandleError(err error){
	if err != nil {
		fmt.Println(err)
		return
	} 
} 

func RenderTemplate(w http.ResponseWriter, html string){
	// Parse template
	parsedTmpl, _ := template.ParseFiles("./templates/" + html)
	// Execute the parse function
	err := parsedTmpl.Execute(w, nil)
	HandleError(err)
}