// Code generated by "enumer -type=SystemTag -transform=snake -trimprefix=Tag -output system_tag_gen.go"; DO NOT EDIT.

//
package metrics

import (
	"fmt"
)

const _SystemTagName = "protosubprotostatusmethodurlraw_urlnamegroupcheckerrorerror_codetls_versionscenarioserviceexpected_responseitervuocsp_statusip"

var _SystemTagMap = map[SystemTag]string{
	1:      _SystemTagName[0:5],
	2:      _SystemTagName[5:13],
	4:      _SystemTagName[13:19],
	8:      _SystemTagName[19:25],
	16:     _SystemTagName[25:28],
	32:     _SystemTagName[28:35],
	64:     _SystemTagName[35:39],
	128:    _SystemTagName[39:44],
	256:    _SystemTagName[44:49],
	512:    _SystemTagName[49:54],
	1024:   _SystemTagName[54:64],
	2048:   _SystemTagName[64:75],
	4096:   _SystemTagName[75:83],
	8192:   _SystemTagName[83:90],
	16384:  _SystemTagName[90:107],
	32768:  _SystemTagName[107:111],
	65536:  _SystemTagName[111:113],
	131072: _SystemTagName[113:124],
	262144: _SystemTagName[124:126],
}

func (i SystemTag) String() string {
	if str, ok := _SystemTagMap[i]; ok {
		return str
	}
	return fmt.Sprintf("SystemTag(%d)", i)
}

var _SystemTagValues = []SystemTag{1, 2, 4, 8, 16, 32, 64, 128, 256, 512, 1024, 2048, 4096, 8192, 16384, 32768, 65536, 131072, 262144}

var _SystemTagNameToValueMap = map[string]SystemTag{
	_SystemTagName[0:5]:     1,
	_SystemTagName[5:13]:    2,
	_SystemTagName[13:19]:   4,
	_SystemTagName[19:25]:   8,
	_SystemTagName[25:28]:   16,
	_SystemTagName[28:35]:   32,
	_SystemTagName[35:39]:   64,
	_SystemTagName[39:44]:   128,
	_SystemTagName[44:49]:   256,
	_SystemTagName[49:54]:   512,
	_SystemTagName[54:64]:   1024,
	_SystemTagName[64:75]:   2048,
	_SystemTagName[75:83]:   4096,
	_SystemTagName[83:90]:   8192,
	_SystemTagName[90:107]:  16384,
	_SystemTagName[107:111]: 32768,
	_SystemTagName[111:113]: 65536,
	_SystemTagName[113:124]: 131072,
	_SystemTagName[124:126]: 262144,
}

// SystemTagString retrieves an enum value from the enum constants string name.
// Throws an error if the param is not part of the enum.
func SystemTagString(s string) (SystemTag, error) {
	if val, ok := _SystemTagNameToValueMap[s]; ok {
		return val, nil
	}
	return 0, fmt.Errorf("%s does not belong to SystemTag values", s)
}

// SystemTagValues returns all values of the enum
func SystemTagValues() []SystemTag {
	return _SystemTagValues
}

// IsASystemTag returns "true" if the value is listed in the enum definition. "false" otherwise
func (i SystemTag) IsASystemTag() bool {
	_, ok := _SystemTagMap[i]
	return ok
}