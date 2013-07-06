/**
 * Created with IntelliJ IDEA.
 * User: Ganaraj
 * Date: 06/07/13
 * Time: 10:52
 * To change this template use File | Settings | File Templates.
 */
package main

import "fmt"

func main() {

	fmt.Printf("Hello world!")

}

type FunctionDef struct{
	name string
	returnType string
	arguments []string
};

func GoFunctionSearch(args...string) (string,[]FunctionDef){

	myfunction := FunctionDef{name:"abc",returnType :"string"}
	funcdefs := make([]FunctionDef,1)
	funcdefs[0] = myfunction
	return myfunction.name,funcdefs
}
