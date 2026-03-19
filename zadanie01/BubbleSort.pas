
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

Procedure BubbleSort(Var arr: IntArray);

Var 
  k, l, temp: integer;
Begin
  For k := 1 To 49 Do
    For l := 1 To 50 - k Do
      If arr[l] > arr[l+1] Then
        Begin
          temp := arr[l];
          arr[l] := arr[l+1];
          arr[l+1] := temp;
        End;
End;

Begin
  clrscr;
  GenerateIntegers(numbers);
  BubbleSort(numbers);

  For j:=1 To 50 Do
    write(numbers[j], ' ');
  writeln;

End.
