package template

import (
	"fmt"
	"regexp"
	"strings"
	"text/template"
)

// Render ...
func Render(tmpl string, vals map[string]interface{}, strict ...bool) (res string, err error) {
	// Basically, what we do here is start with an empty parent template and then
	// build up a list of templates -- one for each file. Once all of the templates
	// have been parsed, we loop through again and execute every template.
	//
	// The idea with this process is to make it possible for more complex templates
	// to share common blocks, but to make the entire thing feel like a file-based
	// template engine.
	defer func() {
		if p := recover(); p != nil {
			err = fmt.Errorf("rendering template failed: %v", p)
		}
	}()
	t := template.New("gotpl")
	if len(strict) > 0 && strict[0] {
		t.Option("missingkey=error")
	} else {
		// Not that zero will attempt to add default values for types it knows,
		// but will still emit <no value> for others. We mitigate that later.
		t.Option("missingkey=zero")
	}
	t.Funcs(funcMap())
	if _, err := t.Parse(tmpl); err != nil {
		fmt.Println("err", err)
		return "", cleanupParseError("gotpl", err)
	}
	var buf strings.Builder
	if err := t.ExecuteTemplate(&buf, "gotpl", vals); err != nil {
		fmt.Println("err", err)
		return "", cleanupExecError("gotpl", err)
	}
	return strings.ReplaceAll(buf.String(), "<no value>", ""), nil
}

const warnStartDelim = "HELM_ERR_START"
const warnEndDelim = "HELM_ERR_END"

var warnRegex = regexp.MustCompile(warnStartDelim + `((?s).*)` + warnEndDelim)

func cleanupExecError(filename string, err error) error {
	if _, isExecError := err.(template.ExecError); !isExecError {
		return err
	}

	tokens := strings.SplitN(err.Error(), ": ", 3)
	if len(tokens) != 3 {
		// This might happen if a non-templating error occurs
		return fmt.Errorf("execution error in (%s): %s", filename, err)
	}

	// The first token is "template"
	// The second token is either "filename:lineno" or "filename:lineNo:columnNo"
	location := tokens[1]

	parts := warnRegex.FindStringSubmatch(tokens[2])
	if len(parts) >= 2 {
		return fmt.Errorf("execution error at (%s): %s", string(location), parts[1])
	}

	return err
}

func cleanupParseError(filename string, err error) error {
	tokens := strings.Split(err.Error(), ": ")
	if len(tokens) == 1 {
		// This might happen if a non-templating error occurs
		return fmt.Errorf("parse error in (%s): %s", filename, err)
	}
	// The first token is "template"
	// The second token is either "filename:lineno" or "filename:lineNo:columnNo"
	location := tokens[1]
	// The remaining tokens make up a stacktrace-like chain, ending with the relevant error
	errMsg := tokens[len(tokens)-1]
	return fmt.Errorf("parse error at (%s): %s", string(location), errMsg)
}
