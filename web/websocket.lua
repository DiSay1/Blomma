options = {
    Address = "/websocket", 
    WebSocket = true,
}

function Handler(req)
    req.write(req.mt, "Ваня ЛОХ!")
end