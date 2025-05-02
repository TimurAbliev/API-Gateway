# api

1. Прописать ручки для api, сделать максимально просто в одном main.go файле, для хранения данных исопльзуй map или slice
Использовать fiber для принятия зарпосов
id - uuid
    1. `POST` `/user` req body{name, age, ...}, resp id, 201, 400, 500
    2. `GET` `/user/:id` resp body{id, name, age, ...}, 200, 400, 500
    3. `DELETE` `/user/:id` resp 200, 400, 500
    4. `PUT` `/user` req body{id, name, age, ...}, resp id, 200, 400, 500

1. слайс vs массив, стурктура слайса, что происходит при append (cap)
2. map, бакеты, коллизии, эвакуации, сложность операции получения value, swisstable в 1.24 go
3. interface, что под капотом, solid, опп в golang (чек примеры)
4. Приведение типов
