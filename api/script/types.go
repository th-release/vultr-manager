package script

type ScriptType string

const (
	ScriptTypeBoot ScriptType = "boot"
	ScriptTypePxe  ScriptType = "pxe"
)

func (s ScriptType) String() string {
	return string(s)
}

func (s ScriptType) IsValid() bool {
	return s == ScriptTypeBoot || s == ScriptTypePxe
}
