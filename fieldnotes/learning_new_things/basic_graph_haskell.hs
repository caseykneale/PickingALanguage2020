module Main where

import Data.IntMap (IntMap)
import qualified Data.IntMap as IntMap

import Data.IntSet (IntSet)
import qualified Data.IntSet as IntSet

import Data.Maybe (Maybe)
import qualified Data.Maybe as Maybe

import qualified Data.Sequence as Seq
import Data.Foldable (toList)

import Data.List () 
import qualified Data.List as List

data Edge w = Edge{ idx::Int, wt::w } deriving (Eq, Ord) 

newtype Graph w = Graph { adjacency :: IntMap (Seq.Seq (Edge w ))} 

data EmptyEdge = EmptyEdge
emptyGraph :: Graph EmptyEdge
emptyGraph = Graph IntMap.empty

emptySimpleGraph :: Graph EmptyEdge
emptySimpleGraph = Graph IntMap.empty

graphLike :: Graph w1 -> Graph w2
graphLike (Graph adj) = 
    case tKeys of
        [] -> Graph tAdj 
        _ -> Graph (foldr (`IntMap.insert` Seq.Empty) tAdj tKeys)
    where
        tAdj = adjacency (Graph IntMap.empty) 
        tKeys = IntMap.keys adj

addNodeToGraph :: IntMap.Key -> Graph w -> Graph w
addNodeToGraph idx (Graph adjacency) = Graph(IntMap.insert idx Seq.empty adjacency)

addNodesToGraph :: Foldable t => t IntMap.Key -> Graph w -> Graph w
addNodesToGraph idxs graph = foldr fn graph idxs
  where fn idx (Graph adjacency) = Graph (IntMap.insert idx Seq.empty adjacency)

