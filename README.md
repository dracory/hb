<img src="screenshots/logo.jpg" width="100%" />

# HTML Builder <a href="https://gitpod.io/#https://github.com/dracory/hb" style="float:right:"><img src="https://gitpod.io/button/open-in-gitpod.svg" alt="Open in Gitpod" loading="lazy"></a>

![tests](https://github.com/gouniverse/html/workflows/tests/badge.svg)
[![Go Report Card](https://goreportcard.com/badge/github.com/dracory/hb)](https://goreportcard.com/report/github.com/dracory/hb)
[![PkgGoDev](https://pkg.go.dev/badge/github.com/dracory/hb)](https://pkg.go.dev/github.com/dracory/hb)

Unpretentious short and sweet declarative HTML Builder.

Demos can be found on: https://golangui.com

## Benefits

- Valid (X)HTML
- Programmatically generate (X)HTML
- Pure GO code
- Full support, and autocomplete by your favorite editor
- No need to transfer HTML files
- No need to embed HTML files
- No need for using template files
- Full reuse of the code
- Fully eliminates the quirks of the standard template package
- Great for email templates too
- Nice looking fluent interface
- Easy to extend and build your own flavor on top
- Dynamic UI with conditions
- Conditional attributes
- Lots of demos and examples
- Works with any CSS library
- Works with any JS library
- Out of the box support for Alpine.js
- Out of the box support for HTMX

## Example

- Section containing div container (Bootstrap) with a message "Hello world"

```go
// imported to global namespace, to avoid typing hb.* each time
import . "github.com/dracory/hb"

// or, regular import, and type hb.* before using it each time
// import "github.com/dracory/hb"

// 1. Create a container div with a "Hello world" message inside
container := Div().Class("container").Text("Hello world")

// 2. Create a section with padding of 40px containing the container
section := Section().Style("padding:40px;").Child(container)

// 3. Render to HTML to display
html := section.ToHTML()
```

!!! For more examples look below in the README

## Installation

```ssh
go get -u github.com/dracory/hb
```

## Implemented Tag Shortcuts

- <b>A()</b> - shortcut for &lt;a> tag
- <b>Abbr()</b> - shortcut for &lt;abbr> tag
- <b>Address()</b> - shortcut for &lt;address> tag
- <b>Article()</b> - shortcut for &lt;article> tag
- <b>Aside()</b> - shortcut for &lt;aside> tag
- <b>BR()</b> / <b>Br()</b> - shortcut for &lt;br> tag
- <b>Button()</b> - shortcut for &lt;button> tag
- <b>Caption()</b> - shortcut for &lt;caption> tag
- <b>Code()</b> - shortcut for &lt;code> tag
- <b>Div()</b> - shortcut for &lt;div> tag
- <b>Form()</b> - shortcut for &lt;form> tag
- <b>I()</b> - shortcut for &lt;i> tag
- <b>Header()</b> - shortcut for &lt;header> tag
- <b>Heading1()</b> - shortcut for &lt;h1> tag
- <b>Heading2()</b> - shortcut for &lt;h2> tag
- <b>Heading3()</b> - shortcut for &lt;h3> tag
- <b>Heading4()</b> - shortcut for &lt;h4> tag
- <b>Heading5()</b> - shortcut for &lt;h5> tag
- <b>Heading6()</b> - shortcut for &lt;h6> tag
- <b>Hyperlink()</b> - shortcut for &lt;a> tag
- <b>HR()</b> / <b>Hr()</b> - shortcut for &lt;hr> tag
- <b>HTML(html string)</b> - creates empty tag with the HTML content
- <b>Image()</b> - shortcut for &lt;img> tag
- <b>Input()</b> - shortcut for &lt;input> tag
- <b>LI()</b> / <b>Li()</b> - shortcut for &lt;li> tag
- <b>Label()</b> - shortcut for &lt;label> tag
- <b>Nav()</b> - shortcut for &lt;nav> tag
- <b>Navbar()</b> - shortcut for &lt;navbar> tag
- <b>OL()</b> / <b>Ol()</b> - shortcut for &lt;ol> tag
- <b>Option()</b> - shortcut for &lt;option> tag
- <b>Paragraph()</b> - shortcut for &lt;p> tag
- <b>PRE()</b> - shortcut for &lt;pre> tag
- <b>Script()</b> - shortcut for &lt;script> tag
- <b>ScriptURL()</b> - shortcut for &lt;script src="{SRC}"> tag
- <b>Select()</b> - shortcut for &lt;select> tag
- <b>Small()</b> - shortcut for &lt;small> tag
- <b>Span()</b> - shortcut for &lt;span> tag
- <b>Style()</b> - shortcut for &lt;style> tag
- <b>StyleURL()</b> - shortcut for &lt;link> tag
- <b>Section()</b> - shortcut for &lt;section> tag
- <b>Sub()</b> - shortcut for &lt;sub> tag
- <b>Sup()</b> - shortcut for &lt;sup> tag
- <b>NewTag(tagName string)</b> - for custom tags
- <b>Table()</b> - shortcut for &lt;table> tag
- <b>Tbody()</b> / <b>TBody()</b> - shortcut for &lt;tbody> tag
- <b>TD()</b> / <b>Td()</b> - shortcut for &lt;td> tag
- <b>Template()</b> - shortcut for &lt;template> tag
- <b>Text(text string)</b> - creates empty tag with the escaped text content
- <b>TextArea()</b> / <b>Textarea()</b> - shortcut for &lt;textarea> tag
- <b>TH()</b> / <b>Th()</b> - shortcut for &lt;th> tag
- <b>Thead()</b> / <b>THead()</b> - shortcut for &lt;thead> tag
- <b>Tfoot()</b> / <b>TFoot()</b> - shortcut for &lt;tfoot> tag
- <b>TR()</b> / <b>Tr()</b> - shortcut for &lt;tr> tag
- <b>UL()</b> / <b>Ul()</b> - shortcut for &lt;ul> tag
- <b>Webpage()</b> - full HTML page withe head, body, meta, styles and scripts
- <b>Wrap()</b> - convenience method to taglessly (without a root wrapping element) group elements together

Examples can be found on: https://golangui.com

## Tag Methods

- <b>Action(action string)</b> - shortcut to add an "action" attribute
- <b>Aria(key, value string)</b> - shortcut for "aria-" attribute
- <b>Attr(key, value string)</b> - shortcut for SetAttribute
- <b>AttrIf(condition bool, key, value string)</b> - conditional setting of attribute
- <b>AttrIfF(condition bool, key string, valueFunc func() string)</b> - conditional setting of attribute using function
- <b>AttrIfElse(condition bool, key, valueIf string, valueElse string)</b> - conditional setting of attribute
- <b>Attrs(map[string]string)</b> - shortcut for setting multiple attributes
- <b>AttrsIf(condition bool, attrs map[string]string)</b> - conditional setting of attributes
- <b>AttrsIfF(condition bool, attrsFunc func() map[string]string)</b> - conditional setting of attributes using function
- <b>AttrsIfElse(condition, attrsIf map[string]string, attrsElse map[string]string)</b> - conditional setting of attributes
- <b>AddChild(tag Tag)</b> - adds a child element
- <b>AddChildren(tag []Tag)</b> - adds an array of child elements
- <b>AddClass(className string)</b> - adds a class name to the "class" attribute
- <b>AddHTML(html string)</b> - adds HTML content to the element
- <b>AddStyle(style string)</b> - adds a style to the "style" attribute
- <b>AddText(text string)</b> - adds escaped text content to the element
- <b>Child(tag Tag)</b> - shortcut for AddChild
- <b>ChildIf(condition bool, tag Tag)</b> - conditional adding of a child
- <b>ChildIfF(condition bool, childFunc func() \*Tag)</b> - conditional adding of a child using function
- <b>ChildIfElse(condition bool, childIf Tag, childElse Tag)</b> - conditional adding of a child
- <b>Children(children []Tag)</b> - shortcut for AddChildren
- <b>ChildrenIf(condition bool, children []Tag)</b> - conditional adding of children
- <b>ChildrenIfF(condition bool, childrenFunc func() []\*Tag)</b> - conditional adding of children using function
- <b>ChildrenIfElse(condition bool, childrenIf []Tag, childrenElse []Tag)</b> - conditional adding of a children
- <b>Class(className string)</b> - shortcut for AddClass
- <b>ClassIf(condition bool, className string)</b> - conditional adding of a class
- <b>ClassIfElse(condition bool, classNameIf string, classNameElse string)</b> - conditional adding of a class
- <b>Data(name string, value string)</b> - shortcut to add a "data-" attribute
- <b>DataIf(condition bool, key, value string)</b> - conditional setting of attribute
- <b>DataIfElse(condition bool, key, valueIf string, valueElse string)</b> - conditional setting of attribute
- <b>Enctype(enctype string)</b> - shortcut to add an "enctype" attribute
- <b>HasClass(className string)</b> - checks if the class is available
- <b>Href(href string)</b> - shortcut to add an "href" attribute
- <b>HTML(html string)</b> - shortcut for AddHTML
- <b>ID(className string)</b> - shortcut to add an "id" attribute
- <b>GetAttribute(key string) string</b>
- <b>Method(method string)</b> - shortcut to add a "method" attribute
- <b>Name(name string)</b> - shortcut to add a "name" attribute
- <b>Placeholder(placeholder string)</b> - shortcut to add a "placeholder" attribute
- <b>OnBlur(js string)</b> - shortcut to add an "onblur" attribute
- <b>OnChange(js string)</b> - shortcut to add an "onchange" attribute
- <b>OnClick(js string)</b> - shortcut to add an "onclick" attribute
- <b>OnDblClick(js string)</b> - shortcut to add an "ondblclick" attribute
- <b>OnFocus(js string)</b> - shortcut to add an "onfocus" attribute
- <b>OnKeyDown(js string)</b> - shortcut to add an "onkeydown" attribute
- <b>OnKeyPress(js string)</b> - shortcut to add an "onkeypress" attribute
- <b>OnKeyUp(js string)</b> - shortcut to add an "onkeyup" attribute
- <b>OnMouseDown(js string)</b> - shortcut to add an "onmousedown" attribute
- <b>OnMouseOut(js string)</b> - shortcut to add an "onmouseout" attribute
- <b>OnMouseUp(js string)</b> - shortcut to add an "onmouseup" attribute
- <b>OnSubmit(js string)</b> - shortcut to add an "onsubmit" attribute
- <b>ReadOnly(isReadOnly bool)</b> - shortcut to add a "readonly" attribute
- <b>Rel(rel string)</b> - shortcut to add a "rel" attribute
- <b>Role(role string)</b> - shortcut to add a "role" attribute
- <b>Required(isRequired bool)</b> - shortcut to add a "required" attribute
- <b>Selected(isSelected bool)</b> - shortcut to add a "selected" attribute
- <b>SetAttribute(key, value string)</b> - sets an attribute (i.e. id, class, etc)
- <b>Src(src string)</b> - shortcut to add a "src" attribute
- <b>Style(style string)</b> - shortcut to add a "style" attribute
- <b>StyleIf(condition bool, style string)</b> - conditional adding of a style
- <b>StyleIfElse(condition bool, styleIf string, styleElse string)</b> - conditional adding of a style
- <b>Target(target string)</b> - shortcut to add a "target" attribute
- <b>TargetIf(condition bool, target string)</b> - conditional adding of a "target" attribute
- <b>Text(html string)</b> - shortcut for AddText
- <b>TextIf(condition bool, text string)</b> - adds escaped text if a condition is met
- <b>TextIfElse(condition bool, textIf string, textElse string)</b> - adds escaped text if a condition is met
- <b>Title(title string)</b> - shortcut for setting the "title" attribute
- <b>TitleIf(condition bool, title string)</b> - sets the title if a condition is met
- <b>ToHTML() string</b> - outputs HTML code
- <b>Type(target string)</b> - shortcut to add a "type" attribute
- <b>TypeIf(condition bool, target string)</b> - conditional setting of "type" attribute
- <b>Value(value string)</b> - shortcut to add a "value" attribute
- <b>ValueIf(condition bool, value string)</b> - conditional setting of "value" attribute

Examples can be found on: https://golangui.com

## Utility

- <b>ToTags[T any](items []T, callback func(item T, index int) *Tag) []*Tag</b> - transforms a slice of anything to a slice of tags
- If(condition bool, trueTag *Tag) *Tag
- IfF(condition bool, trueFunc func() *Tag) *Tag
- IfElse(condition bool, trueTag *Tag, falseTag *Tag) \*Tag
- IfFElseF(condition bool, trueFunc func() *Tag, falseFunc func() *Tag) \*Tag
- Ternary(condition bool, trueTag *Tag, falseTag *Tag) \*Tag
- TernaryF(condition bool, trueFunc func() *Tag, falseFunc func() *Tag) \*Tag

## Tag HTMX Attributes

HTMX (https://htmx.org/reference/) is a great match for Golang, therefore here is a shortcut for working with HTMX.

- Hx(name string, value string) - shortcut for setting an HTMX attribute
- HxConfirm(value string)
- HxDelete(value string)
- HxGet(value string)
- HxInclude(value string)
- HxIndicator(value string)
- HxOn(name string, value string)
- HxPatch(value string)
- HxPost(value string)
- HxPut(value string)
- HxSelect(value string)
- HxSelectOob(value string)
- HxSync(value string)
- HxSwap(value string)
- HxSwapOob(value string)
- HxTarget(value string)
- HxTrigger(value string)
- HxVals(value string)
- HxVars(value string)

Examples can be found on: https://golangui.com

## Tag Alpine.js Attributes

Alpine.js (https://alpinejs.dev/) is a great match for Golang, therefore here is a shortcut for working with HTMX.

- X(name string, value string) - shortcut for setting an AlpineJS attribute
- XBind(name string, value string)
- XOn(name string, value string)

Examples can be found on: https://golangui.com

## Webpage Methods

- <b>AddChild(child \*Tag)</b>
- <b>SetFavicon(favicon string)</b>
- <b>SetTitle(title string)</b>
- <b>AddScripts(scripts []string)</b>
- <b>AddScript(script string)</b>
- <b>AddScriptURLs(scriptURLs []string)</b>
- <b>AddScriptURL(scriptURL string)</b>
- <b>AddStyle(style string)</b>
- <b>AddStyles(styles []string)</b>
- <b>AddStyleURL(styleURL string)</b>
- <b>AddStyleURLs(styleURLs []string)</b>

## Border Layout Methods

- <b>NewBorderLayout()</b> - shortcut to quickly create a border layout with top, bottom, left, right and center slots (see example below how to use it)
- <b>AddTop(tag \*Tag, alignHorizontal string, alignVertical string)</b>
- <b>AddBottom(tag \*Tag, alignHorizontal string, alignVertical string)</b>
- <b>AddLeft(tag \*Tag, alignHorizontal string, alignVertical string)</b>
- <b>AddRight(tag \*Tag, alignHorizontal string, alignVertical string)</b>
- <b>AddCenter(tag \*Tag, alignHorizontal string, alignVertical string)</b>

Usage example:

```go
bl := NewBorderLayout()
	bl.AddTop(NewSpan().Text("TOP"), BORDER_LAYOUT_ALIGN_CENTER, BORDER_LAYOUT_ALIGN_MIDDLE)
	bl.AddCenter(NewSpan().Text("CENTER"), BORDER_LAYOUT_ALIGN_CENTER, BORDER_LAYOUT_ALIGN_MIDDLE)
	bl.AddBottom(NewSpan().Text("BOTTOM"), BORDER_LAYOUT_ALIGN_CENTER, BORDER_LAYOUT_ALIGN_MIDDLE)
	bl.AddLeft(NewSpan().Text("LEFT"), BORDER_LAYOUT_ALIGN_CENTER, BORDER_LAYOUT_ALIGN_MIDDLE)
	bl.AddRight(NewSpan().Text("RIGHT"), BORDER_LAYOUT_ALIGN_CENTER, BORDER_LAYOUT_ALIGN_MIDDLE)
blHtml := bl.ToHTML()
```

## Swal

- <b>NewSwal</b> - shortcut to quickly add Sweetalert 2 dialogs (note requires adding the Sweetalert JS library). See more at: https://sweetalert2.com/download

Usage example in a form:

```go
form.
	ChildIf(data.formErrorMessage != "", hb.Swal(swal.SwalOptions{
		Icon:              "error",
		Title:             "Oops...",
		Text:              data.formErrorMessage,
		ShowCancelButton:  false,
		ConfirmButtonText: "OK",
		ConfirmCallback:   "window.location.href = window.location.href",
	})).
	ChildIf(data.formSuccessMessage != "", hb.Swal(swal.SwalOptions{
		Icon:              "success",
		Title:             "Saved",
		Text:              data.formSuccessMessage,
		ShowCancelButton:  false,
		ConfirmButtonText: "OK",
		ConfirmCallback:   "window.location.href = window.location.href",
	}))
```

## Working with Raw Tags

```go
tag := &Tag{
	TagName: "custom-element",
}
tag.toHTML()
```

## Safe Escaping HTML

For safeguarding HTML, always use the .Text(text) method of the tags, it will automatically escape any HTML tags in the output.

Using the .HTML(html) method of the tags, will output the raw HTML, and leaves you vulnerable to XSS attacks.

You can escape the HTML yourself by using the EscapeString method from the standard HTML library
Link with example: https://golang.org/pkg/html/#EscapeString

## Examples

- Bootstrap login form

```go
// Elements for the form
header := Heading3().text("Please sign in").Style("margin:0px;")
emailLabel := Label().Text("E-mail Address")
emailInput := Input().Class("form-control").Name("email").Placeholder("Enter e-mail address")
emailFormGroup := Div().Class("form-group").Child(emailLabel).Child(emailInput)
passwordLabel := Label().Child(hb.Text("Password"))
passwordInput := Input().Class("form-control").Name("password").Type("password").Placeholder("Enter password")
passwordFormGroup := Div().Class("form-group").Child(passwordLabel).Child(passwordInput)
buttonLogin := Button().Class("btn btn-lg btn-success btn-block").Text("Login")
buttonRegister := Hyperlink().Class("btn btn-lg btn-info float-left").Text("Register").Href("auth/register")
buttonForgotPassword := Hyperlink().Class("btn btn-lg btn-warning float-right").Text("Forgot password?").Href("auth/password-restore")

// Add elements in a card
cardHeader := Div().
	Class("card-header").
	Child(header)

cardBody := Div().
	Class("card-body").
	Children([]*hb.Tag{
		emailFormGroup,
		passwordFormGroup,
		buttonLogin,
	})

cardFooter := Div().
	Class("card-footer").
	Children([]*hb.Tag{
		buttonRegister,
		buttonForgotPassword,
	})

card := Div().
	Class("card card-default").
	Style("margin:0 auto;max-width: 360px;").
	Children([]*hb.Tag{
		cardHeader,
		cardBody,
		cardFooter
	})

// Convert to HTML to display
html := card.ToHTML()
```

Result:

<img src="https://github.com/gouniverse/html/blob/master/screnshots/html-builder-form-login.png" />

- Webpage with title, favicon, Font Awesome icons 4.7.0, jQuery 3.2.1, and Bootstrap 4.5

```go
// 1. Webpage Title
title := "Title"

// 2. Webpage Favicon
favicon := "data:image/x-icon;base64,AAABAAEAEBAQAAEABAAoAQAAFgAAACgAAAAQAAAAIAAAAAEABAAAAAAAgAAAAAAAAAAAAAAAEAAAAAAAAABNTU0AVKH/AOPj4wDExMQAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAACIiAREQEREAIiIBERAREQAiIgIiICIiACIiAiIgIiIAMzMDMzAzMwAzMwMzMDMzACIiAiIgIiIAIiICIiAiIgAzMwMzMDMzADMzAzMwMzMAIiICIiAiIgAiIgIiICIiAAAAAAAAAAAAIiICIiAiIgAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA"

// 3. Webpage
webpage := NewWebpage().
	SetTitle(title).
	SetFavicon(favicon).
	AddStyleURLs([]string{
		"https://maxcdn.bootstrapcdn.com/font-awesome/4.7.0/css/font-awesome.min.css",
		"https://stackpath.bootstrapcdn.com/bootstrap/4.5.0/css/bootstrap.min.css",
	}).
	AddScriptURLs([]string{
		"https://code.jquery.com/jquery-3.2.1.min.js",
		"https://stackpath.bootstrapcdn.com/bootstrap/4.5.0/js/bootstrap.bundle.min.js",
	}).
	AddStyle(`html,body{height:100%;font-family: Ubuntu, sans-serif;}`).AddChild(NewDiv().Text("Hello"))
```

- Wrap webpage in a function to be reused as a master template

```go
// Webpage returns the webpage template for the website
func Webpage(title, content string) *hb.Webpage {
	faviconImgCms := `data:image/x-icon;base64,AAABAAEAEBAQAAEABAAoAQAAFgAAACgAAAAQAAAAIAAAAAEABAAAAAAAgAAAAAAAAAAAAAAAEAAAAAAAAAAAAAAAmzKzAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAABEQEAAQERAAEAAQABAAEAAQABAQEBEQABAAEREQEAAAERARARAREAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAD//wAA//8AAP//AAD//wAA//8AAP//AAD//wAAi6MAALu7AAC6owAAuC8AAIkjAAD//wAA//8AAP//AAD//wAA`
	webpage := hb.Webpage()
	webpage.SetTitle(title)
	webpage.SetFavicon(faviconImgCms)

	webpage.AddStyleURLs([]string{
		"https://cdn.jsdelivr.net/npm/bootstrap@5.0.0-beta1/dist/css/bootstrap.min.css",
	})
	webpage.AddScriptURLs([]string{
		"https://cdn.jsdelivr.net/npm/bootstrap@5.0.0-beta1/dist/js/bootstrap.bundle.min.js",
		"http://code.jquery.com/jquery-3.5.1.min.js",
		"https://unpkg.com/vue@next",
		"https://cdn.jsdelivr.net/npm/sweetalert2@9",
	})
	webpage.AddScripts([]string{})
	webpage.AddStyle(`html,body{height:100%;font-family: Ubuntu, sans-serif;}`)
	webpage.AddStyle(`body {
		font-family: "Nunito", sans-serif;
		font-size: 0.9rem;
		font-weight: 400;
		line-height: 1.6;
		color: #212529;
		text-align: left;
		background-color: #f8fafc;
	}`)
	webpage.AddChild(hb.HTML(content))
	return webpage
}

// Generate HTML
html := webpage("Home", "Hello world").ToHTML()
```

## HTMX Example

Here is an HTMX example which submits the content of all the inputs
in a div to the server (no need to be in a form) and then replaces
the content of the webpage with the result.

```go
input := NewButton().
		Text("Submit")
		Hx("post", "http://test.com").
		Hx("include", "#DivID").
		Hx("target", "#PageID").
		Hx("swap", "outerHTML")
```

## How to Add a Redirection?

```go
webpage.Meta(hb.Meta().Attr("http-equiv", "refresh").Attr("content", "2; url = https://www.yahoo.com"))
```

## Stargazers over time

[![Stargazers over time](https://starchart.cc/dracory/hb.svg)](https://starchart.cc/dracory/hb)

## Similar

- https://github.com/goradd/html5tag - option to have string or stream

## Changelog

2025.07.02 - Added Small() shortcut

2025.02.22 - Updated README.md with latest changes.

2025.01.19 - Added additional Sweetalert options to redirect after delay

2024.11.26 - Added Aria, Readonly, Required, Selected

2024.10.05 - Added HxIndicator

2024.10.01 - Added additional shortcuts for more concise declaration

2023.12.17 - Added ToTags, Ternary, Text

2023.12.10 - Added Swal method for quickly adding Sweetalert

2023.09.16 - Added special case for empty "value" attribute

2023.07.26 - Added NewCaption function and Alpine.js, HTMX functions

2023.07.02 - Added NewHeader function

2023.06.09 - Added functional conditions AttrIfF, AttrsIfF, ChildIfF, ChildrenIfF

2023.05.08 - Added Alt attribute shortcut

2023.05.01 - Added NewWrap function

2023.04.27 - Added OnDblClick, OnInput, OnKeyPress, OnMouseDown, OnMouseUp, and conditionals for data

2023.04.25 - Added Placeholder, and conditionals for attributes

2023.04.15 - Added AddStyle, Src, and conditionals for style

2023.04.14 - Added Type attribute shortcut, conditionals for children and class names

2023.03.26 - Added Hx attribute shortcut for working with HTMX

2023.03.26 - Added OnBlur, OnChange, OnFocus, OnKeyDown, OnKeyUp, attribute shortcuts

2023.03.26 - Added Enctype, Href, Method, Name, target, Value attribute shortcuts

2023.01.22 - Added shortcut for &lt;footer> tag

2023.01.14 - Flow pattern applied to BorderLayout methods

2022.09.07 - Added Child and Children shortcuts

2022.08.29 - Added default favicon to Webpage to fix 404 if missing

2022.01.07 - Added Attrs shortcut for setting multiple attributes

2021.07.30 - Added shortcut for &lt;hr> tag

2021.03.20 - Renamed package name to hb to not conflict/confuse with the standard html package

2021.03.18 - Added shortcuts for &lt;template> tag

2021.02.26 - Added shortcuts for &lt;sub>, &lt;sup> tags

2021.01.03 - Added example for webpage layout, and screenshot

2020.12.28 - Added shortcuts for &lt;code>, &lt;pre> tags, shortcuts sorted alphabetically

2020.12.27 - Added shortcuts for &lt;table>, &lt;thead>, &lt;tbody>, &lt;tr>, &lt;th>, &lt;td> tags

2020.12.26 - Fix for attribute escapes, added tests

2020.09.16 - Added shortcuts for &lt;nav>, &lt;navbar>, &lt;ol>, &lt;ul>, &lt;li> tags

2020.06.16 - Initial commit

## Aternatives

- Stevelacy Daz (https://github.com/stevelacy/daz) | Last update: 19 Apr 2023

- Extemplate (https://github.com/dannyvankooten/extemplate) | Last update: 15 Jun 2021

- ThePlant HTMLGO (https://github.com/theplant/htmlgo) | Last update: 18 Sep 2021

- Julvo HTMLGO (https://github.com/julvo/htmlgo) | Last update: 5 May 2020

- Go Components (https://github.com/maragudk/gomponents)

- Uberswe HTML (https://github.com/uberswe/html) | Last update: 24 Feb 2021

- Elem Go (https://github.com/chasefleming/elem-go)

- HTMGO (https://github.com/maddalax/htmgo)

- Forms from Structs (https://github.com/joncalhoun/form)
  Still writ

- HTMLX (https://github.com/mdigger/htmlx) | Last update: 20 Jun 2021
  An extension to the standard html package. Not lots of documentation

- HTML as functions (https://github.com/jpincas/htmlfunc) | Last update: 23 Jan 2023
  Interesting library, following Elm. Not lots of documentation

- Templ (https://github.com/a-h/templ) | 30 Sep 2023
  Building templates with JSX like syntax. Requires installing 3-rd party binaries, and additional compilation step before being usable.

- GoHTMX (https://gitlab.com/go-htmx/go-htmx)

<picture>
  <source
    media="(prefers-color-scheme: dark)"
    srcset="
      https://api.star-history.com/svg?repos=dracory/hb&type=Date&theme=dark
    "
  />
  <source
    media="(prefers-color-scheme: light)"
    srcset="
      https://api.star-history.com/svg?repos=dracory/hb&type=Date
    "
  />
  <img
    alt="Star History Chart"
    src="https://api.star-history.com/svg?repos=dracory/hb&type=Date"
  />
</picture>
