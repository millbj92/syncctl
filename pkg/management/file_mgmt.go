package management

import (
	"io"
	"log"
	"os"
	"path/filepath"
	"regexp"

	"github.com/millbj92/synctl/libs/models/tasks"
)

func DeleteFiles(args tasks.TaskArgs) error {
	log.Println("Starting Delete Task:", args)
	//Gets all ffiles that match the include flag
	matches, err := filepath.Glob(args.Include)
	var toDelete []string
	if err != nil {
		return err
	}
	//For each file that matches the include flag
	//If the file is not in the exclude flag, delete it
	//else, skip and log it
	for _, match := range matches {
		r, _ := regexp.Compile(args.Exclude)
		m := r.FindStringIndex(match)
		if m == nil {
			log.Println("Deleting: ", match)
			toDelete = append(toDelete, match)
			if err != nil {
				return err
			}

		} else {
			log.Println("Exclude matched, Skipping: ", filepath.Base(match))
		}
	}
	return nil
}

func RunDelete(toDelete []string) error {
	for _, file := range toDelete {
		err := os.Remove(file)
		if err != nil {
			return err
		}
	}
	return nil
}


func CopyFiles(args tasks.CopyArgs) error {
	log.Println("Starting Copy Task:", args)
	matches, err := filepath.Glob(args.Include)
	if err != nil {
		return err
	}
	for _, match := range matches {
		r, _ := regexp.Compile(args.Exclude)
		m := r.FindStringIndex(match)
		if m == nil {
			log.Println("Copying: ", match)
			err := copyFile(match, args.Destination)
			if err != nil {
				return err
			}
		} else {
			log.Println("Exclude matched, Skipping: ", filepath.Base(match))
		}
	}
	return nil
}

func MoveFiles(args tasks.MoveArgs) error {
	log.Println("Starting Move Task:", args)
	matches, err := filepath.Glob(args.Include)
	if err != nil {
		return err
	}
	for _, match := range matches {
		r, _ := regexp.Compile(args.Exclude)
		m := r.FindStringIndex(match)
		if m == nil {
			log.Println("Moving: ", match)
			err := moveFile(match, args.Destination)
			if err != nil {
				return err
			}
		} else {
			log.Println("Exclude matched, Skipping: ", filepath.Base(match))
		}
	}
	return nil
}

func RenameFiles(args tasks.RenameArgs) error {
	log.Println("RenameFiles:", args)
	matches, err := filepath.Glob(args.Include)
	if err != nil {
		return err
	}
	for _, match := range matches {
		r, _ := regexp.Compile(args.Exclude)
		m := r.FindStringIndex(match)
		if m == nil {
			dst := filepath.Dir(match) + args.Prefix + "." + args.Extension
			log.Println("Renaming: ", match)
			err := moveFile(match, dst)
			if err != nil {
				return err
			}
		} else {
			log.Println("Exclude matched, Skipping: ", filepath.Base(match))
		}
	}
	return nil
}

func moveFile(src, dst string) error {
	err := os.Rename(src, dst)
	if err != nil {
		return err
	}
	return nil
}

func copyFile(src, dst string) error {
	in, err := os.Open(src)
	if err != nil {
		return err
	}
	defer in.Close()

	out, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer out.Close()

	_, err = io.Copy(out, in)
	if err != nil {
		return err
	}
	return out.Close()
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
