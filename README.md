BS - изменяемая строка, все операции вызываемые от объекта меняют строку внутри,
чтобы сделать операцию без изменения внутренней строки нужно вызвать метод от bs.str.\
Пример: strings.TrimSpace(bs.str) \
Здесь сделаны не все функции из strings, остальные можно вызвать непосредственно от bs.str,
так как они не меняют исходную строку.\
Пример: strings.Split(bs.str,";") 

Добавлены функции из unicode, они применяют функцию из unicode к элементу в строке
с номером index.
