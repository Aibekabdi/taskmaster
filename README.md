# TaskMaster Микросервис

Это RESTful API микросервис для управления вашими задачами (Todo List). Он предоставляет базовые CRUD операции для управления задачами, а также дополнительные функции, такие как пометка задачи как выполненной.

## Запуск
    
```bash
    make compose-build 
    make compose-up
```

## API Endpoints

1. **Создание новой задачи:**
    POST /api/todo-list/tasks
Тело запроса:

```json
{
   "title":"Купить книгу",
   "activeAt":"2023-08-04"
}
    примичания activeAt должен быть позже нынешнего времени (пример: сейчвс 2023-08-03 значит день должен быть больше 03)
```
2. **Обновления задачи:**
    PUT /api/todo-list/tasks/{ID}
Тело запроса:

```json
{
   "title":"Купить книгу",
   "activeAt":"2023-08-04"
}
```

2. **Удаление задачи:**
    DELETE /api/todo-list/tasks/{ID}

3. **Пометка задачи как выполненной:**
    PUT /api/todo-list/tasks/{ID}/done

4. **Получение списка задач по статусу:**
GET /api/todo-list/tasks?status=active или /api/todo-list/tasks?status=done