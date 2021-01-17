# Picking a Programming Language at the End of 2020? (Opinion)

I'm an unabashed data geek working in industry who specializes in *dealing with* chemical data. I wear hats that span the set from analytical chemist, data scientist, and software developer. With a little bit of everything in between. Over the years I've written hundreds of thousands of lines of code. Some of it quite good, some of it so bad that I hope there are no traces of it on the internet. I do this both as a hobby and as a means of putting food on the table.

I have worked alongside high-flying data scientists - you know the geniuses who can scrap together a new algorithm using mathematics that wasn't known to exist and raise them to production on time. I have also worked alongside low-level systems programmers - the type of people who giggle when their code pushes servers to literally smoke and catch fire. For some developers/scientists, like myself, the line between these two groups gets a bit blurry. I figured I'd offer a mediated opinion on scouting programming languages. 

Every year I scout programming languages to see what advances have been made, and if I should make a change. I wrote this post because I often see or hear about schisms in teams formed from disparities in tooling. Those kinds of divergences can cause fractures in deliverables/products and introduce pain points in large codebases. I also tend to see a lot of hype pouring from influencer types that really muddy the waters for people looking for straightforward suggestions. Thankfully beneath the curtain of advertising and self-promotion, the water really isn't all that murky. Choosing tools is based on technological merit, but also social/business needs. 

# High-level Decision Process

This section serves to detail some items to keep in mind as you drill down the cyclic process of settling on a programming language for a project. If you just interested in learning something new you can skip most of this section. 

## Assess Resources

**Technical ability:** 
How savvy is the team with computer science/development? Can they learn a new language on their own? How long might that take given the complexity of a tool? 

**Incentive:**
Is there an incentive to offer a **change** to be put in place? Is continuous learning a cultural norm? Are there selling points to make a change a "net positive" for the people involved in the project?

**Safety:** 
Is there a culture of psychological safety and merit-based reward systems? If not, trying new things can scare people away or make them feel insecure. 

**Risk:**
Whats the risk of picking "the wrong language"? Is it a total failure for a product, or a few weeks delay in a milestone? How quickly might that realization be made?

## Define the problem statement

**Pick Target Platform(s):** 
Desktop/Embedded/Cloud/On Premises/etc. What do the users need? What do they want? Make a survey, do some research, put yourself in their shoes, and talk to marketing. 

**Identify Sources of Novelty:**
Is your killer feature a huge distributed computing web framework behind a super clean easy to use front end? Might need to look at languages supporting stable parallelism, and their ecosystems. Maybe you need stability, your competitor's offerings are very glitchy, look into languages with good memory safety. Maybe you need a highly interactive GUI - what GUI API's are accessible to the languages you're interested in? Maybe you just need to hoist some fancy new algorithm to the cloud behind an API, does your language have established reliable tools that make this possible (do you need an expert, or another engineer)?

**Identify Technical Need:**
Items to keep in mind: performance, mathematics, networking, parallelism, beauty, maintainability(short-term/long-term), support, etc.

## Estimate Cost and Benefits
Draw up a rough budget for closed source languages (matlab, etc), for training/recruiting staff, licenses needed to use 3rd party tools, etc. Consider a team of 10 developers each needing a 5,000USD$ annual license, how does this compare to losing two extra weeks due to training them with an alternative approach? If you are using on premises hard-ware remember, there is a cost of ownership. Here are some rules of the road:

* Programs consume memory and clock cycles
* Electricity costs money programming languages are not made equal 
* Programs, dependencies, data, and related files all consume disk space.
* Hardware has a life expectancy, and redundancies are necessary.
* Project for growth/success. Better to face worst case upfront and have something left over.

Do these even matter for your project? Usually not for me :). 

*Useful study:*
<https://stefanos1316.github.io/my_curriculum_vitae/GKS17.pdf>. 

# What Options are Available in 2020?

