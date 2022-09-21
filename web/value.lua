local valueController = require "valueController" -- Подключения библиотеки

-- Параметры обработчика
options = {
    Address = "/updateValue", -- Путь к обработчику
}

-- Функция вызываемая при запросе
function Handler (request)
    -- Записываем названия переменной из query параметров.
    local varible = request.getQuery("varible") 

    local value = valueController.getValue(varible)

    value = value + 1 -- Добавляем еденицу

    valueController.updateValue(varible, value) -- Обновляем значения

    request.write("Успешно обновленно!") -- Выводим сообщения
end