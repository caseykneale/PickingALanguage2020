# Picking a Programming Language in 2020? (Opinion)

I'm an unabashed data geek working in industry who specializes in *dealing with* chemical data. I wear hats that span the set from analytical chemist, data scientist, and software developer. With a little bit of everything in between. Over the years I've written hundreds of thousands of lines of code. Some of it quite good, some of it so bad that I hope there are no traces of it on the internet. I do this both as a hobby and as a means of putting food on the table.

I have worked alongside high-flying data scientists - you know the geniuses who can scrap together a new algorithm using mathematics that wasn't known to exist and raise them to production on time. I have also worked alongside low-level systems programmers - the type of people who giggle when their code pushes servers to literally smoke and catch fire (true story). For some developers/scientists, like myself, the line between these two groups gets a bit blurry. I figured I'd offer a mediated opinion on programming language scouting. 

I often see there is a significant disparity in tooling that can cause schisms in teams, and therefore deliverables/products. I also tend to see a lot of hype pouring from influencer types that really muddies the waters for people looking for straightforward answers. Thankfully beneath the curtain of advertising and self-promotion, the water really isn't all that murky.

# High-level Starting Point

This section serves to detail some items to keep in mind as you drill down the cyclic process of settling on a programming language for a project.

## Assess Resources

**Technical ability:** 

How savvy is the team with computer science/development? Can they learn a new language on their own? How long might that take given the complexity of a tool? 

**Incentive:**

Is there an incentive to offer a **change** to be put in place? Is continuous learning a cultural norm? Are there selling points to make a change a "net positive" for the people involved in the project?

**Safety:** 

Is there a culture of psychological safety and merit-based reward systems? If not, trying new things can scare people away. 

## Define the problem statement

**Pick Target Platform(s):** 

Desktop/Embedded/Cloud/On Premises/etc. What do the users need, make a survey/do some research, talk to marketing? 

**Identify Sources of Novelty:**

Is your killer feature a huge distributed computing web framework behind a super clean easy to use front end? Might need to look at languages supporting stable parallelism, and their ecosystems. Maybe you need stability, your competitor offerings are very glitchy, look into type safety. Maybe you need a highly interactive GUI - what API's are accessible to the languages.

**Identify Technical Need:**

Items to keep in mind: performance, mathematics, networking, parallelism, beauty, maintainability(short-term/long-term).

## Estimate Cost

Budget for closed source languages (matlab, etc), for training new staff, licenses are needed to use 3rd party tools, etc. If you are using on premises hard-ware remember, there is a cost of ownership. Here are some rules of the road:

* Programs consume memory and CPU/GPU cycles
* Electricity costs money: all programming languages are not made equal 
* Programs, dependencies, data, and related files, consume disk space.
* Hardware has a life expectancy, and redundancies are necessary.

*Useful study:*
<https://stefanos1316.github.io/my_curriculum_vitae/GKS17.pdf>. 

# What Options are Available in 2020?

There are currently hundreds of programming languages in existence. Thankfully many of them can be rejected for any project that isn't for funsies without effort. For example, no one will ever use Hexagony(<https://esolangs.org/wiki/Hexagony>) in production. The human race is thankful for that. But - even after widdling down the choices there's still probably 50 or so that are feasible for most modern needs.

## Procedural Programming Languages
| Language | Imperitive | Object Oriented | Systems Programming|
| -       | :-: | :-: | :-: |
| [C](https://devdocs.io/c/)       | X |   | X |
| [Fortran](https://devdocs.io/gnu_fortran~4/) | X |   |   |
| Cobol   | X | X | X |

**Please don't ever use Cobol.**

## Object Oriented Programming Languages
| Language | GC | Compiled | Staticly Typed | Parallelism | Difficulty |
| -        | :-: | :-: | :-: | :-: | :-: |
| [C++](https://www.cplusplus.com/doc/)      |   | X | X | X | 3 |
| [C#](https://docs.microsoft.com/en-us/dotnet/csharp/)       | X | X | X | X | 2 |
| [PHP](https://www.php.net/docs.php)      | X |   |   | X | 2 |
| [Java](https://docs.oracle.com/en/java/)     | X | X | X | X | 2 |

** Difficulty is subjective, ? means I've never used it **

## Systems Programming Languages
| Language | Memory Safety | Imperitive | FP | OOP | Parallelism | Difficulty |
| -       | :-: | :-: | :-: | :-: | :-: | :-: |
| [C](https://devdocs.io/c/)       |   | X |   |   | X | 3 |
| [C++](https://www.cplusplus.com/doc/)     |   | X |   | X | X | 3 |
| [Rust](https://www.rust-lang.org/learn)    | X | X | X | X | X | 3 |
| [Go](https://golang.org/doc/)      | X | X | X | X | X | 1 |
| [Nim](https://nim-lang.org/documentation.html)     | X | X | X | X | X | 1 |
| [Swift](https://swift.org/documentation/)   | X | X | X | X | X | ? |
| [D](https://dlang.org/documentation.html)       |   | X | X | X | X | ? |

** Difficulty is subjective, ? means I've never used it **
## JVM Based languages
| Language    | FP | OOP | Parallelism | Difficulty |
| -           | :-: | :-: | :-: | :-: |
| [Java](https://docs.oracle.com/en/java/)       |   | X | X | 2 |
| [Scala](https://docs.scala-lang.org/)       | X | X | X | 3 |
| [Clojure](https://clojuredocs.org/)     | X | X | X | ? |
| [Groovy](https://groovy-lang.org/documentation.html)      | X | X | X | ? |
| [Jython](https://www.jython.org/)      | X | X | X | ? |
| [Kotlin](https://kotlinlang.org/docs/reference/)      | X | X | X | ? |

## Functional Programming Languages
| Language           | FP | OOP | Difficulty |
| -                  | :-: | :-: | :-: |
| [Julia](https://julialang.org/) | X |   | X |
| [Scala](https://docs.scala-lang.org/)              | X | X | 3 |
| [Haskell](https://www.haskell.org/documentation/)            | X |   | ? |
| [Clean](https://clean.cs.ru.nl/Clean)              | X |   | ? |
| [F#](https://fsharp.org/about/)                 | X | X | ? |
| [Erlang](https://www.erlang.org/docs)             | X | X | ? |
| [OCaml](https://ocaml.org/docs/)           | X | X | ? |
| [Lisp](https://lisp-lang.org/)      | X | X | ? |

## Data Science Languages
| Language           | FP | OOP | Static Typing | Compiled | Parallelism | Difficulty |
| -                  | :-: | :-: | :-: | :-: | :-: | :-: |
| [Julia](https://julialang.org/)              | X |   | X | X | X | 1 |
| [Python](https://docs.python.org/3/)             | X | X | X |   |   | 1 |
| [R](https://www.r-project.org/other-docs.html)                  | X | X | X |   | X | 1 |
| [Matlab](https://www.mathworks.com/help/matlab/language-fundamentals.html)/[Octave](https://www.gnu.org/software/octave/index)      | X | X | X |   | X | 1 |

## Others
| Language           | FP | OOP | Static Typing | Compiled | Parallelism | 
| -                  | :-: | :-: | :-: | :-: | :-: |
| [Ruby](https://www.ruby-lang.org/en/documentation/)               | X | X | X | X | X |
| [Perl](https://perldoc.perl.org/)               | X | X | X | X | X |
| Visual Basic       | D | O | N | ' | T |
