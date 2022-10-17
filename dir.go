package mobile

type DirInfo struct {
	Document  string
	Cache     string
	Temporary string
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
