# TeleGraph

## Descripción

La aplicación del desafio retro con la finalidad de ayudar a MELI en sus comunicaciones internas.

## Entorno deploy

Para el despliegue se uso google cloud tomando en cuenta recomendación de MELI [https://appgraph.appspot.com](https://appgraph.appspot.com)

## URL local equivalente

- [http://localhost:8080/](http://localhost:8080/)

## EndPoints

Para traducir código morse a texto

- Method: <b>POST</b>
- [http://localhost:8080/2text](http://localhost:8080/2text)

Ejemplo:

```
request:

curl -X POST "appgraph.appspot.com/2text" -d '{"text": ".... --- .-.. .- / -- . .-.. .."}'
```

```
response:

{"Code":200,"Response":"hola meli"}
```

Para traducir texto a código morse.

- Method: <b>POST</b>
- [http://localhost:8080/2morse](http://localhost:8080/2morse)

Ejemplo:

```
request:

curl -X POST "http://localhost:8080/2morse" -d '{"text": "hola meli"}'
```

```
response:

{"Code":200,"Response":".... --- .-.. .-  -- . .-.. .."}
```

## EndPoints Plus

Para traducir bits a código morse

- Method: <b>POST</b>
- [http://localhost:8080/2bitsmorse](http://localhost:8080/2bitsmorse)

Ejemplo:

```
request:

curl -X POST "appgraph.appspot.com/2bitsmorse" -d '{"text": "000000001101101100111000001111110001111110011111100000001110111111110111011100000001100011111100000111111001111110000000110000110111111110111011100000011011100000000000"}'
```

```
response:

{"Code":200,"Response":" .... --- .-.. .- -- . .-.. .."}
```

Para traducir bits a texto

- Method: <b>POST</b>
- [http://localhost:8080/2bitstext](http://localhost:8080/2bitstext)

Ejemplo:

```
request:

curl -X POST "appgraph.appspot.com/2bitstext" -d '{"text": "000000001101101100111000001111110001111110011111100000001110111111110111011100000001100011111100000111111001111110000000110000110111111110111011100000011011100000000000"}'
```

```
response:

{"Code":200,"Response":"holameli"}
```

Para traducir morse a bits

- Method: <b>POST</b>
- [http://localhost:8080/2morsebits](http://localhost:8080/2morsebits)

Ejemplo:

```
request:

curl -X POST "http://localhost:8080/2morsebits" -d '{"text": ".... --- .-.. .- -- . .-.. .."}'
```

```
response:

{"Code":200 "Response":"110101001100000001111110011110111110000011001111100111011000000011101111110000001111111001111000001100000001110011111110100111000001001110000000000"}
```

Para traducir texto a bits

- Method: <b>POST</b>
- [http://localhost:8080/2humanbit](http://localhost:8080/2humanbit)

Ejemplo:

```
request:

curl -X POST "http://localhost:8080/2humanbits" -d '{"text": "hola meli"}'
```

```
response:

{"Code":200,"Response":"1001110011100111000000111110011111110011110000001110011111110110011100000111001111111000000000001111110011111110000001100000110011111011100111000001110110000000000"}
```

## Documentación

Para el cálculo y desarrollo del servicio se usaron las siguientes fuentes

- La tabla de referencia enviada del pdf enviado.
- [https://es.wikipedia.org/wiki/C%C3%B3digo_morse](https://es.wikipedia.org/wiki/C%C3%B3digo_morse)


## Notas observación:

- Quiza MELI no vio un patron de diseño de codigo en practica
- Quiza los datos debian estar en una base de datos o archivo json de lectura