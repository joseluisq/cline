package cline

// FlagMap defines a hash map of command flags.
type FlagMap struct {
	flags []Flag
}

func (fm *FlagMap) findByKey(flagKey string) FlagValue {
	for _, v := range fm.flags {
		switch fl := v.(type) {
		case FlagBool:
			if flagKey == fl.Name {
				return fl.zflag
			}
			break
		case FlagInt:
			if flagKey == fl.Name {
				return fl.zflag
			}
			break
		case FlagString:
			if flagKey == fl.Name {
				return fl.zflag
			}
			break
		case FlagStringSlice:
			if flagKey == fl.Name {
				return fl.zflag
			}
			break
		}
	}
	return FlagValue("")
}

// Bool gets current flag value as `bool`.
func (fm *FlagMap) Bool(flagName string) (bool, error) {
	return fm.findByKey(flagName).Bool()
}

// Int gets current flag value as `int`.
func (fm *FlagMap) Int(flagName string) (int, error) {
	return fm.findByKey(flagName).Int()
}

// String gets current flag value as `string`.
func (fm *FlagMap) String(flagName string) string {
	return fm.findByKey(flagName).String()
}

// StringSlice gets current flag value as a string slice.
func (fm *FlagMap) StringSlice(flagName string) []string {
	return fm.findByKey(flagName).StringSlice()
}
