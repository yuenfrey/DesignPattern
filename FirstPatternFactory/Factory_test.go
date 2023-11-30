package FirstPatternFactory

import (
	"testing"
)

func TestNewImpl(t *testing.T) {
	t.Run("pass", func(t *testing.T) {
		impl := NewImpl()
		impl.AddSvc("first", NewFirstSvc("127.0.1"))
		impl.AddSvc("second", NewSecondSvc("192.168.0.1"))
		impl.Dispatcher(InspectReq{
			Name:   "test req",
			Policy: "imma",
			Items:  nil,
			Object: nil,
		})
	})

}
