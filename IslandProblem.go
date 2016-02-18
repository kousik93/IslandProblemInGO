package main
import "fmt"

func countIslands(input [][]int) int {

    if len(input)==0 || len(input[0])==0{
     return 0;
    }
    count:= 0;
 
    for i:=0; i<len(input); i++ {
        for j:=0; j<len(input[0]); j++ {
            if input[i][j] == 1 {
                count++;
                merge(input, i, j);
            }
        }
    }
    return count;
}
 func merge(input [][]int, i int, j int){
    
    if i<0 || j<0 || i>len(input)-1 || j>len(input[0])-1 {

        return ;
    }
    //if current cell is water or visited
    if input[i][j] != 1 {
        return;
 }
    //set visited cell to '0'
    input[i][j] =0;

    //merge all adjacent land
    merge(input, i-1, j);
    merge(input, i+1, j);
    merge(input, i, j-1);
    merge(input, i, j+1);
}

func main() {

var input [ ][ ] int
var  sizel,sizew int


fmt.Println("Enter the size of the array(length and breadth)")

//get size
fmt.Scanf("%d", &sizel)
fmt.Scanf("%d", &sizew)

//get array
fmt.Println("Enter the array values")
input = make([][]int, sizel)
for i:=0; i<sizel; i++  {
    input[i] = make([]int, sizew)
    for j:=0;j<sizew;j++ {
        fmt.Scanf("%d", &input[i][j])
    }

}
/*
var a = [][]int{ 

{1,1,0,0,0},
{1,1,0,0,0},
{0,0,1,0,0},
{0,0,0,1,1},
 }
*/
var b int
b=countIslands(input)
fmt.Println(b)


}