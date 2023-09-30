package middleware

import (
	"blog/global"
	"blog/pkg/logger"
	"bytes"
	"time"

	"github.com/gin-gonic/gin"
)

type AccessLogWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func (w AccessLogWriter) Write(p []byte) (int, error) {
	if n, err := w.body.Write(p); err != nil {
		return n, err
	}
	return w.ResponseWriter.Write(p)
}

func AccessLog() gin.HandlerFunc {
	return func(c *gin.Context) {
		bw := &AccessLogWriter{
			body:           bytes.NewBufferString(""),
			ResponseWriter: c.Writer,
		}
		c.Writer = bw

		bTime := time.Now().Unix()
		c.Next()
		eTime := time.Now().Unix()

		fields := logger.Fields{
			"request":  c.Request.PostForm.Encode(),
			"response": bw.body.String(),
		}
		s := "access log: method: %s, status coe: %d begin_time: %d, end_time: %d"
		log := global.Logger.WithFields(fields)
		if global.ServerSetting.IsTracing && global.Tracer != nil {
			log = log.WithContext(c).WithTrace()
		}
		log.Infof(s, c.Request.Method, bw.Status(), bTime, eTime)
	}
}
