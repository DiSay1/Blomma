local valueController = require "valueController" -- Library Connections

-- Handler Options
options = {
    Address = "/", -- Web path to handler
}

-- Function called on request
function Handler (request)
    -- Outputting the values of the desired variable
    request.write("Hello world!")
end