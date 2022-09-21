local valueController = require "valueController" -- Подключения библиотеки

-- Параметры обработчика
options = {
    Address = "/newValue", -- Путь к обработчику
}

-- Функция вызываемая при запросе
function Handler(request)
    -- Записываем названия создаваемой переменной из query параметра
    local varibleName = request.getQuery("varibleName")

    valueController.newValue(varibleName, 0) -- Создаем переменную

    -- Выводим сообщения
    request.write("New varible "..varibleName.." successfully created") 
end