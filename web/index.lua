address = "/"

function Handler (req)
    if req.method == "POST" then
        postData = req.getFormData({"message"})
        req.write(postData.message)
    elseif req.method == "GET" then
        header = req.getHeader("message")
        query = req.getQuery("index")
        req.write(query.."\n")
        req.write(header.."\n")
    end
end