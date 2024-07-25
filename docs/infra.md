# Infraestructura

Las imagenes de go para produccion/desarrollo se construyen con alpine, version de linux mas liviana. Para las de desarrollo se utiliza el comando `air` para realizar hot-reload.

Por otra parte se genera un `compose.yml` con los servicios necesarios para levantar, para usarlo basta con poner `docker compose up`.

Se pensaba en un futuro agregar CI/CD con test y pa11y para verificar la accesibilidad.
