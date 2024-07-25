# Endpoints

Aquellos endpoints que utilicen los verbos HTTP POST/PUT/PATCH, retornan el objeto actualizado.

La autenticación se hace con JWT.

## Sin JWT

- POST signup
- POST login
- GET list posts
- GET post
- GET list comments from post
- GET comment from post
- GET models (nombres + rating)
- POST accesibilize

## Con JWT

- POST create post
- PUT update post
- PATCH like post
- POST create comment of post
- PUT update comment of post
- PATCH like comment of post
- PUT update configuration
- PATCH puntuar accesibilidad

## Listados

Realizar paginado con query params, `page` para determinar que página es y `size` para determinar el tamaño de la página. Los valores por defecto serían los siguientes:
- `page = 0`
- `size = 10`

