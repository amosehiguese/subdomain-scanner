package main

import (
	"context"
	"net/http"
	"time"

	"github.com/google/uuid"
	"go.uber.org/zap"
)

type ctxKeyLog struct{}
type ctxkeyRequestID struct{}

type logHandler struct {
	log  *zap.Logger
	next http.Handler
}

type responseRecorder struct {
	b      int
	status int
	w      http.ResponseWriter
}

func (r *responseRecorder) Header() http.Header {
	return r.w.Header()
}

func (r *responseRecorder) Write(p []byte) (int, error) {
	if r.status == 0 {
		r.status = http.StatusOK
	}
	n, err := r.w.Write(p)
	r.b += n
	return n, err
}

func (r *responseRecorder) WriteHeader(statusCode int) {
	r.status = statusCode
	r.w.WriteHeader(statusCode)
}

func (lh *logHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	requestID, _ := uuid.NewRandom()
	ctx = context.WithValue(ctx, ctxkeyRequestID{}, requestID.String())

	start := time.Now()
	rr := &responseRecorder{w: w}

	log := lh.log.WithOptions(
		zap.Fields(
			zap.String("http.req.path", r.URL.Path),
			zap.String("http.req.method", r.Method),
			zap.String("http.req.id", requestID.String()),
		),
	)

	log.Debug("request started")
	defer func() {
		log.WithOptions(
			zap.Fields(
				zap.Int64("http.resp.took_ms", int64(time.Since(start)/time.Millisecond)),
				zap.Int("http.resp.status", rr.status),
				zap.Int("http.resp.bytes", rr.b),
			),
		).Sugar().Debugf("request complete")
	}()

	ctx = context.WithValue(ctx, ctxKeyLog{}, log)
	r = r.WithContext(ctx)
	lh.next.ServeHTTP(rr, r)
}
