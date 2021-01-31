use std::collections::HashMap;
//use std::collections::Vec;
use std::hash::Hash;
use queues::*;


#[derive(Copy, Clone, Default)]
struct Edge<U:Eq + Hash + Copy, T: Default + Copy = ()> {
    id: U,
    x: T,
}

impl<U:Eq + Hash + Copy, T: Default + Copy > Edge<U, T> {
    fn new(id : U) -> Self {
        return Edge { id: id, x: T::default() };
    }
}

#[derive(Clone, Default)]
struct Graph<U:Eq + Hash + Copy,T: Default + Copy = ()> {
    adjacency: HashMap<U, Vec<Edge<U,T>> >
}

impl<U:Eq + Hash + Copy, T: Default + Copy> Graph<U,T> {
    fn new() -> Self {
        Self { adjacency: HashMap::new() }
    }
    fn add_node(&mut self, node:U) {
        self.adjacency.entry(node).or_insert_with(|| Vec::new());
    }
    fn add_nodes(&mut self, mut nodes:Vec<U>) {
        for node in nodes.drain(..) {
            self.add_node(node);
        }
    }
    fn add_edge(&mut self, id:U, node:Edge<U,T>) {
        (*self.adjacency.entry(id).or_insert( Vec::new() )).push(node);
    }
    fn add_edges(&mut self, mut ids:Vec<U>, mut nodes:Vec<Edge<U,T>>) {
        for (id,node) in ids.drain(..).zip(nodes.drain(..)){
            self.add_edge(id, node)
        }
    }
    fn add_edge_pairs(&mut self, mut id_node_pairs:Vec<(U, Edge<U,T>)>) {
        for (id,node) in id_node_pairs.drain(..){
            self.add_edge(id, node)
        }
    }

    fn get_neighbors(&self, id: U) -> Option<&Vec<Edge<U,T>>> {
        self.adjacency.get(&id)
    }

    fn get_neighbor_ids(&self, id: U) -> Vec<U> {
        match self.adjacency.get(&id) {
            Some(v) => v.iter().map(|n| n.id).collect::<Vec<U>>(),
            _ => Vec::new()
        }
    }

    fn is_neighbor(&self, id: U) -> bool {
        match self.adjacency.get(&id) {
            Some(v) => v.iter().filter(|&n| n.id == id).count() > 0,
            _ => false
        }
    }

    fn get_neighbor_counts(&self, id: U) -> usize {
        match self.adjacency.get(&id) {
            Some(v) => v.len(),
            _ => 0
        }
    }

    fn transpose(&self) -> Graph<U,T> {
        let mut g = Graph::new();
        for (key, _) in &self.adjacency {//Copy keys over
            g.add_node(*key)
        }
        for (key, values) in &self.adjacency{
            for value in values {//mix and add
                g.add_edge(value.id, Edge{ id : *key, x: value.x });
            }
        }
        return g;
    }

    fn breadth_first_traversal(&self, start:U) -> Vec<U> {
        let mut visited = HashMap::new();
        for (key, _) in &self.adjacency {//Copy keys over
            visited.insert( key, false );
        }
        //check that node exists..
        *visited.get_mut(&start).unwrap() = true;
        let mut q: Queue<U> = queue![start];
        let mut path = Vec::new();
        path.push(start);
        while q.size() > 0 {
            let s = q.remove();
            let nb_ids = match s {
                Ok(cur_id) => self.get_neighbor_ids(cur_id),
                Err(_) => Vec::new(),
            };
            for nb_id in nb_ids {
                if !visited[&nb_id]{
                    *visited.get_mut(&nb_id).unwrap() = true;
                    q.add(nb_id);
                    path.push(nb_id);
                }
            }
        }
        return path;
    }

    fn depth_first_traversal_recurse(&self, id:U, v: &mut HashMap<&U, bool>, path:&mut Vec<U>) { 
        let nbs = self.get_neighbor_ids(id);
        if nbs.len() == 0 {
            path.push(id);
        } else {
            for nb_id in nbs.into_iter() {
                if !v[&nb_id] {
                    *v.get_mut(&nb_id).unwrap() = true;
                    path.push(nb_id);
                    self.depth_first_traversal_recurse( nb_id, v, path );
                }
            }
        }
        return;
    }

    fn depth_first_traversal(&self, start:U) -> Vec<U> {
        let mut visited = HashMap::new();
        for (key, _) in &self.adjacency {//Copy keys over
            visited.insert( key, false );
        }
        *visited.get_mut(&start).unwrap() = true;
        let mut path = Vec::new();
        path.push(start);
        self.depth_first_traversal_recurse(start, &mut visited, &mut path);
        return path;
    }
}


// println!("Hello, world!");
// let mut g = Graph::<i32>::new();
// g.add_nodes(vec![ 0,1,2,3 ]);
// g.add_edge_pairs( vec![ (0, Edge::<i32>::new(1) ),
//                         (0, Edge::<i32>::new(2) ),
//                         (1, Edge::<i32>::new(2) ),
//                         (2, Edge::<i32>::new(0) ),
//                         (2, Edge::<i32>::new(3) ),
//                         (3, Edge::<i32>::new(3) ),
//  ] );

//  let nb = g.get_neighbor_ids(1);
//  println!("{:?}",nb);

//  let bfs = g.breadth_first_traversal(2);
//  println!("{:?}",bfs);

//  let dfs = g.depth_first_traversal(2);
//  println!("{:?}",dfs);
