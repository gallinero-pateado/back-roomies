## Pasos para iniciar back roomies.
1. Asegurarse de tener las credenciales necesearias (Archivo ".env" y el "serviceAccountKey"). Asegurarse que los nombres de las variables en el ".env" correspondan a las variables que toma el archivo "API/config/config.go"
2. Dichas credenciales deben estar en la carpeta raíz del proyecto.
3. Ejecutar el comando "go run main.go"

#NOTA SPRINT 5
El cuerpo de petición para generar el reporte es:
{
  "UsuarioReportadoID": 189,
  "UsuarioReportanteID": 200,
  "Motivo": "El usuario publicó contenido inapropiado."
}

El cuerpo para actualizar el reporte es:
{
    "Estado": "Resuelto"
}
