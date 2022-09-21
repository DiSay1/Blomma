local valueController = require "valueController" -- Подключения библиотеки

-- Параметры обработчика
options = {
    Address = "/", -- Путь к обработчику
}

-- Функция вызываемая при запросе
function Handler (request)
    -- Записываем названия переменной из query параметров.
    local varible = request.getQuery("varible")

    -- Выводим значения желоемой переменной
    request.write(varible.." "..valueController.getValue(varible))
end