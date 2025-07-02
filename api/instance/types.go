package instance

type BackupType string

const (
	BackupEnabled  BackupType = "enabled"
	BackupDisabled BackupType = "disabled"
)

func (b BackupType) String() string {
	return string(b)
}

func (b BackupType) IsValid() bool {
	return b == BackupEnabled || b == BackupDisabled
}
