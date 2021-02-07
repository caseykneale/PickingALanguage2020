# Overview
So I did a very basic analysis of the majority of programming languages I've used, and their "objective" properties. 
The analysis done here is pretty bottom of the barrel, IE: we're talking fresher level stuff... The results are interesting nonetheless. The point here is,
what trends "naturally" fall out of language properties, and can I learn anything from them. This is mostly meant to serve as a basis for these kinds of adventures going forward. So each year or so I'll add on more and more... 

# Summary


# Results
## Raw Principal Component Analysis
![pca](https://raw.githubusercontent.com/caseykneale/PickingALanguage2020/main/fieldnotes/analysis/langs.png)

## Knee Jerk Reaction
This is a pretty sensible grouping of languages. Some of it is a little weird, but we did reduce 20 dimensions into 2, and who says any of this is linear.

## Quickly Make Sense of It
![pca](https://raw.githubusercontent.com/caseykneale/PickingALanguage2020/main/fieldnotes/analysis/langpca_lbled.png)

### Interesting Observations
 - A lot of the enterprise languages are at the periphery of the scores plot. This is weird because it implies they are kind of "extreme" from some notion of multipurpose. One would think that languages supporting more paradigms, or mishmashing features, would be more adopted to solve broader problems. 
 
### Some observations lending to consistancy. 
 - C&C++, Java&C#, R&Matlab are right next to each other, etc.
 - There's a pretty clean divide of multiparadigm languages between FP and OOP languages. 
 - Go is closer to C than it is to C++
 - Rust's nearest neighbors are: haskell, F#, C++, and Go. That definitely feels right.
 - Julia being a LISP is close to haskell.
 - The worst programming languages I've ever used are grouped together surprisingly nicely...
 
### Oddities
 - Wouldn't it be nice if C and C++ were primarily OOP languages? Then the left handside of the plot would make more sense. In the past C++ was more strongly OOP than anything else, but overtime it has been adapted to be more paradigm. Sure object oriented C exists, but C, in and of itself is just it's own thing. 
 - Python is smack dab in the middle of special purpose languages and FP languages despite being multiparadigm! This makes sense though... Even though Python is multiparadigm it's OOP is mostly a facade. Most of what python actually does is outright hacky. It makes sense that it's nearer to JavaScript than it is to say C++ or Java despite being a general purpose language and largely touted as one on the internet... Just funny to see it shake out this way...

## My Preferences
![pca](https://raw.githubusercontent.com/caseykneale/PickingALanguage2020/main/fieldnotes/analysis/MyPreferences.png)

*Languages I enjoy are in orange and languages I dislike are in black. The thicker/more consistant the line the stronger feeling I have about them.*
 - There is this sweet spot where my favorite languages have emerged. Right around Haskell, there's Julia and Rust.
 - Languages I don't enjoy are not really to the left of the languages I do enjoy. So it seems I like general purpose tools. That makes sense... I am typically a solodeveloper on a team bringing something from nothing to conception.
 - The upper right hand quadrant has some terrible languages. Some of which you probably could not pay me to write code in. ArrayBasic being the absolute worst programming language I have ever seen. People complain about R, vs Python, etc, my goodness, they have no idea how bad it really gets.

## Criteria
If you look at the [langset.csv](https://github.com/caseykneale/PickingALanguage2020/blob/main/fieldnotes/analysis/langset.csv) file you'll see a small dataset with roughly 
20 features. Most of these are binary encoded or (true = 1, false = 0). Some however are slightly quantitative(based on exemplary Hello World programs in each language).

## Binary Features
- Has REPL : Does the language provide a REPL?
- Is Compiled : is the language compiled or interpretted? 
- Has GOTO : Does it have assembly like GoTo statements
- JIT : Is it Just-In-Time compiled?
- List comprehensions : Does it support sugar/monads for handling lists easily?
- Generics : Can generic types be used?
- GC : Does it have a Garbage Collector?
- OOP : Does it support object oriented programming paradigms. **Note:** not "is this a pure OOP"
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

### Known Caveats
 - Shouldn't really use PCA for binary variables, or count data. 
 - Features are not perfect, I drummed most of them up in a half hour, but had some help from Rust discord with a few!
