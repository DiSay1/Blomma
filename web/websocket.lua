options = {
    WebSocket = true,
}

connectionISOpen = false

function WSHandler(conn)
    connectionISOpen = true
    while connectionISOpen do
        data = conn.read()

        if connectionISOpen then
            conn.write(data.mt, data.data) 
        end
    end
end

function onClose(data)
    connectionISOpen = false
end