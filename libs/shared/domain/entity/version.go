package entity

type Version struct {
	version uint
}

func (e *Version) EntityVersion() uint {
	return e.version
}

func (e *Version) Next() *Version {
	return &Version{e.version + 1}
}

func (e *Version) Previous() *Version {
	return &Version{e.version - 1}
}

func (e *Version) Equals(comparing Version) bool {
	return e.version == comparing.version
}

func CreateNewVersion() *Version {
	return &Version{1}
}

func NewVersion(version uint) *Version {
	return &Version{version}
}
