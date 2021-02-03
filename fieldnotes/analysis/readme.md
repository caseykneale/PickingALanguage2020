So I did a very basic analysis of the majority of programming languages I've used, and their "objective" properties. 
The analysis done here is pretty bottom of the barrel, IE: we're talking fresher level stuff... The results are interesting nonetheless. The point here is,
what trends "naturally" fall out of language properties, and can I learn anything from them. Also establish a basis for this going forward in life now, sounds good. 
So each year or so I'll add on more and more... Again there are obvious problems with this analysis', but I don't have a lot of time to dedicate to this in general.

## Criteria
If you look at the (langset.csv)[https://github.com/caseykneale/PickingALanguage2020/blob/main/fieldnotes/analysis/langset.csv] file you'll see a small dataset with roughly 
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
