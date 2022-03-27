package helper

import (
	"testing"

	log "github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
)

func TestLogContext(t *testing.T) {
	c := "test"
	s := "test"

	t.Run("SUCCESS LOGCONTEXT", func(t *testing.T) {
		assert.NotNil(t, LogContext(c, s))
	})
}

var (
	testConst = "any message"
	testTags  = map[string]interface{}{"test": testConst}
)
var testCasesLog = []struct {
	name    string
	level   log.Level
	message string
	context string
	scope   string
	tags    map[string]interface{}
}{
	{
		name:    "#1 Debug",
		level:   log.DebugLevel,
		message: testConst,
		context: testConst,
		scope:   testConst,
		tags:    testTags,
	},
	{
		name:    "#2 Info",
		level:   log.InfoLevel,
		message: testConst,
		context: testConst,
		scope:   testConst,
		tags:    testTags,
	},
	{
		name:    "#3 Warn",
		level:   log.WarnLevel,
		message: testConst,
		context: testConst,
		scope:   testConst,
		tags:    testTags,
	},
	{
		name:    "#4 Error",
		level:   log.ErrorLevel,
		message: testConst,
		context: testConst,
		scope:   testConst,
		tags:    testTags,
	},
	{
		name:    "#5 Panic",
		level:   log.PanicLevel,
		message: testConst,
		context: testConst,
		scope:   testConst,
		tags:    testTags,
	},
}

func TestLog(t *testing.T) {

	t.Parallel()

	for _, tc := range testCasesLog {
		t.Run(tc.name, func(*testing.T) {
			Log(tc.level, tc.message, tc.context, tc.scope)
		})
	}
}
