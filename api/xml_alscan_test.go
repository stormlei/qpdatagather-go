package api

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"golang.org/x/net/html/charset"
	"io"
	"os"
	"testing"
)

type dataFormat2 struct {
	Measure struct {
		Type string `xml:"type,attr"`
		AL   struct {
			R struct {
				Typical struct {
					AxialLength string `xml:"AxialLength"`
				} `xml:"Typical"`
			} `xml:"R"`
			L struct {
				Typical struct {
					AxialLength string `xml:"AxialLength"`
				} `xml:"Typical"`
			} `xml:"L"`
		} `xml:"AL"`

		KM struct {
			Condition string `xml:"condition,attr"`
			R         struct {
				Typical struct {
					R1 struct {
						Power string `xml:"Power"`
						Axis  string `xml:"Axis"`
					} `xml:"R1"`
					R2 struct {
						Power string `xml:"Power"`
						Axis  string `xml:"Axis"`
					} `xml:"R2"`
				} `xml:"Typical"`
			} `xml:"R"`
			L struct {
				Typical struct {
					R1 struct {
						Power string `xml:"Power"`
						Axis  string `xml:"Axis"`
					} `xml:"R1"`
					R2 struct {
						Power string `xml:"Power"`
						Axis  string `xml:"Axis"`
					} `xml:"R2"`
				} `xml:"Typical"`
			} `xml:"L"`
		} `xml:"KM"`

		CCT struct {
			R struct {
				Typical struct {
					CCT string `xml:"CCT"`
				} `xml:"Typical"`
			} `xml:"R"`
			L struct {
				Typical struct {
					CCT string `xml:"CCT"`
				} `xml:"Typical"`
			} `xml:"L"`
		} `xml:"CCT"`

		ACD struct {
			R struct {
				Typical struct {
					ACD string `xml:"ACD"`
				} `xml:"Typical"`
			} `xml:"R"`
			L struct {
				Typical struct {
					ACD string `xml:"ACD"`
				} `xml:"Typical"`
			} `xml:"L"`
		} `xml:"ACD"`

		WTW struct {
			R struct {
				Typical struct {
					WhiteToWhite string `xml:"WhiteToWhite"`
				} `xml:"Typical"`
			} `xml:"R"`
			L struct {
				Typical struct {
					WhiteToWhite string `xml:"WhiteToWhite"`
				} `xml:"Typical"`
			} `xml:"L"`
		} `xml:"WTW"`

		PS struct {
			R struct {
				Typical struct {
					PupilSize string `xml:"PupilSize"`
				} `xml:"Typical"`
			} `xml:"R"`
			L struct {
				Typical struct {
					PupilSize string `xml:"PupilSize"`
				} `xml:"Typical"`
			} `xml:"L"`
		} `xml:"PS"`
	} `xml:"Measure"`
}

func TestALScanXML(t *testing.T) {
	dataSlice, err := readFileUTF16("../xml/alscan.xml")
	if err != nil {
		fmt.Printf("%s", err)
		os.Exit(1)
	}

	var data dataFormat2
	decoder := xml.NewDecoder(bytes.NewBuffer(dataSlice))
	decoder.CharsetReader = func(charsetT string, input io.Reader) (io.Reader, error) {
		return charset.NewReader(input, charsetT)
	}
	err = decoder.Decode(&data)
	if err != nil {
		fmt.Println("Error unmarshalling from XML", err)
		return
	}

	fmt.Println(data.Measure.KM)

}
