#r "nuget: MathNet.Numerics.FSharp"
open MathNet.Numerics.LinearAlgebra
open System
open FSharp.Collections

type Edge<'G> = struct
    val ID : int 
    val E : 'G option
    new (id, ?e) = {ID = id; E = e}
end

let getID (e:Edge<'w>) = e.ID

type Graph<'w> = struct
    val A : Map< int, seq< Edge< 'w > > >
    new (adjacency) = {A = adjacency}
end 

type SimpleEdge = Edge<option<obj>>
type SimpleGraph = Graph<option<obj>>
let SimpleGraph() = Graph(Map.empty)

type WeightedEdge = Edge<float>
type WeightedGraph = Graph<float>
let WeightedGraph() = Graph<float>(Map.empty)

type ActionEdge<'T> = Edge<Lazy<'T>>
type ActionGraph<'T> = Graph<Lazy<'T>>
let ActionGraph() = Graph<Lazy<'T>>(Map.empty)

type EdgeSeq<'w> = seq<Edge<'w>>
type IDEdgePair<'w> = Tuple<int, EdgeSeq<'w>>
type AdjMap<'w> = Map<int, EdgeSeq<'w>>

let addNode (g : Graph<'w>) (id : int) = 
    match (g.A.TryFind id) with
        | Some _   -> g
        | None     -> 
            Graph( g.A.Add( id, Seq.empty) )

let addNodes (g : Graph<'w>) (ids : seq<int>) = 
    Seq.fold (fun (newG:Graph<'w>) (id:int) -> addNode newG id) g ids

let removeNode (g : Graph<'w>) (id : int) = 
    match (g.A.TryFind id) with
        | Some _    -> 
            Graph( g.A.Remove( id ) )
        | None      -> g

let addEdge (g : Graph<'w>) (id : int) (edge : Edge<'w>) = 
    match (g.A.TryFind id) with
        | Some s    -> 
            Graph( g.A.Add( id, Seq.append s ( Seq.singleton edge ) ) )
        | None      -> 
            Graph( g.A.Add( id, Seq.singleton edge) )

let addEdgeToMap (a : AdjMap<'w>) (id : int) (edge : Edge<'w>) = 
    match (a.TryFind id) with
        | Some s    -> 
            a.Add( id, Seq.append s ( Seq.singleton edge ) ) 
        | None      -> 
            a.Add( id, Seq.singleton edge) 

let pivotEdge (id : int) (edge : Edge<'w>) = ( edge.ID, Edge( id, edge.E ) )

let uncurry f (a,b) = f a b

let uncurry3 f c (a,b) = f c a b

let addDiEdge (g : Graph<'w>) (id : int) (edge : Edge<'w>) = 
    uncurry3 addEdge (addEdge g id edge) (pivotEdge id edge) 

let removeEdge (g : Graph<'w>) (from_ID : int) (to_ID : int) = 
    match (g.A.TryFind from_ID) with
        | Some s    -> 
            Graph( g.A.Add( from_ID, Seq.filter (fun e -> e.ID = to_ID)  s ) )
        | None      -> g

let transpose (g : Graph<'w>) =
    let expand = Seq.map ( fun ((id:int),(es:EdgeSeq<'w>)) ->
        Seq.map ( fun (e : Edge<'w>) -> pivotEdge id e) es ) (Map.toSeq g.A)
    let tg:AdjMap<'w option> = (Seq.concat expand) |> 
        Seq.fold ( fun (m:AdjMap<'w option>) ((id:int),(e:Edge<'w option>)) ->
            addEdgeToMap m id e ) Map.empty 
    Graph(tg)

let getNeighbors (g : Graph<'w>) (id : int) =  
    match (g.A.TryFind id) with
        | Some s    -> s
        | None      -> Seq.empty

let nearestNeighbors<'a when 'a : comparison> (g : Graph<'a>) (id : int) =  
   match (g.A.TryFind id) with
        | Some s    -> Some( Seq.minBy (fun (x:Edge<'a>) -> x.ID) s )
        | None      -> None

let getNeighborCount (g : Graph<'w>) (id : int) = Seq.length(getNeighbors g id)

let isNeighbor (g : Graph<'w>) (id : int) (nb : int) =  
    match (g.A.TryFind id) with
        | Some s    -> 
            match ( Seq.tryFind (fun (e:Edge<'w>) -> e.ID = nb) s ) with
                | Some _ -> true
                | None  -> false
        | None      -> false 

let depthFirstTraversal (g : Graph<'w>) (id : int) =  
    let rec dfsRecur (g : Graph<'w>) (id : int) (path : seq<int>) =
        let nbs = Seq.map getID (getNeighbors g id) 
        if Seq.isEmpty(nbs) then
            path
        else
            Seq.fold ( fun (p:seq<int>) (nb:int) -> 
                if Seq.contains nb p then
                    p
                else
                    dfsRecur g nb (Seq.append p (seq { nb }) ) ) path nbs
    if Seq.isEmpty(getNeighbors g id) then
        Seq.empty : seq<int>
    else
        dfsRecur g id (seq{id})

let depthFirstSearch (g : Graph<'w>) (id : int) (target : int) = 
    let rec dfsRecur (g : Graph<'w>) (id : int) (path : seq<int>) = 
        let nbs = Seq.map getID (getNeighbors g id) 
        if Seq.isEmpty( nbs ) then
            path
        elif Seq.contains target nbs then
            Seq.append path ( seq{target} ) 
        else
            Seq.fold ( fun (p:seq<int>) (nb:int) -> 
                if (Seq.contains nb p) or (Seq.contains target p) then
                    p
                else
                    dfsRecur g nb (Seq.append p (seq { nb }) ) ) path nbs
    if Seq.isEmpty(getNeighbors g id) then
        Seq.empty : seq<int>
    else
        dfsRecur g id (seq{id})

let breadthFirstTraversal (g : Graph<'w>) (id : int) =
    let mutable visitted = 
        List.map ( fun ((id:int),(_)) -> (id, false) ) (Map.toList g.A) |> 
        Map.ofList |> 
        Map.add id true 
    let mutable queue = [ id ] 
    let mutable path = [ id ]
    while not( List.isEmpty queue ) do
        let f = queue.Head
        queue <- queue.Tail
        for nb in (getNeighbors g f) do
            let nbID = getID nb
            if not(visitted.Item nbID) then
                visitted <- Map.add nbID true visitted;
                queue <- queue @ [ nbID ];
                path <- path @ [nbID]
    path

let breadthFirstSearch (g : Graph<'w>) (id : int) (target : int) =
    let mutable visitted = 
        List.map ( fun ((id:int),(_)) -> (id, false) ) (Map.toList g.A) |> 
        Map.ofList |> 
        Map.add id true 
    let mutable queue = [ id ] 
    let mutable path = [ id ]
    while not( List.isEmpty queue ) do
        let f = queue.Head        
        queue <- queue.Tail
        for nb in (getNeighbors g f) do
            let nbID = getID nb
            if not(visitted.Item nbID) then
                visitted <- Map.add nbID true visitted;
                path <- path @ [nbID]
                if nbID <> target then
                    queue <- queue @ [ nbID ];
                else
                    queue <- List.empty
    if (List.rev path).Head <> target then
        List.Empty
    else
        path

let main() =
    printfn "Hello World from F#"
    let gInit = addNodes (SimpleGraph()) (seq {0..4})
    let gf = List.fold ( fun (g) (i, e) -> addEdge g i e ) gInit [ 
            (0, Edge(1, null ))
            (0, Edge(2, null ))
            (1, Edge(2, null )) 
            (2, Edge(0, null ))
            (2, Edge(3, null ))
            (3, Edge(3, null )) 
        ]
    gf
    //2 0 1 3
    //let nbs = Seq.map getID (getNeighbors gf 1)
    //Console.WriteLine (Seq.toList nbs)
    //let p = dfs gf 2
    //let p = breadthFirstSearch gf 2 1
    let p = depthFirstSearch gf 2 1
    Console.WriteLine (sprintf "%A" (Seq.toList p) )
    //Console.WriteLine (isNeighbor gf 1 2) 
    //Console.WriteLine (isNeighbor gf 1 3) 
    //let m = matrix [[ 1.0; 4.0; 7.0 ]
    //                [ 2.0; 5.0; 8.0 ]
    //                [ 3.0; 6.0; 9.0 ]]
    //let m' = m + 2.0*m
    //let a = m'.ToMatrixString()
    //Console.WriteLine a

main()
