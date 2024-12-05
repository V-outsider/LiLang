# LiLang
Study purposes created language to deal with list like objects

Algo:
1) antlr4 -Dlanguage=Go LiLang.g4 -o ./generated # If not exists

2)llc -filetype=obj output.ll -o output.o
3) # Check if output.ll exists and is valid
   cat output.ll
   # Use verbose output for llc
   llc -mtriple=arm64-apple-darwin24.1.0 -filetype=obj output.ll -o output.o
   # Check the file type
   clang -arch arm64 output.o -o output

4) # Use verbose output for clang linking
   clang -v output.o -o output



print func 
clang -c print.c -o print.o

clang output.o print.o -o output
clang output.o -o output -lc