So I did a very basic analysis of the majority of programming languages I've used, and their "objective" properties. 
The analysis done here is pretty bottom of the barrel, IE: we're talking fresher level stuff... The results are interesting nonetheless. The point here is,
what trends "naturally" fall out of language properties, and can I learn anything from them. Also establish a basis for this going forward in life now, sounds good. 
So each year or so I'll add on more and more... Again there are obvious problems with this analysis', but I don't have a lot of time to dedicate to this in general.

## Criteria
If you look at the [langset.csv](https://github.com/caseykneale/PickingALanguage2020/blob/main/fieldnotes/analysis/langset.csv) file you'll see a small dataset with roughly 
20 features. Most of these are binary encoded or (true = 1, false = 0). Some however are slightly quantitative(based on exemplary Hello World programs in each language).

## Binary Features
Has REPL
- Is Compiled : is the language compiled or interpretted? 
- Has GOTO : Does it have assembly like GoTo statements
- JIT : Is it Just-In-Time compiled?
- List comprehensions : Does it support sugar/monads for handling lists easily?
- Generics : Can generic types be used?
- GC : Does it have a Garbage Collector?
- OOP : Does it support object oriented programming paradigms. Note: not "is this a pure OOP"
- FP 	: Does it support functional programming paradigms. I think the criteria here is, is there first class support for fns.
- Metaprogramming : Can you metaprogram in it?
- Cross Platform 	: Can you target other OS's/platforms
- Safety 	: Is memory safety a feature.
- Strong 	: Is it strongly typed.
- General Purpose : Is this language "general purpose" or a "niche". IE: Excel is a niche, so is mathematica. 
- Package Manager : Does the language come with, or commonly support package managers?

## Quantitative Features
- special chars In HelloWorld : Number of special characters in a hello world application.
- avg. specials Per line In HelloWorld 	: Average special characters per line.
- white space In HelloWorld : How much white space is there? (tabs and spaces)
- Hello World LOC : How many lines of code is it
- Hello World Chars : How many characters did it take to write an idiomatic Hello World program

## Principal Component Analysis
![pca](https://raw.githubusercontent.com/caseykneale/PickingALanguage2020/main/fieldnotes/analysis/langs.png)

What I find most interesting here is 
 - There is this sweet spot where my favorite languages have emerged. Right around Haskell, there's Julia and Rust.
 - So languages following these kinds of paradigms are probably where I end up
 - A lot of the enterprise languages are at the periphery of the scores plot. This is weird because it implies they are kind of "extreme" from some notion of multipurpose. One would think that languages supporting more paradigms, or mishmashing features, would be more adopted to solve broader problems. 
 
Some observations lending to consistancy. 
 - C&C++, Java&C#, Python&R&Matlab are right next to each other, etc.
 - There's a pretty clean divide between FP and OOP languages 
 - Go is closer to C than it is to C++
 - Rust's nearest neighbors are: haskell, F#, C++, and Go. This feels right.
 - Julia being a LISP is close to haskell.
 - The worst programming languages I've ever seen are grouped well together.The worst is on the periphery of that group - so, that makes sense...
 
Weird things
 - Fortran and VB are really close to one another. It feels wrong... But that's what these features say anyways...

### Caveats
 - Shouldn't really use PCA for binary variables, or count data. 
 - Features are not perfect, I drummed most of them up in a half hour, but had some help from Rust discord with a few!
