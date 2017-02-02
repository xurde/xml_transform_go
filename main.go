package main

import (
	"fmt"

	"github.com/jbowtie/gokogiri/xml"
	"github.com/jbowtie/ratago/xslt"
)

var globalStyle *xml.XmlDocument
var globalStylesheet *xslt.Stylesheet

func init() {
	fmt.Println("init")
	globalStyle, _ = xml.ReadFile("test.xsl", xml.StrictParseOption)
	globalStylesheet, _ = xslt.ParseStylesheet(globalStyle, "test.xsl")
}

func doTransformation() {
	style, _ := xml.ReadFile("test.xsl", xml.StrictParseOption)
	doc, _ := xml.ReadFile("as.xml", xml.StrictParseOption)
	stylesheet, _ := xslt.ParseStylesheet(style, "test.xsl")
	stylesheet.Process(doc, xslt.StylesheetOptions{true, nil})
	// fmt.Println(output)
}

func doCachedTransformation() {
	doc, _ := xml.ReadFile("as.xml", xml.StrictParseOption)
	globalStylesheet.Process(doc, xslt.StylesheetOptions{true, nil})
}
