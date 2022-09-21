local valueController = require "valueController"

options = {
    Address = "/",
    WebSocket = false,
}

function Handler (req)
    req.write("Close count = "..valueController.getValue("closeCount"))
end