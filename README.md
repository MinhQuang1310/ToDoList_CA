# Các API để test:
-  Get all items: http://localhost:8080/v1/items/getall
-  Get item by id: http://localhost:8080/v1/items/getitem/1
-  Create item: http://localhost:8080/v1/items/create
  Body:
  {
    "title": "abc3",
    "description": "abc3",
    "status": "available"
  }
- Update item: http://localhost:8080/v1/items/update/1
  Body:
  {
    "title": "abc100",
    "description": "abc100",
  }
- Delete item: http://localhost:8080/v1/items/delete/1
