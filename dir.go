package mobile

import "code.olapie.com/log"

type DirInfo struct {
	Document  string
	Cache     string
	Temporary string
}

func (d *DirInfo) MustMakeDirs() {
	if d.Document != "" {
		MustMkdir(d.Document)
	} else {
		log.S().Warn("Document directory is not specified")
	}

	if d.Cache != "" {
		MustMkdir(d.Cache)
	} else {
		log.S().Warn("Cache directory is not specified")
	}

	if d.Temporary != "" {
		MustMkdir(d.Temporary)
	} else {
		log.S().Warn("Temporary directory is not specified")
	}
}

func NewDirInfo() *DirInfo {
	return new(DirInfo)
}

func NewTestDirInfo() *DirInfo {
	return &DirInfo{
		Document:  "testdata/document",
		Cache:     "testdata/cache",
		Temporary: "testdata/temporary",
	}
}
