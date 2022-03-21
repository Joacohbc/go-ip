mv ./bin/go-ip.bin ./bin/go-ip
if [ $? -eq 0 ];then
    echo "Nombre cambiado go-ip.bin -> ./bin/go-ip..."
fi

chmod +x ./bin/go-ip
if [ $? -eq 0 ];then
    echo "Permisos de ejecucion dados..."
fi

sudo mv ./bin/go-ip /usr/bin/
if [ $? -eq 0 ];then
    echo "Archivo movido con exito..."
fi

echo "Instalado con exito"