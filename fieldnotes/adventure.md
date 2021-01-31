Every year or so I do a technical scouting excercise to assess programming languages and similar tools. I have a few "masterpiece" projects that have plagued me over the past 5-10 years. I need very good tooling to work on them by myself, in my free-time. Hacking them up in a scripting language just won't suffice, they are big dreams. So I asked around my primary communities (Julia and Functional Programming communities) "what languages are around that I should try out?". I was directed mostly to Rust, but a few other's.

This post is basically some *personal* anecdotes and code of some things I hacked together while exploring a few new programming languages this year. The code mentioned in each section is hosted [here](https://github.com/caseykneale/PickingALanguage2020/tree/main/fieldnotes/learning_new_things) in it's entirety. Key snippits will be shared as I develop the narrative. The focus here isn't "look I mastered 4 languages and my code is perfect" - pretty much the opposite of that. The idea is to share my experience as someone who had only a short exposure to a variety of a few programming languages that were new to me.

## Rust - Tinkering
I had heard a lot about Rust over the years but assumed it hadn't earned much of a standing when compared to C++. I set out, skim-reading the [Rustinomicon](https://doc.rust-lang.org/nomicon/), and then trying to dive right into writing some code. Thats what I did with Julia, R, Python, C#, C++, etc, so why not? 

I got burned quick. I couldn't get the compiler to accept 99% of my code. I felt like I knew nothing. I was frusterated. It felt so alien to me despite all of the core concepts making complete sense - even sounding brilliant. I gave up on it after 2-3 afternoons of tinkering. **note: my opinion on this changes completely after revisitting it later. You have to keep reading to see that though**

## Go - DIY GUI
Ah-hah the Google Programming language with the cute gopher mascot. Most of the chatter about Go complained that the language was "too simple". But it supported higher order functions, had structs, some notion of pointers/references - seems nice on paper to me. I figured it would be a bit like C, and worked through the [Go tour](https://tour.golang.org/welcome/1). The tour had an interactive online compiler - it felt modern, fast, and nice.

I decided to jump in and do something simple... How about write an algebriacly structured GUI library from SDL2? Put Go to the test, while not exhausting myself.  Also Go's GUI libraries are... a work in progress for the most part. So if I'm going to escape the command-line or backend work without touching JS - I'll need a GUI API I enjoy. Again this target mostly for fun, but it seemed small enough in scope to test my understanding.

I found writing Go, very easy, like boring easy. The syntax was simple. In some cases a little too simple. Spent a fair deal of time grappling with trying to concoct reasonable trait abstractions, but nothing terrible. The code compiled fast. At first I was in love. I saw this as a language with a competitive 80-20. The number of footguns are comparitively small considering what it let's you do. 

Then some wart's started to appear. Iota instead of enumerations? Eh okay. Reference system had me flipping my API around pretty regularly, and made thoughts of multiprocessing more daunting than what I thought was on the label. I really really wanted some Generics, but they just weren't on the languages radar, so I kind of diverted course. 

Over-all, it was "easy" to learn, it is "fast" to write, the compiler was "fast" to compile, the binaries were "smallish" and the code ran "kind of fast". For reference I was easily netting 2k FPS on a simple window with a label and a button + event handling in my own mini GUI library on a cheap laptop. The API I wrote was okay, mostly as intended. Not bad but, not great. In some sense, that's a positive thing. Who hasn't seen someone get a little too creative on a project?
```go
    ...
    var mainWindow = NewSDLEnvironment(winTitle, winWidth, winHeight)

	var mainScene = NewScene()
	mainScene.AddGlobalChain(GlobalChain{ExitProgram})
	mainScene.AddGlobalChain(GlobalChain{MouseButton, DeactivateAll})

	var btnTxt = NewText(mainWindow, 0, 0, "Submit")
	var btnFace = NewPane(mainWindow, Point{32, 32}, 64, 32)
	var newp = mainWindow.NewButton(btnFace, btnTxt)
	mainScene.AddActive(newp)

	mainScene.AddChain(WidgetChain{MouseDown})
	mainScene.AddChain(WidgetChain{MouseUp, customFunction})

	var sigTxt = NewText(mainWindow, 0, 0, "Enter Text Here")
	var sigLine = NewHLine(mainWindow, Point{128, 32}, 96, 2)
	var newt = mainWindow.NewSignature(sigLine, sigTxt)
	mainScene.AddActive(newt)
    ...
```
The full code for this is available [here](https://github.com/caseykneale/PickingALanguage2020/tree/main/fieldnotes/learning_new_things/Go). You'll have to add your own font files though - don't have time to see if I can freely distribute those...

## Change of strategy - Graph Library
At this point I decided making a GUI library from OpenGl/SDL was maybe too laborious as a test. I'll admit I thought Go was going to be my winner after striking out with Rust so early on. I put a week or so of effort into Go off the bat. After deciding to broaden my search scope I decided to change the target task to something more simplistic. Something that tests data structures, control flow, and algorithms? I settled on a basic Graph library (no not a "plotting" library, a [graph library](https://en.wikipedia.org/wiki/Graph_theory))?

The goalposts I set out to try to solve were: generic support for integer indices(ie: int8, int16, int32, etc), sparsely annotated graphs (noncontiguous node labels), optional edges (maybe weights, maybe thunked functions, who knows) and have the ability to write easily write basic traversal algorithms. That's pretty much it. In Julia those tasks are all trivial to do first pass. Harder to do with performance in mind, but that's not what I'm doing here, just learning and assessing.

## Haskell - Graph API
Screw it. These systems programming languages are missing the Functional Programming(FP) charm I've grown to adore in multiparadigm languages like Julia. Rust was hard - I bet learning Haskell is easier. Why not go full FP?

Haskell was one of the first languages that when I read the through the secondary sources for tutorials made complete sense to me. Then I'd go to read the docs and feel like I crash-landed on an alien planet. There's a chicken and egg problem in that, you will not understand the documentation until your understand the language. Start small and work your way up.

If you can get past that, and learn some of the syntax from the great resources that are freely available you'll find something beautiful. I highly reccommend any books by [Stephen Diehl](http://dev.stephendiehl.com/fun/WYAH.pdf), or [learn you a haskell](http://learnyouahaskell.com/) for learning materials. I mean, the language is everything you'd want in a hobby, it makes you think, then it let's you hack, VSCode autocompletes the boilerplate in a way that only a very well thought out language could. Was a real treat.

I wrote my first Monad intentionally in Haskell writing a Depth First Search, 
```haskell
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
```
Then I realized something troubling... I couldn't think of a sound way to write a BreadthFirstSearch in a purely functional way. I mean, I could, but implementing it was awful. It turns out, after some googling, this sort of thing has earned people Journal Articles. Yikes. Let's remember BFS/DFS traversals in C++ are maybe 30 lines of code at a maximum. A common undergraduate excercise... That's when it hit me... 

Haskell is kind of *academic*. Sure, for some tasks(compilers, etc), it is a gold standard for production. For desktop applications and my interests it was more of a gymnasium than a trusty toolchest. Maybe I knew this going into it... The lack of common/easy tools to staticly compile Haskell is an obvious pathology.

Okay so, Haskell wasn't my cup of tea. Despite having only spent 2-3 weeks with it I'll always have a fond reverence for its ambition. Right up there with LISP. To be fair I did find learning basic Haskell easier than learning basic Rust. That said... I can gaurantee difficult haskell is exponentially more difficult than difficult Rust! Anyways, with monads and a few other doo-dads under my belt, I decided something more imperitive would make life easy.

## F# - Graph API
F# was interesting,I was using a *.NET* language on Ubuntu. That's something that in 2015 I didn't think would ever happen. The support for the language via VSCode was quite good. Writing a BFS was trivial because while loops and side effects exist in F#,
```fsharp
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
```
It was truly refreshing to have that bit of imperitive code style available. Felt more like home. But then I started asking myself tougher questions... How is concurrency going to go? How much of a performance penalty am I going to pay if I try to stand up a **real** API in F#? I was really hoping to lose the Garbage Collector... Why learn something new that's slower than C++ afterall? Why did learning F# not feel like I really learned a programming language? Maybe I didn't spend enough time with it. Maybe I didn't want too?

## Rust - Graph API
I wound up stairing down Rust again. This time though, something was different. Now that I had bashed around with traits in Go, Monad's/Maybes,etc in Haskell and dealt with several package managers (ghci, packet, go modules, etc) Rust looked like a thing of beauty.

Rust instantaneously became easy to write, like a step above python or something. I told some of my friends I was really enjoying it - they said things like "wait until you deal with the borrow checker". I didn't even know what they meant? I had been dealing with it, and the compiler was outrageously helpful. It was like writing C++, with someone making sure you aren't shooting yourself in the foot before you even compiled.

I wrote the basics of a graph API in a single afternoon (with some help from the community on some syntax gotchyas). The code was overall pretty sloppy, but, I love choosing between pass by reference and pass by value, mutability, etc. It makes tasks like recursion with mutation easy and clean enough,
```rust
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
```

Rust felt like home. It has FP, but it doesn't box you into any corners like removing side effects. It has OOP, but it's only enough to feel like composition(IE a pipe operator). Nothing like the death by abstraction that requires annual rewrites in many enterprise Java codebases. It is procedural, but only in a way that let's you move onto your next thought, not something that belabours it's own purpose. Sweet-spot.

Sure the compiler is slow when projects get large but it seemed faster than many Scala projects I've worked on atleast. Yea, writing Rust isn't as fast/easy as say Julia is for mathematics, but neither is C++, even with a library. Isn't this what systems level programming is supposed to be? Carefully considering how to orchestrate tasks for a computer to do at a low level. Yet... not being kept up at night, by flashes of race conditions, memory leaks looming on line 20k of some random legacy file. Rust feels like the future to me - when I realized all of the problems it solved for me - I suffered from 2 days of blank canvas intimidation. Not fun, but on the whole, a net positive.

## Conclusions
There are no conclusions here, just options and personal opinions.

## Opinions
I am partial to Rust and Go right now for the tasks that suit them best. I admire Haskell, and to some extent F#. My big takeaway was that trying different languages taught me to appreciate the languages I already regularly use (Julia and C++ mostly). It also made learning more languages easier to do! Learning haskell somehow made Rust feel like a modern but more friendly C++. 

One of my graduate advisors, once told me "The meanest thing you could ever do to someone, is tell them that they can do something that they can't,and then let them do it." I've personally lived by this... I've had to make hard calls here and there that seemed *mean* to some people, but the truth is - the alternative was far-far meaner. Rust is the same way. C and C++ will make you feel like you are a superhero until you hit a segfault 2 weeks into production and spend 2 more unanticipated weeks debugging a subtle casting/pointer issue! 

I think I will be using less and less C++ and more and more Rust. Maybe Go for some small projects with short deadlines though - why not?
