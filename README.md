# AVDM

AVD Manager es una herramienta que envuelve las `Command line tools` de android con el propósito de gestionar los dispositivos android virtuales, ademas de proporcionar una via rápida para instalar el entorno android sin necesidad de Android Studio IDE.

## Requisitos

- JDK

## Roadmap

- [ ] Proveer via rápida es instalación de entorno haciendo uso del `Command line tools` de android
  - [x] Descargar el comprimido del `Command line tools`.
  - [x] Descomprimir `Command line tools` y mover a ruta global de instalación.
  - [x] Limpiar archivos temporales automáticamente
  - [ ] registrar variables de entorno de android (crear entorno de pruebas para esto)
- [ ] Crear dispositivo AVD
  - [ ] Crear estructura de ejecutable de sdkmanager del `Command line tools`
  - [ ] presentar al usuario la lista de plataformas disponibles de android
  - [ ] descargar plataforma objetivo
  - [ ] crear dispositivo usando el comando `avd`
- [ ] Gestión AVD
  - [ ] Leer etiqueta de configuración
  - [ ] Editar etiqueta de configuración
  - [ ] Hacer backup de archivo de configuración
  - [ ] Retornar ruta del archivo de configuración

## Planteamiento

La idea es que el usuario final pueda usar todas las herramientas de terminal de android (adb, avd, emulator, etc) y que el propósito de avdm sea solo de gestionar dispositivos android virtuales (avd), util para cuando se quiera usar una nueva version de android o una version especifica.

Android Studio IDE tiene el avd manager, el cual es una herramienta gráfica para gestionar los distintos dispositivos android virtuales, la idea es recrear estas mecánicas en este programa y reducir la complejidad de hacerlo manualmente desde el `Command line tools` de android.

Ademas de proveer una via rápida para la instalación del entorno de herramientas de android en el caso que el usuario no las tenga y no quiera lidiar con el proceso manual de instalación desde las `Command line tools` de android. Inicialmente hay que instalar un entorno base sin mucha libertad de configuración (tampoco domino que tanto se pueda extender esta configuración inicial por lo que mejor limitarnos al proceso común)

## Desarrollo

## Propuesta de flujo de uso

### 1. Verificación

El usuario debería verificar primero el estado de su entorno de desarrollo android

ejecutando:

```bash
avdm status
```

Este debe retornar:

- Si existe JDK y su version
- Si existe herramientas de android y version:
  - ADB
  - AVD
  - SDKManager
  - Emulator
  - etc
- En caso de existir avd, enumerar los dispositivos virtuales creados y sus nombres

### 2. Inicialización

Si el usuario no tiene el entorno instalado puede hacer

```bash
avdm init
```

este comando procedería a descargar el `Command line tools` desde la pagina de android, descomprimir el archivo y incluir en las variables de entorno los ejecutables de dicho recurso, de esta manera el usuario tendría full acceso a las herramientas de android y avdm serviría de medio para la instalación de estos, ademas de usar algunos de estos.

### 3. Creación de dispositivo

este es el propósito central de la app, crear un dispositivo virtual

```bash
avdm create
```

este comando debe presentar un formulario en interfaz para personalizar las partes relevantes de la creación de un dispositivo virtual, debe haber dos vertientes, la configuración full automatizada y la que permita insertar ciertos parámetros personalizados.

### 4. Edición de dispositivo

Manipular la configuración de un dispositivo virtual

para este creo que hay que ser mas flexibles

una opcion debe ser:

```bash
avdm set [nombre_dispositivo] [nombre_parametro] [valor]
```

de esta manera se puede hacer pequenas modificaciones sin tener que exponer el archivo completo.

otra opcion a considerar debe ser

```bash
# exponer la direccion de la configuracion y abrirla con el editor de preferencia
avdm device [nombre_dispositivo] --path | vim
```

### 5. Eliminación de dispositivo

```bash
avdm delete [nombre_dispositivo]
```

### 6. Actualización de dispositivo

este debería lanzar de nuevo el formulario de creación de dispositivo, pero esta vez ofreciendo las versiones de plataforma (android) disponibles

```bash
avdm update [nombre_dispositivo]
```

no estoy seguro de como manejar esto

## Consideraciones

Este proyecto propone solucionar una problemática real ademas de servir como medio de aprendizaje de aplicaciones en GO.

Cualquier recomendación o mejora es bienvenida!i
