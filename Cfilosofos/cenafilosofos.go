package main
import "fmt"
       



func FILOSOFOS(id int) {
	var estado [3] string
       estado [0]="comiendo"
       estado [1]="esperando"
       estado [2]="terminado"
       var filosofos [5] string
       filosofos [0]="Platon"
       filosofos [1]="Anaximandro"
       filosofos [2]="Aristoteles"
       filosofos [3]="Socrates"
       filosofos [4]="Tales"     
     
       
       switch id {

              case 1:
              		fmt.Println("caso1 :")
              		fmt.Println(filosofos[id-1], estado[0] )
              		fmt.Println(filosofos[id], estado[1] )
              		fmt.Println(filosofos[id+1], estado[1] )
              		fmt.Println(filosofos[id+2], estado[1] )
              		fmt.Println(filosofos[id+3], estado[1] )

              case 2:
              		fmt.Println("cas2 :")
              		fmt.Println(filosofos[id-2], estado[2] )
              		fmt.Println(filosofos[id-1], estado[0] )
              		fmt.Println(filosofos[id], estado[1] )
              		fmt.Println(filosofos[id+1], estado[1] )
              		fmt.Println(filosofos[id+2], estado[1] )
              case 3:
                    fmt.Println("caso3 :")
              		fmt.Println(filosofos[id-3], estado[2] )
              		fmt.Println(filosofos[id-2], estado[2] )
              		fmt.Println(filosofos[id-1], estado[0] )
              		fmt.Println(filosofos[id], estado[1] )
              		fmt.Println(filosofos[id+1], estado[1] )
              case 4:
                    fmt.Println("caso4 :")
              		fmt.Println(filosofos[id-4], estado[2] )
              		fmt.Println(filosofos[id-3], estado[2] )
              		fmt.Println(filosofos[id-2], estado[2] )
              		fmt.Println(filosofos[id-1], estado[0] )
              		fmt.Println(filosofos[id], estado[1] )
              case 5:
                    fmt.Println("caso5 :")
              		fmt.Println(filosofos[id-5], estado[2] )
              		fmt.Println(filosofos[id-4], estado[2] )
              		fmt.Println(filosofos[id-3], estado[2] )
              		fmt.Println(filosofos[id-2], estado[2] )
              		fmt.Println(filosofos[id-1], estado[0] )
              case 6:
                    fmt.Println("caso6 :")
              		fmt.Println(filosofos[id-6], estado[2] )
              		fmt.Println(filosofos[id-5], estado[2] )
              		fmt.Println(filosofos[id-4], estado[2] )
              		fmt.Println(filosofos[id-3], estado[2] )
              		fmt.Println(filosofos[id-2], estado[2] )	

       }

       
}

func main(){

       
       fmt.Println("CENA DE LOS FILOSOFOS")
       
       for i:=0; i<7; i++ {  FILOSOFOS(i) }
     
}