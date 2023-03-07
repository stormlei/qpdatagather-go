package api

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"golang.org/x/net/html/charset"
	"golang.org/x/text/encoding/unicode"
	"golang.org/x/text/transform"
	"io"
	"os"
	"testing"
)

type dataFormat struct {
	Measure struct {
		Type string `xml:"type,attr"`
		SM   struct {
			R struct {
				List struct {
					Num string `xml:"NUM"`
					Cd  string `xml:"CD"`
					Avg string `xml:"AVG"`
					Sd  string `xml:"SD"`
					Cv  string `xml:"CV"`
					Max string `xml:"MAX"`
					Min string `xml:"MIN"`
					Hex string `xml:"HEX"`
					Ct  string `xml:"CT"`
				} `xml:"List"`
			} `xml:"R"`
			L struct {
				List struct {
					Num string `xml:"NUM"`
					Cd  string `xml:"CD"`
					Avg string `xml:"AVG"`
					Sd  string `xml:"SD"`
					Cv  string `xml:"CV"`
					Max string `xml:"MAX"`
					Min string `xml:"MIN"`
					Hex string `xml:"HEX"`
					Ct  string `xml:"CT"`
				} `xml:"List"`
			} `xml:"L"`
		} `xml:"SM"`
	} `xml:"Measure"`
}

func TestCem530XML(t *testing.T) {
	dataSlice, err := readFileUTF16("../xml/cem530.xml")
	if err != nil {
		fmt.Printf("%s", err)
		os.Exit(1)
	}

	var data dataFormat
	decoder := xml.NewDecoder(bytes.NewBuffer(dataSlice))
	decoder.CharsetReader = func(charsetT string, input io.Reader) (io.Reader, error) {
		return charset.NewReader(input, charsetT)
	}
	err = decoder.Decode(&data)
	if err != nil {
		fmt.Println("Error unmarshalling from XML", err)
		return
	}

	fmt.Println(data.Measure.SM)

}

func readFileUTF16(filename string) ([]byte, error) {
	// Read the file into a []byte:
	raw, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	// Make an tranformer that converts MS-Win default to UTF8:
	win16be := unicode.UTF16(unicode.BigEndian, unicode.IgnoreBOM)
	// Make a transformer that is like win16be, but abides by BOM:
	utf16bom := unicode.BOMOverride(win16be.NewDecoder())

	// Make a Reader that uses utf16bom:
	unicodeReader := transform.NewReader(bytes.NewReader(raw), utf16bom)

	// decode and print:
	decoded, err := io.ReadAll(unicodeReader)
	return decoded, err
}
