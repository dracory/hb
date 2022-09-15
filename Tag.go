package hb

import (
	"bytes"
	"sort"
	"strings"
)

// Addslashes addslashes()
func addslashes(str string) string {
	var buf bytes.Buffer
	for _, char := range str {
		switch char {
		//case '\'', '"', '\\':
		case '"', '\\':
			buf.WriteRune('\\')
		}
		buf.WriteRune(char)
	}
	return buf.String()
}

func inArrayString(arr []string, str string) bool {
	for _, a := range arr {
		if a == str {
			return true
		}
	}
	return false
}

// TagInterface represents an HTML tag interface
type TagInterface interface {
	ToHTML() string
}

// Tag represents an HTML tag
type Tag struct {
	TagInterface
	TagName       string
	TagContent    string
	TagAttributes map[string]string
	TagChildren   []*Tag
}

// AddClass adds a new class name to the tag attribute list.
func (t *Tag) AddClass(className string) *Tag {
	classNames := t.GetAttribute("class")
	classNamesArray := strings.Split(classNames, " ")
	classNamesArray = append(classNamesArray, className)
	classNames = strings.Join(classNamesArray, " ")
	classNames = strings.Trim(classNames, " ")
	return t.SetAttribute("class", classNames)
}

// Attr shortcut for SetAttribute
func (t *Tag) Attr(key, value string) *Tag {
	return t.SetAttribute(key, value)
}

// Attrs shortcut for setting multiple attributes
func (t *Tag) Attrs(attrs map[string]string) *Tag {
	for key, value := range attrs {
		t.SetAttribute(key, value)
	}
	return t
}

// AddChild adds a new child tag to this tag
func (t *Tag) AddChild(child *Tag) *Tag {
	t.TagChildren = append(t.TagChildren, child)
	return t
}

// AddChildren adds an array of child tags to this tag
func (t *Tag) AddChildren(children []*Tag) *Tag {
	for _, child := range children {
		t.AddChild(child)
	}
	return t
}

// AddHTML adds an HTML as child tags to this tag
func (t *Tag) AddHTML(html string) *Tag {
	t.AddChild(NewHTML(html))
	return t
}

// AddChild shortcut for AddChild
func (t *Tag) Child(child *Tag) *Tag {
	return t.AddChild(child)
}

// Children shortcut for AddChildren
func (t *Tag) Children(children []*Tag) *Tag {
	return t.AddChildren(children)
}

// Class shortcut for setting the "class" attribute
func (t *Tag) Class(clasName string) *Tag {
	return t.AddClass(clasName)
}

// GetAttribute returns the value of an attribute
func (t *Tag) GetAttribute(key string) string {
	if t.TagAttributes == nil {
		t.TagAttributes = map[string]string{}
	}
	return t.TagAttributes[key]
}

// HasClass returns true if the tag has a class with the specified name.
func (t *Tag) HasClass(className string) bool {
	classNames := t.GetAttribute("class")
	classNamesArray := strings.Split(classNames, " ")
	return inArrayString(classNamesArray, className)
}

// HTML shortcut for AddHTML
func (t *Tag) HTML(html string) *Tag {
	return t.AddHTML(html)
}

// ID shortcut for setting the "id" attribute
func (t *Tag) ID(id string) *Tag {
	return t.SetAttribute("id", id)
}

// OnClick shortcut for setting the "onclick" attribute
func (t *Tag) OnClick(js string) *Tag {
	return t.SetAttribute("onclick", js)
}

// SetAttribute sets the valua of an attribute
func (t *Tag) SetAttribute(key, value string) *Tag {
	if t.TagAttributes == nil {
		t.TagAttributes = map[string]string{}
	}
	t.TagAttributes[key] = value
	return t
}

// Style shortcut for setting the "style" attribute
func (t *Tag) Style(style string) *Tag {
	return t.SetAttribute("style", style)
}

// ToHTML returns HTML from Node
func (t *Tag) ToHTML() string {
	shortTags := []string{
		"br",
		"img",
		"input",
		"link",
		"meta",
	}

	isShortTag := inArrayString(shortTags, t.TagName)

	tagStart := `<` + t.TagName + t.attrToString() + `>`
	tagEnd := `</` + t.TagName + `>`
	if t.TagName == "" {
		tagStart = ""
		tagEnd = ""
	}
	if isShortTag {
		tagStart := `<` + t.TagName + t.attrToString() + ` />`
		tagEnd = ""
	}
	return tagStart + t.TagContent + t.childrenToString() + tagEnd
}

func (t Tag) attrToString() string {
	attrString := ""

	keys := make([]string, 0, len(t.TagAttributes))
	for k := range t.TagAttributes {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	for _, key := range keys {
		value := t.TagAttributes[key]
		//for key, value := range t.TagAttributes {
		if strings.Trim(value, " ") == "" {
			continue
		}
		attrString += ` ` + key + `="` + addslashes(value) + `"`
	}

	if attrString != "" {
		attrString = " " + attrString
	}

	if len(attrString) < 1 {
		return ""
	}

	return " " + strings.Trim(attrString, " ")
}

func (t Tag) childrenToString() string {
	childrenString := ""

	for _, child := range t.TagChildren {
		childrenString += child.ToHTML()
	}

	return childrenString
}