There are currently hundreds of programming languages in existence. Thankfully many languages can be rejected for any project that isn't just for funsies. For example, no one will ever use [Hexagony](<https://esolangs.org/wiki/Hexagony>) in production. The human race is thankful for that. But - even after widdling down the choices there's still probably 50ish languages that are reasonable for modern needs. Below are some categorized lists of languages to help ease someone elses search. Not all languages are included here, just major contenders that I'm aware of. The line of thinking presented here, does extend to other languages though. After this section I offer my preferences and why I like them.

**Please note, in some tables you will see a "Difficulty" score. This is subjective but included to give the reader a "feel" for my experience using a variety of languages. The score is based on my perception of my ability to learn enough syntax and DevOps to create deployable projects. A score of 1 means it's easy, 3 means it's about as had as it gets, and a `?` score means I've never used the language**

## Procedural Programming Languages
| Language | Imperitive | Object Oriented | Systems Programming|
| -       | :-: | :-: | :-: |
| [C](https://devdocs.io/c/)       | X |   | X |
| [Fortran](https://devdocs.io/gnu_fortran~4/) | X |   |   |

## Object Oriented Programming Languages
| Language | GC | Compiled | Staticly Typed | Parallelism | Difficulty |
| -        | :-: | :-: | :-: | :-: | :-: |
| [C++](https://www.cplusplus.com/doc/)      |   | X | X | X | 3 |
| [C#](https://docs.microsoft.com/en-us/dotnet/csharp/)       | X | X | X | X | 2 |
| [PHP](https://www.php.net/docs.php)      | X |   |   | X | 2 |
| [Java](https://docs.oracle.com/en/java/)     | X | X | X | X | 2 |

## Systems Programming Languages
| Language | Embedded Systems | Memory Safety | Imperitive | FP | OOP | Parallelism | Adoption | Difficulty |
| -       | :-: | :-: | :-: | :-: | :-: | :-: | :-: | :-: |
| [C](https://devdocs.io/c/)       | X |   | X |   |   | X | Moderate | 3 |
| [C++](https://www.cplusplus.com/doc/)     | X |   | X |   | X | X | High | 3 |
| [Rust](https://www.rust-lang.org/learn)    | X | X | X | X | X | X | Low | 2 |
| [Go](https://golang.org/doc/)      | X | X | X | X | X | X | Moderate | 1 |
| [Nim](https://nim-lang.org/documentation.html)     | X | X | X | X | X | X | Low | 2 |
| [Swift](https://swift.org/documentation/)   | X | X | X | X | X | X | Moderate | ? |
| [D](https://dlang.org/documentation.html)       | X |  | X | X | X | X | Low | ? |

## JVM Based languages
| Language    | FP | OOP | Parallelism | Adoption | Difficulty |
| -           | :-: | :-: | :-: | :-: | :-: |
| [Java](https://docs.oracle.com/en/java/)       |   | X | X | High | 2 |
| [Scala](https://docs.scala-lang.org/)       | X | X | X | Moderate | 3 |
| [Clojure](https://clojuredocs.org/)     | X | X | X | Moderate  | ?|
| [Groovy](https://groovy-lang.org/documentation.html)      | X | X | X | ? |? |
| [Jython](https://www.jython.org/)      | X | X | X | Niche | ? |
| [Kotlin](https://kotlinlang.org/docs/reference/)      | X | X | X | ? | ? |

## Functional Programming Languages
| Language           | FP | OOP | Difficulty |
| -                  | :-: | :-: | :-: |
| [Julia](https://julialang.org/) | X |   | 1 |
| [Scala](https://docs.scala-lang.org/)              | X | X | 3 |
| [Haskell](https://www.haskell.org/documentation/)            | X |   | 3 |
| [Clean](https://clean.cs.ru.nl/Clean)              | X |   | ? |
| [F#](https://fsharp.org/about/)                 | X | X | 2 |
| [Erlang](https://www.erlang.org/docs)             | X | X | ? |
| [OCaml](https://ocaml.org/docs/)           | X | X | 2 |
| [CLisp](https://lisp-lang.org/)      | X | X | ? |

## Data Science Languages
| Language           | FP | OOP | Static Typing | Compiled | Parallelism | Adoption | Ecosystem | Difficulty |
| -                  | :-: | :-: | :-: | :-: | :-: | :-: | :-: | :-: |
| [Julia](https://julialang.org/)              | X |   | X | X | X | Low | Young | 1 |
| [Python](https://docs.python.org/3/)             | X | X |   |   |   | High | Senior | 1 |
| [R](https://www.r-project.org/other-docs.html)                  | X | X |   |   | X | High | Senior | 2 |
| [MATLAB](https://www.mathworks.com/help/matlab/language-fundamentals.html)/[Octave](https://www.gnu.org/software/octave/index)      | X | X |   |   | X | Moderate/Low | Senior | 1 |

## Others
| Language           | FP | OOP | Static Typing | Compiled | Parallelism | 
| -                  | :-: | :-: | :-: | :-: | :-: |
| [Ruby](https://www.ruby-lang.org/en/documentation/)               | X | X | X |   | X |
| [Perl](https://perldoc.perl.org/)               | X | X | X |   | X |
| Cobol   | P | L | EA | S | E |
| Visual Basic       | D | O | N | ' | T |
| Pascal       | N | O | P | E | ! |

## My Suggestions 

### Data
For my needs I find Julia typically fits the bill. That said - every project is different. Sometimes I run to R for vetted statistical methods. But, I mostly find myself creating new algorithms. So I use Julia for fast prototyping, performance tuning, and quickly turning prototypes into fast deliverables. Keyword - fast. Also - the community is awesome! 

Python is nice, but it has some deficits for what I do. Your mileage will likely vary here; I'm in a weird position. Python has great tools for routinely horking data, building pipelines, and deploying tools. It also has a massive user base for better and for worse. One benefit to these kinds of languages(Julia, Python, and R), is that they tend to have good interoperability with a suite of external tools to build productive stacks. Some stacks are easier to link than others, but you'll often find you can leverage tools across languages if you are careful and creative.

| Primary Need | Secondary Need | Tertiary Need | Platform                    | Okay With  | Language |  
|     :-:      | :-:            | :-:           | :-:                         | :-:        | :-:      |
| Mathematics  | Speed          | Easy          | Mac, Win, Linux, Web, Cloud | JIT/JAOT        | Julia    |
| Statistics   | Easy           | Free APIs     | Mac, Win, Linux, Web, Cloud | Needing C++/APIs | R    |
| Free APIs    | Easy           | Mathematics   | Mac, Win, Linux, Web, Cloud | Needing C++/APIs | Python  |
| A Specific Package    | Mathematics   |  Easy  | Mac, Win, Linux | Proprietary Language | MATLAB  |
| Spark | Custom Data Plumbing  | Parallelism  |  Mac, Win, Linux, Web        | Complexity  | Scala    |

### Applications/Systems Programming
I've been all over the place with the languages in this section. To be honest, I'm sort of on the fence with *most* of these languages. Historically I've ran to C++ when I need this kind of tool - gold standard I guess. For a few years of my life I spent a lot of time using C# for game development, was fun, but the codebases got ugly fast. Lately I've been looking to upgrade my toolchest so I tried some more modern languages.

Some part of me really enjoys Go: I found it safe enough, fast enough, and easy enough for me to quickly create working binaries. Like in a week I was able to write a super basic, but fast, low level GUI library. Is Go a "great" language? No - not really, it's actually kind of *bad*. In the ways that it's bad are a little zen, but also can be frusterating. Like no generic types? No enumerations? But in the same breath - simplistic code without the cognitive overhead is extremely rare for a systems type language. 

Most recently I've ventured into using Rust. It's a seriously alluring language that removes a lot of the ills from C++. It's also pretty easy to write - if you have a background for it. If systems programming isn't in your blood I really reccommend spending at least 2 weeks in C++ and 2 weeks in Haskell before trying Rust. It's surprising how well learning basic Haskell can prime you for Rust. Especially when you consider the languages are not even close to similar in intent. If I were to suggest a language to learn out of all of these, and learn really well, it'd probably be Rust. But, I'm still new to it.

| Primary Need | Secondary Need | Tertiary Need | Platform                    | Okay With   | Language |  
|     :-:      | :-:            | :-:           | :-:                         | :-:         | :-:      |
| Easy         | Speed          | Parallelism   | Mac, Win, Linux, Web, Cloud | No Generics + Static Types | Go       |
| Safety       | Speed          | Binaries      | Mac, Win, Linux, Web, Cloud | Complexity  | Rust     |
| Speed        | Low Level Access | Binaries    | Mac, Win, Linux             | Unsafe Code | C      |
| Low Level Access | Speed      | Easier than C | Mac, Win, Linux             | OOP         | C++      |
| Desktop      | Stability      | GUI           | Win(best), Mac, Linux       | GC          | C#       |

### JVM
I had a falling out with JVM based languages. Java used to be my primary language. Once I became exposed to other languages and crossplatform support became a language requirement - JVM offered me less and less. There are some good things about Java, it's relatively easy to get a group of people to write idiomatic Java and have good coherence across projects. Back in the day, it was pretty fast (relatively speaking), but that's just not the case anymore. 

Scala is neat, some great tools are written in it. The second scariest code I've ever seen was written in Scala. That takes the cake for me - Scala is a little too "liberating", and as such, you can end up with scary syntax between teammates, projects, companies, etc. I know people who have left jobs over Scala... It's a little rough.

Don't even ask about the scariest code I've ever seen, all I can say is - it wasn't my own. 

| Primary Need | Secondary Need | Tertiary Need | Platform                    | Okay With   | Language |  
|     :-:      | :-:            | :-:           | :-:                         | :-:         | :-:      |
| JVM      | Maturity           | Libraries     | Mac, Win, Linux, Web        | GC          | Java     |
| JVM | Functional Programming  | Parallelism   | Mac, Win, Linux, Web        | Complexity  | Scala    |

### Fun
Below are some languages I reccommend trying out if you are looking to learn something new, and have some fun while you're at it. They will open your eyes to whats out there. Some brief comments on them are given below. 

*Julia* - looks, feels, and runs like the future of applied mathematics and data science. It's also downright satisfying to write.

*Haskell* - total paradigm shift from imperitive programming, as challenging as you want it to be. There's something beautiful to be found in Haskell, but - it's not a language you're likely to get paid to use unless you are really good at it.

*Rust* - It's a treat, I mean, it's satisfying to write C++ like code without being kept up at night thinking about pointers before pushing to production. It also has some - good ideas in it, some truly fresh takes, and helpful error messages. Definitely worth learning and kind of healing compared to the C languages.

| Primary Desire | Secondary Desire | Tertiary Desire | Platform                    | Okay With   | Language |  
|     :-:      | :-:            | :-:           | :-:                         | :-:         | :-:      |
| Expressive well-reasoned syntax | Lispyness done right! | Great community | Mac, Win, Linux, Web        | JIT  | Julia    |
| Pure Functional Programming   | Category Theory | Recursion without penalty | Mac, Win, Linux, Web        | Mathematical Reasoning | Haskell     |
| Modern Forward Thinking | Imperitive FP/OOP  | Great community   | Mac, Win, Linux, Web        | Picky Compiler  | Rust    |

# Wrapping up
Pick a language based on your need. My suggestions likely won't suite your needs. If you pick languages, or any tool, based on blog posts, or social media hype, you're probably going to get burned or end up spreading misinformation. Discern your needs first, account for resources available, then weigh the pro's and con's. It involves research, and trying things out - what works for you? 

Technological idealism does not exist in industry nor academia. If it did - there would be a single programming language - right? Not hundreds. My primary choices are based on my ability to easily write stable and performant programs in them. My secondary languages are based on "do I have time to do this myself?". Mostly because I commonly find myself in two positions: innovating, and quickly iterating.

Word of warning - especially my fellow data people. I encourage everyone to know enough about their field and the tools they use. Do not willfully fall victim to choosing technology based on inability to do the work without a specific tool. Be curious, look under the hood, try to build one yourself, engage in open source to learn as much as you can. Make cool things. If that doesn't work out, have has much fun as you can along the way!
