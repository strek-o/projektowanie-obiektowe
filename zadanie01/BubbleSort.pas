
Program BubbleSort;

Uses crt;

Type 
  IntArray = array[1..50] Of integer;

Var 
  numbers: IntArray;
  j: integer;

Procedure GenerateIntegers(Var arr: IntArray);

Var i: integer;
Begin
  randomize;
  For i := 1 To 50 Do
    arr[i] := random(101);
End;

Begin
  clrscr;
  GenerateIntegers(numbers);

  For j:=1 To 50 Do
    write(numbers[j], ' ');
  writeln;

End.
