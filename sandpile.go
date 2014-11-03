package main
import(
	"fmt"
	"os"
	"strconv"
	"time"
)

type Board struct{
	val [][]int
	size int
}
//returns a new board with given size
func CreateBoard(size int) *Board{
	board := make([][]int, size)
	for row := range board{
		board[row] = make([]int, size)
	}
	return &Board{board,size}
}
//topples (r, c) until it can’t be toppled any more.
func (b *Board) Topple(r, c int){
	if b.val[r][c] >= 4{
		if(b.Contains(r-1,c)){
			b.val[r-1][c]++
		}
		if(b.Contains(r,c-1)){
			b.val[r][c-1]++
		}
		if(b.Contains(r+1,c)){
			b.val[r+1][c]++
		}
		if(b.Contains(r,c+1)){
			b.val[r][c+1]++
		}
		b.val[r][c] -= 4
	}
}
//careful here: may cause stack overflow
func (b *Board) topple(r, c int){
	if b.val[r][c] >= 4{
		if(b.Contains(r-1,c)){
			b.val[r-1][c]++
			b.topple(r-1,c)
		}
		if(b.Contains(r,c-1)){
			b.val[r][c-1]++
			b.topple(r,c-1)
		}
		if(b.Contains(r+1,c)){
			b.val[r+1][c]++
			b.topple(r+1,c)
		}
		if(b.Contains(r,c+1)){
			b.val[r][c+1]++
			b.topple(r,c+1)
		}
		b.val[r][c] -= 4
	}
}
//returns true if (r, c) is within the field
func (b *Board) Contains(r,c int) bool{
	return r >= 0 && r < b.size && c >= 0 && c < b.size
}
//sets the value of cell (r, c)
func (b *Board) Set(r,c, value int){
	b.val[r][c] = value
}
//returns the value of the cell (r, c)
func (b *Board) Cell(r, c int) int{
	return b.val[r][c]
}
//returns true if there are no cells with ≥ 4 coins on them
func (b *Board) IsConverged() bool{
	for i := 0; i < b.NumRows(); i++{
		for j := 0 ; j < b.NumCols(); j++{
			if b.val[i][j] >= 4{
				return false
			}
		}
	}
	return true
}
//returns the number of rows on the board
func (b *Board) NumRows() int{
	return b.size
}
//returns the number of columns on the board
func (b *Board) NumCols() int{
	return b.size
} 
//compute the steady state of the board by topple. 
func (b *Board) ComputeSteadyState(){
	for !b.IsConverged(){
		for i := 0; i < b.NumRows(); i++{
			for j:= 0; j < b.NumCols(); j++{
				b.Topple(i,j)
				//b.topple(i,j)
			}
		} 
	}
}
// drawField should draw a representation of the Board 
// on a canvas and save the canvas to a PNG file with given name.
func (b *Board) DrawBoard(filename string) {
	pic := CreateNewCanvas(b.NumRows(), b.NumCols())
	pic.SetLineWidth(1)
	for i := 0; i < b.NumRows(); i++{
		for j := 0; j < b.NumCols(); j++{
			if b.val[i][j] == 3{
					pic.SetFillColor(MakeColor(255,255,255))
					drawSquare(pic,i,j)
					//DrawPoint(pic, i, j)
				}else if b.val[i][j] == 2{
					pic.SetFillColor(MakeColor(170,170,170))
					drawSquare(pic,i,j)
					//DrawPoint(pic, i, j)
				}else if b.val[i][j] == 1{
					pic.SetFillColor(MakeColor(85,85,85))
					drawSquare(pic,i,j)
					//DrawPoint(pic, i, j)
				}else{//b.val[i][j] == 0
					pic.SetFillColor(MakeColor(0,0,0))
					drawSquare(pic,i,j)
					//DrawPoint(pic, i, j)
				}
		}
	}
	pic.SaveToPNG(filename)
}
func DrawPoint(a Canvas, r, c int) {
    a.ClearRect(r,c,r+1,c+1)
}
func drawSquare(a Canvas, r, c int) {
    x1, y1 := float64(r), float64(c)
    x2, y2 := float64(r+1), float64(c+1)
    a.MoveTo(x1, y1)
    a.LineTo(x1, y2)
    a.LineTo(x2, y2)
    a.LineTo(x2, y1)
    a.LineTo(x1, y1)
    //instead of FillStroke()
    a.Fill()
}
func main(){
	// parse the command line
	if len(os.Args) != 3 {
		fmt.Println("Sorry, the # of cmd parameter is not right, try sandpile SIZE PILE")
		return
	}
	size, err1 := strconv.Atoi(os.Args[1])
	if err1 != nil ||size < 0{
		fmt.Println("Sorry, SIZE should be a positive number")
	}
	pile, err2 := strconv.Atoi(os.Args[2])
	if err2 != nil ||pile < 0{
		fmt.Println("Sorry, PILE should be a positive number")
	}
	b := CreateBoard(size)
	b.Set(size/2, size/2, pile)
	i1 := time.Now()
	b.ComputeSteadyState()
	i2 := time.Now()
	b.DrawBoard("board.png")
	i := i2.Sub(i1)
	fmt.Println("The computation runs for" , i)
	// brute force 
	// 200 10000 : 1s59
	// 200 100000 : 22s03
	// 200 1000000 : 4m8s 

}