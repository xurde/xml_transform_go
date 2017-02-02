package main

import (
	"fmt"

	"github.com/jbowtie/gokogiri/xml"
	"github.com/jbowtie/ratago/xslt"
)

var globalStyle, globalDoc *xml.XmlDocument
var globalStylesheet *xslt.Stylesheet

func init() {
	fmt.Println("init")
	globalStyle, _ = xml.ReadFile("test.xsl", xml.StrictParseOption)
	globalStylesheet, _ = xslt.ParseStylesheet(globalStyle, "test.xsl")
	globalDoc, _ = xml.ReadFile("as.xml", xml.StrictParseOption)
}

func doTransformation() {
	style, _ := xml.ReadFile("test.xsl", xml.StrictParseOption)
	doc, _ := xml.ReadFile("as.xml", xml.StrictParseOption)
	stylesheet, _ := xslt.ParseStylesheet(style, "test.xsl")
	stylesheet.Process(doc, xslt.StylesheetOptions{true, nil})
	// fmt.Println(output)
}

func doCachedTransformation() {
	globalStylesheet.Process(globalDoc, xslt.StylesheetOptions{true, nil})
}
