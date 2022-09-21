local valueController = require "valueController"

options = {
    Address = "/websocket", 
    WebSocket = true,
}

function onMessage(req)
    req.write(req.mt, req.data)
end

function onClose(req)
    closeCount = valueController.getValue("closeCount")
    closeCount = closeCount + 1
    valueController.updateValue("closeCount", closeCount)
end