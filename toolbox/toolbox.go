package toolbox

import (
	"crypto/sha1"
	"encoding/hex"
	"fmt"
	"github.com/satori/go.uuid"
	"golang.org/x/crypto/pbkdf2"
	. "neoops/adam/pkg/model/consts"
	"neoops/adam/pkg/module/exception/errors"
	"regexp"
)

func GenerateUUID() string {
	return uuid.NewV4().String()
}

func Encrypt(content string, salt string) string {
	return fmt.Sprintf("%x", pbkdf2.Key([]byte(content), []byte(salt), 4096, 16, sha1.New))
}

func ValidateRecordValue(t DNSRecordType, k, v string) error {
	switch t {
	case DNSRecordTypeNS:
		matched, _ := regexp.MatchString(`^(`+RegexpPatternIPorDomain+`)\.$`, v)
		if !matched {
			return errors.NewFieldError(k, v, "(127.0.0.1 | ::1 | example.org).").Expand(true)
		}
	case DNSRecordTypeCNAME:
		matched, _ := regexp.MatchString(`^(`+RegexpPatternDomain+`)\.$`, v)
		if !matched {
			return errors.NewFieldError(k, v, "www.example.org.").Expand(true)
		}
	case DNSRecordTypeA:
		matched, _ := regexp.MatchString(`^(`+RegexpPatternIPv4+`)$`, v)
		if !matched {
			return errors.NewFieldError(k, v, "127.0.0.1").Expand(true)
		}
	case DNSRecordTypeURL:
		matched, _ := regexp.MatchString(`^(`+RegexpPatternURL+`)$`, v)
		if !matched {
			return errors.NewFieldError(k, v, "http://www.example.org").Expand(true)
		}
	case DNSRecordTypeMX:
		matched, _ := regexp.MatchString(`^[1-9]\d? (`+RegexpPatternIPorDomain+`)\.$`, v)
		if !matched {
			return errors.NewFieldError(k, v, "10 (127.0.0.1 | ::1 | example.org).").Expand(true)
		}
	case DNSRecordTypeAAAA:
		matched, _ := regexp.MatchString(`^(`+RegexpPatternIPv6+`)$`, v)
		if !matched {
			return errors.NewFieldError(k, v, "::1").Expand(true)
		}
	default:
		if v == "" {
			return errors.NewFieldError(k, v).Expand(true)
		}
	}
	return nil
}

func SHA1Sum(ss ...string) string {
	text := ""
	for _, s := range ss {
		text += s
	}
	sha1sum := sha1.Sum([]byte(text))
	return hex.EncodeToString(sha1sum[:])
}
