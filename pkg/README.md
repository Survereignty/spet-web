Код библиотеки, который можно использовать во внешних приложениях (например, / pkg / mypubliclib). Другие проекты будут импортировать эти библиотеки, ожидая, что они будут работать, поэтому подумайте дважды, прежде чем что-то здесь поместить :-) Обратите внимание, что внутренний каталог - лучший способ гарантировать, что ваши частные пакеты не будут импортированы, поскольку он поддерживается Go. Каталог / pkg по-прежнему является хорошим способом явно сообщить, что код в этом каталоге безопасен для использования другими.