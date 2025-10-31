package stringdata

func ToNilIfEmpty(v *string) *string {
	if v == nil || *v == "" {
		return nil
	}
	return v
}