/* La primer declaración siempre debe ser el paquete.
En caso de que en este paquete viva la función main, el paquete debe llamarse "main". */
package main

/* El paquete "fmt" es indicativo de format y se utiliza para la entrada/salida. */
import "fmt"

/* Función "main" que indica el punto de entrada de un programa en GO. */
func main() {
	/* Las variables se pueden declarar con la palabra reservada "var" y puede contener la declaración de una o más variables del mismo tipo.
	Se indica "var" luego la(s) variable(s) y, finalmente, el tipo de dato. */
	var nombre, apellido string

	/* Asignación de variables. */
	nombre = "Rogelio"
	apellido = "Ramírez"

	/* La función pública "Println" se usa para imprimir los datos y agregar un salto de línea al final.
	Esta función añade un espacio entre cada uno de los argumentos pasados. */
	fmt.Println(nombre, apellido)

	/* Las variables pueden adquirir un tipo de dato automáticamente, lo que se conoce como duck typing.
	Go analizará el dato asignado y con el operador ":=" asignará el valor y tipo de dato a la variable.
	No es necesario utilizar la plaabra reservada "var" en este caso. */
	apellidoAuto := "De Tal"
	nombreAuto := "Fulanito"
	edadAuto := 26

	/* El operador ":=" solo debe ser utilizado cuando hay al menos una variable nueva a la izquierda, es decir,
	la intrucción:
	nombreAuto := "Otro Nombre"
	causará un error de compilación, por lo que se debería utilizar la instrucción:
	nombreAuto = "Otro Nombre"
	para lograr el cambio de nombre, pero la instrucción:
	nombreAuto, otraEdad := "Fulanito", 45
	compilará y asignará el nuevo valor y creará la nueva variable con su respectivo valor sin problemas. */
	nombreAuto, otraEdad := "Fulanito", 45

	fmt.Println(nombreAuto, apellidoAuto, "tiene", edadAuto, "años y otra edad:", otraEdad)

	/* Cuando se utiliza el duck typing, se puede simplificar las declaraciones a una sola línea incluso.
	Se deben poner los nombre de las variables, seguidas del operador ":=" y, finalmente, los valores de asignación
	en el orden de declaración de las variables. */
	nombreAuto2, apellidoAuto2, edadAuto2 := "Menganito", "De la Garza", 30

	fmt.Println(nombreAuto2, apellidoAuto2, "tiene", edadAuto2, "años")
}
