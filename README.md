# Evidencia de Aprendizaje 1 — Teoría de Autómatas

Autómata finito determinista que valida cadenas sobre el alfabeto formado por
las letras del nombre completo, los dígitos de la matrícula y el punto.

Una cadena es válida si:

1. El primer símbolo es un dígito.
2. Tiene cualquier combinación intermedia de letras y dígitos del alfabeto.
3. Contiene las iniciales de forma consecutiva al menos una vez.
4. Termina con un punto seguido de la matrícula completa.
5. Acepta puntos intermedios, pero no consecutivos.

## Uso

Requiere Go 1.26 o superior.

```sh
go run ./cmd
```

El programa pide el nombre completo (en minúsculas y sin acentos) y la
matrícula, y después las cadenas a validar hasta que se responda `n`.

Ejemplo con `luis daniel salazar sanchez` y `2222332` (iniciales `ldss`):

| Cadena                              | Resultado |
|-------------------------------------|-----------|
| `23ldssdaniel.2222332`              | Válida    |
| `2luis.danielldss3.sanchez.2222332` | Válida    |
| `2ldss..2222332`                    | Inválida  |
| `2lds.2222332`                      | Inválida  |

## Estructura

- `cmd/main.go`: lectura de datos y ciclo de validación.
- `pkg/fsm`: ejecutor genérico de máquinas de estados.
- `pkg/utils`: obtención del alfabeto y de las iniciales.
- `internal/aut`: estados del autómata.

## Estados

| Estado | Significado |
|--------|-------------|
| Q0 | Inicial: espera un dígito de la matrícula. |
| Q1 | Zona libre antes de las iniciales. |
| Q2 | Se leyó un punto antes de las iniciales. |
| Q3 + cadena de iniciales | Reconocimiento de las iniciales consecutivas. |
| Q4 | Zona libre después de las iniciales. |
| Q5 | Se leyó un punto después de las iniciales. |
| Cadena de matrícula | Reconocimiento de `.matrícula`. |
| Q6 | **Aceptación**: se leyó el punto y la matrícula completa. |

Los estados de las iniciales y de la matrícula se generan dinámicamente.
Cuando un símbolo no coincide, se usa la función de prefijos (KMP) para caer
al estado correcto sin perder lo que ya se reconoció.

## Licencia

MIT. Ver [LICENSE](LICENSE).
