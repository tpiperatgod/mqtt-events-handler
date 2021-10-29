package mqtt_events_handler

import (
    ofctx "github.com/OpenFunction/functions-framework-go/openfunction-context"
    log "k8s.io/klog/v2"
)

func MQTTHandler(ctx *ofctx.OpenFunctionContext, in []byte) ofctx.RetValue {
    content := string(in)
    log.Info(content)
    return ctx.ReturnWithSuccess()
}