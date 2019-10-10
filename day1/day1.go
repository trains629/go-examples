package day1

import (
	"encoding/json"
	"encoding/xml"
)

type MP3Index struct {
	Index xml.Name  `xml:"index"`
	Group *MP3Group `xml:"group"`
}

type MP3Tag struct {
	Text     string `xml:"swac_text,attr" json:"text"`
	Alphaidx string `xml:"swac_alphaidx,attr" json:"alphaidx"`
}

type MP3File struct {
	Path string `xml:"path,attr" json:"path"`
	Tag  MP3Tag `xml:"tag" json:"tag"`
}

type MP3Group struct {
	Authors string    `xml:"swac_coll_authors,attr"`
	Lang    string    `xml:"swac_lang,attr"`
	File    []MP3File `xml:"file" json:"file"`
}

func Xml2Json(b *[]byte) (*[]byte, error) {
	group := MP3Index{Group: &MP3Group{}}

	if err := xml.Unmarshal(*b, &group); err != nil {
		return nil, err
	}

	t2, err := json.Marshal(group.Group.File)
	if err != nil {
		return nil, err
	}

	return &t2, nil
}
