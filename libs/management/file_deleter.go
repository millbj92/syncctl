package management

import (
	"log"
	"os"
	"path/filepath"
	"regexp"

	"github.com/millbj92/synctl/libs/models/tasks"
)

func DeleteFiles(args TaskArgs) error {
	log.Println("DeleteFiles:", args)
	matches, err := filepath.Glob(args.Glob)
	if err != nil {
		return err
	}
	for _, match := range matches {
		log.Println("Matched:", match)
		st, err := os.Stat(match); if err != nil {
			return err
		}
		if !contains(args.Ignore, st.Name()) &&
		   !ContainsRegex(args.Ignore, st.Name()) {
			log.Println("Deleting: ", match)
			//err := os.Remove(match)
			if err != nil {
				return err
			}
	    } else {
			log.Println("Ignore matched, Skipping: ", match)
		}
	}
	return nil
}


func ContainsRegex (s []string, e string) bool {
	for _, a := range s {
		r, _  := regexp.Compile(a)
		m := r.FindStringIndex(a)
		if m != nil {
			log.Println("Regex:", a, " matched:", e)
			return true
		}
	}
	return false
}

func contains (s []string, e string) bool {
	for _, a := range s {
		if a == e {
			log.Println("Ignore:", a, " matched:", e)
			return true
		}
	}
	return false
}
