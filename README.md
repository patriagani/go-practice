# Kasir API Documentation

## Base URL

```
http://localhost:8080/api
```

## Products API

| Method | Route                | Description        |
| ------ | -------------------- | ------------------ |
| GET    | `/api/products`      | Get all products   |
| POST   | `/api/products`      | Create new product |
| GET    | `/api/products/{id}` | Get product by ID  |
| PUT    | `/api/products/{id}` | Update product     |
| DELETE | `/api/products/{id}` | Delete product     |

### Request Body (POST/PUT)

```json
{
  "name": "Product Name",
  "price": 10000,
  "stock": 100
}
```

## Categories API

| Method | Route                  | Description         |
| ------ | ---------------------- | ------------------- |
| GET    | `/api/categories`      | Get all categories  |
| POST   | `/api/categories`      | Create new category |
| GET    | `/api/categories/{id}` | Get category by ID  |
| PUT    | `/api/categories/{id}` | Update category     |
| DELETE | `/api/categories/{id}` | Delete category     |

### Request Body (POST/PUT)

```json
{
  "name": "Category Name",
  "description": "Category Description"
}
```

## Status Codes

- `200` OK
- `201` Created
- `400` Bad Request
- `404` Not Found
- `405` Method Not Allowed
- `500` Internal Server Error
