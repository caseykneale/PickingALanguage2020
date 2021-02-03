vb = """
Imports System
Module Module1
    Sub Main()
        Console.WriteLine("Hello World")
    End Sub 
End Module
"""

qbasic = "PRINT \"Hello World\""

c = """
#include <stdio.h>
int main() {
    printf("Hello, World");
    return 0;
}
"""

cpp = """
#include <iostream>
int main() {
    std::cout << "Hello World";
    return 0;
}
"""

rust = """
fn main() {
    println!("Hello World");
}
"""

java = """
class HelloWorld {
    public static void main(String[] args) {
        System.out.println("Hello, World"); 
    }
}
"""

scala = """
object HelloWorld {
    def main(args: Array[String]): Unit = {
        println("Hello, world")
    }
}
"""

cs = """
namespace HelloWorld
{
    class Hello {         
        static void Main(string[] args)
        {
            System.Console.WriteLine("Hello World");
        }
    }
}
"""

fs = """
open System
[<EntryPoint>]
let main argv =
    printfn "Hello World"
    0
"""

haskell = """
main = putStrLn "Hello World"
"""

javascript = """
console.log('Hello World');
"""

fortran = """
program hello
    print *, "Hello World"
end program
"""

julia = """
println("Hello World")
"""

python = """
print("Hello World")
"""

go = """
package main
import "fmt"
func main() {
    fmt.Println("hello world")
}
"""

excel = "Hello World"
# """
# Sub hello()
#     MsgBox "Hello World"
# End Sub 
# """

r = """
print("Hello, World")
"""

matlab = """
disp('Hello world')
"""

mathematica = """
Print["Hello World"];
"""

#Yes you actually have to make a dialog box - it's horrendous
arraybasic = """
free
    dialogbeg "Hello World"
free
end 
"""

using Statistics

special_chars = only.(split(", \" ' ! @ # \$ % ^ & * ( ) - + [ ] { } ; : < > . / ? \\", " "))
white = ['\t',' ']
count_special(x) = filter( x -> (x in special_chars), x ) |> length
specials_per_line(x) = mean(count_special.( split(x, '\n') ))
count_white(x) = filter( x -> (x in white), x ) |> length
count_lines(x) = filter( x -> (x == 'n'), x ) |> length


langs = [:C,Symbol("C++"),:Go,:Rust,:FORTRAN,:Java,:Scala,:Julia,:Python,:Haskell,:JavaScript,Symbol("F#"),
            :VB,:Excel,:QBASIC,:R,:Matlab,:Mathematica,Symbol("C#"),:ArrayBasic]
hello_worlds = [c,cpp,go,rust,fortran,java,scala,julia,python,haskell,javascript,fs,
            vb,excel,qbasic,r,matlab,mathematica,cs, arraybasic]

specials = count_special.(hello_worlds)
specials_p_line = specials_per_line.(hello_worlds)
ws = count_white.(hello_worlds)
cl = count_lines.(hello_worlds)
chrs = length.(hello_worlds)
