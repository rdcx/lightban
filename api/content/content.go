package content

import (
	"fmt"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v2"
)

type File struct {
	Name     string `yaml:"name" json:"name"`
	Content  string `yaml:"content" json:"content"`
	Readonly bool   `yaml:"readonly" json:"readonly"`
}

type Dialog struct {
	KeyName string `yaml:"key_name" json:"key_name"`
	Trigger string `yaml:"trigger" json:"trigger"`
}

type LessonContent struct {
	ID          uint     `json:"id"`
	KeyName     string   `yaml:"key_name" json:"key_name"`
	Slug        string   `yaml:"slug" json:"slug"`
	Name        string   `yaml:"name" json:"name"`
	Description string   `yaml:"description" json:"description"`
	Position    int      `yaml:"position" json:"position"`
	Files       []File   `json:"files"`
	Dialogs     []Dialog `yaml:"dialogs" json:"dialogs"`
	Asserts     []string `yaml:"asserts" json:"asserts"`
	NextLesson  string   `yaml:"next_lesson" json:"next_lesson"`
}

func GetLessonContent(baseDir, lang, course, lesson string) LessonContent {
	// read all of the files in the lesson directory
	// and return them as a string

	dir := fmt.Sprintf(baseDir+"/%s/%s/%s/files", lang, course, lesson)

	files := []File{}

	filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {

		if err != nil {
			return err
		}

		if info.IsDir() {
			return nil
		}

		data, errRead := os.ReadFile(path)

		if errRead != nil {
			return err
		}

		// file name is the ltrim of the path
		// relative to the lesson directory

		rel, err := filepath.Rel(dir, path)
		if err != nil {
			return err
		}

		files = append(files, File{
			Name:     rel,
			Content:  string(data),
			Readonly: false,
		})

		return nil
	})

	// read the lesson yaml file
	// and return the dialog, helper text, and assert

	lessonFile := fmt.Sprintf(baseDir+"/%s/%s/%s/lesson.yaml", lang, course, lesson)
	data, err := os.ReadFile(lessonFile)

	if err != nil {
		panic(err)
	}

	lc := LessonContent{
		Files: files,
	}

	err = yaml.Unmarshal(data, &lc)

	if err != nil {
		panic(err)
	}

	return lc
}
