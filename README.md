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

Giải thích
1. todo_item_usecase.go: Định nghĩa struct TodoItemUsecase, có trách nhiệm xử lý logic kinh doanh liên quan đến các mục todo. Nó có một phương thức CreateItem để tạo một mục todo mới bằng cách gọi phương thức Create của TodoItemRepository.
2. main.go: Điểm vào của ứng dụng của bạn. Nó nhập các gói cần thiết và khởi tạo kết nối với cơ sở dữ liệu. Nó cũng thiết lập router Gin và định nghĩa các route để xử lý các yêu cầu HTTP.
3. todo_item_repository.go: Định nghĩa struct TodoItemRepository, có trách nhiệm tương tác với cơ sở dữ liệu để thực hiện các thao tác CRUD (Create, Read, Update, Delete) trên các mục todo. Nó có các phương thức như Create, Update, Delete, GetAll, và GetByID để thực hiện các thao tác này.
4. repository.go: Định nghĩa interface TodoItemRepository, xác định các phương thức mà một triển khai cụ thể của repository phải triển khai.
5. todo_item_handler.go: Định nghĩa struct TodoItemHandler, có trách nhiệm xử lý các yêu cầu HTTP liên quan đến các mục todo. Nó có các phương thức như CreateItem, UpdateItem, DeleteItem, và GetAllItems để xử lý các yêu cầu HTTP tương ứng. Nó sử dụng TodoItemUsecase để thực hiện logic kinh doanh và trả về các phản hồi HTTP tương ứng.
