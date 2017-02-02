package main

import (
	"github.com/jbowtie/gokogiri/xml"
	"github.com/jbowtie/ratago/xslt"
)

func doTransformation() {

	style, _ := xml.ReadFile("test.xsl", xml.StrictParseOption)

	doc, _ := xml.ReadFile("as.xml", xml.StrictParseOption)

	stylesheet, _ := xslt.ParseStylesheet(style, "test.xsl")
	// fmt.Println(stylesheet, err)

	stylesheet.Process(doc, xslt.StylesheetOptions{true, nil})

	// fmt.Println(output)
}
