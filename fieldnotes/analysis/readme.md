So I did a very basic analysis of the majority of programming languages I've used, and their "objective" properties. 
The analysis done here is pretty bottom of the barrel, IE: we're talking fresher level stuff... The results are interesting nonetheless. The point here is,
what trends "naturally" fall out of language properties, and can I learn anything from them. Also establish a basis for this going forward in life now, sounds good. 
So each year or so I'll add on more and more... Again there are obvious problems with this analysis', but I don't have a lot of time to dedicate to this in general.

## Criteria
If you look at the (langset.csv)[https://github.com/caseykneale/PickingALanguage2020/blob/main/fieldnotes/analysis/langset.csv] file you'll see a small dataset with roughly 
20 features. Most of these are binary encoded or (true = 1, false = 0). Some however are slightly quantitative(based on exemplary Hello World programs in each language).

## Binary Features
Has REPL
- Is Compiled 	
- Has GOTO 	
- JIT 	
- List comprehensions 	
- Generics 	
- GC 	
- OOP 	
- FP 	
- Metaprogramming 	
- Cross Platform 	
- Safety 	
- Strong 	
- General Purpose 	
- Package Manager

## Quantitative Features
- special chars In HelloWorld 	
- avg. specials Per line In HelloWorld 	
- white space In HelloWorld 	
- Hello World LOC 	
- Hello World Chars
