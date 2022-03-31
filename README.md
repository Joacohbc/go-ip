# Go-ip

Go-ip es un comando que cumple la funcion consultar el IP publica del dispositivo y tambien informacion de las interfaces de red del dispositivo, entre esa información el IP Privada de alguna placa

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

Para consultar la IP de alguna placa, en el ejemplo la placa eth0

```bash
go-ip private -i eth0
```

Si quiero consultar toda la información de toda las placas simplemente pongo

```bash
go-ip private
```

Si quiero consultar toda la información de una sola placa utilizó

```bash
go-ip private -i eth0 -a
```

## Instalación

En caso de querer instalarlo solo hay que mover compilar el binario (o usar el compilado en directorio bin) y moverlo
a la carpeta de binarios del S.0, en Linux generalmente /usr/bin. Este proceso ya lo hace el install.sh.

- Clonar el repositorio

```bash
git clone https://github.com/Joacohbc/go-ip
```

- Entrar a la carpeta

```bash
cd go-ip
```

- Y ejecutar el instalador

```bash
sh ./install.sh
```
