# URL Shortener Go


### data needed
```plaintext
id: int (primary key)
ShortCode:
OriginalURL - str
CreatedAt
ClickCount
```

### app layers
- http layer: *handles requests and responses*
- business logic layer: *generates short codes + validates URLs*
- storage layer: handles *db reads/writes*