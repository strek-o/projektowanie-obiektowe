
Program BubbleSort;

{$mode objfpc}

Uses crt;

Type 
  IntArray = array Of integer;

Var 
  numbers: IntArray;
  j: integer;

Procedure GenerateIntegers(Var arr: IntArray; min, max, len: integer);

Var i: integer;
Begin
  SetLength(arr, len);
  randomize;
  For i := 0 To len - 1 Do
    arr[i] := min + random(max - min + 1);
End;

Procedure BubbleSort(Var arr: IntArray);

Var 
  k, l, temp, len: integer;
Begin
  len := Length(arr);
  If len < 2 Then Exit;
  For k := 0 To len - 2 Do
    For l := 0 To len - 2 - k Do
      If arr[l] > arr[l+1] Then
        Begin
          temp := arr[l];
          arr[l] := arr[l+1];
          arr[l+1] := temp;
        End;
End;

Begin
  clrscr;
  GenerateIntegers(numbers, -100, 100, 75);
  BubbleSort(numbers);

  For j := 0 To Length(numbers) - 1 Do
    write(numbers[j], ' ');
  writeln;

  SetLength(numbers, 0);
End.
