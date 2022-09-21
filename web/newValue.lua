local valueController = require "valueController"

options = {
    Address = "/value",
}

function Handler(request)
    value = request.getQuery("value")

    valueController.newValue(value, 0)

    request.write("New value "..value.." successfully created")
end