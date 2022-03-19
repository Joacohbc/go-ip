# Go-ip

Go-ip es un comando que cumple la funcion consultar el IP publica del dispositivo y tambien informacion de las interfaces de red del dispositivo, entre esa informacion el IP Privada de alguna placa

Para consultar el IP Publico

```bash
#Para consultar el IP Public
go-ip public 
```

Para poder guardar el IP en el portapapeles se pueden usar herramientas como xclip o xsel de la siguiente forma:

```bash
#Para ingresar el IP en el portapapeles
go-ip public | xclip -sel clip

#Para sacarlo del portapapeles
xclip -o sel clip
```

Para consultar le IP de alguna placa, en el ejemplo la placa eth0

```bash
go-ip private -i eth0
```

Si quiero consultar toda la informacion de toda las placas simplemente pongo

```bash
go-ip private
```

Si quiero consultar toda la informacion de una sola placa utilizo

```bash
go-ip private -i eth0 -a
```

## Instalacion

En caso de quere instalarlo solo hay que mover compilar el binario (o usar el compilado en directorio bin) y moverlo
a la carpeta de binarios del S.0, en Linux generalmente /usr/bin.

- Clonar el repositorio

```bash
git clone https://github.com/Joacohbc/go-ip
cd go-ip
```

- Compilar el binario

```bash
go build -o go-ip ./src/*.go
```

- Mover binario a la carpeta de binarios, en este caso /usr/bin

```bash
sudo mv ./go-ip /usr/bin
```
