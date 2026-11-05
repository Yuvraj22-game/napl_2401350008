// Higher-order functions using pure lambda expressions.

let mapList values =
    List.map (fun number -> number * number) values

let filterList values =
    List.filter (fun number -> number % 2 = 0) values

let sumList values =
    List.fold (fun total number -> total + number) 0 values

let firstList = [ 1; 2; 3; 4; 5 ]
let secondList = [ 6; 7; 8; 9; 10 ]

printfn "List.map - squares:"
printfn "%A -> %A" firstList (mapList firstList)
printfn "%A -> %A" secondList (mapList secondList)

printfn "\nList.filter - even numbers:"
printfn "%A -> %A" firstList (filterList firstList)
printfn "%A -> %A" secondList (filterList secondList)

printfn "\nList.fold - sums:"
printfn "%A -> %d" firstList (sumList firstList)
printfn "%A -> %d" secondList (sumList secondList)