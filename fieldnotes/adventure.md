# Intro
I figured I'd share some notes/code I hacked together while exploring some new programming languages this year. The code is hosted [here](https://github.com/caseykneale/PickingALanguage2020/tree/main/fieldnotes/learning_new_things) and key snippits will be shared as I develop the narrative. The focus here isn't "look I mastered 4 languages and my code is perfect" - pretty much the opposite of that. The idea is to share my experience as, someone with exposure to a variety of other programming languages, and what I encountered while learning enough about them to decide if I wanted to learn more.

## How it all started
Every year or so I do a technical scouting of programming languages. I have a few "masterpiece" projects that have plagued me over the past 5-10 years. I need very good tooling to build them by myself in my free-time. Hacking them up in a scripting language just won't suffice. So I asked around my primary communities (Julia and Functional Programming communities) "what should I try out?". I was directed mostly to Rust, but a few other's.

## Rust
I had heard a lot about Rust over the years but assumed it hadn't earned much of a standing when compared to C++. I set out, reading the Rustinomicon, writing some basic doohickies, and then trying to dive right into writing some code. Thats what I did with Julia, R, Python, C#, etc, so why not? 

I got burned quick. I couldn't get the compiler to accept 99% of my code. I felt like I knew nothing. I was frusterated. It felt so alien to me despite all of the core concepts making sense. I gave up on it after 2-3 afternoons of tinkering. **note: my opinion on this changes completely after revisitting it later**

## Go - DIY GUI
Ah-hah the Google Programming language with the cute gopher. Most of the chatter complained that the language was "too simple". But it supported higher order functions, had structs, some notion of pointers/references. I figured it would be a bit like C, and worked through the [Go tour](https://tour.golang.org/welcome/1). Mostly online too - it felt modern.

I decided to jump in. Let's do something simple... How about write an algebriacly structured GUI library from SDL2? Put it to the test, while not exhausting myself.  Also Go's GUI libraries are... a work in progress for the most part. So if I'm going to escape command-line or backend work without touching JS - I'll need a GUI I enjoy. Again a target mostly for fun, but seemed small enough in scope to test my understanding.

I found writing Go, surprisingly... Easy. The syntax was simple. In some cases a little too simple. I couldn't write generics. Spent a fair deal of time grappling with reasonable trait abstractions, but nothing terrible. The code compiled fast. At first I was in love. I saw this as a language with a competitive 80-20.The one great thing about Go is, the number of footguns is comparitively small considering what it let's you do. 

Then some wart's started to appear. Iota instead of enumerations? Eh okay. Reference system had me flipping my API around pretty regularly, and made thoughts of multiprocessing pretty daunting. Tried to make some widget's and really really wanted some Generics, but they just weren't on the languages radar. 

Over-all, it was "easy" to learn, it is "fast" to write, the compiler was "fast" to compile, the binaries were "smallish" and the code ran "kind of fast". For reference I was easily netting 2k FPS on a simple window with a label and a button + event handling in my own mini GUI library on a cheap laptop. Not bad, but, not great. 

## Change of strategy - Graph Library
At this point I decided making a GUI library from OpenGl/SDL was maybe too laborious as a test. Why not something simple to test data structures, control flow, and algorithms? How about a basic Graph library(no not a "plotting" library, a [graph library](https://en.wikipedia.org/wiki/Graph_theory)) with a few goals.

### Goals
- Generic support for integers indices
- Sparsely annotated graphs (noncontiguous indices)
- Optional edges (maybe weights, maybe thunked functions, who knows)
- The ability to write basic traversal algorithms

## Haskell
I said screw it, these systems programming languages are missing the Functional Programming(FP) charm I've grown to adore in multiparadigm languages like Julia. Why not go full FP? Let's learn some Haskell!

Haskell was one of the first languages that when I read the through the secondary sources for tutorials made complete sense to me. okay, everything is a function, easy. Then I'd go to read the docs and feel like I crashlanded on an alien planet. There's a chicken and egg problem in that, you will not understand the documentation until your understand the language.

If you can get past that, and learn some of the syntax from the great resources that are freely available(any books by (Stephen Diehl)[http://dev.stephendiehl.com/fun/WYAH.pdf], or (learn you a haskell)[http://learnyouahaskell.com/]) you'll find something beautiful. I mean, the language is everything you'd want in a hobby, it makes you think, then it let's you hack, VSCode autocompletes the boilerplate in a way that only a very well thought out language could. Was a real treat.

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
Then I realized something troubling... I couldn't think of a sound way to write a BreadthFirstSearch in a purely functional way. I mean, I could, but implementing it was awful. It turns out, after some googling, this sort of thing has earned people Journal Articles. Yikes. Let's remember a basic BFS/DFS traversal in C++ is maybe 100 lines of code as a maximum. An undergraduate excercise. That's when it hit me... Haskell is kind of *academic*. Sure, for some tasks(compilers, etc), it is a gold standard for production. For desktop applications, and my interest, though, well... It's more of a gymnasium than it was a trusty toolchest.

Okay so, Haskell wasn't my cup of tea. Despite having only spent 2-3 weeks with it I'll always have a fond reverence for its ambition. With monads and a few other doo-dads under my belt, I decided something more imperitive would make life easy. What a blast though.

## F#



## Rust

