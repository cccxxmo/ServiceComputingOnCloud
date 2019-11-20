# BlogAPI
## HTTP方法
HTTP方法的描述如下：

| 动词 | 描述 |  
| ---- | ---- |  
| `HEAD`	| 可以针对任何资源发出以仅获取HTTP标头信息。 |
| `GET`	| 用于检索资源。 |
| `POST`	| 用于创建资源。 |
| `PATCH`	| 用于通过部分JSON数据更新资源。例如，问题资源具有`title`和`body`属性。PATCH请求可以接受一个或多个属性以更新资源。PATCH是一个相对较新且不常见的HTTP动词，因此资源端点也接受`POST`请求。 |
| `PUT` |	用于替换资源或集合。对于`PUT`没有`body`属性的请求，请确保将`Content-Length`标头设置为零。 |
| `DELETE`	| 用于删除资源。 |

针对博客所采用的约定如下：  

| 资源 | `POST` | `GET` | `PUT` | `DELETE` |  
| ---- | ---- | ---- | ---- | ---- |
| /users | 创建新用户 | 检索所有用户 | 批量更新用户 | 删除所有用户 |
| /users/userID | 错误 | 检索ID为userID的用户的详细信息 | 如果userID存在，则更新其详细信息 | 删除ID为userID的用户 |
| /users/userID/articles | 创建ID为userID的用户的新文章 | 检索userID的所有文章 | 批量更新userID的文章 | 删除userID的所有文章 |
| /users/userID/articles/articleID | 错误 | 检索ID为articleID的文章的内容 | 如果articleID存在，则更新其内容 | 删除ID为articleID的文章 |
