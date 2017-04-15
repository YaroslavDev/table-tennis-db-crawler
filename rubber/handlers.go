package rubber

import (
	"html/template"
	"net/http"
	"log"
	"upper.io/db.v3/mysql"
)

type rubberController struct {
	template *template.Template
	finder   RubberFinder
	rubberRepository RubberRepository
}

func NewRubberController(template *template.Template, connectionUrl *mysql.ConnectionURL) *rubberController {
	return &rubberController{
		template: template,
		finder: NewTTDBRubberFinder(),
		rubberRepository: NewRubberRepository(connectionUrl),
	}
}

func (r rubberController) RubbersPage(w http.ResponseWriter, req *http.Request) {
	rubbers, err := r.rubberRepository.GetRubbers()
	if err != nil {
		log.Fatalln(err)
	}

	err = r.template.ExecuteTemplate(w, "rubbers.html", rubbers)
	if err != nil {
		log.Fatalln(err)
	}
}

func (r rubberController) SynchronizeRubbers(w http.ResponseWriter, req *http.Request) {
	go func() {
		rubbers, err := r.finder.FindRubbers()
		if err != nil {
			log.Fatalln(err)
		}
		for _, rubber := range rubbers {
			if err = r.rubberRepository.SaveRubber(rubber); err != nil {
				log.Printf("Could not save %v\n", rubber)
				log.Fatalln(err)
			}
		}
	}()
	http.Redirect(w, req, "/rubber", http.StatusSeeOther)
}