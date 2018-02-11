// Esto es un comentario de línea, termina hasta el final de la línea
// Los comentarios en línea se utilzan para documentar el código
/*
  Esto es un comentario de bloque.
  Permite tener un comentario a través de varias líneas.
*/
/*
  Los comentarios de bloque se utilizan para comentar el código.
  Puede servir para que otro desarrollador lo vea (vea un ejemplo de uso),
  por ejemplo:
  var a in
  a = 8
*/
package main

import "fmt"

func main() {
	/* Una constante se declara con la palabra reservada "const" y el valor se le debe asignar de inmediato.
	El compilador sabrá de que tipo es dicha constante. */
	const nombre = "Rogelio"

	/* La siguiente sentencia producirá un error poeque una constante no puede cambiar de valor */
	// nombre = "Juan"

	fmt.Println(nombre)
	/*************************/

	/* int es el alias para int64 en arquitectura de 64 bits o para int32 es arquitectura de 32 bits */
	var a int
	var b int8 = 10
	n := "Pedro"

	a = 121212
	b = 5

	/* Go es muy estricto con el tipo de dato de cada variable, la siguiente línea daría error: */
	//c := a + b

	/* En lugar de tener este tipo de error, se recurre al casting o cambio temporal de tipo de dato */
	c := a + int(b)
	fmt.Println(c)
	fmt.Printf("hola %s ¿cómo estás?, c = %d\n", n, c)
	fmt.Printf("c es de tipo %T\n", c)

	// Prioridad aritmética
	// + - * /   (incorrecta)
	// () % * / + -   (correcta)

	// 6 + (6 * 6) - 6
	d := 6 + 6*6 - 6
	fmt.Println(d)
	/*****************************/
	// Valor cero o inicial de una variable
	// string -> ""
	// int -> 0
	// bool -> false
	var nombre3 string
	var numero int
	var entiendes bool

	fmt.Printf("string: <%s>, int: <%d>, bool: <%t>", nombre3, numero, entiendes)
}