{-- Note: these fn's looksup 2x --}
addEdgeToNode :: IntMap.Key -> Edge w -> Graph w -> Graph w
addEdgeToNode node (Edge idx wt) (Graph adjacency) = case IntMap.lookup idx adjacency of
            Just foundIt -> Graph (IntMap.adjust (Edge idx wt Seq.<|) node adjacency)
            Nothing -> error "Node does not exist in graph." 

addEdgesToNode :: IntMap.Key -> [Edge w] -> Graph w -> Graph w
addEdgesToNode node edgeList (Graph adjacency) = case IntMap.lookup node adjacency of
            Just foundIt -> Graph (IntMap.adjust (Seq.>< Seq.fromList edgeList) node adjacency )
            Nothing -> error "Node does not exist in graph." 

addBiEdgeToNode :: IntMap.Key -> Edge w -> Graph w -> Graph w
addBiEdgeToNode node (Edge edgenode wt) graph = addEdgeToNode edgenode (Edge node wt) 
                                                    ( addEdgeToNode node (Edge edgenode wt) graph )

weights :: [Edge b] -> [b]
weights = map wt

flipEdgeID :: Int -> Edge w -> (Int, Edge w)
flipEdgeID node (Edge id wt) = (id, Edge node wt)
flipEdge :: Int -> Edge w -> Edge w
flipEdge node (Edge id wt) = Edge node wt

flipEdgesAt :: Int -> [Edge w] -> ([Int], [Edge w])
flipEdgesAt node edgeSeqs = (fmap idx edgeSeqs, map (Edge node) (fmap wt edgeSeqs))

flipToFlatList :: Int -> Seq.Seq (Edge w) -> [(Int, Edge w)]
flipToFlatList k v = let lstV = Data.Foldable.toList v
                in fmap (\a -> (idx a, Edge k $ wt a)) lstV

transpose :: Graph w -> Graph w
transpose (Graph adj) = 
    let tAdj = adjacency ( graphLike (Graph adj) ) 
        tAdjLst = concatMap (uncurry flipToFlatList) (IntMap.assocs adj)
    in Graph (foldr (\(id,e) -> IntMap.adjust (e Seq.<|) id) tAdj tAdjLst)
{--

addBiEdgesToNode :: IntMap.Key -> [Edge w] -> Graph w -> Graph w
addBiEdgesToNode node edgeList (Graph adjacency) = case IntMap.lookup node adjacency of
            Just foundIt -> Graph (IntMap.adjust (Seq.>< Seq.fromList edgeList) node adjacency )
            Nothing -> error "Node does not exist in graph." 
--}
removeNode :: IntMap.Key -> Graph w -> IntMap (Seq.Seq (Edge w))
removeNode idx (Graph adjacency) = case IntMap.lookup idx adjacency of
            Just foundIt -> IntMap.delete idx adjacency
            Nothing -> error "Could not remove Node from graph, it does not exist." 

removeNodeFromEdges :: Graph w -> Int -> Graph w
removeNodeFromEdges (Graph adj) id =
    let kvp = IntMap.assocs adj
    in
        Graph (
            foldr ( 
                \(k, vseq) -> IntMap.adjust (
                    return (Seq.filter (\v -> id /= idx v) vseq )
                ) k
            ) adj kvp
        )

getNeighbors :: Graph w -> Int -> Seq.Seq (Edge w)
getNeighbors (Graph adjacency) id = Maybe.fromMaybe Seq.empty (IntMap.lookup id adjacency)

getNeighborIDs :: Graph w -> Int -> Seq.Seq Int
getNeighborIDs (Graph adjacency) id  = fmap idx ( getNeighbors (Graph adjacency) id )

getNeighborCount :: Graph w -> IntMap.Key -> Int
getNeighborCount (Graph adjacency) id = case IntMap.lookup id adjacency of
            Just foundIt -> length foundIt
            Nothing -> error "Could find Node in graph, it does not exist."

--Alias
outDegree :: Graph w -> IntSet.Key -> Int
outDegree = getNeighborCount

inDegree :: Graph w -> IntSet.Key -> Int
inDegree w = getNeighborCount (transpose w) 

isNeighbor :: Graph w -> Int -> Int -> Bool
isNeighbor (Graph adjacency) idA idB = 
        case result of
            -1 -> False 
            _ -> True 
    where
        nbIDs = getNeighborIDs (Graph adjacency) idA
        result = Maybe.fromMaybe (-1) (Seq.findIndexL ( == idB) nbIDs)

--ToDo: 
getNearestNeighbor :: Ord a => Edge a -> [Edge a] -> Edge a
getNearestNeighbor e [] = e
getNearestNeighbor e xs = foldl (\ e x -> if wt x > wt e then e else x) e xs

depthFirstSearchM :: Graph w -> Int -> IntSet -> Seq.Seq Int -> Seq.Seq Int 
depthFirstSearchM (Graph adj) startID elementSet pathSeq = 
    case lookupSeq of 
        Seq.Empty -> pathSeq
        _ ->  do 
            idxAvail <- lookupSeq --every individual index
            let actID = Maybe.fromMaybe (error "your memory is corrupt?") (Seq.lookup idxAvail nbIDs) 
            pathSeq Seq.>< depthFirstSearchM (Graph adj) actID 
                (IntSet.insert actID elementSet) (Seq.singleton actID)
    where
        nbIDs = getNeighborIDs (Graph adj) startID --Sequence of neighbors at node
        lookupVisited = Seq.findIndicesL (`IntSet.notMember` elementSet) nbIDs --list of neighbors indices not at node
        lookupSeq = Seq.fromList lookupVisited

depthFirstSearch :: Graph w -> Int -> [Int]
depthFirstSearch (Graph adj) startidx = 
    List.nub (
        toList (
            depthFirstSearchM (Graph adj) startidx (IntSet.singleton startidx) (Seq.singleton startidx)
        )
    )


bfsM (Graph adj) startID targetID elementSet pathSeq = 
    case lookupSeq of 
        Seq.Empty -> pathSeq
        _ ->  do 
            idxAvail <- lookupSeq --every individual index
            let actID = Maybe.fromMaybe (error "your memory is corrupt?") (Seq.lookup idxAvail nbIDs) 
            pathSeq Seq.>< depthFirstSearchM (Graph adj) actID 
                (IntSet.insert actID elementSet) (Seq.singleton actID)
    where
        nbIDs = getNeighborIDs (Graph adj) startID --Sequence of neighbors at node
        lookupVisited = Seq.findIndicesL (`IntSet.notMember` elementSet) nbIDs --list of neighbors indices not at node
        lookupSeq = Seq.fromList lookupVisited

bfs (Graph adj) startidx target = 
            --dijkstrasM (Graph adj) startidx (IntSet.singleton startidx) (Seq.singleton startidx)

main :: IO ()
main = do
    let gInit = addNodesToGraph [1..4] emptyGraph
        g = foldr (\f x -> f x) gInit [
                    addEdgeToNode 1 (Edge 2 EmptyEdge), 
                    addEdgeToNode 1 (Edge 3 EmptyEdge), 
                    addEdgeToNode 2 (Edge 3 EmptyEdge), 
                    addEdgeToNode 3 (Edge 1 EmptyEdge),
                    addEdgeToNode 3 (Edge 4 EmptyEdge),
                    addEdgeToNode 4 (Edge 4 EmptyEdge) ] 
    let path = depthFirstSearch g 3
    print "Hi World!"
    print path
    print (isNeighbor g 1 2)
    print (isNeighbor g 1 4)
    print "Bye World!"

    let gInit = addNodesToGraph [1..4] emptyGraph
        g = foldr (\f x -> f x) gInit [
                    addEdgeToNode 1 (Edge 2 EmptyEdge), 
                    addEdgeToNode 2 (Edge 3 EmptyEdge), 
                    addEdgeToNode 3 (Edge 4 EmptyEdge),
                    addEdgeToNode 2 (Edge 4 EmptyEdge) ] 
        gi = removeNodeFromEdges g 3
    let path = depthFirstSearch gi 1
    print path
    print "cookies"
    let gInit = addNodesToGraph [1..3] emptyGraph
        g = foldr (\f x -> f x) gInit [
                    addEdgeToNode 1 (Edge 2 EmptyEdge), 
                    addEdgeToNode 2 (Edge 3 EmptyEdge)]
    let path1 = depthFirstSearch g 1
    print path1
    let tg = transpose g
        path2 = depthFirstSearch tg 3
    print path2
    --print tg
    print "Bye World!"
