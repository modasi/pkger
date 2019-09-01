package memware

import (
	"testing"

	"github.com/markbates/pkger/here"
	"github.com/markbates/pkger/pkging/waretest"
	"github.com/stretchr/testify/require"
)

func Test_Warehouse(t *testing.T) {
	r := require.New(t)

	info, err := here.Current()
	r.NoError(err)
	r.NotZero(info)

	wh, err := New(info)
	r.NoError(err)

	WithInfo(wh, info)

	suite, err := waretest.NewSuite(wh)
	r.NoError(err)

	suite.Test(t)
}
