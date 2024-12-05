target triple = "arm64-apple-darwin24.1.0"
target datalayout = "e-m:o-i64:64-i128:128-n32:64-S128"

@x = global i32 0
@z = global i32 0
@printf.str = global [3 x i8] c"%d\0A"

define i32 @main() {
entry:
	%0 = add i32 10, 20
	store i32 %0, i32* @x
	store i32 4, i32* @z
	%1 = load i32, i32* @z
	%2 = call i32 @sum(i32 %1, i32 4)
	store i32 %2, i32* @x
	%3 = load i32, i32* @x
	%4 = getelementptr [3 x i8], [3 x i8]* @printf.str, i64 0, i64 0
	%5 = call i32 (i8*, ...) @printf(i8* %4, i32 %3)
	%6 = load i32, i32* @x
	ret i32 %6
}

define i32 @sum(i32 %a, i32 %b) {
entry:
	%0 = add i32 %a, %b
	ret i32 %0
}

declare i32 @printf(i8* %0, ...)
