local json = require "json"

function Handler(request)
    local jsonString = '[ { "key1": { "key2": "value1" } }, { "key1": "value2" } ]'
    local result = json.Decode(jsonString)

    request.write(result[1].key1.key2)
end