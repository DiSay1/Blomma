options = {
    Address = "/websocket", 
    WebSocket = true,
}

function onMessage(req)
    req.write(req.mt, req.data)
end

function onClose(req)
    print(req.mt)
end